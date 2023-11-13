package elog

import (
	"fmt"
	"gitee.com/xxc_opensource/elog/core"
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
	Level LogLevel
	core  *core.Core
}

// NewElog 创建一个日志组件
func NewElog(logLevel LogLevel, core *core.Core) *Elog {
	elog := &Elog{
		Level: logLevel,
		core:  core,
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
	el.core.Write([]byte(fmt.Sprintf(format, a...)))
}

func (el *Elog) Info(format string, a ...any) {
	if el.Level > Info {
		return
	}
	el.core.Write([]byte(fmt.Sprintf(format, a...)))
}

func (el *Elog) Warn(format string, a ...any) {
	if el.Level > Warn {
		return
	}
	el.core.Write([]byte(fmt.Sprintf(format, a...)))
}

func (el *Elog) Error(format string, a ...any) {
	if el.Level > Error {
		return
	}
	el.core.Write([]byte(fmt.Sprintf(format, a...)))
}

func (el *Elog) Fatal(format string, a ...any) {
	if el.Level > Fatal {
		return
	}
	el.core.Write([]byte(fmt.Sprintf(format, a...)))
}

func (el *Elog) Panic(format string, a ...any) {
	if el.Level > Panic {
		return
	}
	el.core.Write([]byte(fmt.Sprintf(format, a...)))
}