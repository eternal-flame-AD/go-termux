package debug

import (
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
)

func StackOnExit() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		<-c
		pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
		os.Exit(0)
	}()
}
