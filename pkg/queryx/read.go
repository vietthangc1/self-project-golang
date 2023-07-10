package queryx

import (
	"strings"

	"github.com/thangpham4/self-project/pkg/tokenx"
)

func DecodeQuery(url string) (string, error) {
	tokenStruct := tokenx.NewToken(url, "")
	return tokenStruct.ExtractTokenKey()
}

func ReadToMap(q string) map[string]string {
	out := make(map[string]string)
	pairs := strings.Split(q, "&")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		out[kv[0]] = kv[1]
	}
	return out
}

func ReadMoreLink(url string) (map[string]string, error) {
	out := make(map[string]string)
	query, err := DecodeQuery(url)
	if err != nil {
		return out, err
	}
	out = ReadToMap(query)
	return out, nil
}
