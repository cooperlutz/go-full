package baseentitee_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

func TestNewEntityMetadata(t *testing.T) {
	// Arrange & Act
	newEntMeta := baseentitee.NewEntityMetadata()

	// Assert
	newEntMetaId := newEntMeta.GetId()
	assert.NotNil(t, newEntMetaId)
	assert.IsType(t, baseentitee.EntityId{}, newEntMetaId)

	// IdUUID checks
	newEntMetaIdUUID := newEntMeta.GetIdUUID()
	assert.NotNil(t, newEntMetaIdUUID)
	assert.IsType(t, newEntMetaIdUUID, newEntMetaIdUUID)

	// IdString checks
	newEntMetaIdString := newEntMeta.GetIdString()
	assert.NotNil(t, newEntMetaIdString)
	assert.IsType(t, "", newEntMetaIdString)

	// CreatedAt checks
	newEntMetaCreatedAt := newEntMeta.GetCreatedAt()
	assert.NotNil(t, newEntMetaCreatedAt)
	assert.IsType(t, baseentitee.CreatedAt{}, newEntMetaCreatedAt)

	// CreatedAtTime checks
	newEntMetaCreatedAtTime := newEntMeta.GetCreatedAtTime()
	assert.NotNil(t, newEntMetaCreatedAtTime)
	assert.IsType(t, newEntMetaCreatedAtTime, newEntMetaCreatedAtTime)
	assert.WithinDuration(t, time.Now(), newEntMetaCreatedAtTime, time.Second)

	// UpdatedAt checks
	newEntMetaUpdatedAt := newEntMeta.GetUpdatedAt()
	assert.NotNil(t, newEntMetaUpdatedAt)
	assert.IsType(t, baseentitee.UpdatedAt{}, newEntMetaUpdatedAt)

	// UpdatedAtTime checks
	newEntMetaUpdatedAtTime := newEntMeta.GetUpdatedAtTime()
	assert.NotNil(t, newEntMetaUpdatedAtTime)
	assert.IsType(t, newEntMetaUpdatedAtTime, newEntMetaUpdatedAtTime)
	assert.WithinDuration(t, time.Now(), newEntMetaUpdatedAtTime, time.Second)

	// DeletedFlag checks
	newEntMetaDeletedFlag := newEntMeta.GetDeletedFlag()
	assert.NotNil(t, newEntMetaDeletedFlag)
	assert.IsType(t, baseentitee.DeletedFlag(false), newEntMetaDeletedFlag)

	// DeletedAt checks
	newEntMetaDeletedAt := newEntMeta.GetDeletedAt()
	assert.Nil(t, newEntMetaDeletedAt)
	assert.IsType(t, &baseentitee.DeletedAt{}, newEntMetaDeletedAt)

	// DeletedAtTime checks
	assert.NotNil(t, newEntMeta)
	assert.IsType(t, &baseentitee.EntityMetadata{}, newEntMeta)

	// DeletedAt checks
	newEntMetaDeletedAtTime := newEntMeta.GetDeletedAtTime()
	assert.Nil(t, newEntMetaDeletedAtTime)

	// IsDeleted checks
	newEntMetaIsDeleted := newEntMeta.IsDeleted()
	assert.NotNil(t, newEntMetaIsDeleted)
	assert.IsType(t, false, newEntMetaIsDeleted)
}

func TestEntityMetadata_MarkDeleted(t *testing.T) {
	// Arrange
	newEntMeta := baseentitee.NewEntityMetadata()

	// Act
	newEntMeta.MarkDeleted()

	// Assert
	newEntMetaDeletedFlag := newEntMeta.GetDeletedFlag()
	newEntMetaDeletedAt := newEntMeta.GetDeletedAt()
	assert.NotNil(t, newEntMetaDeletedFlag)
	assert.IsType(t, baseentitee.DeletedFlag(true), newEntMetaDeletedFlag)
	assert.True(t, newEntMeta.IsDeleted())
	assert.NotNil(t, newEntMetaDeletedAt)
	assert.WithinDuration(t, time.Now(), *newEntMeta.GetDeletedAtTime(), time.Second)
}

func TestMapToEntityMetadata(t *testing.T) {
	// Arrange
	createdAt := baseentitee.NewCreatedAt()
	updatedAt := baseentitee.NewUpdatedAt()
	deletedFlag := baseentitee.DeletedFlagFromBool(true)
	now := time.Now()
	deletedAt := baseentitee.DeletedAtFromTime(&now)
	newUuid := uuid.New()
	entId := baseentitee.EntityIdFromUUID(newUuid)

	// Act
	mappedEntMeta := baseentitee.MapToEntityMetadata(
		entId,
		createdAt,
		updatedAt,
		deletedFlag,
		deletedAt,
	)

	// Assert
	assert.Equal(t, newUuid, mappedEntMeta.GetIdUUID())
}

func TestMapToEntityMetadataFromCommonTypes(t *testing.T) {
	// Arrange
	createdAt := time.Now().Add(-2 * time.Hour)
	updatedAt := time.Now().Add(-1 * time.Hour)
	deletedFlag := true
	now := time.Now()
	deletedAt := &now
	newUuid := uuid.New()

	// Act
	mappedEntMeta := baseentitee.MapToEntityMetadataFromCommonTypes(
		newUuid,
		createdAt,
		updatedAt,
		deletedFlag,
		deletedAt,
	)

	// Assert
	assert.Equal(t, newUuid, mappedEntMeta.GetIdUUID())
	assert.WithinDuration(t, createdAt, mappedEntMeta.GetCreatedAtTime(), time.Second)
	assert.WithinDuration(t, updatedAt, mappedEntMeta.GetUpdatedAtTime(), time.Second)
	assert.True(t, mappedEntMeta.IsDeleted())
	assert.WithinDuration(t, *deletedAt, *mappedEntMeta.GetDeletedAtTime(), time.Second)
}
