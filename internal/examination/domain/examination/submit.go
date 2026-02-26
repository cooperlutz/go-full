package examination

func (e *Exam) Submit() error {
	err := e.CheckTimeLimit()
	if err != nil {
		return err
	}

	if !e.allQuestionsAnswered() {
		return ErrNotAllQuestionsAnswered{}
	}

	err = e.finishExam()
	if err != nil {
		return err
	}

	return nil
}
