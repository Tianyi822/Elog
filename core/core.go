package core

import (
	"easy-go-log/fileWriter"
	"easy-go-log/interface/logger"
)

// LogLevel 日志级别
type LogLevel byte

const (
	LevelDebug LogLevel = 1 << iota // 1
	LevelInfo                       // 2
	LevelWarn                       // 4
	LevelError                      // 8
	LevelFatal                      // 16
	LevelPanic                      // 32
)

// OutPutType 输出方式
type OutPutType byte

// 用于标识不同的输出模式
const (
	OutPutToConsole OutPutType = 1 << iota
	OutPutToFile
	OutPutToKafka
)

// Core 日志核心组件，用于对接各种输出路径，包含但不限于日志文本文件，MQ 消息队列，Kafka 消息队列，HDFS 分布式集群等
// 这个组件中还要处理一些并发操作，防止出现日志并发写入的问题
type Core struct {
	// 日志内容
	LogChannel chan string

	// 日志写入对象（实现了 LogWriter 接口的实例）
	// 因为会向控制台，文本文件，kafka 等组件写入，所以会针对不同的写入方式创建一个 writer，使用 map 进行管理
	Writers map[string]logger.LogWriter

	// 用于标识该核心是否已经配置完成
	isOk bool
}

// NewCore 无参数配置，默认输出到控制台，且通道默认容量为 2233
func NewCore() *Core {
	return &Core{
		LogChannel: make(chan string, 2233),
		Writers:    nil,
		isOk:       false,
	}
}

// NewBufferCore 创建一个带有缓冲区的 Core
func NewBufferCore(buf int) *Core {
	return &Core{
		LogChannel: make(chan string, buf),
		Writers:    nil,
		isOk:       false,
	}
}

// Create 表示 Core 配置装配已完成，可以使用
func (c *Core) Create() *Core {
	c.isOk = true
	return c
}

// ConfigFileWriter 用于装配文件写入模块
func (c *Core) ConfigFileWriter(fConf *fileWriter.FWConfig) *Core {
	fWriter := fileWriter.CreateFileWriter(fConf)

	c.Writers[fWriter.GetHash()] = fWriter

	return c
}
