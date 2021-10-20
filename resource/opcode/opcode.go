package opcode

type Opcode int

const (
	Disconnect Opcode = 0
	Authenticate = 1
	JoinRoom = 2
	LeaveRoom = 3
	CreateRoom = 4
	AddUser = 5
	RemoveUser = 6
	PlayerState = 7
	SendMessage = 8
	QueueAdd = 9
	QueueRemove = 10
	VideoSet = 11
	VideoEnd = 12
	VideoSkip = 13
	RoomJoinRequest = 14
	SetRoomType = 15
	DispatchControl = 16
)