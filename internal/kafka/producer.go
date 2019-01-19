package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"sync"
)

//func main()  {
//	addrs := []string{"127.0.0.1:9092"}
//	kafka := NewKafka(addrs)
//
//	topic := "test3"
//	value := "sync: this is a message. index=0"
//	msg := &sarama.ProducerMessage{
//		Topic:topic,
//		Value:sarama.ByteEncoder(value),
//		//Partition:0,
//	}
//	part,offset,err := kafka.syncProducer(msg)
//	fmt.Println(part)
//	fmt.Println(offset)
//
//	if err != nil {
//		log.Printf("send message(%s) err=%s \n", value, err)
//	}else {
//		fmt.Printf(value + "发送成功，partition=%d, offset=%d \n", part, offset)
//	}
//
//
//
//	msgc := make(chan *sarama.ProducerMessage,1)
//	ctx,cal := context.WithCancel(context.Background())
//	go func(ctx context.Context) {
//		for i:=0; i<=10;i++  {
//			msgc <- msg
//		}
//		cal()
//	}(ctx)
//	kafka.asyncProducer(ctx,msgc, func(msg *sarama.ProducerMessage) {
//
//	}, func(msg *sarama.ProducerError) {
//
//	})
//	//asyncProducer1(Address)
//}

type successFunc func(msg *sarama.ProducerMessage)
type errorFunc func(msg *sarama.ProducerError)

//同步消息模式
func (k *Kafka) SyncProducer(msg *sarama.ProducerMessage) (part int32, offset int64, err error) {
	part, offset, err = k.sp.SendMessage(msg)
	return
}

func (k *Kafka) AsyncProducer(ctx context.Context, msg chan *sarama.ProducerMessage, sf successFunc, ef errorFunc) {

	var (
		wg sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range k.ap.Successes() {
			sf(msg)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range k.ap.Errors() {
			ef(err)
		}
	}()

ProducerLoop:
	for {
		select {
		case message := <-msg:
			k.ap.Input() <- message

		case <-ctx.Done():
			k.ap.AsyncClose() // Trigger a shutdown of the producer.
			break ProducerLoop
		}
	}

	wg.Wait()
}
