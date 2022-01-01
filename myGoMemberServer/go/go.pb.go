// protoc --go_out=. --go_opt=paths=source_relative \
// --go-grpc_out=. --go-grpc_opt=paths=source_relative \
// mygrpc/go.proto

// protoc --go_out=. \
// --go-grpc_out=. \
// mygrpc/go.proto

//protoc --go_out=. *.proto

// protoc -I mygrpc/ --go_out=plugins=grpc:calculatorPB mygrpc/go.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.9.1
// source: mygrpc/go.proto

package protobuf

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

type QueryLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *QueryLogRequest) Reset() {
	*x = QueryLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mygrpc_go_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryLogRequest) ProtoMessage() {}

func (x *QueryLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mygrpc_go_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryLogRequest.ProtoReflect.Descriptor instead.
func (*QueryLogRequest) Descriptor() ([]byte, []int) {
	return file_mygrpc_go_proto_rawDescGZIP(), []int{0}
}

func (x *QueryLogRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type QueryLogIsExistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *QueryLogIsExistRequest) Reset() {
	*x = QueryLogIsExistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mygrpc_go_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryLogIsExistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryLogIsExistRequest) ProtoMessage() {}

func (x *QueryLogIsExistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mygrpc_go_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryLogIsExistRequest.ProtoReflect.Descriptor instead.
func (*QueryLogIsExistRequest) Descriptor() ([]byte, []int) {
	return file_mygrpc_go_proto_rawDescGZIP(), []int{1}
}

func (x *QueryLogIsExistRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type DeleteLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *DeleteLogRequest) Reset() {
	*x = DeleteLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mygrpc_go_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLogRequest) ProtoMessage() {}

func (x *DeleteLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mygrpc_go_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLogRequest.ProtoReflect.Descriptor instead.
func (*DeleteLogRequest) Descriptor() ([]byte, []int) {
	return file_mygrpc_go_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteLogRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type DeleteLogIsExistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *DeleteLogIsExistRequest) Reset() {
	*x = DeleteLogIsExistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mygrpc_go_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLogIsExistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLogIsExistRequest) ProtoMessage() {}

func (x *DeleteLogIsExistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mygrpc_go_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLogIsExistRequest.ProtoReflect.Descriptor instead.
func (*DeleteLogIsExistRequest) Descriptor() ([]byte, []int) {
	return file_mygrpc_go_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteLogIsExistRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type QueryLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account   string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Logintime []string `protobuf:"bytes,2,rep,name=logintime,proto3" json:"logintime,omitempty"`
}

func (x *QueryLogResponse) Reset() {
	*x = QueryLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mygrpc_go_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryLogResponse) ProtoMessage() {}

func (x *QueryLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mygrpc_go_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryLogResponse.ProtoReflect.Descriptor instead.
func (*QueryLogResponse) Descriptor() ([]byte, []int) {
	return file_mygrpc_go_proto_rawDescGZIP(), []int{4}
}

func (x *QueryLogResponse) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *QueryLogResponse) GetLogintime() []string {
	if x != nil {
		return x.Logintime
	}
	return nil
}

type QueryLogIsExistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Remark    map[string]string `protobuf:"bytes,1,rep,name=remark,proto3" json:"remark,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Dataexist bool              `protobuf:"varint,2,opt,name=dataexist,proto3" json:"dataexist,omitempty"`
}

func (x *QueryLogIsExistResponse) Reset() {
	*x = QueryLogIsExistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mygrpc_go_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryLogIsExistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryLogIsExistResponse) ProtoMessage() {}

func (x *QueryLogIsExistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mygrpc_go_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryLogIsExistResponse.ProtoReflect.Descriptor instead.
func (*QueryLogIsExistResponse) Descriptor() ([]byte, []int) {
	return file_mygrpc_go_proto_rawDescGZIP(), []int{5}
}

func (x *QueryLogIsExistResponse) GetRemark() map[string]string {
	if x != nil {
		return x.Remark
	}
	return nil
}

func (x *QueryLogIsExistResponse) GetDataexist() bool {
	if x != nil {
		return x.Dataexist
	}
	return false
}

type DeleteLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account   string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Logintime []string `protobuf:"bytes,2,rep,name=logintime,proto3" json:"logintime,omitempty"`
}

func (x *DeleteLogResponse) Reset() {
	*x = DeleteLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mygrpc_go_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLogResponse) ProtoMessage() {}

func (x *DeleteLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mygrpc_go_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLogResponse.ProtoReflect.Descriptor instead.
func (*DeleteLogResponse) Descriptor() ([]byte, []int) {
	return file_mygrpc_go_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteLogResponse) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *DeleteLogResponse) GetLogintime() []string {
	if x != nil {
		return x.Logintime
	}
	return nil
}

type DeleteLogIsExistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dataexist bool `protobuf:"varint,1,opt,name=dataexist,proto3" json:"dataexist,omitempty"`
}

func (x *DeleteLogIsExistResponse) Reset() {
	*x = DeleteLogIsExistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mygrpc_go_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLogIsExistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLogIsExistResponse) ProtoMessage() {}

func (x *DeleteLogIsExistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mygrpc_go_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLogIsExistResponse.ProtoReflect.Descriptor instead.
func (*DeleteLogIsExistResponse) Descriptor() ([]byte, []int) {
	return file_mygrpc_go_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteLogIsExistResponse) GetDataexist() bool {
	if x != nil {
		return x.Dataexist
	}
	return false
}

var File_mygrpc_go_proto protoreflect.FileDescriptor

var file_mygrpc_go_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x79, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x2b, 0x0a, 0x0f, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x4c, 0x6f, 0x67, 0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x33, 0x0a, 0x17, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x4a, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xb9, 0x01, 0x0a, 0x17,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x67, 0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x67, 0x49, 0x73, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x65, 0x78, 0x69, 0x73, 0x74, 0x1a, 0x39, 0x0a, 0x0b,
	0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4b, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f,
	0x67, 0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x65, 0x78, 0x69, 0x73, 0x74, 0x32, 0xd3,
	0x02, 0x0a, 0x0d, 0x4d, 0x79, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x58, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x67, 0x49, 0x73, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x67, 0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x67, 0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x08, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4c, 0x6f, 0x67, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5b, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x49, 0x73, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x49, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x49, 0x73, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x09,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x6d, 0x79, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mygrpc_go_proto_rawDescOnce sync.Once
	file_mygrpc_go_proto_rawDescData = file_mygrpc_go_proto_rawDesc
)

func file_mygrpc_go_proto_rawDescGZIP() []byte {
	file_mygrpc_go_proto_rawDescOnce.Do(func() {
		file_mygrpc_go_proto_rawDescData = protoimpl.X.CompressGZIP(file_mygrpc_go_proto_rawDescData)
	})
	return file_mygrpc_go_proto_rawDescData
}

var file_mygrpc_go_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_mygrpc_go_proto_goTypes = []interface{}{
	(*QueryLogRequest)(nil),          // 0: protobuf.QueryLogRequest
	(*QueryLogIsExistRequest)(nil),   // 1: protobuf.QueryLogIsExistRequest
	(*DeleteLogRequest)(nil),         // 2: protobuf.DeleteLogRequest
	(*DeleteLogIsExistRequest)(nil),  // 3: protobuf.DeleteLogIsExistRequest
	(*QueryLogResponse)(nil),         // 4: protobuf.QueryLogResponse
	(*QueryLogIsExistResponse)(nil),  // 5: protobuf.QueryLogIsExistResponse
	(*DeleteLogResponse)(nil),        // 6: protobuf.DeleteLogResponse
	(*DeleteLogIsExistResponse)(nil), // 7: protobuf.DeleteLogIsExistResponse
	nil,                              // 8: protobuf.QueryLogIsExistResponse.RemarkEntry
}
var file_mygrpc_go_proto_depIdxs = []int32{
	8, // 0: protobuf.QueryLogIsExistResponse.remark:type_name -> protobuf.QueryLogIsExistResponse.RemarkEntry
	1, // 1: protobuf.MygrpcService.QueryLogIsExist:input_type -> protobuf.QueryLogIsExistRequest
	0, // 2: protobuf.MygrpcService.QueryLog:input_type -> protobuf.QueryLogRequest
	3, // 3: protobuf.MygrpcService.DeleteLogIsExist:input_type -> protobuf.DeleteLogIsExistRequest
	2, // 4: protobuf.MygrpcService.DeleteLog:input_type -> protobuf.DeleteLogRequest
	5, // 5: protobuf.MygrpcService.QueryLogIsExist:output_type -> protobuf.QueryLogIsExistResponse
	4, // 6: protobuf.MygrpcService.QueryLog:output_type -> protobuf.QueryLogResponse
	7, // 7: protobuf.MygrpcService.DeleteLogIsExist:output_type -> protobuf.DeleteLogIsExistResponse
	6, // 8: protobuf.MygrpcService.DeleteLog:output_type -> protobuf.DeleteLogResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mygrpc_go_proto_init() }
func file_mygrpc_go_proto_init() {
	if File_mygrpc_go_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mygrpc_go_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryLogRequest); i {
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
		file_mygrpc_go_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryLogIsExistRequest); i {
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
		file_mygrpc_go_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLogRequest); i {
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
		file_mygrpc_go_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLogIsExistRequest); i {
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
		file_mygrpc_go_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryLogResponse); i {
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
		file_mygrpc_go_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryLogIsExistResponse); i {
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
		file_mygrpc_go_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLogResponse); i {
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
		file_mygrpc_go_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLogIsExistResponse); i {
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
			RawDescriptor: file_mygrpc_go_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mygrpc_go_proto_goTypes,
		DependencyIndexes: file_mygrpc_go_proto_depIdxs,
		MessageInfos:      file_mygrpc_go_proto_msgTypes,
	}.Build()
	File_mygrpc_go_proto = out.File
	file_mygrpc_go_proto_rawDesc = nil
	file_mygrpc_go_proto_goTypes = nil
	file_mygrpc_go_proto_depIdxs = nil
}
