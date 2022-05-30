package main

import (
	"flag"
	"fmt"

	"github.com/andytyc/docker-cases/casetask"
	"github.com/andytyc/docker-cases/conf"
	"github.com/andytyc/docker-cases/utils"
)

func init() {
	flag.BoolVar(&conf.FlagBlock, conf.FlagBlockName, conf.FlagBlockValue, conf.FlagBlockUsage)
	flag.IntVar(&conf.FlagCase, conf.FlagCaseName, conf.FlagCaseValueC1000HelloWorld, conf.FlagCasePrintOptions())

	flag.Parse()
}

func main() {
	switch conf.FlagCase {
	case conf.FlagCaseValueC1000HelloWorld:
		casetask.C1000HelloWorld()
	case conf.FlagCaseValueC1001OSENV:
		casetask.C1001OSENV()
	default:
		fmt.Printf("value of param [%s] is invaild : %d\n", conf.FlagCaseName, conf.FlagCase)
	}
	if conf.FlagBlock {
		utils.ListenSignal()
	}
}
