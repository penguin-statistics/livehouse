// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: liveconn.proto

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

type ReportBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Generation uint32    `protobuf:"varint,1,opt,name=generation,proto3" json:"generation,omitempty"`
	Report     []*Report `protobuf:"bytes,2,rep,name=report,proto3" json:"report,omitempty"`
}

func (x *ReportBatchRequest) Reset() {
	*x = ReportBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liveconn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportBatchRequest) ProtoMessage() {}

func (x *ReportBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_liveconn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportBatchRequest.ProtoReflect.Descriptor instead.
func (*ReportBatchRequest) Descriptor() ([]byte, []int) {
	return file_liveconn_proto_rawDescGZIP(), []int{0}
}

func (x *ReportBatchRequest) GetGeneration() uint32 {
	if x != nil {
		return x.Generation
	}
	return 0
}

func (x *ReportBatchRequest) GetReport() []*Report {
	if x != nil {
		return x.Report
	}
	return nil
}

type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId uint32  `protobuf:"varint,1,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	Drops   []*Drop `protobuf:"bytes,2,rep,name=drops,proto3" json:"drops,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liveconn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_liveconn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_liveconn_proto_rawDescGZIP(), []int{1}
}

func (x *Report) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *Report) GetDrops() []*Drop {
	if x != nil {
		return x.Drops
	}
	return nil
}

type Drop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId   uint32 `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Quantity uint32 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Drop) Reset() {
	*x = Drop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liveconn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Drop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Drop) ProtoMessage() {}

func (x *Drop) ProtoReflect() protoreflect.Message {
	mi := &file_liveconn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Drop.ProtoReflect.Descriptor instead.
func (*Drop) Descriptor() ([]byte, []int) {
	return file_liveconn_proto_rawDescGZIP(), []int{2}
}

func (x *Drop) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *Drop) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ReportBatchACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Generation uint32 `protobuf:"varint,1,opt,name=generation,proto3" json:"generation,omitempty"`
}

func (x *ReportBatchACK) Reset() {
	*x = ReportBatchACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liveconn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportBatchACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportBatchACK) ProtoMessage() {}

func (x *ReportBatchACK) ProtoReflect() protoreflect.Message {
	mi := &file_liveconn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportBatchACK.ProtoReflect.Descriptor instead.
func (*ReportBatchACK) Descriptor() ([]byte, []int) {
	return file_liveconn_proto_rawDescGZIP(), []int{3}
}

func (x *ReportBatchACK) GetGeneration() uint32 {
	if x != nil {
		return x.Generation
	}
	return 0
}

type MatrixBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Generation uint32    `protobuf:"varint,1,opt,name=generation,proto3" json:"generation,omitempty"`
	Matrix     []*Matrix `protobuf:"bytes,2,rep,name=matrix,proto3" json:"matrix,omitempty"`
}

func (x *MatrixBatchRequest) Reset() {
	*x = MatrixBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liveconn_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatrixBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatrixBatchRequest) ProtoMessage() {}

func (x *MatrixBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_liveconn_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatrixBatchRequest.ProtoReflect.Descriptor instead.
func (*MatrixBatchRequest) Descriptor() ([]byte, []int) {
	return file_liveconn_proto_rawDescGZIP(), []int{4}
}

func (x *MatrixBatchRequest) GetGeneration() uint32 {
	if x != nil {
		return x.Generation
	}
	return 0
}

func (x *MatrixBatchRequest) GetMatrix() []*Matrix {
	if x != nil {
		return x.Matrix
	}
	return nil
}

type Matrix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId  uint32 `protobuf:"varint,1,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	ItemId   uint32 `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Quantity uint64 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Times    uint64 `protobuf:"varint,4,opt,name=times,proto3" json:"times,omitempty"`
}

func (x *Matrix) Reset() {
	*x = Matrix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liveconn_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Matrix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Matrix) ProtoMessage() {}

func (x *Matrix) ProtoReflect() protoreflect.Message {
	mi := &file_liveconn_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Matrix.ProtoReflect.Descriptor instead.
func (*Matrix) Descriptor() ([]byte, []int) {
	return file_liveconn_proto_rawDescGZIP(), []int{5}
}

func (x *Matrix) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *Matrix) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *Matrix) GetQuantity() uint64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Matrix) GetTimes() uint64 {
	if x != nil {
		return x.Times
	}
	return 0
}

type MatrixBatchACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Generation uint32 `protobuf:"varint,1,opt,name=generation,proto3" json:"generation,omitempty"`
}

func (x *MatrixBatchACK) Reset() {
	*x = MatrixBatchACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_liveconn_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatrixBatchACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatrixBatchACK) ProtoMessage() {}

func (x *MatrixBatchACK) ProtoReflect() protoreflect.Message {
	mi := &file_liveconn_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatrixBatchACK.ProtoReflect.Descriptor instead.
func (*MatrixBatchACK) Descriptor() ([]byte, []int) {
	return file_liveconn_proto_rawDescGZIP(), []int{6}
}

func (x *MatrixBatchACK) GetGeneration() uint32 {
	if x != nil {
		return x.Generation
	}
	return 0
}

var File_liveconn_proto protoreflect.FileDescriptor

var file_liveconn_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x69, 0x76, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x55, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x40, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x05,
	0x64, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x44, 0x72,
	0x6f, 0x70, 0x52, 0x05, 0x64, 0x72, 0x6f, 0x70, 0x73, 0x22, 0x3b, 0x0a, 0x04, 0x44, 0x72, 0x6f,
	0x70, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x30, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x41, 0x43, 0x4b, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x12, 0x4d, 0x61, 0x74, 0x72,
	0x69, 0x78, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x0a, 0x06, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x52, 0x06, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x22,
	0x6e, 0x0a, 0x06, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x22,
	0x30, 0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x42, 0x61, 0x74, 0x63, 0x68, 0x41, 0x43,
	0x4b, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x32, 0x8c, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4c,
	0x69, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0f, 0x50, 0x75,
	0x73, 0x68, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x13, 0x2e,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x41, 0x43, 0x4b, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0f, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x61, 0x74,
	0x72, 0x69, 0x78, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x13, 0x2e, 0x4d, 0x61, 0x74, 0x72, 0x69,
	0x78, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x42, 0x61, 0x74, 0x63, 0x68, 0x41, 0x43, 0x4b, 0x22, 0x00,
	0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x65, 0x6e, 0x67, 0x75, 0x69, 0x6e, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_liveconn_proto_rawDescOnce sync.Once
	file_liveconn_proto_rawDescData = file_liveconn_proto_rawDesc
)

func file_liveconn_proto_rawDescGZIP() []byte {
	file_liveconn_proto_rawDescOnce.Do(func() {
		file_liveconn_proto_rawDescData = protoimpl.X.CompressGZIP(file_liveconn_proto_rawDescData)
	})
	return file_liveconn_proto_rawDescData
}

var file_liveconn_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_liveconn_proto_goTypes = []interface{}{
	(*ReportBatchRequest)(nil), // 0: ReportBatchRequest
	(*Report)(nil),             // 1: Report
	(*Drop)(nil),               // 2: Drop
	(*ReportBatchACK)(nil),     // 3: ReportBatchACK
	(*MatrixBatchRequest)(nil), // 4: MatrixBatchRequest
	(*Matrix)(nil),             // 5: Matrix
	(*MatrixBatchACK)(nil),     // 6: MatrixBatchACK
}
var file_liveconn_proto_depIdxs = []int32{
	1, // 0: ReportBatchRequest.report:type_name -> Report
	2, // 1: Report.drops:type_name -> Drop
	5, // 2: MatrixBatchRequest.matrix:type_name -> Matrix
	0, // 3: ConnectedLiveService.PushReportBatch:input_type -> ReportBatchRequest
	4, // 4: ConnectedLiveService.PushMatrixBatch:input_type -> MatrixBatchRequest
	3, // 5: ConnectedLiveService.PushReportBatch:output_type -> ReportBatchACK
	6, // 6: ConnectedLiveService.PushMatrixBatch:output_type -> MatrixBatchACK
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_liveconn_proto_init() }
func file_liveconn_proto_init() {
	if File_liveconn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_liveconn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportBatchRequest); i {
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
		file_liveconn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
		file_liveconn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Drop); i {
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
		file_liveconn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportBatchACK); i {
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
		file_liveconn_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatrixBatchRequest); i {
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
		file_liveconn_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Matrix); i {
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
		file_liveconn_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatrixBatchACK); i {
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
			RawDescriptor: file_liveconn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_liveconn_proto_goTypes,
		DependencyIndexes: file_liveconn_proto_depIdxs,
		MessageInfos:      file_liveconn_proto_msgTypes,
	}.Build()
	File_liveconn_proto = out.File
	file_liveconn_proto_rawDesc = nil
	file_liveconn_proto_goTypes = nil
	file_liveconn_proto_depIdxs = nil
}
