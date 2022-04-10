# speed-tester
Test assignment golang checking connection speed. Full description can be found in [assignment.md](assignment.md)

This library has single exposed method RunSpeedTest of package github.com/itimofeev/speed-tester/pkg/speed/.
Caller should pass context and one of test speed Provider of "OOKLA" or "FastCom" for using speedtest.net and fast.com respectively.

## How to use library
`go get github.com/itimofeev/speed-tester`

Example:

```go
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
}
```

## About tests
Test coverage can be seen by running unit tests with coverage:
`make cover`

In author's opinion it's no so good idea to write unit tests that depends on network, because it can cause of annoying flaky tests.

There's two benchmark test in pkg/speed-test/tester_benchmark_test.go, which were required by assignment. 
In author's opinion benchmark tests in go more effective in clear unit tests which don't depends on third party systems like fastcom, database and so on.
This instrument created for testing performance of code. 

## How to improve quality
- set up build via circle ci for example
- add linters via golangci-lint