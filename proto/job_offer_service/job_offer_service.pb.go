// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: job_offer_service.proto

package job_offer_service

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

type OfferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId      string   `protobuf:"bytes,1,opt,name=companyId,proto3" json:"companyId,omitempty"`
	Position       string   `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	JobDescription string   `protobuf:"bytes,3,opt,name=jobDescription,proto3" json:"jobDescription,omitempty"`
	Preconditions  []string `protobuf:"bytes,4,rep,name=preconditions,proto3" json:"preconditions,omitempty"`
}

func (x *OfferRequest) Reset() {
	*x = OfferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offer_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfferRequest) ProtoMessage() {}

func (x *OfferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_offer_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfferRequest.ProtoReflect.Descriptor instead.
func (*OfferRequest) Descriptor() ([]byte, []int) {
	return file_job_offer_service_proto_rawDescGZIP(), []int{0}
}

func (x *OfferRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *OfferRequest) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *OfferRequest) GetJobDescription() string {
	if x != nil {
		return x.JobDescription
	}
	return ""
}

func (x *OfferRequest) GetPreconditions() []string {
	if x != nil {
		return x.Preconditions
	}
	return nil
}

type Offer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobOfferId     string   `protobuf:"bytes,1,opt,name=jobOfferId,proto3" json:"jobOfferId,omitempty"`
	CompanyId      string   `protobuf:"bytes,2,opt,name=companyId,proto3" json:"companyId,omitempty"`
	Position       string   `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	JobDescription string   `protobuf:"bytes,4,opt,name=jobDescription,proto3" json:"jobDescription,omitempty"`
	Preconditions  []string `protobuf:"bytes,5,rep,name=preconditions,proto3" json:"preconditions,omitempty"`
}

func (x *Offer) Reset() {
	*x = Offer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offer_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Offer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Offer) ProtoMessage() {}

func (x *Offer) ProtoReflect() protoreflect.Message {
	mi := &file_job_offer_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Offer.ProtoReflect.Descriptor instead.
func (*Offer) Descriptor() ([]byte, []int) {
	return file_job_offer_service_proto_rawDescGZIP(), []int{1}
}

func (x *Offer) GetJobOfferId() string {
	if x != nil {
		return x.JobOfferId
	}
	return ""
}

func (x *Offer) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *Offer) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *Offer) GetJobDescription() string {
	if x != nil {
		return x.JobDescription
	}
	return ""
}

func (x *Offer) GetPreconditions() []string {
	if x != nil {
		return x.Preconditions
	}
	return nil
}

type Offers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offers []*Offer `protobuf:"bytes,1,rep,name=offers,proto3" json:"offers,omitempty"`
}

func (x *Offers) Reset() {
	*x = Offers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offer_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Offers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Offers) ProtoMessage() {}

func (x *Offers) ProtoReflect() protoreflect.Message {
	mi := &file_job_offer_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Offers.ProtoReflect.Descriptor instead.
func (*Offers) Descriptor() ([]byte, []int) {
	return file_job_offer_service_proto_rawDescGZIP(), []int{2}
}

func (x *Offers) GetOffers() []*Offer {
	if x != nil {
		return x.Offers
	}
	return nil
}

type CompanyID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyID string `protobuf:"bytes,1,opt,name=companyID,proto3" json:"companyID,omitempty"`
}

func (x *CompanyID) Reset() {
	*x = CompanyID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offer_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyID) ProtoMessage() {}

func (x *CompanyID) ProtoReflect() protoreflect.Message {
	mi := &file_job_offer_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyID.ProtoReflect.Descriptor instead.
func (*CompanyID) Descriptor() ([]byte, []int) {
	return file_job_offer_service_proto_rawDescGZIP(), []int{3}
}

func (x *CompanyID) GetCompanyID() string {
	if x != nil {
		return x.CompanyID
	}
	return ""
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position string `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offer_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_job_offer_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_job_offer_service_proto_rawDescGZIP(), []int{4}
}

func (x *Position) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

type CreatOfferFromAgent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId      string   `protobuf:"bytes,1,opt,name=companyId,proto3" json:"companyId,omitempty"`
	Position       string   `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	JobDescription string   `protobuf:"bytes,3,opt,name=jobDescription,proto3" json:"jobDescription,omitempty"`
	Preconditions  []string `protobuf:"bytes,4,rep,name=preconditions,proto3" json:"preconditions,omitempty"`
	Token          string   `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CreatOfferFromAgent) Reset() {
	*x = CreatOfferFromAgent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offer_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatOfferFromAgent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatOfferFromAgent) ProtoMessage() {}

func (x *CreatOfferFromAgent) ProtoReflect() protoreflect.Message {
	mi := &file_job_offer_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatOfferFromAgent.ProtoReflect.Descriptor instead.
func (*CreatOfferFromAgent) Descriptor() ([]byte, []int) {
	return file_job_offer_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreatOfferFromAgent) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *CreatOfferFromAgent) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *CreatOfferFromAgent) GetJobDescription() string {
	if x != nil {
		return x.JobDescription
	}
	return ""
}

func (x *CreatOfferFromAgent) GetPreconditions() []string {
	if x != nil {
		return x.Preconditions
	}
	return nil
}

func (x *CreatOfferFromAgent) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OfferId string `protobuf:"bytes,1,opt,name=offerId,proto3" json:"offerId,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offer_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_job_offer_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_job_offer_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRequest) GetOfferId() string {
	if x != nil {
		return x.OfferId
	}
	return ""
}

type EmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessage) Reset() {
	*x = EmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_offer_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessage) ProtoMessage() {}

func (x *EmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_job_offer_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessage.ProtoReflect.Descriptor instead.
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return file_job_offer_service_proto_rawDescGZIP(), []int{7}
}

var File_job_offer_service_proto protoreflect.FileDescriptor

var file_job_offer_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x0c, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a,
	0x0e, 0x6a, 0x6f, 0x62, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6a, 0x6f, 0x62, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x05,
	0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x6f, 0x62, 0x4f, 0x66, 0x66, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6a, 0x6f, 0x62, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x26, 0x0a, 0x0e, 0x6a, 0x6f, 0x62, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6a, 0x6f, 0x62, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x32, 0x0a,
	0x06, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66,
	0x66, 0x65, 0x72, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x65, 0x72,
	0x73, 0x22, 0x29, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x22, 0x26, 0x0a, 0x08,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb3, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x6a, 0x6f, 0x62, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x6a, 0x6f, 0x62, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24,
	0x0a, 0x0d, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x29, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x66,
	0x66, 0x65, 0x72, 0x49, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xb0, 0x04, 0x0a, 0x0f, 0x4a, 0x6f, 0x62, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x2d, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x56, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x12, 0x14, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x6a, 0x6f, 0x62, 0x5f,
	0x6f, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x22, 0x17, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x62, 0x79, 0x2d, 0x63, 0x6f,
	0x69, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x5a, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x65,
	0x72, 0x73, 0x42, 0x79, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x6a,
	0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x11, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x4f, 0x66,
	0x66, 0x65, 0x72, 0x73, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x67,
	0x65, 0x74, 0x2d, 0x62, 0x79, 0x2d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01,
	0x2a, 0x12, 0x64, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x1a, 0x10, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65,
	0x72, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22,
	0x13, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x53, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x12, 0x17, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66,
	0x66, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x11, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x4f, 0x66, 0x66,
	0x65, 0x72, 0x73, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x67, 0x65,
	0x74, 0x2d, 0x61, 0x6c, 0x6c, 0x2d, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x12, 0x5a, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x6a, 0x6f,
	0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65,
	0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2d,
	0x6f, 0x66, 0x66, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x42, 0x25, 0x5a, 0x23, 0x58, 0x57, 0x53, 0x2d,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x6f,
	0x62, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_offer_service_proto_rawDescOnce sync.Once
	file_job_offer_service_proto_rawDescData = file_job_offer_service_proto_rawDesc
)

func file_job_offer_service_proto_rawDescGZIP() []byte {
	file_job_offer_service_proto_rawDescOnce.Do(func() {
		file_job_offer_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_offer_service_proto_rawDescData)
	})
	return file_job_offer_service_proto_rawDescData
}

var file_job_offer_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_job_offer_service_proto_goTypes = []interface{}{
	(*OfferRequest)(nil),        // 0: job_offer.OfferRequest
	(*Offer)(nil),               // 1: job_offer.Offer
	(*Offers)(nil),              // 2: job_offer.Offers
	(*CompanyID)(nil),           // 3: job_offer.CompanyID
	(*Position)(nil),            // 4: job_offer.Position
	(*CreatOfferFromAgent)(nil), // 5: job_offer.CreatOfferFromAgent
	(*DeleteRequest)(nil),       // 6: job_offer.DeleteRequest
	(*EmptyMessage)(nil),        // 7: job_offer.EmptyMessage
}
var file_job_offer_service_proto_depIdxs = []int32{
	1, // 0: job_offer.Offers.offers:type_name -> job_offer.Offer
	0, // 1: job_offer.JobOfferService.CreateOffer:input_type -> job_offer.OfferRequest
	3, // 2: job_offer.JobOfferService.GetOffersByCompany:input_type -> job_offer.CompanyID
	4, // 3: job_offer.JobOfferService.GetOffersByPosition:input_type -> job_offer.Position
	5, // 4: job_offer.JobOfferService.CreateAgentOffer:input_type -> job_offer.CreatOfferFromAgent
	7, // 5: job_offer.JobOfferService.GetAllOffers:input_type -> job_offer.EmptyMessage
	6, // 6: job_offer.JobOfferService.DeleteOffer:input_type -> job_offer.DeleteRequest
	1, // 7: job_offer.JobOfferService.CreateOffer:output_type -> job_offer.Offer
	2, // 8: job_offer.JobOfferService.GetOffersByCompany:output_type -> job_offer.Offers
	2, // 9: job_offer.JobOfferService.GetOffersByPosition:output_type -> job_offer.Offers
	1, // 10: job_offer.JobOfferService.CreateAgentOffer:output_type -> job_offer.Offer
	2, // 11: job_offer.JobOfferService.GetAllOffers:output_type -> job_offer.Offers
	7, // 12: job_offer.JobOfferService.DeleteOffer:output_type -> job_offer.EmptyMessage
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_job_offer_service_proto_init() }
func file_job_offer_service_proto_init() {
	if File_job_offer_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_job_offer_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfferRequest); i {
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
		file_job_offer_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Offer); i {
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
		file_job_offer_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Offers); i {
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
		file_job_offer_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyID); i {
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
		file_job_offer_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_job_offer_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatOfferFromAgent); i {
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
		file_job_offer_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_job_offer_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessage); i {
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
			RawDescriptor: file_job_offer_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_job_offer_service_proto_goTypes,
		DependencyIndexes: file_job_offer_service_proto_depIdxs,
		MessageInfos:      file_job_offer_service_proto_msgTypes,
	}.Build()
	File_job_offer_service_proto = out.File
	file_job_offer_service_proto_rawDesc = nil
	file_job_offer_service_proto_goTypes = nil
	file_job_offer_service_proto_depIdxs = nil
}
