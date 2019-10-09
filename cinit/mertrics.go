package cinit

import (
	"time"

	"github.com/xiaomeng79/istio-micro/internal/metrics"
)

func metricsInit(sn string) {
	if Config.Metrics.Enable == "yes" {
		//  Push模式
		m := metrics.NewMetrics()
		//  e.Use(api.MetricsFunc(m))
		m.MemStats()
		//  InfluxDB
		m.InfluxDBWithTags(
			time.Duration(Config.Metrics.Duration)*time.Second,
			Config.Metrics.URL,
			Config.Metrics.Database,
			Config.Metrics.UserName,
			Config.Metrics.Password,
			map[string]string{"service": sn},
		)

		//  Graphite
		// addr, _ := net.ResolveTCPAddr("tcp", Conf.Metrics.Address)
		// m.Graphite(Conf.Metrics.FreqSec*time.Second, "echo-web.node."+hostname, addr)
	}
}
