package resource

import (
	"github.com/sakuraapp/shared/pkg/model"
)

type Message struct {
	Id      string       `json:"id" msgpack:"id"`
	Author  model.UserId `json:"author" msgpack:"author"`
	Content string       `json:"content" msgpack:"content"`
}
