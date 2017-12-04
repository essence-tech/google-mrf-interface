// Code generated by protoc-gen-go.
// source: google_mrf.proto
// DO NOT EDIT!

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
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

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
	Submitter           string                      `protobuf:"bytes,1,opt,name=Submitter,json=submitter" json:"Submitter,omitempty"`
	ID                  string                      `protobuf:"bytes,2,opt,name=ID,json=iD" json:"ID,omitempty"`
	Name                string                      `protobuf:"bytes,3,opt,name=Name,json=name" json:"Name,omitempty"`
	Product             string                      `protobuf:"bytes,4,opt,name=Product,json=product" json:"Product,omitempty"`
	ProductCode         string                      `protobuf:"bytes,5,opt,name=ProductCode,json=productCode" json:"ProductCode,omitempty"`
	SubProduct          string                      `protobuf:"bytes,6,opt,name=SubProduct,json=subProduct" json:"SubProduct,omitempty"`
	CostCenter          string                      `protobuf:"bytes,7,opt,name=CostCenter,json=costCenter" json:"CostCenter,omitempty"`
	PrimaryRegion       string                      `protobuf:"bytes,8,opt,name=PrimaryRegion,json=primaryRegion" json:"PrimaryRegion,omitempty"`
	Year                int32                       `protobuf:"varint,9,opt,name=Year,json=year" json:"Year,omitempty"`
	Quarter             int32                       `protobuf:"varint,10,opt,name=Quarter,json=quarter" json:"Quarter,omitempty"`
	Approved            bool                        `protobuf:"varint,11,opt,name=Approved,json=approved" json:"Approved,omitempty"`
	LongTail            bool                        `protobuf:"varint,12,opt,name=LongTail,json=longTail" json:"LongTail,omitempty"`
	MediaCode           string                      `protobuf:"bytes,13,opt,name=MediaCode,json=mediaCode" json:"MediaCode,omitempty"`
	SubmissionDate      *google_protobuf1.Timestamp `protobuf:"bytes,14,opt,name=SubmissionDate,json=submissionDate" json:"SubmissionDate,omitempty"`
	CloseDate           *google_protobuf1.Timestamp `protobuf:"bytes,15,opt,name=CloseDate,json=closeDate" json:"CloseDate,omitempty"`
	InitialApprovalDate *google_protobuf1.Timestamp `protobuf:"bytes,16,opt,name=InitialApprovalDate,json=initialApprovalDate" json:"InitialApprovalDate,omitempty"`
	HelpText            string                      `protobuf:"bytes,17,opt,name=HelpText,json=helpText" json:"HelpText,omitempty"`
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

func (m *MRF) GetMediaCode() string {
	if m != nil {
		return m.MediaCode
	}
	return ""
}

func (m *MRF) GetSubmissionDate() *google_protobuf1.Timestamp {
	if m != nil {
		return m.SubmissionDate
	}
	return nil
}

func (m *MRF) GetCloseDate() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CloseDate
	}
	return nil
}

func (m *MRF) GetInitialApprovalDate() *google_protobuf1.Timestamp {
	if m != nil {
		return m.InitialApprovalDate
	}
	return nil
}

func (m *MRF) GetHelpText() string {
	if m != nil {
		return m.HelpText
	}
	return ""
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
	// 805 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0x2e, 0x60, 0xc0, 0x3e, 0x2c, 0x2c, 0x19, 0xb4, 0xbb, 0x5e, 0x9a, 0x76, 0x91, 0xd5, 0x0b,
	0x2e, 0xb6, 0xd0, 0x84, 0xaa, 0xea, 0x6d, 0x02, 0xa5, 0x8d, 0x04, 0x25, 0x31, 0x51, 0xaa, 0x5c,
	0x54, 0xd5, 0x00, 0x03, 0x19, 0xc9, 0x9e, 0x71, 0xed, 0x21, 0x2a, 0xaa, 0x7a, 0xd3, 0x57, 0xe8,
	0xa3, 0x55, 0x7d, 0x81, 0xaa, 0x0f, 0x52, 0xcd, 0xd8, 0x63, 0x20, 0x21, 0x22, 0xe4, 0xce, 0xe7,
	0xe7, 0xfb, 0xe6, 0xfc, 0x1b, 0xaa, 0x0b, 0xce, 0x17, 0x1e, 0xf9, 0xc5, 0x0f, 0xe7, 0xad, 0x20,
	0xe4, 0x82, 0x23, 0x2b, 0xd6, 0xf8, 0xe1, 0xbc, 0x7e, 0x1c, 0x7f, 0xb6, 0x71, 0x40, 0xdb, 0x98,
	0x31, 0x2e, 0xb0, 0xa0, 0x9c, 0x45, 0xb1, 0x63, 0xfd, 0x43, 0x62, 0x55, 0xd2, 0x64, 0x39, 0x6f,
	0x0b, 0xea, 0x93, 0x48, 0x60, 0x3f, 0x88, 0x1d, 0x9c, 0x22, 0xe4, 0xbf, 0xf3, 0x03, 0xb1, 0x72,
	0x3a, 0x90, 0xbf, 0x5a, 0x92, 0x70, 0x85, 0x10, 0x18, 0x0c, 0xfb, 0xc4, 0xce, 0x34, 0x32, 0x4d,
	0xcb, 0x55, 0xdf, 0xe8, 0x2d, 0x14, 0x02, 0x1c, 0x12, 0x26, 0xec, 0xac, 0xd2, 0x26, 0x92, 0xf3,
	0x8f, 0x01, 0xb9, 0xa1, 0xdb, 0x47, 0xc7, 0x60, 0x8d, 0x97, 0x13, 0x9f, 0x0a, 0x41, 0xc2, 0x04,
	0x68, 0x45, 0x5a, 0x81, 0x2a, 0x90, 0xbd, 0xe8, 0x25, 0xc8, 0x2c, 0xed, 0xc9, 0x17, 0x7e, 0x94,
	0x2f, 0xe4, 0x36, 0x5e, 0xb0, 0xa1, 0x78, 0x19, 0xf2, 0xd9, 0x72, 0x2a, 0x6c, 0x43, 0xa9, 0x8b,
	0x41, 0x2c, 0xa2, 0x06, 0x94, 0x12, 0x4b, 0x97, 0xcf, 0x88, 0x9d, 0x57, 0xd6, 0x52, 0xb0, 0x56,
	0xa1, 0xcf, 0x01, 0xc6, 0xcb, 0x89, 0x86, 0x17, 0x94, 0x03, 0x44, 0xa9, 0x46, 0xda, 0xbb, 0x3c,
	0x12, 0x5d, 0xc2, 0x64, 0x78, 0xc5, 0xd8, 0x3e, 0x4d, 0x35, 0xe8, 0x0b, 0x28, 0x5f, 0x86, 0xd4,
	0xc7, 0xe1, 0xca, 0x25, 0x0b, 0xca, 0x99, 0x6d, 0x2a, 0x97, 0x72, 0xb0, 0xa9, 0x94, 0x51, 0xdf,
	0x12, 0x1c, 0xda, 0x56, 0x23, 0xd3, 0xcc, 0xbb, 0xc6, 0x8a, 0xe0, 0x50, 0x46, 0x7d, 0xb5, 0xc4,
	0xa1, 0xa4, 0x05, 0xa5, 0x2e, 0xfe, 0x1a, 0x8b, 0xa8, 0x0e, 0xe6, 0x59, 0x10, 0x84, 0xfc, 0x9e,
	0xcc, 0xec, 0x52, 0x23, 0xd3, 0x34, 0x5d, 0x13, 0x27, 0xb2, 0xb4, 0x0d, 0x38, 0x5b, 0x5c, 0x63,
	0xea, 0xd9, 0xaf, 0x62, 0x9b, 0x97, 0xc8, 0xb2, 0x92, 0x43, 0x32, 0xa3, 0x58, 0xe5, 0x5a, 0x8e,
	0x2b, 0xe9, 0x6b, 0x05, 0x3a, 0x87, 0x8a, 0xaa, 0x73, 0x14, 0x51, 0xce, 0x7a, 0x58, 0x10, 0xbb,
	0xd2, 0xc8, 0x34, 0x4b, 0xa7, 0xf5, 0x56, 0xdc, 0xe7, 0x96, 0xee, 0x73, 0xeb, 0x5a, 0xf7, 0xd9,
	0xad, 0x44, 0x5b, 0x08, 0xf4, 0x2d, 0x58, 0x5d, 0x8f, 0x47, 0x44, 0xc1, 0x5f, 0xef, 0x85, 0x5b,
	0x53, 0xed, 0x8c, 0x06, 0x50, 0xbb, 0x60, 0x54, 0x50, 0xec, 0xc5, 0xa9, 0x61, 0x4f, 0x71, 0x54,
	0xf7, 0x72, 0xd4, 0xe8, 0x63, 0x98, 0xac, 0xc2, 0x0f, 0xc4, 0x0b, 0xae, 0xc9, 0x6f, 0xc2, 0x3e,
	0x52, 0x89, 0x9a, 0x77, 0x89, 0xec, 0x7c, 0x09, 0xc5, 0xa1, 0xdb, 0x1f, 0xd0, 0x48, 0x20, 0x07,
	0x8c, 0x61, 0x38, 0x8f, 0xec, 0x4c, 0x23, 0xd7, 0x2c, 0x9d, 0x56, 0x5a, 0xe9, 0xe4, 0xb7, 0x86,
	0x6e, 0xdf, 0x35, 0xfc, 0x70, 0x1e, 0x39, 0x1f, 0xa1, 0x30, 0xa6, 0x6c, 0xe1, 0x91, 0x74, 0xb4,
	0x36, 0x87, 0xf7, 0xc1, 0xf8, 0x39, 0x3d, 0x28, 0xf4, 0xf8, 0x72, 0xf2, 0x84, 0xf7, 0x5b, 0x28,
	0x5c, 0xee, 0x18, 0xf5, 0x84, 0x25, 0xa7, 0x59, 0x4e, 0xff, 0xb5, 0xc0, 0xfa, 0x5e, 0xc5, 0x22,
	0x17, 0x20, 0x0e, 0xf8, 0x82, 0xcd, 0x39, 0xaa, 0x6e, 0x84, 0xa8, 0x36, 0xaa, 0xfe, 0x20, 0x68,
	0xe7, 0x13, 0xf4, 0x11, 0x8c, 0xa1, 0xdb, 0x8f, 0xb6, 0x7c, 0xd5, 0x1a, 0x3e, 0xf6, 0xfd, 0x2a,
	0x83, 0x3a, 0x60, 0x9e, 0x2d, 0x08, 0x9b, 0x52, 0xb2, 0x0b, 0x71, 0xb4, 0xa1, 0x89, 0xab, 0xa0,
	0x40, 0x6d, 0x30, 0x06, 0xa3, 0xf3, 0x03, 0x00, 0x1d, 0x30, 0x93, 0x85, 0xd9, 0x07, 0x8a, 0xab,
	0xa7, 0x40, 0xdf, 0x40, 0x69, 0xbd, 0x7a, 0x07, 0xe0, 0x3a, 0x60, 0x76, 0xef, 0x30, 0x63, 0xc4,
	0x3b, 0x20, 0xc2, 0x13, 0x28, 0xa8, 0xdd, 0x38, 0xe0, 0x9d, 0xaf, 0xd5, 0x61, 0x3a, 0x14, 0x35,
	0x80, 0xd2, 0x0d, 0xf6, 0xe8, 0x0c, 0x0b, 0xd5, 0xdc, 0xfd, 0x1d, 0xfd, 0xf4, 0xcf, 0xbf, 0xff,
	0xfb, 0x2b, 0xfb, 0x06, 0xd5, 0xd4, 0x19, 0xbe, 0x3f, 0x69, 0xcb, 0xb1, 0x6c, 0xff, 0x2e, 0x07,
	0xea, 0x0f, 0x74, 0x03, 0x15, 0xcd, 0xa6, 0xda, 0xb8, 0xda, 0x41, 0xb8, 0x23, 0xe3, 0x0f, 0x8a,
	0xf3, 0x3d, 0x7a, 0xa7, 0x39, 0x71, 0x32, 0x03, 0x9a, 0xf7, 0x27, 0x78, 0xad, 0x79, 0x93, 0x5a,
	0xbe, 0x90, 0x78, 0x9a, 0x74, 0x42, 0x13, 0x8f, 0xd6, 0xe9, 0x0f, 0x46, 0xe7, 0xcf, 0x23, 0x7d,
	0x54, 0x01, 0x8f, 0x4f, 0x52, 0xc2, 0x31, 0x94, 0xd3, 0x7a, 0xca, 0x56, 0x3c, 0x8f, 0xf2, 0x33,
	0x45, 0xf9, 0x0e, 0xbd, 0x49, 0x8b, 0xaa, 0xda, 0xb8, 0x23, 0x7d, 0x7d, 0xe8, 0x5f, 0x96, 0x7e,
	0xf2, 0x2b, 0x49, 0x89, 0x6f, 0xa1, 0xaa, 0x89, 0xf5, 0xec, 0x3c, 0x8f, 0xb9, 0xa1, 0x98, 0xeb,
	0xc8, 0xd6, 0xcc, 0xf2, 0xf2, 0x6e, 0xc5, 0xfc, 0x33, 0xa0, 0x0d, 0xea, 0x83, 0xc2, 0x76, 0x14,
	0xf9, 0x31, 0xaa, 0x6f, 0x90, 0x3f, 0x8c, 0x7c, 0x04, 0xaf, 0xf4, 0x8f, 0xe5, 0x89, 0xf3, 0x82,
	0xb6, 0x07, 0x57, 0x5e, 0x58, 0xe7, 0xbd, 0x62, 0xae, 0xa1, 0xa3, 0x75, 0xeb, 0xd8, 0x42, 0x60,
	0xea, 0x45, 0x93, 0x82, 0x3a, 0xe6, 0x9d, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x01, 0x5b,
	0xa9, 0x82, 0x08, 0x00, 0x00,
}
