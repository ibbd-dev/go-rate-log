package rateLog

import (
	"os"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	// 文件Flag
	fileFlag := os.O_WRONLY | os.O_CREATE | os.O_APPEND
	file, err := os.OpenFile("/tmp/test-rate.log", fileFlag, 0666)
	defer file.Close()
	if err != nil {
		t.Fatal(err)
	}

	logger := New(file, "", time.RFC3339)
	logger.Println("hello world")
	logger.Println("hello world2")
}
