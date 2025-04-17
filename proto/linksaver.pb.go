// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/linksaver.proto

package linksaver

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SaveRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SaveRequest) Reset() {
	*x = SaveRequest{}
	mi := &file_proto_linksaver_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRequest) ProtoMessage() {}

func (x *SaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_linksaver_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveRequest.ProtoReflect.Descriptor instead.
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return file_proto_linksaver_proto_rawDescGZIP(), []int{0}
}

func (x *SaveRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SaveRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type SaveResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SaveResponse) Reset() {
	*x = SaveResponse{}
	mi := &file_proto_linksaver_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveResponse) ProtoMessage() {}

func (x *SaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_linksaver_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveResponse.ProtoReflect.Descriptor instead.
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return file_proto_linksaver_proto_rawDescGZIP(), []int{1}
}

func (x *SaveResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_proto_linksaver_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_linksaver_proto_msgTypes[2]
	if x != nil {
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
	return file_proto_linksaver_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type Page struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Page) Reset() {
	*x = Page{}
	mi := &file_proto_linksaver_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_proto_linksaver_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_proto_linksaver_proto_rawDescGZIP(), []int{3}
}

func (x *Page) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Page) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type PageList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pages         []*Page                `protobuf:"bytes,1,rep,name=pages,proto3" json:"pages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PageList) Reset() {
	*x = PageList{}
	mi := &file_proto_linksaver_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PageList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageList) ProtoMessage() {}

func (x *PageList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_linksaver_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageList.ProtoReflect.Descriptor instead.
func (*PageList) Descriptor() ([]byte, []int) {
	return file_proto_linksaver_proto_rawDescGZIP(), []int{4}
}

func (x *PageList) GetPages() []*Page {
	if x != nil {
		return x.Pages
	}
	return nil
}

type ExistsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Exists        bool                   `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExistsResponse) Reset() {
	*x = ExistsResponse{}
	mi := &file_proto_linksaver_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistsResponse) ProtoMessage() {}

func (x *ExistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_linksaver_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistsResponse.ProtoReflect.Descriptor instead.
func (*ExistsResponse) Descriptor() ([]byte, []int) {
	return file_proto_linksaver_proto_rawDescGZIP(), []int{5}
}

func (x *ExistsResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

var File_proto_linksaver_proto protoreflect.FileDescriptor

const file_proto_linksaver_proto_rawDesc = "" +
	"\n" +
	"\x15proto/linksaver.proto\x12\tlinksaver\";\n" +
	"\vSaveRequest\x12\x10\n" +
	"\x03url\x18\x01 \x01(\tR\x03url\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\"&\n" +
	"\fSaveResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\"\"\n" +
	"\x04User\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\"4\n" +
	"\x04Page\x12\x10\n" +
	"\x03url\x18\x01 \x01(\tR\x03url\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\"1\n" +
	"\bPageList\x12%\n" +
	"\x05pages\x18\x01 \x03(\v2\x0f.linksaver.PageR\x05pages\"(\n" +
	"\x0eExistsResponse\x12\x16\n" +
	"\x06exists\x18\x01 \x01(\bR\x06exists2\xdd\x01\n" +
	"\tLinkSaver\x127\n" +
	"\x04Save\x12\x16.linksaver.SaveRequest\x1a\x17.linksaver.SaveResponse\x12.\n" +
	"\n" +
	"PickRandom\x12\x0f.linksaver.User\x1a\x0f.linksaver.Page\x12/\n" +
	"\aPickAll\x12\x0f.linksaver.User\x1a\x13.linksaver.PageList\x126\n" +
	"\bIsExists\x12\x0f.linksaver.Page\x1a\x19.linksaver.ExistsResponseB.Z,github.com/Gipohub/linksaver/proto;linksaverb\x06proto3"

var (
	file_proto_linksaver_proto_rawDescOnce sync.Once
	file_proto_linksaver_proto_rawDescData []byte
)

func file_proto_linksaver_proto_rawDescGZIP() []byte {
	file_proto_linksaver_proto_rawDescOnce.Do(func() {
		file_proto_linksaver_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_linksaver_proto_rawDesc), len(file_proto_linksaver_proto_rawDesc)))
	})
	return file_proto_linksaver_proto_rawDescData
}

var file_proto_linksaver_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_linksaver_proto_goTypes = []any{
	(*SaveRequest)(nil),    // 0: linksaver.SaveRequest
	(*SaveResponse)(nil),   // 1: linksaver.SaveResponse
	(*User)(nil),           // 2: linksaver.User
	(*Page)(nil),           // 3: linksaver.Page
	(*PageList)(nil),       // 4: linksaver.PageList
	(*ExistsResponse)(nil), // 5: linksaver.ExistsResponse
}
var file_proto_linksaver_proto_depIdxs = []int32{
	3, // 0: linksaver.PageList.pages:type_name -> linksaver.Page
	0, // 1: linksaver.LinkSaver.Save:input_type -> linksaver.SaveRequest
	2, // 2: linksaver.LinkSaver.PickRandom:input_type -> linksaver.User
	2, // 3: linksaver.LinkSaver.PickAll:input_type -> linksaver.User
	3, // 4: linksaver.LinkSaver.IsExists:input_type -> linksaver.Page
	1, // 5: linksaver.LinkSaver.Save:output_type -> linksaver.SaveResponse
	3, // 6: linksaver.LinkSaver.PickRandom:output_type -> linksaver.Page
	4, // 7: linksaver.LinkSaver.PickAll:output_type -> linksaver.PageList
	5, // 8: linksaver.LinkSaver.IsExists:output_type -> linksaver.ExistsResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_linksaver_proto_init() }
func file_proto_linksaver_proto_init() {
	if File_proto_linksaver_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_linksaver_proto_rawDesc), len(file_proto_linksaver_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_linksaver_proto_goTypes,
		DependencyIndexes: file_proto_linksaver_proto_depIdxs,
		MessageInfos:      file_proto_linksaver_proto_msgTypes,
	}.Build()
	File_proto_linksaver_proto = out.File
	file_proto_linksaver_proto_goTypes = nil
	file_proto_linksaver_proto_depIdxs = nil
}
