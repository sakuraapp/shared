package model

type RoomId = int32

type Room struct {
	Id RoomId
	Name string
	OwnerId UserId
	Owner *User   `pg:"rel:has-one"`
	Private bool `pg:",use_zero"`
}

type RoomMember struct {
	User
	Roles []*UserRole `pg:"rel:has-many"`
}