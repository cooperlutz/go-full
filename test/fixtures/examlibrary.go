package fixtures

import (
	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/common"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/pkg/utilitee"
	"github.com/google/uuid"
)

func ptrQuestionType(qt server.QuestionType) *server.QuestionType { return &qt }

var (
	ValidUUID = uuid.UUID([16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
)

var (
	ValidApiExamQuestionMultipleChoice = server.ExamQuestion{
		Index:           utilitee.IntPtr(1),
		QuestionText:    utilitee.StrPtr("What animal is known to bark?"),
		QuestionType:    ptrQuestionType(server.MultipleChoice),
		PossiblePoints:  utilitee.IntPtr(5),
		CorrectAnswer:   utilitee.StrPtr("dog"),
		PossibleAnswers: &[]string{"dog", "cat", "bird", "fish"},
	}
	ValidAppExamQuestionMultipleChoice = common.NewExamQuestion(
		1,
		"What animal is known to bark?",
		"multiple-choice",
		5,
		utilitee.StrPtr("dog"),
		&[]string{"dog", "cat", "bird", "fish"},
	)
	ValidApiExamQuestionShortAnswer = server.ExamQuestion{
		Index:          utilitee.IntPtr(2),
		QuestionText:   utilitee.StrPtr("Describe photosynthesis"),
		QuestionType:   ptrQuestionType(server.ShortAnswer),
		PossiblePoints: utilitee.IntPtr(10),
		CorrectAnswer:  nil,
	}
	ValidAppExamQuestionShortAnswer = common.NewExamQuestion(
		2,
		"Describe photosynthesis",
		"short-answer",
		10,
		nil,
		nil,
	)
	ValidApiExamQuestionEssay = server.ExamQuestion{
		Index:          utilitee.IntPtr(3),
		QuestionText:   utilitee.StrPtr("Explain the theory of relativity"),
		QuestionType:   ptrQuestionType(server.Essay),
		PossiblePoints: utilitee.IntPtr(15),
		CorrectAnswer:  nil,
	}
	ValidAppExamQuestionEssay = common.NewExamQuestion(
		3,
		"Explain the theory of relativity",
		"essay",
		15,
		nil,
		nil,
	)
	ValidApiExamQuestions = []server.ExamQuestion{
		ValidApiExamQuestionMultipleChoice,
		ValidApiExamQuestionShortAnswer,
		ValidApiExamQuestionEssay,
	}
	ValidAppExamQuestions = []common.ExamQuestion{
		ValidAppExamQuestionMultipleChoice,
		ValidAppExamQuestionShortAnswer,
		ValidAppExamQuestionEssay,
	}
	ValidAppExamWithoutQuestions = query.ExamWithoutQuestions{
		ExamID:     "123e4567-e89b-12d3-a456-426614174000",
		Name:       "Sample Exam",
		GradeLevel: 10,
	}
	ValidAppExamsWithoutQuestions = []query.ExamWithoutQuestions{
		{
			ExamID:     "123e4567-e89b-12d3-a456-426614174000",
			Name:       "Sample Exam 1",
			GradeLevel: 10,
		},
		{
			ExamID:     "223e4567-e89b-12d3-a456-426614174001",
			Name:       "Sample Exam 2",
			GradeLevel: 11,
		},
	}
	ValidApiExamsMetadataList = []server.ExamMetadata{
		{
			Id:         utilitee.StrPtr("123e4567-e89b-12d3-a456-426614174000"),
			Name:       utilitee.StrPtr("Sample Exam 1"),
			GradeLevel: utilitee.IntPtr(10),
		},
		{
			Id:         utilitee.StrPtr("223e4567-e89b-12d3-a456-426614174001"),
			Name:       utilitee.StrPtr("Sample Exam 2"),
			GradeLevel: utilitee.IntPtr(11),
		},
	}
	ValidApiExamMetadata = server.ExamMetadata{
		Id:         &ValidAppExamWithoutQuestions.ExamID,
		Name:       &ValidAppExamWithoutQuestions.Name,
		GradeLevel: &ValidAppExamWithoutQuestions.GradeLevel,
	}
	ValidAppFindOneExamByIDResponse = query.FindOneExamByIDResponse{
		ExamID:     "123e4567-e89b-12d3-a456-426614174000",
		Name:       "Sample Exam",
		GradeLevel: 10,
		Questions: &[]common.ExamQuestion{
			common.NewExamQuestion(
				1,
				"What is 2 + 2?",
				"multiple-choice",
				5,
				utilitee.StrPtr("4"),
				&[]string{"3", "4", "5", "6"},
			),
		},
	}
	ValidApiExam = server.Exam{
		Id:         &ValidAppFindOneExamByIDResponse.ExamID,
		Name:       &ValidAppFindOneExamByIDResponse.Name,
		GradeLevel: &ValidAppFindOneExamByIDResponse.GradeLevel,
		Questions: &[]server.ExamQuestion{
			{
				Index:          utilitee.IntPtr(1),
				QuestionText:   utilitee.StrPtr("What is 2 + 2?"),
				QuestionType:   ptrQuestionType(server.MultipleChoice),
				PossiblePoints: utilitee.IntPtr(5),
				CorrectAnswer:  utilitee.StrPtr("4"),
				PossibleAnswers: &[]string{
					"3", "4", "5", "6",
				},
			},
		},
	}
	ValidAppCommandAddExamToLibraryResult = command.AddExamToLibraryResult{
		ExamID:     "123e4567-e89b-12d3-a456-426614174000",
		Name:       "Sample Exam",
		GradeLevel: 10,
		Questions: []common.ExamQuestion{
			common.NewExamQuestion(
				1,
				"What is 2 + 2?",
				"multiple-choice",
				5,
				utilitee.StrPtr("4"),
				&[]string{"3", "4", "5", "6"},
			),
		},
	}
)
