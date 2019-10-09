package socket

import (
	"context"
	"net/http"

	"github.com/xiaomeng79/istio-micro/cinit"
	_ "github.com/xiaomeng79/istio-micro/srv/socket/statik" // 打包静态文件

	"github.com/Shopify/sarama"
	socketio "github.com/googollee/go-socket.io"
	"github.com/rakyll/statik/fs"
	"github.com/xiaomeng79/go-log"
)

const (
	SN = "srv-socket" // 定义services名称
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	// 初始化
	cinit.InitOption(SN, cinit.Trace, cinit.Kafka)
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	go broadcast(ctx, server)
	_ = server.On("connection", func(so socketio.Socket) {
		_ = so.On("init", func(msg string) {
			_ = so.Join("match")
			_ = so.Emit("init", msg)
		})
		_ = so.On("send", func(msg string) {
			_ = so.Emit("recive", "your msg is:"+msg)
		})
		_ = so.On("ack", func(msg string) string {
			return "ack msg:" + msg
		})
		_ = so.On("disconnection", func() {
		})
	})
	_ = server.On("error", func(so socketio.Socket, err error) {
		log.Error(err.Error())
	})

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	http.Handle("/socket.io/", server)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))
	if err := http.ListenAndServe(cinit.Config.SrvSocket.Port, nil); err != nil {
		log.Fatal("failed to listen: " + err.Error())
	}
	defer cancel()
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("socket main:%+v", r)
		}
	}()
}

func broadcast(ctx context.Context, server *socketio.Server) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("socket broadcast:%+v", r)
		}
	}()
	// 接受通知
	// topics := []string{cinit.TOPIC_SRV_KEY_CHANGE}
	// cinit.Kf.ConsumerGroup(ctx, GID, topics, func(message *sarama.ConsumerMessage) {
	// 	log.Debugf("msg:%+s", string(message.Value))
	// 	sign := new(utils.SocketSign)
	// 	sign.K = string(message.Value)
	// 	server.BroadcastTo("match", sign.K, sign.String())
	// })
	cinit.Kf.Consumer(ctx, cinit.TopicSrvKeyChange, 0, -1, func(message *sarama.ConsumerMessage) {
		log.Debugf("msg:%+s", string(message.Value))
		server.BroadcastTo("match", "broadcast", string(message.Value))
	}, func(consumerError *sarama.ConsumerError) {
		log.Errorf("topic:%s,part:%d,error:%+v", consumerError.Topic, consumerError.Partition, consumerError.Err)
	})
}
