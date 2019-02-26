package pprof

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const (
	Addr    = ":38888"
	PidFile = "server.pid"
)

func Run() {
	ch := make(chan os.Signal, 1)
	ch1 := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGUSR1)  //开启pprof
	signal.Notify(ch1, syscall.SIGUSR2) //关闭pprof

	//将pid写入文件
	f, err := os.Create(PidFile)
	if err != nil {
		log.Printf("%v", err)
	}
	f.WriteString(strconv.Itoa(os.Getpid()))
	f.Close()

	log.Print("进程id: ", os.Getpid())
	var server *http.Server
	for {
		select {
		case <-ch:
			go func() {
				server = &http.Server{
					Addr: Addr,
				}
				log.Print("listen addr: ", Addr)

				err := server.ListenAndServe()
				if err != nil {
					log.Printf("listen:%v", err)
				}
			}()
		case <-ch1:
			if server != nil {
				err := server.Close()
				if err != nil {
					log.Printf("server close:%v", err)
				} else {
					server = nil
				}
			}

		}
	}
}
