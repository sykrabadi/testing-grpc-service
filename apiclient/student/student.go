package student

import (
	studentGRPCClient "testing-grpc-service/scripts/grpc/student"

	"google.golang.org/grpc"
)

// NewStudentGRPCService creates actual gRPC 
// connection to student gRPC service
func NewStudentGRPCService(
	addr string,
) (studentGRPCClient.StudentClient, error){
	client, err := grpc.NewClient(addr)
	if err != nil {
		return nil, err
	}

	return studentGRPCClient.NewStudentClient(client), nil
}
