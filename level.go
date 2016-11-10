package rateLog

import (
	"fmt"
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
	return l.Output(LevelDebug, "[DEBEG] "+fmt.Sprintln(v...))
}

func (l *LevelLog) Info(v ...interface{}) error {
	return l.Output(LevelInfo, "[DEBEG] "+fmt.Sprintln(v...))
}

func (l *LevelLog) Warn(v ...interface{}) error {
	return l.Output(LevelWarn, "[DEBEG] "+fmt.Sprintln(v...))
}

func (l *LevelLog) Error(v ...interface{}) error {
	return l.Output(LevelError, "[DEBEG] "+fmt.Sprintln(v...))
}

func (l *LevelLog) Fatal(v ...interface{}) error {
	return l.Output(LevelFatal, "[DEBEG] "+fmt.Sprintln(v...))
}

func (l *LevelLog) Output(level Priority, s string) error {
	if level >= l.level {
		return l.log.Output(s)
	}
	return nil
}
