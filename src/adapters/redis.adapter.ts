import { CacheAdapter } from '@mikro-orm/core'
import { Redis as RedisInstance } from 'ioredis'

const NAMESPACE = 'db_cache'

export interface Options {
    client: RedisInstance
    expiration?: number // expiration in seconds
    namespace?: string
}

export class RedisCacheAdapter implements CacheAdapter {
    private readonly client: RedisInstance
    private readonly expiration: number
    private readonly namespace: string

    constructor(opts: Options) {
        this.client = opts.client
        this.expiration = opts.expiration
        this.namespace = opts.namespace || NAMESPACE
    }

    private serialize<T>(data: T): string {
        return JSON.stringify(data)
    }

    private deserialize<T>(data: string): T {
        return JSON.parse(data)
    }

    private key(name: string): string {
        return `${this.namespace}:${name}`
    }

    async get<T = any>(name: string): Promise<T> {
        const data = await this.client.get(this.key(name))
        let res: T

        if (data) {
            try {
                res = this.deserialize(data)
            } catch(err) {}
        }

        return res
    }

    async set<T = any>(name: string, data: T): Promise<void> {
        const key = this.key(name)

        await this.client.set(key, this.serialize(data))

        if (this.expiration) {
            await this.client.expire(key, this.expiration)
        }
    }

    // Only used by Metadata Cache - not required for Result Cache (but present in interface)
    async clear(): Promise<void> {}
}
