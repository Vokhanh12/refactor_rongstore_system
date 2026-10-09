package grpc

import (
	"context"
	"time"

	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/logger"
	"google.golang.org/grpc"
)

func LoggingUnaryInterceptor(service string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		resp, err := handler(ctx, req)
		duration := time.Since(start).Milliseconds()

		traceId := ctxutil.TraceIDFromContext(ctx)
		userId := ctxutil.UserIDFromContext(ctx)
		clientId := ctxutil.ClientIDFromContext(ctx)
		realmId := ctxutil.RealmIDFromContext(ctx)
		spanId := ctxutil.SpanIDFromContext(ctx)

		ip := peerIP(ctx)

		msg := simplifyMethod(info.FullMethod)

		logger.LogAccess(ctx, msg, logger.AccessParams{
			ServiceInfo: logger.ServiceInfo{
				Name: service,
			},
			RequestContext: logger.RequestContext{
				TraceId:  traceId,
				UserId:   userId,
				ClientId: clientId,
				RealmId:  realmId,
				SpanId:   spanId,
			},
			Path:      info.FullMethod,
			Method:    "POST",
			LatencyMS: duration,
			IP:        ip,
			HTTPCode:  200,
			UserAgent: "",
		})

		return resp, err
	}
}
