package resource

import (
	"github.com/sakuraapp/shared/model"
)

type Room struct {
	Id int64 `json:"id"`
	Name string  `json:"name"`
	Owner *User  `json:"owner"`
	Private bool `json:"private"`
}

func NewRoom(room *model.Room) *Room {
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