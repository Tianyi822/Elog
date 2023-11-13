package format

import "time"

type Format interface {
	genLog(t time.Time, level, trace, msg string) string
	DebugFormat(t time.Time, trace, msg string) string
	InfoFormat(t time.Time, trace, msg string) string
	WarnFormat(t time.Time, trace, msg string) string
	ErrorFormat(t time.Time, trace, msg string) string
	FatalFormat(t time.Time, trace, msg string) string
	PanicFormat(t time.Time, trace, msg string) string
}
