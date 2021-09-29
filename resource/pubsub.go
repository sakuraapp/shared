package resource

import "github.com/sakuraapp/shared/model"

type MessageType int

const (
	NORMAL_MESSAGE MessageType = iota
	BROADCAST_MESSAGE
	SERVER_MESSAGE
)

type MessageTarget struct {
	UserIds map[model.UserId]bool `msgpack:"u,omitempty"`
	IgnoredSessionIds map[string]bool `msgpack:"i,omitempty"`
}

type ServerMessage struct {
	Type MessageType `msgpack:"t,omitempty"`
	Target MessageTarget `msgpack:"tr,omitempty"`
	Data Packet `msgpack:"d,omitempty"`
}