// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: match.proto

package pb

import (
	any1 "github.com/golang/protobuf/ptypes/any"
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

type PhotoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    int32  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	ImageData []byte `protobuf:"bytes,2,opt,name=imageData,proto3" json:"imageData,omitempty"`
	LastChunk bool   `protobuf:"varint,3,opt,name=lastChunk,proto3" json:"lastChunk,omitempty"`
}

func (x *PhotoRequest) Reset() {
	*x = PhotoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhotoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhotoRequest) ProtoMessage() {}

func (x *PhotoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhotoRequest.ProtoReflect.Descriptor instead.
func (*PhotoRequest) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{0}
}

func (x *PhotoRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *PhotoRequest) GetImageData() []byte {
	if x != nil {
		return x.ImageData
	}
	return nil
}

func (x *PhotoRequest) GetLastChunk() bool {
	if x != nil {
		return x.LastChunk
	}
	return false
}

type UserPrefrencesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int32  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Height        string `protobuf:"bytes,2,opt,name=height,proto3" json:"height,omitempty"`
	MaritalStatus string `protobuf:"bytes,3,opt,name=maritalStatus,proto3" json:"maritalStatus,omitempty"`
	Faith         string `protobuf:"bytes,4,opt,name=faith,proto3" json:"faith,omitempty"`
	MotherTounge  string `protobuf:"bytes,5,opt,name=motherTounge,proto3" json:"motherTounge,omitempty"`
	SmokeStatus   string `protobuf:"bytes,6,opt,name=smokeStatus,proto3" json:"smokeStatus,omitempty"`
	AlcoholStatus string `protobuf:"bytes,7,opt,name=alcoholStatus,proto3" json:"alcoholStatus,omitempty"`
	SettleStatus  string `protobuf:"bytes,8,opt,name=settleStatus,proto3" json:"settleStatus,omitempty"`
	Hobbies       string `protobuf:"bytes,9,opt,name=hobbies,proto3" json:"hobbies,omitempty"`
	TeaPerson     string `protobuf:"bytes,10,opt,name=teaPerson,proto3" json:"teaPerson,omitempty"`
	LoveLanguage  string `protobuf:"bytes,11,opt,name=loveLanguage,proto3" json:"loveLanguage,omitempty"`
}

func (x *UserPrefrencesRequest) Reset() {
	*x = UserPrefrencesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPrefrencesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPrefrencesRequest) ProtoMessage() {}

func (x *UserPrefrencesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPrefrencesRequest.ProtoReflect.Descriptor instead.
func (*UserPrefrencesRequest) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{1}
}

func (x *UserPrefrencesRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserPrefrencesRequest) GetHeight() string {
	if x != nil {
		return x.Height
	}
	return ""
}

func (x *UserPrefrencesRequest) GetMaritalStatus() string {
	if x != nil {
		return x.MaritalStatus
	}
	return ""
}

func (x *UserPrefrencesRequest) GetFaith() string {
	if x != nil {
		return x.Faith
	}
	return ""
}

func (x *UserPrefrencesRequest) GetMotherTounge() string {
	if x != nil {
		return x.MotherTounge
	}
	return ""
}

func (x *UserPrefrencesRequest) GetSmokeStatus() string {
	if x != nil {
		return x.SmokeStatus
	}
	return ""
}

func (x *UserPrefrencesRequest) GetAlcoholStatus() string {
	if x != nil {
		return x.AlcoholStatus
	}
	return ""
}

func (x *UserPrefrencesRequest) GetSettleStatus() string {
	if x != nil {
		return x.SettleStatus
	}
	return ""
}

func (x *UserPrefrencesRequest) GetHobbies() string {
	if x != nil {
		return x.Hobbies
	}
	return ""
}

func (x *UserPrefrencesRequest) GetTeaPerson() string {
	if x != nil {
		return x.TeaPerson
	}
	return ""
}

func (x *UserPrefrencesRequest) GetLoveLanguage() string {
	if x != nil {
		return x.LoveLanguage
	}
	return ""
}

type UserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int32 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *UserIdRequest) Reset() {
	*x = UserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdRequest) ProtoMessage() {}

func (x *UserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdRequest.ProtoReflect.Descriptor instead.
func (*UserIdRequest) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{2}
}

func (x *UserIdRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type MatchedUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchedUsers []*MatchedUsers `protobuf:"bytes,1,rep,name=matchedUsers,proto3" json:"matchedUsers,omitempty"`
}

func (x *MatchedUsersResponse) Reset() {
	*x = MatchedUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchedUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchedUsersResponse) ProtoMessage() {}

func (x *MatchedUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchedUsersResponse.ProtoReflect.Descriptor instead.
func (*MatchedUsersResponse) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{3}
}

func (x *MatchedUsersResponse) GetMatchedUsers() []*MatchedUsers {
	if x != nil {
		return x.MatchedUsers
	}
	return nil
}

type MatchedUsers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     int32     `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Fullname   string    `protobuf:"bytes,2,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Age        string    `protobuf:"bytes,3,opt,name=age,proto3" json:"age,omitempty"`
	Location   string    `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	UserImages []*Images `protobuf:"bytes,5,rep,name=userImages,proto3" json:"userImages,omitempty"`
}

func (x *MatchedUsers) Reset() {
	*x = MatchedUsers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchedUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchedUsers) ProtoMessage() {}

func (x *MatchedUsers) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchedUsers.ProtoReflect.Descriptor instead.
func (*MatchedUsers) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{4}
}

func (x *MatchedUsers) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *MatchedUsers) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *MatchedUsers) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

func (x *MatchedUsers) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *MatchedUsers) GetUserImages() []*Images {
	if x != nil {
		return x.UserImages
	}
	return nil
}

type Images struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageId string `protobuf:"bytes,1,opt,name=imageId,proto3" json:"imageId,omitempty"`
}

func (x *Images) Reset() {
	*x = Images{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Images) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Images) ProtoMessage() {}

func (x *Images) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Images.ProtoReflect.Descriptor instead.
func (*Images) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{5}
}

func (x *Images) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

type MatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Error   *any1.Any `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"` // Use google.protobuf.Any for interface{} in Go
	Data    *any1.Any `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`   // Use google.protobuf.Any for interface{} in Go
}

func (x *MatchResponse) Reset() {
	*x = MatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchResponse) ProtoMessage() {}

func (x *MatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchResponse.ProtoReflect.Descriptor instead.
func (*MatchResponse) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{6}
}

func (x *MatchResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MatchResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *MatchResponse) GetError() *any1.Any {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *MatchResponse) GetData() *any1.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_match_proto protoreflect.FileDescriptor

var file_match_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x0c, 0x50, 0x68, 0x6f, 0x74,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c,
	0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0xef, 0x02, 0x0a,
	0x15, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x61, 0x72, 0x69, 0x74, 0x61,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d,
	0x61, 0x72, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x61, 0x69, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x61, 0x69,
	0x74, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x54, 0x6f, 0x75, 0x6e,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x54, 0x6f, 0x75, 0x6e, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x6d, 0x6f,
	0x6b, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x6c, 0x63, 0x6f,
	0x68, 0x6f, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x61, 0x6c, 0x63, 0x6f, 0x68, 0x6f, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x62, 0x62, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x65, 0x61, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x65, 0x61, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x6f,
	0x76, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6c, 0x6f, 0x76, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x27,
	0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x49, 0x0a, 0x14, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x22, 0x99, 0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x22,
	0x0a, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x0d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x28,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xbd, 0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x0c, 0x55, 0x70, 0x6c,
	0x61, 0x6f, 0x64, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x0d, 0x2e, 0x50, 0x68, 0x6f, 0x74,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3e, 0x0a, 0x12,
	0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x12, 0x16, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x0e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_match_proto_rawDescOnce sync.Once
	file_match_proto_rawDescData = file_match_proto_rawDesc
)

func file_match_proto_rawDescGZIP() []byte {
	file_match_proto_rawDescOnce.Do(func() {
		file_match_proto_rawDescData = protoimpl.X.CompressGZIP(file_match_proto_rawDescData)
	})
	return file_match_proto_rawDescData
}

var file_match_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_match_proto_goTypes = []interface{}{
	(*PhotoRequest)(nil),          // 0: PhotoRequest
	(*UserPrefrencesRequest)(nil), // 1: UserPrefrencesRequest
	(*UserIdRequest)(nil),         // 2: UserIdRequest
	(*MatchedUsersResponse)(nil),  // 3: MatchedUsersResponse
	(*MatchedUsers)(nil),          // 4: MatchedUsers
	(*Images)(nil),                // 5: Images
	(*MatchResponse)(nil),         // 6: MatchResponse
	(*any1.Any)(nil),              // 7: google.protobuf.Any
}
var file_match_proto_depIdxs = []int32{
	4, // 0: MatchedUsersResponse.matchedUsers:type_name -> MatchedUsers
	5, // 1: MatchedUsers.userImages:type_name -> Images
	7, // 2: MatchResponse.error:type_name -> google.protobuf.Any
	7, // 3: MatchResponse.data:type_name -> google.protobuf.Any
	0, // 4: MatchService.UplaodPhotos:input_type -> PhotoRequest
	1, // 5: MatchService.SaveUserPrefrences:input_type -> UserPrefrencesRequest
	2, // 6: MatchService.GetMatchedUsers:input_type -> UserIdRequest
	6, // 7: MatchService.UplaodPhotos:output_type -> MatchResponse
	6, // 8: MatchService.SaveUserPrefrences:output_type -> MatchResponse
	3, // 9: MatchService.GetMatchedUsers:output_type -> MatchedUsersResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_match_proto_init() }
func file_match_proto_init() {
	if File_match_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_match_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhotoRequest); i {
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
		file_match_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPrefrencesRequest); i {
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
		file_match_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdRequest); i {
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
		file_match_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchedUsersResponse); i {
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
		file_match_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchedUsers); i {
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
		file_match_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Images); i {
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
		file_match_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchResponse); i {
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
			RawDescriptor: file_match_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_match_proto_goTypes,
		DependencyIndexes: file_match_proto_depIdxs,
		MessageInfos:      file_match_proto_msgTypes,
	}.Build()
	File_match_proto = out.File
	file_match_proto_rawDesc = nil
	file_match_proto_goTypes = nil
	file_match_proto_depIdxs = nil
}
