package file_sys

import (
	"os"
	"time"
)

type FileOp struct {
	file         *os.File
	isOpen       bool // 用于判断是否可以进行操作
	needCompress bool // 是否需要压缩
	maxSize      int  // 以 MB 为单位
	curDate      time.Time
	path         string
}
