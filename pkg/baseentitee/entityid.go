package baseentitee

import (
	"github.com/google/uuid"
)

// EntityId represents the unique identifier for an entity.
type EntityId uuid.UUID

// NewEntityId generates a new unique EntityId.
func NewEntityId() EntityId {
	return EntityId(uuid.New())
}

// getUUID returns the UUID value of the EntityId.
func (id EntityId) getUUID() uuid.UUID {
	return uuid.UUID(id)
}

// string returns the string representation of the EntityId.
func (id EntityId) string() string {
	return uuid.UUID(id).String()
}

// EntityIdFromUUID creates an EntityId from a given uuid.UUID.
func EntityIdFromUUID(u uuid.UUID) EntityId {
	return EntityId(u)
}
