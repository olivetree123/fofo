package common

import "fofo/entity"

const CommandHealthCheckResponse int = 0 // 健康检查返回值
const CommandRegisterService int = 100   // 注册服务
const CommandGetService int = 101        // 获取服务
const CommandSearchService int = 102     // 搜索服务

const RequestQueue string = "FOFO:REQUEST" // 监听的请求队列
const HealthCheckChannel string = "FOFO:HEALTHCHECK"

// Services 已注册的服务
var Services map[string]*entity.Service

func init() {
	Services = make(map[string]*entity.Service)
}
