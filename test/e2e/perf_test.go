package e2e_test

import (
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func TestLoad(t *testing.T) {
	rate := vegeta.Rate{Freq: 100, Per: time.Second}
	duration := 4 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    serverAddr + "/examination/api/v1/exams",
	})
	attacker := vegeta.NewAttacker()
	var runs int

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		runs += 1
		slog.Info("num of runs"+": ", slog.Int("runs", runs))
		metrics.Add(res)
	}
	metrics.Close()

	assert.LessOrEqual(t, metrics.Latencies.P99, 50*time.Millisecond)
}
