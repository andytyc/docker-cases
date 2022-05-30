package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func ListenSignal() os.Signal {
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	return <-chSig
}
