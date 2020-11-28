package main

import (
	"math/rand"
	"os"
	"time"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/core/isolation"
	"github.com/alibaba/sentinel-golang/logging"
)

func main() {
	cfg := config.NewDefaultConfig()
	// for testing, logging output to console
	cfg.Sentinel.Log.Logger = logging.NewConsoleLogger()
	err := sentinel.InitWithConfig(cfg)
	if err != nil {
		logging.Error(err, "fail")
		os.Exit(1)
	}
	logging.ResetGlobalLoggerLevel(logging.DebugLevel)
	ch := make(chan struct{})

	r1 := &isolation.Rule{
		Resource:   "abc",
		MetricType: isolation.Concurrency,
		Threshold:  12,
	}
	_, err = isolation.LoadRules([]*isolation.Rule{r1})
	if err != nil {
		logging.Error(err, "fail")
		os.Exit(1)
	}

	for i := 0; i < 15; i++ {
		go func(batchNum int) {
			for {
				e, b := sentinel.Entry("abc", sentinel.WithBatchCount(1))
				if b != nil {
					logging.Warn("[Isolation] Blocked", "batch", batchNum, "reason", b.BlockType().String(), "rule", b.TriggeredRule(), "snapshot", b.TriggeredValue())
					time.Sleep(time.Duration(rand.Uint64()%5+1) * time.Second)
				} else {
					logging.Info("[Isolation] Passed", "batch", batchNum)
					time.Sleep(time.Duration(rand.Uint64()%5+1) * time.Second)
					e.Exit()
				}
			}
		}(i)
	}
	<-ch
}
