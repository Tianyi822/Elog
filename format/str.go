package format

import (
	"strings"
	"time"
)

// StrFormat 日志格式
// TODO: 日志格式需要自定义几个关键字: time, level, track, msg
// TODO: 通过解析关键字可以自定义日志格式，例如: [time] [level] [trace] [msg]
// TODO: 同时 msg 可以接受 map 类型，输出格式为: time level track json.Marshal(msg)，或者自定义格式
type StrFormat struct {
	DateFormat    string
	ContentFormat string
}

// NewStrFormat 默认日志格式
func NewStrFormat() *StrFormat {
	return &StrFormat{
		DateFormat:    "2006-01-02 15:04:05",
		ContentFormat: "time level trace msg",
	}
}

// CustomStrFormat 自定义日志格式
func CustomStrFormat(dateFormat string, msgFormat string) *StrFormat {
	return &StrFormat{
		DateFormat:    dateFormat,
		ContentFormat: msgFormat,
	}
}

// GenLog 根据格式生成日志
func (sf *StrFormat) genLog(t time.Time, level, trace, msg string) string {
	log := strings.ReplaceAll(sf.ContentFormat, "time", t.Format(sf.DateFormat))
	log = strings.ReplaceAll(log, "level", level)
	log = strings.ReplaceAll(log, "trace", trace)
	log = strings.ReplaceAll(log, "msg", msg)
	return log
}

func (sf *StrFormat) DebugFormat(t time.Time, trace, msg string) string {
	return sf.genLog(t, "DEBUG", trace, msg)
}

func (sf *StrFormat) InfoFormat(t time.Time, trace, msg string) string {
	return sf.genLog(t, "INFO", trace, msg)
}

func (sf *StrFormat) WarnFormat(t time.Time, trace, msg string) string {
	return sf.genLog(t, "WARN", trace, msg)
}

func (sf *StrFormat) ErrorFormat(t time.Time, trace, msg string) string {
	return sf.genLog(t, "ERROR", trace, msg)
}

func (sf *StrFormat) FatalFormat(t time.Time, trace, msg string) string {
	return sf.genLog(t, "FATAL", trace, msg)
}

func (sf *StrFormat) PanicFormat(t time.Time, trace, msg string) string {
	return sf.genLog(t, "PANIC", trace, msg)
}
