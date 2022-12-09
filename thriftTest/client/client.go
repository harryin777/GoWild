package main

import (
	"GoWild/thriftTest/gen-go/simple"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"time"
)

// client.go

func handleClient(client *simple.SimpleServiceClient) {
	res, _ := client.Add(context.Background(), 13, "25")
	fmt.Println("result is ", res)
}

func main() {
	var transport thrift.TTransport
	transport = thrift.NewTSocketConf("localhost:8090", &thrift.TConfiguration{
		ConnectTimeout: time.Second, // Use 0 for no timeout
		SocketTimeout:  time.Second, // Use 0 for no timeout
	})

	conf := &thrift.TConfiguration{
		ConnectTimeout: time.Second,
		SocketTimeout:  time.Second,

		MaxFrameSize: 1024 * 256,

		TBinaryStrictRead:  thrift.BoolPtr(true),
		TBinaryStrictWrite: thrift.BoolPtr(true),
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(conf)

	transportFactory := thrift.NewTTransportFactory()
	transport, _ = transportFactory.GetTransport(transport)
	defer transport.Close()
	transport.Open()

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	handleClient(simple.NewSimpleServiceClient(thrift.NewTStandardClient(iprot, oprot)))
}
