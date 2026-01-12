package usecase

import (
	"context"

	"github.com/cooperlutz/go-full/internal/examlibrary/app/command"
	"github.com/cooperlutz/go-full/internal/examlibrary/app/mapper"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// AddExamToLibrary adds a new exam to the library.
func (uc *examLibraryUseCase) AddExamToLibrary(ctx context.Context, cmd command.AddExamToLibrary) (command.AddExamToLibraryResult, error) {
	ctx, span := telemetree.AddSpan(ctx, "examlibrary.usecase.addexamtolibrary")
	defer span.End()

	entity, err := mapper.FromAppAddExamToLibraryToDomainExam(cmd)
	if err != nil {
		return command.AddExamToLibraryResult{}, err
	}

	if err := uc.Persist.SaveExam(ctx, entity); err != nil {
		return command.AddExamToLibraryResult{}, err
	}

	// emit any domain events that were raised during the entity's lifecycle
	domainEvents := entity.GetDomainEventsAndClear()
	if err := uc.emitEvents(domainEvents); err != nil {
		return command.AddExamToLibraryResult{}, err
	}

	result := mapper.FromDomainExamToAppAddExamToLibraryResult(entity)

	return result, nil
}
