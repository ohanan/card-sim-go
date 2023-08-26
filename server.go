package card

type Server interface {
	OnPluginRegister(ctx ServerContext, pluginInfo *PluginInfo) error
	GetPlugins(ctx ServerContext) ([]PluginInfo, error)
}
