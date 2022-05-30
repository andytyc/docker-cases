package conf

import "fmt"

const (
	FlagBlockName  = "block"
	FlagBlockUsage = "主动阻塞程序退出, 默认:不阻塞"
	FlagBlockValue = false
)

var (
	FlagBlock bool
)

const (
	FlagCaseName                 = "case"
	FlagCaseUsage                = "示例编号, 默认:1000"
	FlagCaseValueC1000HelloWorld = 1000
	FlagCaseValueC1001OSENV      = 1001
)

var FlagCaseUsageMap = map[int]string{
	FlagCaseValueC1000HelloWorld: "Hello World !",
	FlagCaseValueC1001OSENV:      "os env",
}

var FlagCase int

func FlagCasePrintOptions() string {
	msg := FlagCaseUsage + "\n" + "选项, 可参考:\n"
	for key, val := range FlagCaseUsageMap {
		msg += fmt.Sprintf("    %d : %s,\n", key, val)
	}
	return msg
}
