package xzap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"
)

func TestTimeEncoder(t *testing.T) {
	err := InitZLog([]string{"stderr", "/tmp/kang.log"}, zapcore.DebugLevel)
	if err != nil {
		panic(err)
	}

	err = InitZLog([]string{"stderr", "/tmp/kang.log"}, zapcore.DebugLevel)
	if err != nil {
		panic(err)
	}

	for {
		Info("hello", zap.String("name", "kangkang"))
		Info("hello", zap.String("name", "kangkang"))
		time.Sleep(1 * time.Second)
	}
}
