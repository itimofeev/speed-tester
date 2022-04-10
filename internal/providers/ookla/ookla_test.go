package ookla

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestOoklaProvider(t *testing.T) {
	tests := []struct {
		name       string
		getContext func() (context.Context, func())
		wantError  error
	}{
		{
			name: "context_timeout_cased_error",
			getContext: func() (context.Context, func()) {
				return context.WithTimeout(context.Background(), time.Second)
			},
			wantError: context.DeadlineExceeded,
		},
		{
			name: "successful_without_timeout",
			getContext: func() (context.Context, func()) {
				return context.Background(), func() {}
			},
			wantError: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := NewProvider(true)
			ctx, cancel := test.getContext()
			defer cancel()

			result, err := p.RunSpeedTest(ctx)
			if test.wantError != nil {
				require.Error(t, err)
				require.EqualError(t, err, test.wantError.Error())
				return
			}
			require.NoError(t, err)

			require.Positive(t, result.DownloadSpeedBytesPerSecond)
			require.Positive(t, result.UploadSpeedBytesPerSecond)
		})
	}
}
