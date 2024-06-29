// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: metro_service.proto

package metro_service

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

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metro_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_metro_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_metro_service_proto_rawDescGZIP(), []int{0}
}

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UserId string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metro_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_metro_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_metro_service_proto_rawDescGZIP(), []int{1}
}

func (x *Card) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Card) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Card) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Station struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Station) Reset() {
	*x = Station{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metro_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Station) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Station) ProtoMessage() {}

func (x *Station) ProtoReflect() protoreflect.Message {
	mi := &file_metro_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Station.ProtoReflect.Descriptor instead.
func (*Station) Descriptor() ([]byte, []int) {
	return file_metro_service_proto_rawDescGZIP(), []int{2}
}

func (x *Station) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Station) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Terminal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StationId string `protobuf:"bytes,2,opt,name=stationId,proto3" json:"stationId,omitempty"`
}

func (x *Terminal) Reset() {
	*x = Terminal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metro_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Terminal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Terminal) ProtoMessage() {}

func (x *Terminal) ProtoReflect() protoreflect.Message {
	mi := &file_metro_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Terminal.ProtoReflect.Descriptor instead.
func (*Terminal) Descriptor() ([]byte, []int) {
	return file_metro_service_proto_rawDescGZIP(), []int{3}
}

func (x *Terminal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Terminal) GetStationId() string {
	if x != nil {
		return x.StationId
	}
	return ""
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CardId     string `protobuf:"bytes,2,opt,name=cardId,proto3" json:"cardId,omitempty"`
	Amount     string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	TerminalId string `protobuf:"bytes,4,opt,name=terminalId,proto3" json:"terminalId,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metro_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_metro_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_metro_service_proto_rawDescGZIP(), []int{4}
}

func (x *Transaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transaction) GetCardId() string {
	if x != nil {
		return x.CardId
	}
	return ""
}

func (x *Transaction) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Transaction) GetTerminalId() string {
	if x != nil {
		return x.TerminalId
	}
	return ""
}

var File_metro_service_proto protoreflect.FileDescriptor

var file_metro_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x04,
	0x43, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x2d, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x38, 0x0a, 0x08, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x6d, 0x0a, 0x0b, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x64,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x64, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x32, 0xff, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x74,
	0x72, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x43, 0x61,
	0x72, 0x64, 0x73, 0x12, 0x13, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x1a, 0x13, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x6f,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12,
	0x39, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x6d, 0x65,
	0x74, 0x72, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x09, 0x54, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x73, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c,
	0x1a, 0x13, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metro_service_proto_rawDescOnce sync.Once
	file_metro_service_proto_rawDescData = file_metro_service_proto_rawDesc
)

func file_metro_service_proto_rawDescGZIP() []byte {
	file_metro_service_proto_rawDescOnce.Do(func() {
		file_metro_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_metro_service_proto_rawDescData)
	})
	return file_metro_service_proto_rawDescData
}

var file_metro_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_metro_service_proto_goTypes = []interface{}{
	(*Void)(nil),        // 0: metro_service.Void
	(*Card)(nil),        // 1: metro_service.Card
	(*Station)(nil),     // 2: metro_service.Station
	(*Terminal)(nil),    // 3: metro_service.Terminal
	(*Transaction)(nil), // 4: metro_service.Transaction
}
var file_metro_service_proto_depIdxs = []int32{
	1, // 0: metro_service.Metro_service.Cards:input_type -> metro_service.Card
	2, // 1: metro_service.Metro_service.Stations:input_type -> metro_service.Station
	3, // 2: metro_service.Metro_service.Terminals:input_type -> metro_service.Terminal
	4, // 3: metro_service.Metro_service.Transactions:input_type -> metro_service.Transaction
	0, // 4: metro_service.Metro_service.Cards:output_type -> metro_service.Void
	0, // 5: metro_service.Metro_service.Stations:output_type -> metro_service.Void
	0, // 6: metro_service.Metro_service.Terminals:output_type -> metro_service.Void
	0, // 7: metro_service.Metro_service.Transactions:output_type -> metro_service.Void
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_metro_service_proto_init() }
func file_metro_service_proto_init() {
	if File_metro_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metro_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_metro_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_metro_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Station); i {
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
		file_metro_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Terminal); i {
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
		file_metro_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
			RawDescriptor: file_metro_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metro_service_proto_goTypes,
		DependencyIndexes: file_metro_service_proto_depIdxs,
		MessageInfos:      file_metro_service_proto_msgTypes,
	}.Build()
	File_metro_service_proto = out.File
	file_metro_service_proto_rawDesc = nil
	file_metro_service_proto_goTypes = nil
	file_metro_service_proto_depIdxs = nil
}
