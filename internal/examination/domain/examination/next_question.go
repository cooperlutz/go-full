package examination

type ErrNoMoreQuestions struct{}

func (e ErrNoMoreQuestions) Error() string {
	return "no more questions available"
}

func (e *Exam) NextQuestion(currentIndex int32) (*Question, error) {
	if currentIndex >= e.NumberOfQuestions() {
		return nil, ErrNoMoreQuestions{}
	}

	return e.GetQuestionByIndex(currentIndex + 1), nil
}
