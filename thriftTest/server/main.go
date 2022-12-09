package main

import (
	"GoWild/thriftTest/gen-go/simple"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strconv"
	"time"
)

type SimpleServiceHandler struct {
}

func (ssh *SimpleServiceHandler) Add(ctx context.Context, num1 int32, num2 string) (_r int32, _err error) {
	fmt.Println("调用Add的实现")
	num2int32, _ := strconv.Atoi(num2)
	return num1 + int32(num2int32), nil
}

// server.go
func main() {
	conf := &thrift.TConfiguration{
		ConnectTimeout:     time.Second,
		SocketTimeout:      time.Second,
		MaxFrameSize:       1024 * 256,
		TBinaryStrictRead:  thrift.BoolPtr(true),
		TBinaryStrictWrite: thrift.BoolPtr(true),
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(conf)
	transportFactory := thrift.NewTTransportFactory()

	transport, _ := thrift.NewTServerSocket(":8090")

	processor := simple.NewSimpleServiceProcessor(&SimpleServiceHandler{})
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	server.Serve()
}
