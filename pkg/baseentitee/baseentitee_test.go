package baseentitee_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

func TestNewEntityMetadata(t *testing.T) {
	newEntMeta := baseentitee.NewEntityMetadata()

	newEntMetaId := newEntMeta.GetId()
	newEntMetaIdUUID := newEntMeta.GetIdUUID()
	newEntMetaIdString := newEntMeta.GetIdString()
	newEntMetaCreatedAt := newEntMeta.GetCreatedAt()
	newEntMetaCreatedAtTime := newEntMeta.GetCreatedAtTime()
	newEntMetaUpdatedAt := newEntMeta.GetUpdatedAt()
	newEntMetaUpdatedAtTime := newEntMeta.GetUpdatedAtTime()
	newEntMetaDeletedFlag := newEntMeta.GetDeletedFlag()
	newEntMetaDeletedAt := newEntMeta.GetDeletedAt()
	newEntMetaDeletedAtTime := newEntMeta.GetDeletedAtTime()

	assert.NotNil(t, newEntMetaId)
	assert.IsType(t, baseentitee.EntityId{}, newEntMetaId)

	assert.NotNil(t, newEntMetaIdUUID)
	assert.IsType(t, newEntMetaIdUUID, newEntMetaIdUUID)

	assert.NotNil(t, newEntMetaIdString)
	assert.IsType(t, "", newEntMetaIdString)

	assert.NotNil(t, newEntMetaCreatedAt)
	assert.IsType(t, baseentitee.CreatedAt{}, newEntMetaCreatedAt)

	assert.NotNil(t, newEntMetaCreatedAtTime)
	assert.IsType(t, newEntMetaCreatedAtTime, newEntMetaCreatedAtTime)

	assert.NotNil(t, newEntMetaUpdatedAt)
	assert.IsType(t, baseentitee.UpdatedAt{}, newEntMetaUpdatedAt)

	assert.NotNil(t, newEntMetaUpdatedAtTime)
	assert.IsType(t, newEntMetaUpdatedAtTime, newEntMetaUpdatedAtTime)

	assert.NotNil(t, newEntMetaDeletedFlag)
	assert.IsType(t, baseentitee.DeletedFlag(false), newEntMetaDeletedFlag)

	assert.Nil(t, newEntMetaDeletedAt)
	assert.IsType(t, (*baseentitee.DeletedAt)(nil), newEntMetaDeletedAt)

	assert.NotNil(t, newEntMetaDeletedFlag)
	assert.IsType(t, baseentitee.DeletedFlag(true), newEntMetaDeletedFlag)

	assert.Nil(t, newEntMetaDeletedAt)
	assert.IsType(t, &baseentitee.DeletedAt{}, newEntMetaDeletedAt)

	assert.NotNil(t, newEntMetaId)
	assert.IsType(t, baseentitee.EntityId{}, newEntMetaId)

	assert.NotNil(t, newEntMetaCreatedAt)
	assert.IsType(t, baseentitee.CreatedAt{}, newEntMetaCreatedAt)

	assert.NotNil(t, newEntMetaUpdatedAt)
	assert.IsType(t, baseentitee.UpdatedAt{}, newEntMetaUpdatedAt)

	assert.NotNil(t, newEntMeta)
	assert.IsType(t, &baseentitee.EntityMetadata{}, newEntMeta)

	assert.WithinDuration(t, time.Now(), newEntMetaCreatedAtTime, time.Second)
	assert.WithinDuration(t, time.Now(), newEntMetaUpdatedAtTime, time.Second)

	assert.False(t, newEntMeta.IsDeleted())
	assert.Nil(t, newEntMetaDeletedAtTime)
}

func TestEntityMetadata_MarkDeleted(t *testing.T) {
	newEntMeta := baseentitee.NewEntityMetadata()

	newEntMeta.MarkDeleted()

	newEntMetaDeletedFlag := newEntMeta.GetDeletedFlag()
	newEntMetaDeletedAt := newEntMeta.GetDeletedAt()
	// newEntMetaId := newEntMeta.GetId()
	// newEntMetaCreatedAt := newEntMeta.GetCreatedAt()
	// newEntMetaUpdatedAt := newEntMeta.GetUpdatedAt()

	assert.NotNil(t, newEntMetaDeletedFlag)
	assert.IsType(t, baseentitee.DeletedFlag(true), newEntMetaDeletedFlag)
	assert.True(t, newEntMeta.IsDeleted())
	assert.NotNil(t, newEntMetaDeletedAt)
	assert.WithinDuration(t, time.Now(), *newEntMeta.GetDeletedAtTime(), time.Second)
}

func TestMapToEntityMetadata(t *testing.T) {
	createdAt := baseentitee.NewCreatedAt()
	updatedAt := baseentitee.NewUpdatedAt()
	deletedFlag := baseentitee.DeletedFlagFromBool(true)
	now := time.Now()
	deletedAt := baseentitee.DeletedAtFromTime(&now)
	newUuid := uuid.New()
	entId := baseentitee.EntityIdFromUUID(newUuid)

	mappedEntMeta := baseentitee.MapToEntityMetadata(
		entId,
		createdAt,
		updatedAt,
		deletedFlag,
		deletedAt,
	)

	assert.Equal(t, newUuid, mappedEntMeta.GetIdUUID())
}

func TestMapToEntityMetadataFromCommonTypes(t *testing.T) {
	createdAt := time.Now().Add(-2 * time.Hour)
	updatedAt := time.Now().Add(-1 * time.Hour)
	deletedFlag := true
	now := time.Now()
	deletedAt := &now
	newUuid := uuid.New()

	mappedEntMeta := baseentitee.MapToEntityMetadataFromCommonTypes(
		newUuid,
		createdAt,
		updatedAt,
		deletedFlag,
		deletedAt,
	)

	assert.Equal(t, newUuid, mappedEntMeta.GetIdUUID())
	assert.WithinDuration(t, createdAt, mappedEntMeta.GetCreatedAtTime(), time.Second)
	assert.WithinDuration(t, updatedAt, mappedEntMeta.GetUpdatedAtTime(), time.Second)
	assert.True(t, mappedEntMeta.IsDeleted())
	assert.WithinDuration(t, *deletedAt, *mappedEntMeta.GetDeletedAtTime(), time.Second)
}
