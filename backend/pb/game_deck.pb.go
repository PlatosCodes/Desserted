// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: game_deck.proto

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

type GameDeck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameDeckId int64 `protobuf:"varint,1,opt,name=game_deck_id,json=gameDeckId,proto3" json:"game_deck_id,omitempty"`
	GameId     int64 `protobuf:"varint,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	CardId     int32 `protobuf:"varint,3,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
	OrderIndex int32 `protobuf:"varint,4,opt,name=order_index,json=orderIndex,proto3" json:"order_index,omitempty"`
}

func (x *GameDeck) Reset() {
	*x = GameDeck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_deck_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameDeck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameDeck) ProtoMessage() {}

func (x *GameDeck) ProtoReflect() protoreflect.Message {
	mi := &file_game_deck_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameDeck.ProtoReflect.Descriptor instead.
func (*GameDeck) Descriptor() ([]byte, []int) {
	return file_game_deck_proto_rawDescGZIP(), []int{0}
}

func (x *GameDeck) GetGameDeckId() int64 {
	if x != nil {
		return x.GameDeckId
	}
	return 0
}

func (x *GameDeck) GetGameId() int64 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *GameDeck) GetCardId() int32 {
	if x != nil {
		return x.CardId
	}
	return 0
}

func (x *GameDeck) GetOrderIndex() int32 {
	if x != nil {
		return x.OrderIndex
	}
	return 0
}

var File_game_deck_proto protoreflect.FileDescriptor

var file_game_deck_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x7f, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x65, 0x63,
	0x6b, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x44, 0x65, 0x63,
	0x6b, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x6c, 0x61, 0x74, 0x6f, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x73,
	0x2f, 0x64, 0x65, 0x73, 0x73, 0x65, 0x72, 0x74, 0x65, 0x64, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_deck_proto_rawDescOnce sync.Once
	file_game_deck_proto_rawDescData = file_game_deck_proto_rawDesc
)

func file_game_deck_proto_rawDescGZIP() []byte {
	file_game_deck_proto_rawDescOnce.Do(func() {
		file_game_deck_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_deck_proto_rawDescData)
	})
	return file_game_deck_proto_rawDescData
}

var file_game_deck_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_game_deck_proto_goTypes = []interface{}{
	(*GameDeck)(nil), // 0: pb.GameDeck
}
var file_game_deck_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_game_deck_proto_init() }
func file_game_deck_proto_init() {
	if File_game_deck_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_deck_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameDeck); i {
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
			RawDescriptor: file_game_deck_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_deck_proto_goTypes,
		DependencyIndexes: file_game_deck_proto_depIdxs,
		MessageInfos:      file_game_deck_proto_msgTypes,
	}.Build()
	File_game_deck_proto = out.File
	file_game_deck_proto_rawDesc = nil
	file_game_deck_proto_goTypes = nil
	file_game_deck_proto_depIdxs = nil
}
