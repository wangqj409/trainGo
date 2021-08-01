// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: api/shop/v1/supply.proto

package v1

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

type SupplyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SupplyId   uint64 `protobuf:"varint,1,opt,name=supply_id,json=supplyId,proto3" json:"supply_id,omitempty"`
	SupplyName string `protobuf:"bytes,2,opt,name=supply_name,json=supplyName,proto3" json:"supply_name,omitempty"`
}

func (x *SupplyInfo) Reset() {
	*x = SupplyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_supply_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplyInfo) ProtoMessage() {}

func (x *SupplyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_supply_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplyInfo.ProtoReflect.Descriptor instead.
func (*SupplyInfo) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_supply_proto_rawDescGZIP(), []int{0}
}

func (x *SupplyInfo) GetSupplyId() uint64 {
	if x != nil {
		return x.SupplyId
	}
	return 0
}

func (x *SupplyInfo) GetSupplyName() string {
	if x != nil {
		return x.SupplyName
	}
	return ""
}

type CreateSupplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SupplyName string `protobuf:"bytes,1,opt,name=supply_name,json=supplyName,proto3" json:"supply_name,omitempty"`
}

func (x *CreateSupplyRequest) Reset() {
	*x = CreateSupplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_supply_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSupplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSupplyRequest) ProtoMessage() {}

func (x *CreateSupplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_supply_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSupplyRequest.ProtoReflect.Descriptor instead.
func (*CreateSupplyRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_supply_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSupplyRequest) GetSupplyName() string {
	if x != nil {
		return x.SupplyName
	}
	return ""
}

type CreateSupplyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SupplyInfo *SupplyInfo `protobuf:"bytes,1,opt,name=supply_info,json=supplyInfo,proto3" json:"supply_info,omitempty"`
}

func (x *CreateSupplyReply) Reset() {
	*x = CreateSupplyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_supply_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSupplyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSupplyReply) ProtoMessage() {}

func (x *CreateSupplyReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_supply_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSupplyReply.ProtoReflect.Descriptor instead.
func (*CreateSupplyReply) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_supply_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSupplyReply) GetSupplyInfo() *SupplyInfo {
	if x != nil {
		return x.SupplyInfo
	}
	return nil
}

type UpdateSupplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SupplyInfo *SupplyInfo `protobuf:"bytes,1,opt,name=supply_info,json=supplyInfo,proto3" json:"supply_info,omitempty"`
}

func (x *UpdateSupplyRequest) Reset() {
	*x = UpdateSupplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_supply_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSupplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSupplyRequest) ProtoMessage() {}

func (x *UpdateSupplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_supply_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSupplyRequest.ProtoReflect.Descriptor instead.
func (*UpdateSupplyRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_supply_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSupplyRequest) GetSupplyInfo() *SupplyInfo {
	if x != nil {
		return x.SupplyInfo
	}
	return nil
}

type UpdateSupplyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SupplyInfo *SupplyInfo `protobuf:"bytes,1,opt,name=supply_info,json=supplyInfo,proto3" json:"supply_info,omitempty"`
}

func (x *UpdateSupplyReply) Reset() {
	*x = UpdateSupplyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_supply_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSupplyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSupplyReply) ProtoMessage() {}

func (x *UpdateSupplyReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_supply_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSupplyReply.ProtoReflect.Descriptor instead.
func (*UpdateSupplyReply) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_supply_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSupplyReply) GetSupplyInfo() *SupplyInfo {
	if x != nil {
		return x.SupplyInfo
	}
	return nil
}

type DeleteSupplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SupplyInfo *SupplyInfo `protobuf:"bytes,1,opt,name=supply_info,json=supplyInfo,proto3" json:"supply_info,omitempty"`
}

func (x *DeleteSupplyRequest) Reset() {
	*x = DeleteSupplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_supply_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSupplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSupplyRequest) ProtoMessage() {}

func (x *DeleteSupplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_supply_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSupplyRequest.ProtoReflect.Descriptor instead.
func (*DeleteSupplyRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_supply_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteSupplyRequest) GetSupplyInfo() *SupplyInfo {
	if x != nil {
		return x.SupplyInfo
	}
	return nil
}

type DeleteSupplyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteSupplyReply) Reset() {
	*x = DeleteSupplyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_supply_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSupplyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSupplyReply) ProtoMessage() {}

func (x *DeleteSupplyReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_supply_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSupplyReply.ProtoReflect.Descriptor instead.
func (*DeleteSupplyReply) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_supply_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteSupplyReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type GetSupplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SupplyId uint64 `protobuf:"varint,1,opt,name=supply_id,json=supplyId,proto3" json:"supply_id,omitempty"`
}

func (x *GetSupplyRequest) Reset() {
	*x = GetSupplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_supply_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSupplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSupplyRequest) ProtoMessage() {}

func (x *GetSupplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_supply_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSupplyRequest.ProtoReflect.Descriptor instead.
func (*GetSupplyRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_supply_proto_rawDescGZIP(), []int{7}
}

func (x *GetSupplyRequest) GetSupplyId() uint64 {
	if x != nil {
		return x.SupplyId
	}
	return 0
}

type GetSupplyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SupplyInfo *SupplyInfo `protobuf:"bytes,1,opt,name=supply_info,json=supplyInfo,proto3" json:"supply_info,omitempty"`
}

func (x *GetSupplyReply) Reset() {
	*x = GetSupplyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_supply_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSupplyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSupplyReply) ProtoMessage() {}

func (x *GetSupplyReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_supply_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSupplyReply.ProtoReflect.Descriptor instead.
func (*GetSupplyReply) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_supply_proto_rawDescGZIP(), []int{8}
}

func (x *GetSupplyReply) GetSupplyInfo() *SupplyInfo {
	if x != nil {
		return x.SupplyInfo
	}
	return nil
}

type ListSupplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSupplyRequest) Reset() {
	*x = ListSupplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_supply_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSupplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSupplyRequest) ProtoMessage() {}

func (x *ListSupplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_supply_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSupplyRequest.ProtoReflect.Descriptor instead.
func (*ListSupplyRequest) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_supply_proto_rawDescGZIP(), []int{9}
}

type ListSupplyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Supplies []*SupplyInfo `protobuf:"bytes,1,rep,name=supplies,proto3" json:"supplies,omitempty"`
}

func (x *ListSupplyReply) Reset() {
	*x = ListSupplyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_shop_v1_supply_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSupplyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSupplyReply) ProtoMessage() {}

func (x *ListSupplyReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shop_v1_supply_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSupplyReply.ProtoReflect.Descriptor instead.
func (*ListSupplyReply) Descriptor() ([]byte, []int) {
	return file_api_shop_v1_supply_proto_rawDescGZIP(), []int{10}
}

func (x *ListSupplyReply) GetSupplies() []*SupplyInfo {
	if x != nil {
		return x.Supplies
	}
	return nil
}

var File_api_shop_v1_supply_proto protoreflect.FileDescriptor

var file_api_shop_v1_supply_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x0a, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x36, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70,
	0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x38,
	0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4f, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x38, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4d, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x38,
	0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x38, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2b, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x75,
	0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x75, 0x70,
	0x70, 0x6c, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x33, 0x0a, 0x08, 0x73,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x73,
	0x32, 0xf4, 0x03, 0x0a, 0x06, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x61, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x09, 0x22, 0x07, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x61,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x20,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x22, 0x07, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x12, 0x61, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x22, 0x07, 0x2f, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x79, 0x12, 0x64, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x2f, 0x7b,
	0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x5b, 0x0a, 0x0a, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07,
	0x2f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x42, 0x25, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x14, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_shop_v1_supply_proto_rawDescOnce sync.Once
	file_api_shop_v1_supply_proto_rawDescData = file_api_shop_v1_supply_proto_rawDesc
)

func file_api_shop_v1_supply_proto_rawDescGZIP() []byte {
	file_api_shop_v1_supply_proto_rawDescOnce.Do(func() {
		file_api_shop_v1_supply_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_shop_v1_supply_proto_rawDescData)
	})
	return file_api_shop_v1_supply_proto_rawDescData
}

var file_api_shop_v1_supply_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_shop_v1_supply_proto_goTypes = []interface{}{
	(*SupplyInfo)(nil),          // 0: api.shop.v1.SupplyInfo
	(*CreateSupplyRequest)(nil), // 1: api.shop.v1.CreateSupplyRequest
	(*CreateSupplyReply)(nil),   // 2: api.shop.v1.CreateSupplyReply
	(*UpdateSupplyRequest)(nil), // 3: api.shop.v1.UpdateSupplyRequest
	(*UpdateSupplyReply)(nil),   // 4: api.shop.v1.UpdateSupplyReply
	(*DeleteSupplyRequest)(nil), // 5: api.shop.v1.DeleteSupplyRequest
	(*DeleteSupplyReply)(nil),   // 6: api.shop.v1.DeleteSupplyReply
	(*GetSupplyRequest)(nil),    // 7: api.shop.v1.GetSupplyRequest
	(*GetSupplyReply)(nil),      // 8: api.shop.v1.GetSupplyReply
	(*ListSupplyRequest)(nil),   // 9: api.shop.v1.ListSupplyRequest
	(*ListSupplyReply)(nil),     // 10: api.shop.v1.ListSupplyReply
}
var file_api_shop_v1_supply_proto_depIdxs = []int32{
	0,  // 0: api.shop.v1.CreateSupplyReply.supply_info:type_name -> api.shop.v1.SupplyInfo
	0,  // 1: api.shop.v1.UpdateSupplyRequest.supply_info:type_name -> api.shop.v1.SupplyInfo
	0,  // 2: api.shop.v1.UpdateSupplyReply.supply_info:type_name -> api.shop.v1.SupplyInfo
	0,  // 3: api.shop.v1.DeleteSupplyRequest.supply_info:type_name -> api.shop.v1.SupplyInfo
	0,  // 4: api.shop.v1.GetSupplyReply.supply_info:type_name -> api.shop.v1.SupplyInfo
	0,  // 5: api.shop.v1.ListSupplyReply.supplies:type_name -> api.shop.v1.SupplyInfo
	1,  // 6: api.shop.v1.Supply.CreateSupply:input_type -> api.shop.v1.CreateSupplyRequest
	3,  // 7: api.shop.v1.Supply.UpdateSupply:input_type -> api.shop.v1.UpdateSupplyRequest
	5,  // 8: api.shop.v1.Supply.DeleteSupply:input_type -> api.shop.v1.DeleteSupplyRequest
	7,  // 9: api.shop.v1.Supply.GetSupply:input_type -> api.shop.v1.GetSupplyRequest
	9,  // 10: api.shop.v1.Supply.ListSupply:input_type -> api.shop.v1.ListSupplyRequest
	2,  // 11: api.shop.v1.Supply.CreateSupply:output_type -> api.shop.v1.CreateSupplyReply
	4,  // 12: api.shop.v1.Supply.UpdateSupply:output_type -> api.shop.v1.UpdateSupplyReply
	6,  // 13: api.shop.v1.Supply.DeleteSupply:output_type -> api.shop.v1.DeleteSupplyReply
	8,  // 14: api.shop.v1.Supply.GetSupply:output_type -> api.shop.v1.GetSupplyReply
	10, // 15: api.shop.v1.Supply.ListSupply:output_type -> api.shop.v1.ListSupplyReply
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_shop_v1_supply_proto_init() }
func file_api_shop_v1_supply_proto_init() {
	if File_api_shop_v1_supply_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_shop_v1_supply_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SupplyInfo); i {
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
		file_api_shop_v1_supply_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSupplyRequest); i {
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
		file_api_shop_v1_supply_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSupplyReply); i {
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
		file_api_shop_v1_supply_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSupplyRequest); i {
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
		file_api_shop_v1_supply_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSupplyReply); i {
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
		file_api_shop_v1_supply_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSupplyRequest); i {
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
		file_api_shop_v1_supply_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSupplyReply); i {
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
		file_api_shop_v1_supply_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSupplyRequest); i {
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
		file_api_shop_v1_supply_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSupplyReply); i {
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
		file_api_shop_v1_supply_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSupplyRequest); i {
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
		file_api_shop_v1_supply_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSupplyReply); i {
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
			RawDescriptor: file_api_shop_v1_supply_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_shop_v1_supply_proto_goTypes,
		DependencyIndexes: file_api_shop_v1_supply_proto_depIdxs,
		MessageInfos:      file_api_shop_v1_supply_proto_msgTypes,
	}.Build()
	File_api_shop_v1_supply_proto = out.File
	file_api_shop_v1_supply_proto_rawDesc = nil
	file_api_shop_v1_supply_proto_goTypes = nil
	file_api_shop_v1_supply_proto_depIdxs = nil
}
