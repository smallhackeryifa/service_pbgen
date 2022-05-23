// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: service.proto

package model_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ModelServiceClient is the client API for ModelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelServiceClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// 创建模型
	ModelCreate(ctx context.Context, in *ModelCreateRequest, opts ...grpc.CallOption) (*ModelCreateResponse, error)
	// 获取模型
	ModelGet(ctx context.Context, in *ModelGetRequest, opts ...grpc.CallOption) (*ModelGetResponse, error)
	// 模型列表
	ModelList(ctx context.Context, in *ModelListRequest, opts ...grpc.CallOption) (*ModelListResponse, error)
	// 模型搜索
	ModelSearch(ctx context.Context, in *ModelSearchRequest, opts ...grpc.CallOption) (*ModelSearchResponse, error)
	// 删除模型
	ModelDelete(ctx context.Context, in *ModelDeleteRequest, opts ...grpc.CallOption) (*ModelDeleteResponse, error)
	// 分享模型
	ModelShare(ctx context.Context, in *ModelShareRequest, opts ...grpc.CallOption) (*ModelShareResponse, error)
	// 训练模型
	ModelTrain(ctx context.Context, in *ModelTrainRequest, opts ...grpc.CallOption) (*ModelTrainResponse, error)
	// 模型预测
	ModelPredict(ctx context.Context, in *ModelPredictRequest, opts ...grpc.CallOption) (*ModelTrainResponse, error)
}

type modelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModelServiceClient(cc grpc.ClientConnInterface) ModelServiceClient {
	return &modelServiceClient{cc}
}

func (c *modelServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/model_api.ModelService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) ModelCreate(ctx context.Context, in *ModelCreateRequest, opts ...grpc.CallOption) (*ModelCreateResponse, error) {
	out := new(ModelCreateResponse)
	err := c.cc.Invoke(ctx, "/model_api.ModelService/ModelCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) ModelGet(ctx context.Context, in *ModelGetRequest, opts ...grpc.CallOption) (*ModelGetResponse, error) {
	out := new(ModelGetResponse)
	err := c.cc.Invoke(ctx, "/model_api.ModelService/ModelGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) ModelList(ctx context.Context, in *ModelListRequest, opts ...grpc.CallOption) (*ModelListResponse, error) {
	out := new(ModelListResponse)
	err := c.cc.Invoke(ctx, "/model_api.ModelService/ModelList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) ModelSearch(ctx context.Context, in *ModelSearchRequest, opts ...grpc.CallOption) (*ModelSearchResponse, error) {
	out := new(ModelSearchResponse)
	err := c.cc.Invoke(ctx, "/model_api.ModelService/ModelSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) ModelDelete(ctx context.Context, in *ModelDeleteRequest, opts ...grpc.CallOption) (*ModelDeleteResponse, error) {
	out := new(ModelDeleteResponse)
	err := c.cc.Invoke(ctx, "/model_api.ModelService/ModelDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) ModelShare(ctx context.Context, in *ModelShareRequest, opts ...grpc.CallOption) (*ModelShareResponse, error) {
	out := new(ModelShareResponse)
	err := c.cc.Invoke(ctx, "/model_api.ModelService/ModelShare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) ModelTrain(ctx context.Context, in *ModelTrainRequest, opts ...grpc.CallOption) (*ModelTrainResponse, error) {
	out := new(ModelTrainResponse)
	err := c.cc.Invoke(ctx, "/model_api.ModelService/ModelTrain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) ModelPredict(ctx context.Context, in *ModelPredictRequest, opts ...grpc.CallOption) (*ModelTrainResponse, error) {
	out := new(ModelTrainResponse)
	err := c.cc.Invoke(ctx, "/model_api.ModelService/ModelPredict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelServiceServer is the server API for ModelService service.
// All implementations must embed UnimplementedModelServiceServer
// for forward compatibility
type ModelServiceServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	// 创建模型
	ModelCreate(context.Context, *ModelCreateRequest) (*ModelCreateResponse, error)
	// 获取模型
	ModelGet(context.Context, *ModelGetRequest) (*ModelGetResponse, error)
	// 模型列表
	ModelList(context.Context, *ModelListRequest) (*ModelListResponse, error)
	// 模型搜索
	ModelSearch(context.Context, *ModelSearchRequest) (*ModelSearchResponse, error)
	// 删除模型
	ModelDelete(context.Context, *ModelDeleteRequest) (*ModelDeleteResponse, error)
	// 分享模型
	ModelShare(context.Context, *ModelShareRequest) (*ModelShareResponse, error)
	// 训练模型
	ModelTrain(context.Context, *ModelTrainRequest) (*ModelTrainResponse, error)
	// 模型预测
	ModelPredict(context.Context, *ModelPredictRequest) (*ModelTrainResponse, error)
	mustEmbedUnimplementedModelServiceServer()
}

// UnimplementedModelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedModelServiceServer struct {
}

func (UnimplementedModelServiceServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedModelServiceServer) ModelCreate(context.Context, *ModelCreateRequest) (*ModelCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModelCreate not implemented")
}
func (UnimplementedModelServiceServer) ModelGet(context.Context, *ModelGetRequest) (*ModelGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModelGet not implemented")
}
func (UnimplementedModelServiceServer) ModelList(context.Context, *ModelListRequest) (*ModelListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModelList not implemented")
}
func (UnimplementedModelServiceServer) ModelSearch(context.Context, *ModelSearchRequest) (*ModelSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModelSearch not implemented")
}
func (UnimplementedModelServiceServer) ModelDelete(context.Context, *ModelDeleteRequest) (*ModelDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModelDelete not implemented")
}
func (UnimplementedModelServiceServer) ModelShare(context.Context, *ModelShareRequest) (*ModelShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModelShare not implemented")
}
func (UnimplementedModelServiceServer) ModelTrain(context.Context, *ModelTrainRequest) (*ModelTrainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModelTrain not implemented")
}
func (UnimplementedModelServiceServer) ModelPredict(context.Context, *ModelPredictRequest) (*ModelTrainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModelPredict not implemented")
}
func (UnimplementedModelServiceServer) mustEmbedUnimplementedModelServiceServer() {}

// UnsafeModelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelServiceServer will
// result in compilation errors.
type UnsafeModelServiceServer interface {
	mustEmbedUnimplementedModelServiceServer()
}

func RegisterModelServiceServer(s grpc.ServiceRegistrar, srv ModelServiceServer) {
	s.RegisterService(&ModelService_ServiceDesc, srv)
}

func _ModelService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model_api.ModelService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_ModelCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).ModelCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model_api.ModelService/ModelCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).ModelCreate(ctx, req.(*ModelCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_ModelGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).ModelGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model_api.ModelService/ModelGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).ModelGet(ctx, req.(*ModelGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_ModelList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).ModelList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model_api.ModelService/ModelList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).ModelList(ctx, req.(*ModelListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_ModelSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).ModelSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model_api.ModelService/ModelSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).ModelSearch(ctx, req.(*ModelSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_ModelDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).ModelDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model_api.ModelService/ModelDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).ModelDelete(ctx, req.(*ModelDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_ModelShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).ModelShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model_api.ModelService/ModelShare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).ModelShare(ctx, req.(*ModelShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_ModelTrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelTrainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).ModelTrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model_api.ModelService/ModelTrain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).ModelTrain(ctx, req.(*ModelTrainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_ModelPredict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelPredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).ModelPredict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model_api.ModelService/ModelPredict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).ModelPredict(ctx, req.(*ModelPredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelService_ServiceDesc is the grpc.ServiceDesc for ModelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model_api.ModelService",
	HandlerType: (*ModelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _ModelService_SayHello_Handler,
		},
		{
			MethodName: "ModelCreate",
			Handler:    _ModelService_ModelCreate_Handler,
		},
		{
			MethodName: "ModelGet",
			Handler:    _ModelService_ModelGet_Handler,
		},
		{
			MethodName: "ModelList",
			Handler:    _ModelService_ModelList_Handler,
		},
		{
			MethodName: "ModelSearch",
			Handler:    _ModelService_ModelSearch_Handler,
		},
		{
			MethodName: "ModelDelete",
			Handler:    _ModelService_ModelDelete_Handler,
		},
		{
			MethodName: "ModelShare",
			Handler:    _ModelService_ModelShare_Handler,
		},
		{
			MethodName: "ModelTrain",
			Handler:    _ModelService_ModelTrain_Handler,
		},
		{
			MethodName: "ModelPredict",
			Handler:    _ModelService_ModelPredict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
