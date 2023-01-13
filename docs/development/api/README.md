这里是GISB的API文档，此文档为前后端开发人员提供关于GISB的 `HTTP` 服务的说明与解释

# 1. 状态码定义

GISB的后端使用HTTP状态码来表示请求的响应结果，大部分情况下，非 `200(Success)` 状态码都不会返回除状态码的任何信息

下面是GISB的HTTP服务对其返回状态码的说明：

**200(Success)**

此状态表明HTTP请求合法，且服务端响应过程正常，且返回数据正常

**400(BadRequest)**

此状态表明HTTP请求非法，您可以试着分析您的请求，是否将必须的参数提供完整提供，且请求体的格式正确无误

**403(Forbidden)**

此状态表明HTTP请求合法，但发起者对所请求的资源不具备访问权限，您可以试着分析您的请求，是否将鉴权相关的参数/cookie完整提供，或是请求的资源是否正确

**404(NotFound)**

此状态表明HTTP请求合法，但服务端不持有所请求资源，您可以检查您访问的资源是否正确

**412(PreconditionFailed)**

此状态表明HTTP请求合法，但发起者被服务器所限制，一般是由于您的请求频率过快或认证失败次数过多导致的，您可以等待一段时间后重新发起请求

**500(InternalError)**

此状态表明HTTP请求合法，但服务器在响应请求的过程中出现了错误，您可以将此错误反馈给服务器管理人员或本项目的开发者以寻求解决方法

# 2. 配置说明

GISB可以在启动前对其进行配置，此处对其进行说明与解释，配置的声明如下：

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

默认的配置如下：

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
