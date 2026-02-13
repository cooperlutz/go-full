package baseentitee

import (
	"time"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/utilitee"
)

// EntityMetadata represents the baseentitee metadata object for an entity.
type EntityMetadata struct {
	entityId  EntityId
	createdAt CreatedAt
	updatedAt UpdatedAt
	deletedAt *DeletedAt
	deleted   DeletedFlag

	events []any // Domain events associated with the entity
}

// NewEntityMetadata creates a new EntityMetadata with default values.
func NewEntityMetadata() *EntityMetadata {
	return &EntityMetadata{
		entityId:  NewEntityId(),
		createdAt: NewCreatedAt(),
		updatedAt: NewUpdatedAt(),
		deletedAt: NewDeletedAt(),
		deleted:   NewDeletedFlag(),
	}
}

// GetId returns the EntityId of the entity.
func (m EntityMetadata) GetId() EntityId {
	return m.entityId
}

// GetIdUUID returns the UUID of the entity.
func (m EntityMetadata) GetIdUUID() uuid.UUID {
	return m.entityId.getUUID()
}

// GetIdString returns the string representation of the EntityId.
func (m EntityMetadata) GetIdString() string {
	return m.entityId.string()
}

// GetCreatedAt returns the CreatedAt of the entity.
func (m EntityMetadata) GetCreatedAt() CreatedAt {
	return m.createdAt
}

// GetUpdatedAt returns the UpdatedAt of the entity.
func (m EntityMetadata) GetUpdatedAt() UpdatedAt {
	return m.updatedAt
}

// MarkUpdated updates the UpdatedAt timestamp to the current time.
func (m *EntityMetadata) MarkUpdated() {
	m.updatedAt = NewUpdatedAt()
}

// GetDeletedAt returns the DeletedAt of the entity.
func (m EntityMetadata) GetDeletedAt() *DeletedAt {
	return m.deletedAt
}

// GetDeletedAtTime returns the time.Time value of DeletedAt, or nil if not set.
func (m EntityMetadata) GetDeletedAtTime() *time.Time {
	if m.deletedAt == nil {
		return nil
	}

	return m.deletedAt.getTime()
}

// GetCreatedAtTime returns the time.Time value of CreatedAt.
func (m EntityMetadata) GetCreatedAtTime() time.Time {
	return m.createdAt.getTime()
}

// GetUpdatedAtTime returns the time.Time value of UpdatedAt.
func (m EntityMetadata) GetUpdatedAtTime() time.Time {
	return m.updatedAt.getTime()
}

// GetDeletedFlag returns the DeletedFlag of the entity.
func (m EntityMetadata) GetDeletedFlag() DeletedFlag {
	return m.deleted
}

// IsDeleted returns true if the entity is marked as deleted.
func (m EntityMetadata) IsDeleted() bool {
	return m.deleted.getBool()
}

// MarkDeleted marks the entity as deleted and sets the DeletedAt timestamp.
func (m *EntityMetadata) MarkDeleted() {
	m.deleted = DeletedFlagFromBool(true)
	now := utilitee.RightNow()

	dAt := DeletedAtFromTime(&now)
	m.deletedAt = dAt

	m.MarkUpdated()
}

// GetDomainEventsAndClear retrieves and clears the domain events associated with the entity.
func (e *EntityMetadata) GetDomainEventsAndClear() []any {
	// set events to a new slice and return the old one
	if !e.domainEventsExist() {
		return nil
	}

	events := e.events
	// this clears the events by assigning a new slice with a length of 0
	e.events = []any{}

	return events
}

// domainEventsExist checks if there are any domain events associated with the entity.
func (e EntityMetadata) domainEventsExist() bool {
	return len(e.events) > 0
}

// raiseDomainEvent appends a new domain event to the entity's event list.
func (e *EntityMetadata) RaiseDomainEvent(event any) {
	e.events = append(e.events, event)
}

// MapToEntityMetadata maps the baseentitee metadata types to EntityMetadata.
func MapToEntityMetadata(
	id EntityId,
	createdAt CreatedAt,
	updatedAt UpdatedAt,
	deleted DeletedFlag,
	deletedAt *DeletedAt,
) *EntityMetadata {
	return &EntityMetadata{
		entityId:  id,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deleted:   deleted,
		deletedAt: deletedAt,
	}
}

// MapToEntityMetadataFromCommonTypes maps common Go types to EntityMetadata.
func MapToEntityMetadataFromCommonTypes(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	deleted bool,
	deletedAt *time.Time,
) *EntityMetadata {
	return &EntityMetadata{
		entityId:  EntityIdFromUUID(id),
		createdAt: CreatedAtFromTime(createdAt),
		updatedAt: UpdatedAtFromTime(updatedAt),
		deleted:   DeletedFlagFromBool(deleted),
		deletedAt: DeletedAtFromTime(deletedAt),
	}
}
