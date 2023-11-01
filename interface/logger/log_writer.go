package logger

type LogWriter interface {
	Close() error
	WriteLog([]byte) error
}
