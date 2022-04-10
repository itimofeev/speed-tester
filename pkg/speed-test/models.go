package speed

import (
	"context"

	"github.com/itimofeev/speed-tester/internal/providers"
)

// ProviderName type for provider names
type ProviderName string

const (
	OOKLA   ProviderName = "OOKLA"
	FastCom ProviderName = "FastCom"
)

// Provider is an interface which implemented by structs for specific speed test site like fast.com of speedtest.net
type Provider interface {
	RunSpeedTest(ctx context.Context) (*providers.SpeedInfo, error)
}

// Speed contains speed in human-readable format like 79 MB/s
type Speed struct {
	DownloadBitsPerSecond string
	UploadBitsPerSecond   string
}
