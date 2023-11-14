package esWriter

import (
	"gitee.com/xxc_opensource/elog/utils"
	es7 "github.com/elastic/go-elasticsearch/v7"
)

type EsV7Writer struct {
	esClient *es7.Client
	index    string
}

// CreateEsV7WriterByConfig 通过配置项创建一个 v7 版本的 elasticsearch 客户端，并通过这个客户端写入日志
func CreateEsV7WriterByConfig(config es7.Config, index string) *EsV7Writer {
	client, err := es7.NewClient(config)

	if err != nil {
		panic(err)
	}

	return &EsV7Writer{
		index:    index,
		esClient: client,
	}
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
