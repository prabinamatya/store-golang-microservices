// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api_docs/catalogs/protobuf/products/service_clients/product_service_client.proto

package service_clients

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID   string                 `protobuf:"bytes,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Price       float64                `protobuf:"fixed64,4,opt,name=Price,proto3" json:"Price,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Product) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID   string  `protobuf:"bytes,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Price       float64 `protobuf:"fixed64,4,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *CreateProductReq) Reset() {
	*x = CreateProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductReq) ProtoMessage() {}

func (x *CreateProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductReq.ProtoReflect.Descriptor instead.
func (*CreateProductReq) Descriptor() ([]byte, []int) {
	return file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductReq) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *CreateProductReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductReq) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CreateProductRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID string `protobuf:"bytes,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
}

func (x *CreateProductRes) Reset() {
	*x = CreateProductRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRes) ProtoMessage() {}

func (x *CreateProductRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRes.ProtoReflect.Descriptor instead.
func (*CreateProductRes) Descriptor() ([]byte, []int) {
	return file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDescGZIP(), []int{2}
}

func (x *CreateProductRes) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

type UpdateProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID   string  `protobuf:"bytes,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Price       float64 `protobuf:"fixed64,4,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *UpdateProductReq) Reset() {
	*x = UpdateProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductReq) ProtoMessage() {}

func (x *UpdateProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductReq.ProtoReflect.Descriptor instead.
func (*UpdateProductReq) Descriptor() ([]byte, []int) {
	return file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateProductReq) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *UpdateProductReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProductReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateProductReq) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type UpdateProductRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateProductRes) Reset() {
	*x = UpdateProductRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductRes) ProtoMessage() {}

func (x *UpdateProductRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductRes.ProtoReflect.Descriptor instead.
func (*UpdateProductRes) Descriptor() ([]byte, []int) {
	return file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDescGZIP(), []int{4}
}

type GetProductByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID string `protobuf:"bytes,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
}

func (x *GetProductByIdReq) Reset() {
	*x = GetProductByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductByIdReq) ProtoMessage() {}

func (x *GetProductByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductByIdReq.ProtoReflect.Descriptor instead.
func (*GetProductByIdReq) Descriptor() ([]byte, []int) {
	return file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductByIdReq) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

type GetProductByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=Product,proto3" json:"Product,omitempty"`
}

func (x *GetProductByIdRes) Reset() {
	*x = GetProductByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductByIdRes) ProtoMessage() {}

func (x *GetProductByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductByIdRes.ProtoReflect.Descriptor instead.
func (*GetProductByIdRes) Descriptor() ([]byte, []int) {
	return file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDescGZIP(), []int{6}
}

func (x *GetProductByIdRes) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

var File_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto protoreflect.FileDescriptor

var file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDesc = []byte{
	0x0a, 0x50, 0x61, 0x70, 0x69, 0x5f, 0x64, 0x6f, 0x63, 0x73, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7c,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x30, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x22, 0x7c,
	0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x12, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x22, 0x31, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x44, 0x22, 0x47, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x32, 0x99, 0x02, 0x0a,
	0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x55, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x12, 0x58,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDescOnce sync.Once
	file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDescData = file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDesc
)

func file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDescGZIP() []byte {
	file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDescOnce.Do(func() {
		file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDescData)
	})
	return file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDescData
}

var file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_goTypes = []interface{}{
	(*Product)(nil),               // 0: service_clients.Product
	(*CreateProductReq)(nil),      // 1: service_clients.CreateProductReq
	(*CreateProductRes)(nil),      // 2: service_clients.CreateProductRes
	(*UpdateProductReq)(nil),      // 3: service_clients.UpdateProductReq
	(*UpdateProductRes)(nil),      // 4: service_clients.UpdateProductRes
	(*GetProductByIdReq)(nil),     // 5: service_clients.GetProductByIdReq
	(*GetProductByIdRes)(nil),     // 6: service_clients.GetProductByIdRes
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_depIdxs = []int32{
	7, // 0: service_clients.Product.CreatedAt:type_name -> google.protobuf.Timestamp
	7, // 1: service_clients.Product.UpdatedAt:type_name -> google.protobuf.Timestamp
	0, // 2: service_clients.GetProductByIdRes.Product:type_name -> service_clients.Product
	1, // 3: service_clients.ProductsService.CreateProduct:input_type -> service_clients.CreateProductReq
	3, // 4: service_clients.ProductsService.UpdateProduct:input_type -> service_clients.UpdateProductReq
	5, // 5: service_clients.ProductsService.GetProductById:input_type -> service_clients.GetProductByIdReq
	2, // 6: service_clients.ProductsService.CreateProduct:output_type -> service_clients.CreateProductRes
	4, // 7: service_clients.ProductsService.UpdateProduct:output_type -> service_clients.UpdateProductRes
	6, // 8: service_clients.ProductsService.GetProductById:output_type -> service_clients.GetProductByIdRes
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_init()
}
func file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_init() {
	if File_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductReq); i {
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
		file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductRes); i {
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
		file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductReq); i {
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
		file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductRes); i {
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
		file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductByIdReq); i {
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
		file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductByIdRes); i {
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
			RawDescriptor: file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_goTypes,
		DependencyIndexes: file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_depIdxs,
		MessageInfos:      file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_msgTypes,
	}.Build()
	File_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto = out.File
	file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_rawDesc = nil
	file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_goTypes = nil
	file_api_docs_catalogs_protobuf_products_service_clients_product_service_client_proto_depIdxs = nil
}
