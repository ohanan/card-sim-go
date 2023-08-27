package card

import "context"

func (c *_Client) Register(ctx context.Context, pluginInfo *PluginInfo, attr string, token string) (err error) {
	req := &struct {
		Base       *BaseRequest
		PluginInfo *PluginInfo
		Attr       string
		Token      string
	}{
		Base:       newBaseRequestFromContext(ctx),
		PluginInfo: pluginInfo,
		Attr:       attr,
		Token:      token,
	}
	resp := &struct {
		Base *BaseResponse
	}{}
	err = c.c.Call("CS.Register", req, resp)
	return err
}
func (c *_Client) GetPlugins(ctx context.Context) (plugins []*PluginInfo, err error) {
	req := &struct {
		Base *BaseRequest
	}{
		Base: newBaseRequestFromContext(ctx),
	}
	resp := &struct {
		Base    *BaseResponse
		Plugins []*PluginInfo
	}{}
	err = c.c.Call("CS.GetPlugins", req, resp)
	return resp.Plugins, err
}
