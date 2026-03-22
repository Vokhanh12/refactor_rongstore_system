package logger

type LogContext struct {
	ServiceName string
	TraceID     string
	UserID      string
	ClientID    string
	RealmID     string
	SpanID      string
}

type LogEntry struct {
	Context      LogContext
	Code         string
	Key          string
	Message      string
	Cause        string
	CauseDetail  string
	ClientAction string
	ServerAction string
	Expected     bool
	HTTPStatus   int
	GRPCCode     string
}

type AccessLog struct {
	Context   LogContext
	Path      string
	Method    string
	HTTPCode  int
	IP        string
	UserAgent string
	LatencyMS int64
	Extra     map[string]interface{}
}
