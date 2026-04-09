package fixtures

import (
	"time"

	"github.com/google/uuid"

	rest_api_V1 "github.com/cooperlutz/go-full/internal/pingpong/api/rest/v1/server"
	"github.com/cooperlutz/go-full/internal/pingpong/domain/entity"
	"github.com/cooperlutz/go-full/pkg/baseentitee"
)

var (
	// This is NOT an in use private key and should not be used in production. It's only for testing purposes.
	ValidUnusedBase64EncodedPrivateKey = "LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JSUV2UUlCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktjd2dnU2pBZ0VBQW9JQkFRQ242cmFOdHdBVXh0amoKSVlwa3RXTDR3WktmOWE2czg1NHZ4cFdhQlZwVzY4U05zRWRCZEIwRmR4RTQ3NThsTUJHWk5uNmxaMGtaK3BJawpKZkFXRGszQlgzcU5ONWdpRFc5TVVsL05PNHFsRnhla090MG11eFFWenpLM0VMS2dmK0hwZ01BOG5BNzYvYml6Ckp0aisxUkVnNGRRSHlMT3hjWElEMEMyVm5ORnVGN3QvSXBjZjFsSStsV3NzMkxKbkF5MXdKVlZ5VzAvV3F5dmQKVm0rWU43c0lZU2RDK0piWW5XdSt6QjBSbUY2WDRKSzBRVDBOeER3R3A1ajJSQXNlYi9EYm5JdzZCZmtNTXc4dwpyOWI1dVJkTVZBUHdoMXpaRmRTSGhlcEJoYTZVQVVVMWZTaHVkZDRJcnBiVklaVmkxUjYwWmpaVElaNGVpME9LCmRSL2pxRXVCQWdNQkFBRUNnZ0VBTnNxL0FPVXJFSnlHSHk5Tm9ta1J2NHdpbFRHczZ1Q2Z5SVVyUWRQblU1K2gKUm42ejhhYTY1WktodXd0dHljdXlvbVVHMUdBc21WNkxHMktKd2hpdXhKK3NOTEprWVJBek5odFhDWGdaV0o5OQpabWlTeTFXV1N6dU9GL0JKNlZkeFZ5bGRFU24zTlcxZlBJNDdsb1BlRC9YQW96MlpiUmhRbE8vMFAvS280QncwClpWN1FNTUhxaFJoWmhLOGVNYWZmVDdUdjdzRGhyb2daTitXR3VFVUlRZTloWkk1NTRxaTZuWDdMTHJ2RURVRGsKZ2RFMURNWXhYVm1PM1lpWHRmanhnV1plSlU1empjZzhEOXErWVBPV3NEU3hPOVpzUmMzTjlKV1F2bDViemIwegp0MmxhMEhnRHNSZTVpUjBuSzdLRHlGVXZVZjNvMWF5M2YxYW5Rc1p6WlFLQmdRRGxHOURmSUtHVUlYMjlabHNMCkNaaVhtRWRENEJYRFZRS2xTR1VXTlVnak1zL1NrM3JkbStYd2ppeXE0b3BVNVI1a3hzZklrUW1DU0c5Rk5hMkYKa2NmQlpjZ0lHdXJzNC9IWGkyVzdEVWRYQnlHSHVWa0RPQzV2VTdOc0thNmx5dHlYTXZLRjBvT0Z3bnFmRlNwSwpySlZLMFB2WlRQQWJUUmFqSVc4akdSWW1sd0tCZ1FDN29EbW1lZnE2VVpaQXZqWVVvSU5YZmNWTEFBQUtPNmVNCkRYaE5qTEdaMVhQTFhla0d3b0ppOGlwS0pYMlpLSm9PTWlyaEFTK0plQzIzNEcvMXkza1Y4aW1jaU1kUDNNckUKSDliS0lSMytoYUh4SjJ0YVlVYnRSYkJQdDkzLzE1cXNLTGRiTGlhUGVLVFNxbVhnYVVGdEVqajJrdjVmZVFmMgp0V1dJYW42NXB3S0JnUUN6T2VOWDZXcjlPb1htMkY5NGJ2VlVBZWdlRFBwNVo5TStBVGd5KzFPSFZZU0dNcUQzCkk3SHBVUTlZVHdmd0NaeVJNWCt3UXVXeGFZRENINCt5NGF2bWV2elVlbG8xSHN4ZDFjcXJYV1Bsak1xS0psQ2YKMkI5Ykw1czRqaEwxMTVCaEo1WFpZaWxKUmk0dXJKdjg5cTJEYmFEWENQejhIbjRLOTJpOWlDNXYyUUtCZ0E4agpsb05BUkdMVVFuTm01YVkwYldTWjJWbDczb1QwMjBnSnJqTlRydURpd0QwZ2pyVGs3UEJlNmRxaHU4aW41Z1pQCk5mYUJ5ZndsbmtxLzZVakQ2amxER09od2dPU0l4RE9lL2czQkxsSmpKZlgzWUVvaW56NTE2UXdGb2Z3S3dZZ3kKb05qU21IbnJHaUdGYmtMMU96bzV4aTdmSFhkR3BNWGJUUjM4dWN6eEFvR0FkbUQyVFltSE1BUU5kNlJGYTd0dAp3NHBSRW5seWxTY3VPYXVkQ2h6enF0OERyb29EaWVENFhvTGQvdE1UcUlkYTNRQ05LM0ZVaWhXdUdyUEFKa01yCjZYcU96a3NCcnVRcDR4SmdjbWZSb016MGZTRG5sMHNSOFd3K3lMVWlrMVdzSlhMVUJVWnBpaDNlWmh4REw0NGwKaHI0L0lpTXdqVkZZcDE3VnVEZnk2Umc9Ci0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0K"
	ValidMetadata                      = baseentitee.MapToEntityMetadataFromCommonTypes(
		uuid.UUID{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC),
		time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC),
		false,
		nil,
	)
	ValidPing = entity.MapToEntity(
		"ping",
		ValidMetadata,
	)
	ValidPong = entity.MapToEntity(
		"pong",
		ValidMetadata,
	)
	ValidReturningPing = entity.MapToEntity(
		"Ping!",
		ValidMetadata,
	)
	ValidReturningPong = entity.MapToEntity(
		"Pong!",
		ValidMetadata,
	)
	InvalidPingPong = entity.MapToEntity(
		"ring",
		ValidMetadata,
	)
	RestApiV1PingPongRaw = []rest_api_V1.PingPongRaw{
		{
			Message:   new("ping"),
			CreatedAt: new(time.Now()),
			Deleted:   new(false),
			UpdatedAt: new(time.Now()),
			DeletedAt: nil,
			Id:        new(uuid.New().String()),
		},
		{
			Message:   new("pong"),
			CreatedAt: new(time.Now()),
			Deleted:   new(false),
			UpdatedAt: new(time.Now()),
			DeletedAt: nil,
			Id:        new(uuid.New().String()),
		},
	}
	RestApiV1PingPongsRaw = rest_api_V1.PingPongsRaw{
		Pingpongs: &RestApiV1PingPongRaw,
	}
)
