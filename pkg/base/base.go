package base

import (
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/utilitee"
)

type EntityMetadata struct {
	entityId  EntityId
	createdAt CreatedAt
	updatedAt UpdatedAt
	deletedAt *DeletedAt
	deleted   DeletedFlag
}

func NewEntityMetadata() EntityMetadata {
	return EntityMetadata{
		entityId:  NewEntityId(),
		createdAt: NewCreatedAt(),
		updatedAt: NewUpdatedAt(),
		deletedAt: NewDeletedAt(),
		deleted:   NewDeletedFlag(),
	}
}

func (m EntityMetadata) GetId() EntityId {
	return m.entityId
}

func (m EntityMetadata) GetIdUUID() uuid.UUID {
	return m.entityId.uuid()
}

func (m EntityMetadata) GetIdString() string {
	return m.entityId.string()
}

func (m EntityMetadata) GetCreatedAt() CreatedAt {
	return m.createdAt
}

func (m EntityMetadata) GetUpdatedAt() UpdatedAt {
	return m.updatedAt
}

func (m *EntityMetadata) MarkUpdated() {
	m.updatedAt = NewUpdatedAt()
}

func (m EntityMetadata) GetDeletedAt() *DeletedAt {
	return m.deletedAt
}

func (m EntityMetadata) GetDeletedAtTime() *time.Time {
	if m.deletedAt == nil {
		return nil
	}

	return m.deletedAt.getTime()
}

func (m EntityMetadata) GetCreatedAtTime() time.Time {
	return m.createdAt.getTime()
}

func (m EntityMetadata) GetUpdatedAtTime() time.Time {
	return m.updatedAt.getTime()
}

func (m EntityMetadata) GetDeletedFlag() DeletedFlag {
	return m.deleted
}

func (m EntityMetadata) IsDeleted() bool {
	return m.deleted.getBool()
}

func (m *EntityMetadata) MarkDeleted() {
	m.deleted = DeletedFlagFromBool(true)
	now := utilitee.RightNow()
	dAt := DeletedAtFromTime(&now)
	m.deletedAt = dAt
}

func MapToEntityMetadata(
	id EntityId,
	createdAt CreatedAt,
	updatedAt UpdatedAt,
	deleted bool,
	deletedAt *DeletedAt,
) EntityMetadata {
	return EntityMetadata{
		entityId:  id,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deleted:   DeletedFlagFromBool(deleted),
		deletedAt: deletedAt,
	}
}

func MapToEntityMetadataFromCommonTypes(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
) EntityMetadata {
	return EntityMetadata{
		entityId:  EntityIdFromUUID(id),
		createdAt: CreatedAtFromTime(createdAt),
		updatedAt: UpdatedAtFromTime(updatedAt),
		deleted:   DeletedFlagFromBool(deleted),
		deletedAt: DeletedAtFromTime(deletedAt),
	}
}
