//nolint:dupl // these handlers are a bit duplicative but we don't want to abstract them further
package query

import (
	"context"

	"github.com/google/uuid"

	"github.com/cooperlutz/go-full/pkg/telemetree"
)

type FindOneInsuranceProvider struct {
	InsuranceProviderID string
}

type FindOneInsuranceProviderReadModel interface {
	FindOneInsuranceProvider(ctx context.Context, insuranceproviderId uuid.UUID) (InsuranceProvider, error)
}

type FindOneInsuranceProviderHandler struct {
	readModel FindOneInsuranceProviderReadModel
}

func NewFindOneInsuranceProviderHandler(
	readModel FindOneInsuranceProviderReadModel,
) FindOneInsuranceProviderHandler {
	return FindOneInsuranceProviderHandler{readModel: readModel}
}

func (h FindOneInsuranceProviderHandler) Handle(ctx context.Context, qry FindOneInsuranceProvider) (InsuranceProvider, error) {
	ctx, span := telemetree.AddSpan(ctx, "insuranceandclaims.app.query.find_one_insurance_provider.handle")
	defer span.End()

	insuranceprovider, err := h.readModel.FindOneInsuranceProvider(ctx, uuid.MustParse(qry.InsuranceProviderID))
	if err != nil {
		return InsuranceProvider{}, err
	}

	return insuranceprovider, nil
}
