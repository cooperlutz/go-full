package inbound

import (
	"context"

	"github.com/cooperlutz/go-full/internal/reporting/app"
	"github.com/cooperlutz/go-full/internal/reporting/app/query"
	"github.com/cooperlutz/go-full/pkg/telemetree"
)

// HttpServer represents the HTTP server for the Reporting module.
type HttpServer struct {
	app app.Application
}

// NewHttpServer creates a new HttpServer instance with the provided Reporting application.
func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

// StrictHandler returns a strict HTTP handler for the Reporting module.
func (h HttpServer) StrictHandler() ServerInterface {
	return NewStrictHandler(h, nil)
}

// (GET /v1/metrics/{metricName}).
func (h HttpServer) GetMetric(ctx context.Context, request GetMetricRequestObject) (GetMetricResponseObject, error) {
	ctx, span := telemetree.AddSpan(ctx, "reporting.adapters.inbound.http.getmetric")
	defer span.End()

	metric, err := h.app.Queries.FindMetric.Handle(ctx, query.FindMetric{
		Name: request.MetricName,
	})
	if err != nil {
		return nil, err
	}

	metricResponse := queryMetricToHttpMetric(metric)

	return GetMetric200JSONResponse(metricResponse), nil
}

// // (GET /v1/reports/{reportName}).
// func (h HttpServer) GetReport(ctx context.Context, request GetReportRequestObject) (GetReportResponseObject, error) {
// 	ctx, span := telemetree.AddSpan(ctx, "reporting.adapters.inbound.http.getreport")
// 	defer span.End()

// 	report, err := h.app.Queries.FindReport.Handle(ctx, query.FindReport{
// 		Name: request.ReportName,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	reportResponse := queryReportToHttpReport(report)

// 	return GetReport200JSONResponse(reportResponse), nil
// }

// func queryReportToHttpReport(e query.Report) Report {
// 	return Report{
// 		Name:             e.Name,
// 		Description:      e.Description,
// 		ReportComponents: queryReportComponentsToHttpReportComponents(e.Components),
// 	}
// }

// func queryReportComponentsToHttpReportComponents(components []query.ReportComponent) *[]ReportComponent {
// 	var httpComponents []ReportComponent

// 	for _, c := range components {
// 		switch c.Type {
// 		case "metric":
// 			val, ok := c.Value.(float64)
// 			if !ok {
// 				continue
// 			}

// 			httpComponents = append(httpComponents, ReportComponent{
// 				ComponentName:  c.Key,
// 				ComponentValue: &val,
// 			})
// 		}
// 	}

// 	return &httpComponents
// }

func queryMetricToHttpMetric(m query.ReportMetric) Metric {
	return Metric{
		MetricName:  m.Name,
		MetricValue: &m.Value,
	}
}
