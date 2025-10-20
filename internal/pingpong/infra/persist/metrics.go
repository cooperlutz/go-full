package persist

import (
	"context"

	"github.com/cooperlutz/go-full/pkg/telemetree"
	"github.com/cooperlutz/go-full/pkg/types"
)

func (r *pingPongPersistPostgresRepository) TotalNumberOfPingPongs(ctx context.Context) (int64, error) {
	// telemetree: Add a tracing span for the SavePingPong operation
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.totalnumberofpingpongs")
	defer span.End()

	count, err := r.query.TotalNumberOfPingPongs(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *pingPongPersistPostgresRepository) TotalNumberOfPings(ctx context.Context) (int64, error) {
	// telemetree: Add a tracing span for the SavePingPong operation
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.totalnumberofpings")
	defer span.End()

	count, err := r.query.TotalNumberOfPings(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *pingPongPersistPostgresRepository) TotalNumberOfPongs(ctx context.Context) (int64, error) {
	// telemetree: Add a tracing span for the SavePingPong operation
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.totalnumberofpongs")
	defer span.End()

	count, err := r.query.TotalNumberOfPongs(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *pingPongPersistPostgresRepository) AverageNumberOfPingPongsCreatedPerDay(ctx context.Context) (int64, error) {
	// telemetree: Add a tracing span for the SavePingPong operation
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.averagenenumberofpingpongscreatedperday")
	defer span.End()

	var numDays int

	frequencyDist, err := r.query.FrequencyDistributionByDay(ctx)
	if err != nil {
		return 0, err
	}

	totalPingPings, err := r.query.TotalNumberOfPingPongs(ctx)
	if err != nil {
		return 0, err
	}

	// Calculate the number of days based on the length of the frequency distribution
	numDays = len(frequencyDist)

	return totalPingPings / int64(numDays), nil
}

func (r *pingPongPersistPostgresRepository) AverageNumberOfPingsCreatedPerDay(ctx context.Context) (int64, error) {
	// telemetree: Add a tracing span for the SavePingPong operation
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.averagenumberofpingscreatedperday")
	defer span.End()

	var numDays int

	frequencyDist, err := r.query.FrequencyDistributionByDayPing(ctx)
	if err != nil {
		return 0, err
	}

	totalPings, err := r.query.TotalNumberOfPings(ctx)
	if err != nil {
		return 0, err
	}

	numDays = len(frequencyDist)

	return totalPings / int64(numDays), nil
}

func (r *pingPongPersistPostgresRepository) AverageNumberOfPongsCreatedPerDay(ctx context.Context) (int64, error) {
	// telemetree: Add a tracing span for the SavePingPong operation
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.averagenumberofpongscreatedperday")
	defer span.End()

	var numDays int

	frequencyDist, err := r.query.FrequencyDistributionByDayPong(ctx)
	if err != nil {
		return 0, err
	}

	totalPings, err := r.query.TotalNumberOfPongs(ctx)
	if err != nil {
		return 0, err
	}

	numDays = len(frequencyDist)

	return totalPings / int64(numDays), nil
}

func (r *pingPongPersistPostgresRepository) TotalNumberOfPingPongsCreatedPerDay(ctx context.Context) ([]types.MeasureCountbyDateTime, error) {
	// telemetree: Add a tracing span for the SavePingPong operation
	ctx, span := telemetree.AddSpan(ctx, "persist.postgres.totalnumberofpingpongscreatedperday")
	defer span.End()

	countPerDay, err := r.query.CountPerDay(ctx)
	if err != nil {
		return nil, err
	}

	var pingPongsPerDay []types.MeasureCountbyDateTime
	for _, cpd := range countPerDay {
		val := types.MeasureCountbyDateTime{
			DateTime: cpd.CreationDate.Time,
			Count:    int(cpd.CountCreated),
		}
		pingPongsPerDay = append(pingPongsPerDay, val)
	}

	return pingPongsPerDay, nil
}
