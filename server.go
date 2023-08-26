package card

type Server interface {
	OnPluginLoaded(ctx ServerContext, pluginInfo *PluginInfo) error
	GetPlugins(ctx ServerContext) ([]PluginInfo, error)
}
