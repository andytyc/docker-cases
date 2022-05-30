package conf

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
	FlagCaseUsage                = "示例编号, 默认:1000, 一个hello world示例"
	FlagCaseValueC1000HelloWorld = 1000
	FlagCaseValueC1001OSENV      = 1001
)

var FlagCase int
