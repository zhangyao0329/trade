// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/imports/test_import_a1m1.proto

package imports

import (
	test_a_1 "google.golang.org/protobuf/cmd/protoc-gen-go/testdata/imports/test_a_1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type A1M1 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	F             *test_a_1.M1           `protobuf:"bytes,1,opt,name=f,proto3" json:"f,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *A1M1) Reset() {
	*x = A1M1{}
	mi := &file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *A1M1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*A1M1) ProtoMessage() {}

func (x *A1M1) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use A1M1.ProtoReflect.Descriptor instead.
func (*A1M1) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_rawDescGZIP(), []int{0}
}

func (x *A1M1) GetF() *test_a_1.M1 {
	if x != nil {
		return x.F
	}
	return nil
}

var File_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_rawDesc = []byte{
	0x0a, 0x39, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x61, 0x31, 0x6d, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73,
	0x74, 0x1a, 0x34, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x5f, 0x31, 0x2f, 0x6d,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x04, 0x41, 0x31, 0x4d, 0x31, 0x12,
	0x18, 0x0a, 0x01, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x61, 0x2e, 0x4d, 0x31, 0x52, 0x01, 0x66, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_rawDescData = file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_rawDesc
)

func file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_goTypes = []any{
	(*A1M1)(nil),        // 0: test.A1M1
	(*test_a_1.M1)(nil), // 1: test.a.M1
}
var file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_depIdxs = []int32{
	1, // 0: test.A1M1.f:type_name -> test.a.M1
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_init() }
func file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_init() {
	if File_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_depIdxs,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto = out.File
	file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_rawDesc = nil
	file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_imports_test_import_a1m1_proto_depIdxs = nil
}
