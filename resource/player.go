package resource

import "time"

type MediaItemInfo struct {
	Title string `json:"title" redis:"title" msgpack:"title"`
	Icon string `json:"icon" redis:"icon" msgpack:"icon"`
	Url string `json:"url" redis:"url" msgpack:"url"`
}

type MediaItem struct {
	Id string `json:"id" redis:"id" msgpack:"id"`
	*MediaItemInfo
}

type PlayerState struct {
	CurrentTime time.Time `json:"current_time" redis:"current_time" msgpack:"current_time"`
	IsPlaying bool `json:"playing" redis:"playing" msgpack:"playing"`
	PlaybackStart time.Time `json:"-" redis:"playback_start" msgpack:"playback_start"`
}