package examination

import (
	"github.com/cooperlutz/go-full/pkg/utilitee"
)

type ErrNotAllQuestionsAnswered struct{}

func (e ErrNotAllQuestionsAnswered) Error() string {
	return "not all questions have been answered"
}

func (e *Exam) finishExam() error {
	now := utilitee.RightNow()
	e.state = StateCompleted
	e.completedAt = &now
	e.MarkUpdated()

	return nil
}

func (e *Exam) allQuestionsAnswered() bool {
	for _, question := range e.GetQuestions() {
		if !question.answered {
			return false
		}
	}

	return true
}
