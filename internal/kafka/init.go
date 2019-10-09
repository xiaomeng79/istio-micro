package kafka

import (
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/xiaomeng79/go-log"
)

type Kafka struct {
	addrs []string
	c     sarama.Client
	sp    sarama.SyncProducer
	ap    sarama.AsyncProducer
	sc    *cluster.Client
	ss    sarama.Consumer
}

func NewKafka(addrs []string) *Kafka {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	config.Producer.Return.Errors = true

	c, err := sarama.NewClient(addrs, config)
	if err != nil {
		log.Fatal(err.Error())
	}
	cconfig := cluster.NewConfig()
	cconfig.Consumer.Return.Errors = true
	cconfig.Group.Return.Notifications = true
	// cconfig.Group.Mode = cluster.ConsumerModePartitions
	sc, err := cluster.NewClient(addrs, cconfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	sp, err := sarama.NewSyncProducerFromClient(c)
	if err != nil {
		log.Fatalf("sarama.NewSyncProducer err, message=%s \n", err)
	}
	ap, err := sarama.NewAsyncProducerFromClient(c)
	if err != nil {
		log.Fatalf("sarama.NewAsyncProducer err, message=%s \n", err)
	}
	ss, err := sarama.NewConsumerFromClient(c)
	if err != nil {
		log.Fatalf("sarama.NewConsumer err, message=%s \n", err)
	}
	return &Kafka{addrs: addrs, c: c, sc: sc, sp: sp, ap: ap, ss: ss}
}

func (k *Kafka) Close() {
	err := k.c.Close()
	if err != nil {
		log.Error("kafka close:", err.Error())
	}
}
