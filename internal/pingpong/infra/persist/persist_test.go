package persist

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

func TestNewPingPongPersistPostgresRepository(t *testing.T) {
	// Arrange
	unitTests := []struct {
		name           string
		expectedReturn *PingPongPersistPostgresRepository
	}{
		{
			name:           "new ping pong postgres repository returns expected type and is not nil",
			expectedReturn: &PingPongPersistPostgresRepository{},
		},
	}
	for _, tt := range unitTests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			repo := NewPingPongPostgresRepo(&pgxpool.Pool{})
			// Assert
			assert.NotNil(t, repo)
			assert.IsType(t, tt.expectedReturn, repo)
		})
	}
}
