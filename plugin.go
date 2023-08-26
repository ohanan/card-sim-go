package card

type Plugin interface {
	GetPluginInfo() *PluginInfo

	Ping(ctx Context) (string, error)
	Close(ctx Context) error
	OnStart(ctx Context) error
}

type PluginInfo struct {
	Name string
}
