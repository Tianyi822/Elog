package consoleWriter

import (
	"fmt"
	"gitee.com/xxc_opensource/elog/utils"
)

// ConsoleWriter 控制台输出，没必要做什么处理，简单实现 LogWriter 接口即可
type ConsoleWriter struct {
}

func (cs *ConsoleWriter) GetHash() string {
	return utils.GenHash("ConsoleWriter")
}

func (cs *ConsoleWriter) Close() error {
	return nil
}

func (cs *ConsoleWriter) WriteLog(context []byte) error {
	fmt.Println(string(context))
	return nil
}
