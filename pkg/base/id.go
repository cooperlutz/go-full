package base

import (
	"github.com/google/uuid"
)

type EntityId uuid.UUID

func NewEntityId() EntityId {
	return EntityId(uuid.New())
}

func (id EntityId) uuid() uuid.UUID {
	return uuid.UUID(id)
}

func (id EntityId) string() string {
	return uuid.UUID(id).String()
}

func EntityIdFromUUID(u uuid.UUID) EntityId {
	return EntityId(u)
}
