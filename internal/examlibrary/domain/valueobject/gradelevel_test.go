package valueobject

import "testing"

func TestGradeLevelString(t *testing.T) {
	tests := []struct {
		name     string
		level    GradeLevel
		expected string
	}{
		{"Third", GradeLevelThird, "third"},
		{"Fourth", GradeLevelFourth, "fourth"},
		{"Fifth", GradeLevelFifth, "fifth"},
		{"Sixth", GradeLevelSixth, "sixth"},
		{"Seventh", GradeLevelSeventh, "seventh"},
		{"Eighth", GradeLevelEighth, "eighth"},
		{"Ninth", GradeLevelNinth, "ninth"},
		{"Tenth", GradeLevelTenth, "tenth"},
		{"Eleventh", GradeLevelEleventh, "eleventh"},
		{"Twelfth", GradeLevelTwelfth, "twelfth"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.level.String(); got != tt.expected {
				t.Errorf("GradeLevel.String() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestGradeLevelValues(t *testing.T) {
	tests := []struct {
		name     string
		level    GradeLevel
		expected int
	}{
		{"Third", GradeLevelThird, 3},
		{"Fourth", GradeLevelFourth, 4},
		{"Fifth", GradeLevelFifth, 5},
		{"Twelfth", GradeLevelTwelfth, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if int(tt.level) != tt.expected {
				t.Errorf("GradeLevel value = %d, want %d", tt.level, tt.expected)
			}
		})
	}
}

func TestGradeLevelStringInvalid(t *testing.T) {
	invalid := GradeLevel(99)
	if got := invalid.String(); got != "" {
		t.Errorf("GradeLevel.String() for invalid value = %q, want empty string", got)
	}
}
