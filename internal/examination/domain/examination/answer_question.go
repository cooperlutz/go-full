package examination

import "slices"

type ErrInvalidAnswer struct{}

func (e ErrInvalidAnswer) Error() string {
	return "invalid answer provided"
}

func (q *Question) setAnswer(answer string) error {
	valid := q.validateAnswer(answer)
	if !valid {
		return ErrInvalidAnswer{}
	}

	q.providedAnswer = &answer
	q.answered = true
	q.MarkUpdated()

	return nil
}

func (q *Question) validateAnswer(answer string) bool {
	switch q.questionType {
	case QuestionMultipleChoice:
		options := q.GetResponseOptions()
		if options != nil {
			return slices.Contains(*options, answer)
		}
	case QuestionShortAnswer:
		if answer == "" {
			return false
		}
	case QuestionEssay:
		if answer == "" {
			return false
		}
	default:
		return false
	}

	return true
}

func (e *Exam) AnswerQuestion(index int32, answer string) error {
	err := e.CheckTimeLimit()
	if err != nil {
		return err
	}

	question := e.GetQuestionByIndex(index)

	err = question.setAnswer(answer)
	if err != nil {
		return err
	}

	e.MarkUpdated()

	return nil
}
