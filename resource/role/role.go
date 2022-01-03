package role

import "github.com/sakuraapp/shared/resource/permission"

type Role = permission.Permission

const (
	MEMBER Role = permission.QUEUE_ADD
	MANAGER = MEMBER | permission.QUEUE_EDIT | permission.VIDEO_REMOTE
	HOST = permission.ALL
)