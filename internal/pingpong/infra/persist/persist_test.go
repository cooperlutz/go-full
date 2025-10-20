package persist

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

// TestLoadConfigFromEnvVars tests the LoadConfigFromEnvVars function.
// Test Cases:
//
// 1. Success - All environment variables are set correctly.
//
// 2. Invalid DB_PORT - Non-integer value for DB_PORT, should fallback to default.
func TestNewPingPongPersistPostgresRepository(t *testing.T) {
	// mockPgConn := mocks.NewMockDBTX(t)
	// mockQuerier := mocks.NewMockQuerier(t)

	unitTests := []struct {
		name           string
		expectedReturn *pingPongPersistPostgresRepository
	}{
		{
			name:           "new ping pong postgres repository returns expected type and is not nil",
			expectedReturn: &pingPongPersistPostgresRepository{},
		},
	}
	for _, tt := range unitTests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewPingPongPostgresRepo(&pgxpool.Pool{})
			// Assert
			//
			assert.NotNil(t, repo)
			assert.IsType(t, tt.expectedReturn, repo)
		})
	}
}
