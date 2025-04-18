// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: messages.proto

package protos

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

type ActionType int32

const (
	ActionType_ACTION_TYPE_UNSPECIFIED  ActionType = 0
	ActionType_ACTION_TYPE_GET          ActionType = 1
	ActionType_ACTION_TYPE_LIST         ActionType = 2
	ActionType_ACTION_TYPE_CREATE       ActionType = 3
	ActionType_ACTION_TYPE_UPDATE       ActionType = 4
	ActionType_ACTION_TYPE_DELETE       ActionType = 5
	ActionType_ACTION_TYPE_NOTIFICATION ActionType = 6
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0: "ACTION_TYPE_UNSPECIFIED",
		1: "ACTION_TYPE_GET",
		2: "ACTION_TYPE_LIST",
		3: "ACTION_TYPE_CREATE",
		4: "ACTION_TYPE_UPDATE",
		5: "ACTION_TYPE_DELETE",
		6: "ACTION_TYPE_NOTIFICATION",
	}
	ActionType_value = map[string]int32{
		"ACTION_TYPE_UNSPECIFIED":  0,
		"ACTION_TYPE_GET":          1,
		"ACTION_TYPE_LIST":         2,
		"ACTION_TYPE_CREATE":       3,
		"ACTION_TYPE_UPDATE":       4,
		"ACTION_TYPE_DELETE":       5,
		"ACTION_TYPE_NOTIFICATION": 6,
	}
)

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}

func (x ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_proto_enumTypes[0].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_messages_proto_enumTypes[0]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

// ----- SHARED -----
type InputMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Meta          *Meta                  `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Input         *Input                 `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InputMessage) Reset() {
	*x = InputMessage{}
	mi := &file_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InputMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputMessage) ProtoMessage() {}

func (x *InputMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputMessage.ProtoReflect.Descriptor instead.
func (*InputMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *InputMessage) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *InputMessage) GetInput() *Input {
	if x != nil {
		return x.Input
	}
	return nil
}

type OutputMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Meta          *Meta                  `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Output        *Output                `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OutputMessage) Reset() {
	*x = OutputMessage{}
	mi := &file_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OutputMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputMessage) ProtoMessage() {}

func (x *OutputMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputMessage.ProtoReflect.Descriptor instead.
func (*OutputMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

func (x *OutputMessage) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *OutputMessage) GetOutput() *Output {
	if x != nil {
		return x.Output
	}
	return nil
}

type Meta struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Action        ActionType             `protobuf:"varint,1,opt,name=action,proto3,enum=v1.messages.ActionType" json:"action,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Meta) Reset() {
	*x = Meta{}
	mi := &file_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2}
}

func (x *Meta) GetAction() ActionType {
	if x != nil {
		return x.Action
	}
	return ActionType_ACTION_TYPE_UNSPECIFIED
}

// ----- INPUT -----
type Input struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Data:
	//
	//	*Input_UpdateJwtToken
	//	*Input_AppserverListing
	//	*Input_AppserverDetails
	//	*Input_CreateAppserver
	//	*Input_DeleteAppserver
	//	*Input_CreateChannel
	//	*Input_ChannelListing
	//	*Input_JoinAppserver
	//	*Input_CreateAppserverRole
	//	*Input_AppserverRolesListing
	//	*Input_AppserverUserListing
	Data          isInput_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Input) Reset() {
	*x = Input{}
	mi := &file_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{3}
}

func (x *Input) GetData() isInput_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Input) GetUpdateJwtToken() *UpdateJwtToken {
	if x != nil {
		if x, ok := x.Data.(*Input_UpdateJwtToken); ok {
			return x.UpdateJwtToken
		}
	}
	return nil
}

func (x *Input) GetAppserverListing() *AppserverListingRequest {
	if x != nil {
		if x, ok := x.Data.(*Input_AppserverListing); ok {
			return x.AppserverListing
		}
	}
	return nil
}

func (x *Input) GetAppserverDetails() *GetByIdAppserverRequest {
	if x != nil {
		if x, ok := x.Data.(*Input_AppserverDetails); ok {
			return x.AppserverDetails
		}
	}
	return nil
}

func (x *Input) GetCreateAppserver() *CreateAppserverRequest {
	if x != nil {
		if x, ok := x.Data.(*Input_CreateAppserver); ok {
			return x.CreateAppserver
		}
	}
	return nil
}

func (x *Input) GetDeleteAppserver() *DeleteAppserverRequest {
	if x != nil {
		if x, ok := x.Data.(*Input_DeleteAppserver); ok {
			return x.DeleteAppserver
		}
	}
	return nil
}

func (x *Input) GetCreateChannel() *CreateChannelRequest {
	if x != nil {
		if x, ok := x.Data.(*Input_CreateChannel); ok {
			return x.CreateChannel
		}
	}
	return nil
}

func (x *Input) GetChannelListing() *ListChannelsRequest {
	if x != nil {
		if x, ok := x.Data.(*Input_ChannelListing); ok {
			return x.ChannelListing
		}
	}
	return nil
}

func (x *Input) GetJoinAppserver() *CreateAppserverSubRequest {
	if x != nil {
		if x, ok := x.Data.(*Input_JoinAppserver); ok {
			return x.JoinAppserver
		}
	}
	return nil
}

func (x *Input) GetCreateAppserverRole() *CreateAppserverRoleRequest {
	if x != nil {
		if x, ok := x.Data.(*Input_CreateAppserverRole); ok {
			return x.CreateAppserverRole
		}
	}
	return nil
}

func (x *Input) GetAppserverRolesListing() *GetAllAppserverRolesRequest {
	if x != nil {
		if x, ok := x.Data.(*Input_AppserverRolesListing); ok {
			return x.AppserverRolesListing
		}
	}
	return nil
}

func (x *Input) GetAppserverUserListing() *GetAllUsersAppserverSubsRequest {
	if x != nil {
		if x, ok := x.Data.(*Input_AppserverUserListing); ok {
			return x.AppserverUserListing
		}
	}
	return nil
}

type isInput_Data interface {
	isInput_Data()
}

type Input_UpdateJwtToken struct {
	UpdateJwtToken *UpdateJwtToken `protobuf:"bytes,1,opt,name=update_jwt_token,json=updateJwtToken,proto3,oneof"`
}

type Input_AppserverListing struct {
	AppserverListing *AppserverListingRequest `protobuf:"bytes,2,opt,name=appserver_listing,json=appserverListing,proto3,oneof"`
}

type Input_AppserverDetails struct {
	AppserverDetails *GetByIdAppserverRequest `protobuf:"bytes,3,opt,name=appserver_details,json=appserverDetails,proto3,oneof"`
}

type Input_CreateAppserver struct {
	CreateAppserver *CreateAppserverRequest `protobuf:"bytes,4,opt,name=create_appserver,json=createAppserver,proto3,oneof"`
}

type Input_DeleteAppserver struct {
	DeleteAppserver *DeleteAppserverRequest `protobuf:"bytes,5,opt,name=delete_appserver,json=deleteAppserver,proto3,oneof"`
}

type Input_CreateChannel struct {
	CreateChannel *CreateChannelRequest `protobuf:"bytes,6,opt,name=create_channel,json=createChannel,proto3,oneof"`
}

type Input_ChannelListing struct {
	ChannelListing *ListChannelsRequest `protobuf:"bytes,7,opt,name=channel_listing,json=channelListing,proto3,oneof"`
}

type Input_JoinAppserver struct {
	JoinAppserver *CreateAppserverSubRequest `protobuf:"bytes,8,opt,name=join_appserver,json=joinAppserver,proto3,oneof"`
}

type Input_CreateAppserverRole struct {
	CreateAppserverRole *CreateAppserverRoleRequest `protobuf:"bytes,9,opt,name=create_appserver_role,json=createAppserverRole,proto3,oneof"`
}

type Input_AppserverRolesListing struct {
	AppserverRolesListing *GetAllAppserverRolesRequest `protobuf:"bytes,10,opt,name=appserver_roles_listing,json=appserverRolesListing,proto3,oneof"`
}

type Input_AppserverUserListing struct {
	AppserverUserListing *GetAllUsersAppserverSubsRequest `protobuf:"bytes,11,opt,name=appserver_user_listing,json=appserverUserListing,proto3,oneof"`
}

func (*Input_UpdateJwtToken) isInput_Data() {}

func (*Input_AppserverListing) isInput_Data() {}

func (*Input_AppserverDetails) isInput_Data() {}

func (*Input_CreateAppserver) isInput_Data() {}

func (*Input_DeleteAppserver) isInput_Data() {}

func (*Input_CreateChannel) isInput_Data() {}

func (*Input_ChannelListing) isInput_Data() {}

func (*Input_JoinAppserver) isInput_Data() {}

func (*Input_CreateAppserverRole) isInput_Data() {}

func (*Input_AppserverRolesListing) isInput_Data() {}

func (*Input_AppserverUserListing) isInput_Data() {}

type UpdateJwtToken struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Access        string                 `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateJwtToken) Reset() {
	*x = UpdateJwtToken{}
	mi := &file_messages_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateJwtToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJwtToken) ProtoMessage() {}

func (x *UpdateJwtToken) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJwtToken.ProtoReflect.Descriptor instead.
func (*UpdateJwtToken) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateJwtToken) GetAccess() string {
	if x != nil {
		return x.Access
	}
	return ""
}

// ----- OUTPUT -----
type Output struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Data:
	//
	//	*Output_AppserverListing
	//	*Output_AppserverDetails
	//	*Output_ChannelListing
	//	*Output_UpdateAddAppserver
	//	*Output_AppserverRolesListing
	//	*Output_AppserverUserListing
	//	*Output_UpdateRemoveAppserver
	//	*Output_UpdateNewChannel
	//	*Output_UpdateDeleteChannel
	Data          isOutput_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Output) Reset() {
	*x = Output{}
	mi := &file_messages_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{5}
}

func (x *Output) GetData() isOutput_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Output) GetAppserverListing() *GetUserAppserverSubsResponse {
	if x != nil {
		if x, ok := x.Data.(*Output_AppserverListing); ok {
			return x.AppserverListing
		}
	}
	return nil
}

func (x *Output) GetAppserverDetails() *GetByIdAppserverResponse {
	if x != nil {
		if x, ok := x.Data.(*Output_AppserverDetails); ok {
			return x.AppserverDetails
		}
	}
	return nil
}

func (x *Output) GetChannelListing() *ListChannelsResponse {
	if x != nil {
		if x, ok := x.Data.(*Output_ChannelListing); ok {
			return x.ChannelListing
		}
	}
	return nil
}

func (x *Output) GetUpdateAddAppserver() *Appserver {
	if x != nil {
		if x, ok := x.Data.(*Output_UpdateAddAppserver); ok {
			return x.UpdateAddAppserver
		}
	}
	return nil
}

func (x *Output) GetAppserverRolesListing() *GetAllAppserverRolesResponse {
	if x != nil {
		if x, ok := x.Data.(*Output_AppserverRolesListing); ok {
			return x.AppserverRolesListing
		}
	}
	return nil
}

func (x *Output) GetAppserverUserListing() *GetAllUsersAppserverSubsResponse {
	if x != nil {
		if x, ok := x.Data.(*Output_AppserverUserListing); ok {
			return x.AppserverUserListing
		}
	}
	return nil
}

func (x *Output) GetUpdateRemoveAppserver() string {
	if x != nil {
		if x, ok := x.Data.(*Output_UpdateRemoveAppserver); ok {
			return x.UpdateRemoveAppserver
		}
	}
	return ""
}

func (x *Output) GetUpdateNewChannel() *Channel {
	if x != nil {
		if x, ok := x.Data.(*Output_UpdateNewChannel); ok {
			return x.UpdateNewChannel
		}
	}
	return nil
}

func (x *Output) GetUpdateDeleteChannel() *Channel {
	if x != nil {
		if x, ok := x.Data.(*Output_UpdateDeleteChannel); ok {
			return x.UpdateDeleteChannel
		}
	}
	return nil
}

type isOutput_Data interface {
	isOutput_Data()
}

type Output_AppserverListing struct {
	AppserverListing *GetUserAppserverSubsResponse `protobuf:"bytes,1,opt,name=appserver_listing,json=appserverListing,proto3,oneof"`
}

type Output_AppserverDetails struct {
	AppserverDetails *GetByIdAppserverResponse `protobuf:"bytes,2,opt,name=appserver_details,json=appserverDetails,proto3,oneof"`
}

type Output_ChannelListing struct {
	ChannelListing *ListChannelsResponse `protobuf:"bytes,3,opt,name=channel_listing,json=channelListing,proto3,oneof"`
}

type Output_UpdateAddAppserver struct {
	UpdateAddAppserver *Appserver `protobuf:"bytes,4,opt,name=update_add_appserver,json=updateAddAppserver,proto3,oneof"`
}

type Output_AppserverRolesListing struct {
	AppserverRolesListing *GetAllAppserverRolesResponse `protobuf:"bytes,5,opt,name=appserver_roles_listing,json=appserverRolesListing,proto3,oneof"`
}

type Output_AppserverUserListing struct {
	AppserverUserListing *GetAllUsersAppserverSubsResponse `protobuf:"bytes,6,opt,name=appserver_user_listing,json=appserverUserListing,proto3,oneof"`
}

type Output_UpdateRemoveAppserver struct {
	UpdateRemoveAppserver string `protobuf:"bytes,7,opt,name=update_remove_appserver,json=updateRemoveAppserver,proto3,oneof"`
}

type Output_UpdateNewChannel struct {
	UpdateNewChannel *Channel `protobuf:"bytes,8,opt,name=update_new_channel,json=updateNewChannel,proto3,oneof"`
}

type Output_UpdateDeleteChannel struct {
	UpdateDeleteChannel *Channel `protobuf:"bytes,9,opt,name=update_delete_channel,json=updateDeleteChannel,proto3,oneof"`
}

func (*Output_AppserverListing) isOutput_Data() {}

func (*Output_AppserverDetails) isOutput_Data() {}

func (*Output_ChannelListing) isOutput_Data() {}

func (*Output_UpdateAddAppserver) isOutput_Data() {}

func (*Output_AppserverRolesListing) isOutput_Data() {}

func (*Output_AppserverUserListing) isOutput_Data() {}

func (*Output_UpdateRemoveAppserver) isOutput_Data() {}

func (*Output_UpdateNewChannel) isOutput_Data() {}

func (*Output_UpdateDeleteChannel) isOutput_Data() {}

// LISTING
type AppserverListingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AppserverListingRequest) Reset() {
	*x = AppserverListingRequest{}
	mi := &file_messages_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppserverListingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppserverListingRequest) ProtoMessage() {}

func (x *AppserverListingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppserverListingRequest.ProtoReflect.Descriptor instead.
func (*AppserverListingRequest) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{6}
}

var File_messages_proto protoreflect.FileDescriptor

var file_messages_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x0d, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x61, 0x70,
	0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a,
	0x0c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x63,
	0x0a, 0x0d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x25, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x06, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x37, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbe, 0x07, 0x0a,
	0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x47, 0x0a, 0x10, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6a, 0x77, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x00, 0x52,
	0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x53, 0x0a, 0x11, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x10, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x54, 0x0a, 0x11, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x41, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x10, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x51, 0x0a, 0x10, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x51, 0x0a,
	0x10, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x70,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x0f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x49, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x4a, 0x0a, 0x0f, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x50, 0x0a, 0x0e, 0x6a, 0x6f, 0x69, 0x6e, 0x5f,
	0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x75,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x6a, 0x6f, 0x69, 0x6e,
	0x41, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x5e, 0x0a, 0x15, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70,
	0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70,
	0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x13, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x63, 0x0a, 0x17, 0x61, 0x70, 0x70,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x76, 0x31, 0x2e,
	0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x41, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x15, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x65,
	0x0a, 0x16, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x14, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x28, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xf4, 0x05, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x59, 0x0a, 0x11, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x76, 0x31, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x75, 0x62,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x10, 0x61, 0x70, 0x70,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x55, 0x0a,
	0x11, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70,
	0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x41,
	0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x00, 0x52, 0x10, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x4b, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x00, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x4b, 0x0a, 0x14, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x5f,
	0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41,
	0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x00, 0x52, 0x12, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x64, 0x64, 0x41, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x64,
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x15, 0x61,
	0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x66, 0x0a, 0x16, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41,
	0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x14, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x38, 0x0a, 0x17,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x61, 0x70,
	0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x15, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x70, 0x70,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x12, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x48, 0x00, 0x52, 0x10, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4e, 0x65, 0x77, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x49, 0x0a, 0x15, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x48,
	0x00, 0x52, 0x13, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x19,
	0x0a, 0x17, 0x41, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2a, 0xba, 0x01, 0x0a, 0x0a, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x45, 0x54, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x02,
	0x12, 0x16, 0x0a, 0x12, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x04,
	0x12, 0x16, 0x0a, 0x12, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x05, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_proto_rawDescOnce sync.Once
	file_messages_proto_rawDescData = file_messages_proto_rawDesc
)

func file_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto_rawDescData)
	})
	return file_messages_proto_rawDescData
}

var file_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_messages_proto_goTypes = []any{
	(ActionType)(0),                          // 0: v1.messages.ActionType
	(*InputMessage)(nil),                     // 1: v1.messages.InputMessage
	(*OutputMessage)(nil),                    // 2: v1.messages.OutputMessage
	(*Meta)(nil),                             // 3: v1.messages.Meta
	(*Input)(nil),                            // 4: v1.messages.Input
	(*UpdateJwtToken)(nil),                   // 5: v1.messages.UpdateJwtToken
	(*Output)(nil),                           // 6: v1.messages.Output
	(*AppserverListingRequest)(nil),          // 7: v1.messages.AppserverListingRequest
	(*GetByIdAppserverRequest)(nil),          // 8: v1.appserver.GetByIdAppserverRequest
	(*CreateAppserverRequest)(nil),           // 9: v1.appserver.CreateAppserverRequest
	(*DeleteAppserverRequest)(nil),           // 10: v1.appserver.DeleteAppserverRequest
	(*CreateChannelRequest)(nil),             // 11: v1.channel.CreateChannelRequest
	(*ListChannelsRequest)(nil),              // 12: v1.channel.ListChannelsRequest
	(*CreateAppserverSubRequest)(nil),        // 13: v1.appserver.CreateAppserverSubRequest
	(*CreateAppserverRoleRequest)(nil),       // 14: v1.appserver.CreateAppserverRoleRequest
	(*GetAllAppserverRolesRequest)(nil),      // 15: v1.appserver.GetAllAppserverRolesRequest
	(*GetAllUsersAppserverSubsRequest)(nil),  // 16: v1.appserver.GetAllUsersAppserverSubsRequest
	(*GetUserAppserverSubsResponse)(nil),     // 17: v1.appserver.GetUserAppserverSubsResponse
	(*GetByIdAppserverResponse)(nil),         // 18: v1.appserver.GetByIdAppserverResponse
	(*ListChannelsResponse)(nil),             // 19: v1.channel.ListChannelsResponse
	(*Appserver)(nil),                        // 20: v1.appserver.Appserver
	(*GetAllAppserverRolesResponse)(nil),     // 21: v1.appserver.GetAllAppserverRolesResponse
	(*GetAllUsersAppserverSubsResponse)(nil), // 22: v1.appserver.GetAllUsersAppserverSubsResponse
	(*Channel)(nil),                          // 23: v1.channel.Channel
}
var file_messages_proto_depIdxs = []int32{
	3,  // 0: v1.messages.InputMessage.meta:type_name -> v1.messages.Meta
	4,  // 1: v1.messages.InputMessage.input:type_name -> v1.messages.Input
	3,  // 2: v1.messages.OutputMessage.meta:type_name -> v1.messages.Meta
	6,  // 3: v1.messages.OutputMessage.output:type_name -> v1.messages.Output
	0,  // 4: v1.messages.Meta.action:type_name -> v1.messages.ActionType
	5,  // 5: v1.messages.Input.update_jwt_token:type_name -> v1.messages.UpdateJwtToken
	7,  // 6: v1.messages.Input.appserver_listing:type_name -> v1.messages.AppserverListingRequest
	8,  // 7: v1.messages.Input.appserver_details:type_name -> v1.appserver.GetByIdAppserverRequest
	9,  // 8: v1.messages.Input.create_appserver:type_name -> v1.appserver.CreateAppserverRequest
	10, // 9: v1.messages.Input.delete_appserver:type_name -> v1.appserver.DeleteAppserverRequest
	11, // 10: v1.messages.Input.create_channel:type_name -> v1.channel.CreateChannelRequest
	12, // 11: v1.messages.Input.channel_listing:type_name -> v1.channel.ListChannelsRequest
	13, // 12: v1.messages.Input.join_appserver:type_name -> v1.appserver.CreateAppserverSubRequest
	14, // 13: v1.messages.Input.create_appserver_role:type_name -> v1.appserver.CreateAppserverRoleRequest
	15, // 14: v1.messages.Input.appserver_roles_listing:type_name -> v1.appserver.GetAllAppserverRolesRequest
	16, // 15: v1.messages.Input.appserver_user_listing:type_name -> v1.appserver.GetAllUsersAppserverSubsRequest
	17, // 16: v1.messages.Output.appserver_listing:type_name -> v1.appserver.GetUserAppserverSubsResponse
	18, // 17: v1.messages.Output.appserver_details:type_name -> v1.appserver.GetByIdAppserverResponse
	19, // 18: v1.messages.Output.channel_listing:type_name -> v1.channel.ListChannelsResponse
	20, // 19: v1.messages.Output.update_add_appserver:type_name -> v1.appserver.Appserver
	21, // 20: v1.messages.Output.appserver_roles_listing:type_name -> v1.appserver.GetAllAppserverRolesResponse
	22, // 21: v1.messages.Output.appserver_user_listing:type_name -> v1.appserver.GetAllUsersAppserverSubsResponse
	23, // 22: v1.messages.Output.update_new_channel:type_name -> v1.channel.Channel
	23, // 23: v1.messages.Output.update_delete_channel:type_name -> v1.channel.Channel
	24, // [24:24] is the sub-list for method output_type
	24, // [24:24] is the sub-list for method input_type
	24, // [24:24] is the sub-list for extension type_name
	24, // [24:24] is the sub-list for extension extendee
	0,  // [0:24] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	file_channel_proto_init()
	file_appserver_proto_init()
	file_messages_proto_msgTypes[3].OneofWrappers = []any{
		(*Input_UpdateJwtToken)(nil),
		(*Input_AppserverListing)(nil),
		(*Input_AppserverDetails)(nil),
		(*Input_CreateAppserver)(nil),
		(*Input_DeleteAppserver)(nil),
		(*Input_CreateChannel)(nil),
		(*Input_ChannelListing)(nil),
		(*Input_JoinAppserver)(nil),
		(*Input_CreateAppserverRole)(nil),
		(*Input_AppserverRolesListing)(nil),
		(*Input_AppserverUserListing)(nil),
	}
	file_messages_proto_msgTypes[5].OneofWrappers = []any{
		(*Output_AppserverListing)(nil),
		(*Output_AppserverDetails)(nil),
		(*Output_ChannelListing)(nil),
		(*Output_UpdateAddAppserver)(nil),
		(*Output_AppserverRolesListing)(nil),
		(*Output_AppserverUserListing)(nil),
		(*Output_UpdateRemoveAppserver)(nil),
		(*Output_UpdateNewChannel)(nil),
		(*Output_UpdateDeleteChannel)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_depIdxs,
		EnumInfos:         file_messages_proto_enumTypes,
		MessageInfos:      file_messages_proto_msgTypes,
	}.Build()
	File_messages_proto = out.File
	file_messages_proto_rawDesc = nil
	file_messages_proto_goTypes = nil
	file_messages_proto_depIdxs = nil
}
