package examination

import "github.com/cooperlutz/go-full/pkg/utilitee"

func (e Exam) AnsweredQuestionsCount() int32 {
	var count int32

	for _, q := range e.questions {
		if q.IsAnswered() {
			count++
		}
	}

	return count
}

func (e Exam) NumberOfQuestions() int32 {
	length := len(e.questions)

	return utilitee.SafeIntToInt32(&length)
}
