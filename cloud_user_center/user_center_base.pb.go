// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: user_center_base.proto

package cloud_user_center

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

// 基础用户结构
type BaseUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                           //用户id
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                              // 用户名
	Email           string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`                                            // 邮箱
	InstitutionName string `protobuf:"bytes,4,opt,name=institution_name,json=institutionName,proto3" json:"institution_name,omitempty"` // 机构名
	PhoneNumber     string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`             // 联系方式
	Grade           string `protobuf:"bytes,6,opt,name=grade,proto3" json:"grade,omitempty"`                                            // 职级
	Area            string `protobuf:"bytes,7,opt,name=area,proto3" json:"area,omitempty"`                                              // 区域
	VncPort         int32  `protobuf:"varint,8,opt,name=vnc_port,json=vncPort,proto3" json:"vnc_port,omitempty"`                        // vnc登录端口
	UserHost        string `protobuf:"bytes,9,opt,name=user_host,json=userHost,proto3" json:"user_host,omitempty"`                      // 用户主机
	PasswordLevel   int32  `protobuf:"varint,10,opt,name=password_level,json=passwordLevel,proto3" json:"password_level,omitempty"`     // 密码等级
}

func (x *BaseUserInfo) Reset() {
	*x = BaseUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseUserInfo) ProtoMessage() {}

func (x *BaseUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseUserInfo.ProtoReflect.Descriptor instead.
func (*BaseUserInfo) Descriptor() ([]byte, []int) {
	return file_user_center_base_proto_rawDescGZIP(), []int{0}
}

func (x *BaseUserInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BaseUserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BaseUserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *BaseUserInfo) GetInstitutionName() string {
	if x != nil {
		return x.InstitutionName
	}
	return ""
}

func (x *BaseUserInfo) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *BaseUserInfo) GetGrade() string {
	if x != nil {
		return x.Grade
	}
	return ""
}

func (x *BaseUserInfo) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *BaseUserInfo) GetVncPort() int32 {
	if x != nil {
		return x.VncPort
	}
	return 0
}

func (x *BaseUserInfo) GetUserHost() string {
	if x != nil {
		return x.UserHost
	}
	return ""
}

func (x *BaseUserInfo) GetPasswordLevel() int32 {
	if x != nil {
		return x.PasswordLevel
	}
	return 0
}

// 创建的用户结构
type CreateUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                           //用户id
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                              // 用户名
	Email           string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`                                            // 邮箱
	InstitutionName string `protobuf:"bytes,4,opt,name=institution_name,json=institutionName,proto3" json:"institution_name,omitempty"` // 机构名
	PhoneNumber     string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`             // 联系方式
	Grade           string `protobuf:"bytes,6,opt,name=grade,proto3" json:"grade,omitempty"`                                            // 职级
	Area            string `protobuf:"bytes,7,opt,name=area,proto3" json:"area,omitempty"`                                              // 区域
	Password        string `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`                                      // 密码, 对原始密码md5后的值
	PasswordLevel   int32  `protobuf:"varint,9,opt,name=password_level,json=passwordLevel,proto3" json:"password_level,omitempty"`      // 密码等级，业务方自定义
	VncPort         int32  `protobuf:"varint,10,opt,name=vnc_port,json=vncPort,proto3" json:"vnc_port,omitempty"`                       // optional vnc 端口，不传则表示不开通该功能，
	UserHost        string `protobuf:"bytes,11,opt,name=user_host,json=userHost,proto3" json:"user_host,omitempty"`                     // optional 用户使用主机ip:port, 不传则表示不开通该功能，
}

func (x *CreateUserInfo) Reset() {
	*x = CreateUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserInfo) ProtoMessage() {}

func (x *CreateUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserInfo.ProtoReflect.Descriptor instead.
func (*CreateUserInfo) Descriptor() ([]byte, []int) {
	return file_user_center_base_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateUserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserInfo) GetInstitutionName() string {
	if x != nil {
		return x.InstitutionName
	}
	return ""
}

func (x *CreateUserInfo) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreateUserInfo) GetGrade() string {
	if x != nil {
		return x.Grade
	}
	return ""
}

func (x *CreateUserInfo) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *CreateUserInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserInfo) GetPasswordLevel() int32 {
	if x != nil {
		return x.PasswordLevel
	}
	return 0
}

func (x *CreateUserInfo) GetVncPort() int32 {
	if x != nil {
		return x.VncPort
	}
	return 0
}

func (x *CreateUserInfo) GetUserHost() string {
	if x != nil {
		return x.UserHost
	}
	return ""
}

// 更新用户结构
type UpdateUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                           // 用户id
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                              // 用户名
	Password        string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`                                      // 密码
	PasswordLevel   int32  `protobuf:"varint,4,opt,name=password_level,json=passwordLevel,proto3" json:"password_level,omitempty"`      // 密码等级
	InstitutionName string `protobuf:"bytes,5,opt,name=institution_name,json=institutionName,proto3" json:"institution_name,omitempty"` // 机构名
	Grade           string `protobuf:"bytes,6,opt,name=grade,proto3" json:"grade,omitempty"`                                            // 职级
	Area            string `protobuf:"bytes,7,opt,name=area,proto3" json:"area,omitempty"`                                              // 区域
}

func (x *UpdateUserInfo) Reset() {
	*x = UpdateUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfo) ProtoMessage() {}

func (x *UpdateUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfo.ProtoReflect.Descriptor instead.
func (*UpdateUserInfo) Descriptor() ([]byte, []int) {
	return file_user_center_base_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateUserInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateUserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateUserInfo) GetPasswordLevel() int32 {
	if x != nil {
		return x.PasswordLevel
	}
	return 0
}

func (x *UpdateUserInfo) GetInstitutionName() string {
	if x != nil {
		return x.InstitutionName
	}
	return ""
}

func (x *UpdateUserInfo) GetGrade() string {
	if x != nil {
		return x.Grade
	}
	return ""
}

func (x *UpdateUserInfo) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                        //用户id
	RecordId       int64  `protobuf:"varint,2,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`                  // 记录id
	UserName       string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`                   // 用户名
	MessageTitle   string `protobuf:"bytes,4,opt,name=message_title,json=messageTitle,proto3" json:"message_title,omitempty"`       // 消息标题
	MessageContent string `protobuf:"bytes,5,opt,name=message_content,json=messageContent,proto3" json:"message_content,omitempty"` // 消息内容
	MessageDetail  string `protobuf:"bytes,6,opt,name=message_detail,json=messageDetail,proto3" json:"message_detail,omitempty"`    // 消息详情
	HasRead        int32  `protobuf:"varint,7,opt,name=has_read,json=hasRead,proto3" json:"has_read,omitempty"`                     // 是否已读, 0表示未读，其他表示已读
	CreatedAt      int64  `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`               // 创建时间
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_center_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_user_center_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_user_center_base_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Message) GetRecordId() int64 {
	if x != nil {
		return x.RecordId
	}
	return 0
}

func (x *Message) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Message) GetMessageTitle() string {
	if x != nil {
		return x.MessageTitle
	}
	return ""
}

func (x *Message) GetMessageContent() string {
	if x != nil {
		return x.MessageContent
	}
	return ""
}

func (x *Message) GetMessageDetail() string {
	if x != nil {
		return x.MessageDetail
	}
	return ""
}

func (x *Message) GetHasRead() int32 {
	if x != nil {
		return x.HasRead
	}
	return 0
}

func (x *Message) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

var File_user_center_base_proto protoreflect.FileDescriptor

var file_user_center_base_proto_rawDesc = []byte{
	0x0a, 0x16, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x02, 0x0a, 0x0c, 0x42, 0x61, 0x73,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x29, 0x0a, 0x10,
	0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x72, 0x65, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x6e, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x6e, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x22, 0xc6, 0x02, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x73,
	0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x65,
	0x61, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x6e, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x6e, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x22, 0xd5, 0x01, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x29, 0x0a, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x69,
	0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x72, 0x65, 0x61, 0x22, 0x8b, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f,
	0x72, 0x65, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x68, 0x61, 0x73, 0x52,
	0x65, 0x61, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x42, 0x27, 0x5a, 0x25, 0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x70, 0x62, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_user_center_base_proto_rawDescOnce sync.Once
	file_user_center_base_proto_rawDescData = file_user_center_base_proto_rawDesc
)

func file_user_center_base_proto_rawDescGZIP() []byte {
	file_user_center_base_proto_rawDescOnce.Do(func() {
		file_user_center_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_center_base_proto_rawDescData)
	})
	return file_user_center_base_proto_rawDescData
}

var file_user_center_base_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_center_base_proto_goTypes = []interface{}{
	(*BaseUserInfo)(nil),   // 0: BaseUserInfo
	(*CreateUserInfo)(nil), // 1: CreateUserInfo
	(*UpdateUserInfo)(nil), // 2: UpdateUserInfo
	(*Message)(nil),        // 3: Message
}
var file_user_center_base_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_center_base_proto_init() }
func file_user_center_base_proto_init() {
	if File_user_center_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_center_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseUserInfo); i {
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
		file_user_center_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserInfo); i {
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
		file_user_center_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserInfo); i {
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
		file_user_center_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_user_center_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_center_base_proto_goTypes,
		DependencyIndexes: file_user_center_base_proto_depIdxs,
		MessageInfos:      file_user_center_base_proto_msgTypes,
	}.Build()
	File_user_center_base_proto = out.File
	file_user_center_base_proto_rawDesc = nil
	file_user_center_base_proto_goTypes = nil
	file_user_center_base_proto_depIdxs = nil
}
