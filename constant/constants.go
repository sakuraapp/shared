package constant

import "time"

const (
	UserSessionsFmt = "user_sessions.%v"
	SessionFmt = "session.%v"

	RoomFmt = "room.%v"
	RoomInviteFmt = RoomFmt + ".invites"
	RoomUsersFmt = RoomFmt + ".users"
	RoomUserSessionsFmt = RoomUsersFmt + "%v.sessions"
	RoomQueueFmt = RoomFmt + ".queue"
	RoomQueueItemsFmt = RoomQueueFmt + ".items"
	RoomCurrentItemFmt = RoomFmt + ".currentItem"
	RoomStateFmt = RoomFmt + ".state"
	RoomVideoEndAckFmt = RoomFmt + ".endAck"

	GatewayFmt = "gateway.%v"
	BroadcastChName = "gateway.broadcast"

	UserCacheFmt = "c.user.%v"
	UserCacheTTL = 15 * time.Minute

	RoomCacheFmt = "c.room.%v"
	RoomCacheTTL = 15 * time.Minute
)
