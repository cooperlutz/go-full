package outbound

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/cooperlutz/go-full/internal/reporting/app/query"
	"github.com/cooperlutz/go-full/internal/reporting/domain/reporting"
	"github.com/cooperlutz/go-full/pkg/deebee"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// PostgresAdapter implements the reporting repository using Postgres as the data store.
type PostgresAdapter struct {
	Handler IQuerierReporting
}

// NewPostgresAdapter creates a new instance of PostgresAdapter.
func NewPostgresAdapter(db deebee.IDatabase) PostgresAdapter {
	return PostgresAdapter{
		Handler: NewQueriesWrapper(db),
	}
}

// FindReport retrieves a report by its ID and maps it to a query.ReportDefinition.
func (p PostgresAdapter) FindReport(ctx context.Context, name string) (query.ReportDefinition, error) {
	// ctx, span := telemetree.AddSpan(ctx, "reporting.adapters.outbound.postgres.findreport")
	// defer span.End()

	// report, err := p.GetReport(ctx, name)
	// if err != nil {
	// 	telemetree.RecordError(ctx, err)

	// 	return query.ReportDefinition{}, err
	// }

	// return mapEntityReportToQuery(report), nil
	return query.ReportDefinition{}, nil
}

// FindMetric retrieves a metric by its ID and maps it to a query.ReportMetric.
func (p PostgresAdapter) FindMetric(ctx context.Context, name reporting.MetricName) (query.ReportMetric, error) {
	ctx, span := telemetree.AddSpan(ctx, "reporting.adapters.outbound.postgres.find_metric")
	defer span.End()

	metric, err := p.GetMetric(ctx, name)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return query.ReportMetric{}, err
	}

	return mapEntityMetricToQuery(metric), nil
}

func (p PostgresAdapter) AddReport(ctx context.Context, report *reporting.Report) error {
	return nil
}

// GetReport retrieves a report by its name from the database.
func (p PostgresAdapter) GetReport(ctx context.Context, name string) (*reporting.Report, error) {
	// ctx, span := telemetree.AddSpan(ctx, "reporting.adapters.outbound.postgres.getreport")
	// defer span.End()

	// reportDef, err := p.Handler.GetReportByName(
	// 	ctx,
	// 	GetReportByNameParams{Name: name},
	// )
	// if err != nil {
	// 	telemetree.RecordError(ctx, err)

	// 	return nil, err
	// }
	return nil, nil
}

// UpdateReport updates an existing report in the database.
func (p PostgresAdapter) UpdateReport(
	ctx context.Context,
	name string,
	updateFn func(r *reporting.Report) (*reporting.Report, error),
) error {
	// ctx, span := telemetree.AddSpan(ctx, "reporting.adapters.outbound.postgres.updatereport")
	// defer span.End()

	// tx, err := p.Handler.Begin(ctx)
	// if err != nil {
	// 	return err
	// }

	// defer func() {
	// 	err = p.finishTransaction(ctx, err, tx)
	// }()

	// updatedReport, err := updateFn(report)
	// if err != nil {
	// 	telemetree.RecordError(ctx, err)

	// 	return err
	// }

	// dbReport := mapEntityReportToDB(updatedReport)
	// err = p.Handler.UpdateReport(ctx, UpdateReportParams(dbReport))

	// for _, question := range updatedReport.GetQuestions() {
	// 	dbQuestion := mapEntityQuestionToDB(question, updatedReport.GetIdUUID())

	// 	err = p.Handler.UpdateQuestion(ctx, UpdateQuestionParams(dbQuestion))
	// 	if err != nil {
	// 		telemetree.RecordError(ctx, err)

	// 		return err
	// 	}
	// }

	// if err != nil {
	// 	telemetree.RecordError(ctx, err)

	// 	return err
	// }
	return nil
}

func (p PostgresAdapter) AddMetric(ctx context.Context, metric *reporting.Metric) error {
	ctx, span := telemetree.AddSpan(ctx, "reporting.adapters.outbound.postgres.add_metric")
	defer span.End()

	dbMetric := mapEntityMetricToDb(metric)

	err := p.Handler.AddMetric(ctx, AddMetricParams(dbMetric))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

// GetMetric retrieves a metric by its name from the database.
func (p PostgresAdapter) GetMetric(ctx context.Context, name reporting.MetricName) (*reporting.Metric, error) {
	ctx, span := telemetree.AddSpan(ctx, "reporting.adapters.outbound.postgres.get_metric")
	defer span.End()

	metric, err := p.Handler.GetMetricByName(
		ctx,
		GetMetricByNameParams{Name: name.String()},
	)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return nil, err
	}

	return metric.toDomain()
}

// UpdateMetric updates an existing metric in the database.
func (p PostgresAdapter) UpdateMetric(
	ctx context.Context,
	name reporting.MetricName,
	updateFn func(r *reporting.Metric) (*reporting.Metric, error),
) error {
	ctx, span := telemetree.AddSpan(ctx, "reporting.adapters.outbound.postgres.update_metric")
	defer span.End()

	tx, err := p.Handler.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = p.finishTransaction(ctx, err, tx)
	}()

	metric, err := p.GetMetric(ctx, name)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	updatedMetric, err := updateFn(metric)
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	dbMetric := mapEntityMetricToDb(updatedMetric)

	err = p.Handler.UpdateMetric(ctx, UpdateMetricParams(dbMetric))
	if err != nil {
		telemetree.RecordError(ctx, err)

		return err
	}

	return nil
}

// finishTransaction commits or rolls back the transaction based on the error state.
func (p PostgresAdapter) finishTransaction(ctx context.Context, err error, tx pgx.Tx) error {
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			telemetree.RecordError(ctx, rollbackErr, "failed to rollback tx")

			return rollbackErr
		}

		return err
	} else {
		if commitErr := tx.Commit(ctx); commitErr != nil {
			telemetree.RecordError(ctx, commitErr, "failed to commit tx")

			return commitErr
		}

		return nil
	}
}
