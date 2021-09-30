// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: mpc.proto

package types

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

type MpcPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation string `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"` //操作
	NodeName  string `protobuf:"bytes,2,opt,name=nodeName,proto3" json:"nodeName,omitempty"`   //节点名（地址）
	NodeAddr  string `protobuf:"bytes,3,opt,name=nodeAddr,proto3" json:"nodeAddr,omitempty"`   //节点rpc地址
	TaskId    string `protobuf:"bytes,4,opt,name=taskId,proto3" json:"taskId,omitempty"`       //task唯一标识 id
	TaskType  string `protobuf:"bytes,5,opt,name=taskType,proto3" json:"taskType,omitempty"`   //task类型
	ExecNode  string `protobuf:"bytes,6,opt,name=execNode,proto3" json:"execNode,omitempty"`   //task执行节点
	Result    []byte `protobuf:"bytes,7,opt,name=result,proto3" json:"result,omitempty"`       //计算结果
	DataName  string `protobuf:"bytes,8,opt,name=dataName,proto3" json:"dataName,omitempty"`   //数据名
	DataType  string `protobuf:"bytes,9,opt,name=dataType,proto3" json:"dataType,omitempty"`   //数据类型
	DataHash  string `protobuf:"bytes,10,opt,name=dataHash,proto3" json:"dataHash,omitempty"`  //数据hash
}

func (x *MpcPayload) Reset() {
	*x = MpcPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpcPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpcPayload) ProtoMessage() {}

func (x *MpcPayload) ProtoReflect() protoreflect.Message {
	mi := &file_mpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpcPayload.ProtoReflect.Descriptor instead.
func (*MpcPayload) Descriptor() ([]byte, []int) {
	return file_mpc_proto_rawDescGZIP(), []int{0}
}

func (x *MpcPayload) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *MpcPayload) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *MpcPayload) GetNodeAddr() string {
	if x != nil {
		return x.NodeAddr
	}
	return ""
}

func (x *MpcPayload) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *MpcPayload) GetTaskType() string {
	if x != nil {
		return x.TaskType
	}
	return ""
}

func (x *MpcPayload) GetExecNode() string {
	if x != nil {
		return x.ExecNode
	}
	return ""
}

func (x *MpcPayload) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *MpcPayload) GetDataName() string {
	if x != nil {
		return x.DataName
	}
	return ""
}

func (x *MpcPayload) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *MpcPayload) GetDataHash() string {
	if x != nil {
		return x.DataHash
	}
	return ""
}

type MpcPayloadList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MpcPayload []*MpcPayload `protobuf:"bytes,1,rep,name=mpcPayload,proto3" json:"mpcPayload,omitempty"`
}

func (x *MpcPayloadList) Reset() {
	*x = MpcPayloadList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpcPayloadList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpcPayloadList) ProtoMessage() {}

func (x *MpcPayloadList) ProtoReflect() protoreflect.Message {
	mi := &file_mpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpcPayloadList.ProtoReflect.Descriptor instead.
func (*MpcPayloadList) Descriptor() ([]byte, []int) {
	return file_mpc_proto_rawDescGZIP(), []int{1}
}

func (x *MpcPayloadList) GetMpcPayload() []*MpcPayload {
	if x != nil {
		return x.MpcPayload
	}
	return nil
}

var File_mpc_proto protoreflect.FileDescriptor

var file_mpc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x22, 0x9e, 0x02, 0x0a, 0x0a, 0x4d, 0x70, 0x63, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65,
	0x78, 0x65, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x78, 0x65, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x48,
	0x61, 0x73, 0x68, 0x22, 0x43, 0x0a, 0x0e, 0x4d, 0x70, 0x63, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0a, 0x6d, 0x70, 0x63, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x4d, 0x70, 0x63, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0a, 0x6d, 0x70,
	0x63, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x33, 0x33, 0x63, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x33, 0x33, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mpc_proto_rawDescOnce sync.Once
	file_mpc_proto_rawDescData = file_mpc_proto_rawDesc
)

func file_mpc_proto_rawDescGZIP() []byte {
	file_mpc_proto_rawDescOnce.Do(func() {
		file_mpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mpc_proto_rawDescData)
	})
	return file_mpc_proto_rawDescData
}

var file_mpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mpc_proto_goTypes = []interface{}{
	(*MpcPayload)(nil),     // 0: types.MpcPayload
	(*MpcPayloadList)(nil), // 1: types.MpcPayloadList
}
var file_mpc_proto_depIdxs = []int32{
	0, // 0: types.MpcPayloadList.mpcPayload:type_name -> types.MpcPayload
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mpc_proto_init() }
func file_mpc_proto_init() {
	if File_mpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpcPayload); i {
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
		file_mpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpcPayloadList); i {
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
			RawDescriptor: file_mpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mpc_proto_goTypes,
		DependencyIndexes: file_mpc_proto_depIdxs,
		MessageInfos:      file_mpc_proto_msgTypes,
	}.Build()
	File_mpc_proto = out.File
	file_mpc_proto_rawDesc = nil
	file_mpc_proto_goTypes = nil
	file_mpc_proto_depIdxs = nil
}
