package card

import (
	"context"
)

type Server interface {
	Register(ctx context.Context, pluginInfo *PluginInfo, attr string, token string) (err error)
	GetPlugins(ctx context.Context) (plugins []*PluginInfo, err error)
}
