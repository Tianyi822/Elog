package elog

import (
	"fmt"
	"gitee.com/xxc_opensource/elog/core"
	"gitee.com/xxc_opensource/elog/writers/fileWriter"
	"testing"
	"time"
)

func TestElog(t *testing.T) {
	fConf := &fileWriter.FWConfig{
		NeedCompress: true,
		MaxSize:      1,
		Path:         "E:\\MineProgram\\go-workplace\\elog\\log\\test.log",
	}

	c := core.NewCore().AddFileWriter(fConf).AddConsoleWriter().Create()
	elog := NewElog(Panic, c)

	for i := 0; i < 10000; i++ {
		elog.Debug(fmt.Sprintf("DEBUG test %d", i))
		elog.Error(fmt.Sprintf("Error test %d", i))
		elog.Fatal(fmt.Sprintf("Fatal test %d", i))
		elog.Panic(fmt.Sprintf("Panic test %d", i))
		elog.Warn(fmt.Sprintf("Warn test %d", i))
	}
	time.Sleep(2 * time.Second)

	elog.Close()
}
