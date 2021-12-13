package main

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
)

var consulReg registry.Registry

func init() {
	consulReg = consul.NewRegistry(registry.Addrs(":8500")) // 告知consul的端口号，如果走默认可以不填写
}

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, name *string, msg *string) error {
	*msg = "Hello world !!!" + *name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(consulReg),
	)

	// initialise command line
	service.Init()

	// set the handler
	micro.RegisterHandler(service.Server(), new(Greeter))

	// run service
	service.Run()
}
