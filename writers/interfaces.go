package writers

type LogWriter interface {
	GetHash() string
	Close() error
	WriteLog([]byte) error
}
