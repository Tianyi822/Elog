package esWriter

import (
	"gitee.com/xxc_opensource/elog/utils"
	es7 "github.com/elastic/go-elasticsearch/v7"
	"sync"
)

type EsV7Writer struct {
	esClient *es7.Client
	index    string
}

var createEs7ClientOnce sync.Once

// CreateEsV7WriterByConfig 通过配置项创建一个 v7 版本的 elasticsearch 客户端，并通过这个客户端写入日志
func CreateEsV7WriterByConfig(config es7.Config, index string) *EsV7Writer {
	writer := &EsV7Writer{
		index: index,
	}

	// 如果让日志自己来初始化一个 ES 的客户端，那在整个项目的生命周期中，只执行一次，且没有多余的操作
	// 所以还是建议直接传入一个已经配置好了的 ES 客户端，对 EsWriter 进行初始化
	createEs7ClientOnce.Do(func() {
		var err error
		writer.esClient, err = es7.NewClient(config)

		if err != nil {
			panic(err)
		}
	})

	return writer
}

// CreateEsV7WriterByClient 通过一个 v7 版本的 elasticsearch 客户端写入日志
func CreateEsV7WriterByClient(client *es7.Client, index string) *EsV7Writer {
	return &EsV7Writer{
		index:    index,
		esClient: client,
	}
}

func (ew *EsV7Writer) GetHash() string {
	return utils.GenHash(ew.index)
}

func (ew *EsV7Writer) ready() error {
	return nil
}

func (ew *EsV7Writer) Close() error {
	return nil
}

func (ew *EsV7Writer) WriteLog(content []byte) error {
	return nil
}
