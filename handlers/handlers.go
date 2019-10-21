package handlers

import (
	. "fofo/common"
	"fofo/entity"
	"fofo/storage"
	"fofo/utils"
	"github.com/micro/go-micro/errors"
	"github.com/mitchellh/mapstructure"
	"time"
)

// CommandHandler 处理命令
func CommandHandler(cmd *entity.Command, responseChannel chan *entity.Response) error {
	var err error
	var r interface{}
	if cmd.Code == CommandRegisterService {
		//Logger.Info("注册服务")
		r, err = RegisterServiceHandler(cmd)
	} else if cmd.Code == CommandHealthCheckResponse {
		//Logger.Info("健康检查返回")
		HealthCheckResponseHandler(cmd)
		return nil
	} else if cmd.Code == CommandGetService {
		//Logger.Info("获取服务")
		r, err = GetServiceHandler(cmd)
	} else if cmd.Code == CommandSearchService {
		//Logger.Info("搜索服务")
		r, err = SearchServiceHandler(cmd)
	}
	if err != nil {
		return err
	}
	var resp entity.Response
	resp.RequestID = cmd.RequestID
	resp.Content = r
	responseChannel <- &resp
	return nil
}

// registerServiceHandler 注册服务处理器
func RegisterServiceHandler(cmd *entity.Command) (*entity.Service, error) {
	var service entity.Service
	err := mapstructure.Decode(cmd.Args, &service)
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	if service.ID == "" {
		service.ID = utils.NewUUID()
	}
	service.HealthCheckTime = time.Now().Unix()
	err = storage.AddService(&service)
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	return &service, nil
}

// GetServiceParam 获取服务详情的参数
type GetServiceParam struct {
	ServiceID string `json:"service_id" mapstructure:"service_id"`
}

// SearchServiceParam 搜索服务的参数
type SearchServiceParam struct {
	Name  string
	Group string
}

// GetServiceHandler 获取服务详情
func GetServiceHandler(cmd *entity.Command) (*entity.Service, error) {
	var param GetServiceParam
	err := mapstructure.Decode(cmd.Args, &param)
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	Logger.Info("service = ", param)
	if param.ServiceID == "" {
		Logger.Error("ServiceID should not be null")
		return nil, errors.BadRequest("123", "ServiceID should not be null")
	}
	service, err := storage.GetService(param.ServiceID)
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	return service, nil
}

// SearchServiceHandler 搜索服务
func SearchServiceHandler(cmd *entity.Command) ([]*entity.Service, error) {
	var param SearchServiceParam
	err := mapstructure.Decode(cmd.Args, &param)
	if err != nil {
		Logger.Error(err)
		return nil, err
	}
	return storage.SearchService(param.Name, param.Group)
}

// HealthCheckResponseHandler 健康检查处理器
func HealthCheckResponseHandler(cmd *entity.Command) {
	var resp entity.HealthCheckResponse
	err := mapstructure.Decode(cmd.Args, &resp)
	if err != nil {
		Logger.Error(err)
		return
	}
	if resp.ServiceID == "" {
		Logger.Error("ServiceID should not be null")
	}
	_, err = storage.UpdateServiceTime(resp.ServiceID, resp.HealthCheckTime)
	if err != nil {
		Logger.Error(err)
		return
	}
	return
}
