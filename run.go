package card

import (
	"flag"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"
	"sync"
)

type runOptions struct {
	port int
}
type RunOption func(options *runOptions)

func WithServerPort(port int) RunOption {
	return func(options *runOptions) {
		options.port = port
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
	listen, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		return err
	}
	rpr := &RegisterPluginResp{}
	err = c.Call("System.Register", &RegisterPluginReq{
		Info: plugin.GetPluginInfo(),
		Addr: listen.Addr().String(),
	}, rpr)
	if err != nil {
		return err
	}
	s := rpc.NewServer()
	err = s.RegisterName("CSPlugin", &_Plugin{p: plugin})
	if err != nil {
		return err
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			return err
		}
		go s.ServeCodec(jsonrpc.NewServerCodec(accept))
	}
}

type serverSystem struct {
	server  Server
	Addr    string
	Plugins sync.Map // map[string]*rpc.Client
}

func (s *serverSystem) getPlugin(name string) *rpc.Client {
	value, ok := s.Plugins.Load(name)
	if !ok {
		return nil
	}
	return value.(*rpc.Client)
}

func (s *serverSystem) Register(req *RegisterPluginReq, resp *RegisterPluginResp) error {
	info := req.Info
	_, ok := s.Plugins.Load(info.Name)
	if ok {
		resp.Error = fmt.Sprintf("failed to register plugin, has existed one: %v", info.Name)
		return nil
	}
	client, err := jsonrpc.Dial("tcp", req.Addr)
	if err != nil {
		resp.Error = fmt.Sprintf("failed to dial plugin, error: %v", err)
		return nil
	}
	_, loaded := s.Plugins.LoadOrStore(info.Name, client)
	if loaded {
		resp.Error = fmt.Sprintf("failed to register plugin, has existed one: %v", info.Name)
		return nil
	}
	if err := s.server.OnPluginLoaded(nil, info); err != nil {
		return err
	}
	return nil
}

func RunServer(server Server, onClose <-chan struct{}, options ...RunOption) error {
	ro := getOptions(options)
	s := rpc.NewServer()
	ss := &serverSystem{server: server}
	err := s.RegisterName("System", ss)
	if err != nil {
		return err
	}
	err = s.RegisterName("CS", &_Server{
		s:      server,
		system: ss,
	})
	if err != nil {
		return err
	}
	listen, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(ro.port))
	if err != nil {
		return err
	}
	ss.Addr = listen.Addr().String()
	go func() {
		<-onClose
		_ = listen.Close()
	}()
	for {
		accept, err := listen.Accept()
		if err != nil {
			return err
		}
		go s.ServeCodec(jsonrpc.NewServerCodec(accept))
	}
}
