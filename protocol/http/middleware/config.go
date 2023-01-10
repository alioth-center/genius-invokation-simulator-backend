package middleware

type Config struct {
	UUIDKey                 string `json:"uuid_key"`
	IPTranceKey             string `json:"ip_trace_key"`
	InterdictorTraceKey     string `json:"interdictor_trace_key"`
	InterdictorBlockedTime  uint   `json:"interdictor_blocked_time"`
	InterdictorTriggerCount uint   `json:"interdictor_trigger_count"`
	QPSLimitTime            uint   `json:"qps_limit_time"`
}
