package service_test

import (
	"context"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestGetStudent(t *testing.T){
	cctx := context.Background()
	svcProxy := newServiceProxy()

	t.Run("success get student", func(t *testing.T) {
		resp, err := svcProxy.svcV1.GetStudentV1(cctx, "",1)

		// will return error, because test is relying to actual gRPC connection, 
		// but no gRPC server is live currently
		assert.NoError(t, err)
		assert.Equal(t, 1, resp)
	})

	t.Run("success get student", func(t *testing.T) {
		resp, err := svcProxy.svc.GetStudent(cctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, resp)
	})
}
