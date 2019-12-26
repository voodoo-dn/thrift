package thrift

import (
	"context"
)

type HandlerInterceptor func(ctx context.Context, arg interface{}, next HandlerInterceptor) (result interface{}, err error)

func ChainedHandlerInterceptor(interceptors ...HandlerInterceptor) HandlerInterceptor {
	interceptorCount := len(interceptors)

	if interceptorCount == 0 {
		return nil
	} else {
		interceptorCount--
	}

	chainer := func(currentInterceptor HandlerInterceptor, nextInterceptor HandlerInterceptor) HandlerInterceptor {
		return func(ctx context.Context, arg interface{}, nextInterceptor HandlerInterceptor) (result interface{}, err error) {
			return currentInterceptor(ctx, arg, nextInterceptor)
		}
	}

	chainedHandlerInterceptor := interceptors[interceptorCount]

	for i := interceptorCount - 1; i >= 0; i-- {
		chainedHandlerInterceptor = chainer(chainedHandlerInterceptor, interceptors[i])
	}

	return chainedHandlerInterceptor
}
