package resource

import (
	"github.com/sakuraapp/shared/pkg/model"
	"github.com/sakuraapp/shared/pkg/resource/role"
)

type Room struct {
	Id      model.RoomId `json:"id"`
	Name    string        `json:"name"`
	Owner   *User         `json:"owner"`
	Private bool          `json:"private"`
}

func (b *Builder) NewRoom(room *model.Room) *Room {
	if room.Owner == nil && room.OwnerId != 0 {
		room.Owner = &model.User{Id: room.OwnerId}
	}

	return &Room{
		room.Id,
		room.Name,
		b.NewUser(room.Owner),
		room.Private,
	}
}

func (b *Builder) NewRoomList(rooms []model.Room) []*Room {
	list := make([]*Room, len(rooms))

	for i, v := range rooms {
		list[i] = b.NewRoom(&v)
	}

	return list
}

type RoomMember struct {
	User  *User     `json:"user"`
	Roles []role.Id `json:"roles" json:"roles"`
}

func (b *Builder) NewRoomMember(member *model.RoomMember) *RoomMember {
	user := b.NewUser(&member.User)

	roles := make([]role.Id, 0, len(member.Roles) + 1)
	roles = append(roles, role.MEMBER)

	for _, userRole := range member.Roles {
		roles = append(roles, userRole.RoleId)
	}

	return &RoomMember{
		User: user,
		Roles: roles,
	}
}

func (b *Builder) NewRoomMemberList(members []model.RoomMember) []*RoomMember {
	list := make([]*RoomMember, len(members))

	for i, v := range members {
		list[i] = b.NewRoomMember(&v)
	}

	return list
}