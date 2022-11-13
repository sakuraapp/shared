package constant

import "time"

const (
	NodeTopicFmt = "gateway.%v"
	SessionTopicFmt = "session.%v"
	UserTopicFmt = "user.%v"
	RoomTopicFmt = "room.%v"

	UserSessionsFmt = "user_sessions.%v"
	SessionFmt = "session.%v"

	RoomFmt = "room.%v"
	RoomJoinRequestsFmt = RoomFmt + ".requests"
	RoomUsersFmt = RoomFmt + ".users"
	RoomUserSessionsFmt = RoomUsersFmt + ".%v.sessions"
	RoomQueueFmt = RoomFmt + ".queue"
	RoomQueueItemsFmt = RoomQueueFmt + ".items"
	RoomCurrentItemFmt = RoomFmt + ".currentItem"
	RoomStateFmt = RoomFmt + ".state"
	RoomVideoEndAckFmt = RoomFmt + ".endAck"

	UserCacheFmt = "c.user.%v"
	UserCacheTTL = 15 * time.Minute

	RoomCacheFmt = "c.room.%v"
	RoomCacheTTL = 15 * time.Minute
)
