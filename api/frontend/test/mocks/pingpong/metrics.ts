import { http, HttpResponse } from 'msw'

const metricsHandlers = [
    http.get('/pingpong/api/v1/metrics/dailyDistribution', () => {
      return HttpResponse.json({
  "dimensionKeys": [
    "2025-10-26 00:00:00 +0000 UTC",
    "2025-11-01 00:00:00 +0000 UTC",
    "2025-11-05 00:00:00 +0000 UTC"
  ],
  "dimensionValues": [10, 2, 2]
})
    }),

    http.get('/pingpong/api/v1/metrics/totalPingPongs', () => {
      return HttpResponse.json(14)
    }),

    http.get('/pingpong/api/v1/metrics/totalPings', () => {
      return HttpResponse.json(7)
    }),

    http.get('/pingpong/api/v1/metrics/totalPongs', () => {
      return HttpResponse.json(7)
    }),
];

export default metricsHandlers;