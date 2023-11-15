package esWriter

type BufferSize uint32

const (
	ByteSize  BufferSize = 1
	KByteSize            = 1000 * ByteSize
	MByteSize            = 1000 * KByteSize
	GByteSize            = 1000 * MByteSize
)

type Interval uint8

const (
	Second Interval = 1
	Minute          = 60 * Second
	Hour            = 60 * Minute
	Day             = 24 * Hour
)

// IndexConfig Elasticsearch 中创建索引时需要添加的配置项
type IndexConfig struct {
	Name               string         // 索引名称
	Mapping            map[string]any // 索引映射对象
	NumberOfShards     int            // 索引的主分片数量
	NumberOfReplicas   int            // 每个主分片的副本数量
	RefreshInterval    Interval       // 刷新间隔，控制索引的刷新频率
	MaxResultWindow    int            // 控制单个查询的结果窗口大小
	IndexingBufferSize BufferSize     // 写入索引时的缓冲区大小
}
