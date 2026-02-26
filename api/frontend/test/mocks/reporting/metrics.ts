import { http, HttpResponse } from "msw";
import { type Metric } from "~/reporting/services";

const metricsHandlers = [
  http.get<{ metricName: string }>(
    "/api/reporting/v1/metrics/:metricName",
    ({ params }) => {
      const { metricName } = params;
      return HttpResponse.json({
        metricName,
        metricValue: Math.floor(Math.random() * 100),
      } as Metric);
    },
  ),
];

export default metricsHandlers;
