package file_sys

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// IsExist 判断路径是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// MustOpenFile 直接打开文件，使用该方法的前提是确定文件一定存在
func MustOpenFile(realPath string) (*os.File, error) {
	file, err := os.OpenFile(realPath, os.O_APPEND|os.O_RDWR, 0666)
	return file, err
}

// CreateFile 创建文件，先检查文件是否存在，存在就报错，不存在就创建
func CreateFile(path string) (*os.File, error) {
	dir := filepath.Dir(path)
	exist := IsExist(dir)
	if !exist {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	// 文件路径
	exist = IsExist(path)
	if !exist {
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

// Remove 删除文件
func Remove(path string) error {
	return os.RemoveAll(path)
}

// CompressToTarGz 打包压缩成 .tar.gz 文件
func CompressToTarGz(src string) error {

	dir := filepath.Dir(src)
	filePrefixName := strings.Split(filepath.Base(src), ".")[0]
	dst := filepath.Join(dir, filePrefixName+".tar.gz")

	// 创建目标文件
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(destFile *os.File) {
		err := destFile.Close()
		if err != nil {
			panic(err)
		}
	}(destFile)

	// 创建Gzip压缩写入器
	gzw := gzip.NewWriter(destFile)
	defer func(gzw *gzip.Writer) {
		err := gzw.Close()
		if err != nil {
			panic(err)
		}
	}(gzw)

	// 创建Tar写入器
	tw := tar.NewWriter(gzw)
	defer func(tw *tar.Writer) {
		err := tw.Close()
		if err != nil {
			panic(err)
		}
	}(tw)

	// 遍历源目录并将文件添加到Tar包中
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 构建文件头信息
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		// 更新文件头中的路径信息
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		header.Name = relPath

		// 写入文件头
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// 如果不是目录，将文件内容写入Tar包
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					panic(err)
				}
			}(file)

			// 将文件内容复制到Tar包中
			_, err = io.Copy(tw, file)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
