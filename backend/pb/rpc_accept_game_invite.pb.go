// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpc_accept_game_invite.proto

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

type AcceptGameInviteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InviteeUsername string `protobuf:"bytes,1,opt,name=invitee_username,json=inviteeUsername,proto3" json:"invitee_username,omitempty"`
	GameId          int64  `protobuf:"varint,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (x *AcceptGameInviteRequest) Reset() {
	*x = AcceptGameInviteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_accept_game_invite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptGameInviteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptGameInviteRequest) ProtoMessage() {}

func (x *AcceptGameInviteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_accept_game_invite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptGameInviteRequest.ProtoReflect.Descriptor instead.
func (*AcceptGameInviteRequest) Descriptor() ([]byte, []int) {
	return file_rpc_accept_game_invite_proto_rawDescGZIP(), []int{0}
}

func (x *AcceptGameInviteRequest) GetInviteeUsername() string {
	if x != nil {
		return x.InviteeUsername
	}
	return ""
}

func (x *AcceptGameInviteRequest) GetGameId() int64 {
	if x != nil {
		return x.GameId
	}
	return 0
}

var File_rpc_accept_game_invite_proto protoreflect.FileDescriptor

var file_rpc_accept_game_invite_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x67, 0x61, 0x6d,
	0x65, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x5d, 0x0a, 0x17, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a,
	0x10, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49,
	0x64, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x50, 0x6c, 0x61, 0x74, 0x6f, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x73, 0x73,
	0x65, 0x72, 0x74, 0x65, 0x64, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_accept_game_invite_proto_rawDescOnce sync.Once
	file_rpc_accept_game_invite_proto_rawDescData = file_rpc_accept_game_invite_proto_rawDesc
)

func file_rpc_accept_game_invite_proto_rawDescGZIP() []byte {
	file_rpc_accept_game_invite_proto_rawDescOnce.Do(func() {
		file_rpc_accept_game_invite_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_accept_game_invite_proto_rawDescData)
	})
	return file_rpc_accept_game_invite_proto_rawDescData
}

var file_rpc_accept_game_invite_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_rpc_accept_game_invite_proto_goTypes = []interface{}{
	(*AcceptGameInviteRequest)(nil), // 0: pb.AcceptGameInviteRequest
}
var file_rpc_accept_game_invite_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_accept_game_invite_proto_init() }
func file_rpc_accept_game_invite_proto_init() {
	if File_rpc_accept_game_invite_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_accept_game_invite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptGameInviteRequest); i {
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
			RawDescriptor: file_rpc_accept_game_invite_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_accept_game_invite_proto_goTypes,
		DependencyIndexes: file_rpc_accept_game_invite_proto_depIdxs,
		MessageInfos:      file_rpc_accept_game_invite_proto_msgTypes,
	}.Build()
	File_rpc_accept_game_invite_proto = out.File
	file_rpc_accept_game_invite_proto_rawDesc = nil
	file_rpc_accept_game_invite_proto_goTypes = nil
	file_rpc_accept_game_invite_proto_depIdxs = nil
}
