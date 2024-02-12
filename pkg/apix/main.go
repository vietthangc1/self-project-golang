package apix

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/thangpham4/self-project/pkg/logger"
)

var _ APICaller = &APICallerImpl{}

type APICallerImpl struct {
	client http.Client
	logger logger.Logger
}

type APICaller interface {
	Get(ctx context.Context, uri string, headers map[string]string) ([]byte, error)
	Post(ctx context.Context, uri string, headers map[string]string, data []byte) ([]byte, error)
}

func NewAPICaller() *APICallerImpl {
	return &APICallerImpl{
		logger: logger.Factory("APICaller"),
		client: http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *APICallerImpl) Get(
	ctx context.Context,
	uri string,
	headers map[string]string,
) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, http.NoBody)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err != nil {
		c.logger.Error(err, "error in calling api", "uri", uri)
		return nil, err
	}
	defer resp.Body.Close()
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error(err, "error in calling api", "uri", uri)
		return nil, err
	}
	return responseData, nil
}

func (c *APICallerImpl) Post(
	ctx context.Context,
	uri string,
	headers map[string]string,
	data []byte,
) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(res["json"])
	if err != nil {
		return nil, err
	}
	return body, nil
}
