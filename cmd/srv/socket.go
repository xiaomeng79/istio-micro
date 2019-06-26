//+build srv_socket

package main

import (
	"github.com/xiaomeng79/istio-micro/cmd/cmd"
	"github.com/xiaomeng79/istio-micro/srv/socket"
	_ "github.com/xiaomeng79/istio-micro/version"
)

func main() {
	cmd.Execute()
	socket.Run()
}
