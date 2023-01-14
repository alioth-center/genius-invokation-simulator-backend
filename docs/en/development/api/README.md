Here is the API documentation for GISB. This document provides fore-end and back-end developers with instructions and explanations about GISB's HTTP service

# 1. Status Code definition

The GISB back-end uses HTTP status codes to represent the response results of requests, and in most cases, anything other than a '200(Success)' status code will not be returned

Here's what the GISB HTTP service says about its return status code:

**200(Success)**

This status indicates that the HTTP request is valid and that the server is responding properly and returning normal data

**400(BadRequest)**

This status indicates that the HTTP request is invalid, and you can try to analyze your request to see if the required parameters are supplied in full and the request body is properly formatted

**403(Forbidden)**

This state indicates that the HTTP request is valid, but the originator does not have access to the requested resource. You can try to analyze your request to see if all auth parameters/cookies are provided, or if the requested resource is correct

**404(NotFound)**

This state indicates that the HTTP request is valid, but the server does not hold the requested resource, and you can check that you are accessing the correct resource

**412(PreconditionFailed)**

This state indicates that the HTTP request is valid, but the initiator is limited by the server, usually due to your request rate is too fast or the number of authentication failures is too high, you can wait for a period of time and then make a new request

**500(InternalError)**

This status indicates that the HTTP request is valid, but there was an error in the server's response to the request. You can report this error to the server administrator or the project developer for resolution

# 2. Configuration instructions

GISB can be configured before starting, which is described and explained here. The declaration of the configuration is as follows:

```go
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
```

The default configuration is as follows：

```go
var (
    config = &EngineConfig{
            Middleware: MiddlewareConfig{
                UUIDKey:                 "uuid",
                IPTranceKey:             "ip",
                InterdictorTraceKey:     "interdicted",
                InterdictorBlockedTime:  600,
                InterdictorTriggerCount: 5,
                QPSLimitTime:            1,
                TokenIDKey:              "gisb_token_id",
                TokenKey:                "gisb_token",
                TokenRefreshTime:        7200,
                TokenRemainingTime:      1800,
                CookieDomain:            "localhost",
            },
            Service: ServiceConfig{
                MaxRooms: 100,
            },
        }
    }
)
```
