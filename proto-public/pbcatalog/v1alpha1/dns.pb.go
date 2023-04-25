// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: pbcatalog/v1alpha1/dns.proto

package catalogv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DNSPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workloads *Selector `protobuf:"bytes,1,opt,name=workloads,proto3" json:"workloads,omitempty"`
	Weights   *Weights  `protobuf:"bytes,2,opt,name=weights,proto3" json:"weights,omitempty"`
}

func (x *DNSPolicy) Reset() {
	*x = DNSPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbcatalog_v1alpha1_dns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSPolicy) ProtoMessage() {}

func (x *DNSPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_pbcatalog_v1alpha1_dns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSPolicy.ProtoReflect.Descriptor instead.
func (*DNSPolicy) Descriptor() ([]byte, []int) {
	return file_pbcatalog_v1alpha1_dns_proto_rawDescGZIP(), []int{0}
}

func (x *DNSPolicy) GetWorkloads() *Selector {
	if x != nil {
		return x.Workloads
	}
	return nil
}

func (x *DNSPolicy) GetWeights() *Weights {
	if x != nil {
		return x.Weights
	}
	return nil
}

type Weights struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Passing uint64 `protobuf:"varint,1,opt,name=passing,proto3" json:"passing,omitempty"`
	Warning uint64 `protobuf:"varint,2,opt,name=warning,proto3" json:"warning,omitempty"`
}

func (x *Weights) Reset() {
	*x = Weights{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbcatalog_v1alpha1_dns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Weights) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Weights) ProtoMessage() {}

func (x *Weights) ProtoReflect() protoreflect.Message {
	mi := &file_pbcatalog_v1alpha1_dns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Weights.ProtoReflect.Descriptor instead.
func (*Weights) Descriptor() ([]byte, []int) {
	return file_pbcatalog_v1alpha1_dns_proto_rawDescGZIP(), []int{1}
}

func (x *Weights) GetPassing() uint64 {
	if x != nil {
		return x.Passing
	}
	return 0
}

func (x *Weights) GetWarning() uint64 {
	if x != nil {
		return x.Warning
	}
	return 0
}

var File_pbcatalog_v1alpha1_dns_proto protoreflect.FileDescriptor

var file_pbcatalog_v1alpha1_dns_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x62, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x1a, 0x21, 0x70, 0x62, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x09, 0x44, 0x4e, 0x53, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x49, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x44, 0x0a,
	0x07, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x52, 0x07, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x22, 0x3d, 0x0a, 0x07, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x70, 0x61, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x42, 0xa5, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x08, 0x44, 0x6e,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2f, 0x70, 0x62, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x43, 0xaa, 0x02, 0x21, 0x48, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca,
	0x02, 0x21, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x5c, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xe2, 0x02, 0x2d, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x24, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a,
	0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pbcatalog_v1alpha1_dns_proto_rawDescOnce sync.Once
	file_pbcatalog_v1alpha1_dns_proto_rawDescData = file_pbcatalog_v1alpha1_dns_proto_rawDesc
)

func file_pbcatalog_v1alpha1_dns_proto_rawDescGZIP() []byte {
	file_pbcatalog_v1alpha1_dns_proto_rawDescOnce.Do(func() {
		file_pbcatalog_v1alpha1_dns_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbcatalog_v1alpha1_dns_proto_rawDescData)
	})
	return file_pbcatalog_v1alpha1_dns_proto_rawDescData
}

var file_pbcatalog_v1alpha1_dns_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pbcatalog_v1alpha1_dns_proto_goTypes = []interface{}{
	(*DNSPolicy)(nil), // 0: hashicorp.consul.catalog.v1alpha1.DNSPolicy
	(*Weights)(nil),   // 1: hashicorp.consul.catalog.v1alpha1.Weights
	(*Selector)(nil),  // 2: hashicorp.consul.catalog.v1alpha1.Selector
}
var file_pbcatalog_v1alpha1_dns_proto_depIdxs = []int32{
	2, // 0: hashicorp.consul.catalog.v1alpha1.DNSPolicy.workloads:type_name -> hashicorp.consul.catalog.v1alpha1.Selector
	1, // 1: hashicorp.consul.catalog.v1alpha1.DNSPolicy.weights:type_name -> hashicorp.consul.catalog.v1alpha1.Weights
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pbcatalog_v1alpha1_dns_proto_init() }
func file_pbcatalog_v1alpha1_dns_proto_init() {
	if File_pbcatalog_v1alpha1_dns_proto != nil {
		return
	}
	file_pbcatalog_v1alpha1_selector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pbcatalog_v1alpha1_dns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbcatalog_v1alpha1_dns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Weights); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbcatalog_v1alpha1_dns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbcatalog_v1alpha1_dns_proto_goTypes,
		DependencyIndexes: file_pbcatalog_v1alpha1_dns_proto_depIdxs,
		MessageInfos:      file_pbcatalog_v1alpha1_dns_proto_msgTypes,
	}.Build()
	File_pbcatalog_v1alpha1_dns_proto = out.File
	file_pbcatalog_v1alpha1_dns_proto_rawDesc = nil
	file_pbcatalog_v1alpha1_dns_proto_goTypes = nil
	file_pbcatalog_v1alpha1_dns_proto_depIdxs = nil
}
