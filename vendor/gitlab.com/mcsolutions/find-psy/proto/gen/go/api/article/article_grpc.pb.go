// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/article/article.proto

package article

import (
	context "context"
	common "gitlab.com/mcsolutions/find-psy/proto/gen/go/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ArticleClient is the client API for Article service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleClient interface {
	Create(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*ArticleReply, error)
	Get(ctx context.Context, in *common.IdRequest, opts ...grpc.CallOption) (*ArticleReply, error)
	Update(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*ArticleReply, error)
	Delete(ctx context.Context, in *common.IdRequest, opts ...grpc.CallOption) (*common.Empty, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	ListMy(ctx context.Context, in *ListMyRequest, opts ...grpc.CallOption) (*ListReply, error)
	Publish(ctx context.Context, in *common.IdRequest, opts ...grpc.CallOption) (*common.Empty, error)
	Revoke(ctx context.Context, in *common.IdRequest, opts ...grpc.CallOption) (*common.Empty, error)
	// article-admin
	ListToProcess(ctx context.Context, in *ListToProcessRequest, opts ...grpc.CallOption) (*ListReply, error)
	SetProcessed(ctx context.Context, in *SetProcessedRequest, opts ...grpc.CallOption) (*common.Empty, error)
}

type articleClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleClient(cc grpc.ClientConnInterface) ArticleClient {
	return &articleClient{cc}
}

func (c *articleClient) Create(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*ArticleReply, error) {
	out := new(ArticleReply)
	err := c.cc.Invoke(ctx, "/api.article.Article/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) Get(ctx context.Context, in *common.IdRequest, opts ...grpc.CallOption) (*ArticleReply, error) {
	out := new(ArticleReply)
	err := c.cc.Invoke(ctx, "/api.article.Article/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) Update(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*ArticleReply, error) {
	out := new(ArticleReply)
	err := c.cc.Invoke(ctx, "/api.article.Article/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) Delete(ctx context.Context, in *common.IdRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/api.article.Article/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, "/api.article.Article/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) ListMy(ctx context.Context, in *ListMyRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, "/api.article.Article/ListMy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) Publish(ctx context.Context, in *common.IdRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/api.article.Article/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) Revoke(ctx context.Context, in *common.IdRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/api.article.Article/Revoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) ListToProcess(ctx context.Context, in *ListToProcessRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, "/api.article.Article/ListToProcess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) SetProcessed(ctx context.Context, in *SetProcessedRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/api.article.Article/SetProcessed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServer is the server API for Article service.
// All implementations must embed UnimplementedArticleServer
// for forward compatibility
type ArticleServer interface {
	Create(context.Context, *CreateArticleRequest) (*ArticleReply, error)
	Get(context.Context, *common.IdRequest) (*ArticleReply, error)
	Update(context.Context, *UpdateArticleRequest) (*ArticleReply, error)
	Delete(context.Context, *common.IdRequest) (*common.Empty, error)
	List(context.Context, *ListRequest) (*ListReply, error)
	ListMy(context.Context, *ListMyRequest) (*ListReply, error)
	Publish(context.Context, *common.IdRequest) (*common.Empty, error)
	Revoke(context.Context, *common.IdRequest) (*common.Empty, error)
	// article-admin
	ListToProcess(context.Context, *ListToProcessRequest) (*ListReply, error)
	SetProcessed(context.Context, *SetProcessedRequest) (*common.Empty, error)
	mustEmbedUnimplementedArticleServer()
}

// UnimplementedArticleServer must be embedded to have forward compatible implementations.
type UnimplementedArticleServer struct {
}

func (UnimplementedArticleServer) Create(context.Context, *CreateArticleRequest) (*ArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedArticleServer) Get(context.Context, *common.IdRequest) (*ArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedArticleServer) Update(context.Context, *UpdateArticleRequest) (*ArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedArticleServer) Delete(context.Context, *common.IdRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedArticleServer) List(context.Context, *ListRequest) (*ListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedArticleServer) ListMy(context.Context, *ListMyRequest) (*ListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMy not implemented")
}
func (UnimplementedArticleServer) Publish(context.Context, *common.IdRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedArticleServer) Revoke(context.Context, *common.IdRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Revoke not implemented")
}
func (UnimplementedArticleServer) ListToProcess(context.Context, *ListToProcessRequest) (*ListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListToProcess not implemented")
}
func (UnimplementedArticleServer) SetProcessed(context.Context, *SetProcessedRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProcessed not implemented")
}
func (UnimplementedArticleServer) mustEmbedUnimplementedArticleServer() {}

// UnsafeArticleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServer will
// result in compilation errors.
type UnsafeArticleServer interface {
	mustEmbedUnimplementedArticleServer()
}

func RegisterArticleServer(s grpc.ServiceRegistrar, srv ArticleServer) {
	s.RegisterService(&Article_ServiceDesc, srv)
}

func _Article_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.article.Article/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).Create(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.article.Article/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).Get(ctx, req.(*common.IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.article.Article/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).Update(ctx, req.(*UpdateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.article.Article/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).Delete(ctx, req.(*common.IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.article.Article/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_ListMy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).ListMy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.article.Article/ListMy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).ListMy(ctx, req.(*ListMyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.article.Article/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).Publish(ctx, req.(*common.IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.article.Article/Revoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).Revoke(ctx, req.(*common.IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_ListToProcess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListToProcessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).ListToProcess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.article.Article/ListToProcess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).ListToProcess(ctx, req.(*ListToProcessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_SetProcessed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetProcessedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).SetProcessed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.article.Article/SetProcessed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).SetProcessed(ctx, req.(*SetProcessedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Article_ServiceDesc is the grpc.ServiceDesc for Article service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Article_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.article.Article",
	HandlerType: (*ArticleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Article_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Article_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Article_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Article_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Article_List_Handler,
		},
		{
			MethodName: "ListMy",
			Handler:    _Article_ListMy_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _Article_Publish_Handler,
		},
		{
			MethodName: "Revoke",
			Handler:    _Article_Revoke_Handler,
		},
		{
			MethodName: "ListToProcess",
			Handler:    _Article_ListToProcess_Handler,
		},
		{
			MethodName: "SetProcessed",
			Handler:    _Article_SetProcessed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/article/article.proto",
}