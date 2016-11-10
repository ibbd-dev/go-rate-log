# 控制写入频率的log

按时间周期写入，保证一个周期内，只会写入一次。

对于很多写log的情况，我们都需要控制一定的输出频率，避免log文件被写爆掉。

## Install

```sh
go get -u github.com/ibbd-dev/go-rate-log
```

## Example

```go
package main

import (
	"os"
	"time"

    "github.com/ibbd-dev/go-rate-log"
)

func main() {
	// 文件Flag
	fileFlag := os.O_WRONLY | os.O_CREATE | os.O_APPEND
	file, err := os.OpenFile("/tmp/test-rate.log", fileFlag, 0666)
	defer file.Close()
	if err != nil {
		panic(err)
	}


	logger := rateLog.New(file, "", time.RFC3339)
	logger.SetDuration(time.Millisecond * 100) // 每100毫秒写入一次
	logger.Println("hello world")
	logger.Println("hello world2")
	time.Sleep(time.Millisecond * 105)
	logger.Println("hello world3")
	logger.Println("hello world3")
	time.Sleep(time.Millisecond * 10)
	logger.Println("hello world4")
}
```

## 根据周期自动切割文件

配合`github.com/ibbd-dev/go-rotate-file`，即可实现自动切割文件。

```go
package main

import (
	"os"
	"time"

	"github.com/ibbd-dev/go-rotate-file"
    "github.com/ibbd-dev/go-rate-log"
)

func main() {
	// 最终保存时，文件名如：/tmp/test-rotate.log.161109
    // 其中后缀161109表示2016-11-09，默认是按照日期对文件进行切割
	file := rotateFile.Open("/tmp/test-rotate.log")
	defer file.Close()

	logger := rateLog.New(file, "", time.RFC3339)
	logger.SetDuration(time.Millisecond * 100) // 每100毫秒写入一次
	logger.Println("hello world")
	logger.Println("hello world2")
	time.Sleep(time.Millisecond * 105)
	logger.Println("hello world3")
	logger.Println("hello world3")
	time.Sleep(time.Millisecond * 10)
	logger.Println("hello world4")
}
```

## 根据错误类型写日志

```go
package main

import (
	"time"

	"github.com/ibbd-dev/go-rotate-file"
    "github.com/ibbd-dev/go-rate-log"
)

func main() {
	file := rotateFile.Open("/tmp/test-level-log.log")
	defer file.Close()

	logger := rateLog.New(file, "", time.RFC3339)
	logger.SetDuration(time.Millisecond * 100)

	ll := rateLog.NewLevelLog(logger, rateLog.LevelWarn)
	ll.Debug("Debug")
	ll.Info("Info")
	ll.Warn("Warn")
	ll.Error("Error")
	time.Sleep(time.Second)
	ll.Fatal("Fatal")
}
```

