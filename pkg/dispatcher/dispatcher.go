package dispatcher

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sakuraapp/shared/pkg/constant"
	"github.com/sakuraapp/shared/pkg/model"
	"github.com/sakuraapp/shared/pkg/resource"
	"github.com/vmihailenco/msgpack/v5"
)

type LocalDispatcher interface {
	DispatchLocal(msg resource.ServerMessage) error
	DispatchRoomLocal(roomId model.RoomId, msg resource.ServerMessage) error
	HandleServerMessage(msg *resource.ServerMessage)
}

type Dispatcher interface {
	LocalDispatcher
	Dispatch(msg resource.ServerMessage) error
	DispatchRoom(roomId model.RoomId, msg resource.ServerMessage) error
}

type RedisDispatcher struct {
	LocalDispatcher
	nodeId string
	ctx context.Context
	rdb *redis.Client
}

func (d *RedisDispatcher) Dispatch(msg resource.ServerMessage) error {
	msg.Origin = d.nodeId

	switch msg.Type {
	case resource.BROADCAST_MESSAGE:
		if d.LocalDispatcher != nil {
			err := d.DispatchLocal(msg)

			if err != nil {
				return err
			}
		}

		return d.rdb.Publish(d.ctx, constant.BroadcastChName, msg).Err()
	case resource.NORMAL_MESSAGE:
		fallthrough
	case resource.SERVER_MESSAGE:
		// note: both normal and server messages can be dispatched toward certain users
		// todo: re-investigate how server messages should be targeted

		pipe := d.rdb.Pipeline()
		locNodeId := d.nodeId
		roomId := msg.Target.RoomId

		for _, userId := range msg.Target.UserIds {
			if roomId == 0 {
				pipe.SMembers(d.ctx, fmt.Sprintf(constant.UserSessionsFmt, userId))
			} else {
				pipe.SMembers(d.ctx, fmt.Sprintf(constant.RoomUserSessionsFmt, roomId, userId))
			}
		}

		results, err := pipe.Exec(d.ctx)

		if err != nil {
			return err
		}

		pipe = d.rdb.Pipeline()

		var sessionKey string

		for _, result := range results {
			sessions := result.(*redis.StringSliceCmd).Val()

			for _, session := range sessions {
				sessionKey = fmt.Sprintf(constant.SessionFmt, session)
				pipe.HGet(d.ctx, sessionKey, "node_id")
			}
		}

		results, err = pipe.Exec(d.ctx)

		if err != nil && err != redis.Nil {
			return err
		}

		bytes, err := msgpack.Marshal(msg)

		if err != nil {
			return err
		}

		pipe = d.rdb.Pipeline()
		nodes := map[string]bool{}
		var nodeKey string

		for _, result := range results {
			cmd := result.(*redis.StringCmd)

			if cmd.Err() == redis.Nil {
				continue
			}

			nodeId := cmd.Val()

			if !nodes[nodeId] {
				nodes[nodeId] = true

				if nodeId == locNodeId {
					if d.LocalDispatcher != nil {
						if msg.Type == resource.SERVER_MESSAGE {
							d.HandleServerMessage(&msg)
						} else {
							d.DispatchLocal(msg)
						}
					}
				} else {
					nodeKey = fmt.Sprintf(constant.GatewayFmt, nodeId)
					pipe.Publish(d.ctx, nodeKey, bytes)
				}
			}
		}

		_, err = pipe.Exec(d.ctx)

		if err != nil {
			return err
		}
	}

	return nil
}

func (d *RedisDispatcher) DispatchRoom(roomId model.RoomId, msg resource.ServerMessage) error {
	msg.Origin = d.nodeId

	if d.LocalDispatcher != nil {
		err := d.DispatchRoomLocal(roomId, msg)

		if err != nil {
			return err
		}
	}

	roomKey := fmt.Sprintf(constant.RoomFmt, roomId)
	bytes, err := msgpack.Marshal(msg)

	if err != nil {
		return err
	}

	return d.rdb.Publish(d.ctx, roomKey, bytes).Err()
}

func NewRedis(ctx context.Context, ld LocalDispatcher, nodeId string, rdb *redis.Client) *RedisDispatcher {
	return &RedisDispatcher{
		LocalDispatcher: ld,
		ctx: ctx,
		nodeId: nodeId,
		rdb: rdb,
	}
}