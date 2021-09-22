package resources

import (
	"github.com/sakuraapp/shared/models"
)

type Room struct {
	Id int64 `json:"id"`
	Name string  `json:"name"`
	Owner *User  `json:"owner"`
	Private bool `json:"private"`
}

func NewRoom(room *models.Room) *Room {
	return &Room{
		room.Id,
		room.Name,
		NewUser(room.Owner),
		room.Private,
	}
}

func NewRoomList(rooms []models.Room) []*Room {
	list := make([]*Room, len(rooms))

	for i, v := range rooms {
		list[i] = NewRoom(&v)
	}

	return list
}