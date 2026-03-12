//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneInsuranceClaim struct {
	InsuranceClaimID string
}

type FindOneInsuranceClaimReadModel interface {
	FindOneInsuranceClaim(ctx context.Context, insuranceclaimId uuid.UUID) (InsuranceClaim, error)
}

type FindOneInsuranceClaimHandler struct {
	readModel FindOneInsuranceClaimReadModel
}

func NewFindOneInsuranceClaimHandler(
	readModel FindOneInsuranceClaimReadModel,
) FindOneInsuranceClaimHandler {
	return FindOneInsuranceClaimHandler{readModel: readModel}
}

func (h FindOneInsuranceClaimHandler) Handle(ctx context.Context, qry FindOneInsuranceClaim) (InsuranceClaim, error) {
	ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.app.query.find_one_insurance_claim.handle")
	defer span.End()

	insuranceclaim, err := h.readModel.FindOneInsuranceClaim(ctx, uuid.MustParse(qry.InsuranceClaimID))
	if err != nil {
		return InsuranceClaim{}, err
	}

	return insuranceclaim, nil
}
