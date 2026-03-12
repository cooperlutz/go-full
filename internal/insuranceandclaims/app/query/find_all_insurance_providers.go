//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllInsuranceProvidersReadModel interface {
	FindAllInsuranceProviders(ctx context.Context) ([]InsuranceProvider, error)
}

type FindAllInsuranceProvidersHandler struct {
	readModel FindAllInsuranceProvidersReadModel
}

func NewFindAllInsuranceProvidersHandler(
	readModel FindAllInsuranceProvidersReadModel,
) FindAllInsuranceProvidersHandler {
	return FindAllInsuranceProvidersHandler{readModel: readModel}
}

func (h FindAllInsuranceProvidersHandler) Handle(ctx context.Context) ([]InsuranceProvider, error) {
	ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.app.query.find_all_insurance_providers.handle")
	defer span.End()

	exams, err := h.readModel.FindAllInsuranceProviders(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
