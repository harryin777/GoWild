package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
)

func main() {

	consulReg := consul.NewRegistry(registry.Addrs(":8500"))
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Registry(consulReg),
	)

	service.Init()
	c := service.Client()

	request := c.NewRequest("greeter", "Greeter.Hello", "john", client.WithContentType("application/json"))
	var response string
	if err := c.Call(context.TODO(), request, &response); err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)
}
