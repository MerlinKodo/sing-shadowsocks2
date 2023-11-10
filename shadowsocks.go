package shadowsocks

import (
	"context"

	C "github.com/MerlinKodo/sing-shadowsocks/cipher"
	_ "github.com/MerlinKodo/sing-shadowsocks/shadowaead"
	_ "github.com/MerlinKodo/sing-shadowsocks/shadowaead_2022"
	_ "github.com/MerlinKodo/sing-shadowsocks/shadowstream"
)

type (
	Method        = C.Method
	MethodOptions = C.MethodOptions
)

func CreateMethod(ctx context.Context, method string, options MethodOptions) (Method, error) {
	return C.CreateMethod(ctx, method, options)
}
