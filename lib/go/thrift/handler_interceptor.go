package thrift

import (
	"context"
)

type HandlerInterceptor func(ctx context.Context, arg interface{}) (result interface{}, err error)
