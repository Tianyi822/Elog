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
	path         string // 文件保存路径
	fileName     string // 文件保存名称
}

// CreateFileOp 只是创建一个文件操作对象，但不代表要立即操作这个文件，所以 isOpen 默认为 false
func CreateFileOp(path, fileName string, maxSize int, needCompress bool) *FileOp {
	return &FileOp{
		path:         path,
		fileName:     fileName,
		needCompress: needCompress,
		isOpen:       false,
		maxSize:      maxSize,
	}
}

// ready 用于进行文件操作前的准备工作
func (fo *FileOp) ready() (err error) {
	if fo.file == nil {
		if IsExists(fo.path) {
			fo.file, err = MustOpenFile(fo.path)
			if err != nil {
				return err
			}
		} else {
			fo.file, err = CreateFile(fo.path, fo.fileName)
			if err != nil {
				return err
			}
		}
	}
	fo.isOpen = true
	fo.curDate = time.Now()
	return nil
}
