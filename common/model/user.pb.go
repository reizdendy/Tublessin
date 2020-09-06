// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: user.proto

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

type UserResponeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response  string         `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Code      string         `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Result    *UserAccount   `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	TotalUser string         `protobuf:"bytes,4,opt,name=total_user,json=totalUser,proto3" json:"total_user,omitempty"`
	List      []*UserAccount `protobuf:"bytes,5,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *UserResponeMessage) Reset() {
	*x = UserResponeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponeMessage) ProtoMessage() {}

func (x *UserResponeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponeMessage.ProtoReflect.Descriptor instead.
func (*UserResponeMessage) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserResponeMessage) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *UserResponeMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UserResponeMessage) GetResult() *UserAccount {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *UserResponeMessage) GetTotalUser() string {
	if x != nil {
		return x.TotalUser
	}
	return ""
}

func (x *UserResponeMessage) GetList() []*UserAccount {
	if x != nil {
		return x.List
	}
	return nil
}

type UserAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string       `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password      string       `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	StatusAccount string       `protobuf:"bytes,4,opt,name=status_account,json=statusAccount,proto3" json:"status_account,omitempty"`
	Profile       *UserProfile `protobuf:"bytes,5,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *UserAccount) Reset() {
	*x = UserAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAccount) ProtoMessage() {}

func (x *UserAccount) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAccount.ProtoReflect.Descriptor instead.
func (*UserAccount) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserAccount) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserAccount) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserAccount) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserAccount) GetStatusAccount() string {
	if x != nil {
		return x.StatusAccount
	}
	return ""
}

func (x *UserAccount) GetProfile() *UserProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type UserProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname   string        `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname    string        `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Gender      string        `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	PhoneNumber string        `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Email       string        `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	ImageURL    string        `protobuf:"bytes,7,opt,name=imageURL,proto3" json:"imageURL,omitempty"`
	DateUpdated string        `protobuf:"bytes,8,opt,name=date_updated,json=dateUpdated,proto3" json:"date_updated,omitempty"`
	DateCreated string        `protobuf:"bytes,9,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	Location    *UserLocation `protobuf:"bytes,10,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserProfile) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserProfile) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *UserProfile) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *UserProfile) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserProfile) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserProfile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserProfile) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *UserProfile) GetDateUpdated() string {
	if x != nil {
		return x.DateUpdated
	}
	return ""
}

func (x *UserProfile) GetDateCreated() string {
	if x != nil {
		return x.DateCreated
	}
	return ""
}

func (x *UserProfile) GetLocation() *UserLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

type UserLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude    float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude   float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	DateUpdated string  `protobuf:"bytes,3,opt,name=date_updated,json=dateUpdated,proto3" json:"date_updated,omitempty"`
}

func (x *UserLocation) Reset() {
	*x = UserLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLocation) ProtoMessage() {}

func (x *UserLocation) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLocation.ProtoReflect.Descriptor instead.
func (*UserLocation) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserLocation) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *UserLocation) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *UserLocation) GetDateUpdated() string {
	if x != nil {
		return x.DateUpdated
	}
	return ""
}

type UserPagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Page    string `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit   string `protobuf:"bytes,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Status  string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	OrderBy string `protobuf:"bytes,5,opt,name=orderBy,proto3" json:"orderBy,omitempty"`
	Sort    string `protobuf:"bytes,6,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *UserPagination) Reset() {
	*x = UserPagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPagination) ProtoMessage() {}

func (x *UserPagination) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPagination.ProtoReflect.Descriptor instead.
func (*UserPagination) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserPagination) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *UserPagination) GetPage() string {
	if x != nil {
		return x.Page
	}
	return ""
}

func (x *UserPagination) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

func (x *UserPagination) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UserPagination) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *UserPagination) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x22, 0xb7, 0x01, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xaa, 0x01,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0xbb, 0x02, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52,
	0x4c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52,
	0x4c, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6b, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x9a, 0x01, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x32, 0xae, 0x04, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x0f, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12,
	0x31, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x12, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x12, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x1a,
	0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x18,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x19, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x1a, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x12, 0x45, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x19, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_proto_goTypes = []interface{}{
	(*UserResponeMessage)(nil), // 0: model.UserResponeMessage
	(*UserAccount)(nil),        // 1: model.UserAccount
	(*UserProfile)(nil),        // 2: model.UserProfile
	(*UserLocation)(nil),       // 3: model.UserLocation
	(*UserPagination)(nil),     // 4: model.UserPagination
}
var file_user_proto_depIdxs = []int32{
	1,  // 0: model.UserResponeMessage.result:type_name -> model.UserAccount
	1,  // 1: model.UserResponeMessage.list:type_name -> model.UserAccount
	2,  // 2: model.UserAccount.profile:type_name -> model.UserProfile
	3,  // 3: model.UserProfile.location:type_name -> model.UserLocation
	1,  // 4: model.User.RegisterNewUser:input_type -> model.UserAccount
	1,  // 5: model.User.Login:input_type -> model.UserAccount
	2,  // 6: model.User.UpdateUserProfileById:input_type -> model.UserProfile
	2,  // 7: model.User.UpdateUserProfilePicture:input_type -> model.UserProfile
	2,  // 8: model.User.UpdateUserLocation:input_type -> model.UserProfile
	1,  // 9: model.User.GetUserProfileById:input_type -> model.UserAccount
	1,  // 10: model.User.DeleteUserByID:input_type -> model.UserAccount
	4,  // 11: model.User.GetAllUserSummary:input_type -> model.UserPagination
	0,  // 12: model.User.RegisterNewUser:output_type -> model.UserResponeMessage
	1,  // 13: model.User.Login:output_type -> model.UserAccount
	0,  // 14: model.User.UpdateUserProfileById:output_type -> model.UserResponeMessage
	0,  // 15: model.User.UpdateUserProfilePicture:output_type -> model.UserResponeMessage
	0,  // 16: model.User.UpdateUserLocation:output_type -> model.UserResponeMessage
	0,  // 17: model.User.GetUserProfileById:output_type -> model.UserResponeMessage
	0,  // 18: model.User.DeleteUserByID:output_type -> model.UserResponeMessage
	0,  // 19: model.User.GetAllUserSummary:output_type -> model.UserResponeMessage
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponeMessage); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAccount); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProfile); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLocation); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPagination); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	RegisterNewUser(ctx context.Context, in *UserAccount, opts ...grpc.CallOption) (*UserResponeMessage, error)
	Login(ctx context.Context, in *UserAccount, opts ...grpc.CallOption) (*UserAccount, error)
	UpdateUserProfileById(ctx context.Context, in *UserProfile, opts ...grpc.CallOption) (*UserResponeMessage, error)
	UpdateUserProfilePicture(ctx context.Context, in *UserProfile, opts ...grpc.CallOption) (*UserResponeMessage, error)
	UpdateUserLocation(ctx context.Context, in *UserProfile, opts ...grpc.CallOption) (*UserResponeMessage, error)
	GetUserProfileById(ctx context.Context, in *UserAccount, opts ...grpc.CallOption) (*UserResponeMessage, error)
	DeleteUserByID(ctx context.Context, in *UserAccount, opts ...grpc.CallOption) (*UserResponeMessage, error)
	GetAllUserSummary(ctx context.Context, in *UserPagination, opts ...grpc.CallOption) (*UserResponeMessage, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) RegisterNewUser(ctx context.Context, in *UserAccount, opts ...grpc.CallOption) (*UserResponeMessage, error) {
	out := new(UserResponeMessage)
	err := c.cc.Invoke(ctx, "/model.User/RegisterNewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Login(ctx context.Context, in *UserAccount, opts ...grpc.CallOption) (*UserAccount, error) {
	out := new(UserAccount)
	err := c.cc.Invoke(ctx, "/model.User/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserProfileById(ctx context.Context, in *UserProfile, opts ...grpc.CallOption) (*UserResponeMessage, error) {
	out := new(UserResponeMessage)
	err := c.cc.Invoke(ctx, "/model.User/UpdateUserProfileById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserProfilePicture(ctx context.Context, in *UserProfile, opts ...grpc.CallOption) (*UserResponeMessage, error) {
	out := new(UserResponeMessage)
	err := c.cc.Invoke(ctx, "/model.User/UpdateUserProfilePicture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserLocation(ctx context.Context, in *UserProfile, opts ...grpc.CallOption) (*UserResponeMessage, error) {
	out := new(UserResponeMessage)
	err := c.cc.Invoke(ctx, "/model.User/UpdateUserLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserProfileById(ctx context.Context, in *UserAccount, opts ...grpc.CallOption) (*UserResponeMessage, error) {
	out := new(UserResponeMessage)
	err := c.cc.Invoke(ctx, "/model.User/GetUserProfileById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUserByID(ctx context.Context, in *UserAccount, opts ...grpc.CallOption) (*UserResponeMessage, error) {
	out := new(UserResponeMessage)
	err := c.cc.Invoke(ctx, "/model.User/DeleteUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetAllUserSummary(ctx context.Context, in *UserPagination, opts ...grpc.CallOption) (*UserResponeMessage, error) {
	out := new(UserResponeMessage)
	err := c.cc.Invoke(ctx, "/model.User/GetAllUserSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	RegisterNewUser(context.Context, *UserAccount) (*UserResponeMessage, error)
	Login(context.Context, *UserAccount) (*UserAccount, error)
	UpdateUserProfileById(context.Context, *UserProfile) (*UserResponeMessage, error)
	UpdateUserProfilePicture(context.Context, *UserProfile) (*UserResponeMessage, error)
	UpdateUserLocation(context.Context, *UserProfile) (*UserResponeMessage, error)
	GetUserProfileById(context.Context, *UserAccount) (*UserResponeMessage, error)
	DeleteUserByID(context.Context, *UserAccount) (*UserResponeMessage, error)
	GetAllUserSummary(context.Context, *UserPagination) (*UserResponeMessage, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) RegisterNewUser(context.Context, *UserAccount) (*UserResponeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNewUser not implemented")
}
func (*UnimplementedUserServer) Login(context.Context, *UserAccount) (*UserAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedUserServer) UpdateUserProfileById(context.Context, *UserProfile) (*UserResponeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfileById not implemented")
}
func (*UnimplementedUserServer) UpdateUserProfilePicture(context.Context, *UserProfile) (*UserResponeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfilePicture not implemented")
}
func (*UnimplementedUserServer) UpdateUserLocation(context.Context, *UserProfile) (*UserResponeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserLocation not implemented")
}
func (*UnimplementedUserServer) GetUserProfileById(context.Context, *UserAccount) (*UserResponeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfileById not implemented")
}
func (*UnimplementedUserServer) DeleteUserByID(context.Context, *UserAccount) (*UserResponeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserByID not implemented")
}
func (*UnimplementedUserServer) GetAllUserSummary(context.Context, *UserPagination) (*UserResponeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUserSummary not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_RegisterNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RegisterNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.User/RegisterNewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RegisterNewUser(ctx, req.(*UserAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*UserAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserProfileById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserProfileById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.User/UpdateUserProfileById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserProfileById(ctx, req.(*UserProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserProfilePicture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserProfilePicture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.User/UpdateUserProfilePicture",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserProfilePicture(ctx, req.(*UserProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.User/UpdateUserLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserLocation(ctx, req.(*UserProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserProfileById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserProfileById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.User/GetUserProfileById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserProfileById(ctx, req.(*UserAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.User/DeleteUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUserByID(ctx, req.(*UserAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetAllUserSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPagination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetAllUserSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.User/GetAllUserSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetAllUserSummary(ctx, req.(*UserPagination))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNewUser",
			Handler:    _User_RegisterNewUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "UpdateUserProfileById",
			Handler:    _User_UpdateUserProfileById_Handler,
		},
		{
			MethodName: "UpdateUserProfilePicture",
			Handler:    _User_UpdateUserProfilePicture_Handler,
		},
		{
			MethodName: "UpdateUserLocation",
			Handler:    _User_UpdateUserLocation_Handler,
		},
		{
			MethodName: "GetUserProfileById",
			Handler:    _User_GetUserProfileById_Handler,
		},
		{
			MethodName: "DeleteUserByID",
			Handler:    _User_DeleteUserByID_Handler,
		},
		{
			MethodName: "GetAllUserSummary",
			Handler:    _User_GetAllUserSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
