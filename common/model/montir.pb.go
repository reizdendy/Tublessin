// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: montir.proto

package model

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MontirResponeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string         `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Code     string         `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Result   *MontirAccount `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MontirResponeMessage) Reset() {
	*x = MontirResponeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_montir_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MontirResponeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MontirResponeMessage) ProtoMessage() {}

func (x *MontirResponeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_montir_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MontirResponeMessage.ProtoReflect.Descriptor instead.
func (*MontirResponeMessage) Descriptor() ([]byte, []int) {
	return file_montir_proto_rawDescGZIP(), []int{0}
}

func (x *MontirResponeMessage) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *MontirResponeMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MontirResponeMessage) GetResult() *MontirAccount {
	if x != nil {
		return x.Result
	}
	return nil
}

type MontirAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string         `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password      string         `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	StatusAccount string         `protobuf:"bytes,5,opt,name=status_account,json=statusAccount,proto3" json:"status_account,omitempty"`
	Profile       *MontirProfile `protobuf:"bytes,6,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *MontirAccount) Reset() {
	*x = MontirAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_montir_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MontirAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MontirAccount) ProtoMessage() {}

func (x *MontirAccount) ProtoReflect() protoreflect.Message {
	mi := &file_montir_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MontirAccount.ProtoReflect.Descriptor instead.
func (*MontirAccount) Descriptor() ([]byte, []int) {
	return file_montir_proto_rawDescGZIP(), []int{1}
}

func (x *MontirAccount) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MontirAccount) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *MontirAccount) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *MontirAccount) GetStatusAccount() string {
	if x != nil {
		return x.StatusAccount
	}
	return ""
}

func (x *MontirAccount) GetProfile() *MontirProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type MontirProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname       string          `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname        string          `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	BornDate        string          `protobuf:"bytes,4,opt,name=born_date,json=bornDate,proto3" json:"born_date,omitempty"`
	Gender          string          `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	Ktp             string          `protobuf:"bytes,6,opt,name=ktp,proto3" json:"ktp,omitempty"`
	Address         string          `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	City            string          `protobuf:"bytes,8,opt,name=city,proto3" json:"city,omitempty"`
	Email           string          `protobuf:"bytes,9,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber     string          `protobuf:"bytes,10,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	ImageURL        string          `protobuf:"bytes,11,opt,name=imageURL,proto3" json:"imageURL,omitempty"`
	VerifiedAccount string          `protobuf:"bytes,12,opt,name=verified_account,json=verifiedAccount,proto3" json:"verified_account,omitempty"`
	DateUpdated     string          `protobuf:"bytes,13,opt,name=date_updated,json=dateUpdated,proto3" json:"date_updated,omitempty"`
	DateCreated     string          `protobuf:"bytes,14,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	Status          *MontirStatus   `protobuf:"bytes,15,opt,name=status,proto3" json:"status,omitempty"`
	RatingList      []*MontirRating `protobuf:"bytes,16,rep,name=rating_list,json=ratingList,proto3" json:"rating_list,omitempty"`
	Location        *MontirLocation `protobuf:"bytes,17,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *MontirProfile) Reset() {
	*x = MontirProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_montir_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MontirProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MontirProfile) ProtoMessage() {}

func (x *MontirProfile) ProtoReflect() protoreflect.Message {
	mi := &file_montir_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MontirProfile.ProtoReflect.Descriptor instead.
func (*MontirProfile) Descriptor() ([]byte, []int) {
	return file_montir_proto_rawDescGZIP(), []int{2}
}

func (x *MontirProfile) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MontirProfile) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *MontirProfile) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *MontirProfile) GetBornDate() string {
	if x != nil {
		return x.BornDate
	}
	return ""
}

func (x *MontirProfile) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *MontirProfile) GetKtp() string {
	if x != nil {
		return x.Ktp
	}
	return ""
}

func (x *MontirProfile) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MontirProfile) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *MontirProfile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *MontirProfile) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *MontirProfile) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *MontirProfile) GetVerifiedAccount() string {
	if x != nil {
		return x.VerifiedAccount
	}
	return ""
}

func (x *MontirProfile) GetDateUpdated() string {
	if x != nil {
		return x.DateUpdated
	}
	return ""
}

func (x *MontirProfile) GetDateCreated() string {
	if x != nil {
		return x.DateCreated
	}
	return ""
}

func (x *MontirProfile) GetStatus() *MontirStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *MontirProfile) GetRatingList() []*MontirRating {
	if x != nil {
		return x.RatingList
	}
	return nil
}

func (x *MontirProfile) GetLocation() *MontirLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

type MontirRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating      int32  `protobuf:"varint,1,opt,name=rating,proto3" json:"rating,omitempty"`
	RaterId     string `protobuf:"bytes,2,opt,name=rater_id,json=raterId,proto3" json:"rater_id,omitempty"`
	Review      string `protobuf:"bytes,3,opt,name=review,proto3" json:"review,omitempty"`
	DateCreated string `protobuf:"bytes,4,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
}

func (x *MontirRating) Reset() {
	*x = MontirRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_montir_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MontirRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MontirRating) ProtoMessage() {}

func (x *MontirRating) ProtoReflect() protoreflect.Message {
	mi := &file_montir_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MontirRating.ProtoReflect.Descriptor instead.
func (*MontirRating) Descriptor() ([]byte, []int) {
	return file_montir_proto_rawDescGZIP(), []int{3}
}

func (x *MontirRating) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *MontirRating) GetRaterId() string {
	if x != nil {
		return x.RaterId
	}
	return ""
}

func (x *MontirRating) GetReview() string {
	if x != nil {
		return x.Review
	}
	return ""
}

func (x *MontirRating) GetDateCreated() string {
	if x != nil {
		return x.DateCreated
	}
	return ""
}

type MontirLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude    float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude   float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	DateUpdated string  `protobuf:"bytes,3,opt,name=date_updated,json=dateUpdated,proto3" json:"date_updated,omitempty"`
}

func (x *MontirLocation) Reset() {
	*x = MontirLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_montir_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MontirLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MontirLocation) ProtoMessage() {}

func (x *MontirLocation) ProtoReflect() protoreflect.Message {
	mi := &file_montir_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MontirLocation.ProtoReflect.Descriptor instead.
func (*MontirLocation) Descriptor() ([]byte, []int) {
	return file_montir_proto_rawDescGZIP(), []int{4}
}

func (x *MontirLocation) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *MontirLocation) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *MontirLocation) GetDateUpdated() string {
	if x != nil {
		return x.DateUpdated
	}
	return ""
}

type MontirStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusOperational string `protobuf:"bytes,1,opt,name=status_operational,json=statusOperational,proto3" json:"status_operational,omitempty"`
	StatusActivity    string `protobuf:"bytes,2,opt,name=status_activity,json=statusActivity,proto3" json:"status_activity,omitempty"`
	DateUpdated       string `protobuf:"bytes,3,opt,name=date_updated,json=dateUpdated,proto3" json:"date_updated,omitempty"`
}

func (x *MontirStatus) Reset() {
	*x = MontirStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_montir_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MontirStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MontirStatus) ProtoMessage() {}

func (x *MontirStatus) ProtoReflect() protoreflect.Message {
	mi := &file_montir_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MontirStatus.ProtoReflect.Descriptor instead.
func (*MontirStatus) Descriptor() ([]byte, []int) {
	return file_montir_proto_rawDescGZIP(), []int{5}
}

func (x *MontirStatus) GetStatusOperational() string {
	if x != nil {
		return x.StatusOperational
	}
	return ""
}

func (x *MontirStatus) GetStatusActivity() string {
	if x != nil {
		return x.StatusActivity
	}
	return ""
}

func (x *MontirStatus) GetDateUpdated() string {
	if x != nil {
		return x.DateUpdated
	}
	return ""
}

var File_montir_proto protoreflect.FileDescriptor

var file_montir_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x74, 0x0a, 0x14, 0x4d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xae, 0x01, 0x0a, 0x0d,
	0x4d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0xaa, 0x04, 0x0a,
	0x0d, 0x4d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x72, 0x6e,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x72,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x74, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x74, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x52, 0x4c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x52, 0x4c, 0x12, 0x29, 0x0a, 0x10, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x6e, 0x74,
	0x69, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x34, 0x0a, 0x0b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f,
	0x6e, 0x74, 0x69, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7c, 0x0a, 0x0c, 0x4d, 0x6f, 0x6e,
	0x74, 0x69, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x6d, 0x0a, 0x0e, 0x4d, 0x6f, 0x6e, 0x74, 0x69,
	0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x0c, 0x4d, 0x6f, 0x6e, 0x74, 0x69,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x32, 0xf9, 0x02, 0x0a, 0x06, 0x4d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x12, 0x48, 0x0a,
	0x11, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x4d, 0x6f, 0x6e, 0x74,
	0x69, 0x72, 0x12, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x69,
	0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d,
	0x6f, 0x6e, 0x74, 0x69, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x4e,
	0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x1a,
	0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x51,
	0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x1a, 0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x69,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x12, 0x4b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a,
	0x1b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x69, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_montir_proto_rawDescOnce sync.Once
	file_montir_proto_rawDescData = file_montir_proto_rawDesc
)

func file_montir_proto_rawDescGZIP() []byte {
	file_montir_proto_rawDescOnce.Do(func() {
		file_montir_proto_rawDescData = protoimpl.X.CompressGZIP(file_montir_proto_rawDescData)
	})
	return file_montir_proto_rawDescData
}

var file_montir_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_montir_proto_goTypes = []interface{}{
	(*MontirResponeMessage)(nil), // 0: model.MontirResponeMessage
	(*MontirAccount)(nil),        // 1: model.MontirAccount
	(*MontirProfile)(nil),        // 2: model.MontirProfile
	(*MontirRating)(nil),         // 3: model.MontirRating
	(*MontirLocation)(nil),       // 4: model.MontirLocation
	(*MontirStatus)(nil),         // 5: model.MontirStatus
}
var file_montir_proto_depIdxs = []int32{
	1,  // 0: model.MontirResponeMessage.result:type_name -> model.MontirAccount
	2,  // 1: model.MontirAccount.profile:type_name -> model.MontirProfile
	5,  // 2: model.MontirProfile.status:type_name -> model.MontirStatus
	3,  // 3: model.MontirProfile.rating_list:type_name -> model.MontirRating
	4,  // 4: model.MontirProfile.location:type_name -> model.MontirLocation
	1,  // 5: model.Montir.RegisterNewMontir:input_type -> model.MontirAccount
	1,  // 6: model.Montir.Login:input_type -> model.MontirAccount
	2,  // 7: model.Montir.UpdateMontirProfileByID:input_type -> model.MontirProfile
	2,  // 8: model.Montir.UpdateMontirProfilePicture:input_type -> model.MontirProfile
	1,  // 9: model.Montir.GetMontirProfileByID:input_type -> model.MontirAccount
	0,  // 10: model.Montir.RegisterNewMontir:output_type -> model.MontirResponeMessage
	1,  // 11: model.Montir.Login:output_type -> model.MontirAccount
	0,  // 12: model.Montir.UpdateMontirProfileByID:output_type -> model.MontirResponeMessage
	0,  // 13: model.Montir.UpdateMontirProfilePicture:output_type -> model.MontirResponeMessage
	0,  // 14: model.Montir.GetMontirProfileByID:output_type -> model.MontirResponeMessage
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_montir_proto_init() }
func file_montir_proto_init() {
	if File_montir_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_montir_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MontirResponeMessage); i {
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
		file_montir_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MontirAccount); i {
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
		file_montir_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MontirProfile); i {
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
		file_montir_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MontirRating); i {
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
		file_montir_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MontirLocation); i {
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
		file_montir_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MontirStatus); i {
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
			RawDescriptor: file_montir_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_montir_proto_goTypes,
		DependencyIndexes: file_montir_proto_depIdxs,
		MessageInfos:      file_montir_proto_msgTypes,
	}.Build()
	File_montir_proto = out.File
	file_montir_proto_rawDesc = nil
	file_montir_proto_goTypes = nil
	file_montir_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MontirClient is the client API for Montir service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MontirClient interface {
	RegisterNewMontir(ctx context.Context, in *MontirAccount, opts ...grpc.CallOption) (*MontirResponeMessage, error)
	Login(ctx context.Context, in *MontirAccount, opts ...grpc.CallOption) (*MontirAccount, error)
	UpdateMontirProfileByID(ctx context.Context, in *MontirProfile, opts ...grpc.CallOption) (*MontirResponeMessage, error)
	UpdateMontirProfilePicture(ctx context.Context, in *MontirProfile, opts ...grpc.CallOption) (*MontirResponeMessage, error)
	GetMontirProfileByID(ctx context.Context, in *MontirAccount, opts ...grpc.CallOption) (*MontirResponeMessage, error)
}

type montirClient struct {
	cc grpc.ClientConnInterface
}

func NewMontirClient(cc grpc.ClientConnInterface) MontirClient {
	return &montirClient{cc}
}

func (c *montirClient) RegisterNewMontir(ctx context.Context, in *MontirAccount, opts ...grpc.CallOption) (*MontirResponeMessage, error) {
	out := new(MontirResponeMessage)
	err := c.cc.Invoke(ctx, "/model.Montir/RegisterNewMontir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *montirClient) Login(ctx context.Context, in *MontirAccount, opts ...grpc.CallOption) (*MontirAccount, error) {
	out := new(MontirAccount)
	err := c.cc.Invoke(ctx, "/model.Montir/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *montirClient) UpdateMontirProfileByID(ctx context.Context, in *MontirProfile, opts ...grpc.CallOption) (*MontirResponeMessage, error) {
	out := new(MontirResponeMessage)
	err := c.cc.Invoke(ctx, "/model.Montir/UpdateMontirProfileByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *montirClient) UpdateMontirProfilePicture(ctx context.Context, in *MontirProfile, opts ...grpc.CallOption) (*MontirResponeMessage, error) {
	out := new(MontirResponeMessage)
	err := c.cc.Invoke(ctx, "/model.Montir/UpdateMontirProfilePicture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *montirClient) GetMontirProfileByID(ctx context.Context, in *MontirAccount, opts ...grpc.CallOption) (*MontirResponeMessage, error) {
	out := new(MontirResponeMessage)
	err := c.cc.Invoke(ctx, "/model.Montir/GetMontirProfileByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MontirServer is the server API for Montir service.
type MontirServer interface {
	RegisterNewMontir(context.Context, *MontirAccount) (*MontirResponeMessage, error)
	Login(context.Context, *MontirAccount) (*MontirAccount, error)
	UpdateMontirProfileByID(context.Context, *MontirProfile) (*MontirResponeMessage, error)
	UpdateMontirProfilePicture(context.Context, *MontirProfile) (*MontirResponeMessage, error)
	GetMontirProfileByID(context.Context, *MontirAccount) (*MontirResponeMessage, error)
}

// UnimplementedMontirServer can be embedded to have forward compatible implementations.
type UnimplementedMontirServer struct {
}

func (*UnimplementedMontirServer) RegisterNewMontir(context.Context, *MontirAccount) (*MontirResponeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNewMontir not implemented")
}
func (*UnimplementedMontirServer) Login(context.Context, *MontirAccount) (*MontirAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedMontirServer) UpdateMontirProfileByID(context.Context, *MontirProfile) (*MontirResponeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMontirProfileByID not implemented")
}
func (*UnimplementedMontirServer) UpdateMontirProfilePicture(context.Context, *MontirProfile) (*MontirResponeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMontirProfilePicture not implemented")
}
func (*UnimplementedMontirServer) GetMontirProfileByID(context.Context, *MontirAccount) (*MontirResponeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMontirProfileByID not implemented")
}

func RegisterMontirServer(s *grpc.Server, srv MontirServer) {
	s.RegisterService(&_Montir_serviceDesc, srv)
}

func _Montir_RegisterNewMontir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MontirAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MontirServer).RegisterNewMontir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Montir/RegisterNewMontir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MontirServer).RegisterNewMontir(ctx, req.(*MontirAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Montir_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MontirAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MontirServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Montir/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MontirServer).Login(ctx, req.(*MontirAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Montir_UpdateMontirProfileByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MontirProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MontirServer).UpdateMontirProfileByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Montir/UpdateMontirProfileByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MontirServer).UpdateMontirProfileByID(ctx, req.(*MontirProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _Montir_UpdateMontirProfilePicture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MontirProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MontirServer).UpdateMontirProfilePicture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Montir/UpdateMontirProfilePicture",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MontirServer).UpdateMontirProfilePicture(ctx, req.(*MontirProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _Montir_GetMontirProfileByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MontirAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MontirServer).GetMontirProfileByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Montir/GetMontirProfileByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MontirServer).GetMontirProfileByID(ctx, req.(*MontirAccount))
	}
	return interceptor(ctx, in, info, handler)
}

var _Montir_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Montir",
	HandlerType: (*MontirServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNewMontir",
			Handler:    _Montir_RegisterNewMontir_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Montir_Login_Handler,
		},
		{
			MethodName: "UpdateMontirProfileByID",
			Handler:    _Montir_UpdateMontirProfileByID_Handler,
		},
		{
			MethodName: "UpdateMontirProfilePicture",
			Handler:    _Montir_UpdateMontirProfilePicture_Handler,
		},
		{
			MethodName: "GetMontirProfileByID",
			Handler:    _Montir_GetMontirProfileByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "montir.proto",
}
