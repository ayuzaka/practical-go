package chapter13

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type contextTimeKey string

const timeKey contextTimeKey = "timeKey"

func CurrentTime(ctx context.Context) time.Time {
	v := ctx.Value(timeKey)
	if t, ok := v.(time.Time); ok {
		return t
	}

	return time.Now()
}

func SetFixTime(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, timeKey, t)
}

func NextMonth(ctx context.Context) time.Month {
	now := CurrentTime(ctx)

	return now.AddDate(0, 1, 0).Month()
}

func TestNextMonth(t *testing.T) {
	ctx := SetFixTime(context.Background(), time.Date(1980, time.December, 1, 0, 0, 0, 0, time.Local))

	assert.Equal(t, time.January, NextMonth(ctx))
}
