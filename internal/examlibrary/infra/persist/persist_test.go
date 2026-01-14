package persist

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

func TestNewExamLibraryPersistPostgresRepository(t *testing.T) {
	// Arrange
	unitTests := []struct {
		name           string
		expectedReturn *examLibraryPersistPostgresRepository
	}{
		{
			name:           "new exam library postgres repository returns expected type and is not nil",
			expectedReturn: &examLibraryPersistPostgresRepository{},
		},
	}
	for _, tt := range unitTests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			repo := NewExamLibraryPostgresRepo(&pgxpool.Pool{})
			// Assert
			assert.NotNil(t, repo)
			assert.IsType(t, tt.expectedReturn, repo)
		})
	}
}
