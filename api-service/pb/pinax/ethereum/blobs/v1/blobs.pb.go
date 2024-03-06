// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: pinax/ethereum/blobs/v1/blobs.proto

package pbbl

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

type Spec int32

const (
	Spec_UNSPECIFIED Spec = 0
	Spec_PHASE0      Spec = 1
	Spec_ALTAIR      Spec = 2
	Spec_BELLATRIX   Spec = 3
	Spec_CAPELLA     Spec = 4
	Spec_DENEB       Spec = 5
)

// Enum value maps for Spec.
var (
	Spec_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "PHASE0",
		2: "ALTAIR",
		3: "BELLATRIX",
		4: "CAPELLA",
		5: "DENEB",
	}
	Spec_value = map[string]int32{
		"UNSPECIFIED": 0,
		"PHASE0":      1,
		"ALTAIR":      2,
		"BELLATRIX":   3,
		"CAPELLA":     4,
		"DENEB":       5,
	}
)

func (x Spec) Enum() *Spec {
	p := new(Spec)
	*p = x
	return p
}

func (x Spec) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Spec) Descriptor() protoreflect.EnumDescriptor {
	return file_pinax_ethereum_blobs_v1_blobs_proto_enumTypes[0].Descriptor()
}

func (Spec) Type() protoreflect.EnumType {
	return &file_pinax_ethereum_blobs_v1_blobs_proto_enumTypes[0]
}

func (x Spec) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Spec.Descriptor instead.
func (Spec) EnumDescriptor() ([]byte, []int) {
	return file_pinax_ethereum_blobs_v1_blobs_proto_rawDescGZIP(), []int{0}
}

type Slot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot          uint64                 `protobuf:"varint,1,opt,name=slot,proto3" json:"slot"`
	Spec          Spec                   `protobuf:"varint,2,opt,name=spec,proto3,enum=pinax.ethereum.blobs.v1.Spec" json:"spec"`
	Root          []byte                 `protobuf:"bytes,3,opt,name=root,proto3" json:"root"`
	ProposerIndex uint64                 `protobuf:"varint,4,opt,name=proposer_index,json=proposerIndex,proto3" json:"proposer_index"`
	ParentRoot    []byte                 `protobuf:"bytes,5,opt,name=parent_root,json=parentRoot,proto3" json:"parent_root"`
	StateRoot     []byte                 `protobuf:"bytes,6,opt,name=state_root,json=stateRoot,proto3" json:"state_root"`
	BodyRoot      []byte                 `protobuf:"bytes,7,opt,name=body_root,json=bodyRoot,proto3" json:"body_root"`
	Signature     []byte                 `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=timestamp,proto3" json:"timestamp"`
	Blobs         []*Blob                `protobuf:"bytes,20,rep,name=blobs,proto3" json:"blobs"`
}

func (x *Slot) Reset() {
	*x = Slot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pinax_ethereum_blobs_v1_blobs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slot) ProtoMessage() {}

func (x *Slot) ProtoReflect() protoreflect.Message {
	mi := &file_pinax_ethereum_blobs_v1_blobs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slot.ProtoReflect.Descriptor instead.
func (*Slot) Descriptor() ([]byte, []int) {
	return file_pinax_ethereum_blobs_v1_blobs_proto_rawDescGZIP(), []int{0}
}

func (x *Slot) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *Slot) GetSpec() Spec {
	if x != nil {
		return x.Spec
	}
	return Spec_UNSPECIFIED
}

func (x *Slot) GetRoot() []byte {
	if x != nil {
		return x.Root
	}
	return nil
}

func (x *Slot) GetProposerIndex() uint64 {
	if x != nil {
		return x.ProposerIndex
	}
	return 0
}

func (x *Slot) GetParentRoot() []byte {
	if x != nil {
		return x.ParentRoot
	}
	return nil
}

func (x *Slot) GetStateRoot() []byte {
	if x != nil {
		return x.StateRoot
	}
	return nil
}

func (x *Slot) GetBodyRoot() []byte {
	if x != nil {
		return x.BodyRoot
	}
	return nil
}

func (x *Slot) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Slot) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Slot) GetBlobs() []*Blob {
	if x != nil {
		return x.Blobs
	}
	return nil
}

type Blob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index                       uint32   `protobuf:"varint,1,opt,name=index,proto3" json:"index"`
	Blob                        []byte   `protobuf:"bytes,2,opt,name=blob,proto3" json:"blob"`
	KzgCommitment               []byte   `protobuf:"bytes,3,opt,name=kzg_commitment,json=kzgCommitment,proto3" json:"kzg_commitment"`
	KzgProof                    []byte   `protobuf:"bytes,4,opt,name=kzg_proof,json=kzgProof,proto3" json:"kzg_proof"`
	KzgCommitmentInclusionProof [][]byte `protobuf:"bytes,6,rep,name=kzg_commitment_inclusion_proof,json=kzgCommitmentInclusionProof,proto3" json:"kzg_commitment_inclusion_proof"`
}

func (x *Blob) Reset() {
	*x = Blob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pinax_ethereum_blobs_v1_blobs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blob) ProtoMessage() {}

func (x *Blob) ProtoReflect() protoreflect.Message {
	mi := &file_pinax_ethereum_blobs_v1_blobs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blob.ProtoReflect.Descriptor instead.
func (*Blob) Descriptor() ([]byte, []int) {
	return file_pinax_ethereum_blobs_v1_blobs_proto_rawDescGZIP(), []int{1}
}

func (x *Blob) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Blob) GetBlob() []byte {
	if x != nil {
		return x.Blob
	}
	return nil
}

func (x *Blob) GetKzgCommitment() []byte {
	if x != nil {
		return x.KzgCommitment
	}
	return nil
}

func (x *Blob) GetKzgProof() []byte {
	if x != nil {
		return x.KzgProof
	}
	return nil
}

func (x *Blob) GetKzgCommitmentInclusionProof() [][]byte {
	if x != nil {
		return x.KzgCommitmentInclusionProof
	}
	return nil
}

var File_pinax_ethereum_blobs_v1_blobs_proto protoreflect.FileDescriptor

var file_pinax_ethereum_blobs_v1_blobs_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x69, 0x6e, 0x61, 0x78, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x70, 0x69, 0x6e, 0x61, 0x78, 0x2e, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf2, 0x02, 0x0a, 0x04, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x31, 0x0a, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x69, 0x6e,
	0x61, 0x78, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x62,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x72,
	0x6f, 0x6f, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f,
	0x64, 0x79, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x62,
	0x6f, 0x64, 0x79, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x33, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x70, 0x69, 0x6e, 0x61, 0x78, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e,
	0x62, 0x6c, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x05, 0x62,
	0x6c, 0x6f, 0x62, 0x73, 0x22, 0xb9, 0x01, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x12, 0x25, 0x0a, 0x0e, 0x6b, 0x7a, 0x67, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0d, 0x6b, 0x7a, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x6b, 0x7a, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x6b, 0x7a, 0x67, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x43, 0x0a, 0x1e, 0x6b,
	0x7a, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x1b, 0x6b, 0x7a, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x2a, 0x56, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x48, 0x41,
	0x53, 0x45, 0x30, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x4c, 0x54, 0x41, 0x49, 0x52, 0x10,
	0x02, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x45, 0x4c, 0x4c, 0x41, 0x54, 0x52, 0x49, 0x58, 0x10, 0x03,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x41, 0x50, 0x45, 0x4c, 0x4c, 0x41, 0x10, 0x04, 0x12, 0x09, 0x0a,
	0x05, 0x44, 0x45, 0x4e, 0x45, 0x42, 0x10, 0x05, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69, 0x6e, 0x61, 0x78, 0x2d, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x69, 0x6e, 0x61, 0x78, 0x2f, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62,
	0x62, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pinax_ethereum_blobs_v1_blobs_proto_rawDescOnce sync.Once
	file_pinax_ethereum_blobs_v1_blobs_proto_rawDescData = file_pinax_ethereum_blobs_v1_blobs_proto_rawDesc
)

func file_pinax_ethereum_blobs_v1_blobs_proto_rawDescGZIP() []byte {
	file_pinax_ethereum_blobs_v1_blobs_proto_rawDescOnce.Do(func() {
		file_pinax_ethereum_blobs_v1_blobs_proto_rawDescData = protoimpl.X.CompressGZIP(file_pinax_ethereum_blobs_v1_blobs_proto_rawDescData)
	})
	return file_pinax_ethereum_blobs_v1_blobs_proto_rawDescData
}

var file_pinax_ethereum_blobs_v1_blobs_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pinax_ethereum_blobs_v1_blobs_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pinax_ethereum_blobs_v1_blobs_proto_goTypes = []interface{}{
	(Spec)(0),                     // 0: pinax.ethereum.blobs.v1.Spec
	(*Slot)(nil),                  // 1: pinax.ethereum.blobs.v1.Slot
	(*Blob)(nil),                  // 2: pinax.ethereum.blobs.v1.Blob
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_pinax_ethereum_blobs_v1_blobs_proto_depIdxs = []int32{
	0, // 0: pinax.ethereum.blobs.v1.Slot.spec:type_name -> pinax.ethereum.blobs.v1.Spec
	3, // 1: pinax.ethereum.blobs.v1.Slot.timestamp:type_name -> google.protobuf.Timestamp
	2, // 2: pinax.ethereum.blobs.v1.Slot.blobs:type_name -> pinax.ethereum.blobs.v1.Blob
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pinax_ethereum_blobs_v1_blobs_proto_init() }
func file_pinax_ethereum_blobs_v1_blobs_proto_init() {
	if File_pinax_ethereum_blobs_v1_blobs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pinax_ethereum_blobs_v1_blobs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slot); i {
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
		file_pinax_ethereum_blobs_v1_blobs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blob); i {
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
			RawDescriptor: file_pinax_ethereum_blobs_v1_blobs_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pinax_ethereum_blobs_v1_blobs_proto_goTypes,
		DependencyIndexes: file_pinax_ethereum_blobs_v1_blobs_proto_depIdxs,
		EnumInfos:         file_pinax_ethereum_blobs_v1_blobs_proto_enumTypes,
		MessageInfos:      file_pinax_ethereum_blobs_v1_blobs_proto_msgTypes,
	}.Build()
	File_pinax_ethereum_blobs_v1_blobs_proto = out.File
	file_pinax_ethereum_blobs_v1_blobs_proto_rawDesc = nil
	file_pinax_ethereum_blobs_v1_blobs_proto_goTypes = nil
	file_pinax_ethereum_blobs_v1_blobs_proto_depIdxs = nil
}