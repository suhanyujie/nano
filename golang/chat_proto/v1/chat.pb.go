// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: chat.proto

package v1

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

type ConnMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgType int32   `protobuf:"varint,1,opt,name=MsgType,proto3" json:"MsgType,omitempty"` // 消息类型：1 控制消息；2 应用消息；
	Sys     *SysMsg `protobuf:"bytes,2,opt,name=sys,proto3" json:"sys,omitempty"`          // 如果是控制消息，这该字段有值
	Data    *AppMsg `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`        // 如果是应用消息，则该字段有值
}

func (x *ConnMsg) Reset() {
	*x = ConnMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnMsg) ProtoMessage() {}

func (x *ConnMsg) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnMsg.ProtoReflect.Descriptor instead.
func (*ConnMsg) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ConnMsg) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *ConnMsg) GetSys() *SysMsg {
	if x != nil {
		return x.Sys
	}
	return nil
}

func (x *ConnMsg) GetData() *AppMsg {
	if x != nil {
		return x.Data
	}
	return nil
}

type SysMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ver  string `protobuf:"bytes,1,opt,name=ver,proto3" json:"ver,omitempty"`
	Typ  string `protobuf:"bytes,2,opt,name=typ,proto3" json:"typ,omitempty"`   // 控制类型
	User *User  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"` // 当前连接对应的用户信息
}

func (x *SysMsg) Reset() {
	*x = SysMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysMsg) ProtoMessage() {}

func (x *SysMsg) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysMsg.ProtoReflect.Descriptor instead.
func (*SysMsg) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *SysMsg) GetVer() string {
	if x != nil {
		return x.Ver
	}
	return ""
}

func (x *SysMsg) GetTyp() string {
	if x != nil {
		return x.Typ
	}
	return ""
}

func (x *SysMsg) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AppMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 应用消息类型。发文本消息、进入房间、投掷标枪、退出房间、投降
	// 每种类型都有对应的消息体
	Action string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	// Types that are assignable to Data:
	//	*AppMsg_ChatMsg
	//	*AppMsg_JoinUser
	Data isAppMsg_Data `protobuf_oneof:"data"`
}

func (x *AppMsg) Reset() {
	*x = AppMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppMsg) ProtoMessage() {}

func (x *AppMsg) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppMsg.ProtoReflect.Descriptor instead.
func (*AppMsg) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *AppMsg) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (m *AppMsg) GetData() isAppMsg_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *AppMsg) GetChatMsg() *AppMsgChat {
	if x, ok := x.GetData().(*AppMsg_ChatMsg); ok {
		return x.ChatMsg
	}
	return nil
}

func (x *AppMsg) GetJoinUser() *AppMsgUserJoin {
	if x, ok := x.GetData().(*AppMsg_JoinUser); ok {
		return x.JoinUser
	}
	return nil
}

type isAppMsg_Data interface {
	isAppMsg_Data()
}

type AppMsg_ChatMsg struct {
	ChatMsg *AppMsgChat `protobuf:"bytes,2,opt,name=chatMsg,proto3,oneof"` // 应用数据：用户发消息聊天
}

type AppMsg_JoinUser struct {
	JoinUser *AppMsgUserJoin `protobuf:"bytes,3,opt,name=joinUser,proto3,oneof"` // 应用数据：投资
}

func (*AppMsg_ChatMsg) isAppMsg_Data() {}

func (*AppMsg_JoinUser) isAppMsg_Data() {}

type AppMsgChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetUid int64  `protobuf:"varint,1,opt,name=targetUid,proto3" json:"targetUid,omitempty"` // -1 表示所有人
	Content   string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`      // 聊天内容
}

func (x *AppMsgChat) Reset() {
	*x = AppMsgChat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppMsgChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppMsgChat) ProtoMessage() {}

func (x *AppMsgChat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppMsgChat.ProtoReflect.Descriptor instead.
func (*AppMsgChat) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{4}
}

func (x *AppMsgChat) GetTargetUid() int64 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

func (x *AppMsgChat) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type AppMsgUserJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AppMsgUserJoin) Reset() {
	*x = AppMsgUserJoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppMsgUserJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppMsgUserJoin) ProtoMessage() {}

func (x *AppMsgUserJoin) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppMsgUserJoin.ProtoReflect.Descriptor instead.
func (*AppMsgUserJoin) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{5}
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x19, 0x0a, 0x03, 0x73, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x53, 0x79, 0x73, 0x4d, 0x73, 0x67, 0x52, 0x03, 0x73, 0x79, 0x73, 0x12, 0x1b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x41, 0x70, 0x70,
	0x4d, 0x73, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x47, 0x0a, 0x06, 0x53, 0x79, 0x73,
	0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x76, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x79, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x74, 0x79, 0x70, 0x12, 0x19, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x1c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x80, 0x01, 0x0a, 0x06, 0x41, 0x70, 0x70, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x70, 0x70, 0x4d, 0x73, 0x67, 0x43, 0x68, 0x61,
	0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x2d, 0x0a, 0x08,
	0x6a, 0x6f, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x41, 0x70, 0x70, 0x4d, 0x73, 0x67, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x48,
	0x00, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x44, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x4d, 0x73, 0x67, 0x43, 0x68, 0x61,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x41, 0x70, 0x70,
	0x4d, 0x73, 0x67, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x6f, 0x69, 0x6e, 0x42, 0x2e, 0x5a, 0x2c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6f, 0x6e, 0x6e, 0x67, 0x2f,
	0x6e, 0x61, 0x6e, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_chat_proto_goTypes = []interface{}{
	(*ConnMsg)(nil),        // 0: ConnMsg
	(*SysMsg)(nil),         // 1: SysMsg
	(*User)(nil),           // 2: User
	(*AppMsg)(nil),         // 3: AppMsg
	(*AppMsgChat)(nil),     // 4: AppMsgChat
	(*AppMsgUserJoin)(nil), // 5: AppMsgUserJoin
}
var file_chat_proto_depIdxs = []int32{
	1, // 0: ConnMsg.sys:type_name -> SysMsg
	3, // 1: ConnMsg.data:type_name -> AppMsg
	2, // 2: SysMsg.user:type_name -> User
	4, // 3: AppMsg.chatMsg:type_name -> AppMsgChat
	5, // 4: AppMsg.joinUser:type_name -> AppMsgUserJoin
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnMsg); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysMsg); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppMsg); i {
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
		file_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppMsgChat); i {
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
		file_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppMsgUserJoin); i {
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
	file_chat_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*AppMsg_ChatMsg)(nil),
		(*AppMsg_JoinUser)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
