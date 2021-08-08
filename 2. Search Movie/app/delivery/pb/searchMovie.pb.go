// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: searchMovie.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_searchMovie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_searchMovie_proto_msgTypes[0]
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
	return file_searchMovie_proto_rawDescGZIP(), []int{0}
}

type SearchMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Searchword string `protobuf:"bytes,1,opt,name=searchword,proto3" json:"searchword,omitempty"`
	Pagination string `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *SearchMovieRequest) Reset() {
	*x = SearchMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_searchMovie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMovieRequest) ProtoMessage() {}

func (x *SearchMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_searchMovie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMovieRequest.ProtoReflect.Descriptor instead.
func (*SearchMovieRequest) Descriptor() ([]byte, []int) {
	return file_searchMovie_proto_rawDescGZIP(), []int{1}
}

func (x *SearchMovieRequest) GetSearchword() string {
	if x != nil {
		return x.Searchword
	}
	return ""
}

func (x *SearchMovieRequest) GetPagination() string {
	if x != nil {
		return x.Pagination
	}
	return ""
}

type Search struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=Year,proto3" json:"Year,omitempty"`
	ImdbID string `protobuf:"bytes,3,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=Poster,proto3" json:"Poster,omitempty"`
}

func (x *Search) Reset() {
	*x = Search{}
	if protoimpl.UnsafeEnabled {
		mi := &file_searchMovie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Search) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Search) ProtoMessage() {}

func (x *Search) ProtoReflect() protoreflect.Message {
	mi := &file_searchMovie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Search.ProtoReflect.Descriptor instead.
func (*Search) Descriptor() ([]byte, []int) {
	return file_searchMovie_proto_rawDescGZIP(), []int{2}
}

func (x *Search) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Search) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Search) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *Search) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Search) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type OMDBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search       []*Search `protobuf:"bytes,1,rep,name=Search,proto3" json:"Search,omitempty"`
	TotalResults string    `protobuf:"bytes,2,opt,name=TotalResults,proto3" json:"TotalResults,omitempty"`
	Response     string    `protobuf:"bytes,3,opt,name=Response,proto3" json:"Response,omitempty"`
	Error        string    `protobuf:"bytes,4,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *OMDBResponse) Reset() {
	*x = OMDBResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_searchMovie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OMDBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OMDBResponse) ProtoMessage() {}

func (x *OMDBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_searchMovie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OMDBResponse.ProtoReflect.Descriptor instead.
func (*OMDBResponse) Descriptor() ([]byte, []int) {
	return file_searchMovie_proto_rawDescGZIP(), []int{3}
}

func (x *OMDBResponse) GetSearch() []*Search {
	if x != nil {
		return x.Search
	}
	return nil
}

func (x *OMDBResponse) GetTotalResults() string {
	if x != nil {
		return x.TotalResults
	}
	return ""
}

func (x *OMDBResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *OMDBResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_searchMovie_proto protoreflect.FileDescriptor

var file_searchMovie_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x54, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x76, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d,
	0x64, 0x62, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x22, 0x88,
	0x01, 0x0a, 0x0c, 0x4f, 0x4d, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x06, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x4a, 0x0a, 0x0d, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x4d, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_searchMovie_proto_rawDescOnce sync.Once
	file_searchMovie_proto_rawDescData = file_searchMovie_proto_rawDesc
)

func file_searchMovie_proto_rawDescGZIP() []byte {
	file_searchMovie_proto_rawDescOnce.Do(func() {
		file_searchMovie_proto_rawDescData = protoimpl.X.CompressGZIP(file_searchMovie_proto_rawDescData)
	})
	return file_searchMovie_proto_rawDescData
}

var file_searchMovie_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_searchMovie_proto_goTypes = []interface{}{
	(*Empty)(nil),              // 0: pb.Empty
	(*SearchMovieRequest)(nil), // 1: pb.SearchMovieRequest
	(*Search)(nil),             // 2: pb.Search
	(*OMDBResponse)(nil),       // 3: pb.OMDBResponse
}
var file_searchMovie_proto_depIdxs = []int32{
	2, // 0: pb.OMDBResponse.Search:type_name -> pb.Search
	1, // 1: pb.SearchService.SearchMovie:input_type -> pb.SearchMovieRequest
	3, // 2: pb.SearchService.SearchMovie:output_type -> pb.OMDBResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_searchMovie_proto_init() }
func file_searchMovie_proto_init() {
	if File_searchMovie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_searchMovie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_searchMovie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMovieRequest); i {
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
		file_searchMovie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Search); i {
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
		file_searchMovie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OMDBResponse); i {
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
			RawDescriptor: file_searchMovie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_searchMovie_proto_goTypes,
		DependencyIndexes: file_searchMovie_proto_depIdxs,
		MessageInfos:      file_searchMovie_proto_msgTypes,
	}.Build()
	File_searchMovie_proto = out.File
	file_searchMovie_proto_rawDesc = nil
	file_searchMovie_proto_goTypes = nil
	file_searchMovie_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchServiceClient interface {
	SearchMovie(ctx context.Context, in *SearchMovieRequest, opts ...grpc.CallOption) (*OMDBResponse, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) SearchMovie(ctx context.Context, in *SearchMovieRequest, opts ...grpc.CallOption) (*OMDBResponse, error) {
	out := new(OMDBResponse)
	err := c.cc.Invoke(ctx, "/pb.SearchService/SearchMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
type SearchServiceServer interface {
	SearchMovie(context.Context, *SearchMovieRequest) (*OMDBResponse, error)
}

// UnimplementedSearchServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (*UnimplementedSearchServiceServer) SearchMovie(context.Context, *SearchMovieRequest) (*OMDBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMovie not implemented")
}

func RegisterSearchServiceServer(s *grpc.Server, srv SearchServiceServer) {
	s.RegisterService(&_SearchService_serviceDesc, srv)
}

func _SearchService_SearchMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SearchService/SearchMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchMovie(ctx, req.(*SearchMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchMovie",
			Handler:    _SearchService_SearchMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "searchMovie.proto",
}
