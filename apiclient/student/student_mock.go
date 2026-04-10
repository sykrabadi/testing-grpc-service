package student

import (
	"context"

	studentGRPCClient "testing-grpc-service/scripts/grpc/student"

	"google.golang.org/grpc"
)

type StudentGRPCMock struct{}

func NewStudentGRPCMock() *StudentGRPCMock{
	return &StudentGRPCMock{}
}

func (m *StudentGRPCMock) GetStudent(ctx context.Context, in *studentGRPCClient.GetStudentRequest, opts ...grpc.CallOption) (*studentGRPCClient.GetStudentResponse, error) {
	return &studentGRPCClient.GetStudentResponse{
		Id: 1,
	}, nil
}
