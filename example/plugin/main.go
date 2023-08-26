package main

import "github.com/ohanan/card-sim-go"

type demoPlugin struct {
}

func (d demoPlugin) GetPluginInfo() *card.PluginInfo {
	return &card.PluginInfo{
		Name: "demo",
	}
}

func (d demoPlugin) Ping(ctx card.Context) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (d demoPlugin) Close(ctx card.Context) error {
	// TODO implement me
	panic("implement me")
}

func (d demoPlugin) OnStart(ctx card.Context) error {
	// TODO implement me
	panic("implement me")
}

func main() {
	card.Run(&demoPlugin{}, card.WithServerPort(53072))
}
