package resource

import (
	"github.com/sakuraapp/shared/pkg/model"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	Id       model.UserId `json:"id"`
	Username string       `json:"username,omitempty"`
	Discriminator string `json:"discriminator,omitempty"`
	Avatar null.String `json:"avatar"`
}

func (b *Builder) NewUser(user *model.User) *User {
	u := &User{
		Id: user.Id,
		Username: user.Username,
		Discriminator: user.Discriminator.ValueOrZero(),
		Avatar: user.Avatar,
	}

	if b.userFormatter != nil {
		u = b.userFormatter(u)
	}

	return u
}

func (b *Builder) NewUserList(users []model.User) []*User {
	list := make([]*User, len(users))

	for i, v := range users {
		list[i] = b.NewUser(&v)
	}

	return list
}