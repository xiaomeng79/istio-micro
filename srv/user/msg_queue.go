package user

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/xiaomeng79/go-log"
	"github.com/xiaomeng79/istio-micro/cinit"
)

func msgNotify(ctx context.Context, msg_str string) {
	topic := cinit.TOPIC_SRV_KEY_CHANGE
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg_str),
		//Partition:0,
	}
	part, offset, err := cinit.Kf.SyncProducer(msg)
	if err != nil {
		log.Error("msg notify error:"+err.Error(), ctx)
	}
	log.Infof("topic:%s,part:%d,offset:%d", topic, part, offset, ctx)
}
