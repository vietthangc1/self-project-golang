package main

import (
	"context"
	"net/http"

	"github.com/thangpham4/self-project/config"
	"github.com/thangpham4/self-project/pkg/logger"
)

func main() {
	ctx := context.TODO()
	l := logger.Factory("CMD/Server")
	server, err := BuildServer(ctx)
	if err != nil {
		l.Error(err, "Build Server failed!")
		panic(err.Error())
	}
	host := config.Domain
	l.Info("Prepare for running server", "host", host)
	err = http.ListenAndServe(host, server)
	if err != nil {
		panic(err.Error())
	}
}
