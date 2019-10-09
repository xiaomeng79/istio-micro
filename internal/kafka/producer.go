package kafka

import (
	"context"
	"sync"

	"github.com/Shopify/sarama"
)

type successFunc func(msg *sarama.ProducerMessage)
type errorFunc func(msg *sarama.ProducerError)

// 同步消息模式
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
			k.ap.AsyncClose() //  Trigger a shutdown of the producer.
			break ProducerLoop
		}
	}

	wg.Wait()
}
