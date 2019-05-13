// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/customer_negative_criterion_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v1/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import status "google.golang.org/genproto/googleapis/rpc/status"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for
// [CustomerNegativeCriterionService.GetCustomerNegativeCriterion][google.ads.googleads.v1.services.CustomerNegativeCriterionService.GetCustomerNegativeCriterion].
type GetCustomerNegativeCriterionRequest struct {
	// The resource name of the criterion to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerNegativeCriterionRequest) Reset()         { *m = GetCustomerNegativeCriterionRequest{} }
func (m *GetCustomerNegativeCriterionRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerNegativeCriterionRequest) ProtoMessage()    {}
func (*GetCustomerNegativeCriterionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_negative_criterion_service_bef57624ec14ad68, []int{0}
}
func (m *GetCustomerNegativeCriterionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerNegativeCriterionRequest.Unmarshal(m, b)
}
func (m *GetCustomerNegativeCriterionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerNegativeCriterionRequest.Marshal(b, m, deterministic)
}
func (dst *GetCustomerNegativeCriterionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerNegativeCriterionRequest.Merge(dst, src)
}
func (m *GetCustomerNegativeCriterionRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerNegativeCriterionRequest.Size(m)
}
func (m *GetCustomerNegativeCriterionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerNegativeCriterionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerNegativeCriterionRequest proto.InternalMessageInfo

func (m *GetCustomerNegativeCriterionRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [CustomerNegativeCriterionService.MutateCustomerNegativeCriteria][google.ads.googleads.v1.services.CustomerNegativeCriterionService.MutateCustomerNegativeCriteria].
type MutateCustomerNegativeCriteriaRequest struct {
	// The ID of the customer whose criteria are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual criteria.
	Operations []*CustomerNegativeCriterionOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCustomerNegativeCriteriaRequest) Reset()         { *m = MutateCustomerNegativeCriteriaRequest{} }
func (m *MutateCustomerNegativeCriteriaRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerNegativeCriteriaRequest) ProtoMessage()    {}
func (*MutateCustomerNegativeCriteriaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_negative_criterion_service_bef57624ec14ad68, []int{1}
}
func (m *MutateCustomerNegativeCriteriaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaRequest.Unmarshal(m, b)
}
func (m *MutateCustomerNegativeCriteriaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaRequest.Marshal(b, m, deterministic)
}
func (dst *MutateCustomerNegativeCriteriaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerNegativeCriteriaRequest.Merge(dst, src)
}
func (m *MutateCustomerNegativeCriteriaRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaRequest.Size(m)
}
func (m *MutateCustomerNegativeCriteriaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerNegativeCriteriaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerNegativeCriteriaRequest proto.InternalMessageInfo

func (m *MutateCustomerNegativeCriteriaRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCustomerNegativeCriteriaRequest) GetOperations() []*CustomerNegativeCriterionOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCustomerNegativeCriteriaRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCustomerNegativeCriteriaRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create or remove) on a customer level negative criterion.
type CustomerNegativeCriterionOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CustomerNegativeCriterionOperation_Create
	//	*CustomerNegativeCriterionOperation_Remove
	Operation            isCustomerNegativeCriterionOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *CustomerNegativeCriterionOperation) Reset()         { *m = CustomerNegativeCriterionOperation{} }
func (m *CustomerNegativeCriterionOperation) String() string { return proto.CompactTextString(m) }
func (*CustomerNegativeCriterionOperation) ProtoMessage()    {}
func (*CustomerNegativeCriterionOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_negative_criterion_service_bef57624ec14ad68, []int{2}
}
func (m *CustomerNegativeCriterionOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerNegativeCriterionOperation.Unmarshal(m, b)
}
func (m *CustomerNegativeCriterionOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerNegativeCriterionOperation.Marshal(b, m, deterministic)
}
func (dst *CustomerNegativeCriterionOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerNegativeCriterionOperation.Merge(dst, src)
}
func (m *CustomerNegativeCriterionOperation) XXX_Size() int {
	return xxx_messageInfo_CustomerNegativeCriterionOperation.Size(m)
}
func (m *CustomerNegativeCriterionOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerNegativeCriterionOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerNegativeCriterionOperation proto.InternalMessageInfo

type isCustomerNegativeCriterionOperation_Operation interface {
	isCustomerNegativeCriterionOperation_Operation()
}

type CustomerNegativeCriterionOperation_Create struct {
	Create *resources.CustomerNegativeCriterion `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CustomerNegativeCriterionOperation_Remove struct {
	Remove string `protobuf:"bytes,2,opt,name=remove,proto3,oneof"`
}

func (*CustomerNegativeCriterionOperation_Create) isCustomerNegativeCriterionOperation_Operation() {}

func (*CustomerNegativeCriterionOperation_Remove) isCustomerNegativeCriterionOperation_Operation() {}

func (m *CustomerNegativeCriterionOperation) GetOperation() isCustomerNegativeCriterionOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CustomerNegativeCriterionOperation) GetCreate() *resources.CustomerNegativeCriterion {
	if x, ok := m.GetOperation().(*CustomerNegativeCriterionOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CustomerNegativeCriterionOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CustomerNegativeCriterionOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CustomerNegativeCriterionOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CustomerNegativeCriterionOperation_OneofMarshaler, _CustomerNegativeCriterionOperation_OneofUnmarshaler, _CustomerNegativeCriterionOperation_OneofSizer, []interface{}{
		(*CustomerNegativeCriterionOperation_Create)(nil),
		(*CustomerNegativeCriterionOperation_Remove)(nil),
	}
}

func _CustomerNegativeCriterionOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CustomerNegativeCriterionOperation)
	// operation
	switch x := m.Operation.(type) {
	case *CustomerNegativeCriterionOperation_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *CustomerNegativeCriterionOperation_Remove:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("CustomerNegativeCriterionOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _CustomerNegativeCriterionOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CustomerNegativeCriterionOperation)
	switch tag {
	case 1: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.CustomerNegativeCriterion)
		err := b.DecodeMessage(msg)
		m.Operation = &CustomerNegativeCriterionOperation_Create{msg}
		return true, err
	case 2: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &CustomerNegativeCriterionOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _CustomerNegativeCriterionOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CustomerNegativeCriterionOperation)
	// operation
	switch x := m.Operation.(type) {
	case *CustomerNegativeCriterionOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CustomerNegativeCriterionOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for customer negative criterion mutate.
type MutateCustomerNegativeCriteriaResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCustomerNegativeCriteriaResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *MutateCustomerNegativeCriteriaResponse) Reset() {
	*m = MutateCustomerNegativeCriteriaResponse{}
}
func (m *MutateCustomerNegativeCriteriaResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerNegativeCriteriaResponse) ProtoMessage()    {}
func (*MutateCustomerNegativeCriteriaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_negative_criterion_service_bef57624ec14ad68, []int{3}
}
func (m *MutateCustomerNegativeCriteriaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaResponse.Unmarshal(m, b)
}
func (m *MutateCustomerNegativeCriteriaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaResponse.Marshal(b, m, deterministic)
}
func (dst *MutateCustomerNegativeCriteriaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerNegativeCriteriaResponse.Merge(dst, src)
}
func (m *MutateCustomerNegativeCriteriaResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaResponse.Size(m)
}
func (m *MutateCustomerNegativeCriteriaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerNegativeCriteriaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerNegativeCriteriaResponse proto.InternalMessageInfo

func (m *MutateCustomerNegativeCriteriaResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCustomerNegativeCriteriaResponse) GetResults() []*MutateCustomerNegativeCriteriaResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the criterion mutate.
type MutateCustomerNegativeCriteriaResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCustomerNegativeCriteriaResult) Reset()         { *m = MutateCustomerNegativeCriteriaResult{} }
func (m *MutateCustomerNegativeCriteriaResult) String() string { return proto.CompactTextString(m) }
func (*MutateCustomerNegativeCriteriaResult) ProtoMessage()    {}
func (*MutateCustomerNegativeCriteriaResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_customer_negative_criterion_service_bef57624ec14ad68, []int{4}
}
func (m *MutateCustomerNegativeCriteriaResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaResult.Unmarshal(m, b)
}
func (m *MutateCustomerNegativeCriteriaResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaResult.Marshal(b, m, deterministic)
}
func (dst *MutateCustomerNegativeCriteriaResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCustomerNegativeCriteriaResult.Merge(dst, src)
}
func (m *MutateCustomerNegativeCriteriaResult) XXX_Size() int {
	return xxx_messageInfo_MutateCustomerNegativeCriteriaResult.Size(m)
}
func (m *MutateCustomerNegativeCriteriaResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCustomerNegativeCriteriaResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCustomerNegativeCriteriaResult proto.InternalMessageInfo

func (m *MutateCustomerNegativeCriteriaResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCustomerNegativeCriterionRequest)(nil), "google.ads.googleads.v1.services.GetCustomerNegativeCriterionRequest")
	proto.RegisterType((*MutateCustomerNegativeCriteriaRequest)(nil), "google.ads.googleads.v1.services.MutateCustomerNegativeCriteriaRequest")
	proto.RegisterType((*CustomerNegativeCriterionOperation)(nil), "google.ads.googleads.v1.services.CustomerNegativeCriterionOperation")
	proto.RegisterType((*MutateCustomerNegativeCriteriaResponse)(nil), "google.ads.googleads.v1.services.MutateCustomerNegativeCriteriaResponse")
	proto.RegisterType((*MutateCustomerNegativeCriteriaResult)(nil), "google.ads.googleads.v1.services.MutateCustomerNegativeCriteriaResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerNegativeCriterionServiceClient is the client API for CustomerNegativeCriterionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerNegativeCriterionServiceClient interface {
	// Returns the requested criterion in full detail.
	GetCustomerNegativeCriterion(ctx context.Context, in *GetCustomerNegativeCriterionRequest, opts ...grpc.CallOption) (*resources.CustomerNegativeCriterion, error)
	// Creates or removes criteria. Operation statuses are returned.
	MutateCustomerNegativeCriteria(ctx context.Context, in *MutateCustomerNegativeCriteriaRequest, opts ...grpc.CallOption) (*MutateCustomerNegativeCriteriaResponse, error)
}

type customerNegativeCriterionServiceClient struct {
	cc *grpc.ClientConn
}

func NewCustomerNegativeCriterionServiceClient(cc *grpc.ClientConn) CustomerNegativeCriterionServiceClient {
	return &customerNegativeCriterionServiceClient{cc}
}

func (c *customerNegativeCriterionServiceClient) GetCustomerNegativeCriterion(ctx context.Context, in *GetCustomerNegativeCriterionRequest, opts ...grpc.CallOption) (*resources.CustomerNegativeCriterion, error) {
	out := new(resources.CustomerNegativeCriterion)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CustomerNegativeCriterionService/GetCustomerNegativeCriterion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerNegativeCriterionServiceClient) MutateCustomerNegativeCriteria(ctx context.Context, in *MutateCustomerNegativeCriteriaRequest, opts ...grpc.CallOption) (*MutateCustomerNegativeCriteriaResponse, error) {
	out := new(MutateCustomerNegativeCriteriaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CustomerNegativeCriterionService/MutateCustomerNegativeCriteria", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerNegativeCriterionServiceServer is the server API for CustomerNegativeCriterionService service.
type CustomerNegativeCriterionServiceServer interface {
	// Returns the requested criterion in full detail.
	GetCustomerNegativeCriterion(context.Context, *GetCustomerNegativeCriterionRequest) (*resources.CustomerNegativeCriterion, error)
	// Creates or removes criteria. Operation statuses are returned.
	MutateCustomerNegativeCriteria(context.Context, *MutateCustomerNegativeCriteriaRequest) (*MutateCustomerNegativeCriteriaResponse, error)
}

func RegisterCustomerNegativeCriterionServiceServer(s *grpc.Server, srv CustomerNegativeCriterionServiceServer) {
	s.RegisterService(&_CustomerNegativeCriterionService_serviceDesc, srv)
}

func _CustomerNegativeCriterionService_GetCustomerNegativeCriterion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerNegativeCriterionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerNegativeCriterionServiceServer).GetCustomerNegativeCriterion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CustomerNegativeCriterionService/GetCustomerNegativeCriterion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerNegativeCriterionServiceServer).GetCustomerNegativeCriterion(ctx, req.(*GetCustomerNegativeCriterionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerNegativeCriterionService_MutateCustomerNegativeCriteria_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerNegativeCriteriaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerNegativeCriterionServiceServer).MutateCustomerNegativeCriteria(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CustomerNegativeCriterionService/MutateCustomerNegativeCriteria",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerNegativeCriterionServiceServer).MutateCustomerNegativeCriteria(ctx, req.(*MutateCustomerNegativeCriteriaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerNegativeCriterionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CustomerNegativeCriterionService",
	HandlerType: (*CustomerNegativeCriterionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerNegativeCriterion",
			Handler:    _CustomerNegativeCriterionService_GetCustomerNegativeCriterion_Handler,
		},
		{
			MethodName: "MutateCustomerNegativeCriteria",
			Handler:    _CustomerNegativeCriterionService_MutateCustomerNegativeCriteria_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/customer_negative_criterion_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/customer_negative_criterion_service.proto", fileDescriptor_customer_negative_criterion_service_bef57624ec14ad68)
}

var fileDescriptor_customer_negative_criterion_service_bef57624ec14ad68 = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0x37, 0x69, 0xa9, 0x76, 0xb6, 0x2a, 0x8c, 0x88, 0x4b, 0x29, 0x75, 0x49, 0x5b, 0x2d, 0x7b,
	0x48, 0xd8, 0xf5, 0x16, 0xdb, 0xc3, 0x6e, 0x6d, 0x77, 0xad, 0xd8, 0x96, 0x14, 0xf6, 0x20, 0x0b,
	0x71, 0x9a, 0xbc, 0x86, 0x40, 0x36, 0x13, 0x67, 0x26, 0x91, 0x52, 0x7a, 0xf1, 0xdc, 0x9b, 0xdf,
	0x40, 0x6f, 0x7e, 0x14, 0x8f, 0xfa, 0x15, 0x04, 0xc1, 0xcf, 0xa0, 0x20, 0xf9, 0x33, 0xb1, 0x2d,
	0xa4, 0x59, 0xa8, 0xb7, 0x97, 0xf7, 0x5e, 0x7e, 0xef, 0xfd, 0x7e, 0xf3, 0xe6, 0x0d, 0xda, 0xf5,
	0x28, 0xf5, 0x02, 0x30, 0x88, 0xcb, 0x8d, 0xdc, 0x4c, 0xad, 0xa4, 0x63, 0x70, 0x60, 0x89, 0xef,
	0x00, 0x37, 0x9c, 0x98, 0x0b, 0x3a, 0x01, 0x66, 0x87, 0xe0, 0x11, 0xe1, 0x27, 0x60, 0x3b, 0xcc,
	0x17, 0xc0, 0x7c, 0x1a, 0xda, 0x45, 0x92, 0x1e, 0x31, 0x2a, 0x28, 0x6e, 0xe5, 0x00, 0x3a, 0x71,
	0xb9, 0x5e, 0x62, 0xe9, 0x49, 0x47, 0x97, 0x58, 0x8b, 0x5b, 0x55, 0xd5, 0x18, 0x70, 0x1a, 0xb3,
	0x9a, 0x72, 0x79, 0x99, 0xc5, 0x25, 0x09, 0x12, 0xf9, 0x06, 0x09, 0x43, 0x2a, 0x88, 0xf0, 0x69,
	0xc8, 0x8b, 0xe8, 0x72, 0x11, 0xcd, 0xbe, 0x8e, 0xe2, 0x63, 0xe3, 0x3d, 0x23, 0x51, 0x04, 0x4c,
	0xc6, 0x1f, 0x15, 0x71, 0x16, 0x39, 0x06, 0x17, 0x44, 0xc4, 0x45, 0x40, 0xdb, 0x45, 0x2b, 0x03,
	0x10, 0x5b, 0x45, 0xf9, 0xbd, 0xa2, 0xfa, 0x96, 0x2c, 0x6e, 0xc1, 0xbb, 0x18, 0xb8, 0xc0, 0x2b,
	0xe8, 0xae, 0x6c, 0xd6, 0x0e, 0xc9, 0x04, 0x9a, 0x4a, 0x4b, 0x59, 0x9f, 0xb7, 0x16, 0xa4, 0x73,
	0x8f, 0x4c, 0x40, 0xfb, 0xad, 0xa0, 0xb5, 0xd7, 0xb1, 0x20, 0x02, 0x2a, 0xf0, 0x88, 0x84, 0x7b,
	0x8c, 0x1a, 0x25, 0x63, 0xdf, 0x2d, 0xc0, 0x90, 0x74, 0xbd, 0x74, 0xb1, 0x8b, 0x10, 0x8d, 0x80,
	0xe5, 0x1c, 0x9b, 0x6a, 0x6b, 0x66, 0xbd, 0xd1, 0x7d, 0xa1, 0xd7, 0x29, 0xad, 0x57, 0xf2, 0xd8,
	0x97, 0x60, 0xd6, 0x05, 0x5c, 0xfc, 0x14, 0xdd, 0x8f, 0x08, 0x13, 0x3e, 0x09, 0xec, 0x63, 0xe2,
	0x07, 0x31, 0x83, 0xe6, 0x4c, 0x4b, 0x59, 0xbf, 0x63, 0xdd, 0x2b, 0xdc, 0x3b, 0xb9, 0x37, 0xa5,
	0x9f, 0x90, 0xc0, 0x77, 0x89, 0x00, 0x9b, 0x86, 0xc1, 0x49, 0x73, 0x36, 0x4b, 0x5b, 0x90, 0xce,
	0xfd, 0x30, 0x38, 0xd1, 0x3e, 0x2b, 0x48, 0xab, 0x6f, 0x00, 0x8f, 0xd0, 0x9c, 0xc3, 0x80, 0x88,
	0x5c, 0xc3, 0x46, 0x77, 0xa3, 0x92, 0x56, 0x39, 0x1e, 0xd5, 0xbc, 0x86, 0xb7, 0xac, 0x02, 0x0d,
	0x37, 0xd1, 0x1c, 0x83, 0x09, 0x4d, 0xa0, 0xa9, 0xa6, 0x72, 0xa6, 0x91, 0xfc, 0xbb, 0xdf, 0x40,
	0xf3, 0x25, 0x69, 0xed, 0x9b, 0x82, 0x9e, 0xd4, 0x1d, 0x12, 0x8f, 0x68, 0xc8, 0x01, 0xef, 0xa0,
	0x87, 0x57, 0xe4, 0xb1, 0x81, 0x31, 0xca, 0x32, 0x91, 0x1a, 0x5d, 0x2c, 0x1b, 0x67, 0x91, 0xa3,
	0x1f, 0x66, 0x43, 0x65, 0x3d, 0xb8, 0x2c, 0xdc, 0x76, 0x9a, 0x8e, 0xdf, 0xa2, 0xdb, 0x0c, 0x78,
	0x1c, 0x08, 0x79, 0x92, 0x3b, 0xf5, 0x27, 0x59, 0xdb, 0x62, 0x1c, 0x08, 0x4b, 0xc2, 0x6a, 0xaf,
	0xd0, 0xea, 0x34, 0x3f, 0x4c, 0x35, 0xc6, 0xdd, 0xf3, 0x59, 0xd4, 0xaa, 0x14, 0xfc, 0x30, 0x6f,
	0x10, 0xff, 0x54, 0xd0, 0xd2, 0x75, 0x17, 0x07, 0x6f, 0xd7, 0x73, 0x9c, 0xe2, 0xe2, 0x2d, 0xde,
	0x68, 0x3a, 0xb4, 0xfe, 0x87, 0xef, 0x3f, 0x3e, 0xaa, 0x1b, 0xd8, 0x4c, 0xb7, 0xcd, 0xe9, 0x25,
	0xea, 0x9b, 0xf2, 0xb6, 0x71, 0xa3, 0x5d, 0xae, 0x9f, 0xab, 0xb2, 0x19, 0xed, 0x33, 0xfc, 0x47,
	0x41, 0xcb, 0xd7, 0x8b, 0x8b, 0x07, 0x37, 0x3f, 0xcf, 0x9c, 0xed, 0xf0, 0x3f, 0x0c, 0x46, 0x36,
	0xbb, 0xda, 0x30, 0x63, 0xde, 0xd7, 0x36, 0x53, 0xe6, 0xff, 0xa8, 0x9e, 0x5e, 0x58, 0x3b, 0x9b,
	0xed, 0xb3, 0x4a, 0xe2, 0xe6, 0x24, 0x2b, 0x63, 0x2a, 0xed, 0xfe, 0xb9, 0x8a, 0x56, 0x1d, 0x3a,
	0xa9, 0xed, 0xac, 0xbf, 0x56, 0x37, 0x34, 0x07, 0xe9, 0xc6, 0x3d, 0x50, 0xde, 0x0c, 0x0b, 0x28,
	0x8f, 0x06, 0x24, 0xf4, 0x74, 0xca, 0x3c, 0xc3, 0x83, 0x30, 0xdb, 0xc7, 0xf2, 0x7d, 0x88, 0x7c,
	0x5e, 0xfd, 0x38, 0x3d, 0x97, 0xc6, 0x27, 0x75, 0x66, 0xd0, 0xeb, 0x7d, 0x51, 0x5b, 0x83, 0x1c,
	0xb0, 0xe7, 0x72, 0x3d, 0x37, 0x53, 0x6b, 0xd4, 0xd1, 0x8b, 0xc2, 0xfc, 0xab, 0x4c, 0x19, 0xf7,
	0x5c, 0x3e, 0x2e, 0x53, 0xc6, 0xa3, 0xce, 0x58, 0xa6, 0xfc, 0x52, 0x57, 0x73, 0xbf, 0x69, 0xf6,
	0x5c, 0x6e, 0x9a, 0x65, 0x92, 0x69, 0x8e, 0x3a, 0xa6, 0x29, 0xd3, 0x8e, 0xe6, 0xb2, 0x3e, 0x9f,
	0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x87, 0xdb, 0xa3, 0xde, 0x43, 0x07, 0x00, 0x00,
}
