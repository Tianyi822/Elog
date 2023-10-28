package file_sys

import "testing"

func TestCreateFileOp(t *testing.T) {
	path := "/Users/chentianyi/Program/Goland-workplace/easy-go-log/log"

	fo := CreateFileOp(path, "test.elog", 1, false)
	err := fo.ready()
	if err != nil {
		t.Fatalf("日志文件准备发生错误: %v", err)
	}
}
