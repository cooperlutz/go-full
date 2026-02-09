package grading

// unexported method to finalize grading and calculate total points received and grade
// this should only be called once all questions have been graded, and should not be called directly outside of the domain.
func (e *Exam) finalizeGrading() {
	e.CalculateTotalPointsReceived()
	e.gradingCompleted = true
	e.MarkUpdated()
}

// CheckIfGradingCompletedAndFinalize checks if grading is completed and finalizes grading if it is.
func (e *Exam) CheckIfGradingCompletedAndFinalize() bool {
	if e.IsCompleted() {
		return true
	}

	for _, q := range e.GetQuestions() {
		if !q.graded {
			return false
		}
	}

	e.finalizeGrading()

	return true
}
