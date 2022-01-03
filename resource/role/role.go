package role

import "github.com/sakuraapp/shared/resource/permission"

type Role int32

const (
	MEMBER Role = 0
	MANAGER = 1
	HOST = 2
)

var permissions = map[Role]permission.Permission{
	MEMBER: permission.QUEUE_ADD,
	MANAGER: permission.QUEUE_ADD | permission.QUEUE_EDIT | permission.VIDEO_REMOTE,
	HOST: permission.ALL,
}

func (r Role) Permissions() permission.Permission {
	return permissions[r]
}