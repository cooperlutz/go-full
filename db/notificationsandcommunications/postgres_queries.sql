
-- name: GetNotification :one
SELECT * FROM notificationsandcommunications.notifications
WHERE notification_id = $1;

-- name: AddNotification :exec
INSERT INTO notificationsandcommunications.notifications (
    notification_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
    --,
    --notification_id
    --,
    --recipient_id
    --,
    --recipient_type
    --,
    --channel
    --,
    --subject
    --,
    --message_body
    --,
    --status
    --,
    --sent_at
    --,
    --notification_type
    --
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdateNotification :exec
UPDATE notificationsandcommunications.notifications
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
    --,
    --notification_id
    --,
    --recipient_id
    --,
    --recipient_type
    --,
    --channel
    --,
    --subject
    --,
    --message_body
    --,
    --status
    --,
    --sent_at
    --,
    --notification_type
    --
    -- TODO
WHERE notification_id = $1;

-- name: FindOneNotification :one
SELECT * FROM notificationsandcommunications.notifications
WHERE notification_id = $1;

-- name: FindAllNotifications :many
SELECT * FROM notificationsandcommunications.notifications;


-- name: GetNotificationTemplate :one
SELECT * FROM notificationsandcommunications.notification_templates
WHERE notification_template_id = $1;

-- name: AddNotificationTemplate :exec
INSERT INTO notificationsandcommunications.notification_templates (
    notification_template_id,
    created_at,
    updated_at,
    deleted_at,
    deleted
    --,
    --template_id
    --,
    --name
    --,
    --notification_type
    --,
    --channel
    --,
    --subject_template
    --,
    --body_template
    --,
    --is_active
    --
    -- TODO
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    -- TODO
);

-- name: UpdateNotificationTemplate :exec
UPDATE notificationsandcommunications.notification_templates
SET
    created_at = $2,
    updated_at = $3,
    deleted_at = $4,
    deleted = $5
    --,
    --template_id
    --,
    --name
    --,
    --notification_type
    --,
    --channel
    --,
    --subject_template
    --,
    --body_template
    --,
    --is_active
    --
    -- TODO
WHERE notification_template_id = $1;

-- name: FindOneNotificationTemplate :one
SELECT * FROM notificationsandcommunications.notification_templates
WHERE notification_template_id = $1;

-- name: FindAllNotificationTemplates :many
SELECT * FROM notificationsandcommunications.notification_templates;

