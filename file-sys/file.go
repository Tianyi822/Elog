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

// CreateFileOp 只是创建一个文件操作对象，但不代表要立即操作这个文件，所以 isOpen 默认为 false
func CreateFileOp(path string, maxSize int, needCompress bool) *FileOp {
	return &FileOp{
		path:         path,
		needCompress: needCompress,
		isOpen:       false,
		maxSize:      maxSize,
	}
}
