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
	flag.IntVar(&conf.FlagCase, conf.FlagCaseName, conf.FlagCaseValueHelloWorld, conf.FlagCaseUsage)

	flag.Parse()
}

func main() {
	switch conf.FlagCase {
	case conf.FlagCaseValueHelloWorld:
		casetask.C100HelloWorld()
	default:
		fmt.Printf("value of param [%s] is invaild : %d\n", conf.FlagCaseName, conf.FlagCase)
	}
	if conf.FlagBlock {
		utils.ListenSignal()
	}
}
