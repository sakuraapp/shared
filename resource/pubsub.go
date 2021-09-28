package resource

type MessageType int

const (
	NORMAL_MESSAGE MessageType = iota
	BROADCAST_MESSAGE
	SERVER_MESSAGE
)

type MessageTarget struct {
	UserIds []string `msgpack:"u,omitempty"`
	IgnoredUserIds []string `msgpack:"i,omitempty"`
}

type ServerMessage struct {
	Type MessageType `msgpack:"t,omitempty"`
	Target MessageTarget `msgpack:"tr,omitempty"`
	Data Packet `msgpack:"d,omitempty"`
}