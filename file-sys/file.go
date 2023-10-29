package file_sys

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileOp struct {
	file         *os.File
	isOpen       bool // 用于判断是否可以进行操作
	needCompress bool // 是否需要压缩
	maxSize      int  // 以 MB 为单位
	curDate      time.Time
	dirPath      string // 文件保存路径
	fileName     string // 文件保存名称
	realPath     string // 文件真实路径
}

// CreateFileOp 只是创建一个文件操作对象，但不代表要立即操作这个文件，所以 isOpen 默认为 false
func CreateFileOp(path, fileName string, maxSize int, needCompress bool) *FileOp {
	return &FileOp{
		dirPath:      path,
		fileName:     fileName,
		realPath:     filepath.Join(path, fileName),
		needCompress: needCompress,
		isOpen:       false,
		maxSize:      maxSize,
	}
}

// ready 用于进行文件操作前的准备工作
func (fo *FileOp) ready() (err error) {
	if fo.file == nil {
		if IsExist(fo.realPath) {
			fo.file, err = MustOpenFile(fo.realPath)
			if err != nil {
				return err
			}
		} else {
			fo.file, err = CreateFile(fo.dirPath, fo.fileName)
			if err != nil {
				return err
			}
		}
	}
	fo.isOpen = true
	fo.curDate = time.Now()
	return nil
}

// Close 关闭文件
func (fo *FileOp) Close() error {
	fo.isOpen = false
	err := fo.file.Close()
	fo.file = nil
	return err
}

// needSplit 判断是否需要进行分片
func (fo *FileOp) needSplit() bool {
	// 判断是否需要进行分片
	if fo.maxSize <= 0 {
		return false
	}

	// 判断是否需要进行分片
	if fo.curDate.Day() != time.Now().Day() {
		return true
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
			// 获取文件名和后缀
			fInfo := strings.Split(fo.fileName, ".")
			dst := filepath.Join(fo.dirPath, fInfo[0]+"_"+fo.curDate.Format("2006-01-02_15:04:05")+"."+fInfo[1])
			err = Compress(dst+".zip", dst)
			if err != nil {
				return err
			}
			// 压缩之后删除原来的文件
			err := Remove(dst)
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
	_, err := fo.file.Write(buf)
	return err
}
