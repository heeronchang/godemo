package hystrix_test

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"testing"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func server() {
	r := gin.Default()

	var count int64

	r.GET("/ping", func(ctx *gin.Context) {
		if atomic.AddInt64(&count, 1) < 3 {
			ctx.String(http.StatusInternalServerError, "pong")
			return
		}

		ctx.String(http.StatusOK, "pong")
	})

	_ = r.Run(":8080")
}

func TestQuickStart(t *testing.T) {
	go server()

	hystrix.ConfigureCommand("test-api-ping", hystrix.CommandConfig{
		// 执行命令超时时间, 默认值 1 秒
		Timeout: 0,

		// 最大并发请求量, 默认值 10
		MaxConcurrentRequests: 100,

		// 熔断开启前需要达到的最小请求数量, 默认值 20
		RequestVolumeThreshold: 5,

		// 熔断器开启后，重试服务是否恢复的等待时间，默认值 5 秒
		// 这里修改为 0.5 秒
		SleepWindow: 500,

		// 请求错误百分比阈值，超过阈值后熔断开启
		ErrorPercentThreshold: 20,
	})

	for i := 0; i < 20; i++ {
		_ = hystrix.Do("test-api-ping", func() error {
			// resp, _ := resty.New().R().Get("https://api.github.com/")
			resp, _ := resty.New().R().Get("http://localhost:8080/ping")
			if resp.IsError() {
				return fmt.Errorf("err code: %s", resp.Status())
			}
			return nil
		}, func(err error) error {
			fmt.Println("fallback err", err)
			return err
		})

		time.Sleep(100 * time.Millisecond)
	}
}
