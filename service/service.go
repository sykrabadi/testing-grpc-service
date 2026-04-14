package service

import "testing-grpc-service/scripts/grpc/student"

// Service is enhanced version of ServiceV1, 
// with injected student gRPC client interface
type Service struct {
	studentGRPCClient student.StudentClient
}

func NewService(
	studentGRPCClient student.StudentClient,
) Service {
	return Service{
		studentGRPCClient: studentGRPCClient,
	}
}

type ServiceV1 struct{}

func NewServiceV1() ServiceV1 {
	return ServiceV1{}
}
