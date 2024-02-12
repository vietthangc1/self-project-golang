package config

import "github.com/thangpham4/self-project/pkg/envx"

var (
	Domain        = envx.String("HOST", "0.0.0.0:8080")
	ProductAPI    = envx.String("PRODUCT_HOST", "0.0.0.0:8081")
	ProductAPIExt = envx.String("PRODUCT_HOST_EXT", "0.0.0.0:8082")
)
