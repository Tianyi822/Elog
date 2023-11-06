package elog

import (
	"fmt"
	"time"
)

// LogFormat 日志格式
type LogFormat struct {
	DateFormat string
	MsgFormat  string
}

// DefaultLogFormat 默认日志格式
func DefaultLogFormat() *LogFormat {
	return &LogFormat{
		DateFormat: "2006-01-02 15:04:05",
		MsgFormat:  "[%s] [%s] %s\n",
	}
}

// CustomLogFormat 自定义日志格式
func CustomLogFormat(dateFormat string, msgFormat string) *LogFormat {
	return &LogFormat{
		DateFormat: dateFormat,
		MsgFormat:  msgFormat,
	}
}

func (lf *LogFormat) Build(level LogLevel, logContent string) string {
	return fmt.Sprintf(lf.MsgFormat, time.Now().Format(lf.DateFormat), level, logContent)
}
