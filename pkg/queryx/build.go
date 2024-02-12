package queryx

import (
	"fmt"
	"strings"

	"github.com/thangpham4/self-project/pkg/tokenx"
)

func BuildFromMap(_map map[string]string) string {
	_str := ""
	for k, v := range _map {
		_str += fmt.Sprintf("%s=%s", k, v)
		_str += "&"
	}
	return strings.TrimRight(_str, "&")
}

func EncodeQuery(query string) (string, error) {
	tokenStruct := tokenx.NewToken("", query)
	return tokenStruct.GenerateToken()
}

func GenerateMoreLink(_map map[string]string) (string, error) {
	q := BuildFromMap(_map)
	url, err := EncodeQuery(q)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("page_token=%s", url), nil
}
