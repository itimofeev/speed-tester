package fastcom

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFastCom(t *testing.T) {
	tests := []struct {
		name                 string
		getContext           func() (context.Context, func())
		expectedTestDuration time.Duration
		expectedAccuracy     time.Duration
	}{
		{
			name: "run_with_big_timeout",
			getContext: func() (context.Context, func()) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
				return ctx, cancel
			},
			expectedTestDuration: time.Second * 10,
			expectedAccuracy:     time.Second * 2,
		},
		{
			name: "run_with_small_timeout",
			getContext: func() (context.Context, func()) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
				return ctx, cancel
			},
			expectedTestDuration: time.Second * 2,
			expectedAccuracy:     time.Second * 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := NewProvider()

			ctx, cancel := test.getContext()
			defer cancel()

			start := time.Now()

			result, err := p.RunSpeedTest(ctx)
			require.NoError(t, err)
			require.NotZero(t, result.DownloadSpeedBytesPerSecond)
			require.Equal(t, float64(-1), result.UploadSpeedBytesPerSecond)
			require.WithinDuration(t, start.Add(test.expectedTestDuration), time.Now(), test.expectedAccuracy)
		})
	}
}
