package resource

import "github.com/sakuraapp/shared/model"

type Message struct {
	Id string `json:"id" msgpack:"id"`
	Author model.UserId `json:"author" msgpack:"author"`
	Content string `json:"content" msgpack:"content"`
}
