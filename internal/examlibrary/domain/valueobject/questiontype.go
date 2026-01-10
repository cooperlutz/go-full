package valueobject

import "github.com/cooperlutz/go-full/internal/examlibrary/domain/exception"

type QuestionType int

const (
	QuestionMultipleChoice QuestionType = iota
	QuestionEssay
	QuestionShortAnswer
	QuestionUnknown
)

var questionName = map[QuestionType]string{ //nolint:gochecknoglobals // global is ok here
	QuestionMultipleChoice: "multiple_choice",
	QuestionEssay:          "essay",
	QuestionShortAnswer:    "short_answer",
	QuestionUnknown:        "unknown",
}

func (qt QuestionType) String() string {
	return questionName[qt]
}

func QuestionTypeFromString(s string) (QuestionType, error) {
	for qt, name := range questionName {
		if name == s {
			return qt, nil
		}
	}

	return QuestionUnknown, exception.ErrInvalidQuestionType{}
}
