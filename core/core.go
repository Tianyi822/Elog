package core

import (
	"gitee.com/xxc_opensource/elog"
	"gitee.com/xxc_opensource/elog/consoleWriter"
	"gitee.com/xxc_opensource/elog/fileWriter"
)

// Core 日志核心组件，用于对接各种输出路径，包含但不限于日志文本文件，MQ 消息队列，Kafka 消息队列，HDFS 分布式集群等
// 这个组件中还要处理一些并发操作，防止出现日志并发写入的问题
type Core struct {
	// 日志内容
	logChannel chan []byte

	// 日志写入对象（实现了 LogWriter 接口的实例）
	// 因为会向控制台，文本文件，kafka 等组件写入，所以会针对不同的写入方式创建一个 writer，使用 map 进行管理
	writers map[string]writers.LogWriter
}

// writeLog 用于写入日志
func (c *Core) writeLog(content []byte) {
	for _, writer := range c.writers {
		err := writer.WriteLog(content)
		if err != nil {
			panic(err)
		}
	}
}

// closeWriters 用于关闭所有的写入器
func (c *Core) closeWriters() {
	for _, writer := range c.writers {
		err := writer.Close()
		if err != nil {
			panic(err)
		}
	}
}

// Start 启动日志核心
func (c *Core) Start() {
	go func() {
		for {
			select {
			case content, ok := <-c.logChannel:
				if ok {
					c.writeLog(content)
				} else {
					c.writeLog([]byte("The log channel is closed and the log core exits"))
					c.closeWriters()
					return
				}
			}
		}
	}()
}

// Write 向日志核心写入日志
func (c *Core) Write(content []byte) {
	c.logChannel <- content
}

// Close 关闭日志核心
func (c *Core) Close() {
	close(c.logChannel)
}

// Config 链式配置 Core 工具
type Config struct {
	core *Core
	isOk bool // 用于标识该核心是否已经配置完成
}

// NewCore 创建一个 Core 实例
func NewCore() *Config {
	core := &Core{
		logChannel: make(chan []byte, 256),
		writers:    nil,
	}

	return &Config{
		core: core,
		isOk: false,
	}
}

// NewCoreWithBuffer 创建一个 Core 实例，指定缓冲区大小
func NewCoreWithBuffer(buf int) *Config {
	return &Config{
		core: &Core{
			logChannel: make(chan []byte, buf),
			writers:    nil,
		},
		isOk: false,
	}
}

// Create 表示 Core 配置装配已完成，可以使用
func (conf *Config) Create() *Core {
	conf.isOk = true
	return conf.core
}

// checkWriters 用于检查 writers 是否被初始化
func (conf *Config) addWriter(writer writers.LogWriter) *Config {
	if conf.core.writers == nil {
		conf.core.writers = make(map[string]writers.LogWriter)
	}
	conf.core.writers[writer.GetHash()] = writer

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
