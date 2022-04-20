// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: nft_explorer.proto

package nft_explorer_proto

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

// Account represents the address of the wallet of an account
type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_explorer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_nft_explorer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_nft_explorer_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// Artwork represents the content of a nft token.
type Artwork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl    string `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *Artwork) Reset() {
	*x = Artwork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_explorer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Artwork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Artwork) ProtoMessage() {}

func (x *Artwork) ProtoReflect() protoreflect.Message {
	mi := &file_nft_explorer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Artwork.ProtoReflect.Descriptor instead.
func (*Artwork) Descriptor() ([]byte, []int) {
	return file_nft_explorer_proto_rawDescGZIP(), []int{1}
}

func (x *Artwork) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Artwork) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Artwork) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

var File_nft_explorer_proto protoreflect.FileDescriptor

var file_nft_explorer_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x66, 0x74, 0x5f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6e, 0x66, 0x74, 0x5f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72,
	0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5c, 0x0a,
	0x07, 0x41, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x32, 0x60, 0x0a, 0x0b, 0x4e,
	0x46, 0x54, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x41, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12,
	0x1b, 0x2e, 0x6e, 0x66, 0x74, 0x5f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x1b, 0x2e, 0x6e,
	0x66, 0x74, 0x5f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x72, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x42, 0x6a, 0x5a,
	0x68, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6c, 0x61, 0x63,
	0x75, 0x69, 0x6c, 0x6f, 0x73, 0x65, 0x2f, 0x6e, 0x66, 0x74, 0x2d, 0x65, 0x78, 0x70, 0x6c, 0x6f,
	0x72, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x66,
	0x74, 0x5f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6e, 0x66, 0x74, 0x5f, 0x65, 0x78, 0x70, 0x6c, 0x6f,
	0x72, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_nft_explorer_proto_rawDescOnce sync.Once
	file_nft_explorer_proto_rawDescData = file_nft_explorer_proto_rawDesc
)

func file_nft_explorer_proto_rawDescGZIP() []byte {
	file_nft_explorer_proto_rawDescOnce.Do(func() {
		file_nft_explorer_proto_rawDescData = protoimpl.X.CompressGZIP(file_nft_explorer_proto_rawDescData)
	})
	return file_nft_explorer_proto_rawDescData
}

var file_nft_explorer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_nft_explorer_proto_goTypes = []interface{}{
	(*Account)(nil), // 0: nft_explorer_proto.Account
	(*Artwork)(nil), // 1: nft_explorer_proto.Artwork
}
var file_nft_explorer_proto_depIdxs = []int32{
	0, // 0: nft_explorer_proto.NFTExplorer.ListOwnedArtworks:input_type -> nft_explorer_proto.Account
	1, // 1: nft_explorer_proto.NFTExplorer.ListOwnedArtworks:output_type -> nft_explorer_proto.Artwork
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nft_explorer_proto_init() }
func file_nft_explorer_proto_init() {
	if File_nft_explorer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nft_explorer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_nft_explorer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Artwork); i {
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
			RawDescriptor: file_nft_explorer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nft_explorer_proto_goTypes,
		DependencyIndexes: file_nft_explorer_proto_depIdxs,
		MessageInfos:      file_nft_explorer_proto_msgTypes,
	}.Build()
	File_nft_explorer_proto = out.File
	file_nft_explorer_proto_rawDesc = nil
	file_nft_explorer_proto_goTypes = nil
	file_nft_explorer_proto_depIdxs = nil
}