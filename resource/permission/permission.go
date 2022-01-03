package permission

type Permission int

const (
	QUEUE_ADD Permission = 1
	QUEUE_EDIT = 2
	VIDEO_REMOTE = 4
	ALL = QUEUE_ADD | QUEUE_EDIT | VIDEO_REMOTE
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