package speed

import (
	"context"

	"github.com/itimofeev/speed-tester/internal/providers"
)

type ProviderName string

const (
	OOKLA   ProviderName = "OOKLA"
	FastCom ProviderName = "FastCom"
)

type Provider interface {
	RunSpeedTest(ctx context.Context) (*providers.SpeedInfo, error)
}

type Speed struct {
	DownloadBitsPerSecond string
	UploadBitsPerSecond   string
}
