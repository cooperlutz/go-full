package valueobject

type GradeLevel int

const (
	GradeLevelThird GradeLevel = iota + 3
	GradeLevelFourth
	GradeLevelFifth
	GradeLevelSixth
	GradeLevelSeventh
	GradeLevelEighth
	GradeLevelNinth
	GradeLevelTenth
	GradeLevelEleventh
	GradeLevelTwelfth
)

var gradeLevelName = map[GradeLevel]string{ //nolint:gochecknoglobals // global is ok here
	GradeLevelThird:    "third",
	GradeLevelFourth:   "fourth",
	GradeLevelFifth:    "fifth",
	GradeLevelSixth:    "sixth",
	GradeLevelSeventh:  "seventh",
	GradeLevelEighth:   "eighth",
	GradeLevelNinth:    "ninth",
	GradeLevelTenth:    "tenth",
	GradeLevelEleventh: "eleventh",
	GradeLevelTwelfth:  "twelfth",
}

func (gl GradeLevel) String() string {
	return gradeLevelName[gl]
}

func (gl GradeLevel) Int() int {
	return int(gl)
}
