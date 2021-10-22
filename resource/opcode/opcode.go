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
	SetPlaying = 8
	SetTime = 9
	SendMessage = 10
	QueueAdd = 11
	QueueRemove = 12
	VideoSet = 13
	VideoEnd = 14
	VideoSkip = 15
	RoomJoinRequest = 16
	SetRoomType = 17
	DispatchControl = 18
)