package core

import (
	"easy-go-log/consoleWriter"
	"easy-go-log/fileWriter"
	"easy-go-log/logger"
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
}

// Config 链式配置 Core 工具
type Config struct {
	core *Core
	isOk bool // 用于标识该核心是否已经配置完成
}

// NewCore 无参数配置，且通道默认容量为 2233
func NewCore() *Config {
	core := &Core{
		LogChannel: make(chan string, 2233),
		Writers:    nil,
	}

	return &Config{
		core: core,
		isOk: false,
	}
}

// NewBufferCore 创建一个带有缓冲区的配置
func NewBufferCore(buf int) *Config {
	core := &Core{
		LogChannel: make(chan string, buf),
		Writers:    nil,
	}

	return &Config{
		core: core,
		isOk: false,
	}
}

// Create 表示 Core 配置装配已完成，可以使用
func (conf *Config) Create() *Core {
	conf.isOk = true
	return conf.core
}

// checkWriters 用于检查 Writers 是否被初始化
func (conf *Config) addWriter(writer logger.LogWriter) *Config {
	if conf.core.Writers == nil {
		conf.core.Writers = make(map[string]logger.LogWriter)
	}
	conf.core.Writers[writer.GetHash()] = writer

	return conf
}

// AddConsoleWriter 装配控制台写入模块
func (conf *Config) AddConsoleWriter() *Config {
	cWriter := &consoleWriter.ConsoleWriter{}
	conf.addWriter(cWriter)

	return conf
}

// AddFileWriter 装配文件写入模块
func (conf *Config) AddFileWriter(fConf *fileWriter.FWConfig) *Config {
	fWriter := fileWriter.CreateFileWriter(fConf)
	conf.addWriter(fWriter)

	return conf
}
