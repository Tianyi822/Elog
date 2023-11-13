package esWriter

import "github.com/elastic/go-elasticsearch/v7"

type EsV7Writer struct {
	esClient *elasticsearch.Client
	index    string
}

// V7Config 用于配置 ElasticSearch v7 版本的客户端配置
type V7Config struct {
	Addresses            []string
	Username             string // BasicAuth 身份验证
	Pwd                  string // BasicAuth 身份验证
	DiscoverNodesOnStart bool   // 启东时节点嗅探功能
	CompressRequestBody  bool   // 请求压缩
}

// CreateEsV7WriterWithConfig 通过配置项创建一个 v7 版本的 elasticsearch 客户端，并通过这个客户端写入日志
func CreateEsV7WriterWithConfig(config *V7Config, index string) *EsV7Writer {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses:            config.Addresses,
		Username:             config.Username,
		Password:             config.Pwd,
		DiscoverNodesOnStart: config.DiscoverNodesOnStart,
		CompressRequestBody:  config.CompressRequestBody,
	})

	if err != nil {
		panic(err)
	}

	return &EsV7Writer{
		index:    index,
		esClient: client,
	}
}

func (ew *EsV7Writer) GetHash() string {
	return ""
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
