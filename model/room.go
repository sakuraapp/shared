package model

type Room struct {
	Id int32
	Name string
	OwnerId UserId ``
	Owner *User   `pg:"rel:has-one"`
	Private bool
}