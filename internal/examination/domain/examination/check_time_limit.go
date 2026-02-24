package examination

import "github.com/cooperlutz/go-full/pkg/utilitee"

type ErrTimeLimitExceeded struct{}

func (e ErrTimeLimitExceeded) Error() string {
	return "time limit for exam has been exceeded"
}

func (e Exam) timeLimitExceeded() bool {
	if e.timeOfTimeLimit == nil {
		return false
	}

	return utilitee.RightNow().After(*e.timeOfTimeLimit)
}

func (e Exam) checkTimeLimit() error {
	if e.timeLimitExceeded() {
		err := e.finishExam()
		if err != nil {
			return err
		}

		return ErrTimeLimitExceeded{}
	}

	return nil
}
