// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type Query_Type int32

const (
	Query_MATCH  Query_Type = 0
	Query_PHRASE Query_Type = 1
)

// Enum value maps for Query_Type.
var (
	Query_Type_name = map[int32]string{
		0: "MATCH",
		1: "PHRASE",
	}
	Query_Type_value = map[string]int32{
		"MATCH":  0,
		"PHRASE": 1,
	}
)

func (x Query_Type) Enum() *Query_Type {
	p := new(Query_Type)
	*p = x
	return p
}

func (x Query_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Query_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (Query_Type) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x Query_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Query_Type.Descriptor instead.
func (Query_Type) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1, 0}
}

// Document represents an indexed document.
type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LinkId    []byte                 `protobuf:"bytes,1,opt,name=link_id,json=linkId,proto3" json:"link_id,omitempty"`
	Url       string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Title     string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content   string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	IndexedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=indexed_at,json=indexedAt,proto3" json:"indexed_at,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *Document) GetLinkId() []byte {
	if x != nil {
		return x.LinkId
	}
	return nil
}

func (x *Document) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Document) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Document) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Document) GetIndexedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.IndexedAt
	}
	return nil
}

// Query represents a search query.
type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       Query_Type `protobuf:"varint,1,opt,name=type,proto3,enum=proto.Query_Type" json:"type,omitempty"`
	Expression string     `protobuf:"bytes,2,opt,name=expression,proto3" json:"expression,omitempty"`
	Offset     uint64     `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *Query) GetType() Query_Type {
	if x != nil {
		return x.Type
	}
	return Query_MATCH
}

func (x *Query) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

func (x *Query) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// QueryResult contains either the total count of results for a query or a
// single document from the resultset.
type QueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//	*QueryResult_DocCount
	//	*QueryResult_Doc
	Result isQueryResult_Result `protobuf_oneof:"result"`
}

func (x *QueryResult) Reset() {
	*x = QueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResult) ProtoMessage() {}

func (x *QueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResult.ProtoReflect.Descriptor instead.
func (*QueryResult) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (m *QueryResult) GetResult() isQueryResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *QueryResult) GetDocCount() uint64 {
	if x, ok := x.GetResult().(*QueryResult_DocCount); ok {
		return x.DocCount
	}
	return 0
}

func (x *QueryResult) GetDoc() *Document {
	if x, ok := x.GetResult().(*QueryResult_Doc); ok {
		return x.Doc
	}
	return nil
}

type isQueryResult_Result interface {
	isQueryResult_Result()
}

type QueryResult_DocCount struct {
	DocCount uint64 `protobuf:"varint,1,opt,name=doc_count,json=docCount,proto3,oneof"`
}

type QueryResult_Doc struct {
	Doc *Document `protobuf:"bytes,2,opt,name=doc,proto3,oneof"`
}

func (*QueryResult_DocCount) isQueryResult_Result() {}

func (*QueryResult_Doc) isQueryResult_Result() {}

// UpdateScoreRequest encapsulates the parameters for the UpdateScore RPC.
type UpdateScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LinkId        []byte  `protobuf:"bytes,1,opt,name=link_id,json=linkId,proto3" json:"link_id,omitempty"`
	PageRankScore float64 `protobuf:"fixed64,2,opt,name=page_rank_score,json=pageRankScore,proto3" json:"page_rank_score,omitempty"`
}

func (x *UpdateScoreRequest) Reset() {
	*x = UpdateScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScoreRequest) ProtoMessage() {}

func (x *UpdateScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScoreRequest.ProtoReflect.Descriptor instead.
func (*UpdateScoreRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateScoreRequest) GetLinkId() []byte {
	if x != nil {
		return x.LinkId
	}
	return nil
}

func (x *UpdateScoreRequest) GetPageRankScore() float64 {
	if x != nil {
		return x.PageRankScore
	}
	return 0
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa0, 0x01, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x6c, 0x69, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x25, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x1d, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x50, 0x48, 0x52, 0x41, 0x53, 0x45, 0x10, 0x01, 0x22, 0x5b, 0x0a, 0x0b, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x64, 0x6f,
	0x63, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52,
	0x08, 0x64, 0x6f, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x64, 0x6f, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x03, 0x64, 0x6f, 0x63, 0x42, 0x08,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x55, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x6c, 0x69, 0x6e, 0x6b, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x70, 0x61, 0x67, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x32,
	0xa8, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x12,
	0x29, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x30, 0x01, 0x12, 0x40, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_proto_goTypes = []interface{}{
	(Query_Type)(0),               // 0: proto.Query.Type
	(*Document)(nil),              // 1: proto.Document
	(*Query)(nil),                 // 2: proto.Query
	(*QueryResult)(nil),           // 3: proto.QueryResult
	(*UpdateScoreRequest)(nil),    // 4: proto.UpdateScoreRequest
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_api_proto_depIdxs = []int32{
	5, // 0: proto.Document.indexed_at:type_name -> google.protobuf.Timestamp
	0, // 1: proto.Query.type:type_name -> proto.Query.Type
	1, // 2: proto.QueryResult.doc:type_name -> proto.Document
	1, // 3: proto.TextIndexer.Index:input_type -> proto.Document
	2, // 4: proto.TextIndexer.Search:input_type -> proto.Query
	4, // 5: proto.TextIndexer.UpdateScore:input_type -> proto.UpdateScoreRequest
	1, // 6: proto.TextIndexer.Index:output_type -> proto.Document
	3, // 7: proto.TextIndexer.Search:output_type -> proto.QueryResult
	6, // 8: proto.TextIndexer.UpdateScore:output_type -> google.protobuf.Empty
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResult); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateScoreRequest); i {
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
	file_api_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*QueryResult_DocCount)(nil),
		(*QueryResult_Doc)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TextIndexerClient is the client API for TextIndexer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TextIndexerClient interface {
	// Index inserts a new document to the index or updates the index entry for
	// and existing document.
	Index(ctx context.Context, in *Document, opts ...grpc.CallOption) (*Document, error)
	// Search the index for a particular query and stream the results back to
	// the client. The first response will include the total result count while
	// all subsequent responses will include documents from the resultset.
	Search(ctx context.Context, in *Query, opts ...grpc.CallOption) (TextIndexer_SearchClient, error)
	// UpdateScore updates the PageRank score for a document with the specified
	// link ID.
	UpdateScore(ctx context.Context, in *UpdateScoreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type textIndexerClient struct {
	cc grpc.ClientConnInterface
}

func NewTextIndexerClient(cc grpc.ClientConnInterface) TextIndexerClient {
	return &textIndexerClient{cc}
}

func (c *textIndexerClient) Index(ctx context.Context, in *Document, opts ...grpc.CallOption) (*Document, error) {
	out := new(Document)
	err := c.cc.Invoke(ctx, "/proto.TextIndexer/Index", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textIndexerClient) Search(ctx context.Context, in *Query, opts ...grpc.CallOption) (TextIndexer_SearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TextIndexer_serviceDesc.Streams[0], "/proto.TextIndexer/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &textIndexerSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TextIndexer_SearchClient interface {
	Recv() (*QueryResult, error)
	grpc.ClientStream
}

type textIndexerSearchClient struct {
	grpc.ClientStream
}

func (x *textIndexerSearchClient) Recv() (*QueryResult, error) {
	m := new(QueryResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *textIndexerClient) UpdateScore(ctx context.Context, in *UpdateScoreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.TextIndexer/UpdateScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TextIndexerServer is the server API for TextIndexer service.
type TextIndexerServer interface {
	// Index inserts a new document to the index or updates the index entry for
	// and existing document.
	Index(context.Context, *Document) (*Document, error)
	// Search the index for a particular query and stream the results back to
	// the client. The first response will include the total result count while
	// all subsequent responses will include documents from the resultset.
	Search(*Query, TextIndexer_SearchServer) error
	// UpdateScore updates the PageRank score for a document with the specified
	// link ID.
	UpdateScore(context.Context, *UpdateScoreRequest) (*emptypb.Empty, error)
}

// UnimplementedTextIndexerServer can be embedded to have forward compatible implementations.
type UnimplementedTextIndexerServer struct {
}

func (*UnimplementedTextIndexerServer) Index(context.Context, *Document) (*Document, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Index not implemented")
}
func (*UnimplementedTextIndexerServer) Search(*Query, TextIndexer_SearchServer) error {
	return status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedTextIndexerServer) UpdateScore(context.Context, *UpdateScoreRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScore not implemented")
}

func RegisterTextIndexerServer(s *grpc.Server, srv TextIndexerServer) {
	s.RegisterService(&_TextIndexer_serviceDesc, srv)
}

func _TextIndexer_Index_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Document)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextIndexerServer).Index(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TextIndexer/Index",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextIndexerServer).Index(ctx, req.(*Document))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextIndexer_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Query)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TextIndexerServer).Search(m, &textIndexerSearchServer{stream})
}

type TextIndexer_SearchServer interface {
	Send(*QueryResult) error
	grpc.ServerStream
}

type textIndexerSearchServer struct {
	grpc.ServerStream
}

func (x *textIndexerSearchServer) Send(m *QueryResult) error {
	return x.ServerStream.SendMsg(m)
}

func _TextIndexer_UpdateScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextIndexerServer).UpdateScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TextIndexer/UpdateScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextIndexerServer).UpdateScore(ctx, req.(*UpdateScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TextIndexer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TextIndexer",
	HandlerType: (*TextIndexerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Index",
			Handler:    _TextIndexer_Index_Handler,
		},
		{
			MethodName: "UpdateScore",
			Handler:    _TextIndexer_UpdateScore_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Search",
			Handler:       _TextIndexer_Search_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}
