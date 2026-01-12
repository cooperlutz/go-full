package valueobject

import "github.com/cooperlutz/go-full/internal/examlibrary/domain/exception"

// QuestionType represents the type of a question in an exam.
// It is defined as an enumeration.
type QuestionType int

const (
	QuestionMultipleChoice QuestionType = iota
	QuestionEssay
	QuestionShortAnswer
	QuestionUnknown
)

var questionName = map[QuestionType]string{ //nolint:gochecknoglobals // global is ok here
	QuestionMultipleChoice: "multiple-choice",
	QuestionEssay:          "essay",
	QuestionShortAnswer:    "short-answer",
	QuestionUnknown:        "unknown",
}

func (qt QuestionType) String() string {
	return questionName[qt]
}

func (qt QuestionType) Int() int {
	return int(qt)
}

func QuestionTypeFromString(s string) (QuestionType, error) {
	for qt, name := range questionName {
		if name == s {
			return qt, nil
		}
	}

	return QuestionUnknown, exception.ErrInvalidQuestionType{}
}
