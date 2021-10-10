package resource

import (
	"github.com/sakuraapp/shared/model"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	Id model.UserId `json:"id"`
	Username string `json:"username,omitempty"`
	Discriminator string `json:"discriminator,omitempty"`
	Avatar null.String `json:"avatar"`
}

func NewUser(user *model.User) *User {
	return &User{
		Id: user.Id,
		Username: user.Username,
		Discriminator: user.Discriminator.ValueOrZero(),
		Avatar: user.Avatar,
	}
}

func NewUserList(users []model.User) []*User {
	list := make([]*User, len(users))

	for i, v := range users {
		list[i] = NewUser(&v)
	}

	return list
}