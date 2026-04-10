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
		resp, err := svcProxy.svc.GetStudent(cctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, resp)
	})
}
