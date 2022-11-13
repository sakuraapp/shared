package gateway

import (
	"github.com/sakuraapp/pubsub"
	"github.com/sakuraapp/shared/pkg/model"
)

type Dispatcher struct {
	pubsub.Dispatcher
}

func (d *Dispatcher) DispatchToRoom(id model.RoomId, message pubsub.Message) error {
	return d.DispatchTo(NewRoomTarget(id), message)
}

func NewDispatcher(dispatcher pubsub.Dispatcher) *Dispatcher {
	return &Dispatcher{dispatcher}
}