// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/movie.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MoviesList struct {
	Movies []*MoviesList_Movie `protobuf:"bytes,1,rep,name=movies" json:"movies,omitempty"`
}

func (m *MoviesList) Reset()                    { *m = MoviesList{} }
func (m *MoviesList) String() string            { return proto1.CompactTextString(m) }
func (*MoviesList) ProtoMessage()               {}
func (*MoviesList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *MoviesList) GetMovies() []*MoviesList_Movie {
	if m != nil {
		return m.Movies
	}
	return nil
}

type MoviesList_IMDBMeta struct {
	Genre      string  `protobuf:"bytes,1,opt,name=genre" json:"genre,omitempty"`
	MpaaRating string  `protobuf:"bytes,2,opt,name=mpaa_rating,json=mpaaRating" json:"mpaa_rating,omitempty"`
	Score      float32 `protobuf:"fixed32,3,opt,name=score" json:"score,omitempty"`
}

func (m *MoviesList_IMDBMeta) Reset()                    { *m = MoviesList_IMDBMeta{} }
func (m *MoviesList_IMDBMeta) String() string            { return proto1.CompactTextString(m) }
func (*MoviesList_IMDBMeta) ProtoMessage()               {}
func (*MoviesList_IMDBMeta) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *MoviesList_IMDBMeta) GetGenre() string {
	if m != nil {
		return m.Genre
	}
	return ""
}

func (m *MoviesList_IMDBMeta) GetMpaaRating() string {
	if m != nil {
		return m.MpaaRating
	}
	return ""
}

func (m *MoviesList_IMDBMeta) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

type MoviesList_RottenTomatoesMeta struct {
	TomatoScore        int32  `protobuf:"varint,1,opt,name=tomato_score,json=tomatoScore" json:"tomato_score,omitempty"`
	PopcornScore       int32  `protobuf:"varint,2,opt,name=popcorn_score,json=popcornScore" json:"popcorn_score,omitempty"`
	TheaterReleaseDate string `protobuf:"bytes,3,opt,name=theater_release_date,json=theaterReleaseDate" json:"theater_release_date,omitempty"`
	MpaaRating         string `protobuf:"bytes,4,opt,name=mpaa_rating,json=mpaaRating" json:"mpaa_rating,omitempty"`
	Synopsis           string `protobuf:"bytes,5,opt,name=synopsis" json:"synopsis,omitempty"`
	SynopsisType       string `protobuf:"bytes,6,opt,name=synopsis_type,json=synopsisType" json:"synopsis_type,omitempty"`
	Runtime            string `protobuf:"bytes,7,opt,name=runtime" json:"runtime,omitempty"`
}

func (m *MoviesList_RottenTomatoesMeta) Reset()         { *m = MoviesList_RottenTomatoesMeta{} }
func (m *MoviesList_RottenTomatoesMeta) String() string { return proto1.CompactTextString(m) }
func (*MoviesList_RottenTomatoesMeta) ProtoMessage()    {}
func (*MoviesList_RottenTomatoesMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{0, 1}
}

func (m *MoviesList_RottenTomatoesMeta) GetTomatoScore() int32 {
	if m != nil {
		return m.TomatoScore
	}
	return 0
}

func (m *MoviesList_RottenTomatoesMeta) GetPopcornScore() int32 {
	if m != nil {
		return m.PopcornScore
	}
	return 0
}

func (m *MoviesList_RottenTomatoesMeta) GetTheaterReleaseDate() string {
	if m != nil {
		return m.TheaterReleaseDate
	}
	return ""
}

func (m *MoviesList_RottenTomatoesMeta) GetMpaaRating() string {
	if m != nil {
		return m.MpaaRating
	}
	return ""
}

func (m *MoviesList_RottenTomatoesMeta) GetSynopsis() string {
	if m != nil {
		return m.Synopsis
	}
	return ""
}

func (m *MoviesList_RottenTomatoesMeta) GetSynopsisType() string {
	if m != nil {
		return m.SynopsisType
	}
	return ""
}

func (m *MoviesList_RottenTomatoesMeta) GetRuntime() string {
	if m != nil {
		return m.Runtime
	}
	return ""
}

type MoviesList_Movie struct {
	Title              string                         `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	ImdbMeta           *MoviesList_IMDBMeta           `protobuf:"bytes,2,opt,name=imdb_meta,json=imdbMeta" json:"imdb_meta,omitempty"`
	RottenTomatoesMeta *MoviesList_RottenTomatoesMeta `protobuf:"bytes,3,opt,name=rotten_tomatoes_meta,json=rottenTomatoesMeta" json:"rotten_tomatoes_meta,omitempty"`
}

func (m *MoviesList_Movie) Reset()                    { *m = MoviesList_Movie{} }
func (m *MoviesList_Movie) String() string            { return proto1.CompactTextString(m) }
func (*MoviesList_Movie) ProtoMessage()               {}
func (*MoviesList_Movie) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 2} }

func (m *MoviesList_Movie) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MoviesList_Movie) GetImdbMeta() *MoviesList_IMDBMeta {
	if m != nil {
		return m.ImdbMeta
	}
	return nil
}

func (m *MoviesList_Movie) GetRottenTomatoesMeta() *MoviesList_RottenTomatoesMeta {
	if m != nil {
		return m.RottenTomatoesMeta
	}
	return nil
}

type PostMoviesResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *PostMoviesResponse) Reset()                    { *m = PostMoviesResponse{} }
func (m *PostMoviesResponse) String() string            { return proto1.CompactTextString(m) }
func (*PostMoviesResponse) ProtoMessage()               {}
func (*PostMoviesResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *PostMoviesResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto1.RegisterType((*MoviesList)(nil), "proto.MoviesList")
	proto1.RegisterType((*MoviesList_IMDBMeta)(nil), "proto.MoviesList.IMDBMeta")
	proto1.RegisterType((*MoviesList_RottenTomatoesMeta)(nil), "proto.MoviesList.RottenTomatoesMeta")
	proto1.RegisterType((*MoviesList_Movie)(nil), "proto.MoviesList.Movie")
	proto1.RegisterType((*PostMoviesResponse)(nil), "proto.PostMoviesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MovieService service

type MovieServiceClient interface {
	BulkIndex(ctx context.Context, in *MoviesList, opts ...grpc.CallOption) (*PostMoviesResponse, error)
	Get(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*MoviesList, error)
}

type movieServiceClient struct {
	cc *grpc.ClientConn
}

func NewMovieServiceClient(cc *grpc.ClientConn) MovieServiceClient {
	return &movieServiceClient{cc}
}

func (c *movieServiceClient) BulkIndex(ctx context.Context, in *MoviesList, opts ...grpc.CallOption) (*PostMoviesResponse, error) {
	out := new(PostMoviesResponse)
	err := grpc.Invoke(ctx, "/proto.MovieService/BulkIndex", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) Get(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*MoviesList, error) {
	out := new(MoviesList)
	err := grpc.Invoke(ctx, "/proto.MovieService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MovieService service

type MovieServiceServer interface {
	BulkIndex(context.Context, *MoviesList) (*PostMoviesResponse, error)
	Get(context.Context, *google_protobuf1.Empty) (*MoviesList, error)
}

func RegisterMovieServiceServer(s *grpc.Server, srv MovieServiceServer) {
	s.RegisterService(&_MovieService_serviceDesc, srv)
}

func _MovieService_BulkIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoviesList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).BulkIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MovieService/BulkIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).BulkIndex(ctx, req.(*MoviesList))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MovieService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).Get(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _MovieService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MovieService",
	HandlerType: (*MovieServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BulkIndex",
			Handler:    _MovieService_BulkIndex_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _MovieService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/movie.proto",
}

func init() { proto1.RegisterFile("proto/movie.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0xcd, 0x8a, 0x13, 0x41,
	0x10, 0x66, 0x92, 0xcd, 0x5f, 0x25, 0x22, 0xdb, 0x04, 0x1d, 0x47, 0xc1, 0x18, 0x3d, 0x04, 0x0f,
	0x13, 0x89, 0x07, 0xc1, 0x9b, 0xcb, 0x8a, 0x2c, 0x18, 0x90, 0xde, 0x45, 0x8f, 0x43, 0x67, 0x52,
	0xc6, 0xc1, 0x4c, 0x77, 0xd3, 0x5d, 0x59, 0xcc, 0xd5, 0xa3, 0x57, 0x5f, 0xc0, 0x27, 0xf0, 0x65,
	0xf4, 0x11, 0x7c, 0x10, 0x99, 0xea, 0x99, 0xdd, 0xc3, 0x9c, 0xba, 0xbf, 0xaf, 0xbe, 0xfe, 0xea,
	0xa7, 0x29, 0x38, 0xb5, 0xce, 0x90, 0x59, 0x96, 0xe6, 0xba, 0xc0, 0x94, 0xef, 0xa2, 0xc7, 0x47,
	0xf2, 0x68, 0x67, 0xcc, 0x6e, 0x8f, 0x4b, 0x65, 0x8b, 0xa5, 0xd2, 0xda, 0x90, 0xa2, 0xc2, 0x68,
	0x1f, 0x44, 0xc9, 0xc3, 0x3a, 0xca, 0x68, 0x73, 0xf8, 0xbc, 0xc4, 0xd2, 0xd2, 0x31, 0x04, 0xe7,
	0x7f, 0x4f, 0x00, 0xd6, 0x95, 0xa3, 0x7f, 0x5f, 0x78, 0x12, 0x4b, 0xe8, 0xb3, 0xbf, 0x8f, 0xa3,
	0x59, 0x77, 0x31, 0x5e, 0xdd, 0x0f, 0xb2, 0xf4, 0x56, 0x12, 0xae, 0xb2, 0x96, 0x25, 0x9f, 0x60,
	0x78, 0xb1, 0x3e, 0x3f, 0x5b, 0x23, 0x29, 0x31, 0x85, 0xde, 0x0e, 0xb5, 0xc3, 0x38, 0x9a, 0x45,
	0x8b, 0x91, 0x0c, 0x40, 0x3c, 0x86, 0x71, 0x69, 0x95, 0xca, 0x9c, 0xa2, 0x42, 0xef, 0xe2, 0x0e,
	0xc7, 0xa0, 0xa2, 0x24, 0x33, 0xd5, 0x33, 0x9f, 0x1b, 0x87, 0x71, 0x77, 0x16, 0x2d, 0x3a, 0x32,
	0x80, 0xe4, 0x47, 0x07, 0x84, 0x34, 0x44, 0xa8, 0xaf, 0x4c, 0xa9, 0xc8, 0xa0, 0xe7, 0x1c, 0x4f,
	0x60, 0x42, 0x8c, 0xb3, 0xf0, 0xa6, 0x4a, 0xd5, 0x93, 0xe3, 0xc0, 0x5d, 0x56, 0x94, 0x78, 0x0a,
	0x77, 0xac, 0xb1, 0xb9, 0x71, 0xba, 0xd6, 0x74, 0x58, 0x33, 0xa9, 0xc9, 0x20, 0x7a, 0x01, 0x53,
	0xfa, 0x82, 0x8a, 0xd0, 0x65, 0x0e, 0xf7, 0xa8, 0x3c, 0x66, 0x5b, 0x45, 0xa1, 0x86, 0x91, 0x14,
	0x75, 0x4c, 0x86, 0xd0, 0xb9, 0xa2, 0x56, 0x1f, 0x27, 0xad, 0x3e, 0x12, 0x18, 0xfa, 0xa3, 0x36,
	0xd6, 0x17, 0x3e, 0xee, 0x71, 0xf4, 0x06, 0x57, 0x35, 0x35, 0xf7, 0x8c, 0x8e, 0x16, 0xe3, 0x3e,
	0x0b, 0x26, 0x0d, 0x79, 0x75, 0xb4, 0x28, 0x62, 0x18, 0xb8, 0x83, 0xa6, 0xa2, 0xc4, 0x78, 0xc0,
	0xe1, 0x06, 0x26, 0xbf, 0x23, 0xe8, 0xf1, 0xdc, 0xab, 0x61, 0x51, 0x41, 0xfb, 0x9b, 0x19, 0x33,
	0x10, 0xaf, 0x60, 0x54, 0x94, 0xdb, 0x4d, 0x56, 0x22, 0x29, 0x6e, 0x77, 0xbc, 0x4a, 0xda, 0x3f,
	0xd7, 0x7c, 0x94, 0x1c, 0x56, 0x62, 0x1e, 0xe7, 0x47, 0x98, 0x3a, 0x1e, 0x72, 0x46, 0xf5, 0x94,
	0x83, 0x47, 0x97, 0x3d, 0x9e, 0xb5, 0x3d, 0xda, 0x5f, 0x22, 0x85, 0x6b, 0x71, 0xf3, 0x14, 0xc4,
	0x07, 0xe3, 0x29, 0x3c, 0x94, 0xe8, 0xad, 0xd1, 0x9e, 0x1b, 0xf4, 0x87, 0x3c, 0x47, 0xef, 0xb9,
	0xfc, 0xa1, 0x6c, 0xe0, 0xea, 0x57, 0x04, 0x13, 0x16, 0x5f, 0xa2, 0xbb, 0x2e, 0x72, 0x14, 0x6b,
	0x18, 0x9d, 0x1d, 0xf6, 0x5f, 0x2f, 0xf4, 0x16, 0xbf, 0x89, 0xd3, 0x56, 0x1d, 0xc9, 0x83, 0x9a,
	0x6a, 0x67, 0x99, 0x8b, 0xef, 0x7f, 0xfe, 0xfd, 0xec, 0x4c, 0xe6, 0x83, 0xb0, 0x2a, 0xfe, 0x75,
	0xf4, 0x5c, 0xbc, 0x81, 0xee, 0x3b, 0x24, 0x71, 0x2f, 0x0d, 0xbb, 0x90, 0x36, 0xbb, 0x90, 0xbe,
	0xad, 0x76, 0x21, 0x69, 0x27, 0x98, 0xdf, 0x65, 0x97, 0x91, 0x68, 0x5c, 0x36, 0x7d, 0x96, 0xbc,
	0xfc, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x7f, 0x29, 0x5b, 0x87, 0x03, 0x00, 0x00,
}
