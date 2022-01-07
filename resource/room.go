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
	User  User        `json:"user"`
	Roles []role.Role `json:"roles" json:"roles"`
}