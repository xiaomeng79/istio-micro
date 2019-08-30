//+build srv_user

package main

import (
	"github.com/xiaomeng79/istio-micro/cmd/cmd"
	"github.com/xiaomeng79/istio-micro/srv/user"
	_ "github.com/xiaomeng79/istio-micro/version"
)

func main() {
	cmd.Execute()
	user.Run()
}
