//+build srv_account

package main

import (
	"github.com/xiaomeng79/istio-micro/cmd/cmd"
	"github.com/xiaomeng79/istio-micro/srv/account"
	_ "github.com/xiaomeng79/istio-micro/version"
)

func main() {
	cmd.Execute()
	account.Run()
}
