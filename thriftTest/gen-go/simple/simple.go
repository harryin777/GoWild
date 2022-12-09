// Code generated by Thrift Compiler (0.17.0). DO NOT EDIT.

package simple

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"time"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = errors.New
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

type SimpleService interface {
	// Parameters:
	//  - Num1
	//  - Num2
	Add(ctx context.Context, num1 int32, num2 string) (_r int32, _err error)
}

type SimpleServiceClient struct {
	c    thrift.TClient
	meta thrift.ResponseMeta
}

func NewSimpleServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *SimpleServiceClient {
	return &SimpleServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewSimpleServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *SimpleServiceClient {
	return &SimpleServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewSimpleServiceClient(c thrift.TClient) *SimpleServiceClient {
	return &SimpleServiceClient{
		c: c,
	}
}

func (p *SimpleServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *SimpleServiceClient) LastResponseMeta_() thrift.ResponseMeta {
	return p.meta
}

func (p *SimpleServiceClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
	p.meta = meta
}

// Parameters:
//   - Num1
//   - Num2
func (p *SimpleServiceClient) Add(ctx context.Context, num1 int32, num2 string) (_r int32, _err error) {
	var _args0 SimpleServiceAddArgs
	_args0.Num1 = num1
	_args0.Num2 = num2
	var _result2 SimpleServiceAddResult
	var _meta1 thrift.ResponseMeta
	_meta1, _err = p.Client_().Call(ctx, "add", &_args0, &_result2)
	p.SetLastResponseMeta_(_meta1)
	if _err != nil {
		return
	}
	return _result2.GetSuccess(), nil
}

type SimpleServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      SimpleService
}

func (p *SimpleServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *SimpleServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *SimpleServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewSimpleServiceProcessor(handler SimpleService) *SimpleServiceProcessor {

	self3 := &SimpleServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self3.processorMap["add"] = &simpleServiceProcessorAdd{handler: handler}
	return self3
}

func (p *SimpleServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
	if err2 != nil {
		return false, thrift.WrapTException(err2)
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(ctx, thrift.STRUCT)
	iprot.ReadMessageEnd(ctx)
	x4 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
	x4.Write(ctx, oprot)
	oprot.WriteMessageEnd(ctx)
	oprot.Flush(ctx)
	return false, x4

}

type simpleServiceProcessorAdd struct {
	handler SimpleService
}

func (p *simpleServiceProcessorAdd) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	var _write_err5 error
	args := SimpleServiceAddArgs{}
	if err2 := args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "add", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}
	iprot.ReadMessageEnd(ctx)

	tickerCancel := func() {}
	// Start a goroutine to do server side connectivity check.
	if thrift.ServerConnectivityCheckInterval > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()
		var tickerCtx context.Context
		tickerCtx, tickerCancel = context.WithCancel(context.Background())
		defer tickerCancel()
		go func(ctx context.Context, cancel context.CancelFunc) {
			ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if !iprot.Transport().IsOpen() {
						cancel()
						return
					}
				}
			}
		}(tickerCtx, cancel)
	}

	result := SimpleServiceAddResult{}
	if retval, err2 := p.handler.Add(ctx, args.Num1, args.Num2); err2 != nil {
		tickerCancel()
		err = thrift.WrapTException(err2)
		if errors.Is(err2, thrift.ErrAbandonRequest) {
			return false, thrift.WrapTException(err2)
		}
		_exc6 := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing add: "+err2.Error())
		if err2 := oprot.WriteMessageBegin(ctx, "add", thrift.EXCEPTION, seqId); err2 != nil {
			_write_err5 = thrift.WrapTException(err2)
		}
		if err2 := _exc6.Write(ctx, oprot); _write_err5 == nil && err2 != nil {
			_write_err5 = thrift.WrapTException(err2)
		}
		if err2 := oprot.WriteMessageEnd(ctx); _write_err5 == nil && err2 != nil {
			_write_err5 = thrift.WrapTException(err2)
		}
		if err2 := oprot.Flush(ctx); _write_err5 == nil && err2 != nil {
			_write_err5 = thrift.WrapTException(err2)
		}
		if _write_err5 != nil {
			return false, thrift.WrapTException(_write_err5)
		}
		return true, err
	} else {
		result.Success = &retval
	}
	tickerCancel()
	if err2 := oprot.WriteMessageBegin(ctx, "add", thrift.REPLY, seqId); err2 != nil {
		_write_err5 = thrift.WrapTException(err2)
	}
	if err2 := result.Write(ctx, oprot); _write_err5 == nil && err2 != nil {
		_write_err5 = thrift.WrapTException(err2)
	}
	if err2 := oprot.WriteMessageEnd(ctx); _write_err5 == nil && err2 != nil {
		_write_err5 = thrift.WrapTException(err2)
	}
	if err2 := oprot.Flush(ctx); _write_err5 == nil && err2 != nil {
		_write_err5 = thrift.WrapTException(err2)
	}
	if _write_err5 != nil {
		return false, thrift.WrapTException(_write_err5)
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//   - Num1
//   - Num2
type SimpleServiceAddArgs struct {
	Num1 int32  `thrift:"num1,1" db:"num1" json:"num1"`
	Num2 string `thrift:"num2,2" db:"num2" json:"num2"`
}

func NewSimpleServiceAddArgs() *SimpleServiceAddArgs {
	return &SimpleServiceAddArgs{}
}

func (p *SimpleServiceAddArgs) GetNum1() int32 {
	return p.Num1
}

func (p *SimpleServiceAddArgs) GetNum2() string {
	return p.Num2
}
func (p *SimpleServiceAddArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *SimpleServiceAddArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Num1 = v
	}
	return nil
}

func (p *SimpleServiceAddArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Num2 = v
	}
	return nil
}

func (p *SimpleServiceAddArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "add_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SimpleServiceAddArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "num1", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:num1: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.Num1)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.num1 (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:num1: ", p), err)
	}
	return err
}

func (p *SimpleServiceAddArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "num2", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:num2: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Num2)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.num2 (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:num2: ", p), err)
	}
	return err
}

func (p *SimpleServiceAddArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SimpleServiceAddArgs(%+v)", *p)
}

// Attributes:
//   - Success
type SimpleServiceAddResult struct {
	Success *int32 `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewSimpleServiceAddResult() *SimpleServiceAddResult {
	return &SimpleServiceAddResult{}
}

var SimpleServiceAddResult_Success_DEFAULT int32

func (p *SimpleServiceAddResult) GetSuccess() int32 {
	if !p.IsSetSuccess() {
		return SimpleServiceAddResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *SimpleServiceAddResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SimpleServiceAddResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *SimpleServiceAddResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *SimpleServiceAddResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "add_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SimpleServiceAddResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.I32, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteI32(ctx, int32(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *SimpleServiceAddResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SimpleServiceAddResult(%+v)", *p)
}
