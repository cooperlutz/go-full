package base_test

import (
	"testing"
	"time"

	"github.com/cooperlutz/go-full/pkg/base"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_NewEntityMetadata(t *testing.T) {

	newEntMeta := base.NewEntityMetadata()

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
	assert.IsType(t, base.EntityId{}, newEntMetaId)

	assert.NotNil(t, newEntMetaIdUUID)
	assert.IsType(t, newEntMetaIdUUID, newEntMetaIdUUID)

	assert.NotNil(t, newEntMetaIdString)
	assert.IsType(t, "", newEntMetaIdString)

	assert.NotNil(t, newEntMetaCreatedAt)
	assert.IsType(t, base.CreatedAt{}, newEntMetaCreatedAt)

	assert.NotNil(t, newEntMetaCreatedAtTime)
	assert.IsType(t, newEntMetaCreatedAtTime, newEntMetaCreatedAtTime)

	assert.NotNil(t, newEntMetaUpdatedAt)
	assert.IsType(t, base.UpdatedAt{}, newEntMetaUpdatedAt)

	assert.NotNil(t, newEntMetaUpdatedAtTime)
	assert.IsType(t, newEntMetaUpdatedAtTime, newEntMetaUpdatedAtTime)

	assert.NotNil(t, newEntMetaDeletedFlag)
	assert.IsType(t, base.DeletedFlag(false), newEntMetaDeletedFlag)

	assert.Nil(t, newEntMetaDeletedAt)
	assert.IsType(t, (*base.DeletedAt)(nil), newEntMetaDeletedAt)

	assert.NotNil(t, newEntMetaDeletedFlag)
	assert.IsType(t, base.DeletedFlag(true), newEntMetaDeletedFlag)

	assert.Nil(t, newEntMetaDeletedAt)
	assert.IsType(t, &base.DeletedAt{}, newEntMetaDeletedAt)

	assert.NotNil(t, newEntMetaId)
	assert.IsType(t, base.EntityId{}, newEntMetaId)

	assert.NotNil(t, newEntMetaCreatedAt)
	assert.IsType(t, base.CreatedAt{}, newEntMetaCreatedAt)

	assert.NotNil(t, newEntMetaUpdatedAt)
	assert.IsType(t, base.UpdatedAt{}, newEntMetaUpdatedAt)

	assert.NotNil(t, newEntMeta)
	assert.IsType(t, base.EntityMetadata{}, newEntMeta)

	assert.WithinDuration(t, time.Now(), newEntMetaCreatedAtTime, time.Second)
	assert.WithinDuration(t, time.Now(), newEntMetaUpdatedAtTime, time.Second)

	assert.False(t, newEntMeta.IsDeleted())
	assert.Nil(t, newEntMetaDeletedAtTime)
}

func Test_EntityMetadata_MarkDeleted(t *testing.T) {

	newEntMeta := base.NewEntityMetadata()

	newEntMeta.MarkDeleted()

	newEntMetaDeletedFlag := newEntMeta.GetDeletedFlag()
	newEntMetaDeletedAt := newEntMeta.GetDeletedAt()
	// newEntMetaId := newEntMeta.GetId()
	// newEntMetaCreatedAt := newEntMeta.GetCreatedAt()
	// newEntMetaUpdatedAt := newEntMeta.GetUpdatedAt()

	assert.NotNil(t, newEntMetaDeletedFlag)
	assert.IsType(t, base.DeletedFlag(true), newEntMetaDeletedFlag)
	assert.True(t, newEntMeta.IsDeleted())
	assert.NotNil(t, newEntMetaDeletedAt)
	assert.WithinDuration(t, time.Now(), *newEntMeta.GetDeletedAtTime(), time.Second)

}

func TestMapToEntityMetadata(t *testing.T) {
	createdAt := base.NewCreatedAt()
	updatedAt := base.NewUpdatedAt()
	deletedFlag := base.DeletedFlagFromBool(true)
	now := time.Now()
	deletedAt := base.DeletedAtFromTime(&now)
	newUuid := uuid.New()
	entId := base.EntityIdFromUUID(newUuid)

	mappedEntMeta := base.MapToEntityMetadata(
		entId,
		createdAt,
		updatedAt,
		deletedFlag,
		deletedAt,
	)

	assert.Equal(t, newUuid, mappedEntMeta.GetIdUUID())

}
