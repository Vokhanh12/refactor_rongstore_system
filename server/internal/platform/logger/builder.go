package logger

import "go.uber.org/zap"

func buildFields(entry LogEntry, extra map[string]interface{}) []zap.Field {
	fields := []zap.Field{
		zap.String("service", entry.Context.ServiceName),
		zap.String("code", entry.Code),
		zap.String("key", entry.Key),
		zap.String("message", entry.Message),
		zap.String("cause", entry.Cause),
		zap.String("client_action", entry.ClientAction),
		zap.String("server_action", entry.ServerAction),
		zap.Int("http_status", entry.HTTPStatus),
		zap.String("grpc_code", entry.GRPCCode),
		zap.Bool("expected", entry.Expected),
	}

	if entry.CauseDetail != "" {
		fields = append(fields, zap.String("cause_detail", entry.CauseDetail))
	}

	if entry.Context.TraceID != "" {
		fields = append(fields,
			zap.String("trace_id", entry.Context.TraceID),
			zap.String("user_id", entry.Context.UserID),
			zap.String("client_id", entry.Context.ClientID),
			zap.String("realm_id", entry.Context.RealmID),
			zap.String("span_id", entry.Context.SpanID),
		)
	}

	for k, v := range extra {
		fields = append(fields, zap.Any(k, v))
	}

	return fields
}

func buildAccessFields(access AccessLog) []zap.Field {
	fields := []zap.Field{
		zap.String("path", access.Path),
		zap.String("method", access.Method),
		zap.Int("http_code", access.HTTPCode),
		zap.String("ip", access.IP),
		zap.String("user_agent", access.UserAgent),
		zap.Int64("latency_ms", access.LatencyMS),
	}

	if access.Context.TraceID != "" {
		fields = append(fields,
			zap.String("trace_id", access.Context.TraceID),
			zap.String("user_id", access.Context.UserID),
			zap.String("client_id", access.Context.ClientID),
			zap.String("realm_id", access.Context.RealmID),
			zap.String("span_id", access.Context.SpanID),
		)
	}

	for k, v := range access.Extra {
		fields = append(fields, zap.Any(k, v))
	}

	return fields
}
