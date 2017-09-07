// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google_mrf.proto

/*
Package googlemrf is a generated protocol buffer package.

It is generated from these files:
	google_mrf.proto

It has these top-level messages:
	Empty
	Query
	MRF
	MRFList
	Single
	Double
*/
package googlemrf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Describes the query options available.
type Query struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parent string `protobuf:"bytes,2,opt,name=parent" json:"parent,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Query) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Query) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

// Describes an MRF code.
type MRF struct {
	Submitter     string `protobuf:"bytes,1,opt,name=Submitter,json=submitter" json:"Submitter,omitempty"`
	ID            string `protobuf:"bytes,2,opt,name=ID,json=iD" json:"ID,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=Name,json=name" json:"Name,omitempty"`
	Product       string `protobuf:"bytes,4,opt,name=Product,json=product" json:"Product,omitempty"`
	ProductCode   string `protobuf:"bytes,5,opt,name=ProductCode,json=productCode" json:"ProductCode,omitempty"`
	SubProduct    string `protobuf:"bytes,6,opt,name=SubProduct,json=subProduct" json:"SubProduct,omitempty"`
	CostCenter    string `protobuf:"bytes,7,opt,name=CostCenter,json=costCenter" json:"CostCenter,omitempty"`
	PrimaryRegion string `protobuf:"bytes,8,opt,name=PrimaryRegion,json=primaryRegion" json:"PrimaryRegion,omitempty"`
	Year          int32  `protobuf:"varint,9,opt,name=Year,json=year" json:"Year,omitempty"`
	Quarter       int32  `protobuf:"varint,10,opt,name=Quarter,json=quarter" json:"Quarter,omitempty"`
	Approved      bool   `protobuf:"varint,11,opt,name=Approved,json=approved" json:"Approved,omitempty"`
	LongTail      bool   `protobuf:"varint,12,opt,name=LongTail,json=longTail" json:"LongTail,omitempty"`
}

func (m *MRF) Reset()                    { *m = MRF{} }
func (m *MRF) String() string            { return proto.CompactTextString(m) }
func (*MRF) ProtoMessage()               {}
func (*MRF) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MRF) GetSubmitter() string {
	if m != nil {
		return m.Submitter
	}
	return ""
}

func (m *MRF) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *MRF) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MRF) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *MRF) GetProductCode() string {
	if m != nil {
		return m.ProductCode
	}
	return ""
}

func (m *MRF) GetSubProduct() string {
	if m != nil {
		return m.SubProduct
	}
	return ""
}

func (m *MRF) GetCostCenter() string {
	if m != nil {
		return m.CostCenter
	}
	return ""
}

func (m *MRF) GetPrimaryRegion() string {
	if m != nil {
		return m.PrimaryRegion
	}
	return ""
}

func (m *MRF) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *MRF) GetQuarter() int32 {
	if m != nil {
		return m.Quarter
	}
	return 0
}

func (m *MRF) GetApproved() bool {
	if m != nil {
		return m.Approved
	}
	return false
}

func (m *MRF) GetLongTail() bool {
	if m != nil {
		return m.LongTail
	}
	return false
}

// A list of MRFs.
type MRFList struct {
	Mrfs []*MRF `protobuf:"bytes,1,rep,name=Mrfs,json=mrfs" json:"Mrfs,omitempty"`
}

func (m *MRFList) Reset()                    { *m = MRFList{} }
func (m *MRFList) String() string            { return proto.CompactTextString(m) }
func (*MRFList) ProtoMessage()               {}
func (*MRFList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MRFList) GetMrfs() []*MRF {
	if m != nil {
		return m.Mrfs
	}
	return nil
}

// Describes a model with no parent, {Agency, Channel, LOB}
type Single struct {
	Name string `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
	ID   string `protobuf:"bytes,2,opt,name=ID,json=iD" json:"ID,omitempty"`
}

func (m *Single) Reset()                    { *m = Single{} }
func (m *Single) String() string            { return proto.CompactTextString(m) }
func (*Single) ProtoMessage()               {}
func (*Single) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Single) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Single) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

// Describes a model with a parent, {Media, Product, SubMedia, SubProduct}
type Double struct {
	Name   string `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
	Parent string `protobuf:"bytes,2,opt,name=Parent,json=parent" json:"Parent,omitempty"`
	ID     string `protobuf:"bytes,3,opt,name=ID,json=iD" json:"ID,omitempty"`
}

func (m *Double) Reset()                    { *m = Double{} }
func (m *Double) String() string            { return proto.CompactTextString(m) }
func (*Double) ProtoMessage()               {}
func (*Double) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Double) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Double) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Double) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "googlemrf.Empty")
	proto.RegisterType((*Query)(nil), "googlemrf.Query")
	proto.RegisterType((*MRF)(nil), "googlemrf.MRF")
	proto.RegisterType((*MRFList)(nil), "googlemrf.MRFList")
	proto.RegisterType((*Single)(nil), "googlemrf.Single")
	proto.RegisterType((*Double)(nil), "googlemrf.Double")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GoogleMRF service

type GoogleMRFClient interface {
	// Obtains a single MRF code and its related info.
	//
	// If it is not found an empty MRF is returned.
	MRFInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*MRF, error)
	// Obtains all MRF codes known.
	MRFs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GoogleMRF_MRFsClient, error)
	// Obtains all Agencies.
	Agencies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GoogleMRF_AgenciesClient, error)
	// Obtains all LOBs known.
	LOBs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GoogleMRF_LOBsClient, error)
	// Obtains all Products known.
	Products(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GoogleMRF_ProductsClient, error)
	// Obtains all SubProducts known.
	SubProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GoogleMRF_SubProductsClient, error)
	// Obtains all Channels known.
	Channels(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GoogleMRF_ChannelsClient, error)
	// Obtains all Medias known.
	Medias(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GoogleMRF_MediasClient, error)
	// Obtains all SubMedias known.
	SubMedias(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GoogleMRF_SubMediasClient, error)
	// The following methods check to see if the
	// described object exists.
	//
	// In the case of something like a product, the query
	// can be provided with both the name of the product
	// and the parent line of business. If the combination
	// of product and LOB does not exist empty will be
	// returned.
	//
	// Returns an empty instance if it is not found.
	ValidateMRF(ctx context.Context, in *Query, opts ...grpc.CallOption) (*MRF, error)
	ValidateAgency(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Single, error)
	ValidateChannel(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Single, error)
	ValidateLOB(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Single, error)
	ValidateMedia(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Single, error)
	ValidateProduct(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Single, error)
	ValidateSubMedia(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Single, error)
	ValidateSubProduct(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Single, error)
	// All 'long tail' mrfs.
	LongTailMRFs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MRFList, error)
}

type googleMRFClient struct {
	cc *grpc.ClientConn
}

func NewGoogleMRFClient(cc *grpc.ClientConn) GoogleMRFClient {
	return &googleMRFClient{cc}
}

func (c *googleMRFClient) MRFInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*MRF, error) {
	out := new(MRF)
	err := grpc.Invoke(ctx, "/googlemrf.GoogleMRF/MRFInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleMRFClient) MRFs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GoogleMRF_MRFsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GoogleMRF_serviceDesc.Streams[0], c.cc, "/googlemrf.GoogleMRF/MRFs", opts...)
	if err != nil {
		return nil, err
	}
	x := &googleMRFMRFsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoogleMRF_MRFsClient interface {
	Recv() (*MRF, error)
	grpc.ClientStream
}

type googleMRFMRFsClient struct {
	grpc.ClientStream
}

func (x *googleMRFMRFsClient) Recv() (*MRF, error) {
	m := new(MRF)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *googleMRFClient) Agencies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GoogleMRF_AgenciesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GoogleMRF_serviceDesc.Streams[1], c.cc, "/googlemrf.GoogleMRF/Agencies", opts...)
	if err != nil {
		return nil, err
	}
	x := &googleMRFAgenciesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoogleMRF_AgenciesClient interface {
	Recv() (*Single, error)
	grpc.ClientStream
}

type googleMRFAgenciesClient struct {
	grpc.ClientStream
}

func (x *googleMRFAgenciesClient) Recv() (*Single, error) {
	m := new(Single)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *googleMRFClient) LOBs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GoogleMRF_LOBsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GoogleMRF_serviceDesc.Streams[2], c.cc, "/googlemrf.GoogleMRF/LOBs", opts...)
	if err != nil {
		return nil, err
	}
	x := &googleMRFLOBsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoogleMRF_LOBsClient interface {
	Recv() (*Single, error)
	grpc.ClientStream
}

type googleMRFLOBsClient struct {
	grpc.ClientStream
}

func (x *googleMRFLOBsClient) Recv() (*Single, error) {
	m := new(Single)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *googleMRFClient) Products(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GoogleMRF_ProductsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GoogleMRF_serviceDesc.Streams[3], c.cc, "/googlemrf.GoogleMRF/Products", opts...)
	if err != nil {
		return nil, err
	}
	x := &googleMRFProductsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoogleMRF_ProductsClient interface {
	Recv() (*Double, error)
	grpc.ClientStream
}

type googleMRFProductsClient struct {
	grpc.ClientStream
}

func (x *googleMRFProductsClient) Recv() (*Double, error) {
	m := new(Double)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *googleMRFClient) SubProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GoogleMRF_SubProductsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GoogleMRF_serviceDesc.Streams[4], c.cc, "/googlemrf.GoogleMRF/SubProducts", opts...)
	if err != nil {
		return nil, err
	}
	x := &googleMRFSubProductsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoogleMRF_SubProductsClient interface {
	Recv() (*Double, error)
	grpc.ClientStream
}

type googleMRFSubProductsClient struct {
	grpc.ClientStream
}

func (x *googleMRFSubProductsClient) Recv() (*Double, error) {
	m := new(Double)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *googleMRFClient) Channels(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GoogleMRF_ChannelsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GoogleMRF_serviceDesc.Streams[5], c.cc, "/googlemrf.GoogleMRF/Channels", opts...)
	if err != nil {
		return nil, err
	}
	x := &googleMRFChannelsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoogleMRF_ChannelsClient interface {
	Recv() (*Single, error)
	grpc.ClientStream
}

type googleMRFChannelsClient struct {
	grpc.ClientStream
}

func (x *googleMRFChannelsClient) Recv() (*Single, error) {
	m := new(Single)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *googleMRFClient) Medias(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GoogleMRF_MediasClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GoogleMRF_serviceDesc.Streams[6], c.cc, "/googlemrf.GoogleMRF/Medias", opts...)
	if err != nil {
		return nil, err
	}
	x := &googleMRFMediasClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoogleMRF_MediasClient interface {
	Recv() (*Double, error)
	grpc.ClientStream
}

type googleMRFMediasClient struct {
	grpc.ClientStream
}

func (x *googleMRFMediasClient) Recv() (*Double, error) {
	m := new(Double)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *googleMRFClient) SubMedias(ctx context.Context, in *Empty, opts ...grpc.CallOption) (GoogleMRF_SubMediasClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GoogleMRF_serviceDesc.Streams[7], c.cc, "/googlemrf.GoogleMRF/SubMedias", opts...)
	if err != nil {
		return nil, err
	}
	x := &googleMRFSubMediasClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoogleMRF_SubMediasClient interface {
	Recv() (*Double, error)
	grpc.ClientStream
}

type googleMRFSubMediasClient struct {
	grpc.ClientStream
}

func (x *googleMRFSubMediasClient) Recv() (*Double, error) {
	m := new(Double)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *googleMRFClient) ValidateMRF(ctx context.Context, in *Query, opts ...grpc.CallOption) (*MRF, error) {
	out := new(MRF)
	err := grpc.Invoke(ctx, "/googlemrf.GoogleMRF/ValidateMRF", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleMRFClient) ValidateAgency(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Single, error) {
	out := new(Single)
	err := grpc.Invoke(ctx, "/googlemrf.GoogleMRF/ValidateAgency", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleMRFClient) ValidateChannel(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Single, error) {
	out := new(Single)
	err := grpc.Invoke(ctx, "/googlemrf.GoogleMRF/ValidateChannel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleMRFClient) ValidateLOB(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Single, error) {
	out := new(Single)
	err := grpc.Invoke(ctx, "/googlemrf.GoogleMRF/ValidateLOB", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleMRFClient) ValidateMedia(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Single, error) {
	out := new(Single)
	err := grpc.Invoke(ctx, "/googlemrf.GoogleMRF/ValidateMedia", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleMRFClient) ValidateProduct(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Single, error) {
	out := new(Single)
	err := grpc.Invoke(ctx, "/googlemrf.GoogleMRF/ValidateProduct", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleMRFClient) ValidateSubMedia(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Single, error) {
	out := new(Single)
	err := grpc.Invoke(ctx, "/googlemrf.GoogleMRF/ValidateSubMedia", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleMRFClient) ValidateSubProduct(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Single, error) {
	out := new(Single)
	err := grpc.Invoke(ctx, "/googlemrf.GoogleMRF/ValidateSubProduct", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *googleMRFClient) LongTailMRFs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MRFList, error) {
	out := new(MRFList)
	err := grpc.Invoke(ctx, "/googlemrf.GoogleMRF/LongTailMRFs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoogleMRF service

type GoogleMRFServer interface {
	// Obtains a single MRF code and its related info.
	//
	// If it is not found an empty MRF is returned.
	MRFInfo(context.Context, *Query) (*MRF, error)
	// Obtains all MRF codes known.
	MRFs(*Empty, GoogleMRF_MRFsServer) error
	// Obtains all Agencies.
	Agencies(*Empty, GoogleMRF_AgenciesServer) error
	// Obtains all LOBs known.
	LOBs(*Empty, GoogleMRF_LOBsServer) error
	// Obtains all Products known.
	Products(*Empty, GoogleMRF_ProductsServer) error
	// Obtains all SubProducts known.
	SubProducts(*Empty, GoogleMRF_SubProductsServer) error
	// Obtains all Channels known.
	Channels(*Empty, GoogleMRF_ChannelsServer) error
	// Obtains all Medias known.
	Medias(*Empty, GoogleMRF_MediasServer) error
	// Obtains all SubMedias known.
	SubMedias(*Empty, GoogleMRF_SubMediasServer) error
	// The following methods check to see if the
	// described object exists.
	//
	// In the case of something like a product, the query
	// can be provided with both the name of the product
	// and the parent line of business. If the combination
	// of product and LOB does not exist empty will be
	// returned.
	//
	// Returns an empty instance if it is not found.
	ValidateMRF(context.Context, *Query) (*MRF, error)
	ValidateAgency(context.Context, *Query) (*Single, error)
	ValidateChannel(context.Context, *Query) (*Single, error)
	ValidateLOB(context.Context, *Query) (*Single, error)
	ValidateMedia(context.Context, *Query) (*Single, error)
	ValidateProduct(context.Context, *Query) (*Single, error)
	ValidateSubMedia(context.Context, *Query) (*Single, error)
	ValidateSubProduct(context.Context, *Query) (*Single, error)
	// All 'long tail' mrfs.
	LongTailMRFs(context.Context, *Empty) (*MRFList, error)
}

func RegisterGoogleMRFServer(s *grpc.Server, srv GoogleMRFServer) {
	s.RegisterService(&_GoogleMRF_serviceDesc, srv)
}

func _GoogleMRF_MRFInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleMRFServer).MRFInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/googlemrf.GoogleMRF/MRFInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleMRFServer).MRFInfo(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleMRF_MRFs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoogleMRFServer).MRFs(m, &googleMRFMRFsServer{stream})
}

type GoogleMRF_MRFsServer interface {
	Send(*MRF) error
	grpc.ServerStream
}

type googleMRFMRFsServer struct {
	grpc.ServerStream
}

func (x *googleMRFMRFsServer) Send(m *MRF) error {
	return x.ServerStream.SendMsg(m)
}

func _GoogleMRF_Agencies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoogleMRFServer).Agencies(m, &googleMRFAgenciesServer{stream})
}

type GoogleMRF_AgenciesServer interface {
	Send(*Single) error
	grpc.ServerStream
}

type googleMRFAgenciesServer struct {
	grpc.ServerStream
}

func (x *googleMRFAgenciesServer) Send(m *Single) error {
	return x.ServerStream.SendMsg(m)
}

func _GoogleMRF_LOBs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoogleMRFServer).LOBs(m, &googleMRFLOBsServer{stream})
}

type GoogleMRF_LOBsServer interface {
	Send(*Single) error
	grpc.ServerStream
}

type googleMRFLOBsServer struct {
	grpc.ServerStream
}

func (x *googleMRFLOBsServer) Send(m *Single) error {
	return x.ServerStream.SendMsg(m)
}

func _GoogleMRF_Products_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoogleMRFServer).Products(m, &googleMRFProductsServer{stream})
}

type GoogleMRF_ProductsServer interface {
	Send(*Double) error
	grpc.ServerStream
}

type googleMRFProductsServer struct {
	grpc.ServerStream
}

func (x *googleMRFProductsServer) Send(m *Double) error {
	return x.ServerStream.SendMsg(m)
}

func _GoogleMRF_SubProducts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoogleMRFServer).SubProducts(m, &googleMRFSubProductsServer{stream})
}

type GoogleMRF_SubProductsServer interface {
	Send(*Double) error
	grpc.ServerStream
}

type googleMRFSubProductsServer struct {
	grpc.ServerStream
}

func (x *googleMRFSubProductsServer) Send(m *Double) error {
	return x.ServerStream.SendMsg(m)
}

func _GoogleMRF_Channels_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoogleMRFServer).Channels(m, &googleMRFChannelsServer{stream})
}

type GoogleMRF_ChannelsServer interface {
	Send(*Single) error
	grpc.ServerStream
}

type googleMRFChannelsServer struct {
	grpc.ServerStream
}

func (x *googleMRFChannelsServer) Send(m *Single) error {
	return x.ServerStream.SendMsg(m)
}

func _GoogleMRF_Medias_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoogleMRFServer).Medias(m, &googleMRFMediasServer{stream})
}

type GoogleMRF_MediasServer interface {
	Send(*Double) error
	grpc.ServerStream
}

type googleMRFMediasServer struct {
	grpc.ServerStream
}

func (x *googleMRFMediasServer) Send(m *Double) error {
	return x.ServerStream.SendMsg(m)
}

func _GoogleMRF_SubMedias_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoogleMRFServer).SubMedias(m, &googleMRFSubMediasServer{stream})
}

type GoogleMRF_SubMediasServer interface {
	Send(*Double) error
	grpc.ServerStream
}

type googleMRFSubMediasServer struct {
	grpc.ServerStream
}

func (x *googleMRFSubMediasServer) Send(m *Double) error {
	return x.ServerStream.SendMsg(m)
}

func _GoogleMRF_ValidateMRF_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleMRFServer).ValidateMRF(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/googlemrf.GoogleMRF/ValidateMRF",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleMRFServer).ValidateMRF(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleMRF_ValidateAgency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleMRFServer).ValidateAgency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/googlemrf.GoogleMRF/ValidateAgency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleMRFServer).ValidateAgency(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleMRF_ValidateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleMRFServer).ValidateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/googlemrf.GoogleMRF/ValidateChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleMRFServer).ValidateChannel(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleMRF_ValidateLOB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleMRFServer).ValidateLOB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/googlemrf.GoogleMRF/ValidateLOB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleMRFServer).ValidateLOB(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleMRF_ValidateMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleMRFServer).ValidateMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/googlemrf.GoogleMRF/ValidateMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleMRFServer).ValidateMedia(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleMRF_ValidateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleMRFServer).ValidateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/googlemrf.GoogleMRF/ValidateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleMRFServer).ValidateProduct(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleMRF_ValidateSubMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleMRFServer).ValidateSubMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/googlemrf.GoogleMRF/ValidateSubMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleMRFServer).ValidateSubMedia(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleMRF_ValidateSubProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleMRFServer).ValidateSubProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/googlemrf.GoogleMRF/ValidateSubProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleMRFServer).ValidateSubProduct(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoogleMRF_LongTailMRFs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoogleMRFServer).LongTailMRFs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/googlemrf.GoogleMRF/LongTailMRFs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoogleMRFServer).LongTailMRFs(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoogleMRF_serviceDesc = grpc.ServiceDesc{
	ServiceName: "googlemrf.GoogleMRF",
	HandlerType: (*GoogleMRFServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MRFInfo",
			Handler:    _GoogleMRF_MRFInfo_Handler,
		},
		{
			MethodName: "ValidateMRF",
			Handler:    _GoogleMRF_ValidateMRF_Handler,
		},
		{
			MethodName: "ValidateAgency",
			Handler:    _GoogleMRF_ValidateAgency_Handler,
		},
		{
			MethodName: "ValidateChannel",
			Handler:    _GoogleMRF_ValidateChannel_Handler,
		},
		{
			MethodName: "ValidateLOB",
			Handler:    _GoogleMRF_ValidateLOB_Handler,
		},
		{
			MethodName: "ValidateMedia",
			Handler:    _GoogleMRF_ValidateMedia_Handler,
		},
		{
			MethodName: "ValidateProduct",
			Handler:    _GoogleMRF_ValidateProduct_Handler,
		},
		{
			MethodName: "ValidateSubMedia",
			Handler:    _GoogleMRF_ValidateSubMedia_Handler,
		},
		{
			MethodName: "ValidateSubProduct",
			Handler:    _GoogleMRF_ValidateSubProduct_Handler,
		},
		{
			MethodName: "LongTailMRFs",
			Handler:    _GoogleMRF_LongTailMRFs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MRFs",
			Handler:       _GoogleMRF_MRFs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Agencies",
			Handler:       _GoogleMRF_Agencies_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LOBs",
			Handler:       _GoogleMRF_LOBs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Products",
			Handler:       _GoogleMRF_Products_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubProducts",
			Handler:       _GoogleMRF_SubProducts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Channels",
			Handler:       _GoogleMRF_Channels_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Medias",
			Handler:       _GoogleMRF_Medias_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubMedias",
			Handler:       _GoogleMRF_SubMedias_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "google_mrf.proto",
}

func init() { proto.RegisterFile("google_mrf.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 690 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x95, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0xc7, 0x6f, 0x12, 0xc7, 0xb1, 0x4f, 0x80, 0xcb, 0x9d, 0x2b, 0xc0, 0x4d, 0x69, 0x1b, 0x59,
	0x5d, 0x64, 0x41, 0x49, 0x21, 0x55, 0xf7, 0x90, 0x34, 0x15, 0x52, 0x52, 0xc0, 0xa9, 0xa8, 0x58,
	0x54, 0xd5, 0x24, 0x99, 0xa4, 0x23, 0xd9, 0x33, 0xee, 0xd8, 0x41, 0x8a, 0xaa, 0x6e, 0xfa, 0x0a,
	0x7d, 0xb0, 0x2e, 0xfa, 0x0a, 0x55, 0x9f, 0xa3, 0x9a, 0xb1, 0xc7, 0x18, 0x08, 0x82, 0xb0, 0xf3,
	0x9c, 0x8f, 0xdf, 0xfc, 0xe7, 0x9c, 0x93, 0x13, 0x58, 0x9f, 0x72, 0x3e, 0xf5, 0xc9, 0xa7, 0x40,
	0x4c, 0x76, 0x43, 0xc1, 0x63, 0x8e, 0xec, 0xc4, 0x12, 0x88, 0x49, 0x6d, 0x3b, 0xf9, 0x6c, 0xe2,
	0x90, 0x36, 0x31, 0x63, 0x3c, 0xc6, 0x31, 0xe5, 0x2c, 0x4a, 0x02, 0xdd, 0x0a, 0x94, 0xdf, 0x04,
	0x61, 0x3c, 0x77, 0x5b, 0x50, 0x3e, 0x9d, 0x11, 0x31, 0x47, 0x08, 0x0c, 0x86, 0x03, 0xe2, 0x14,
	0xea, 0x85, 0x86, 0xed, 0xa9, 0x6f, 0xb4, 0x09, 0x66, 0x88, 0x05, 0x61, 0xb1, 0x53, 0x54, 0xd6,
	0xf4, 0xe4, 0xfe, 0x2c, 0x42, 0xa9, 0xef, 0x75, 0xd1, 0x36, 0xd8, 0x83, 0xd9, 0x30, 0xa0, 0x71,
	0x4c, 0x44, 0x9a, 0x68, 0x47, 0xda, 0x80, 0xd6, 0xa0, 0x78, 0xd4, 0x49, 0x33, 0x8b, 0xb4, 0x23,
	0x6f, 0x78, 0x27, 0x6f, 0x28, 0xe5, 0x6e, 0x70, 0xa0, 0x72, 0x22, 0xf8, 0x78, 0x36, 0x8a, 0x1d,
	0x43, 0x99, 0x2b, 0x61, 0x72, 0x44, 0x75, 0xa8, 0xa6, 0x9e, 0x36, 0x1f, 0x13, 0xa7, 0xac, 0xbc,
	0xd5, 0xf0, 0xd2, 0x84, 0x9e, 0x02, 0x0c, 0x66, 0x43, 0x9d, 0x6e, 0xaa, 0x00, 0x88, 0x32, 0x8b,
	0xf4, 0xb7, 0x79, 0x14, 0xb7, 0x09, 0x93, 0xf2, 0x2a, 0x89, 0x7f, 0x94, 0x59, 0xd0, 0x73, 0x58,
	0x3d, 0x11, 0x34, 0xc0, 0x62, 0xee, 0x91, 0x29, 0xe5, 0xcc, 0xb1, 0x54, 0xc8, 0x6a, 0x98, 0x37,
	0x4a, 0xd5, 0xe7, 0x04, 0x0b, 0xc7, 0xae, 0x17, 0x1a, 0x65, 0xcf, 0x98, 0x13, 0x2c, 0xa4, 0xea,
	0xd3, 0x19, 0x16, 0x12, 0x0b, 0xca, 0x5c, 0xf9, 0x92, 0x1c, 0x51, 0x0d, 0xac, 0x83, 0x30, 0x14,
	0xfc, 0x82, 0x8c, 0x9d, 0x6a, 0xbd, 0xd0, 0xb0, 0x3c, 0x0b, 0xa7, 0x67, 0xe9, 0xeb, 0x71, 0x36,
	0x7d, 0x8f, 0xa9, 0xef, 0xac, 0x24, 0x3e, 0x3f, 0x3d, 0xbb, 0x2f, 0xa0, 0xd2, 0xf7, 0xba, 0x3d,
	0x1a, 0xc5, 0xc8, 0x05, 0xa3, 0x2f, 0x26, 0x91, 0x53, 0xa8, 0x97, 0x1a, 0xd5, 0xfd, 0xb5, 0xdd,
	0xac, 0xa5, 0xbb, 0x7d, 0xaf, 0xeb, 0x19, 0x81, 0x98, 0x44, 0xee, 0x0e, 0x98, 0x03, 0xca, 0xa6,
	0x3e, 0xc9, 0x8a, 0x9a, 0x6f, 0xdb, 0xb5, 0xc2, 0xbb, 0x1d, 0x30, 0x3b, 0x7c, 0x36, 0xbc, 0x25,
	0x7a, 0x13, 0xcc, 0x93, 0x05, 0x4d, 0x4e, 0x29, 0x25, 0x4d, 0xd9, 0xff, 0x63, 0x83, 0xfd, 0x56,
	0x69, 0x91, 0xad, 0x4f, 0x04, 0x1f, 0xb1, 0x09, 0x47, 0xeb, 0x39, 0x89, 0x6a, 0x96, 0x6a, 0xd7,
	0x44, 0xbb, 0xff, 0xa0, 0x1d, 0x30, 0xfa, 0x5e, 0x37, 0xba, 0x12, 0xab, 0x06, 0xf0, 0x66, 0xec,
	0xcb, 0x02, 0x6a, 0x81, 0x75, 0x30, 0x25, 0x6c, 0x44, 0xc9, 0xa2, 0x8c, 0xff, 0x72, 0x96, 0xa4,
	0x0a, 0x2a, 0xa9, 0x09, 0x46, 0xef, 0xf8, 0x70, 0x89, 0x84, 0x16, 0x58, 0xe9, 0xa8, 0xdc, 0x95,
	0x94, 0x54, 0x4f, 0x25, 0xbd, 0x86, 0xea, 0xe5, 0xd0, 0x2d, 0x91, 0xd7, 0x02, 0xab, 0xfd, 0x19,
	0x33, 0x46, 0xfc, 0x25, 0x14, 0xee, 0x81, 0xd9, 0x27, 0x63, 0x8a, 0x97, 0xb8, 0xe7, 0x95, 0xfa,
	0x49, 0x2e, 0x9b, 0xd5, 0x83, 0xea, 0x19, 0xf6, 0xe9, 0x18, 0xc7, 0xaa, 0xb9, 0x77, 0x77, 0xf4,
	0xf1, 0xf7, 0x5f, 0xbf, 0x7f, 0x14, 0x37, 0xd0, 0xff, 0x6a, 0xbf, 0x5c, 0xec, 0x35, 0xe5, 0x58,
	0x36, 0xbf, 0xca, 0x81, 0xfa, 0x86, 0xce, 0x60, 0x4d, 0xd3, 0x54, 0x1b, 0xe7, 0x0b, 0x80, 0x0b,
	0x5e, 0xfc, 0x4c, 0x31, 0x1f, 0xa1, 0x2d, 0xcd, 0xc4, 0xe9, 0x0c, 0x68, 0xee, 0x07, 0xf8, 0x57,
	0x73, 0xd3, 0x5a, 0x3e, 0x10, 0x3c, 0x4a, 0x3b, 0xa1, 0xc1, 0xc7, 0x97, 0xcf, 0xef, 0x1d, 0x1f,
	0xde, 0x0f, 0x7a, 0xa3, 0x02, 0x3e, 0x1f, 0x66, 0xc0, 0x01, 0xac, 0x66, 0xf5, 0x94, 0xad, 0xb8,
	0x1f, 0xf2, 0x89, 0x42, 0x6e, 0xa1, 0x8d, 0xac, 0xa8, 0xaa, 0x8d, 0x0b, 0x9e, 0xaf, 0x57, 0xdc,
	0xc3, 0x9e, 0x9f, 0x2e, 0xd1, 0x0c, 0x7c, 0x0e, 0xeb, 0x1a, 0xac, 0x67, 0xe7, 0x7e, 0xe4, 0xba,
	0x22, 0xd7, 0x90, 0xa3, 0xc9, 0x72, 0xf9, 0x5f, 0xd1, 0xfc, 0x11, 0x50, 0x0e, 0xbd, 0x94, 0x6c,
	0x57, 0xc1, 0xb7, 0x51, 0x2d, 0x07, 0xbf, 0xae, 0x7c, 0x00, 0x2b, 0x7a, 0xa5, 0xde, 0xb2, 0x5e,
	0xd0, 0xd5, 0xc1, 0x95, 0x1b, 0x76, 0x41, 0x9d, 0xe5, 0xf0, 0xca, 0x5d, 0x1c, 0x63, 0xea, 0x0f,
	0x4d, 0xf5, 0x17, 0xd9, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xc3, 0x60, 0xc3, 0x5f, 0x07,
	0x00, 0x00,
}
