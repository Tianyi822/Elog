package fileSys

import (
	"fmt"
	"testing"
)

func TestCreateFileOp(t *testing.T) {
	path := "/Users/chentianyi/Program/Goland-workplace/easy-go-log/log/test.log"

	fo := CreateFileOp(path, 1, false)
	err := fo.ready()
	if err != nil {
		t.Fatalf("日志文件准备发生错误: %v", err)
	}
}

func TestIsExists(t *testing.T) {
	path := "/Users/chentianyi/Program/Goland-workplace/easy-go-log/log/"

	exist := IsExist(path)
	if !exist {
		t.Fatalf("路径不存在: %v", exist)
	}
	fmt.Printf("路径存在: %v", exist)
}

func TestFileOp_Write(t *testing.T) {
	path := "/Users/chentianyi/Program/Goland-workplace/easy-go-log/log/test.log"

	fo := CreateFileOp(path, 1, true)
	for {
		err := fo.Write([]byte("test"))
		if err != nil {
			t.Fatalf("写入日志发生错误: %v", err)
		}
	}
}
