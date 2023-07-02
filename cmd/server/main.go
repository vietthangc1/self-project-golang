package main

import (
	"context"
	"net/http"

	"github.com/thangpham4/self-project/pkg/envx"
	"github.com/thangpham4/self-project/pkg/logger"
)

const (
	defaultHost string = "localhost:8080"
)

func main() {
	ctx := context.TODO()
	l := logger.Factory("CMD/Server")
	server, err := BuildServer(ctx)
	if err != nil {
		l.Error(err, "Build Server failed!")
		panic(err.Error())
	}
	host := envx.String("HOST", defaultHost)
	l.Info("Prepare for running server", "host", host)
	err = http.ListenAndServe(host, server)
	if err != nil {
		panic(err.Error())
	}
}
