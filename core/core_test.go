package core

import (
	"easy-go-log/fileWriter"
	"fmt"
	"testing"
	"time"
)

func TestNewCore(t *testing.T) {
	fConf := &fileWriter.FWConfig{
		NeedCompress: true,
		MaxSize:      1,
		Path:         "/Users/chentianyi/Program/Goland-workplace/easy-go-log/log/test.log",
	}
	ch := make(chan []byte, 100)

	core := NewCore(ch).AddFileWriter(fConf).AddConsoleWriter().Create()

	fmt.Println(core)
}

func TestCore_Write(t *testing.T) {
	fConf := &fileWriter.FWConfig{
		NeedCompress: true,
		MaxSize:      1,
		Path:         "/Users/chentianyi/Program/Goland-workplace/easy-go-log/log/test.log",
	}
	ch := make(chan []byte, 10)

	core := NewCore(ch).AddFileWriter(fConf).AddConsoleWriter().Create()
	core.Start()

	for i := 0; i < 10000; i++ {
		content := fmt.Sprintf("test %d", i)
		ch <- []byte(content)
	}
	time.Sleep(5 * time.Second)

	close(ch)
}
