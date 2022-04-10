package speed

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkOOKLA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := RunSpeedTest(context.Background(), OOKLA)
		require.NoError(b, err)
	}
}

func BenchmarkFastcom(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := RunSpeedTest(context.Background(), FastCom)
		require.NoError(b, err)
	}
}
