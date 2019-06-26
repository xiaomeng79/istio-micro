//+build api_backend

package main

import (
	"github.com/xiaomeng79/istio-micro/api/backend"
	"github.com/xiaomeng79/istio-micro/cmd/cmd"
	_ "github.com/xiaomeng79/istio-micro/version"
)

func main() {
	cmd.Execute()
	backend.Run()
}
