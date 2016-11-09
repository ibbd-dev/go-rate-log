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


	logger := New(file, "", time.RFC3339)
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


