package grading

type ErrExamGradingAlreadyInProgress struct{}

func (e ErrExamGradingAlreadyInProgress) Error() string {
	return "exam is already being graded"
}

type ErrExamGradingAlreadyCompleted struct{}

func (e ErrExamGradingAlreadyCompleted) Error() string {
	return "exam grading is already completed"
}

func (e *Exam) startGrading() {
	switch e.GetState() {
	case StateNotStarted:
		e.state = StateInProgress
		e.MarkUpdated()
	case StateCompleted, StateInProgress:
		return
	}
}
