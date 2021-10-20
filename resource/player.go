package resource

import (
	"github.com/sakuraapp/shared/resource/opcode"
	"github.com/vmihailenco/msgpack/v5"
	"time"
)

type MediaItemInfo struct {
	Title string `json:"title" redis:"title" msgpack:"title"`
	Icon string `json:"icon" redis:"icon" msgpack:"icon"`
	Url string `json:"url" redis:"url" msgpack:"url"`
}

type MediaItem struct {
	Id string `json:"id" redis:"id" msgpack:"id"`
	*MediaItemInfo
}

type rawMediaItem MediaItem

func (i MediaItem) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal((rawMediaItem)(i))
}

func (i *MediaItem) UnmarshalBinary(b []byte) error {
	return msgpack.Unmarshal(b, (*rawMediaItem)(i))
}

type PlayerState struct {
	CurrentTime time.Duration `json:"currentTime" redis:"currentTime" msgpack:"currentTime"`
	IsPlaying bool `json:"playing" redis:"playing" msgpack:"playing"`
	PlaybackStart time.Time `json:"-" redis:"playbackStart" msgpack:"playbackStart"`
}

type rawPlayerState PlayerState

func (s PlayerState) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal((rawPlayerState)(s))
}

func (s *PlayerState) UnmarshalBinary(b []byte) error {
	return msgpack.Unmarshal(b, (*rawPlayerState)(s))
}

func (s *PlayerState) BuildPacket() Packet {
	return BuildPacket(opcode.PlayerState, s)
}