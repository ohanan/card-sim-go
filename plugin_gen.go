package card

import (
	"context"
)

var _ Plugin = (*PluginImpl)(nil)

type PluginImpl struct {
}

func (p PluginImpl) GetPluginInfo(ctx context.Context) (info *PluginInfo, err error) {
	return nil, nil
}

func (p PluginImpl) Init(ctx context.Context, params map[string]any) (err error) {
	return nil
}

func (p PluginImpl) Start(ctx context.Context) (err error) {
	return nil
}

func (p PluginImpl) Close(ctx context.Context) (err error) {
	return nil
}
