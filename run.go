package card

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type runOptions struct {
	port  int
	token string
}
type RunOption func(options *runOptions)

func WithServerPort(port int) RunOption {
	return func(options *runOptions) {
		options.port = port
	}
}
func WithToken(token string) RunOption {
	return func(options *runOptions) {
		options.token = token
	}
}

type RegisterPluginReq struct {
	Info *PluginInfo
	Addr string
}
type RegisterPluginResp struct {
	Error string
}

func getOptions(options []RunOption) runOptions {
	var ro runOptions
	for _, option := range options {
		option(&ro)
	}
	port := flag.Int("port", 0, "server port")
	flag.Parse()
	if port != nil && *port > 0 {
		ro.port = *port
	}
	return ro
}
func Run(plugin Plugin, options ...RunOption) error {
	ro := getOptions(options)
	c, err := jsonrpc.Dial("tcp", fmt.Sprintf("localhost:%d", ro.port))
	if err != nil {
		return err
	}
	listen, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	s := &_Client{c: c}
	info, err := plugin.GetPluginInfo(context.Background())
	if err != nil {
		return err
	}
	err = s.Register(context.Background(), info, listen.Addr().String(), "")
	if err != nil {
		return err
	}
	server := rpc.NewServer()
	err = server.RegisterName("CP", &_Server{c: plugin})
	if err != nil {
		return err
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			return err
		}
		server.ServeCodec(jsonrpc.NewServerCodec(accept))
	}
}

type BaseRequest struct {
}
type BaseResponse struct {
}

func newContextFromBaseRequest(request *BaseRequest) context.Context {
	return context.Background()
}
func newBaseRequestFromContext(ctx context.Context) *BaseRequest {
	return &BaseRequest{}
}

type _Client struct {
	c *rpc.Client
}
type _Server struct {
	c Plugin
}
