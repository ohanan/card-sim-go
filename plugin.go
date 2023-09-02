package card

import "context"

type Plugin interface {
	GetPluginInfo(ctx context.Context) (info *PluginInfo, err error)
	Init(ctx context.Context, params map[string]any) (err error)
	Start(ctx context.Context) (err error)
	Close(ctx context.Context) (err error)
}

type PluginInfo struct {
	Name       string
	InitConfig *InitConfig
}
type InitConfig struct {
}

type Card struct {
}
