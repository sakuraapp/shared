import { EventEmitter } from 'events'
import { encode, decode } from '@msgpack/msgpack'
import Redis, { Redis as RedisInstance } from 'ioredis'
import { v4 } from 'uuid'

export interface BrokerMessage<T = unknown> {
    d: T // data
    s: string // source broker id
}

export interface BrokerOptions {
    client: RedisInstance
    allowOwnMessages?: boolean
}

export class Broker {
    public readonly id = v4()

    public pub: RedisInstance
    public sub: RedisInstance

    public readonly messages = new EventEmitter()

    private readonly opts: BrokerOptions

    constructor(opts: BrokerOptions) {
        this.opts = opts
    }

    private serialize<T>(message: BrokerMessage<T>): Buffer {
        return Buffer.from(encode(message))
    }

    private deserialize<T>(data: Buffer): BrokerMessage<T> {
        return decode(data) as BrokerMessage<T>
    }

    private cloneClient(): RedisInstance {
        return new Redis(this.opts.client.options)
    }

    private initPub(): void {
        let client: RedisInstance = this.opts.client

        if (this.sub) {
            client = this.cloneClient()
        }

        this.pub = client
    }

    private initSub(): void {
        let client: RedisInstance = this.opts.client

        if (this.pub) {
            client = this.cloneClient()
        }

        this.sub = client
        this.sub.on('messageBuffer', (channel, data) => {
            const message: BrokerMessage = this.deserialize(data)

            if (message.s === this.id && !this.opts.allowOwnMessages) {
                return
            }
            
            this.messages.emit(channel, message.d)
        })
    }

    publish<T>(channel: string, data: T): Promise<number> {
        const message: BrokerMessage<T> = {
            d: data,
            s: this.id,
        }

        if (!this.pub) {
            this.initPub()
        }

        return this.pub.send_command('PUBLISH', channel, this.serialize(message))
    }

    async subscribe<T>(channel: string, handler: (data: T) => void) {
        if (!this.sub) {
            this.initSub()
        }

        if (this.messages.listenerCount(channel) === 0) {
            await this.sub.subscribe(channel)
        }

        this.messages.on(channel, handler)
    }

    async unsubscribe<T>(channel: string, handler: (data: T) => void) {
        this.messages.off(channel, handler)

        if (this.messages.listenerCount(channel) === 0) {
            await this.sub.unsubscribe(channel)
        }
    }
}
