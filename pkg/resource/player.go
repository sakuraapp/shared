package resource

import (
	"github.com/sakuraapp/shared/pkg/model"
	"time"
)

const (
	MediaItemTypeNormal = 0
	MediaItemTypeChakra = 1
)

type MediaItemType int

type MediaItemInfo struct {
	Title string `json:"title" redis:"title" msgpack:"title"`
	Icon string `json:"icon" redis:"icon" msgpack:"icon"`
	Url string `json:"url" redis:"url" msgpack:"url"`
}

type MediaItem struct {
	Id     string       `json:"id" redis:"id" msgpack:"id"`
	Author model.UserId `json:"author" redis:"author" msgpack:"author"`
	Type   MediaItemType `json:"type" redis:"type" msgpack:"type"`
	*MediaItemInfo
}

/* type rawMediaItem MediaItem

func (i MediaItem) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal((rawMediaItem)(i))
}

func (i *MediaItem) UnmarshalBinary(b []byte) error {
	return msgpack.Unmarshal(b, (*rawMediaItem)(i))
} */

type PlayerState struct {
	CurrentTime float64 `json:"currentTime" redis:"currentTime" msgpack:"currentTime"`
	IsPlaying bool `json:"playing" redis:"playing" msgpack:"playing"`
	PlaybackStart time.Time `json:"-" redis:"playbackStart" msgpack:"playbackStart"`
}