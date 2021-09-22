package resources

import (
	"github.com/sakuraapp/shared/models"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	Id int64 `json:"id"`
	Username string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar null.String `json:"avatar"`
}

func NewUser(user *models.User) *User {
	return &User{
		Id: user.Id,
		Username: user.Username,
		Discriminator: user.Discriminator.ValueOrZero(),
		Avatar: user.Avatar,
	}
}