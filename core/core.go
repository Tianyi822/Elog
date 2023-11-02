package core

import (
	fileWriter "easy-go-log/fileWriter"
	"easy-go-log/interface/logger"
)

type LogLevel byte

const (
	// LevelDebug 日志级别
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// OutPutType 输出方式
type OutPutType byte

const (
	// OutPutToFile 输出到文件
	OutPutToFile OutPutType = iota
	// OutPutToConsole 输出到控制台
	OutPutToConsole
	// TODO: 输出到数据库，Kafka，MQ 等，后续再实现
)

// Core 日志核心组件，用于对接各种输出路径，包含但不限于日志文本文件，MQ 消息队列，Kafka 消息队列，HDFS 分布式集群等
// 这个组件中还要处理一些并发操作，防止出现日志并发写入的问题
type Core struct {
	LogContext chan string      // 日志内容
	Writer     logger.LogWriter // 日志写入对象（实现了 LogWriter 接口的实例）
}

// NewCoreWithFileLogConf 创建一个日志核心对象，输出到文件
func NewCoreWithFileLogConf(config *fileWriter.FWConfig) *Core {
	fop := fileWriter.CreateFileWriter(config)

	return &Core{
		LogContext: make(chan string, 1000),
		Writer:     fop,
	}
}

// NewCore 创建一个日志核心对象
func NewCore() *Core {
	return &Core{
		LogContext: make(chan string, 1000),
		Writer:     nil,
	}
}

// Output 将日志内容写入到文件中
func (c *Core) Output() {
	for {
		select {
		case log := <-c.LogContext:
			err := c.Writer.WriteLog([]byte(log))
			if err != nil {
				panic(err)
			}
		}
	}
}
