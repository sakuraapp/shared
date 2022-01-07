package model

type RoomId = int32

type Room struct {
	Id RoomId
	Name string
	OwnerId UserId
	Owner *User   `pg:"rel:has-one"`
	Private bool `pg:",use_zero"`
}