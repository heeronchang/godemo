package sentinel_test

import (
	"log"
	"math/rand"
	"net/http"
	"sync/atomic"
	"testing"
	"time"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/util"
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

func TestSentinel(tt *testing.T) {
	go server()

	// 资源 (resource) 是 Sentinel 中的最核心概念之一，Sentinel 中所有的限流熔断机制都是基于资源生效的，不同资源的限流熔断规则互相隔离互不影响。
	// 在 Sentinel 中，用户可以灵活的定义资源埋点。资源可以是应用、接口、函数、甚至是一段代码。我们的流量治理机制都是为了保护这段资源运行如预期一样。

	// 用户通过 Sentinel api 包里面的接口可以把资源访问包起来，这一步称为“埋点”。每个埋点都有一个资源名称（resource），代表触发了这个资源的调用或访问。
	// 有了资源埋点之后，我们就可以针对资源埋点配置流量治理规则。即使没有配置任何规则，资源埋点仍然会产生 metric 统计。

	err := sentinel.InitDefault()
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
	}

	// 配置一条限流规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "test-ping",
			Threshold:              10,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
		},
	})

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("fuck")
	ch := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			for {
				// 埋点逻辑，埋点资源名为 test-ping
				e, b := sentinel.Entry("test-ping")
				if b != nil {
					log.Println("will sleep x millsecond.")
					// 请求被拒绝，在此处进行处理
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
				} else {
					// 请求允许通过，此处编写业务逻辑
					resp, err := resty.New().R().Get("http://localhost:8080/ping")
					if err != nil {
						log.Printf("request err:%s\n", err.Error())
					}
					log.Printf("resp:%v", resp)
					log.Println(util.CurrentTimeMillis(), "Passed")
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)

					// 务必保证业务结束后调用 Exit
					e.Exit()
				}
			}
		}()
		log.Println(i)
	}

	<-ch
}
