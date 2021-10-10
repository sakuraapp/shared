package constant

const (
	UserSessionsFmt = "user_sessions.%v"
	SessionFmt = "session.%v"

	RoomFmt = "room.%v"
	RoomInviteFmt = RoomFmt + ".invites"
	RoomUsersFmt = RoomFmt + ".users"
	RoomUserSessionsFmt = RoomUsersFmt + "%v.sessions"

	GatewayFmt = "gateway.%v"
	BroadcastChName = "gateway.broadcast"
)
