//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindAllInsuranceClaimsReadModel interface {
	FindAllInsuranceClaims(ctx context.Context) ([]InsuranceClaim, error)
}

type FindAllInsuranceClaimsHandler struct {
	readModel FindAllInsuranceClaimsReadModel
}

func NewFindAllInsuranceClaimsHandler(
	readModel FindAllInsuranceClaimsReadModel,
) FindAllInsuranceClaimsHandler {
	return FindAllInsuranceClaimsHandler{readModel: readModel}
}

func (h FindAllInsuranceClaimsHandler) Handle(ctx context.Context) ([]InsuranceClaim, error) {
	ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.app.query.find_all_insurance_claims.handle")
	defer span.End()

	exams, err := h.readModel.FindAllInsuranceClaims(ctx)
	if err != nil {
		return nil, err
	}

	return exams, nil
}
