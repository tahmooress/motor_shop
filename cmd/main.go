package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tahmooress/motor-shop/internal/adapters/database"
	"github.com/tahmooress/motor-shop/internal/pkg/logger"
	"github.com/tahmooress/motor-shop/internal/port/http"
	"github.com/tahmooress/motor-shop/internal/usecases"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	logger := logger.New(ctx, ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	logger.Info.Println("msg", fmt.Sprintln("run  with commit  & commit date by version"),
		"func", "cmd", "when", "Bootstrapping project")

	iDatabase, err := database.New(ctx)
	if err != nil {
		logger.Error.Println("can't connect to db")
		panic(err)
	}

	iUseCases, err := usecases.New(ctx, iDatabase, logger)
	if err != nil {
		logger.Error.Println("err", err, "msg", "Error occurred for new use cases", "func", "main")
		panic(err)
	}

	httpServer, err := http.New(ctx, iUseCases, logger)
	if err != nil {
		logger.Error.Println("err", err, "msg", "Error occurred for new http server", "func", "main")
		panic(err)
	}

	go interruptHook(ctx, cancelFunc, iDatabase, httpServer, logger)

	os.Exit(1)

}
