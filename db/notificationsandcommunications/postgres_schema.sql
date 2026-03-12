CREATE SCHEMA IF NOT EXISTS notificationsandcommunications;


CREATE TABLE IF NOT EXISTS notificationsandcommunications.notifications (
    notification_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
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
);

-- create index to optimize queries searching by notification_id
CREATE INDEX IF NOT EXISTS idx_notifications_and_communications_notifications_id
ON notificationsandcommunications.notifications (notification_id);

CREATE TABLE IF NOT EXISTS notificationsandcommunications.notification_templates (
    notification_template_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    deleted BOOLEAN NOT NULL DEFAULT FALSE
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
);

-- create index to optimize queries searching by notification_template_id
CREATE INDEX IF NOT EXISTS idx_notifications_and_communications_notification_templates_id
ON notificationsandcommunications.notification_templates (notification_template_id);
