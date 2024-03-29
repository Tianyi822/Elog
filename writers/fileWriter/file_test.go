package fileWriter

import (
	"fmt"
	"testing"
)

func TestCreateFileOp(t *testing.T) {
	config := &FWConfig{
		Path:         "/Users/chentianyi/Program/Goland-workplace/easy-go-log/log/test.log",
		MaxSize:      1,
		NeedCompress: false,
	}

	fo := CreateFileWriter(config)
	err := fo.ready()
	if err != nil {
		t.Fatalf("日志文件准备发生错误: %v", err)
	}
}

func TestIsExists(t *testing.T) {
	path := "/Users/chentianyi/Program/Goland-workplace/easy-go-log/log/"

	exist := isExist(path)
	if !exist {
		t.Fatalf("路径不存在: %v", exist)
	}
	fmt.Printf("路径存在: %v", exist)
}

func TestFileOp_Write(t *testing.T) {
	config := &FWConfig{
		Path:         "E:\\MineProgram\\go-workplace\\elog\\log\\test.log",
		MaxSize:      1,
		NeedCompress: true,
	}

	fo := CreateFileWriter(config)
	times := 0
	for {
		err := fo.WriteLog([]byte("chentyit OHHHHHHHHH ttttt"))
		if err != nil {
			t.Fatalf("写入日志发生错误: %v", err)
		}
		times += 1
		if times == 100_0000 {
			err = fo.Close()
			if err != nil {
				t.Fatalf("结束写入发生错误: %v", err)
			}
			break
		}
	}
}

func TestFileOp_Write2(t *testing.T) {
	config := &FWConfig{
		Path:         "E:\\MineProgram\\go-workplace\\elog\\log\\test.log",
		MaxSize:      1,
		NeedCompress: true,
	}

	fo := CreateFileWriter(config)
	err := fo.WriteLog([]byte("chentyit OHHHHHHHHH ttttt"))
	if err != nil {
		t.Fatalf("写入日志发生错误: %v", err)
	}
}
