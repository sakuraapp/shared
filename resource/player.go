package resource

type MediaItemInfo struct {
	Title string `json:"title" redis:"title" msgpack:"title"`
	Icon string `json:"icon" redis:"icon" msgpack:"icon"`
	Url string `json:"url" redis:"url" msgpack:"url"`
}

type MediaItem struct {
	Id string `json:"id" redis:"id" msgpack:"id"`
	*MediaItemInfo
}