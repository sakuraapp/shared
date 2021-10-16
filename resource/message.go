package resource

import "github.com/sakuraapp/shared/model"

type Message struct {
	Id string `json:"id"`
	Author model.UserId `json:"author"`
	Content string `json:"content"`
}
