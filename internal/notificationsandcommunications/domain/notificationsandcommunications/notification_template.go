package notificationsandcommunications

import (
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

type NotificationTemplate struct {
	*baseentitee.EntityMetadata
	//
	//templateId string
	//
	//name string
	//
	//notificationType string
	//
	//channel string
	//
	//subjectTemplate *string
	//
	//bodyTemplate string
	//
	//isActive bool
	//
	// TODO
}

func NewNotificationTemplate(
// templateId string,
//
// name string,
//
// notificationType string,
//
// channel string,
//
// subjectTemplate *string,
//
// bodyTemplate string,
//
// isActive bool,
) *NotificationTemplate {
	return &NotificationTemplate{
		EntityMetadata: baseentitee.NewEntityMetadata(),
		//
		//templateId: templateId,
		//
		//name: name,
		//
		//notificationType: notificationType,
		//
		//channel: channel,
		//
		//subjectTemplate: subjectTemplate,
		//
		//bodyTemplate: bodyTemplate,
		//
		//isActive: isActive,
		//
	}
}
