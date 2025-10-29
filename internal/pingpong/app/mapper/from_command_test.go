package mapper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cooperlutz/go-full/internal/pingpong/app/command"
	"github.com/cooperlutz/go-full/internal/pingpong/app/mapper"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/exception"
)

func TestMapFromPingPongCommand(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		command             command.PingPongCommand
		wantEntityMsg       string
		wantEntityUpdatedAt time.Time
		wantEntityDeletedAt *time.Time
		wantEntityDeleted   bool
		wantErr             error
	}{
		{
			name: "maps command to entity",
			command: command.PingPongCommand{
				Message: "ping",
			},
			wantEntityMsg:       "ping",
			wantEntityUpdatedAt: time.Now(),
			wantEntityDeletedAt: nil,
			wantEntityDeleted:   false,
			wantErr:             nil,
		},
		{
			name: "an invalid command message returns a validation error",
			command: command.PingPongCommand{
				Message: "NONSENSE_MESSAGE",
			},
			wantEntityMsg:       "NONSENSE_MESSAGE",
			wantEntityUpdatedAt: time.Now(),
			wantEntityDeletedAt: nil,
			wantEntityDeleted:   false,
			wantErr:             exception.ErrPingPongMsgValidation{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			got, err := mapper.MapFromPingPongCommand(tt.command)
			// Assert
			// if we receive an error, we should have wanted one, and the error we received should match the one we wanted
			if err != nil {
				assert.NotNil(t, tt.wantErr)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
				assert.IsType(t, tt.wantErr, err)
				return
			}

			// the returning message should match the input message
			assert.Equal(t, tt.wantEntityMsg, got.Message)
			// the returning entity updatedAt timestamp should be close to now
			assert.WithinDuration(t, tt.wantEntityUpdatedAt, got.UpdatedAt, time.Second)
			// the returning entity createdAt timestamp should be close to now
			assert.WithinDuration(t, time.Now(), got.CreatedAt, time.Second)
			// the returning entity deletedAt timestamp should be nil
			assert.Nil(t, got.DeletedAt)
			// the returning entity deleted flag should be false
			assert.Equal(t, tt.wantEntityDeleted, got.Deleted)
		})
	}
}
