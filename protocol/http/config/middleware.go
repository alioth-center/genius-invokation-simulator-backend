package config

type MiddlewareConfig struct {
	UUIDKey                 string `json:"uuid_key" yaml:"uuid_key" xml:"uuid_key"`
	IPTranceKey             string `json:"ip_trace_key" yaml:"ip_trance_key" xml:"ip_trance_key"`
	InterdictorTraceKey     string `json:"interdictor_trace_key" yaml:"interdictor_trace_key" xml:"interdictor_trace_key"`
	InterdictorBlockedTime  uint   `json:"interdictor_blocked_time" yaml:"interdictor_blocked_time" xml:"interdictor_blocked_time"`
	InterdictorTriggerCount uint   `json:"interdictor_trigger_count" yaml:"interdictor_trigger_count" xml:"interdictor_trigger_count"`
	QPSLimitTime            uint   `json:"qps_limit_time" yaml:"qps_limit_time" xml:"qps_limit_time"`
	TokenIDKey              string `json:"token_id_key" yaml:"token_id_key" xml:"token_id_key"`
	TokenKey                string `json:"token_key" yaml:"token_key" xml:"token_key"`
	TokenRefreshTime        uint   `json:"token_refresh_time" yaml:"token_refresh_time" xml:"token_refresh_time"`
	TokenRemainingTime      uint   `json:"token_remaining_time" yaml:"token_remaining_time" xml:"token_remaining_time"`
	CookieDomain            string `json:"cookie_domain" yaml:"cookie_domain" xml:"cookie_domain"`
}
