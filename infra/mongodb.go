package infra

import (
	"context"

	"github.com/thangpham4/self-project/pkg/envx"
	"github.com/thangpham4/self-project/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBConnection() (*mongo.Client, error) {
	l := logger.Factory("Setup MongoDB")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := envx.String("MONGO_ADDR", "")
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// Send a ping to confirm a successful connection
	if err := client.Database("self-project").RunCommand(
		context.TODO(),
		bson.D{primitive.E{Key: "ping", Value: 1}},
	).Err(); err != nil {
		l.V(logger.LogErrorLevel).Error(err, "failed to set up mongodb", "uri", uri)
		return nil, err
	}
	l.V(logger.LogInfoLevel).Info("successfully set up mongodb", "uri", uri)

	return client, nil
}
