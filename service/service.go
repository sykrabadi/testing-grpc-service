package service

import "testing-grpc-service/scripts/grpc/student"

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
