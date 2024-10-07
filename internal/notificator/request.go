package notificator

import (
	"gitlab.com/tokene/keyserver-svc/resources"
	"time"
)

type NotificationPriority int32

type CreateNotificationRequest struct {
	Data CreateNotificationData `json:"data"`
}

type CreateNotificationData struct {
	resources.Key
	Attributes    CreateNotificationAttributes    `json:"attributes"`
	Relationships CreateNotificationRelationships `json:"relationships"`
}

type CreateNotificationAttributes struct {
	Channel      *string               `json:"channel,omitempty"`
	Message      Message               `json:"message"`
	Priority     *NotificationPriority `json:"priority,omitempty"`
	ScheduledFor *time.Time            `json:"scheduled_for,omitempty"`
	Token        *string               `json:"token,omitempty"`
	Topic        string                `json:"topic"`
}

type Message struct {
	Type       string       `json:"type"`
	Attributes MessageAttrs `json:"attributes"`
}

type MessageAttrs struct {
	Payload interface{} `json:"payload"`
}

type CreateNotificationRelationships struct {
	Destinations resources.RelationCollection `json:"destinations"`
}

type VerificationPayload struct {
	Code string `json:"Code"`
}
