// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.9.1
// source: none.proto

package types

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NoneAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*NoneAction_CommitDelayTx
	Value isNoneAction_Value `protobuf_oneof:"value"`
	Ty    int32              `protobuf:"varint,2,opt,name=Ty,proto3" json:"Ty,omitempty"`
}

func (x *NoneAction) Reset() {
	*x = NoneAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_none_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoneAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoneAction) ProtoMessage() {}

func (x *NoneAction) ProtoReflect() protoreflect.Message {
	mi := &file_none_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoneAction.ProtoReflect.Descriptor instead.
func (*NoneAction) Descriptor() ([]byte, []int) {
	return file_none_proto_rawDescGZIP(), []int{0}
}

func (m *NoneAction) GetValue() isNoneAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *NoneAction) GetCommitDelayTx() *CommitDelayTx {
	if x, ok := x.GetValue().(*NoneAction_CommitDelayTx); ok {
		return x.CommitDelayTx
	}
	return nil
}

func (x *NoneAction) GetTy() int32 {
	if x != nil {
		return x.Ty
	}
	return 0
}

type isNoneAction_Value interface {
	isNoneAction_Value()
}

type NoneAction_CommitDelayTx struct {
	CommitDelayTx *CommitDelayTx `protobuf:"bytes,1,opt,name=commitDelayTx,proto3,oneof"`
}

func (*NoneAction_CommitDelayTx) isNoneAction_Value() {}

// 提交延时交易类型
type CommitDelayTx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DelayTx             string `protobuf:"bytes,1,opt,name=delayTx,proto3" json:"delayTx,omitempty"`                          //延时交易, 16进制格式
	RelativeDelayHeight int64  `protobuf:"varint,2,opt,name=relativeDelayHeight,proto3" json:"relativeDelayHeight,omitempty"` //相对延时时长，相对区块高度
}

func (x *CommitDelayTx) Reset() {
	*x = CommitDelayTx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_none_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitDelayTx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitDelayTx) ProtoMessage() {}

func (x *CommitDelayTx) ProtoReflect() protoreflect.Message {
	mi := &file_none_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitDelayTx.ProtoReflect.Descriptor instead.
func (*CommitDelayTx) Descriptor() ([]byte, []int) {
	return file_none_proto_rawDescGZIP(), []int{1}
}

func (x *CommitDelayTx) GetDelayTx() string {
	if x != nil {
		return x.DelayTx
	}
	return ""
}

func (x *CommitDelayTx) GetRelativeDelayHeight() int64 {
	if x != nil {
		return x.RelativeDelayHeight
	}
	return 0
}

// 提交延时交易回执
type CommitDelayTxLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Submitter        string `protobuf:"bytes,1,opt,name=submitter,proto3" json:"submitter,omitempty"`                // 提交者
	DelayTxHash      string `protobuf:"bytes,2,opt,name=delayTxHash,proto3" json:"delayTxHash,omitempty"`            // 延时交易哈希
	DelayBeginHeight int64  `protobuf:"varint,3,opt,name=delayBeginHeight,proto3" json:"delayBeginHeight,omitempty"` // 延时开始区块高度
}

func (x *CommitDelayTxLog) Reset() {
	*x = CommitDelayTxLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_none_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitDelayTxLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitDelayTxLog) ProtoMessage() {}

func (x *CommitDelayTxLog) ProtoReflect() protoreflect.Message {
	mi := &file_none_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitDelayTxLog.ProtoReflect.Descriptor instead.
func (*CommitDelayTxLog) Descriptor() ([]byte, []int) {
	return file_none_proto_rawDescGZIP(), []int{2}
}

func (x *CommitDelayTxLog) GetSubmitter() string {
	if x != nil {
		return x.Submitter
	}
	return ""
}

func (x *CommitDelayTxLog) GetDelayTxHash() string {
	if x != nil {
		return x.DelayTxHash
	}
	return ""
}

func (x *CommitDelayTxLog) GetDelayBeginHeight() int64 {
	if x != nil {
		return x.DelayBeginHeight
	}
	return 0
}

var File_none_proto protoreflect.FileDescriptor

var file_none_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x22, 0x63, 0x0a, 0x0a, 0x4e, 0x6f, 0x6e, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3c, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79,
	0x54, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x54, 0x78, 0x48, 0x00,
	0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x54, 0x78, 0x12,
	0x0e, 0x0a, 0x02, 0x54, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x54, 0x79, 0x42,
	0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5b, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x54, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c,
	0x61, 0x79, 0x54, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x61,
	0x79, 0x54, 0x78, 0x12, 0x30, 0x0a, 0x13, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x44,
	0x65, 0x6c, 0x61, 0x79, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x13, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x48,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x7e, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x44,
	0x65, 0x6c, 0x61, 0x79, 0x54, 0x78, 0x4c, 0x6f, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x61, 0x79,
	0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x6c, 0x61, 0x79, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x65, 0x6c,
	0x61, 0x79, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x48,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_none_proto_rawDescOnce sync.Once
	file_none_proto_rawDescData = file_none_proto_rawDesc
)

func file_none_proto_rawDescGZIP() []byte {
	file_none_proto_rawDescOnce.Do(func() {
		file_none_proto_rawDescData = protoimpl.X.CompressGZIP(file_none_proto_rawDescData)
	})
	return file_none_proto_rawDescData
}

var file_none_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_none_proto_goTypes = []interface{}{
	(*NoneAction)(nil),       // 0: types.NoneAction
	(*CommitDelayTx)(nil),    // 1: types.CommitDelayTx
	(*CommitDelayTxLog)(nil), // 2: types.CommitDelayTxLog
}
var file_none_proto_depIdxs = []int32{
	1, // 0: types.NoneAction.commitDelayTx:type_name -> types.CommitDelayTx
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_none_proto_init() }
func file_none_proto_init() {
	if File_none_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_none_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoneAction); i {
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
		file_none_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitDelayTx); i {
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
		file_none_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitDelayTxLog); i {
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
	file_none_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*NoneAction_CommitDelayTx)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_none_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_none_proto_goTypes,
		DependencyIndexes: file_none_proto_depIdxs,
		MessageInfos:      file_none_proto_msgTypes,
	}.Build()
	File_none_proto = out.File
	file_none_proto_rawDesc = nil
	file_none_proto_goTypes = nil
	file_none_proto_depIdxs = nil
}
