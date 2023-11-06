package elog

import (
	"fmt"
	"time"
)

// LogFormat 日志格式
// TODO: 日志格式需要自定义几个关键字: time, level, track, msg
// TODO: 通过解析关键字可以自定义日志格式，例如: [time] [level] [track] [msg]
// TODO: 同时 msg 可以接受 map 类型，输出格式为: time level track json.Marshal(msg)，或者自定义格式
type LogFormat struct {
	DateFormat    string
	ContentFormat string
}

// DefaultLogFormat 默认日志格式
func DefaultLogFormat() *LogFormat {
	return &LogFormat{
		DateFormat:    "2006-01-02 15:04:05",
		ContentFormat: "[%s] [%s] %s\n",
	}
}

// CustomLogFormat 自定义日志格式
func CustomLogFormat(dateFormat string, msgFormat string) *LogFormat {
	return &LogFormat{
		DateFormat:    dateFormat,
		ContentFormat: msgFormat,
	}
}

func (lf *LogFormat) Build(level LogLevel, logContent string) string {
	return fmt.Sprintf(lf.ContentFormat, time.Now().Format(lf.DateFormat), level, logContent)
}
