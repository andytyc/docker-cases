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
	FlagCaseName            = "case"
	FlagCaseUsage           = "示例编号, 默认:100, 一个hello world示例"
	FlagCaseValueHelloWorld = 100
)

var FlagCase int
