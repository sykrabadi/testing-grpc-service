package service

import (
	"context"
	"testing-grpc-service/scripts/grpc/student"
)

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
