package core

import (
	"fmt"
	"gitee.com/xxc_opensource/elog/writers/fileWriter"
	"testing"
	"time"
)

func TestNewCore(t *testing.T) {
	fConf := &fileWriter.FWConfig{
		NeedCompress: true,
		MaxSize:      1,
		Path:         "/Users/chentianyi/Program/Goland-workplace/easy-go-log/log/test.log",
	}

	core := NewCore().AddFileWriter(fConf).AddConsoleWriter().Create()

	fmt.Println(core)
}

func TestCore_Write(t *testing.T) {
	fConf := &fileWriter.FWConfig{
		NeedCompress: true,
		MaxSize:      1,
		Path:         "/Users/chentianyi/Program/Goland-workplace/easy-go-log/log/test.log",
	}

	core := NewCore().AddFileWriter(fConf).AddConsoleWriter().Create()
	core.Start()

	for i := 0; i < 10000; i++ {
		core.Write([]byte(fmt.Sprintf("test %d", i)))
	}
	time.Sleep(2 * time.Second)

	core.Close()
}
