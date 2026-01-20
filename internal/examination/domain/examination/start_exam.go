package examination

import (
	"github.com/cooperlutz/go-full/pkg/utilitee"
)

type ErrExamAlreadyStarted struct{}

func (e ErrExamAlreadyStarted) Error() string {
	return "exam has already been started"
}

func (e *Exam) StartExam() error {
	if e.startedAt != nil {
		return ErrExamAlreadyStarted{}
	}

	now := utilitee.RightNow()
	e.startedAt = &now
	e.MarkUpdated()

	return nil
}
