package fixtures

import (
	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/common"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/valueobject"
	"github.com/cooperlutz/go-full/pkg/utilitee"
)

func ptrQuestionType(qt server.QuestionType) *server.QuestionType { return &qt }

var ValidUUID = uuid.UUID([16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

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
		ExamID:     ValidMetadata.GetIdString(),
		Name:       "Sample Exam",
		GradeLevel: 10,
		Questions:  &ValidAppExamQuestions,
	}
	ValidApiExam = server.Exam{
		Id:         &ValidAppFindOneExamByIDResponse.ExamID,
		Name:       &ValidAppFindOneExamByIDResponse.Name,
		GradeLevel: &ValidAppFindOneExamByIDResponse.GradeLevel,
		Questions:  &ValidApiExamQuestions,
	}
	ValidAppCommandAddExamToLibraryResult = command.AddExamToLibraryResult{
		ExamID:     ValidMetadata.GetIdString(),
		Name:       "Sample Exam",
		GradeLevel: 10,
		Questions:  ValidAppExamQuestions,
	}
	ValidAppCommandAddExamToLibrary = command.NewAddExamToLibrary(
		"Sample Exam",
		10,
		ValidAppExamQuestions,
	)
	ValidDomainExamQuestions = []entity.ExamQuestion{
		entity.MapToExamQuestion(
			ValidMetadata.GetIdUUID(),
			ValidMetadata.GetCreatedAtTime(),
			ValidMetadata.GetUpdatedAtTime(),
			ValidMetadata.IsDeleted(),
			ValidMetadata.GetDeletedAtTime(),
			"What animal is known to bark?",
			valueobject.QuestionMultipleChoice,
			5,
			utilitee.StrPtr("dog"),
			&[]string{"dog", "cat", "bird", "fish"},
			1,
		),
		entity.MapToExamQuestion(
			ValidMetadata.GetIdUUID(),
			ValidMetadata.GetCreatedAtTime(),
			ValidMetadata.GetUpdatedAtTime(),
			ValidMetadata.IsDeleted(),
			ValidMetadata.GetDeletedAtTime(),
			"Describe photosynthesis",
			valueobject.QuestionShortAnswer,
			10,
			nil,
			nil,
			2,
		),
		entity.MapToExamQuestion(
			ValidMetadata.GetIdUUID(),
			ValidMetadata.GetCreatedAtTime(),
			ValidMetadata.GetUpdatedAtTime(),
			ValidMetadata.IsDeleted(),
			ValidMetadata.GetDeletedAtTime(),
			"Explain the theory of relativity",
			valueobject.QuestionEssay,
			15,
			nil,
			nil,
			3,
		),
	}
	ValidDomainExam = entity.MapToExamEntity(
		ValidMetadata.GetIdUUID(),
		ValidMetadata.GetCreatedAtTime(),
		ValidMetadata.GetUpdatedAtTime(),
		ValidMetadata.IsDeleted(),
		ValidMetadata.GetDeletedAtTime(),
		"Sample Exam",
		10,
		&ValidDomainExamQuestions,
	)
	ValidDomainExamWithNoQuestions = entity.MapToExamEntity(
		ValidMetadata.GetIdUUID(),
		ValidMetadata.GetCreatedAtTime(),
		ValidMetadata.GetUpdatedAtTime(),
		ValidMetadata.IsDeleted(),
		ValidMetadata.GetDeletedAtTime(),
		"Sample Exam",
		10,
		nil,
	)
	ValidAppCommandAddExamToLibraryResultWithNoQuestions = command.NewAddExamToLibraryResult(
		ValidMetadata.GetIdString(),
		"Sample Exam",
		10,
		[]common.ExamQuestion{},
	)
)
