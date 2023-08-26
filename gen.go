package card

type _Server struct {
	s      Server
	system *serverSystem
}
type _ServerGetPluginsRequest struct {
}

func (s *_Server) GetPlugins(_ int, resp *[]PluginInfo) (err error) {
	*resp, err = s.s.GetPlugins(nil)
	return
}

type _Plugin struct {
	p Plugin
}

func (p *_Plugin) Ping(_ struct{}, resp *string) (err error) {
	*resp, err = p.p.Ping(nil)
	return
}
func (p *_Plugin) Close(_ struct{}, _ *struct{}) (err error) {
	err = p.p.Close(nil)
	return
}
