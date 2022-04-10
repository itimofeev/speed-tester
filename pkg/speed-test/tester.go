package speed

import (
	"context"
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/itimofeev/speed-tester/internal/providers/fastcom"
	"github.com/itimofeev/speed-tester/internal/providers/ookla"
)

// speedTester internal struct with providerMap field needed to write unit tests on RunSpeedTest method
type speedTester struct {
	providerMap map[ProviderName]Provider
}

func RunSpeedTest(ctx context.Context, providerName ProviderName) (*Speed, error) {
	return defaultSpeedTester.RunSpeedTest(ctx, providerName)
}

var defaultSpeedTester = speedTester{
	providerMap: map[ProviderName]Provider{
		OOKLA:   ookla.NewProvider(true),
		FastCom: fastcom.NewProvider(),
	},
}

func (st *speedTester) RunSpeedTest(ctx context.Context, providerName ProviderName) (*Speed, error) {
	provider, ok := st.providerMap[providerName]
	if !ok {
		return nil, fmt.Errorf("unknown provider %s", providerName)
	}

	speedInfo, err := provider.RunSpeedTest(ctx)
	if err != nil {
		return nil, fmt.Errorf("error running speed test: %w", err)
	}
	speed := &Speed{
		DownloadBitsPerSecond: "not measured",
		UploadBitsPerSecond:   "not measured",
	}
	if speedInfo.DownloadSpeedBytesPerSecond > 0 {
		speed.DownloadBitsPerSecond = humanize.Bytes(uint64(speedInfo.DownloadSpeedBytesPerSecond)) + "/s"
	}
	if speedInfo.UploadSpeedBytesPerSecond > 0 {
		speed.UploadBitsPerSecond = humanize.Bytes(uint64(speedInfo.UploadSpeedBytesPerSecond)) + "/s"
	}

	return speed, nil
}
