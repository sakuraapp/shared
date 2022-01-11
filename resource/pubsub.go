package resource

import (
	"github.com/sakuraapp/shared/model"
	"github.com/vmihailenco/msgpack/v5"
)

type MessageType int

const (
	NORMAL_MESSAGE MessageType = iota
	BROADCAST_MESSAGE
	SERVER_MESSAGE
)

type MessageTarget struct {
	UserIds []model.UserId `msgpack:"u,omitempty"`
	IgnoredSessionIds map[string]bool `msgpack:"i,omitempty"`
}

type ServerMessage struct {
	Type MessageType `msgpack:"t,omitempty"`
	Target MessageTarget `msgpack:"tr,omitempty"`
	Data Packet `msgpack:"d,omitempty"`
	Origin string `msgpack:"o,omitempty"` // source/origin node of the message
}

type rawServerMessage ServerMessage

func (m ServerMessage) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal((rawServerMessage)(m))
}

func (m *ServerMessage) UnmarshalBinary(b []byte) error {
	return msgpack.Unmarshal(b, (*rawServerMessage)(m))
}