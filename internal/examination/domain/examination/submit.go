package examination

import (
	"github.com/cooperlutz/go-full/pkg/utilitee"
)

func (e *Exam) Submit() error {
	now := utilitee.RightNow()
	e.completedAt = &now
	e.completed = true
	e.MarkUpdated()

	return nil
}
