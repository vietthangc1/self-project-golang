package tokenx

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/pkg/envx"
)

const (
	tokenHourLifespanDefault = 24
)

type Token struct {
	Token    string
	Lifespan int32
	Key      string
}

func NewToken(
	token string,
	key string,
) *Token {
	return &Token{
		Token:    token,
		Key:      key,
		Lifespan: envx.Int("TOKEN_HOUR_LIFESPAN", tokenHourLifespanDefault),
	}
}

func (t *Token) GenerateToken() (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["key"] = t.Key
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(t.Lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(envx.String("TOKEN_SECRET_KEY", "")))
}

func (t *Token) CheckValidToken() error {
	tokenString := t.Token
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, commonx.ErrorMessages(
				commonx.ErrWrongMethod,
				fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]),
			)
		}
		return []byte(envx.String("TOKEN_SECRET_KEY", "")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *Token) ExtractTokenKey() (string, error) {
	tokenString := t.Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, commonx.ErrorMessages(
				commonx.ErrWrongMethod,
				fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]),
			)
		}
		return []byte(envx.String("TOKEN_SECRET_KEY", "")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		key, exist := claims["key"].(string)
		if !exist {
			return "", commonx.ErrorMessages(commonx.ErrKeyNotFound, fmt.Sprintf("not found key: %s", key))
		}
		return key, nil
	}
	return "", nil
}
