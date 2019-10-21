package entity

import "time"

// 命令
type Command struct {
	Code      int                    `json:"code"`
	RequestID string                 `json:"request_id"`
	Args      map[string]interface{} `json:"args"`
}

type Response struct {
	RequestID string
	Content   interface{}
}

// Service 服务
type Service struct {
	ID              string                 `json:"id"`
	Name            string                 `json:"name"`
	Group           string                 `json:"group"`
	Port            int                    `json:"port"`
	Address         string                 `json:"address"`
	Extra           map[string]interface{} `json:"extra"`
	HealthCheckTime int64                  `json:"-"`
}

// HealthCheckResponse 健康检查返回值
type HealthCheckResponse struct {
	ServiceID       string `json:"service_id" mapstructure:"service_id"`
	HealthCheckTime int64  `json:"healthCheckTime" mapstructure:"healthCheckTime"`
}

func (service *Service) IsValid() bool {
	// 一次健康检查未返回，则定义为不可用的服务。服务搜索时不返回。
	now := time.Now().Unix()
	if service.HealthCheckTime+6 >= now {
		return true
	}
	return false
}
