package resource

import (
	"github.com/sakuraapp/shared/model"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	Id model.UserId `json:"id"`
	Username string `json:"username"`
	Discriminator string `json:"discriminator"`
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