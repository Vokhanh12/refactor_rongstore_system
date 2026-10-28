package interceptor

import (
	"context"

	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func MetaUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		// 1. Khởi tạo request context mặc định (sinh RequestID mới nếu chưa có)
		reqCtx := ctxutil.NewRequestContext()

		// 2. Đọc metadata từ Envoy / Client g gửi xuống
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if ips := md.Get("x-forwarded-for"); len(ips) > 0 {
				reqCtx.ClientIP = ips[0]
			} else if ips := md.Get("x-real-ip"); len(ips) > 0 {
				reqCtx.ClientIP = ips[0]
			}

			if uas := md.Get("user-agent"); len(uas) > 0 {
				reqCtx.UserAgent = uas[0]
			}

			if tids := md.Get("x-tenant-id"); len(tids) > 0 {
				reqCtx.TenantID = tids[0]
			}

			// Nếu Envoy có truyền sẵn x-request-id lên thì tận dụng luôn
			if reqIDs := md.Get("x-request-id"); len(reqIDs) > 0 {
				reqCtx.RequestID = reqIDs[0]
			}
		}

		// 3. Đưa vào context thông qua hàm chuẩn của ctxutil
		newCtx := ctxutil.WithRequest(ctx, reqCtx)

		return handler(newCtx, req)
	}
}
