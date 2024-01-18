package infra

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/thangpham4/self-project/pkg/envx"
	"github.com/thangpham4/self-project/pkg/logger"
)

func NewBlobConnection() (*azblob.Client, error) {
	l := logger.Factory("Setup Blob Connection")
	// uri := envx.String("BLOB_HOST", "")
	// accessKey := envx.String("BLOB_ACCESS_KEY", "")

	connectionString := envx.String("BLOB_CONNECTION_STRING", "")

	client, err := azblob.NewClientFromConnectionString(connectionString, nil)
	if err != nil {
		l.V(logger.LogErrorLevel).Error(err, "failed to connect Blob")
		return nil, err
	}

	return client, nil
}
