package service

import (
	"context"
	"testing-grpc-service/scripts/grpc/student"

	"google.golang.org/grpc"
)

// GetStudent is enhanced version of GetStudentV1. Instead of 
// creating gRPC client on service level, it uses interface to 
// call gRPC services, it makes easier to test because test will 
// no longer depends on actual gRPC connection
func (s *Service) GetStudent(
	ctx context.Context,
	id int32,
) (int32, error) {
	responseID, err := s.studentGRPCClient.GetStudent(ctx, &student.GetStudentRequest{
		Id: id,
	})
	if err != nil {
		return 0, err
	}

	return responseID.Id, nil
}

// GetStudentV1 create gRPC client on service level. This creates 
// tight coupling and makes it harder to test
func (s *ServiceV1) GetStudentV1(
	ctx context.Context,
	studentClientAddr string,
	id int32,
) (int32, error) {
	client, err := grpc.NewClient(studentClientAddr)
	if err != nil {
		return 0, err
	}

	studentGRPCClient := student.NewStudentClient(client)

	responseID, err := studentGRPCClient.GetStudent(ctx, &student.GetStudentRequest{
		Id: id,
	})
	if err != nil {
		return 0, err
	}

	return responseID.Id, nil
}
