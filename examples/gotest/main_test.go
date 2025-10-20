package gotest_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	exampletesting "github.com/cooperlutz/go-full/examples/gotest"
)

// TestLoadConfigFromEnvVars tests the LoadConfigFromEnvVars function.
// Test Cases:
//
// 1. Success - All environment variables are set correctly.
//
// 2. Invalid DB_PORT - Non-integer value for DB_PORT, should fallback to default.
func Test_COPYABLE_UNIT_TEST(t *testing.T) {
	unitTests := []struct {
		name           string
		input          *exampletesting.SampleStruct
		expectedReturn *exampletesting.SampleStruct
	}{
		{
			name: "a test case",
			input: &exampletesting.SampleStruct{
				Field1: "a silly string",
				Field2: 8675309,
			},
			expectedReturn: &exampletesting.SampleStruct{
				Field1: "a silly string",
				Field2: 8675309,
			},
		},
	}
	for _, tt := range unitTests {
		t.Run(tt.name, func(t *testing.T) {
			// Assert
			//
			assert.Equal(t, tt.input, tt.expectedReturn)
			// The fields are equal
			assert.Equal(t, tt.expectedReturn.Field1, tt.input.Field1)
			assert.Equal(t, tt.expectedReturn.Field2, tt.input.Field2)
			// Field types is equal to the expected type
			assert.IsType(t, "", tt.input.Field1)
			assert.IsType(t, 0, tt.expectedReturn.Field2)

			gotBoolVal, gotErr := tt.input.AMethod()
			// No error is returned
			assert.NoError(t, gotErr)
			// The bool return value is true
			assert.True(t, gotBoolVal)
		})
	}
}

// TestLoadConfigFromEnvVars tests the LoadConfigFromEnvVars function.
// Test Cases:
//
// 1. Method returns an error when Field1 is "METHOD_ERROR".
//
// 2. Method returns does not return an error when Field1 is anything else.
func Test_COPYABLE_UNIT_TEST_METHOD(t *testing.T) {
	// Arrange
	unitTests := []struct {
		name              string
		input             *exampletesting.SampleStruct
		expectedReturnVal bool
		expectedErr       error
		wantErr           bool
	}{
		{
			name: "method should return an error if Field1 is METHOD_ERROR",
			input: &exampletesting.SampleStruct{
				Field1: "METHOD_ERROR",
				Field2: 8675309,
			},
			wantErr:           true,
			expectedReturnVal: false,
			expectedErr:       &exampletesting.ErrSample{Sample: "YOU GOT AN ERROR"},
		},
		{
			name: "method should NOT return an error if Field1 is not METHOD_ERROR",
			input: &exampletesting.SampleStruct{
				Field1: "no error here",
				Field2: 8675309,
			},
			wantErr:           false,
			expectedReturnVal: true,
			expectedErr:       nil,
		},
	}
	for _, tt := range unitTests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			gotReturnVal, gotErr := tt.input.AMethod()

			// Assert
			if tt.wantErr {
				// An error is returned
				assert.Error(t, gotErr)
				// The error is the expected error
				assert.Equal(t, tt.expectedErr, gotErr)
			} else {
				// No error is returned
				assert.NoError(t, gotErr)
				// The return value is the expected value
				assert.Equal(t, tt.expectedReturnVal, gotReturnVal)
			}
		})
	}
}
