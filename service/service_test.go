package service_test

import (
	"testing-grpc-service/apiclient/student"
	"testing-grpc-service/service"
)

type serviceProxy struct {
	svc service.Service
}

func newServiceProxy(
) serviceProxy{
	mockStudentClient := student.NewStudentGRPCMock()
	svc := service.NewService(mockStudentClient)
 
	return serviceProxy{
		svc: svc,
	}
}
