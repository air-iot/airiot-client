// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/core.proto

/*
Package core is a generated protocol buffer package.

It is generated from these files:

	core/core.proto

It has these top-level messages:

	GetDeviceRequest
	GetRequestName
	QueryDataRequest
	GetOrDeleteDataRequest
	UpdateDataRequest
	MultiUpdateDataRequest
	CreateDataRequest
	LoginUserRequest
	UploadFileRequest
	DownloadFileResponse
*/
package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/air-iot/api-client-go/v4/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetDeviceRequest struct {
	Driver string `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	Group  string `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
}

func (m *GetDeviceRequest) Reset()                    { *m = GetDeviceRequest{} }
func (m *GetDeviceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDeviceRequest) ProtoMessage()               {}
func (*GetDeviceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetDeviceRequest) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *GetDeviceRequest) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

type GetRequestName struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetRequestName) Reset()                    { *m = GetRequestName{} }
func (m *GetRequestName) String() string            { return proto.CompactTextString(m) }
func (*GetRequestName) ProtoMessage()               {}
func (*GetRequestName) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetRequestName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type QueryDataRequest struct {
	Table string `protobuf:"bytes,1,opt,name=table" json:"table,omitempty"`
	Query []byte `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
}

func (m *QueryDataRequest) Reset()                    { *m = QueryDataRequest{} }
func (m *QueryDataRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryDataRequest) ProtoMessage()               {}
func (*QueryDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *QueryDataRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *QueryDataRequest) GetQuery() []byte {
	if m != nil {
		return m.Query
	}
	return nil
}

type GetOrDeleteDataRequest struct {
	Table string `protobuf:"bytes,1,opt,name=table" json:"table,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *GetOrDeleteDataRequest) Reset()                    { *m = GetOrDeleteDataRequest{} }
func (m *GetOrDeleteDataRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOrDeleteDataRequest) ProtoMessage()               {}
func (*GetOrDeleteDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetOrDeleteDataRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *GetOrDeleteDataRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateDataRequest struct {
	Table string `protobuf:"bytes,1,opt,name=table" json:"table,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Data  []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *UpdateDataRequest) Reset()                    { *m = UpdateDataRequest{} }
func (m *UpdateDataRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateDataRequest) ProtoMessage()               {}
func (*UpdateDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateDataRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *UpdateDataRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateDataRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type MultiUpdateDataRequest struct {
	Table string `protobuf:"bytes,1,opt,name=table" json:"table,omitempty"`
	Query []byte `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Data  []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MultiUpdateDataRequest) Reset()                    { *m = MultiUpdateDataRequest{} }
func (m *MultiUpdateDataRequest) String() string            { return proto.CompactTextString(m) }
func (*MultiUpdateDataRequest) ProtoMessage()               {}
func (*MultiUpdateDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MultiUpdateDataRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *MultiUpdateDataRequest) GetQuery() []byte {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *MultiUpdateDataRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CreateDataRequest struct {
	Table string `protobuf:"bytes,1,opt,name=table" json:"table,omitempty"`
	Data  []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *CreateDataRequest) Reset()                    { *m = CreateDataRequest{} }
func (m *CreateDataRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDataRequest) ProtoMessage()               {}
func (*CreateDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CreateDataRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *CreateDataRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type LoginUserRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginUserRequest) Reset()                    { *m = LoginUserRequest{} }
func (m *LoginUserRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginUserRequest) ProtoMessage()               {}
func (*LoginUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LoginUserRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UploadFileRequest struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *UploadFileRequest) Reset()                    { *m = UploadFileRequest{} }
func (m *UploadFileRequest) String() string            { return proto.CompactTextString(m) }
func (*UploadFileRequest) ProtoMessage()               {}
func (*UploadFileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UploadFileRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type DownloadFileResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *DownloadFileResponse) Reset()                    { *m = DownloadFileResponse{} }
func (m *DownloadFileResponse) String() string            { return proto.CompactTextString(m) }
func (*DownloadFileResponse) ProtoMessage()               {}
func (*DownloadFileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DownloadFileResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GetDeviceRequest)(nil), "core.GetDeviceRequest")
	proto.RegisterType((*GetRequestName)(nil), "core.GetRequestName")
	proto.RegisterType((*QueryDataRequest)(nil), "core.QueryDataRequest")
	proto.RegisterType((*GetOrDeleteDataRequest)(nil), "core.GetOrDeleteDataRequest")
	proto.RegisterType((*UpdateDataRequest)(nil), "core.UpdateDataRequest")
	proto.RegisterType((*MultiUpdateDataRequest)(nil), "core.MultiUpdateDataRequest")
	proto.RegisterType((*CreateDataRequest)(nil), "core.CreateDataRequest")
	proto.RegisterType((*LoginUserRequest)(nil), "core.LoginUserRequest")
	proto.RegisterType((*UploadFileRequest)(nil), "core.UploadFileRequest")
	proto.RegisterType((*DownloadFileResponse)(nil), "core.DownloadFileResponse")
}

func init() { proto.RegisterFile("core/core.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1010 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x58, 0x41, 0x6f, 0x1b, 0x45,
	0x14, 0xee, 0x3a, 0x8d, 0x69, 0x5f, 0xb0, 0xeb, 0x2c, 0x21, 0x44, 0x81, 0x03, 0x5a, 0x55, 0x22,
	0x82, 0x76, 0x13, 0xb6, 0x29, 0x08, 0x85, 0x54, 0xc4, 0x76, 0x62, 0x59, 0x8a, 0xa1, 0xd8, 0x0d,
	0x20, 0x6e, 0x93, 0xdd, 0x87, 0x33, 0xf2, 0x7a, 0x66, 0x3b, 0x3b, 0x2e, 0xf8, 0x1f, 0xf0, 0x17,
	0xe0, 0x04, 0xa7, 0x4a, 0x9c, 0x91, 0xf8, 0x41, 0x5c, 0xb9, 0xf1, 0x07, 0xd0, 0xcc, 0xec, 0x6e,
	0xa2, 0xd8, 0x86, 0xd9, 0x88, 0xf6, 0x50, 0xe5, 0xb2, 0xda, 0x59, 0xbf, 0xef, 0xbd, 0x79, 0xef,
	0xfb, 0x66, 0xde, 0x8c, 0xe1, 0x4e, 0xc8, 0x05, 0x6e, 0xab, 0x87, 0x9f, 0x08, 0x2e, 0xb9, 0x7b,
	0x53, 0xbd, 0x6f, 0xd6, 0x48, 0x42, 0xb7, 0x49, 0x42, 0xcd, 0x47, 0xef, 0x33, 0x68, 0x74, 0x50,
	0xb6, 0xf1, 0x19, 0x0d, 0xb1, 0x8f, 0x4f, 0x27, 0x98, 0x4a, 0x77, 0x1d, 0xaa, 0x91, 0xa0, 0xcf,
	0x50, 0x6c, 0x38, 0xef, 0x3a, 0x5b, 0xb7, 0xfb, 0xd9, 0xc8, 0x5d, 0x83, 0xe5, 0xa1, 0xe0, 0x93,
	0x64, 0xa3, 0xa2, 0x3f, 0x9b, 0x81, 0x77, 0x17, 0xea, 0x1d, 0x94, 0x19, 0xf6, 0x73, 0x32, 0x46,
	0xd7, 0x85, 0x9b, 0x8c, 0x8c, 0x31, 0x43, 0xeb, 0x77, 0xef, 0x11, 0x34, 0xbe, 0x9c, 0xa0, 0x98,
	0xb6, 0x89, 0x24, 0x79, 0x9c, 0x35, 0x58, 0x96, 0xe4, 0x34, 0xce, 0x0d, 0xcd, 0x40, 0x7d, 0x7d,
	0xaa, 0x2c, 0x75, 0x94, 0xd7, 0xfb, 0x66, 0xe0, 0x3d, 0x82, 0xf5, 0x0e, 0xca, 0x2f, 0x44, 0x1b,
	0x63, 0x94, 0xf8, 0xdf, 0x5e, 0xea, 0x50, 0xa1, 0x51, 0x36, 0xd1, 0x0a, 0x8d, 0xbc, 0x1e, 0xac,
	0x9e, 0x24, 0x11, 0xb9, 0x02, 0x54, 0xa5, 0x13, 0x11, 0x49, 0x36, 0x96, 0xf4, 0x7c, 0xf4, 0xbb,
	0xf7, 0x0d, 0xac, 0xf7, 0x26, 0xb1, 0xa4, 0xb6, 0x3e, 0xe7, 0x26, 0x35, 0xd7, 0xf3, 0x3e, 0xac,
	0xb6, 0x04, 0x5a, 0x39, 0xcd, 0xe1, 0x95, 0x0b, 0xf0, 0x2d, 0x68, 0x1c, 0xf3, 0x21, 0x65, 0x27,
	0x29, 0x8a, 0x8b, 0x68, 0x3e, 0x42, 0x56, 0xa0, 0xd5, 0xc0, 0x7b, 0x4f, 0x55, 0x24, 0xe6, 0x24,
	0x3a, 0xa2, 0x71, 0x41, 0x7d, 0xee, 0xd2, 0xb9, 0xe0, 0xf2, 0x7d, 0x58, 0x6b, 0xf3, 0xef, 0xd9,
	0xb9, 0x69, 0x9a, 0x70, 0x96, 0xe2, 0x3c, 0xdb, 0xe0, 0xb7, 0x0a, 0xd4, 0x8f, 0x69, 0x88, 0x2c,
	0xc5, 0x01, 0x0a, 0x25, 0x2a, 0x77, 0x57, 0xeb, 0x43, 0x21, 0xb3, 0x1f, 0xdc, 0x55, 0x5f, 0xe9,
	0x4f, 0xcb, 0x21, 0x8b, 0xbb, 0x59, 0xd3, 0x9f, 0x72, 0xd7, 0xde, 0x0d, 0x77, 0x07, 0xe0, 0x24,
	0x2d, 0x85, 0x78, 0x08, 0x77, 0x8e, 0x28, 0x8b, 0x7a, 0x24, 0x3c, 0xa3, 0x0c, 0x5b, 0x3c, 0xb2,
	0x83, 0x7d, 0x6a, 0x16, 0x80, 0x56, 0x78, 0x1e, 0xee, 0x2d, 0x6d, 0x74, 0x41, 0x6f, 0x0b, 0xd1,
	0x7b, 0x50, 0x33, 0x45, 0x3c, 0x87, 0xea, 0x15, 0x37, 0x53, 0xd9, 0x19, 0xe8, 0x96, 0x13, 0xfc,
	0xec, 0x00, 0x1c, 0x24, 0x49, 0x5e, 0x28, 0x1f, 0x6e, 0x75, 0x50, 0x3e, 0x51, 0xe4, 0x64, 0x33,
	0xd7, 0xef, 0x0b, 0x63, 0x7f, 0x00, 0xcb, 0x3a, 0x35, 0xab, 0x34, 0xb7, 0x61, 0xa9, 0x83, 0xd2,
	0x3e, 0xb3, 0xe0, 0x27, 0x07, 0xe0, 0x98, 0x0f, 0xf3, 0xc9, 0xdd, 0x87, 0xaa, 0x91, 0xa5, 0xeb,
	0x6a, 0x4b, 0x33, 0xf8, 0xdf, 0xc2, 0x95, 0x4a, 0x26, 0xf8, 0xab, 0x02, 0x2b, 0x4a, 0xe0, 0xf9,
	0xe4, 0x5e, 0x68, 0x25, 0xdc, 0x00, 0xaa, 0xc6, 0xa2, 0x04, 0xe6, 0x3e, 0x54, 0xcd, 0xd6, 0x90,
	0x95, 0xcb, 0x0c, 0x16, 0x9a, 0xfb, 0xf0, 0x5a, 0x1f, 0x93, 0x98, 0x84, 0x96, 0xf6, 0x25, 0xd9,
	0xd8, 0x07, 0xb7, 0x83, 0xb2, 0x35, 0x11, 0x02, 0x99, 0x54, 0x85, 0xeb, 0xb2, 0xef, 0xb8, 0xbb,
	0x6e, 0xa4, 0x7a, 0x79, 0xbb, 0x98, 0x2d, 0xf7, 0x9f, 0x4b, 0xe0, 0x3e, 0x51, 0x3b, 0xce, 0x20,
	0x3c, 0xc3, 0x31, 0xb9, 0x52, 0xd5, 0x8f, 0xe0, 0x6d, 0xb3, 0xff, 0xeb, 0x4e, 0xd3, 0x9c, 0x9a,
	0x05, 0x77, 0xc0, 0xa2, 0x8e, 0x6a, 0x22, 0xf9, 0x5c, 0x2e, 0xb7, 0xa2, 0x59, 0x3f, 0x0f, 0xa0,
	0xa6, 0xfd, 0x1c, 0x8e, 0x27, 0x31, 0x91, 0x5c, 0x5c, 0x53, 0xbe, 0x70, 0xff, 0x1c, 0x48, 0x22,
	0x53, 0xfb, 0x45, 0xf5, 0xf7, 0x32, 0x34, 0x34, 0xcb, 0xaa, 0xf1, 0xe4, 0x1c, 0xef, 0xe4, 0x1c,
	0x67, 0x04, 0x5d, 0xee, 0xe1, 0xb3, 0x81, 0x3f, 0x81, 0xba, 0x36, 0x6a, 0x4e, 0xb5, 0xb3, 0x6e,
	0x64, 0x0f, 0xdd, 0x35, 0x34, 0xbd, 0x53, 0x68, 0x61, 0x4e, 0xbb, 0x9f, 0x45, 0x7d, 0x5c, 0x70,
	0x55, 0x12, 0xf8, 0x10, 0xc0, 0x58, 0xf5, 0x08, 0x2b, 0x91, 0x60, 0x50, 0xf0, 0x5c, 0xec, 0xf5,
	0x97, 0xce, 0x00, 0xf3, 0x54, 0x5b, 0x90, 0x6d, 0x0f, 0x0a, 0x0a, 0xc6, 0x33, 0xcc, 0xcc, 0xb9,
	0x60, 0x16, 0xf3, 0x11, 0x80, 0xb1, 0xd2, 0x39, 0xd9, 0xe3, 0x0e, 0x60, 0x53, 0x35, 0xcf, 0x82,
	0xff, 0x36, 0x26, 0xb2, 0x39, 0x55, 0xcf, 0x6e, 0x3b, 0xb5, 0x53, 0xdc, 0x1e, 0x80, 0x49, 0x4a,
	0x87, 0xce, 0xb8, 0x98, 0x7f, 0x48, 0x9a, 0xa7, 0x9a, 0x37, 0x3a, 0x28, 0xbf, 0x26, 0x82, 0x51,
	0x36, 0x3c, 0xa2, 0xb1, 0x44, 0xa1, 0x02, 0x1b, 0xdd, 0x1e, 0x8e, 0x13, 0xf9, 0x6f, 0x0d, 0x7c,
	0xc5, 0x4c, 0x7d, 0xd8, 0x9c, 0x76, 0xdb, 0x25, 0x45, 0x10, 0xfc, 0x5e, 0xc9, 0xf6, 0xb6, 0x3e,
	0x86, 0x5c, 0x44, 0xd7, 0x1d, 0xc5, 0x8a, 0xec, 0x20, 0x86, 0x7a, 0x0f, 0xd3, 0x94, 0x0c, 0xf1,
	0x8a, 0x07, 0x84, 0x52, 0xfd, 0xfe, 0xb9, 0x03, 0x0d, 0xc5, 0xa2, 0xb6, 0xca, 0x03, 0x7e, 0x08,
	0xf0, 0x98, 0xa7, 0xf2, 0x98, 0x48, 0x7d, 0x70, 0xb5, 0x09, 0x6a, 0x4e, 0x58, 0xf6, 0xc4, 0xee,
	0xc0, 0x6d, 0x15, 0xc2, 0x00, 0xac, 0xea, 0xf2, 0xab, 0x03, 0x2b, 0x7d, 0x1e, 0xe3, 0xcb, 0xd1,
	0xd1, 0x2e, 0xd4, 0x0f, 0xa2, 0x31, 0x65, 0x2a, 0x62, 0xeb, 0x0c, 0xc3, 0x91, 0xcd, 0x82, 0x09,
	0x18, 0xd4, 0x5b, 0x44, 0x92, 0xf8, 0xfc, 0x70, 0xf7, 0x62, 0x4f, 0x92, 0x23, 0x58, 0x51, 0x1b,
	0xc9, 0xcb, 0x09, 0xb6, 0x0f, 0xf5, 0x01, 0x4a, 0x49, 0xd9, 0x95, 0x92, 0x0b, 0xfe, 0xa8, 0xc0,
	0x9b, 0x83, 0x69, 0x2a, 0x71, 0xfc, 0x15, 0x11, 0x54, 0x9f, 0x79, 0xae, 0x77, 0x04, 0x2b, 0xe5,
	0x3f, 0x5f, 0x82, 0x5a, 0x93, 0x84, 0xa3, 0x49, 0xf2, 0xca, 0x54, 0xec, 0x1e, 0x54, 0xbb, 0xe3,
	0x84, 0x0b, 0x69, 0x95, 0xc1, 0x3d, 0xa8, 0x1e, 0xfe, 0x60, 0x6d, 0xbd, 0xab, 0xa6, 0xa2, 0xee,
	0x88, 0x65, 0x6e, 0x8c, 0x6e, 0x0b, 0x6e, 0xe5, 0x57, 0xf1, 0xc5, 0x69, 0x6f, 0x1a, 0x87, 0xf3,
	0xee, 0xec, 0xde, 0x8d, 0x1d, 0x27, 0x6b, 0x79, 0xe9, 0xa8, 0x47, 0x18, 0x19, 0x5e, 0x5f, 0xa2,
	0x2c, 0x05, 0xde, 0xdc, 0x83, 0xbb, 0x21, 0xf3, 0x09, 0x15, 0x94, 0x4b, 0x3f, 0x8d, 0x46, 0x7e,
	0x18, 0x53, 0x64, 0xd2, 0x8f, 0x26, 0xa7, 0xa7, 0xdc, 0x1f, 0x8a, 0x24, 0xd4, 0x75, 0x7f, 0xec,
	0x7c, 0xbb, 0xe2, 0xeb, 0xbf, 0xdd, 0xf6, 0xd4, 0xe3, 0x47, 0xc7, 0xf9, 0xc5, 0x71, 0x4e, 0xab,
	0xfa, 0xdf, 0xb6, 0x07, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x15, 0x22, 0xf8, 0x31, 0x95, 0x13,
	0x00, 0x00,
}
