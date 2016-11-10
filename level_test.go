package rateLog

import (
	"testing"
	"time"

	"github.com/ibbd-dev/go-rotate-file"
)

func TestLevelLog(t *testing.T) {
	file := rotateFile.Open("/tmp/test-level-log.log")
	defer file.Close()

	logger := New(file, "", time.RFC3339)
	logger.SetDuration(time.Millisecond * 100)

	ll := NewLevelLog(logger, LevelWarn)
	ll.Debug("Debug")
	ll.Info("Info")
	ll.Warn("Warn")
	ll.Error("Error")
	time.Sleep(time.Second)
	ll.Fatal("Fatal")
}
