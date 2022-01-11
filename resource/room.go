package resource

import (
	"github.com/sakuraapp/shared/model"
	"github.com/sakuraapp/shared/resource/role"
)

type Room struct {
	Id model.RoomId `json:"id"`
	Name string  `json:"name"`
	Owner *User  `json:"owner"`
	Private bool `json:"private"`
}

func NewRoom(room *model.Room) *Room {
	if room.Owner == nil && room.OwnerId != 0 {
		room.Owner = &model.User{Id: room.OwnerId}
	}

	return &Room{
		room.Id,
		room.Name,
		NewUser(room.Owner),
		room.Private,
	}
}

func NewRoomList(rooms []model.Room) []*Room {
	list := make([]*Room, len(rooms))

	for i, v := range rooms {
		list[i] = NewRoom(&v)
	}

	return list
}

type RoomMember struct {
	User  *User       `json:"user"`
	Roles []role.Id `json:"roles" json:"roles"`
}

func NewRoomMember(member *model.RoomMember) *RoomMember {
	user := NewUser(&member.User)
	roles := make([]role.Id, 0, len(member.Roles))

	for _, userRole := range member.Roles {
		roles = append(roles, userRole.RoleId)
	}

	return &RoomMember{
		User: user,
		Roles: roles,
	}
}

func NewRoomMemberList(members []model.RoomMember) []*RoomMember {
	list := make([]*RoomMember, len(members))

	for i, v := range members {
		list[i] = NewRoomMember(&v)
	}

	return list
}