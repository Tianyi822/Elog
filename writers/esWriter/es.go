package esWriter

import (
	"errors"
	"fmt"
	"gitee.com/xxc_opensource/elog/utils"
	es7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io"
	"sync"
)

type EsV7Writer struct {
	esClient *es7.Client
	index    string
	isReady  bool
}

var createEs7ClientOnce sync.Once

// CreateEsV7WriterByConfig 通过配置项创建一个 v7 版本的 elasticsearch 客户端，并通过这个客户端写入日志
func CreateEsV7WriterByConfig(config es7.Config, index string) *EsV7Writer {
	writer := &EsV7Writer{
		index:   index,
		isReady: false,
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
		isReady:  false,
	}
}

func (ew *EsV7Writer) GetHash() string {
	return utils.GenHash(ew.index)
}

// ready 因为创建了 EsWriter 后，并不会直接进行操作
// 当第一次操作的时候，需要使用 ready 方法让 writer 准备好
// EsWriter 的 ready 函数会先检查一下 ES 中是否确实存在指定的 index，不存在则创建
func (ew *EsV7Writer) ready() (err error) {
	req := esapi.IndicesExistsAliasRequest{
		Index: []string{ew.index},
	}

	// TODO: 目前没有使用 context 对整个日志框架进行控制，所以这里的第一个参数先传入一个 nil
	res, err := req.Do(nil, ew.esClient)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	if res.IsError() {
		return errors.New(fmt.Sprintf("Error checking index existence: %s", res.String()))
	}

	if res.StatusCode == 200 {
		ew.isReady = true
	} else {
		// TODO: index 不存在则需要创建对应的index
	}

	return nil
}

func (ew *EsV7Writer) Close() error {
	return nil
}

func (ew *EsV7Writer) WriteLog(content []byte) error {
	return nil
}
