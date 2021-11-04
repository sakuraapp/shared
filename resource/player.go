package resource

import (
	"github.com/sakuraapp/shared/model"
	"time"
)

type MediaItemInfo struct {
	Title string `json:"title" redis:"title" msgpack:"title"`
	Icon string `json:"icon" redis:"icon" msgpack:"icon"`
	Url string `json:"url" redis:"url" msgpack:"url"`
}

type MediaItem struct {
	Id string `json:"id" redis:"id" msgpack:"id"`
	Author model.UserId `json:"author" redis:"author" msgpack:"author"`
	*MediaItemInfo
}

type PlayerState struct {
	CurrentTime time.Duration `json:"currentTime" redis:"currentTime" msgpack:"currentTime"`
	IsPlaying bool `json:"playing" redis:"playing" msgpack:"playing"`
	PlaybackStart time.Time `json:"-" redis:"playbackStart" msgpack:"playbackStart"`
}