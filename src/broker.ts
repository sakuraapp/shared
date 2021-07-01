import { EventEmitter } from 'events'
import { encode, decode } from '@msgpack/msgpack'
import Redis, { Redis as RedisInstance } from 'ioredis'
import { v4 } from 'uuid'

export type BrokerHandler<T> = (data: T) => void

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
    public readonly messages = new EventEmitter()

    protected readonly opts: BrokerOptions
    protected pub: RedisInstance
    protected sub: RedisInstance

    constructor(opts: BrokerOptions) {
        this.opts = opts
    }

    private serialize<T>(message: BrokerMessage<T>): Buffer {
        return Buffer.from(encode(message))
    }

    private deserialize<T>(data: Buffer): BrokerMessage<T> {
        return decode(data) as BrokerMessage<T>
    }

    protected cloneClient(): RedisInstance {
        return new Redis(this.opts.client.options)
    }

    protected initPub(): void {
        let client: RedisInstance = this.opts.client

        if (this.sub) {
            client = this.cloneClient()
        }

        this.pub = client
    }

    protected initSub(): void {
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
            
            this.onMessage(channel, message.d)
        })
    }

    protected onMessage<T>(channel: string, data: T): void {
        this.messages.emit(channel, data)
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

    async subscribe<T>(channel: string, handler: BrokerHandler<T>) {
        if (!this.sub) {
            this.initSub()
        }

        if (this.messages.listenerCount(channel) === 0) {
            await this.sub.subscribe(channel)
        }

        this.messages.on(channel, handler)
    }

    async unsubscribe<T>(channel: string, handler: BrokerHandler<T>) {
        this.messages.off(channel, handler)

        if (this.messages.listenerCount(channel) === 0) {
            await this.sub.unsubscribe(channel)
        }
    }
}
