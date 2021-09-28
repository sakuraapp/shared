package resource

import "github.com/sakuraapp/shared/model"

type MessageType int

const (
	NORMAL_MESSAGE MessageType = iota
	BROADCAST_MESSAGE
	SERVER_MESSAGE
)

type MessageTarget struct {
	UserIds []model.UserId `msgpack:"u,omitempty"`
	IgnoredSessionIds []string `msgpack:"i,omitempty"`
}

type ServerMessage struct {
	Type MessageType `msgpack:"t,omitempty"`
	Target MessageTarget `msgpack:"tr,omitempty"`
	Data Packet `msgpack:"d,omitempty"`
}