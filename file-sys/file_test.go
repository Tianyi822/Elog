package file_sys

import "testing"

func TestCreateFileOp(t *testing.T) {
	path := "E:\\MineProgram\\go-workplace\\easy-go-log\\log\\test.log"

	fo := CreateFileOp(path, 1, false)
	err := fo.ready()
	if err != nil {
		t.Fatalf("日志文件准备发生错误: %v", err)
	}
}
