package persist

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
	"github.com/cooperlutz/go-full/pkg/types"
)

// TotalNumberOfPingPongs - Returns the total number of pingpongs in the database.
func (r *PingPongPersistPostgresRepository) TotalNumberOfPingPongs(ctx context.Context) (types.QuantityMetric, error) {
	// telemetree: Add a tracing span for the SavePingPong operation
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.totalnumberofpingpongs")
	defer span.End()

	count, err := r.query.TotalNumberOfPingPongs(ctx)
	if err != nil {
		return types.QuantityMetric{Quantity: 0}, err
	}

	return types.QuantityMetric{Quantity: count}, nil
}

// TotalNumberOfPings - Returns the total number of pings in the database.
func (r *PingPongPersistPostgresRepository) TotalNumberOfPings(ctx context.Context) (types.QuantityMetric, error) {
	// telemetree: Add a tracing span for the SavePingPong operation
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.totalnumberofpings")
	defer span.End()

	count, err := r.query.TotalNumberOfPings(ctx)
	if err != nil {
		return types.QuantityMetric{Quantity: 0}, err
	}

	return types.QuantityMetric{Quantity: count}, nil
}

// TotalNumberOfPongs - Returns the total number of pongs in the database.
func (r *PingPongPersistPostgresRepository) TotalNumberOfPongs(ctx context.Context) (types.QuantityMetric, error) {
	// telemetree: Add a tracing span for the SavePingPong operation
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.totalnumberofpongs")
	defer span.End()

	count, err := r.query.TotalNumberOfPongs(ctx)
	if err != nil {
		return types.QuantityMetric{Quantity: 0}, err
	}

	return types.QuantityMetric{Quantity: count}, nil
}

// AverageNumberOfPingPongsCreatedPerDay - Returns the average number of pingpongs created per day.
func (r *PingPongPersistPostgresRepository) AverageNumberOfPingPongsCreatedPerDay(ctx context.Context) (types.QuantityMetric, error) {
	// telemetree: Add a tracing span for the SavePingPong operation
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.averagenenumberofpingpongscreatedperday")
	defer span.End()

	var numDays int

	frequencyDist, err := r.query.FrequencyDistributionByDay(ctx)
	if err != nil {
		return types.QuantityMetric{Quantity: 0}, err
	}

	totalPingPings, err := r.query.TotalNumberOfPingPongs(ctx)
	if err != nil {
		return types.QuantityMetric{Quantity: 0}, err
	}

	// Calculate the number of days based on the length of the frequency distribution
	numDays = len(frequencyDist)

	avgNumPerDay := totalPingPings / int64(numDays)

	return types.QuantityMetric{Quantity: avgNumPerDay}, nil
}

// TotalNumberOfPingPongsCreatedPerDay - Returns the total number of pingpongs created per day as a slice of MeasureCountbyDateTime.
func (r *PingPongPersistPostgresRepository) TotalNumberOfPingPongsCreatedPerDay(ctx context.Context) ([]types.MeasureCountbyDateTimeMetric, error) {
	// telemetree: Add a tracing span for the SavePingPong operation
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.totalnumberofpingpongscreatedperday")
	defer span.End()

	countPerDay, err := r.query.CountPerDay(ctx)
	if err != nil {
		return nil, err
	}

	var pingPongsPerDay []types.MeasureCountbyDateTimeMetric
	for _, cpd := range countPerDay {
		val := types.MeasureCountbyDateTimeMetric{
			DateTime: cpd.CreationDate.Time,
			Count:    int(cpd.CountCreated),
		}
		pingPongsPerDay = append(pingPongsPerDay, val)
	}

	return pingPongsPerDay, nil
}
