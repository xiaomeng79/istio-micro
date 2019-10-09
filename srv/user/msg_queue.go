package user

import (
	"context"

	"github.com/xiaomeng79/istio-micro/cinit"

	"github.com/Shopify/sarama"
	"github.com/xiaomeng79/go-log"
)

func msgNotify(ctx context.Context, msgStr string) {
	topic := cinit.TopicSrvKeyChange
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msgStr),
		// Partition:0,
	}
	part, offset, err := cinit.Kf.SyncProducer(msg)
	if err != nil {
		log.Error("msg notify error:"+err.Error(), ctx)
	}
	log.Infof("topic:%s,part:%d,offset:%d", topic, part, offset, ctx)
}
