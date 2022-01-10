package role

import (
	"github.com/sakuraapp/shared/model"
	"github.com/sakuraapp/shared/resource/permission"
)

type Id int32

const (
	MEMBER Id = 0
	MANAGER = 1
	HOST = 2
)

type Role struct {
	id Id
	permissions permission.Permission
	order int
}

func (r *Role) Id() Id {
	return r.id
}

func (r *Role) Permissions() permission.Permission {
	return r.permissions
}

func (r *Role) Order() int {
	return r.order
}

var roles = map[Id]*Role{
	MEMBER: {
		id: MEMBER,
		permissions: permission.QUEUE_ADD,
		order: 0,
	},
	MANAGER: {
		id: MANAGER,
		permissions: permission.QUEUE_ADD | permission.QUEUE_EDIT | permission.VIDEO_REMOTE | permission.KICK_MEMBERS,
		order: 1,
	},
	HOST: {
		id: HOST,
		permissions: permission.ALL,
		order: 2,
	},
}

func GetRole(id Id) *Role {
	return roles[id]
}

type Manager struct {
	roles map[Id]bool
	permissions permission.Permission
}

func (m *Manager) recalculatePerms() {
	m.permissions = 0

	for id := range m.roles {
		m.permissions |= GetRole(id).Permissions()
	}
}

func (m *Manager) Add(id Id) {
	m.roles[id] = true
	m.permissions |= GetRole(id).Permissions()
}

func (m *Manager) Remove(id Id) {
	delete(m.roles, id)
	m.recalculatePerms()
}

func (m *Manager) Has(id Id) bool {
	return m.roles[id]
}

func (m *Manager) Max() *Role {
	max := new(Role)

	for id := range m.roles {
		role := GetRole(id)

		if max == nil || role.Order() > max.Order() {
			max = role
		}
	}

	return max
}

func (m *Manager) Slice() []Id {
	roleList := make([]Id, len(m.roles))

	for id := range m.roles {
		roleList = append(roleList, id)
	}

	return roleList
}

func (m *Manager) Permissions() permission.Permission {
	return m.permissions
}

func (m *Manager) HasPermission(perm permission.Permission) bool {
	return m.permissions.Has(perm)
}

func NewManager() *Manager {
	return &Manager{
		roles: map[Id]bool{},
		permissions: 0,
	}
}

func BuildManager(userRoles []model.UserRole, isRoomOwner bool) *Manager {
	m := NewManager()
	m.Add(MEMBER)

	for _, userRole := range userRoles {
		m.Add(userRole.RoleId)
	}

	if isRoomOwner {
		m.Add(HOST)
	}

	return m
}