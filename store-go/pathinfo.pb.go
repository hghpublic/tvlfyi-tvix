// SPDX-License-Identifier: MIT
// Copyright © 2022 The Tvix Authors

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tvix/store/protos/pathinfo.proto

package storev1

import (
	protos "code.tvl.fyi/tvix/castore-go"
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

// PathInfo shows information about a Nix Store Path.
// That's a single element inside /nix/store.
type PathInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path can be a directory, file or symlink.
	Node *protos.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// List of references (output path hashes)
	// This really is the raw *bytes*, after decoding nixbase32, and not a
	// base32-encoded string.
	References [][]byte `protobuf:"bytes,2,rep,name=references,proto3" json:"references,omitempty"`
	// see below.
	Narinfo *NARInfo `protobuf:"bytes,3,opt,name=narinfo,proto3" json:"narinfo,omitempty"`
	// The StorePath of the .drv file producing this output.
	// The .drv suffix is omitted in its `name` field.
	Deriver *StorePath `protobuf:"bytes,4,opt,name=deriver,proto3" json:"deriver,omitempty"`
}

func (x *PathInfo) Reset() {
	*x = PathInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tvix_store_protos_pathinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathInfo) ProtoMessage() {}

func (x *PathInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tvix_store_protos_pathinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathInfo.ProtoReflect.Descriptor instead.
func (*PathInfo) Descriptor() ([]byte, []int) {
	return file_tvix_store_protos_pathinfo_proto_rawDescGZIP(), []int{0}
}

func (x *PathInfo) GetNode() *protos.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *PathInfo) GetReferences() [][]byte {
	if x != nil {
		return x.References
	}
	return nil
}

func (x *PathInfo) GetNarinfo() *NARInfo {
	if x != nil {
		return x.Narinfo
	}
	return nil
}

func (x *PathInfo) GetDeriver() *StorePath {
	if x != nil {
		return x.Deriver
	}
	return nil
}

// Represents a path in the Nix store (a direct child of STORE_DIR).
// It is commonly formatted by a nixbase32-encoding the digest, and
// concatenating the name, separated by a `-`.
type StorePath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The string after digest and `-`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The digest (20 bytes).
	Digest []byte `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (x *StorePath) Reset() {
	*x = StorePath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tvix_store_protos_pathinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorePath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorePath) ProtoMessage() {}

func (x *StorePath) ProtoReflect() protoreflect.Message {
	mi := &file_tvix_store_protos_pathinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorePath.ProtoReflect.Descriptor instead.
func (*StorePath) Descriptor() ([]byte, []int) {
	return file_tvix_store_protos_pathinfo_proto_rawDescGZIP(), []int{1}
}

func (x *StorePath) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StorePath) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

// Nix C++ uses NAR (Nix Archive) as a format to transfer store paths,
// and stores metadata and signatures in NARInfo files.
// Store all these attributes in a separate message.
//
// This is useful to render .narinfo files to clients, or to preserve/validate
// these signatures.
// As verifying these signatures requires the whole NAR file to be synthesized,
// moving to another signature scheme is desired.
// Even then, it still makes sense to hold this data, for old clients.
type NARInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This size of the NAR file, in bytes.
	NarSize uint64 `protobuf:"varint,1,opt,name=nar_size,json=narSize,proto3" json:"nar_size,omitempty"`
	// The sha256 of the NAR file representation.
	NarSha256 []byte `protobuf:"bytes,2,opt,name=nar_sha256,json=narSha256,proto3" json:"nar_sha256,omitempty"`
	// The signatures in a .narinfo file.
	Signatures []*NARInfo_Signature `protobuf:"bytes,3,rep,name=signatures,proto3" json:"signatures,omitempty"`
	// A list of references. To validate .narinfo signatures, a fingerprint
	// needs to be constructed.
	// This fingerprint doesn't just contain the hashes of the output paths of
	// all references (like PathInfo.references), but their whole (base)names,
	// so we need to keep them somewhere.
	ReferenceNames []string `protobuf:"bytes,4,rep,name=reference_names,json=referenceNames,proto3" json:"reference_names,omitempty"`
}

func (x *NARInfo) Reset() {
	*x = NARInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tvix_store_protos_pathinfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NARInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NARInfo) ProtoMessage() {}

func (x *NARInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tvix_store_protos_pathinfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NARInfo.ProtoReflect.Descriptor instead.
func (*NARInfo) Descriptor() ([]byte, []int) {
	return file_tvix_store_protos_pathinfo_proto_rawDescGZIP(), []int{2}
}

func (x *NARInfo) GetNarSize() uint64 {
	if x != nil {
		return x.NarSize
	}
	return 0
}

func (x *NARInfo) GetNarSha256() []byte {
	if x != nil {
		return x.NarSha256
	}
	return nil
}

func (x *NARInfo) GetSignatures() []*NARInfo_Signature {
	if x != nil {
		return x.Signatures
	}
	return nil
}

func (x *NARInfo) GetReferenceNames() []string {
	if x != nil {
		return x.ReferenceNames
	}
	return nil
}

// This represents a (parsed) signature line in a .narinfo file.
type NARInfo_Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *NARInfo_Signature) Reset() {
	*x = NARInfo_Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tvix_store_protos_pathinfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NARInfo_Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NARInfo_Signature) ProtoMessage() {}

func (x *NARInfo_Signature) ProtoReflect() protoreflect.Message {
	mi := &file_tvix_store_protos_pathinfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NARInfo_Signature.ProtoReflect.Descriptor instead.
func (*NARInfo_Signature) Descriptor() ([]byte, []int) {
	return file_tvix_store_protos_pathinfo_proto_rawDescGZIP(), []int{2, 0}
}

func (x *NARInfo_Signature) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NARInfo_Signature) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_tvix_store_protos_pathinfo_proto protoreflect.FileDescriptor

var file_tvix_store_protos_pathinfo_proto_rawDesc = []byte{
	0x0a, 0x20, 0x74, 0x76, 0x69, 0x78, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x70, 0x61, 0x74, 0x68, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x74, 0x76, 0x69, 0x78, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x21, 0x74, 0x76, 0x69, 0x78, 0x2f, 0x63, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x08, 0x50, 0x61, 0x74, 0x68, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x29, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x74, 0x76, 0x69, 0x78, 0x2e, 0x63, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x07,
	0x6e, 0x61, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x74, 0x76, 0x69, 0x78, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x41,
	0x52, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x6e, 0x61, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x32,
	0x0a, 0x07, 0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x74, 0x76, 0x69, 0x78, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x74, 0x68, 0x52, 0x07, 0x64, 0x65, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x22, 0x37, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x22, 0xe3, 0x01, 0x0a, 0x07,
	0x4e, 0x41, 0x52, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x61, 0x72, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6e, 0x61, 0x72, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x61, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6e, 0x61, 0x72, 0x53, 0x68, 0x61, 0x32, 0x35,
	0x36, 0x12, 0x40, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x76, 0x69, 0x78, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x41, 0x52, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x33, 0x0a, 0x09,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x28, 0x5a, 0x26, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x74, 0x76, 0x6c, 0x2e, 0x66, 0x79,
	0x69, 0x2f, 0x74, 0x76, 0x69, 0x78, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tvix_store_protos_pathinfo_proto_rawDescOnce sync.Once
	file_tvix_store_protos_pathinfo_proto_rawDescData = file_tvix_store_protos_pathinfo_proto_rawDesc
)

func file_tvix_store_protos_pathinfo_proto_rawDescGZIP() []byte {
	file_tvix_store_protos_pathinfo_proto_rawDescOnce.Do(func() {
		file_tvix_store_protos_pathinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_tvix_store_protos_pathinfo_proto_rawDescData)
	})
	return file_tvix_store_protos_pathinfo_proto_rawDescData
}

var file_tvix_store_protos_pathinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tvix_store_protos_pathinfo_proto_goTypes = []interface{}{
	(*PathInfo)(nil),          // 0: tvix.store.v1.PathInfo
	(*StorePath)(nil),         // 1: tvix.store.v1.StorePath
	(*NARInfo)(nil),           // 2: tvix.store.v1.NARInfo
	(*NARInfo_Signature)(nil), // 3: tvix.store.v1.NARInfo.Signature
	(*protos.Node)(nil),       // 4: tvix.castore.v1.Node
}
var file_tvix_store_protos_pathinfo_proto_depIdxs = []int32{
	4, // 0: tvix.store.v1.PathInfo.node:type_name -> tvix.castore.v1.Node
	2, // 1: tvix.store.v1.PathInfo.narinfo:type_name -> tvix.store.v1.NARInfo
	1, // 2: tvix.store.v1.PathInfo.deriver:type_name -> tvix.store.v1.StorePath
	3, // 3: tvix.store.v1.NARInfo.signatures:type_name -> tvix.store.v1.NARInfo.Signature
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tvix_store_protos_pathinfo_proto_init() }
func file_tvix_store_protos_pathinfo_proto_init() {
	if File_tvix_store_protos_pathinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tvix_store_protos_pathinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathInfo); i {
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
		file_tvix_store_protos_pathinfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorePath); i {
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
		file_tvix_store_protos_pathinfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NARInfo); i {
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
		file_tvix_store_protos_pathinfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NARInfo_Signature); i {
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
			RawDescriptor: file_tvix_store_protos_pathinfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tvix_store_protos_pathinfo_proto_goTypes,
		DependencyIndexes: file_tvix_store_protos_pathinfo_proto_depIdxs,
		MessageInfos:      file_tvix_store_protos_pathinfo_proto_msgTypes,
	}.Build()
	File_tvix_store_protos_pathinfo_proto = out.File
	file_tvix_store_protos_pathinfo_proto_rawDesc = nil
	file_tvix_store_protos_pathinfo_proto_goTypes = nil
	file_tvix_store_protos_pathinfo_proto_depIdxs = nil
}
