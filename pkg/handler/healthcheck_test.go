package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_getHealth(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Successful health check",
			args: args{},
			want: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &healthCheckHandler{}
			gin.SetMode(gin.TestMode)
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			h.getHealth(ctx)
			if ctx.Writer.Status() != tt.want {
				t.Errorf("getHealth() = %v, want %v", tt.args.c.Request.Response.StatusCode, tt.want)
			}
		})
	}
}
