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
	err = server.Start(ctx)
	defer server.Stop()
	if err != nil {
		panic(err.Error())
	}
}
