package elog

import (
	"fmt"
	"gitee.com/xxc_opensource/elog/core"
	"gitee.com/xxc_opensource/elog/format"
	"time"
)

type LogLevel byte

const (
	Debug LogLevel = 1 << iota
	Info
	Warn
	Error
	Fatal
	Panic
)

type Elog struct {
	Level  LogLevel
	format format.Format
	core   *core.Core
}

// NewElog 创建一个日志组件
func NewElog(logLevel LogLevel, format format.Format, core *core.Core) *Elog {
	elog := &Elog{
		Level:  logLevel,
		core:   core,
		format: format,
	}
	elog.core.Start()

	return elog
}

// Close 关闭日志组件
func (el *Elog) Close() {
	el.core.Close()
}

func (el *Elog) Debug(format string, a ...any) {
	if el.Level > Debug {
		return
	}
	log := el.format.DebugFormat(time.Now(), "", fmt.Sprintf(format, a...))
	el.core.Write(log)
}

func (el *Elog) Info(format string, a ...any) {
	if el.Level > Info {
		return
	}
	log := el.format.InfoFormat(time.Now(), "", fmt.Sprintf(format, a...))
	el.core.Write(log)
}

func (el *Elog) Warn(format string, a ...any) {
	if el.Level > Warn {
		return
	}
	log := el.format.WarnFormat(time.Now(), "", fmt.Sprintf(format, a...))
	el.core.Write(log)
}

func (el *Elog) Error(format string, a ...any) {
	if el.Level > Error {
		return
	}
	log := el.format.ErrorFormat(time.Now(), "", fmt.Sprintf(format, a...))
	el.core.Write(log)
}

func (el *Elog) Fatal(format string, a ...any) {
	if el.Level > Fatal {
		return
	}
	log := el.format.FatalFormat(time.Now(), "", fmt.Sprintf(format, a...))
	el.core.Write(log)
}

func (el *Elog) Panic(format string, a ...any) {
	if el.Level > Panic {
		return
	}
	log := el.format.PanicFormat(time.Now(), "", fmt.Sprintf(format, a...))
	el.core.Write(log)
}
