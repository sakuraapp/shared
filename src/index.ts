import { BaseRepository } from './repositories/base.repository'
import { NameRepository } from './repositories/name.repository'
import { Discriminator } from './entities/discriminator.entity'
import { Name } from './entities/name.entity'
import { User, Credentials, Profile } from './entities/user.entity'
import { Room } from './entities/room.entity'
import { MikroORM, Options } from '@mikro-orm/core'
import { TsMorphMetadataProvider } from '@mikro-orm/reflection'

export function createDatabase(options: Options) {
    return MikroORM.init({
        type: 'mongo',
        metadataProvider: TsMorphMetadataProvider,
        entityRepository: BaseRepository,
        ...options,
    })
}

export {
    BaseRepository,
    NameRepository,
    Discriminator,
    Name,
    User,
    Credentials,
    Profile,
    Room
}
