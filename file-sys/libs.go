package file_sys

import (
	"errors"
	"io/fs"
	"os"
)

// IsExists 判断路径是否存在
func IsExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, fs.ErrNotExist)
}

// MustOpenFile 直接打开文件，使用该方法的前提是确定文件一定存在
func MustOpenFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR, 0666)
	return file, err
}

// CreateFile 创建文件，先检查文件是否存在，存在就报错，不存在就创建
func CreateFile(path, fileName string) (*os.File, error) {
	exist := IsExists(path)
	if !exist {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	// 文件路径
	path = path + "/" + fileName
	exist = IsExists(path)
	if exist {
		return nil, errors.New("文件已存在")
	} else {
		_, err := os.Create(path)
		if err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	return file, nil
}
