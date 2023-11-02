package core

import (
	"easy-go-log/fileWriter"
	"fmt"
	"testing"
)

func TestNewCore(t *testing.T) {
	fConf := &fileWriter.FWConfig{
		NeedCompress: true,
		MaxSize:      1,
		Path:         "/Users/chentianyi/Program/Goland-workplace/easy-go-log/log/test.log",
	}

	core := NewBufferCore(666).AddFileWriter(fConf).AddConsoleWriter().Create()

	fmt.Println(core)
}
