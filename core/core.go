package core

import (
	fileWriter "easy-go-log/fileWriter"
	"easy-go-log/interface/logger"
	"path/filepath"
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

// 用于标识不同的输出模式
const (
	OutPutToConsole OutPutType = 1
	OutPutToFile    OutPutType = 1 << OutPutToConsole
	OutPutToKafka   OutPutType = 2 << OutPutToConsole
)

// Core 日志核心组件，用于对接各种输出路径，包含但不限于日志文本文件，MQ 消息队列，Kafka 消息队列，HDFS 分布式集群等
// 这个组件中还要处理一些并发操作，防止出现日志并发写入的问题
type Core struct {
	// 日志内容
	LogContext chan string

	// 日志写入对象（实现了 LogWriter 接口的实例）
	// 因为会向控制台，文本文件，kafka 等组件写入，所以会针对不同的写入方式创建一个 writer，使用 map 进行管理
	Writers map[string]logger.LogWriter
}

// NewCoreWithFileLogConf 创建一个日志核心对象，输出到文件
func NewCoreWithFileLogConf(config *fileWriter.FWConfig) *Core {
	fop := fileWriter.CreateFileWriter(config)
	fileName := filepath.Base(config.Path)

	core := &Core{
		LogContext: make(chan string, 1000),
		Writers:    make(map[string]logger.LogWriter),
	}

	core.Writers[fileName] = fop

	return core
}

// NewCore 创建一个日志核心对象
func NewCore() *Core {
	return &Core{
		LogContext: make(chan string, 1000),
		Writers:    nil,
	}
}
