package main

import (
	server "github.com/tubfuzzy/banraiphisan-reservation/pkg/server"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	serv, err := server.New()
	if err != nil {
		panic(err)
	}

	if err := serv.App().Listen(serv.Config().Server.Port); err != nil {
		serv.Logger().Fatalf("%s", err)
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	err = serv.App().Shutdown()

	defer serv.Cache().Close()
	// defer serv.DB().Close()

	if err != nil {
		serv.Logger().Fatalf("%s", err)
	}
}
