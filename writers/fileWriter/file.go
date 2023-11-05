package fileWriter

import (
	"bufio"
	"fmt"
	"gitee.com/xxc_opensource/elog/uitls"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type FileWriter struct {
	file           *os.File
	writer         *bufio.Writer
	isOpen         bool   // 用于判断是否可以进行操作
	needCompress   bool   // 是否需要压缩
	maxSize        int    // 以 MB 为单位
	hash           string // 每个 FileWriter 有且只有一个唯一的 hash 值
	path           string // 文件路径
	filePrefixName string // 文件前缀名
	fileSuffixName string // 文件后缀名
}

// FWConfig 日志文件配置项
type FWConfig struct {
	NeedCompress bool   // 是否需要压缩
	MaxSize      int    // 以 MB 为单位
	Path         string // 文件保存路径
}

// CreateFileWriter 只是创建一个文件操作对象，但不代表要立即操作这个文件，所以 isOpen 默认为 false
func CreateFileWriter(config *FWConfig) *FileWriter {
	fileInfo := strings.Split(filepath.Base(config.Path), ".")

	return &FileWriter{
		hash:           uitls.GenHash(config.Path),
		filePrefixName: fileInfo[0],
		fileSuffixName: fileInfo[1],
		path:           config.Path,
		needCompress:   config.NeedCompress,
		isOpen:         false,
		maxSize:        config.MaxSize,
	}
}

func (fw *FileWriter) GetHash() string {
	return fw.hash
}

// ready 用于进行文件操作前的准备工作
func (fw *FileWriter) ready() (err error) {
	if fw.file == nil {
		if isExist(fw.path) {
			fw.file, err = mustOpenFile(fw.path)
			if err != nil {
				return err
			}
		} else {
			fw.file, err = createFile(fw.path)
			if err != nil {
				return err
			}
		}
		fw.writer = bufio.NewWriter(fw.file)
	}
	fw.isOpen = true
	return nil
}

// Close 关闭文件
func (fw *FileWriter) Close() error {
	fw.isOpen = false

	// 将缓存中的数据落盘
	err := fw.writer.Flush()
	if err != nil {
		return err
	}

	err = fw.file.Close()
	if err != nil {
		return err
	}

	fw.file = nil
	fw.writer = nil
	return err
}

// needSplit 判断是否需要进行分片
func (fw *FileWriter) needSplit() bool {
	// 判断是否需要进行分片
	if fw.maxSize <= 0 {
		return false
	}

	// 判断文件大小是否超过最大值
	fileInfo, err := fw.file.Stat()
	if err != nil {
		return false
	}

	return fileInfo.Size() > int64(fw.maxSize*1024*1024)
}

// WriteLog 写入日志数据
// 该函数不做并发处理，传入的数据都是通过 channel 传递过来的，所以不需要考虑并发问题
// 并不会出现多个协程往同一个文件里面写数据，文件操作模块主要集中于对日志文件的分片管理，对历史日志打包
func (fw *FileWriter) WriteLog(context []byte) error {
	if !fw.isOpen {
		return fw.ready()
	}

	// 判断是否需要进行分片
	if fw.needSplit() {
		// 关闭文件
		err := fw.Close()
		if err != nil {
			return err
		}

		// 修改文件名
		newFileName := fmt.Sprintf("%v_%v.%v", fw.filePrefixName, strconv.FormatInt(time.Now().Unix(), 10), fw.fileSuffixName)
		destPath := filepath.Join(filepath.Dir(fw.path), newFileName)
		err = os.Rename(fw.path, destPath)
		if err != nil {
			return err
		}

		// 压缩文件
		if fw.needCompress {
			err = compressFileToTarGz(destPath)
			if err != nil {
				return err
			}
			// 删除原文件
			err = remove(destPath)
			if err != nil {
				return err
			}
		}

		// 重新打开文件
		err = fw.ready()
		if err != nil {
			return err
		}
	}

	// 写入数据
	buf := append(context, '\n')
	_, err := fw.writer.Write(buf)
	if err != nil {
		return err
	}
	// 数据落盘
	err = fw.writer.Flush()
	if err != nil {
		return err
	}
	return err
}
