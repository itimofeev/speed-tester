package fastcom

import (
	"context"
	"fmt"

	"github.com/itimofeev/speed-tester/internal/providers"
	"gopkg.in/ddo/go-fast.v0"
)

// Provider that runs speed test using fast.com
type Provider struct {
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) RunSpeedTest(ctx context.Context) (*providers.SpeedInfo, error) {
	fastCom := fast.New()

	if err := fastCom.Init(); err != nil {
		return nil, fmt.Errorf("error initing fastcom: %w", err)
	}

	// get urls
	urls, err := fastCom.GetUrls()
	if err != nil {
		return nil, fmt.Errorf("error getting urls for fastcom: %w", err)
	}

	kbpsChan := make(chan float64, 10)

	var measureError error
	go func() {
		if err := fastCom.Measure(urls, kbpsChan); err != nil {
			measureError = err
			return
		}
	}()

	countOfMeasures := 0
	sumBytesPerSecond := float64(0)
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case kbps, ok := <-kbpsChan:
			sumBytesPerSecond += kbps * 1000
			countOfMeasures++
			if !ok {
				break loop
			}
		}
	}
	if measureError != nil {
		return nil, fmt.Errorf("measure error for fastcom: %w", err)
	}
	return &providers.SpeedInfo{
		DownloadSpeedBytesPerSecond: sumBytesPerSecond / float64(countOfMeasures),
		UploadSpeedBytesPerSecond:   -1,
	}, nil
}
