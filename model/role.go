package model

import "github.com/sakuraapp/shared/resource/role"

type UserRoleId = int32

type UserRole struct {
	Id UserRoleId
	RoleId role.Id
}