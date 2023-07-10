package queryx

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/thangpham4/self-project/pkg/tokenx"
)

func DecodeQuery(uri string) (string, error) {
	tokenStruct := tokenx.NewToken(uri, "")
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

func ReadMoreLink(uri string) (map[string]string, error) {
	out := make(map[string]string)
	query, err := DecodeQuery(uri)
	if err != nil {
		return out, err
	}
	out = ReadToMap(query)
	return out, nil
}

func ReadAndParseIntVariable(params url.Values, varName string, i *int32) error {
	_str := params.Get(varName)
	if _str == "" {
		return nil
	}
	_int, err := strconv.ParseInt(_str, 10, 32)
	if err != nil {
		return err
	}
	_int32 := int32(_int)
	*i = _int32
	return nil
}
