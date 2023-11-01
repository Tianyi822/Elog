package core

import filesys "easy-go-log/file-sys"

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

// Config 核心配置项
type Config struct {
	Level        LogLevel   // 日志级别
	NeedCompress bool       // 是否需要压缩
	MaxSize      int        // 以 MB 为单位
	DirPath      string     // 文件保存路径
	FileName     string     // 文件保存名称
	OutPutTo     OutPutType // 日志输出位置
}

// Core 日志核心组件，用于对接各种输出路径，包含但不限于日志文本文件，MQ 消息队列，Kafka 消息队列，HDFS 分布式集群等
// 这个组件中还要处理一些并发操作，防止出现日志并发写入的问题
type Core struct {
	LogContext chan string     // 日志内容
	fileOp     *filesys.FileOp // 文件操作对象
	*Config                    // 核心配置项
}

// NewCoreWithConf 创建一个日志核心对象
func NewCoreWithConf(config *Config) *Core {
	fop := filesys.CreateFileOp(config.DirPath, config.FileName, config.MaxSize, config.NeedCompress)

	return &Core{
		LogContext: make(chan string, 1000),
		fileOp:     fop,
		Config:     config,
	}
}

// NewCore 创建一个日志核心对象
func NewCore() *Core {
	return &Core{
		LogContext: make(chan string, 1000),
		fileOp:     nil,
		Config: &Config{
			Level:    LevelDebug,
			OutPutTo: OutPutToConsole,
		},
	}
}

// Output 将日志内容写入到文件中
func (c *Core) Output() {
	for {
		select {
		case log := <-c.LogContext:
			err := c.fileOp.Write([]byte(log))
			if err != nil {
				panic(err)
			}
		}
	}
}
