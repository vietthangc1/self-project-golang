package main

import (
	"context"

	"github.com/thangpham4/self-project/pkg/logger"
)

func main() {
	ctx := context.TODO()
	l := logger.Factory("CMD/Worker")
	server, err := BuildServer(ctx)
	if err != nil {
		l.Error(err, "Build Server failed!")
		panic(err.Error())
	}
	go func() {
		err := server.Start(ctx)
		logger.Error(err, "error when starting server")
	}()
}
