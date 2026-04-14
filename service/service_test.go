package service_test

import (
	"testing-grpc-service/apiclient/student"
	"testing-grpc-service/service"
)

type serviceProxy struct {
	svc service.Service
	svcV1 service.ServiceV1
}

func newServiceProxy(
) serviceProxy{
	mockStudentClient := student.NewStudentGRPCMock()
	svc := service.NewService(mockStudentClient)
 
	svcV1 := service.NewServiceV1()
	return serviceProxy{
		svc: svc,
		svcV1: svcV1,
	}
}
