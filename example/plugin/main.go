package main

import (
	"context"
	"github.com/ohanan/card-sim-go"
)

type demoPlugin struct {
}

func (d demoPlugin) GetPluginInfo(ctx context.Context) (info *card.PluginInfo, err error) {
	return &card.PluginInfo{
		Name: "demo",
	}, nil
}

func (d demoPlugin) Init(ctx context.Context) (err error) {
	println("init")
	return nil
}

func (d demoPlugin) Start(ctx context.Context) (err error) {
	println("start")
	return nil
}

func (d demoPlugin) Close(ctx context.Context) (err error) {
	println("close")
	return nil
}

func main() {
	err := card.Run(&demoPlugin{},
		card.WithServerPort(10280),
	)
	if err != nil {
		panic(err)
	}
}
