package ookla

import (
	"context"
	"fmt"

	"github.com/itimofeev/speed-tester/internal/providers"
	"github.com/showwin/speedtest-go/speedtest"
)

type Provider struct {
	onlyClosest bool
}

func NewProvider(onlyClosest bool) *Provider {
	return &Provider{
		onlyClosest: onlyClosest,
	}
}

const megabytesToBits = 1000 * 1000

func (p *Provider) RunSpeedTest(ctx context.Context) (*providers.SpeedInfo, error) {
	user, err := speedtest.FetchUserInfoContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetchgin user info: %w", err)
	}

	serverList, err := speedtest.FetchServerListContext(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error getching servers: %w", err)
	}

	if p.onlyClosest {
		// underlying library returns servers sorted by distance starting from closest one
		serverList = serverList[:1]
	}

	var lastError error
	var successfulServer *speedtest.Server
	for _, s := range serverList {
		lastError = nil
		if err := s.DownloadTestContext(ctx, false); err != nil {
			lastError = err
			continue
		}
		if err := s.UploadTestContext(ctx, false); err != nil {
			lastError = err
			continue
		}
		successfulServer = s
		break
	}
	if lastError != nil {
		return nil, lastError
	}

	return &providers.SpeedInfo{
		DownloadSpeedBytesPerSecond: successfulServer.DLSpeed * megabytesToBits,
		UploadSpeedBytesPerSecond:   successfulServer.ULSpeed * megabytesToBits,
	}, nil
}
