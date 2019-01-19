package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"github.com/xiaomeng79/go-log"
)

//var group_id = "group_id"
//func main() {
//	addrs := []string{"127.0.0.1:9092"}
//	kafka := NewKafka(addrs)
//	fmt.Printf("consumer_test")
//	topics := []string{"test3", "my_topic"}
//	ctx ,cal := context.WithCancel(context.Background())
//	kafka.single(ctx,group_id,topics, func(msg *sarama.ConsumerMessage) {
//		fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
//	})
//	defer cal()
//
//}

type rmsgFunc func(*sarama.ConsumerMessage)
type rerrFunc func(*sarama.ConsumerError)
type errFunc func(error)

func (k *Kafka) ConsumerGroup(ctx context.Context, gid string, topics []string, rf rmsgFunc) {
	//topics := []string{"test3", "my_topic"}
	consumer, err := cluster.NewConsumerFromClient(k.sc, gid, topics)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Error("kafka consumer Error:" + err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Infof("Rebalanced: %+v\n", ntf)
		}
	}()

	// consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				rf(msg)
				//fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				consumer.MarkOffset(msg, "") // mark message as processed
			}
		case <-ctx.Done():
			return
		}
	}
}

func (k *Kafka) Consumer(ctx context.Context, topic string, partition int32, offset int64, rf rmsgFunc, ef rerrFunc) {

	partitionConsumer, err := k.ss.ConsumePartition(topic, partition, offset)
	if err != nil {
		log.Error("kafka consumer Error:" + err.Error())
	}

	//defer func() {
	//	if err := partitionConsumer.Close(); err != nil {
	//		log.Error("kafka consumer Error:"+err.Error())
	//	}
	//}()

ConsumerLoop:
	for {
		select {
		case err := <-partitionConsumer.Errors():
			ef(err)
		case msg := <-partitionConsumer.Messages():
			rf(msg)
		case <-ctx.Done():
			break ConsumerLoop
		}
	}

}
