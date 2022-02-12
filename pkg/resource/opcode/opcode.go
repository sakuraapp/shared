package opcode

type Opcode int

const (
	Disconnect   Opcode = 0
	Authenticate        = 1
	JoinRoom = 2
	LeaveRoom = 3
	CreateRoom = 4
	AddUser = 5
	RemoveUser = 6
	PlayerState = 7
	Seek = 8
	SendMessage = 9
	QueueAdd = 10
	QueueRemove = 11
	VideoSet = 12
	VideoEnd = 13
	VideoSkip = 14
	RoomJoinRequest = 15
	DispatchControl = 16
	UpdateRoom = 17
	KickUser = 18
	AddRole = 19
	RemoveRole = 20
	UpdatePermissions = 21
	AddNotification = 22
	Error = 23
)