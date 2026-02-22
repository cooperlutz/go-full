package fixtures

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/cooperlutz/go-full/internal/examlibrary/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/common"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/query"
	"github.com/cooperlutz/go-full/internal/examlibrary/domain/entity"
	examlibrary_postgres "github.com/cooperlutz/go-full/internal/examlibrary/infra/persist/postgres"
	"github.com/cooperlutz/go-full/pkg/deebee/pgxutil"
)

func ptrQuestionType(qt server.QuestionType) *server.QuestionType { return new(qt) }

var (
	ValidUUID                          = uuid.UUID([16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	ValidApiExamQuestionMultipleChoice = server.ExamQuestion{
		Index:           new(1),
		QuestionText:    new("What animal is known to bark?"),
		QuestionType:    ptrQuestionType(server.MultipleChoice),
		PossiblePoints:  new(5),
		CorrectAnswer:   new("dog"),
		PossibleAnswers: &[]string{"dog", "cat", "bird", "fish"},
	}
	ValidAppExamQuestionMultipleChoice = common.NewExamQuestion(
		1,
		"What animal is known to bark?",
		"multiple-choice",
		5,
		new("dog"),
		&[]string{"dog", "cat", "bird", "fish"},
	)
	ValidApiExamQuestionShortAnswer = server.ExamQuestion{
		Index:          new(2),
		QuestionText:   new("Describe photosynthesis"),
		QuestionType:   ptrQuestionType(server.ShortAnswer),
		PossiblePoints: new(10),
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
		Index:          new(3),
		QuestionText:   new("Explain the theory of relativity"),
		QuestionType:   ptrQuestionType(server.Essay),
		PossiblePoints: new(15),
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
		ExamID:     "00000000-0000-0000-0000-000000000001",
		Name:       "Sample Exam",
		GradeLevel: 10,
		TimeLimit:  3600,
	}
	ValidAppExamWithQuestions = query.ExamWithQuestions{
		ExamID:     "00000000-0000-0000-0000-000000000001",
		Name:       "Sample Exam",
		GradeLevel: 10,
		TimeLimit:  3600,
		Questions:  ValidAppExamQuestions,
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
			Id:         new("123e4567-e89b-12d3-a456-426614174000"),
			Name:       new("Sample Exam 1"),
			GradeLevel: new(10),
		},
		{
			Id:         new("223e4567-e89b-12d3-a456-426614174001"),
			Name:       new("Sample Exam 2"),
			GradeLevel: new(11),
		},
	}
	ValidApiExamMetadata = server.ExamMetadata{
		Id:         &ValidAppExamWithoutQuestions.ExamID,
		Name:       &ValidAppExamWithoutQuestions.Name,
		GradeLevel: &ValidAppExamWithoutQuestions.GradeLevel,
		TimeLimit:  &ValidAppExamWithoutQuestions.TimeLimit,
	}
	ValidAppFindOneExamByIDResponse = query.FindOneExamByIDResponse{
		ExamID:     ValidMetadata.GetIdString(),
		Name:       "Sample Exam",
		GradeLevel: 10,
		TimeLimit:  3600,
		Questions:  &ValidAppExamQuestions,
	}
	ValidApiExam = server.Exam{
		Id:         &ValidAppFindOneExamByIDResponse.ExamID,
		Name:       &ValidAppFindOneExamByIDResponse.Name,
		GradeLevel: &ValidAppFindOneExamByIDResponse.GradeLevel,
		TimeLimit:  &ValidAppFindOneExamByIDResponse.TimeLimit,
		Questions:  &ValidApiExamQuestions,
	}
	ValidAppCommandAddExamToLibraryResult = command.AddExamToLibraryResult{
		ExamID:     ValidMetadata.GetIdString(),
		Name:       "Sample Exam",
		GradeLevel: 10,
		TimeLimit:  3600,
		Questions:  ValidAppExamQuestions,
	}
	ValidAppCommandAddExamToLibrary = command.NewAddExamToLibrary(
		"Sample Exam",
		10,
		3600,
		ValidAppExamQuestions,
	)
	ValidDomainExamQuestionMultipleChoice, _ = entity.MapToExamQuestion(
		ValidMetadata.GetIdUUID(),
		ValidMetadata.GetCreatedAtTime(),
		ValidMetadata.GetUpdatedAtTime(),
		ValidMetadata.IsDeleted(),
		ValidMetadata.GetDeletedAtTime(),
		"What animal is known to bark?",
		"multiple-choice",
		5,
		new("dog"),
		&[]string{"dog", "cat", "bird", "fish"},
		1,
	)
	ValidDomainExamQuestionShortAnswer, _ = entity.MapToExamQuestion(
		ValidMetadata.GetIdUUID(),
		ValidMetadata.GetCreatedAtTime(),
		ValidMetadata.GetUpdatedAtTime(),
		ValidMetadata.IsDeleted(),
		ValidMetadata.GetDeletedAtTime(),
		"Describe photosynthesis",
		"short-answer",
		10,
		nil,
		nil,
		2,
	)
	ValidDomainExamQuestionEssay, _ = entity.MapToExamQuestion(
		ValidMetadata.GetIdUUID(),
		ValidMetadata.GetCreatedAtTime(),
		ValidMetadata.GetUpdatedAtTime(),
		ValidMetadata.IsDeleted(),
		ValidMetadata.GetDeletedAtTime(),
		"Explain the theory of relativity",
		"essay",
		15,
		nil,
		nil,
		3,
	)
	ValidDomainExamQuestions = []entity.ExamQuestion{
		ValidDomainExamQuestionMultipleChoice,
		ValidDomainExamQuestionShortAnswer,
		ValidDomainExamQuestionEssay,
	}
	ValidDBExamQuestionShortAnswer = examlibrary_postgres.SaveExamQuestionParams{
		ExamQuestionID: pgxutil.UUIDToPgtypeUUID(ValidDomainExamQuestions[1].GetIdUUID()),
		CreatedAt:      pgxutil.TimeToTimestampz(&ValidDBExamCreatedAt),
		UpdatedAt:      pgxutil.TimeToTimestampz(&ValidDBExamUpdatedAt),
		DeletedAt:      pgxutil.TimeToTimestampz(ValidDomainExamQuestions[1].GetDeletedAtTime()),
		Deleted:        ValidDomainExamQuestions[1].IsDeleted(),
		ExamID:         pgxutil.UUIDToPgtypeUUID(ValidMetadata.GetIdUUID()),
		Index:          int32(ValidDomainExamQuestions[1].GetIndex()),
		QuestionText:   ValidDomainExamQuestions[1].GetQuestionText(),
		AnswerText:     pgtype.Text{Valid: false},
		PossiblePoints: int32(ValidDomainExamQuestions[1].GetPossiblePoints()),
		QuestionType:   ValidDomainExamQuestions[1].GetQuestionType().String(),
	}
	ValidDomainExam = entity.MapToExamEntity(
		ValidMetadata.GetIdUUID(),
		ValidMetadata.GetCreatedAtTime(),
		ValidMetadata.GetUpdatedAtTime(),
		ValidMetadata.IsDeleted(),
		ValidMetadata.GetDeletedAtTime(),
		"Sample Exam",
		10,
		3600,
		&ValidDomainExamQuestions,
	)
	ValidDBExamCreatedAt = ValidMetadata.GetCreatedAtTime()
	ValidDBExamUpdatedAt = ValidMetadata.GetUpdatedAtTime()
	ValidDBExam          = examlibrary_postgres.SaveExamParams{
		ExamID:     pgxutil.UUIDToPgtypeUUID(ValidMetadata.GetIdUUID()),
		CreatedAt:  pgxutil.TimeToTimestampz(&ValidDBExamCreatedAt),
		UpdatedAt:  pgxutil.TimeToTimestampz(&ValidDBExamUpdatedAt),
		DeletedAt:  pgxutil.TimeToTimestampz(ValidMetadata.GetDeletedAtTime()),
		Deleted:    ValidMetadata.IsDeleted(),
		Name:       "Sample Exam",
		GradeLevel: pgtype.Int4{Int32: 10, Valid: true},
		TimeLimit:  pgtype.Int8{Int64: 3600, Valid: true},
	}
	ValidDBExamLibraryExam = examlibrary_postgres.ExamLibraryExam{
		ExamID:    pgxutil.UUIDToPgtypeUUID(ValidMetadata.GetIdUUID()),
		CreatedAt: pgxutil.TimeToTimestampz(&ValidDBExamCreatedAt),
		UpdatedAt: pgxutil.TimeToTimestampz(&ValidDBExamUpdatedAt),
		Deleted:   ValidMetadata.IsDeleted(),
		DeletedAt: pgxutil.TimeToTimestampz(ValidMetadata.GetDeletedAtTime()),
		Name:      "Sample Exam",
		GradeLevel: pgtype.Int4{
			Int32: 10,
			Valid: true,
		},
		TimeLimit: pgtype.Int8{
			Int64: 3600,
			Valid: true,
		},
	}
	responseOpts                      = ValidDomainExamQuestions[0].GetResponseOptions()
	ValidDBExamQuestionMultipleChoice = examlibrary_postgres.ExamLibraryExamQuestion{
		ExamQuestionID: pgxutil.UUIDToPgtypeUUID(ValidDomainExamQuestions[0].GetIdUUID()),
		CreatedAt:      pgxutil.TimeToTimestampz(&ValidDBExamCreatedAt),
		UpdatedAt:      pgxutil.TimeToTimestampz(&ValidDBExamUpdatedAt),
		DeletedAt:      pgxutil.TimeToTimestampz(ValidDomainExamQuestions[0].GetDeletedAtTime()),
		Deleted:        ValidDomainExamQuestions[0].IsDeleted(),
		ExamID:         pgxutil.UUIDToPgtypeUUID(ValidMetadata.GetIdUUID()),
		Index:          int32(ValidDomainExamQuestions[0].GetIndex()),
		QuestionText:   ValidDomainExamQuestions[0].GetQuestionText(),
		AnswerText: pgtype.Text{
			String: *ValidDomainExamQuestions[0].GetCorrectAnswer(),
			Valid:  true,
		},
		PossiblePoints:  int32(ValidDomainExamQuestions[0].GetPossiblePoints()),
		QuestionType:    ValidDomainExamQuestions[0].GetQuestionType().String(),
		ResponseOptions: *responseOpts,
	}
	ValidDBExamQuestion = examlibrary_postgres.SaveExamQuestionParams{
		ExamQuestionID: pgxutil.UUIDToPgtypeUUID(ValidDomainExamQuestions[0].GetIdUUID()),
		CreatedAt:      pgxutil.TimeToTimestampz(&ValidDBExamCreatedAt),
		UpdatedAt:      pgxutil.TimeToTimestampz(&ValidDBExamUpdatedAt),
		DeletedAt:      pgxutil.TimeToTimestampz(ValidDomainExamQuestions[0].GetDeletedAtTime()),
		Deleted:        ValidDomainExamQuestions[0].IsDeleted(),
		ExamID:         pgxutil.UUIDToPgtypeUUID(ValidMetadata.GetIdUUID()),
		Index:          int32(ValidDomainExamQuestions[0].GetIndex()),
		QuestionText:   ValidDomainExamQuestions[0].GetQuestionText(),
		AnswerText:     pgxutil.StrToPgtypeText(ValidDomainExamQuestions[0].GetCorrectAnswer()),
		PossiblePoints: int32(ValidDomainExamQuestions[0].GetPossiblePoints()),
		QuestionType:   ValidDomainExamQuestions[0].GetQuestionType().String(),
	}
	ValidDomainExamWithNoQuestions = entity.MapToExamEntity(
		ValidMetadata.GetIdUUID(),
		ValidMetadata.GetCreatedAtTime(),
		ValidMetadata.GetUpdatedAtTime(),
		ValidMetadata.IsDeleted(),
		ValidMetadata.GetDeletedAtTime(),
		"Sample Exam",
		10,
		3600,
		nil,
	)
	ValidAppCommandAddExamToLibraryResultWithNoQuestions = command.NewAddExamToLibraryResult(
		ValidMetadata.GetIdString(),
		"Sample Exam",
		10,
		3600,
		[]common.ExamQuestion{},
	)
)
