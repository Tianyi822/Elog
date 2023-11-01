package file_sys

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type FileOp struct {
	file           *os.File
	writer         *bufio.Writer
	isOpen         bool   // 用于判断是否可以进行操作
	needCompress   bool   // 是否需要压缩
	maxSize        int    // 以 MB 为单位
	path           string // 文件路径
	filePrefixName string // 文件前缀名
	fileSuffixName string // 文件后缀名
}

// CreateFileOp 只是创建一个文件操作对象，但不代表要立即操作这个文件，所以 isOpen 默认为 false
func CreateFileOp(path string, maxSize int, needCompress bool) *FileOp {
	fileInfo := strings.Split(filepath.Base(path), ".")

	return &FileOp{
		filePrefixName: fileInfo[0],
		fileSuffixName: fileInfo[1],
		path:           path,
		needCompress:   needCompress,
		isOpen:         false,
		maxSize:        maxSize,
	}
}

// ready 用于进行文件操作前的准备工作
func (fo *FileOp) ready() (err error) {
	if fo.file == nil {
		if IsExist(fo.path) {
			fo.file, err = MustOpenFile(fo.path)
			if err != nil {
				return err
			}
		} else {
			fo.file, err = CreateFile(fo.path)
			if err != nil {
				return err
			}
		}
		fo.writer = bufio.NewWriter(fo.file)
	}
	fo.isOpen = true
	return nil
}

// Close 关闭文件
func (fo *FileOp) Close() error {
	fo.isOpen = false

	// 将缓存中的数据落盘
	err := fo.writer.Flush()
	if err != nil {
		return err
	}

	err = fo.file.Close()
	if err != nil {
		return err
	}

	fo.file = nil
	fo.writer = nil
	return err
}

// needSplit 判断是否需要进行分片
func (fo *FileOp) needSplit() bool {
	// 判断是否需要进行分片
	if fo.maxSize <= 0 {
		return false
	}

	// 判断文件大小是否超过最大值
	fileInfo, err := fo.file.Stat()
	if err != nil {
		return false
	}

	return fileInfo.Size() > int64(fo.maxSize*1024*1024)
}

// Write 写入数据
// 该函数不做并发处理，传入的数据都是通过 channel 传递过来的，所以不需要考虑并发问题
// 并不会出现多个协程往同一个文件里面写数据，文件操作模块主要集中于对日志文件的分片管理，对历史日志打包
func (fo *FileOp) Write(context []byte) error {
	if !fo.isOpen {
		return fo.ready()
	}

	// 判断是否需要进行分片
	if fo.needSplit() {
		// 关闭文件
		err := fo.Close()
		if err != nil {
			return err
		}

		// 压缩文件
		if fo.needCompress {
			// 获取文件前缀名\
			src := filepath.Join(fo.dirPath, fo.fileName)
			err = CompressToTarGz(src)
			if err != nil {
				return err
			}
		}

		// 重新打开文件
		err = fo.ready()
		if err != nil {
			return err
		}
	}

	// 写入数据
	buf := append(context, '\n')
	_, err := fo.writer.Write(buf)
	return err
}
