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

type Manager struct {
	roles map[Role]bool
	permissions permission.Permission
}

func (m *Manager) recalculatePerms() {
	m.permissions = 0

	for role := range m.roles {
		m.permissions |= role.Permissions()
	}
}

func (m *Manager) Add(role Role) {
	m.roles[role] = true
	m.permissions |= role.Permissions()
}

func (m *Manager) Remove(role Role) {
	delete(m.roles, role)
	m.recalculatePerms()
}

func (m *Manager) Has(role Role) bool {
	return m.roles[role]
}

func (m *Manager) Roles() []Role {
	roles := make([]Role, len(m.roles))

	for role := range m.roles {
		roles = append(roles, role)
	}

	return roles
}

func (m *Manager) Permissions() permission.Permission {
	return m.permissions
}

func (m *Manager) HasPermission(perm permission.Permission) bool {
	return m.permissions.Has(perm)
}

func NewManager() *Manager {
	return &Manager{
		roles: map[Role]bool{},
		permissions: 0,
	}
}