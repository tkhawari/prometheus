// Code generated by protoc-gen-go.
// source: remote.proto
// DO NOT EDIT!

/*
Package remote is a generated protocol buffer package.

It is generated from these files:
	remote.proto

It has these top-level messages:
	Sample
	LabelPair
	TimeSeries
	WriteRequest
*/
package remote

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Sample struct {
	Value       float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	TimestampMs int64   `protobuf:"varint,2,opt,name=timestamp_ms,json=timestampMs" json:"timestamp_ms,omitempty"`
}

func (m *Sample) Reset()                    { *m = Sample{} }
func (m *Sample) String() string            { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()               {}
func (*Sample) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LabelPair struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *LabelPair) Reset()                    { *m = LabelPair{} }
func (m *LabelPair) String() string            { return proto.CompactTextString(m) }
func (*LabelPair) ProtoMessage()               {}
func (*LabelPair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type TimeSeries struct {
	Labels []*LabelPair `protobuf:"bytes,1,rep,name=labels" json:"labels,omitempty"`
	// Sorted by time, oldest sample first.
	Samples []*Sample `protobuf:"bytes,2,rep,name=samples" json:"samples,omitempty"`
}

func (m *TimeSeries) Reset()                    { *m = TimeSeries{} }
func (m *TimeSeries) String() string            { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()               {}
func (*TimeSeries) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TimeSeries) GetLabels() []*LabelPair {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TimeSeries) GetSamples() []*Sample {
	if m != nil {
		return m.Samples
	}
	return nil
}

type WriteRequest struct {
	Timeseries []*TimeSeries `protobuf:"bytes,1,rep,name=timeseries" json:"timeseries,omitempty"`
}

func (m *WriteRequest) Reset()                    { *m = WriteRequest{} }
func (m *WriteRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()               {}
func (*WriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WriteRequest) GetTimeseries() []*TimeSeries {
	if m != nil {
		return m.Timeseries
	}
	return nil
}

func init() {
	proto.RegisterType((*Sample)(nil), "remote.Sample")
	proto.RegisterType((*LabelPair)(nil), "remote.LabelPair")
	proto.RegisterType((*TimeSeries)(nil), "remote.TimeSeries")
	proto.RegisterType((*WriteRequest)(nil), "remote.WriteRequest")
}

func init() { proto.RegisterFile("remote.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0x3f, 0x4f, 0x80, 0x30,
	0x10, 0xc5, 0x03, 0x68, 0x0d, 0x07, 0x31, 0xf1, 0xe2, 0xc0, 0xa8, 0x9d, 0x70, 0x61, 0xc0, 0xf8,
	0x01, 0x74, 0xd6, 0xc4, 0x14, 0x13, 0x47, 0x53, 0x92, 0x1b, 0x9a, 0xb4, 0x82, 0x6d, 0xf1, 0xf3,
	0x5b, 0x5a, 0xfe, 0xb8, 0xf5, 0xdd, 0xbd, 0x7b, 0xf7, 0xeb, 0x41, 0x6d, 0xc9, 0x4c, 0x9e, 0xba,
	0xd9, 0x4e, 0x7e, 0x42, 0x96, 0x14, 0x7f, 0x06, 0x36, 0x48, 0x33, 0x6b, 0xc2, 0x5b, 0xb8, 0xfc,
	0x95, 0x7a, 0xa1, 0x26, 0xbb, 0xcb, 0xda, 0x4c, 0x24, 0x81, 0xf7, 0x50, 0x7b, 0x65, 0xc8, 0xf9,
	0x60, 0xfa, 0x32, 0xae, 0xc9, 0x43, 0xb3, 0x10, 0xd5, 0x51, 0x7b, 0x73, 0xfc, 0x09, 0xca, 0x57,
	0x39, 0x92, 0x7e, 0x97, 0xca, 0x22, 0xc2, 0xc5, 0xb7, 0x34, 0x29, 0xa4, 0x14, 0xf1, 0x7d, 0x26,
	0xe7, 0xb1, 0x98, 0x04, 0x97, 0x00, 0x1f, 0x21, 0x65, 0x20, 0xab, 0xc8, 0xe1, 0x03, 0x30, 0xbd,
	0x86, 0xb8, 0x30, 0x59, 0xb4, 0x55, 0x7f, 0xd3, 0x6d, 0xb8, 0x47, 0xb4, 0xd8, 0x0c, 0xd8, 0xc2,
	0x95, 0x8b, 0xc8, 0x2b, 0xcd, 0xea, 0xbd, 0xde, 0xbd, 0xe9, 0x27, 0x62, 0x6f, 0xf3, 0x17, 0xa8,
	0x3f, 0xad, 0xf2, 0x24, 0xe8, 0x67, 0x09, 0xb8, 0xd8, 0x03, 0x44, 0xf0, 0xb8, 0x72, 0x5b, 0x84,
	0xfb, 0xf0, 0x09, 0x23, 0xfe, 0xb9, 0x46, 0x16, 0xef, 0xf5, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff,
	0x73, 0xb4, 0xd1, 0xb6, 0x3f, 0x01, 0x00, 0x00,
}
