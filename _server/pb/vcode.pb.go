// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: vcode.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type DriverEnum int32

const (
	DriverEnum_Digit    DriverEnum = 0
	DriverEnum_String   DriverEnum = 1
	DriverEnum_Math     DriverEnum = 2
	DriverEnum_Chinese  DriverEnum = 3
	DriverEnum_Audio    DriverEnum = 4
	DriverEnum_Language DriverEnum = 5
)

// Enum value maps for DriverEnum.
var (
	DriverEnum_name = map[int32]string{
		0: "Digit",
		1: "String",
		2: "Math",
		3: "Chinese",
		4: "Audio",
		5: "Language",
	}
	DriverEnum_value = map[string]int32{
		"Digit":    0,
		"String":   1,
		"Math":     2,
		"Chinese":  3,
		"Audio":    4,
		"Language": 5,
	}
)

func (x DriverEnum) Enum() *DriverEnum {
	p := new(DriverEnum)
	*p = x
	return p
}

func (x DriverEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DriverEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_vcode_proto_enumTypes[0].Descriptor()
}

func (DriverEnum) Type() protoreflect.EnumType {
	return &file_vcode_proto_enumTypes[0]
}

func (x DriverEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DriverEnum.Descriptor instead.
func (DriverEnum) EnumDescriptor() ([]byte, []int) {
	return file_vcode_proto_rawDescGZIP(), []int{0}
}

type LangEnum int32

const (
	LangEnum_Zh     LangEnum = 0
	LangEnum_Latin  LangEnum = 1
	LangEnum_Ko     LangEnum = 2
	LangEnum_Jp     LangEnum = 3
	LangEnum_Ru     LangEnum = 4
	LangEnum_Th     LangEnum = 5
	LangEnum_Greek  LangEnum = 6
	LangEnum_Arabic LangEnum = 7
	LangEnum_Hebrew LangEnum = 8
)

// Enum value maps for LangEnum.
var (
	LangEnum_name = map[int32]string{
		0: "Zh",
		1: "Latin",
		2: "Ko",
		3: "Jp",
		4: "Ru",
		5: "Th",
		6: "Greek",
		7: "Arabic",
		8: "Hebrew",
	}
	LangEnum_value = map[string]int32{
		"Zh":     0,
		"Latin":  1,
		"Ko":     2,
		"Jp":     3,
		"Ru":     4,
		"Th":     5,
		"Greek":  6,
		"Arabic": 7,
		"Hebrew": 8,
	}
)

func (x LangEnum) Enum() *LangEnum {
	p := new(LangEnum)
	*p = x
	return p
}

func (x LangEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LangEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_vcode_proto_enumTypes[1].Descriptor()
}

func (LangEnum) Type() protoreflect.EnumType {
	return &file_vcode_proto_enumTypes[1]
}

func (x LangEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LangEnum.Descriptor instead.
func (LangEnum) EnumDescriptor() ([]byte, []int) {
	return file_vcode_proto_rawDescGZIP(), []int{1}
}

// Generate请求参数
type GenerateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  int32      `protobuf:"varint,1,opt,name=AppID,proto3" json:"AppID,omitempty"`
	IP     string     `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	Driver DriverEnum `protobuf:"varint,3,opt,name=Driver,proto3,enum=pb.DriverEnum" json:"Driver,omitempty"`
	Height int32      `protobuf:"varint,4,opt,name=Height,proto3" json:"Height,omitempty"`
	Width  int32      `protobuf:"varint,5,opt,name=Width,proto3" json:"Width,omitempty"`
	Length int32      `protobuf:"varint,6,opt,name=Length,proto3" json:"Length,omitempty"`
	Lang   LangEnum   `protobuf:"varint,7,opt,name=Lang,proto3,enum=pb.LangEnum" json:"Lang,omitempty"`
}

func (x *GenerateReq) Reset() {
	*x = GenerateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vcode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateReq) ProtoMessage() {}

func (x *GenerateReq) ProtoReflect() protoreflect.Message {
	mi := &file_vcode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateReq.ProtoReflect.Descriptor instead.
func (*GenerateReq) Descriptor() ([]byte, []int) {
	return file_vcode_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateReq) GetAppID() int32 {
	if x != nil {
		return x.AppID
	}
	return 0
}

func (x *GenerateReq) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *GenerateReq) GetDriver() DriverEnum {
	if x != nil {
		return x.Driver
	}
	return DriverEnum_Digit
}

func (x *GenerateReq) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *GenerateReq) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *GenerateReq) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *GenerateReq) GetLang() LangEnum {
	if x != nil {
		return x.Lang
	}
	return LangEnum_Zh
}

// Generate返回结果
type GenerateRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	ID   string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GenerateRsp) Reset() {
	*x = GenerateRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vcode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRsp) ProtoMessage() {}

func (x *GenerateRsp) ProtoReflect() protoreflect.Message {
	mi := &file_vcode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRsp.ProtoReflect.Descriptor instead.
func (*GenerateRsp) Descriptor() ([]byte, []int) {
	return file_vcode_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateRsp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *GenerateRsp) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

// Verify请求参数
type VerifyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  int32  `protobuf:"varint,1,opt,name=AppID,proto3" json:"AppID,omitempty"`
	IP     string `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	ID     string `protobuf:"bytes,3,opt,name=ID,proto3" json:"ID,omitempty"`
	Answer string `protobuf:"bytes,4,opt,name=Answer,proto3" json:"Answer,omitempty"`
}

func (x *VerifyReq) Reset() {
	*x = VerifyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vcode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyReq) ProtoMessage() {}

func (x *VerifyReq) ProtoReflect() protoreflect.Message {
	mi := &file_vcode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyReq.ProtoReflect.Descriptor instead.
func (*VerifyReq) Descriptor() ([]byte, []int) {
	return file_vcode_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyReq) GetAppID() int32 {
	if x != nil {
		return x.AppID
	}
	return 0
}

func (x *VerifyReq) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *VerifyReq) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *VerifyReq) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

// 获取评论返回结果
type VerifyRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OK bool `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
}

func (x *VerifyRsp) Reset() {
	*x = VerifyRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vcode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRsp) ProtoMessage() {}

func (x *VerifyRsp) ProtoReflect() protoreflect.Message {
	mi := &file_vcode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRsp.ProtoReflect.Descriptor instead.
func (*VerifyRsp) Descriptor() ([]byte, []int) {
	return file_vcode_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyRsp) GetOK() bool {
	if x != nil {
		return x.OK
	}
	return false
}

// GenRawCode请求参数
type GenRawCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  int32  `protobuf:"varint,1,opt,name=AppID,proto3" json:"AppID,omitempty"`
	IP     string `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	Length int32  `protobuf:"varint,3,opt,name=Length,proto3" json:"Length,omitempty"`
	IDTag  string `protobuf:"bytes,4,opt,name=IDTag,proto3" json:"IDTag,omitempty"`
}

func (x *GenRawCodeReq) Reset() {
	*x = GenRawCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vcode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenRawCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenRawCodeReq) ProtoMessage() {}

func (x *GenRawCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_vcode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenRawCodeReq.ProtoReflect.Descriptor instead.
func (*GenRawCodeReq) Descriptor() ([]byte, []int) {
	return file_vcode_proto_rawDescGZIP(), []int{4}
}

func (x *GenRawCodeReq) GetAppID() int32 {
	if x != nil {
		return x.AppID
	}
	return 0
}

func (x *GenRawCodeReq) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *GenRawCodeReq) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *GenRawCodeReq) GetIDTag() string {
	if x != nil {
		return x.IDTag
	}
	return ""
}

var File_vcode_proto protoreflect.FileDescriptor

var file_vcode_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc3, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x26, 0x0a, 0x06, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x48,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x04, 0x4c, 0x61, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x45, 0x6e, 0x75, 0x6d, 0x52,
	0x04, 0x4c, 0x61, 0x6e, 0x67, 0x22, 0x31, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x59, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x22, 0x1b, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x4b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4f, 0x4b,
	0x22, 0x63, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x52, 0x61, 0x77, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12,
	0x14, 0x0a, 0x05, 0x49, 0x44, 0x54, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x49, 0x44, 0x54, 0x61, 0x67, 0x2a, 0x53, 0x0a, 0x0a, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x45,
	0x6e, 0x75, 0x6d, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x69, 0x67, 0x69, 0x74, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x61,
	0x74, 0x68, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x65, 0x10,
	0x03, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x10, 0x05, 0x2a, 0x60, 0x0a, 0x08, 0x4c, 0x61,
	0x6e, 0x67, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x06, 0x0a, 0x02, 0x5a, 0x68, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x4c, 0x61, 0x74, 0x69, 0x6e, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x4b, 0x6f, 0x10,
	0x02, 0x12, 0x06, 0x0a, 0x02, 0x4a, 0x70, 0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x52, 0x75, 0x10,
	0x04, 0x12, 0x06, 0x0a, 0x02, 0x54, 0x68, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x72, 0x65,
	0x65, 0x6b, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x72, 0x61, 0x62, 0x69, 0x63, 0x10, 0x07,
	0x12, 0x0a, 0x0a, 0x06, 0x48, 0x65, 0x62, 0x72, 0x65, 0x77, 0x10, 0x08, 0x32, 0xda, 0x01, 0x0a,
	0x0c, 0x56, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a,
	0x08, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x22, 0x14, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x3a, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x0d, 0x2e, 0x70, 0x62,
	0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x73, 0x70, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0c, 0x22, 0x07, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x4a, 0x0a,
	0x0a, 0x47, 0x65, 0x6e, 0x52, 0x61, 0x77, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x6e, 0x52, 0x61, 0x77, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0f,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x22,
	0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x67, 0x65, 0x6e, 0x5f, 0x72, 0x61,
	0x77, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x78, 0x69, 0x61, 0x6f, 0x78, 0x69, 0x6e,
	0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x43, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x2f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vcode_proto_rawDescOnce sync.Once
	file_vcode_proto_rawDescData = file_vcode_proto_rawDesc
)

func file_vcode_proto_rawDescGZIP() []byte {
	file_vcode_proto_rawDescOnce.Do(func() {
		file_vcode_proto_rawDescData = protoimpl.X.CompressGZIP(file_vcode_proto_rawDescData)
	})
	return file_vcode_proto_rawDescData
}

var file_vcode_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_vcode_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_vcode_proto_goTypes = []interface{}{
	(DriverEnum)(0),       // 0: pb.DriverEnum
	(LangEnum)(0),         // 1: pb.LangEnum
	(*GenerateReq)(nil),   // 2: pb.GenerateReq
	(*GenerateRsp)(nil),   // 3: pb.GenerateRsp
	(*VerifyReq)(nil),     // 4: pb.VerifyReq
	(*VerifyRsp)(nil),     // 5: pb.VerifyRsp
	(*GenRawCodeReq)(nil), // 6: pb.GenRawCodeReq
}
var file_vcode_proto_depIdxs = []int32{
	0, // 0: pb.GenerateReq.Driver:type_name -> pb.DriverEnum
	1, // 1: pb.GenerateReq.Lang:type_name -> pb.LangEnum
	2, // 2: pb.VcodeService.Generate:input_type -> pb.GenerateReq
	4, // 3: pb.VcodeService.Verify:input_type -> pb.VerifyReq
	6, // 4: pb.VcodeService.GenRawCode:input_type -> pb.GenRawCodeReq
	3, // 5: pb.VcodeService.Generate:output_type -> pb.GenerateRsp
	5, // 6: pb.VcodeService.Verify:output_type -> pb.VerifyRsp
	3, // 7: pb.VcodeService.GenRawCode:output_type -> pb.GenerateRsp
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_vcode_proto_init() }
func file_vcode_proto_init() {
	if File_vcode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vcode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateReq); i {
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
		file_vcode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateRsp); i {
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
		file_vcode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyReq); i {
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
		file_vcode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyRsp); i {
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
		file_vcode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenRawCodeReq); i {
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
			RawDescriptor: file_vcode_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vcode_proto_goTypes,
		DependencyIndexes: file_vcode_proto_depIdxs,
		EnumInfos:         file_vcode_proto_enumTypes,
		MessageInfos:      file_vcode_proto_msgTypes,
	}.Build()
	File_vcode_proto = out.File
	file_vcode_proto_rawDesc = nil
	file_vcode_proto_goTypes = nil
	file_vcode_proto_depIdxs = nil
}
