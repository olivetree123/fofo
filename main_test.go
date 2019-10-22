package main

import (
	"fofo/common"
	"fofo/entity"
	"fofo/handlers"
	"fofo/utils"
	"testing"
)

func testRegisterService(t *testing.T) {
	service := entity.Service{
		ID:      "myService",
		Name:    "myService",
		Group:   "myGroup",
		Port:    6666,
		Address: "localhost",
	}
	serviceMap := utils.Struct2Map(service)
	cmd := entity.Command{
		Code:      common.CommandRegisterService,
		RequestID: "TestRegisterService",
		Args:      serviceMap,
	}
	r, err := handlers.RegisterServiceHandler(&cmd)
	if err != nil {
		t.Error(err)
	}
	t.Log("register service success, service = ", r)
}

func testGetService(t *testing.T) {
	paramMap := make(map[string]interface{})
	paramMap["service_id"] = "myService"
	cmd := entity.Command{
		Code:      common.CommandGetService,
		RequestID: "TestGetService",
		Args:      paramMap,
	}
	r, err := handlers.GetServiceHandler(&cmd)
	if err != nil {
		t.Error(err)
	}
	t.Log("get service success, service = ", r)
}

func TestAll(t *testing.T) {
	t.Run("TestRegisterService", testRegisterService)
	t.Run("TestGetService", testGetService)
}
