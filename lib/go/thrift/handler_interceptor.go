package thrift

import (
	"context"
)

type HandlerInterceptor func(ctx context.Context, arg interface{}, next HandlerInterceptor) (result interface{}, err error)
