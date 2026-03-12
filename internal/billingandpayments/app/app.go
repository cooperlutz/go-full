package app

import (
	"github.com/cooperlutz/go-full/internal/billingandpayments/adapters/outbound"
	"github.com/cooperlutz/go-full/internal/billingandpayments/app/command"
	"github.com/cooperlutz/go-full/internal/billingandpayments/app/event"
	"github.com/cooperlutz/go-full/internal/billingandpayments/app/query"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/eeventdriven"
)

type Application struct {
	Commands Commands
	Queries  Queries
	Events   Events
}

type Commands struct {
	GenerateInvoice command.GenerateInvoiceHandler

	ApplyDiscountToInvoice command.ApplyDiscountToInvoiceHandler

	ProcessPayment command.ProcessPaymentHandler

	IssueRefund command.IssueRefundHandler

	VoidInvoice command.VoidInvoiceHandler
}

type Queries struct {
	FindAllInvoices query.FindAllInvoicesHandler
	FindOneInvoice  query.FindOneInvoiceHandler

	FindAllPayments query.FindAllPaymentsHandler
	FindOnePayment  query.FindOnePaymentHandler

	FindAllRefunds query.FindAllRefundsHandler
	FindOneRefund  query.FindOneRefundHandler
}

type Events struct {
	InvoiceGenerated event.InvoiceGeneratedHandler

	DiscountAppliedToInvoice event.DiscountAppliedToInvoiceHandler

	PaymentProcessed event.PaymentProcessedHandler

	ServicePaymentCompleted event.ServicePaymentCompletedHandler

	RefundIssued event.RefundIssuedHandler

	InvoiceVoided event.InvoiceVoidedHandler

	AppointmentCompleted event.AppointmentCompletedHandler
}

// NewApplication initializes the BillingAndPayments application with its dependencies.
func NewApplication( //nolint:funlen // it's fine
	pgconn deebee.IDatabase,
	pubSub eeventdriven.IPubSubEventProcessor,
) (Application, error) {
	invoiceRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	paymentRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	refundRepository := outbound.NewPostgresAdapter(
		pgconn,
	)

	app := Application{
		Commands: Commands{
			GenerateInvoice: command.NewGenerateInvoiceHandler(

				invoiceRepository,

				paymentRepository,

				refundRepository,
			),
			ApplyDiscountToInvoice: command.NewApplyDiscountToInvoiceHandler(

				invoiceRepository,

				paymentRepository,

				refundRepository,
			),
			ProcessPayment: command.NewProcessPaymentHandler(

				invoiceRepository,

				paymentRepository,

				refundRepository,
			),
			IssueRefund: command.NewIssueRefundHandler(

				invoiceRepository,

				paymentRepository,

				refundRepository,
			),
			VoidInvoice: command.NewVoidInvoiceHandler(

				invoiceRepository,

				paymentRepository,

				refundRepository,
			),
		},
		Queries: Queries{
			FindAllInvoices: query.NewFindAllInvoicesHandler(
				invoiceRepository,
			),
			FindOneInvoice: query.NewFindOneInvoiceHandler(
				invoiceRepository,
			),

			FindAllPayments: query.NewFindAllPaymentsHandler(
				paymentRepository,
			),
			FindOnePayment: query.NewFindOnePaymentHandler(
				paymentRepository,
			),

			FindAllRefunds: query.NewFindAllRefundsHandler(
				refundRepository,
			),
			FindOneRefund: query.NewFindOneRefundHandler(
				refundRepository,
			),
		},
		Events: Events{
			InvoiceGenerated: event.NewInvoiceGeneratedHandler(
				pubSub,
			),

			DiscountAppliedToInvoice: event.NewDiscountAppliedToInvoiceHandler(
				pubSub,
			),

			PaymentProcessed: event.NewPaymentProcessedHandler(
				pubSub,
			),

			ServicePaymentCompleted: event.NewServicePaymentCompletedHandler(
				pubSub,
			),

			RefundIssued: event.NewRefundIssuedHandler(
				pubSub,
			),

			InvoiceVoided: event.NewInvoiceVoidedHandler(
				pubSub,
			),

			AppointmentCompleted: event.NewAppointmentCompletedHandler(
				pubSub,
			),
		},
	}

	return app, nil
}
