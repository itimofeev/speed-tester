package speed

import (
	"context"
	"errors"
	"testing"

	"github.com/itimofeev/speed-tester/internal/providers"
	"github.com/stretchr/testify/require"
)

func TestTester(t *testing.T) {
	const providerName = "someProviderName"
	tests := []struct {
		name        string
		providerMap map[ProviderName]Provider
		wantError   error
		wantSpeed   *Speed
	}{
		{
			name:        "error_on_unknown_provider",
			providerMap: map[ProviderName]Provider{},
			wantError:   errors.New("unknown provider someProviderName"),
		},
		{
			name:        "error_on_error_provider",
			providerMap: map[ProviderName]Provider{providerName: errorProvider{}},
			wantError:   errors.New("error running speed test: some error"),
		},
		{
			name:        "not_measured",
			providerMap: map[ProviderName]Provider{providerName: notMeasuredProvider{}},
			wantSpeed: &Speed{
				DownloadBitsPerSecond: "not measured",
				UploadBitsPerSecond:   "not measured",
			},
		},
		{
			name:        "ok",
			providerMap: map[ProviderName]Provider{providerName: okProvider{}},
			wantSpeed: &Speed{
				DownloadBitsPerSecond: "1.0 MB/s",
				UploadBitsPerSecond:   "1.0 kB/s",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			st := speedTester{
				providerMap: test.providerMap,
			}

			result, err := st.RunSpeedTest(context.Background(), providerName)
			if test.wantError != nil {
				require.Error(t, err)
				require.EqualError(t, err, test.wantError.Error())
				require.Nil(t, result)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, result)
			require.Equal(t, *test.wantSpeed, *result)
		})
	}
}

type notMeasuredProvider struct {
}

func (n notMeasuredProvider) RunSpeedTest(_ context.Context) (*providers.SpeedInfo, error) {
	return &providers.SpeedInfo{
		DownloadSpeedBytesPerSecond: -1,
		UploadSpeedBytesPerSecond:   -1,
	}, nil
}

type errorProvider struct {
}

func (n errorProvider) RunSpeedTest(_ context.Context) (*providers.SpeedInfo, error) {
	return nil, errors.New("some error")
}

type okProvider struct {
}

func (n okProvider) RunSpeedTest(_ context.Context) (*providers.SpeedInfo, error) {
	return &providers.SpeedInfo{
		DownloadSpeedBytesPerSecond: 1024 * 1024,
		UploadSpeedBytesPerSecond:   1024,
	}, nil
}
