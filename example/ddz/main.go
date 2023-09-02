package main

import (
	"context"

	"github.com/ohanan/card-sim-go"
)

type ddzPlugin struct {
}

func (p *ddzPlugin) GetPluginInfo(ctx context.Context) (info *card.PluginInfo, err error) {
	return &card.PluginInfo{Name: "ddz"}, nil
}

func (p *ddzPlugin) Init(ctx context.Context) (err error) {
	return nil
}

func (p *ddzPlugin) Start(ctx context.Context) (err error) {
	return nil
}

func (p *ddzPlugin) Close(ctx context.Context) (err error) {
	return nil
}

func main() {
	err := card.Run(&ddzPlugin{},
		card.WithServerPort(10280),
	)
	if err != nil {
		panic(err)
	}
}
