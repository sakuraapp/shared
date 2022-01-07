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

type RoleManager struct {
	roles map[Role]bool
	permissions permission.Permission
}

func (m *RoleManager) recalculatePerms() {
	m.permissions = 0

	for role := range m.roles {
		m.permissions |= role.Permissions()
	}
}

func (m *RoleManager) Add(role Role) {
	m.roles[role] = true
	m.permissions |= role.Permissions()
}

func (m *RoleManager) Remove(role Role) {
	delete(m.roles, role)
	m.recalculatePerms()
}

func (m *RoleManager) Has(role Role) bool {
	return m.roles[role]
}

func (m *RoleManager) Roles() []Role {
	roles := make([]Role, len(m.roles))

	for role := range m.roles {
		roles = append(roles, role)
	}

	return roles
}

func (m *RoleManager) Permissions() permission.Permission {
	return m.permissions
}

func (m *RoleManager) HasPermission(perm permission.Permission) bool {
	return m.permissions.Has(perm)
}

func NewRoleManager() *RoleManager {
	return &RoleManager{
		roles: map[Role]bool{},
		permissions: 0,
	}
}