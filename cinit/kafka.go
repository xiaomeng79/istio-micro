package cinit

import (
	"strings"

	"github.com/xiaomeng79/istio-micro/internal/kafka"
)

var Kf *kafka.Kafka

// 初始化连接
func KafkaInit() {
	addrs := strings.Split(Config.Kafka.Addrs, ",")
	Kf = kafka.NewKafka(addrs)
}

// 关闭
func KafkaClose() {
	Kf.Close()
}
