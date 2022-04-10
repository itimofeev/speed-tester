package main

import (
	"context"
	"fmt"

	speedTest "github.com/itimofeev/speed-tester/pkg/speed-test"
)

func main() {
	test, err := speedTest.RunSpeedTest(context.Background(), speedTest.OOKLA)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Speed according to OOKLA's Speed Test is %s (download) and %s (upload)\n", test.DownloadBitsPerSecond, test.UploadBitsPerSecond)

	test, err = speedTest.RunSpeedTest(context.Background(), speedTest.FastCom)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Speed according to fast.com Speed Test is %s (download) and %s (upload)\n", test.DownloadBitsPerSecond, test.UploadBitsPerSecond)
}
