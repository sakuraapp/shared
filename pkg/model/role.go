package model

import (
	"github.com/sakuraapp/shared/pkg/resource/role"
)

type UserRoleId = int32

type UserRole struct {
	Id     UserRoleId
	UserId UserId
	RoomId RoomId
	RoleId role.Id
}

func BuildRoleManager(userRoles []UserRole) *role.Manager {
	m := role.NewManager()
	m.Add(role.MEMBER)

	for _, userRole := range userRoles {
		m.Add(userRole.RoleId)
	}

	return m
}