package app

import (
	"github.com/cooperlutz/go-full/internal/insuranceandclaims/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/insuranceandclaims/app/command"
	"github.com/cooperlutz/go-full/internal/insuranceandclaims/app/event"
	"github.com/cooperlutz/go-full/internal/insuranceandclaims/app/query"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type Application struct {
	Commands Commands
	Queries  Queries
	Events   Events
}

type Commands struct {
	RegisterInsuranceProvider command.RegisterInsuranceProviderHandler

	SubmitInsuranceClaim command.SubmitInsuranceClaimHandler

	ApproveInsuranceClaim command.ApproveInsuranceClaimHandler

	RejectInsuranceClaim command.RejectInsuranceClaimHandler
}

type Queries struct {
	FindAllInsuranceProviders query.FindAllInsuranceProvidersHandler
	FindOneInsuranceProvider  query.FindOneInsuranceProviderHandler

	FindAllInsuranceClaims query.FindAllInsuranceClaimsHandler
	FindOneInsuranceClaim  query.FindOneInsuranceClaimHandler
}

type Events struct {
	InsuranceProviderRegistered event.InsuranceProviderRegisteredHandler

	InsuranceClaimSubmitted event.InsuranceClaimSubmittedHandler

	InsuranceClaimApproved event.InsuranceClaimApprovedHandler

	InsuranceClaimRejected event.InsuranceClaimRejectedHandler
}

// NewApplication initializes the InsuranceAndClaims application with its dependencies.
func NewApplication( //nolint:funlen // it's fine
	pgconn deebee.IDatabase,
	pubSub eeventdriven.IPubSubEventProcessor,
) (Application, error) {
	insuranceproviderRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	insuranceclaimRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	app := Application{
		Commands: Commands{
			RegisterInsuranceProvider: command.NewRegisterInsuranceProviderHandler(

				insuranceproviderRepository,

				insuranceclaimRepository,
			),
			SubmitInsuranceClaim: command.NewSubmitInsuranceClaimHandler(

				insuranceproviderRepository,

				insuranceclaimRepository,
			),
			ApproveInsuranceClaim: command.NewApproveInsuranceClaimHandler(

				insuranceproviderRepository,

				insuranceclaimRepository,
			),
			RejectInsuranceClaim: command.NewRejectInsuranceClaimHandler(

				insuranceproviderRepository,

				insuranceclaimRepository,
			),
		},
		Queries: Queries{
			FindAllInsuranceProviders: query.NewFindAllInsuranceProvidersHandler(
				insuranceproviderRepository,
			),
			FindOneInsuranceProvider: query.NewFindOneInsuranceProviderHandler(
				insuranceproviderRepository,
			),

			FindAllInsuranceClaims: query.NewFindAllInsuranceClaimsHandler(
				insuranceclaimRepository,
			),
			FindOneInsuranceClaim: query.NewFindOneInsuranceClaimHandler(
				insuranceclaimRepository,
			),
		},
		Events: Events{
			InsuranceProviderRegistered: event.NewInsuranceProviderRegisteredHandler(
				pubSub,
			),

			InsuranceClaimSubmitted: event.NewInsuranceClaimSubmittedHandler(
				pubSub,
			),

			InsuranceClaimApproved: event.NewInsuranceClaimApprovedHandler(
				pubSub,
			),

			InsuranceClaimRejected: event.NewInsuranceClaimRejectedHandler(
				pubSub,
			),
		},
	}

	return app, nil
}
