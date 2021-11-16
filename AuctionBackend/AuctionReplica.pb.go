// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: AuctionBackend/AuctionReplica.proto

package AuctionReplica

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

type BidReplicaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int32 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *BidReplicaRequest) Reset() {
	*x = BidReplicaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AuctionBackend_AuctionReplica_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidReplicaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidReplicaRequest) ProtoMessage() {}

func (x *BidReplicaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_AuctionBackend_AuctionReplica_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidReplicaRequest.ProtoReflect.Descriptor instead.
func (*BidReplicaRequest) Descriptor() ([]byte, []int) {
	return file_AuctionBackend_AuctionReplica_proto_rawDescGZIP(), []int{0}
}

func (x *BidReplicaRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BidReplicaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack string `protobuf:"bytes,1,opt,name=ack,proto3" json:"ack,omitempty"`
}

func (x *BidReplicaResponse) Reset() {
	*x = BidReplicaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AuctionBackend_AuctionReplica_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidReplicaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidReplicaResponse) ProtoMessage() {}

func (x *BidReplicaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_AuctionBackend_AuctionReplica_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidReplicaResponse.ProtoReflect.Descriptor instead.
func (*BidReplicaResponse) Descriptor() ([]byte, []int) {
	return file_AuctionBackend_AuctionReplica_proto_rawDescGZIP(), []int{1}
}

func (x *BidReplicaResponse) GetAck() string {
	if x != nil {
		return x.Ack
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AuctionBackend_AuctionReplica_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_AuctionBackend_AuctionReplica_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_AuctionBackend_AuctionReplica_proto_rawDescGZIP(), []int{2}
}

type ResultReplicaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ResultReplicaResponse) Reset() {
	*x = ResultReplicaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AuctionBackend_AuctionReplica_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultReplicaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultReplicaResponse) ProtoMessage() {}

func (x *ResultReplicaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_AuctionBackend_AuctionReplica_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultReplicaResponse.ProtoReflect.Descriptor instead.
func (*ResultReplicaResponse) Descriptor() ([]byte, []int) {
	return file_AuctionBackend_AuctionReplica_proto_rawDescGZIP(), []int{3}
}

func (x *ResultReplicaResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_AuctionBackend_AuctionReplica_proto protoreflect.FileDescriptor

var file_AuctionBackend_AuctionReplica_proto_rawDesc = []byte{
	0x0a, 0x23, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x22, 0x2b, 0x0a, 0x11, 0x62, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x26, 0x0a, 0x12, 0x62, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x2f, 0x0a, 0x15, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x32, 0xa3, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x12, 0x4e, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x12, 0x21, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x2e, 0x62, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x41, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x2e, 0x62, 0x69, 0x64, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x48, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x15, 0x2e, 0x41, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x25, 0x2e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f,
	0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AuctionBackend_AuctionReplica_proto_rawDescOnce sync.Once
	file_AuctionBackend_AuctionReplica_proto_rawDescData = file_AuctionBackend_AuctionReplica_proto_rawDesc
)

func file_AuctionBackend_AuctionReplica_proto_rawDescGZIP() []byte {
	file_AuctionBackend_AuctionReplica_proto_rawDescOnce.Do(func() {
		file_AuctionBackend_AuctionReplica_proto_rawDescData = protoimpl.X.CompressGZIP(file_AuctionBackend_AuctionReplica_proto_rawDescData)
	})
	return file_AuctionBackend_AuctionReplica_proto_rawDescData
}

var file_AuctionBackend_AuctionReplica_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_AuctionBackend_AuctionReplica_proto_goTypes = []interface{}{
	(*BidReplicaRequest)(nil),     // 0: AuctionReplica.bidReplicaRequest
	(*BidReplicaResponse)(nil),    // 1: AuctionReplica.bidReplicaResponse
	(*Empty)(nil),                 // 2: AuctionReplica.Empty
	(*ResultReplicaResponse)(nil), // 3: AuctionReplica.resultReplicaResponse
}
var file_AuctionBackend_AuctionReplica_proto_depIdxs = []int32{
	0, // 0: AuctionReplica.Replica.bid:input_type -> AuctionReplica.bidReplicaRequest
	2, // 1: AuctionReplica.Replica.result:input_type -> AuctionReplica.Empty
	1, // 2: AuctionReplica.Replica.bid:output_type -> AuctionReplica.bidReplicaResponse
	3, // 3: AuctionReplica.Replica.result:output_type -> AuctionReplica.resultReplicaResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AuctionBackend_AuctionReplica_proto_init() }
func file_AuctionBackend_AuctionReplica_proto_init() {
	if File_AuctionBackend_AuctionReplica_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AuctionBackend_AuctionReplica_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidReplicaRequest); i {
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
		file_AuctionBackend_AuctionReplica_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidReplicaResponse); i {
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
		file_AuctionBackend_AuctionReplica_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_AuctionBackend_AuctionReplica_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultReplicaResponse); i {
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
			RawDescriptor: file_AuctionBackend_AuctionReplica_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_AuctionBackend_AuctionReplica_proto_goTypes,
		DependencyIndexes: file_AuctionBackend_AuctionReplica_proto_depIdxs,
		MessageInfos:      file_AuctionBackend_AuctionReplica_proto_msgTypes,
	}.Build()
	File_AuctionBackend_AuctionReplica_proto = out.File
	file_AuctionBackend_AuctionReplica_proto_rawDesc = nil
	file_AuctionBackend_AuctionReplica_proto_goTypes = nil
	file_AuctionBackend_AuctionReplica_proto_depIdxs = nil
}
