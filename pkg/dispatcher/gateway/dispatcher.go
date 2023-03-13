package gateway

import (
	"fmt"
	"github.com/sakuraapp/pubsub"
	"github.com/sakuraapp/shared/pkg/constant"
	"github.com/sakuraapp/shared/pkg/model"
	"github.com/sakuraapp/shared/pkg/resource"
	"github.com/sakuraapp/shared/pkg/resource/permission"
)

type Message = pubsub.Message[resource.Packet]
type SubscriptionManager = pubsub.SubscriptionManager[*Message]

func NewMessage(payload resource.Packet, opts ...pubsub.FilterMap) *Message {
	return pubsub.NewMessage[resource.Packet](payload, opts...)
}

func NewSubscriptionManager(client pubsub.Client) *SubscriptionManager {
	return pubsub.NewSubscriptionManager[*Message](client)
}

const (
	MessageFilterType pubsub.MessageFilterKind = iota
	MessageFilterIgnoredSession
	MessageFilterPermissions
	MessageFilterRoom
)

type MessageType int

const (
	NormalMessage MessageType = iota
	ServerMessage
)

type FilterMap pubsub.FilterMap

func NewFilterMap() FilterMap {
	return FilterMap{}
}

func (m FilterMap) WithType(kind MessageType) FilterMap {
	m[MessageFilterType] = kind

	return m
}

func (m FilterMap) WithIgnoredSession(sessionId string) FilterMap {
	m[MessageFilterIgnoredSession] = sessionId

	return m
}

func (m FilterMap) WithPermissions(perm permission.Permission) FilterMap {
	m[MessageFilterPermissions] = perm

	return m
}

func (m FilterMap) WithRoom(roomId model.RoomId) FilterMap {
	m[MessageFilterRoom] = roomId

	return m
}

type MessageTargetNode string

func (t MessageTargetNode) Build() string {
	return fmt.Sprintf(constant.NodeTopicFmt, t)
}

func NewNodeTarget(id string) MessageTargetNode {
	return MessageTargetNode(id)
}

type MessageTargetSession string

func (t MessageTargetSession) Build() string {
	return fmt.Sprintf(constant.SessionTopicFmt, t)
}

func NewSessionTarget(id string) MessageTargetSession {
	return MessageTargetSession(id)
}

type MessageTargetUser model.UserId

func (t MessageTargetUser) Build() string {
	return fmt.Sprintf(constant.UserTopicFmt, t)
}

func NewUserTarget(id model.UserId) MessageTargetUser {
	return MessageTargetUser(id)
}

type MessageTargetRoom model.RoomId

func (t MessageTargetRoom) Build() string {
	return fmt.Sprintf(constant.RoomTopicFmt, t)
}

func NewRoomTarget(id model.RoomId) MessageTargetRoom {
	return MessageTargetRoom(id)
}
