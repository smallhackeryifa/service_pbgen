// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: service.proto

package label_dataset_srv

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
	0x11, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73,
	0x72, 0x76, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f,
	0x73, 0x72, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd2, 0x0e, 0x0a, 0x0c, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x76, 0x0a, 0x0b, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x5f, 0x73, 0x72, 0x76, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x22, 0x0d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a,
	0x01, 0x2a, 0x12, 0x6a, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x65, 0x47, 0x65, 0x74, 0x12, 0x22,
	0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73,
	0x72, 0x76, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22,
	0x0a, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x6e,
	0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x5f, 0x73, 0x72, 0x76, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x76,
	0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72,
	0x76, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x76, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x76,
	0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x25, 0x2e,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72,
	0x76, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x3a, 0x01, 0x2a, 0x12, 0x8f, 0x01, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x53, 0x61, 0x76, 0x65, 0x12, 0x2b, 0x2e, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x53, 0x61,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x53, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22,
	0x14, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x2f, 0x73, 0x61, 0x76, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x8b, 0x01, 0x0a, 0x10, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x47, 0x65, 0x74, 0x12, 0x2a, 0x2e,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72,
	0x76, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2f,
	0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x7e, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x5f, 0x73, 0x72, 0x76, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x22, 0x0f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x72, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x47, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x76, 0x0a, 0x0b, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x5f, 0x73, 0x72, 0x76, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12,
	0x22, 0x0d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a,
	0x01, 0x2a, 0x12, 0x7e, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x27, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x3a,
	0x01, 0x2a, 0x12, 0x7e, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x3a,
	0x01, 0x2a, 0x12, 0x7e, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a,
	0x01, 0x2a, 0x12, 0x7a, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x12, 0x26, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x27,
	0x5a, 0x25, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x70, 0x62, 0x67, 0x65, 0x6e, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x5f, 0x73, 0x72, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_proto_goTypes = []interface{}{
	(*StateCreateRequest)(nil),        // 0: label_dataset_srv.StateCreateRequest
	(*StateGetRequest)(nil),           // 1: label_dataset_srv.StateGetRequest
	(*StateListRequest)(nil),          // 2: label_dataset_srv.StateListRequest
	(*StateUpdateRequest)(nil),        // 3: label_dataset_srv.StateUpdateRequest
	(*StateDeleteRequest)(nil),        // 4: label_dataset_srv.StateDeleteRequest
	(*StateSearchRequest)(nil),        // 5: label_dataset_srv.StateSearchRequest
	(*StateSnapshotSaveRequest)(nil),  // 6: label_dataset_srv.StateSnapshotSaveRequest
	(*StateSnapshotGetRequest)(nil),   // 7: label_dataset_srv.StateSnapshotGetRequest
	(*DatasetCreateRequest)(nil),      // 8: label_dataset_srv.DatasetCreateRequest
	(*DatasetGetRequest)(nil),         // 9: label_dataset_srv.DatasetGetRequest
	(*DatasetListRequest)(nil),        // 10: label_dataset_srv.DatasetListRequest
	(*DatasetSearchRequest)(nil),      // 11: label_dataset_srv.DatasetSearchRequest
	(*DatasetUpdateRequest)(nil),      // 12: label_dataset_srv.DatasetUpdateRequest
	(*DatasetDeleteRequest)(nil),      // 13: label_dataset_srv.DatasetDeleteRequest
	(*DatasetShareRequest)(nil),       // 14: label_dataset_srv.DatasetShareRequest
	(*StateCreateResponse)(nil),       // 15: label_dataset_srv.StateCreateResponse
	(*StateGetResponse)(nil),          // 16: label_dataset_srv.StateGetResponse
	(*StateListResponse)(nil),         // 17: label_dataset_srv.StateListResponse
	(*StateUpdateResponse)(nil),       // 18: label_dataset_srv.StateUpdateResponse
	(*StateDeleteResponse)(nil),       // 19: label_dataset_srv.StateDeleteResponse
	(*StateSearchResponse)(nil),       // 20: label_dataset_srv.StateSearchResponse
	(*StateSnapshotSaveResponse)(nil), // 21: label_dataset_srv.StateSnapshotSaveResponse
	(*StateSnapshotGetResponse)(nil),  // 22: label_dataset_srv.StateSnapshotGetResponse
	(*DatasetCreateResponse)(nil),     // 23: label_dataset_srv.DatasetCreateResponse
	(*DatasetGetResponse)(nil),        // 24: label_dataset_srv.DatasetGetResponse
	(*DatasetListResponse)(nil),       // 25: label_dataset_srv.DatasetListResponse
	(*DatasetSearchResponse)(nil),     // 26: label_dataset_srv.DatasetSearchResponse
	(*DatasetUpdateResponse)(nil),     // 27: label_dataset_srv.DatasetUpdateResponse
	(*DatasetDeleteResponse)(nil),     // 28: label_dataset_srv.DatasetDeleteResponse
	(*DatasetShareResponse)(nil),      // 29: label_dataset_srv.DatasetShareResponse
}
var file_service_proto_depIdxs = []int32{
	0,  // 0: label_dataset_srv.LabelDataset.StateCreate:input_type -> label_dataset_srv.StateCreateRequest
	1,  // 1: label_dataset_srv.LabelDataset.StateGet:input_type -> label_dataset_srv.StateGetRequest
	2,  // 2: label_dataset_srv.LabelDataset.StateList:input_type -> label_dataset_srv.StateListRequest
	3,  // 3: label_dataset_srv.LabelDataset.StateUpdate:input_type -> label_dataset_srv.StateUpdateRequest
	4,  // 4: label_dataset_srv.LabelDataset.StateDelete:input_type -> label_dataset_srv.StateDeleteRequest
	5,  // 5: label_dataset_srv.LabelDataset.StateSearch:input_type -> label_dataset_srv.StateSearchRequest
	6,  // 6: label_dataset_srv.LabelDataset.StateSnapshotSave:input_type -> label_dataset_srv.StateSnapshotSaveRequest
	7,  // 7: label_dataset_srv.LabelDataset.StateSnapshotGet:input_type -> label_dataset_srv.StateSnapshotGetRequest
	8,  // 8: label_dataset_srv.LabelDataset.DatasetCreate:input_type -> label_dataset_srv.DatasetCreateRequest
	9,  // 9: label_dataset_srv.LabelDataset.DatasetGet:input_type -> label_dataset_srv.DatasetGetRequest
	10, // 10: label_dataset_srv.LabelDataset.DatasetList:input_type -> label_dataset_srv.DatasetListRequest
	11, // 11: label_dataset_srv.LabelDataset.DatasetSearch:input_type -> label_dataset_srv.DatasetSearchRequest
	12, // 12: label_dataset_srv.LabelDataset.DatasetUpdate:input_type -> label_dataset_srv.DatasetUpdateRequest
	13, // 13: label_dataset_srv.LabelDataset.DatasetDelete:input_type -> label_dataset_srv.DatasetDeleteRequest
	14, // 14: label_dataset_srv.LabelDataset.DatasetShare:input_type -> label_dataset_srv.DatasetShareRequest
	15, // 15: label_dataset_srv.LabelDataset.StateCreate:output_type -> label_dataset_srv.StateCreateResponse
	16, // 16: label_dataset_srv.LabelDataset.StateGet:output_type -> label_dataset_srv.StateGetResponse
	17, // 17: label_dataset_srv.LabelDataset.StateList:output_type -> label_dataset_srv.StateListResponse
	18, // 18: label_dataset_srv.LabelDataset.StateUpdate:output_type -> label_dataset_srv.StateUpdateResponse
	19, // 19: label_dataset_srv.LabelDataset.StateDelete:output_type -> label_dataset_srv.StateDeleteResponse
	20, // 20: label_dataset_srv.LabelDataset.StateSearch:output_type -> label_dataset_srv.StateSearchResponse
	21, // 21: label_dataset_srv.LabelDataset.StateSnapshotSave:output_type -> label_dataset_srv.StateSnapshotSaveResponse
	22, // 22: label_dataset_srv.LabelDataset.StateSnapshotGet:output_type -> label_dataset_srv.StateSnapshotGetResponse
	23, // 23: label_dataset_srv.LabelDataset.DatasetCreate:output_type -> label_dataset_srv.DatasetCreateResponse
	24, // 24: label_dataset_srv.LabelDataset.DatasetGet:output_type -> label_dataset_srv.DatasetGetResponse
	25, // 25: label_dataset_srv.LabelDataset.DatasetList:output_type -> label_dataset_srv.DatasetListResponse
	26, // 26: label_dataset_srv.LabelDataset.DatasetSearch:output_type -> label_dataset_srv.DatasetSearchResponse
	27, // 27: label_dataset_srv.LabelDataset.DatasetUpdate:output_type -> label_dataset_srv.DatasetUpdateResponse
	28, // 28: label_dataset_srv.LabelDataset.DatasetDelete:output_type -> label_dataset_srv.DatasetDeleteResponse
	29, // 29: label_dataset_srv.LabelDataset.DatasetShare:output_type -> label_dataset_srv.DatasetShareResponse
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_label_dataset_srv_proto_init()
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
