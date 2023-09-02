package card

func (s *_Server) GetPluginInfo(req *struct {
	Base *BaseRequest
}, resp *struct {
	Base *BaseResponse
	Info *PluginInfo
}) error {
	info, err := s.c.GetPluginInfo(newContextFromBaseRequest(req.Base))
	resp.Base = &BaseResponse{}
	resp.Info = info
	return err
}
func (s *_Server) Init(req *struct {
	Base *BaseRequest
}, resp *struct {
	Base *BaseResponse
}) error {
	err := s.c.Init(newContextFromBaseRequest(req.Base), nil)
	resp.Base = &BaseResponse{}
	return err
}
func (s *_Server) Start(req *struct {
	Base *BaseRequest
}, resp *struct {
	Base *BaseResponse
}) error {
	err := s.c.Start(newContextFromBaseRequest(req.Base))
	resp.Base = &BaseResponse{}
	return err
}
func (s *_Server) Close(req *struct {
	Base *BaseRequest
}, resp *struct {
	Base *BaseResponse
}) error {
	err := s.c.Close(newContextFromBaseRequest(req.Base))
	resp.Base = &BaseResponse{}
	return err
}
