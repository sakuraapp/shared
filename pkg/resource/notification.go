package resource

type NotificationType int

const (
	NotificationJoinRequest = 0
)

type Notification struct {
	Id   string           `json:"id"`
	Type NotificationType `json:"type"`
	Data interface{}      `json:"data,omitempty"`
}