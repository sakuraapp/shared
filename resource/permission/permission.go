package permission

type Permission int32

const (
	QUEUE_ADD Permission = 1
	QUEUE_EDIT = 2
	VIDEO_REMOTE = 4
	KICK_MEMBERS = 8
	MANAGE_ROLES = 16
	MANAGE_ROOM = 32
	ALL = (1 << 6) - 1 // (2 ^ no. of roles) - 1
)

func (p Permission) Has(perm Permission) bool {
	return (p & perm) > 0
}

func (p Permission) Add(perms ...Permission) Permission {
	for _, perm := range perms {
		p |= perm
	}

	 return p
}

func (p Permission) Remove(perms ...Permission) Permission {
	for _, perm := range perms {
		if p.Has(perm) {
			p ^= perm
		}
	}

	return p
}