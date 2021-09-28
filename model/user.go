package model

import (
	"gopkg.in/guregu/null.v4"
)

type UserId = int32

type User struct {
	Id UserId
	Username string
	Avatar null.String
	Provider string
	AccessToken null.String
	RefreshToken null.String
	ExternalUserID null.String
	Discriminator null.String `pg:"-"`
}