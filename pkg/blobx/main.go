package blobx

import (
	"bytes"
	"context"
	"encoding/csv"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/thangpham4/self-project/pkg/envx"
	"github.com/thangpham4/self-project/pkg/logger"
)

var _ BlobService = &BlobServiceImpl{}

type BlobServiceImpl struct {
	ContainerName string
	Client        *azblob.Client
	Logger        logger.Logger
}

func NewBlobService(client *azblob.Client) *BlobServiceImpl {
	containerName := envx.String("BLOB_CONTAINER", "")
	return &BlobServiceImpl{
		ContainerName: containerName,
		Client:        client,
		Logger:        logger.Factory("BlobService"),
	}
}

type BlobService interface {
	GetCSV(path string) ([][]string, error)
	ListBlob() ([]string, error)
	DownloadToBuffer(path string) (*bytes.Buffer, error)
}

func (b *BlobServiceImpl) GetCSV(path string) ([][]string, error) {
	buf, err := b.DownloadToBuffer(path)
	if err != nil {
		return [][]string{}, err
	}
	reader := csv.NewReader(buf)
	records, err := reader.ReadAll()
	if err != nil {
		b.Logger.Error(err, "file not formated as csv")
		return [][]string{}, err
	}
	return records, nil
}

func (b *BlobServiceImpl) DownloadToBuffer(path string) (*bytes.Buffer, error) {
	ctx := context.Background()
	get, err := b.Client.DownloadStream(ctx, b.ContainerName, path, nil)
	if err != nil {
		b.Logger.Error(err, "error in getting blob", "blob_name", path)
		return nil, err
	}

	downloadedData := bytes.Buffer{}
	retryReader := get.NewRetryReader(ctx, &azblob.RetryReaderOptions{})
	_, err = downloadedData.ReadFrom(retryReader)
	if err != nil {
		b.Logger.Error(err, "error in reading blob", "blob_name", path)
		return nil, err
	}

	err = retryReader.Close()
	if err != nil {
		b.Logger.Error(err, "error in closing blob stream", "blob_name", path)
		return nil, err
	}

	return &downloadedData, nil
}

func (b *BlobServiceImpl) ListBlob() ([]string, error) {
	b.Logger.Info("Listing blob...", "container", b.ContainerName)
	pager := b.Client.NewListBlobsFlatPager(b.ContainerName, &azblob.ListBlobsFlatOptions{
		Include: azblob.ListBlobsInclude{Snapshots: true, Versions: true},
	})

	listBlob := []string{}

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		if err != nil {
			b.Logger.Error(err, "error in getting list blob")
			return listBlob, err
		}
		for _, blob := range resp.Segment.BlobItems {
			listBlob = append(listBlob, *blob.Name)
		}
	}

	return listBlob, nil
}
