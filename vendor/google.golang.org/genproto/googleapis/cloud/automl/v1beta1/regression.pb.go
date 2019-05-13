// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/regression.proto

package automl // import "google.golang.org/genproto/googleapis/cloud/automl/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Metrics for regression problems.
type RegressionEvaluationMetrics struct {
	// Output only. Root Mean Squared Error (RMSE).
	RootMeanSquaredError float32 `protobuf:"fixed32,1,opt,name=root_mean_squared_error,json=rootMeanSquaredError,proto3" json:"root_mean_squared_error,omitempty"`
	// Output only. Mean Absolute Error (MAE).
	MeanAbsoluteError float32 `protobuf:"fixed32,2,opt,name=mean_absolute_error,json=meanAbsoluteError,proto3" json:"mean_absolute_error,omitempty"`
	// Output only. Mean absolute percentage error. Only set if all ground truth
	// values are are positive.
	MeanAbsolutePercentageError float32 `protobuf:"fixed32,3,opt,name=mean_absolute_percentage_error,json=meanAbsolutePercentageError,proto3" json:"mean_absolute_percentage_error,omitempty"`
	// Output only. R squared.
	RSquared             float32  `protobuf:"fixed32,4,opt,name=r_squared,json=rSquared,proto3" json:"r_squared,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegressionEvaluationMetrics) Reset()         { *m = RegressionEvaluationMetrics{} }
func (m *RegressionEvaluationMetrics) String() string { return proto.CompactTextString(m) }
func (*RegressionEvaluationMetrics) ProtoMessage()    {}
func (*RegressionEvaluationMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_regression_ca949e981cd1d01b, []int{0}
}
func (m *RegressionEvaluationMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegressionEvaluationMetrics.Unmarshal(m, b)
}
func (m *RegressionEvaluationMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegressionEvaluationMetrics.Marshal(b, m, deterministic)
}
func (dst *RegressionEvaluationMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegressionEvaluationMetrics.Merge(dst, src)
}
func (m *RegressionEvaluationMetrics) XXX_Size() int {
	return xxx_messageInfo_RegressionEvaluationMetrics.Size(m)
}
func (m *RegressionEvaluationMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_RegressionEvaluationMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_RegressionEvaluationMetrics proto.InternalMessageInfo

func (m *RegressionEvaluationMetrics) GetRootMeanSquaredError() float32 {
	if m != nil {
		return m.RootMeanSquaredError
	}
	return 0
}

func (m *RegressionEvaluationMetrics) GetMeanAbsoluteError() float32 {
	if m != nil {
		return m.MeanAbsoluteError
	}
	return 0
}

func (m *RegressionEvaluationMetrics) GetMeanAbsolutePercentageError() float32 {
	if m != nil {
		return m.MeanAbsolutePercentageError
	}
	return 0
}

func (m *RegressionEvaluationMetrics) GetRSquared() float32 {
	if m != nil {
		return m.RSquared
	}
	return 0
}

func init() {
	proto.RegisterType((*RegressionEvaluationMetrics)(nil), "google.cloud.automl.v1beta1.RegressionEvaluationMetrics")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/regression.proto", fileDescriptor_regression_ca949e981cd1d01b)
}

var fileDescriptor_regression_ca949e981cd1d01b = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0x99, 0xfe, 0x3f, 0xa2, 0xd9, 0x88, 0x55, 0xb0, 0x38, 0xa2, 0xe2, 0xca, 0x85, 0x24,
	0x14, 0x71, 0xe5, 0xaa, 0x2d, 0xc5, 0x55, 0xa1, 0x54, 0x70, 0x21, 0x85, 0x72, 0x3b, 0xbd, 0x84,
	0x81, 0x34, 0x77, 0xbc, 0x49, 0xfa, 0x22, 0xbe, 0x95, 0x0f, 0xe2, 0x73, 0xc8, 0x24, 0x69, 0x55,
	0x10, 0x97, 0xe1, 0x7c, 0xdf, 0xe1, 0x72, 0x22, 0x6e, 0x35, 0x91, 0x36, 0xa8, 0x2a, 0x43, 0x61,
	0xa5, 0x20, 0x78, 0x5a, 0x1b, 0xb5, 0xe9, 0x2f, 0xd1, 0x43, 0x5f, 0x31, 0x6a, 0x46, 0xe7, 0x6a,
	0xb2, 0xb2, 0x61, 0xf2, 0xd4, 0x2d, 0x13, 0x2d, 0x23, 0x2d, 0x13, 0x2d, 0x33, 0x7d, 0x76, 0x9e,
	0xab, 0xa0, 0xa9, 0x15, 0x58, 0x4b, 0x1e, 0x7c, 0x4d, 0xd6, 0x25, 0xf5, 0xfa, 0xa3, 0x10, 0xe5,
	0x6c, 0xd7, 0x37, 0xde, 0x80, 0x09, 0x31, 0x9f, 0xa0, 0xe7, 0xba, 0x72, 0xdd, 0x7b, 0x71, 0xca,
	0x44, 0x7e, 0xb1, 0x46, 0xb0, 0x0b, 0xf7, 0x1a, 0x80, 0x71, 0xb5, 0x40, 0x66, 0xe2, 0x5e, 0x71,
	0x55, 0xdc, 0x74, 0x66, 0x27, 0x6d, 0x3c, 0x41, 0xb0, 0x4f, 0x29, 0x1c, 0xb7, 0x59, 0x57, 0x8a,
	0xe3, 0x68, 0xc0, 0xd2, 0x91, 0x09, 0x1e, 0xb3, 0xd2, 0x89, 0xca, 0x51, 0x1b, 0x0d, 0x72, 0x92,
	0xf8, 0x91, 0xb8, 0xf8, 0xc9, 0x37, 0xc8, 0x15, 0x5a, 0x0f, 0x7a, 0xab, 0xfe, 0x8b, 0x6a, 0xf9,
	0x5d, 0x9d, 0xee, 0x98, 0x54, 0x52, 0x8a, 0x03, 0xde, 0xde, 0xd8, 0xfb, 0x1f, 0xf9, 0x7d, 0xce,
	0x67, 0x0d, 0xdf, 0x0a, 0x71, 0x59, 0xd1, 0x5a, 0xfe, 0x31, 0xd5, 0xf0, 0xf0, 0x6b, 0x89, 0x69,
	0xbb, 0xce, 0xcb, 0x20, 0xd3, 0x9a, 0x0c, 0x58, 0x2d, 0x89, 0xb5, 0xd2, 0x68, 0xe3, 0x72, 0x2a,
	0x45, 0xd0, 0xd4, 0xee, 0xd7, 0x5f, 0x7a, 0x48, 0xcf, 0xf7, 0x4e, 0xf9, 0x18, 0xc1, 0xf9, 0xa8,
	0x85, 0xe6, 0x83, 0xe0, 0x69, 0x62, 0xe6, 0xcf, 0x09, 0x5a, 0xee, 0xc5, 0xae, 0xbb, 0xcf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa6, 0x7b, 0x3d, 0x93, 0xf0, 0x01, 0x00, 0x00,
}
