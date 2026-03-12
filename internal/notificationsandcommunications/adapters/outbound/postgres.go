//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package outbound

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/cooperlutz/go-full/internal/notificationsandcommunications/app/query"
	"github.com/cooperlutz/go-full/internal/notificationsandcommunications/domain/notificationsandcommunications"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// PostgresAdapter implements the examination repository using Postgres as the data store.
type PostgresAdapter struct {
	Handler IQuerierNotificationsAndCommunications
}

// NewPostgresAdapter creates a new instance of PostgresAdapter.
func NewPostgresAdapter(db deebee.IDatabase) PostgresAdapter {
	return PostgresAdapter{
		Handler: NewQueriesWrapper(db),
	}
}

func (p PostgresAdapter) FindAllNotifications(ctx context.Context) ([]query.Notification, error) {
	ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.adapters.outbound.postgres.find_all_notification")
	defer span.End()

	notifications, err := p.Handler.FindAllNotifications(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return notificationsandcommunicationsNotificationsToQuery(notifications)
}

func (p PostgresAdapter) FindOneNotification(ctx context.Context, id uuid.UUID) (query.Notification, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_notification")
	defer span.End()

	notification, err := p.GetNotification(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.Notification{}, err
	}

	return mapEntityNotificationToQuery(notification), nil
}

// AddNotification adds a new exam to the database.
func (p PostgresAdapter) AddNotification(ctx context.Context, notification *notificationsandcommunications.Notification) error {
	ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.adapters.outbound.postgres.add_notification")
	defer span.End()

	dbNotification := mapEntityNotificationToDB(notification)

	err := p.Handler.AddNotification(ctx, AddNotificationParams(dbNotification))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetNotification(ctx context.Context, id uuid.UUID) (*notificationsandcommunications.Notification, error) {
	ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.adapters.outbound.postgres.get_notification")
	defer span.End()

	notification, err := p.Handler.GetNotification(
		ctx,
		GetNotificationParams{NotificationID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return notification.toDomain()
}

func (p PostgresAdapter) UpdateNotification(
	ctx context.Context,
	notificationId uuid.UUID,
	updateFn func(e *notificationsandcommunications.Notification) (*notificationsandcommunications.Notification, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.adapters.outbound.postgres.update_notification")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	notification, err := p.GetNotification(ctx, notificationId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedNotification, err := updateFn(notification)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbNotification := mapEntityNotificationToDB(updatedNotification)

	err = p.Handler.UpdateNotification(ctx, UpdateNotificationParams(dbNotification))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) FindAllNotificationTemplates(ctx context.Context) ([]query.NotificationTemplate, error) {
	ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.adapters.outbound.postgres.find_all_notification_template")
	defer span.End()

	notificationtemplates, err := p.Handler.FindAllNotificationTemplates(ctx)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return notificationsandcommunicationsNotificationTemplatesToQuery(notificationtemplates)
}

func (p PostgresAdapter) FindOneNotificationTemplate(ctx context.Context, id uuid.UUID) (query.NotificationTemplate, error) {
	ctx, span := telemetree.AddSpan(ctx, "examination.adapters.outbound.postgres.find_one_notification_template")
	defer span.End()

	notificationtemplate, err := p.GetNotificationTemplate(ctx, id)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.NotificationTemplate{}, err
	}

	return mapEntityNotificationTemplateToQuery(notificationtemplate), nil
}

// AddNotificationTemplate adds a new exam to the database.
func (p PostgresAdapter) AddNotificationTemplate(ctx context.Context, notificationtemplate *notificationsandcommunications.NotificationTemplate) error {
	ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.adapters.outbound.postgres.add_notification_template")
	defer span.End()

	dbNotificationTemplate := mapEntityNotificationTemplateToDB(notificationtemplate)

	err := p.Handler.AddNotificationTemplate(ctx, AddNotificationTemplateParams(dbNotificationTemplate))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

func (p PostgresAdapter) GetNotificationTemplate(ctx context.Context, id uuid.UUID) (*notificationsandcommunications.NotificationTemplate, error) {
	ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.adapters.outbound.postgres.get_notification_template")
	defer span.End()

	notificationtemplate, err := p.Handler.GetNotificationTemplate(
		ctx,
		GetNotificationTemplateParams{NotificationTemplateID: pgxutil.UUIDToPgtypeUUID(id)},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return notificationtemplate.toDomain()
}

func (p PostgresAdapter) UpdateNotificationTemplate(
	ctx context.Context,
	notificationtemplateId uuid.UUID,
	updateFn func(e *notificationsandcommunications.NotificationTemplate) (*notificationsandcommunications.NotificationTemplate, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "notificationsandcommunications.adapters.outbound.postgres.update_notificationtemplate")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	notificationtemplate, err := p.GetNotificationTemplate(ctx, notificationtemplateId)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	updatedNotificationTemplate, err := updateFn(notificationtemplate)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbNotificationTemplate := mapEntityNotificationTemplateToDB(updatedNotificationTemplate)

	err = p.Handler.UpdateNotificationTemplate(ctx, UpdateNotificationTemplateParams(dbNotificationTemplate))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

// finishTransaction commits or rolls back the transaction based on the error state.
func (p PostgresAdapter) finishTransaction(ctx context.Context, err error, tx pgx.Tx) error {
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			telemetree.RecordError(ctx, rollbackErr, "failed to rollback tx")

			return rollbackErr
		}

		return err
	} else {
		if commitErr := tx.Commit(ctx); commitErr != nil {
			telemetree.RecordError(ctx, commitErr, "failed to commit tx")

			return commitErr
		}

		return nil
	}
}
