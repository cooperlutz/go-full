package grading

// the values within the option are not relevant for grading each question type (e.g. multiple choice).
type GradeQuestionOption struct {
	Feedback string
	Points   int32
}

func (e *Exam) GradeQuestion(index int32, options GradeQuestionOption) (bool, error) {
	question := e.GetQuestionByIndex(index)

	err := question.gradeQuestion(options)
	if err != nil {
		return false, err
	}

	e.MarkUpdated()
	completed := e.checkIfGradingCompletedAndFinalize()

	return completed, nil
}

func (q *Question) gradeQuestion(options GradeQuestionOption) error {
	switch q.questionType {
	case QuestionMultipleChoice:
		gradeMultipleChoiceQuestion(q)
	case QuestionShortAnswer, QuestionEssay:
		q.feedback = &options.Feedback

		err := q.SetPointsReceived(options.Points)
		if err != nil {
			return err
		}
	}

	q.markAsGraded()
	q.MarkUpdated()

	return nil
}

type ErrPointsExceedPossiblePoints struct{}

func (e ErrPointsExceedPossiblePoints) Error() string {
	return "points received cannot exceed possible points"
}

func (q *Question) SetPointsReceived(points int32) error {
	if points > q.pointsPossible {
		return ErrPointsExceedPossiblePoints{}
	}

	q.pointsReceived = &points

	q.MarkUpdated()

	return nil
}

func gradeMultipleChoiceQuestion(q *Question) {
	var correct bool
	if q.providedAnswer == *q.correctAnswer {
		correct = true
		q.correctlyAnswered = &correct
		points := q.pointsPossible

		err := q.SetPointsReceived(points)
		if err != nil {
			// This should never happen since possiblePoints is the max
			panic(err)
		}
	} else {
		correct = false
		q.correctlyAnswered = &correct

		err := q.SetPointsReceived(int32(0))
		if err != nil {
			// This should never happen since 0 is less than possiblePoints
			panic(err)
		}
	}
}
