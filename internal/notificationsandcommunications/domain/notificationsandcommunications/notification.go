package notificationsandcommunications

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type Notification struct {
	*baseentitee.EntityMetadata
	//
	//notificationId string
	//
	//recipientId string
	//
	//recipientType string
	//
	//channel string
	//
	//subject *string
	//
	//messageBody string
	//
	//status string
	//
	//sentAt *string
	//
	//notificationType string
	//
	// TODO
}

func NewNotification(
// notificationId string,
//
// recipientId string,
//
// recipientType string,
//
// channel string,
//
// subject *string,
//
// messageBody string,
//
// status string,
//
// sentAt *string,
//
// notificationType string,
) *Notification {
	return &Notification{
		EntityMetadata: baseentitee.NewEntityMetadata(),
		//
		//notificationId: notificationId,
		//
		//recipientId: recipientId,
		//
		//recipientType: recipientType,
		//
		//channel: channel,
		//
		//subject: subject,
		//
		//messageBody: messageBody,
		//
		//status: status,
		//
		//sentAt: sentAt,
		//
		//notificationType: notificationType,
		//
	}
}
