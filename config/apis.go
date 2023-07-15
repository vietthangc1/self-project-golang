package config

import "github.com/thangpham4/self-project/pkg/envx"

var (
	Domain     = envx.String("HOST", "localhost:8080")
	ProductAPI = envx.String("PRODUCT_HOST", "localhost:8081")
)
