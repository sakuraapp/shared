import { BaseRepository } from './repositories/base.repository'
import { NameRepository } from './repositories/name.repository'
import { Discriminator } from './entities/discriminator.entity'
import { Name } from './entities/name.entity'
import { User, Credentials, Profile } from './entities/user.entity'
import { Room } from './entities/room.entity'
import { MikroORM, Options } from '@mikro-orm/core'
import { TsMorphMetadataProvider } from '@mikro-orm/reflection'
import { RedisCacheAdapter } from './adapters/redis.adapter'
import {
    Broker,
    BrokerOptions,
    BrokerMessage,
    BrokerHandler,
} from './broker'

export function createDatabase(options: Options): Promise<MikroORM> {
    return MikroORM.init({
        type: 'mongo',
        metadataProvider: TsMorphMetadataProvider,
        entityRepository: BaseRepository,
        ...options,
    })
}

export {
    MikroORM,
    BaseRepository,
    NameRepository,
    Discriminator,
    Name,
    User,
    Credentials,
    Profile,
    Room,
    RedisCacheAdapter,
    Broker,
    BrokerOptions,
    BrokerMessage,
    BrokerHandler,
}
