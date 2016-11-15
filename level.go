package rateLog

import (
	"fmt"
)

// 日志中的前缀标题
const (
	titleDebug = "[DEBUG]"
	titleInfo  = "[INFO]"
	titleWarn  = "[WARN]"
	titleError = "[ERROR]"
	titleFatal = "[FATAL]"
)

// 日志优先级
type Priority int

type LevelLog struct {
	log   ILogger
	level Priority
}

const (
	LevelAll Priority = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelOff
)

// NewLevelLog 写入等级日志
// 级别高于logLevel才会被写入
func NewLevelLog(log ILogger, logLevel Priority) *LevelLog {
	return &LevelLog{
		log:   log,
		level: logLevel,
	}
}

func (l *LevelLog) SetLevel(logLevel Priority) {
	l.level = logLevel
}

func (l *LevelLog) Debug(v ...interface{}) error {
	return l.output(LevelDebug, titleDebug, v...)
}

func (l *LevelLog) Info(v ...interface{}) error {
	return l.output(LevelInfo, titleInfo, v...)
}

func (l *LevelLog) Warn(v ...interface{}) error {
	return l.output(LevelWarn, titleWarn, v...)
}

func (l *LevelLog) Error(v ...interface{}) error {
	return l.output(LevelError, titleError, v...)
}

func (l *LevelLog) Fatal(v ...interface{}) error {
	return l.output(LevelFatal, titleFatal, v...)
}

func (l *LevelLog) Output(s string) error {
	return l.log.Output(s)
}

//************************** Private ********************

func (l *LevelLog) output(level Priority, title string, v ...interface{}) error {
	if level >= l.level {
		return l.log.Output(title + " " + fmt.Sprint(v...))
	}
	return nil
}
