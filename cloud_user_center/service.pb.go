// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: service.proto

package cloud_user_center

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xe2, 0x0e, 0x0a, 0x0f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x7e, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x76, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x41, 0x75, 0x74, 0x68, 0x12, 0x26, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x22, 0x0a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x3a, 0x01, 0x2a, 0x12,
	0x7e, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x28, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x7e, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x28, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x72, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x47, 0x65, 0x74, 0x12, 0x25,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74,
	0x3a, 0x01, 0x2a, 0x12, 0x77, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x4d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x8c, 0x01, 0x0a,
	0x11, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x2b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74,
	0x5f, 0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x7e, 0x0a, 0x0e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x28, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x3a, 0x01, 0x2a, 0x12, 0x8f, 0x01, 0x0a, 0x11,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x2b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x94, 0x01,
	0x0a, 0x12, 0x4d, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x87, 0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x83,
	0x01, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x65,
	0x74, 0x12, 0x28, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x65,
	0x74, 0x3a, 0x01, 0x2a, 0x12, 0x8f, 0x01, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x8f, 0x01, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22,
	0x14, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x27, 0x5a, 0x25, 0x2e, 0x2e, 0x2f, 0x2e,
	0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x62, 0x67, 0x65, 0x6e, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*UserInfoCreateRequest)(nil),      // 0: cloud_user_center.UserInfoCreateRequest
	(*UserInfoAuthRequest)(nil),        // 1: cloud_user_center.UserInfoAuthRequest
	(*UserInfoUpdateRequest)(nil),      // 2: cloud_user_center.UserInfoUpdateRequest
	(*UserInfoDeleteRequest)(nil),      // 3: cloud_user_center.UserInfoDeleteRequest
	(*UserInfoGetRequest)(nil),         // 4: cloud_user_center.UserInfoGetRequest
	(*UserInfoListRequest)(nil),        // 5: cloud_user_center.UserInfoListRequest
	(*UserInfoGetByNameRequest)(nil),   // 6: cloud_user_center.UserInfoGetByNameRequest
	(*UserInfoSearchRequest)(nil),      // 7: cloud_user_center.UserInfoSearchRequest
	(*UserMessageCreateRequest)(nil),   // 8: cloud_user_center.UserMessageCreateRequest
	(*MUserMessageCreateRequest)(nil),  // 9: cloud_user_center.MUserMessageCreateRequest
	(*UserMessageListRequest)(nil),     // 10: cloud_user_center.UserMessageListRequest
	(*UserMessageGetRequest)(nil),      // 11: cloud_user_center.UserMessageGetRequest
	(*UserMessageUpdateRequest)(nil),   // 12: cloud_user_center.UserMessageUpdateRequest
	(*UserMessageDeleteRequest)(nil),   // 13: cloud_user_center.UserMessageDeleteRequest
	(*UserInfoCreateResponse)(nil),     // 14: cloud_user_center.UserInfoCreateResponse
	(*UserInfoAuthResponse)(nil),       // 15: cloud_user_center.UserInfoAuthResponse
	(*UserInfoUpdateResponse)(nil),     // 16: cloud_user_center.UserInfoUpdateResponse
	(*UserInfoDeleteResponse)(nil),     // 17: cloud_user_center.UserInfoDeleteResponse
	(*UserInfoGetResponse)(nil),        // 18: cloud_user_center.UserInfoGetResponse
	(*MUserInfoListResponse)(nil),      // 19: cloud_user_center.MUserInfoListResponse
	(*UserInfoGetByNameResponse)(nil),  // 20: cloud_user_center.UserInfoGetByNameResponse
	(*UserInfoSearchResponse)(nil),     // 21: cloud_user_center.UserInfoSearchResponse
	(*UserMessageCreateResponse)(nil),  // 22: cloud_user_center.UserMessageCreateResponse
	(*MUserMessageCreateResponse)(nil), // 23: cloud_user_center.MUserMessageCreateResponse
	(*UserMessageListResponse)(nil),    // 24: cloud_user_center.UserMessageListResponse
	(*UserMessageGetResponse)(nil),     // 25: cloud_user_center.UserMessageGetResponse
	(*UserMessageUpdateResponse)(nil),  // 26: cloud_user_center.UserMessageUpdateResponse
	(*UserMessageDeleteResponse)(nil),  // 27: cloud_user_center.UserMessageDeleteResponse
}
var file_service_proto_depIdxs = []int32{
	0,  // 0: cloud_user_center.CloudUserCenter.UserInfoCreate:input_type -> cloud_user_center.UserInfoCreateRequest
	1,  // 1: cloud_user_center.CloudUserCenter.UserInfoAuth:input_type -> cloud_user_center.UserInfoAuthRequest
	2,  // 2: cloud_user_center.CloudUserCenter.UserInfoUpdate:input_type -> cloud_user_center.UserInfoUpdateRequest
	3,  // 3: cloud_user_center.CloudUserCenter.UserInfoDelete:input_type -> cloud_user_center.UserInfoDeleteRequest
	4,  // 4: cloud_user_center.CloudUserCenter.UserInfoGet:input_type -> cloud_user_center.UserInfoGetRequest
	5,  // 5: cloud_user_center.CloudUserCenter.UserInfoList:input_type -> cloud_user_center.UserInfoListRequest
	6,  // 6: cloud_user_center.CloudUserCenter.UserInfoGetByName:input_type -> cloud_user_center.UserInfoGetByNameRequest
	7,  // 7: cloud_user_center.CloudUserCenter.UserInfoSearch:input_type -> cloud_user_center.UserInfoSearchRequest
	8,  // 8: cloud_user_center.CloudUserCenter.UserMessageCreate:input_type -> cloud_user_center.UserMessageCreateRequest
	9,  // 9: cloud_user_center.CloudUserCenter.MUserMessageCreate:input_type -> cloud_user_center.MUserMessageCreateRequest
	10, // 10: cloud_user_center.CloudUserCenter.UserMessageList:input_type -> cloud_user_center.UserMessageListRequest
	11, // 11: cloud_user_center.CloudUserCenter.UserMessageGet:input_type -> cloud_user_center.UserMessageGetRequest
	12, // 12: cloud_user_center.CloudUserCenter.UserMessageUpdate:input_type -> cloud_user_center.UserMessageUpdateRequest
	13, // 13: cloud_user_center.CloudUserCenter.UserMessageDelete:input_type -> cloud_user_center.UserMessageDeleteRequest
	14, // 14: cloud_user_center.CloudUserCenter.UserInfoCreate:output_type -> cloud_user_center.UserInfoCreateResponse
	15, // 15: cloud_user_center.CloudUserCenter.UserInfoAuth:output_type -> cloud_user_center.UserInfoAuthResponse
	16, // 16: cloud_user_center.CloudUserCenter.UserInfoUpdate:output_type -> cloud_user_center.UserInfoUpdateResponse
	17, // 17: cloud_user_center.CloudUserCenter.UserInfoDelete:output_type -> cloud_user_center.UserInfoDeleteResponse
	18, // 18: cloud_user_center.CloudUserCenter.UserInfoGet:output_type -> cloud_user_center.UserInfoGetResponse
	19, // 19: cloud_user_center.CloudUserCenter.UserInfoList:output_type -> cloud_user_center.MUserInfoListResponse
	20, // 20: cloud_user_center.CloudUserCenter.UserInfoGetByName:output_type -> cloud_user_center.UserInfoGetByNameResponse
	21, // 21: cloud_user_center.CloudUserCenter.UserInfoSearch:output_type -> cloud_user_center.UserInfoSearchResponse
	22, // 22: cloud_user_center.CloudUserCenter.UserMessageCreate:output_type -> cloud_user_center.UserMessageCreateResponse
	23, // 23: cloud_user_center.CloudUserCenter.MUserMessageCreate:output_type -> cloud_user_center.MUserMessageCreateResponse
	24, // 24: cloud_user_center.CloudUserCenter.UserMessageList:output_type -> cloud_user_center.UserMessageListResponse
	25, // 25: cloud_user_center.CloudUserCenter.UserMessageGet:output_type -> cloud_user_center.UserMessageGetResponse
	26, // 26: cloud_user_center.CloudUserCenter.UserMessageUpdate:output_type -> cloud_user_center.UserMessageUpdateResponse
	27, // 27: cloud_user_center.CloudUserCenter.UserMessageDelete:output_type -> cloud_user_center.UserMessageDeleteResponse
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_user_center_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
