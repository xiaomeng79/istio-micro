//+build api_frontend

package main

import (
	"github.com/xiaomeng79/istio-micro/api/frontend"
	"github.com/xiaomeng79/istio-micro/cmd/cmd"
	_ "github.com/xiaomeng79/istio-micro/version"
)

func main() {
	cmd.Execute()
	frontend.Run()
}
