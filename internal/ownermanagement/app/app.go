package app

import (
	"github.com/cooperlutz/go-full/internal/ownermanagement/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/ownermanagement/app/command"
	"github.com/cooperlutz/go-full/internal/ownermanagement/app/event"
	"github.com/cooperlutz/go-full/internal/ownermanagement/app/query"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type Application struct {
	Commands Commands
	Queries  Queries
	Events   Events
}

type Commands struct {
	RegisterOwner command.RegisterOwnerHandler

	UpdateOwnerProfile command.UpdateOwnerProfileHandler

	EnrollInLoyaltyProgram command.EnrollInLoyaltyProgramHandler

	AwardLoyaltyPoints command.AwardLoyaltyPointsHandler

	RedeemLoyaltyPoints command.RedeemLoyaltyPointsHandler

	DeactivateOwner command.DeactivateOwnerHandler
}

type Queries struct {
	FindAllOwners query.FindAllOwnersHandler
	FindOneOwner  query.FindOneOwnerHandler

	FindAllLoyaltyAccounts query.FindAllLoyaltyAccountsHandler
	FindOneLoyaltyAccount  query.FindOneLoyaltyAccountHandler
}

type Events struct {
	OwnerRegistered event.OwnerRegisteredHandler

	OwnerProfileUpdated event.OwnerProfileUpdatedHandler

	OwnerEnrolledInLoyaltyProgram event.OwnerEnrolledInLoyaltyProgramHandler

	LoyaltyPointsAwarded event.LoyaltyPointsAwardedHandler

	LoyaltyPointsRedeemed event.LoyaltyPointsRedeemedHandler

	OwnerDeactivated event.OwnerDeactivatedHandler

	ServicePaymentCompleted event.ServicePaymentCompletedHandler
}

// NewApplication initializes the OwnerManagement application with its dependencies.
func NewApplication( //nolint:funlen // it's fine
	pgconn deebee.IDatabase,
	pubSub eeventdriven.IPubSubEventProcessor,
) (Application, error) {
	ownerRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	loyaltyaccountRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	app := Application{
		Commands: Commands{
			RegisterOwner: command.NewRegisterOwnerHandler(

				ownerRepository,

				loyaltyaccountRepository,
			),
			UpdateOwnerProfile: command.NewUpdateOwnerProfileHandler(

				ownerRepository,

				loyaltyaccountRepository,
			),
			EnrollInLoyaltyProgram: command.NewEnrollInLoyaltyProgramHandler(

				ownerRepository,

				loyaltyaccountRepository,
			),
			AwardLoyaltyPoints: command.NewAwardLoyaltyPointsHandler(

				ownerRepository,

				loyaltyaccountRepository,
			),
			RedeemLoyaltyPoints: command.NewRedeemLoyaltyPointsHandler(

				ownerRepository,

				loyaltyaccountRepository,
			),
			DeactivateOwner: command.NewDeactivateOwnerHandler(

				ownerRepository,

				loyaltyaccountRepository,
			),
		},
		Queries: Queries{
			FindAllOwners: query.NewFindAllOwnersHandler(
				ownerRepository,
			),
			FindOneOwner: query.NewFindOneOwnerHandler(
				ownerRepository,
			),

			FindAllLoyaltyAccounts: query.NewFindAllLoyaltyAccountsHandler(
				loyaltyaccountRepository,
			),
			FindOneLoyaltyAccount: query.NewFindOneLoyaltyAccountHandler(
				loyaltyaccountRepository,
			),
		},
		Events: Events{
			OwnerRegistered: event.NewOwnerRegisteredHandler(
				pubSub,
			),

			OwnerProfileUpdated: event.NewOwnerProfileUpdatedHandler(
				pubSub,
			),

			OwnerEnrolledInLoyaltyProgram: event.NewOwnerEnrolledInLoyaltyProgramHandler(
				pubSub,
			),

			LoyaltyPointsAwarded: event.NewLoyaltyPointsAwardedHandler(
				pubSub,
			),

			LoyaltyPointsRedeemed: event.NewLoyaltyPointsRedeemedHandler(
				pubSub,
			),

			OwnerDeactivated: event.NewOwnerDeactivatedHandler(
				pubSub,
			),

			ServicePaymentCompleted: event.NewServicePaymentCompletedHandler(
				pubSub,
			),
		},
	}

	return app, nil
}
