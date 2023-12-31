// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpc_cart.proto

package pb

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

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId   int32  `protobuf:"varint,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	ProductName string `protobuf:"bytes,2,opt,name=productName,proto3" json:"productName,omitempty"`
	Price       int32  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Amounts     int32  `protobuf:"varint,5,opt,name=amounts,proto3" json:"amounts,omitempty"`
	Image       string `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	Status      string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_rpc_cart_proto_rawDescGZIP(), []int{0}
}

func (x *CartItem) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CartItem) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *CartItem) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CartItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CartItem) GetAmounts() int32 {
	if x != nil {
		return x.Amounts
	}
	return 0
}

func (x *CartItem) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CartItem) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Carts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string      `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items  []*CartItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Carts) Reset() {
	*x = Carts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Carts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Carts) ProtoMessage() {}

func (x *Carts) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Carts.ProtoReflect.Descriptor instead.
func (*Carts) Descriptor() ([]byte, []int) {
	return file_rpc_cart_proto_rawDescGZIP(), []int{1}
}

func (x *Carts) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Carts) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type AddCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32       `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items  []*CartItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *AddCartRequest) Reset() {
	*x = AddCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCartRequest) ProtoMessage() {}

func (x *AddCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCartRequest.ProtoReflect.Descriptor instead.
func (*AddCartRequest) Descriptor() ([]byte, []int) {
	return file_rpc_cart_proto_rawDescGZIP(), []int{2}
}

func (x *AddCartRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddCartRequest) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type AddCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32       `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Item   []*CartItem `protobuf:"bytes,2,rep,name=item,proto3" json:"item,omitempty"`
}

func (x *AddCartResponse) Reset() {
	*x = AddCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCartResponse) ProtoMessage() {}

func (x *AddCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCartResponse.ProtoReflect.Descriptor instead.
func (*AddCartResponse) Descriptor() ([]byte, []int) {
	return file_rpc_cart_proto_rawDescGZIP(), []int{3}
}

func (x *AddCartResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddCartResponse) GetItem() []*CartItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type GetCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetCartRequest) Reset() {
	*x = GetCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartRequest) ProtoMessage() {}

func (x *GetCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartRequest.ProtoReflect.Descriptor instead.
func (*GetCartRequest) Descriptor() ([]byte, []int) {
	return file_rpc_cart_proto_rawDescGZIP(), []int{4}
}

func (x *GetCartRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32       `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items  []*CartItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetCartResponse) Reset() {
	*x = GetCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartResponse) ProtoMessage() {}

func (x *GetCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartResponse.ProtoReflect.Descriptor instead.
func (*GetCartResponse) Descriptor() ([]byte, []int) {
	return file_rpc_cart_proto_rawDescGZIP(), []int{5}
}

func (x *GetCartResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetCartResponse) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type DeleteCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteCartRequest) Reset() {
	*x = DeleteCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartRequest) ProtoMessage() {}

func (x *DeleteCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartRequest.ProtoReflect.Descriptor instead.
func (*DeleteCartRequest) Descriptor() ([]byte, []int) {
	return file_rpc_cart_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCartRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DeleteCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *CartItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *DeleteCartResponse) Reset() {
	*x = DeleteCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartResponse) ProtoMessage() {}

func (x *DeleteCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartResponse.ProtoReflect.Descriptor instead.
func (*DeleteCartResponse) Descriptor() ([]byte, []int) {
	return file_rpc_cart_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCartResponse) GetItem() *CartItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type DeleteItemCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId int32 `protobuf:"varint,2,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
}

func (x *DeleteItemCartRequest) Reset() {
	*x = DeleteItemCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cart_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemCartRequest) ProtoMessage() {}

func (x *DeleteItemCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cart_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemCartRequest.ProtoReflect.Descriptor instead.
func (*DeleteItemCartRequest) Descriptor() ([]byte, []int) {
	return file_rpc_cart_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteItemCartRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DeleteItemCartRequest) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type DeleteItemCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *CartItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *DeleteItemCartResponse) Reset() {
	*x = DeleteItemCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cart_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemCartResponse) ProtoMessage() {}

func (x *DeleteItemCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cart_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemCartResponse.ProtoReflect.Descriptor instead.
func (*DeleteItemCartResponse) Descriptor() ([]byte, []int) {
	return file_rpc_cart_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteItemCartResponse) GetItem() *CartItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32       `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items  []*CartItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cart_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cart_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_rpc_cart_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateRequest) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *CartItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_cart_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_cart_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_rpc_cart_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateResponse) GetItem() *CartItem {
	if x != nil {
		return x.Item
	}
	return nil
}

var File_rpc_cart_proto protoreflect.FileDescriptor

var file_rpc_cart_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0xca, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x44, 0x0a, 0x05, 0x43, 0x61, 0x72, 0x74, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x4d, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x4c, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x4e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x36, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x4e, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x22, 0x4c, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x32, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69,
	0x74, 0x65, 0x6d, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x61, 0x72, 0x68, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x61, 0x6e, 0x2f, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_cart_proto_rawDescOnce sync.Once
	file_rpc_cart_proto_rawDescData = file_rpc_cart_proto_rawDesc
)

func file_rpc_cart_proto_rawDescGZIP() []byte {
	file_rpc_cart_proto_rawDescOnce.Do(func() {
		file_rpc_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_cart_proto_rawDescData)
	})
	return file_rpc_cart_proto_rawDescData
}

var file_rpc_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_rpc_cart_proto_goTypes = []interface{}{
	(*CartItem)(nil),               // 0: pb.CartItem
	(*Carts)(nil),                  // 1: pb.Carts
	(*AddCartRequest)(nil),         // 2: pb.AddCartRequest
	(*AddCartResponse)(nil),        // 3: pb.AddCartResponse
	(*GetCartRequest)(nil),         // 4: pb.GetCartRequest
	(*GetCartResponse)(nil),        // 5: pb.GetCartResponse
	(*DeleteCartRequest)(nil),      // 6: pb.DeleteCartRequest
	(*DeleteCartResponse)(nil),     // 7: pb.DeleteCartResponse
	(*DeleteItemCartRequest)(nil),  // 8: pb.DeleteItemCartRequest
	(*DeleteItemCartResponse)(nil), // 9: pb.DeleteItemCartResponse
	(*UpdateRequest)(nil),          // 10: pb.UpdateRequest
	(*UpdateResponse)(nil),         // 11: pb.UpdateResponse
}
var file_rpc_cart_proto_depIdxs = []int32{
	0, // 0: pb.Carts.items:type_name -> pb.CartItem
	0, // 1: pb.AddCartRequest.items:type_name -> pb.CartItem
	0, // 2: pb.AddCartResponse.item:type_name -> pb.CartItem
	0, // 3: pb.GetCartResponse.items:type_name -> pb.CartItem
	0, // 4: pb.DeleteCartResponse.item:type_name -> pb.CartItem
	0, // 5: pb.DeleteItemCartResponse.item:type_name -> pb.CartItem
	0, // 6: pb.UpdateRequest.items:type_name -> pb.CartItem
	0, // 7: pb.UpdateResponse.item:type_name -> pb.CartItem
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_rpc_cart_proto_init() }
func file_rpc_cart_proto_init() {
	if File_rpc_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
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
		file_rpc_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Carts); i {
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
		file_rpc_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCartRequest); i {
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
		file_rpc_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCartResponse); i {
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
		file_rpc_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartRequest); i {
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
		file_rpc_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartResponse); i {
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
		file_rpc_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartRequest); i {
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
		file_rpc_cart_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartResponse); i {
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
		file_rpc_cart_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemCartRequest); i {
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
		file_rpc_cart_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemCartResponse); i {
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
		file_rpc_cart_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_rpc_cart_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
			RawDescriptor: file_rpc_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_cart_proto_goTypes,
		DependencyIndexes: file_rpc_cart_proto_depIdxs,
		MessageInfos:      file_rpc_cart_proto_msgTypes,
	}.Build()
	File_rpc_cart_proto = out.File
	file_rpc_cart_proto_rawDesc = nil
	file_rpc_cart_proto_goTypes = nil
	file_rpc_cart_proto_depIdxs = nil
}
