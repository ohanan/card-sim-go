package main

import (
	"fmt"

	"github.com/ohanan/card-sim-go"
)

type DemoServer struct {
}

func (d DemoServer) OnPluginLoaded(ctx card.ServerContext, pluginInfo *card.PluginInfo) error {
	println("Plugin loaded")
	return nil
}

func (d DemoServer) GetPlugins(ctx card.ServerContext) ([]card.PluginInfo, error) {
	fmt.Println("hhh")
	return nil, nil
}

func main() {
	c := make(chan struct{})
	_ = card.RunServer(&DemoServer{}, c, card.WithServerPort(53072))
}
