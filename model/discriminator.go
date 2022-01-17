package model

type Discriminator struct {
	Id int32
	Name string
	Value string
	OwnerId UserId
	Owner *User `pg:"rel:has-one"`
}