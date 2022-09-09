// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: edge-service.proto

package api

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

type SetDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	NumberOfPeople int64                  `protobuf:"varint,2,opt,name=number_of_people,json=numberOfPeople,proto3" json:"number_of_people,omitempty"`
	CameraImage    []byte                 `protobuf:"bytes,3,opt,name=camera_image,json=cameraImage,proto3" json:"camera_image,omitempty"`
}

func (x *SetDataRequest) Reset() {
	*x = SetDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_edge_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDataRequest) ProtoMessage() {}

func (x *SetDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_edge_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDataRequest.ProtoReflect.Descriptor instead.
func (*SetDataRequest) Descriptor() ([]byte, []int) {
	return file_edge_service_proto_rawDescGZIP(), []int{0}
}

func (x *SetDataRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *SetDataRequest) GetNumberOfPeople() int64 {
	if x != nil {
		return x.NumberOfPeople
	}
	return 0
}

func (x *SetDataRequest) GetCameraImage() []byte {
	if x != nil {
		return x.CameraImage
	}
	return nil
}

var File_edge_service_proto protoreflect.FileDescriptor

var file_edge_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x64, 0x67, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x75, 0x74, 0x69, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f,
	0x66, 0x5f, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x32, 0x62, 0x0a, 0x0b, 0x45, 0x64, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x25, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x0a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x72, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x2d, 0x6d, 0x6f, 0x6e,
	0x6b, 0x65, 0x79, 0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_edge_service_proto_rawDescOnce sync.Once
	file_edge_service_proto_rawDescData = file_edge_service_proto_rawDesc
)

func file_edge_service_proto_rawDescGZIP() []byte {
	file_edge_service_proto_rawDescOnce.Do(func() {
		file_edge_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_edge_service_proto_rawDescData)
	})
	return file_edge_service_proto_rawDescData
}

var file_edge_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_edge_service_proto_goTypes = []interface{}{
	(*SetDataRequest)(nil),        // 0: api.SetDataRequest
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*Empty)(nil),                 // 2: api.Empty
}
var file_edge_service_proto_depIdxs = []int32{
	1, // 0: api.SetDataRequest.timestamp:type_name -> google.protobuf.Timestamp
	2, // 1: api.EdgeService.Heartbeat:input_type -> api.Empty
	0, // 2: api.EdgeService.SetData:input_type -> api.SetDataRequest
	2, // 3: api.EdgeService.Heartbeat:output_type -> api.Empty
	2, // 4: api.EdgeService.SetData:output_type -> api.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_edge_service_proto_init() }
func file_edge_service_proto_init() {
	if File_edge_service_proto != nil {
		return
	}
	file_utils_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_edge_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDataRequest); i {
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
			RawDescriptor: file_edge_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_edge_service_proto_goTypes,
		DependencyIndexes: file_edge_service_proto_depIdxs,
		MessageInfos:      file_edge_service_proto_msgTypes,
	}.Build()
	File_edge_service_proto = out.File
	file_edge_service_proto_rawDesc = nil
	file_edge_service_proto_goTypes = nil
	file_edge_service_proto_depIdxs = nil
}
