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
	// 1096 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x58, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xee, 0x3a, 0x89, 0xdb, 0x1c, 0x93, 0x6d, 0xb2, 0x84, 0x10, 0x05, 0x2e, 0x90, 0x55, 0x89,
	0x08, 0x5a, 0x27, 0x6c, 0xd3, 0x96, 0x62, 0x52, 0x1a, 0xdb, 0x89, 0x15, 0x29, 0x86, 0xe2, 0x34,
	0x80, 0xb8, 0x9b, 0xec, 0x1e, 0x9c, 0x91, 0xd7, 0x33, 0xdb, 0xd9, 0x71, 0xc1, 0x6f, 0xc0, 0x2b,
	0xc0, 0x55, 0xb9, 0xaa, 0xc4, 0x35, 0x12, 0x0f, 0xc0, 0xeb, 0xf0, 0x10, 0x68, 0x66, 0x76, 0xd7,
	0x56, 0xec, 0xc0, 0x6c, 0x44, 0x7b, 0x51, 0xe5, 0xc6, 0xda, 0x59, 0x9f, 0xef, 0xfc, 0x7e, 0x73,
	0xe6, 0xec, 0xc0, 0xcd, 0x80, 0x0b, 0xdc, 0x52, 0x3f, 0xb5, 0x58, 0x70, 0xc9, 0xbd, 0x79, 0xf5,
	0xbc, 0xb1, 0x44, 0x62, 0xba, 0x45, 0x62, 0x6a, 0x5e, 0x56, 0x1f, 0xc3, 0x72, 0x1b, 0x65, 0x0b,
	0x9f, 0xd3, 0x00, 0xbb, 0xf8, 0x6c, 0x88, 0x89, 0xf4, 0xd6, 0xa0, 0x1c, 0x0a, 0xfa, 0x1c, 0xc5,
	0xba, 0xf3, 0x81, 0xb3, 0xb9, 0xd8, 0x4d, 0x57, 0xde, 0x2a, 0x2c, 0xf4, 0x04, 0x1f, 0xc6, 0xeb,
	0x25, 0xfd, 0xda, 0x2c, 0xaa, 0xb7, 0xc0, 0x6d, 0xa3, 0x4c, 0xb1, 0x5f, 0x92, 0x01, 0x7a, 0x1e,
	0xcc, 0x33, 0x32, 0xc0, 0x14, 0xad, 0x9f, 0xab, 0x8f, 0x60, 0xf9, 0xeb, 0x21, 0x8a, 0x51, 0x8b,
	0x48, 0x92, 0xd9, 0x59, 0x85, 0x05, 0x49, 0x4e, 0xa3, 0x4c, 0xd0, 0x2c, 0xd4, 0xdb, 0x67, 0x4a,
	0x52, 0x5b, 0x79, 0xab, 0x6b, 0x16, 0xd5, 0x47, 0xb0, 0xd6, 0x46, 0xf9, 0x95, 0x68, 0x61, 0x84,
	0x12, 0xff, 0x5b, 0x8b, 0x0b, 0x25, 0x1a, 0xa6, 0x8e, 0x96, 0x68, 0x58, 0xed, 0xc0, 0xca, 0x49,
	0x1c, 0x92, 0x4b, 0x40, 0x55, 0x38, 0x21, 0x91, 0x64, 0x7d, 0x4e, 0xfb, 0xa3, 0x9f, 0xab, 0xdf,
	0xc1, 0x5a, 0x67, 0x18, 0x49, 0x6a, 0xab, 0x73, 0x66, 0x50, 0x33, 0x35, 0xef, 0xc2, 0x4a, 0x53,
	0xa0, 0x95, 0xd2, 0x0c, 0x5e, 0x9a, 0x80, 0x6f, 0xc2, 0xf2, 0x11, 0xef, 0x51, 0x76, 0x92, 0xa0,
	0x98, 0x44, 0xf3, 0x3e, 0xb2, 0x1c, 0xad, 0x16, 0xd5, 0x0f, 0x55, 0x46, 0x22, 0x4e, 0xc2, 0x03,
	0x1a, 0xe5, 0xa5, 0xcf, 0x54, 0x3a, 0x13, 0x2a, 0x3f, 0x82, 0xd5, 0x16, 0xff, 0x91, 0x8d, 0x45,
	0x93, 0x98, 0xb3, 0x04, 0x67, 0xc9, 0xfa, 0xbf, 0x97, 0xc0, 0x3d, 0xa2, 0x01, 0xb2, 0x04, 0x8f,
	0x51, 0x28, 0x52, 0x79, 0x3b, 0x9a, 0x1f, 0x0a, 0x99, 0xfe, 0xe1, 0xad, 0xd4, 0x14, 0xff, 0x34,
	0x1d, 0x52, 0xbb, 0x1b, 0x4b, 0xfa, 0x55, 0xa6, 0xba, 0x7a, 0xcd, 0xdb, 0x06, 0x38, 0x49, 0x0a,
	0x21, 0xee, 0xc1, 0xcd, 0x03, 0xca, 0xc2, 0x0e, 0x09, 0xce, 0x28, 0xc3, 0x26, 0x0f, 0xed, 0x60,
	0x9f, 0x9b, 0x0d, 0xa0, 0x19, 0x9e, 0x99, 0x7b, 0x57, 0x0b, 0x4d, 0xf0, 0xed, 0x42, 0x74, 0x1d,
	0x96, 0x4c, 0x12, 0xc7, 0x50, 0xbd, 0xe3, 0xa6, 0x32, 0x3b, 0x05, 0xdd, 0x74, 0xfc, 0x5f, 0x1d,
	0x80, 0xbd, 0x38, 0xce, 0x12, 0x55, 0x83, 0x1b, 0x6d, 0x94, 0x4f, 0x55, 0x71, 0x52, 0xcf, 0xf5,
	0xf3, 0x85, 0xb6, 0x3f, 0x86, 0x05, 0x1d, 0x9a, 0x55, 0x98, 0x5b, 0x30, 0xd7, 0x46, 0x69, 0x1f,
	0x99, 0xff, 0x8b, 0x03, 0x70, 0xc4, 0x7b, 0x99, 0x73, 0x77, 0xa0, 0x6c, 0x68, 0xe9, 0x79, 0x5a,
	0xd2, 0x2c, 0xfe, 0x37, 0x73, 0x85, 0x82, 0xf1, 0xff, 0x2e, 0x41, 0x45, 0x11, 0x3c, 0x73, 0xee,
	0x95, 0x66, 0xc2, 0xf3, 0xa1, 0x6c, 0x24, 0x0a, 0x60, 0xee, 0x40, 0xd9, 0xb4, 0x86, 0x34, 0x5d,
	0x66, 0x71, 0xa1, 0x78, 0x0d, 0xae, 0x77, 0x31, 0x8e, 0x48, 0x60, 0x29, 0x5f, 0xb0, 0x1a, 0xbb,
	0xe0, 0xb5, 0x51, 0x36, 0x87, 0x42, 0x20, 0x93, 0x2a, 0x71, 0x87, 0xec, 0x07, 0xee, 0xad, 0x19,
	0xaa, 0x9e, 0x6f, 0x17, 0xd3, 0xe9, 0x7e, 0x31, 0x0f, 0xde, 0x53, 0xd5, 0x71, 0x8e, 0x83, 0x33,
	0x1c, 0x90, 0x4b, 0x65, 0xfd, 0x00, 0xde, 0x33, 0xfd, 0x5f, 0x9f, 0x34, 0x8d, 0x91, 0xd9, 0x70,
	0x7b, 0x2c, 0x6c, 0xab, 0x43, 0x24, 0xf3, 0xe5, 0xfc, 0x51, 0x34, 0xad, 0xe7, 0x2e, 0x2c, 0x69,
	0x3d, 0xfb, 0x83, 0x61, 0x44, 0x24, 0x17, 0x57, 0x25, 0xbf, 0xb0, 0x7f, 0x1e, 0x4b, 0x22, 0x13,
	0xfb, 0x0a, 0xd5, 0x4d, 0xff, 0x6c, 0xf2, 0xc1, 0x80, 0xb0, 0xb0, 0x31, 0x3a, 0x6c, 0x15, 0xe8,
	0x16, 0x7f, 0x01, 0x2c, 0x6b, 0x8a, 0xa8, 0x53, 0x2b, 0x23, 0xc8, 0x76, 0x46, 0x90, 0xb4, 0xba,
	0xe7, 0x07, 0x80, 0x69, 0x1f, 0x1e, 0x82, 0xab, 0x85, 0x1a, 0x23, 0xad, 0xec, 0x30, 0xb4, 0x87,
	0xee, 0x98, 0x1a, 0xbf, 0x9f, 0x13, 0x69, 0xc6, 0xac, 0x30, 0x8d, 0x7a, 0x90, 0x17, 0xba, 0x20,
	0xf0, 0x1e, 0x80, 0x91, 0xea, 0x10, 0x56, 0x20, 0x40, 0x3f, 0x27, 0x49, 0x7e, 0x50, 0x9c, 0x1b,
	0x20, 0x66, 0x51, 0x3e, 0x67, 0x8a, 0x3d, 0xc8, 0xcf, 0xe9, 0x92, 0x62, 0xa6, 0x86, 0x8a, 0x69,
	0xcc, 0x7d, 0x00, 0x23, 0xa5, 0x63, 0xb2, 0xc7, 0xed, 0xc1, 0x86, 0x62, 0x4e, 0x5e, 0xff, 0x16,
	0xc6, 0xb2, 0x31, 0x52, 0xbf, 0x87, 0xad, 0xc4, 0x8e, 0xae, 0x75, 0x00, 0x13, 0x94, 0x36, 0x9d,
	0xd6, 0x62, 0xf6, 0x84, 0x35, 0x8b, 0x35, 0x6f, 0xb7, 0x51, 0x7e, 0x4b, 0x04, 0xa3, 0xac, 0x77,
	0x40, 0x23, 0x89, 0x42, 0x19, 0x36, 0xa4, 0xdf, 0x1f, 0xc4, 0xf2, 0xdf, 0x4e, 0xff, 0x8a, 0x71,
	0xbd, 0xa7, 0x09, 0x5f, 0x90, 0x04, 0x8f, 0xa7, 0xb7, 0x4c, 0x41, 0x0d, 0x3b, 0xb0, 0x98, 0x12,
	0xbe, 0xd5, 0xb0, 0x67, 0xd1, 0xa7, 0x70, 0xbd, 0x8d, 0x52, 0x63, 0x0a, 0xda, 0xab, 0x67, 0xb4,
	0xbd, 0x0c, 0xf8, 0x21, 0xb8, 0x63, 0xce, 0x17, 0xf3, 0xf8, 0x7e, 0x56, 0x5f, 0x0d, 0xb3, 0xa7,
	0xf1, 0x03, 0xa8, 0xa4, 0xdc, 0x2f, 0x08, 0xcc, 0xb9, 0x3c, 0x89, 0xb3, 0xe0, 0xf2, 0x67, 0xe0,
	0x8e, 0xf7, 0x40, 0x41, 0xec, 0x17, 0xe0, 0x8e, 0x49, 0x3c, 0x99, 0x60, 0x4b, 0x22, 0xfb, 0x7f,
	0x94, 0xd2, 0x83, 0xb6, 0x8b, 0x01, 0x17, 0xe1, 0xd5, 0x78, 0x63, 0xd5, 0x3c, 0xfc, 0x08, 0xdc,
	0x0e, 0x26, 0x09, 0xe9, 0xe1, 0x25, 0xa7, 0xd5, 0x42, 0xc3, 0xe7, 0x4b, 0x07, 0x96, 0x55, 0x15,
	0xb5, 0x54, 0x66, 0xf0, 0x13, 0x80, 0x27, 0x3c, 0x91, 0x47, 0x44, 0xea, 0xaf, 0x28, 0x1b, 0xa3,
	0x66, 0xdc, 0xb7, 0x2f, 0xec, 0x36, 0x2c, 0x2a, 0x13, 0x06, 0x60, 0x95, 0x97, 0xdf, 0x1c, 0xa8,
	0x74, 0x79, 0x84, 0xaf, 0x87, 0x47, 0x3b, 0xe0, 0xee, 0x85, 0x03, 0xca, 0x94, 0xc5, 0xe6, 0x19,
	0x06, 0x7d, 0x9b, 0x06, 0xec, 0x33, 0x70, 0x9b, 0x44, 0x92, 0x68, 0xfc, 0xa5, 0xf1, 0x6a, 0x3f,
	0x6b, 0xfa, 0x50, 0x51, 0x07, 0xd3, 0xeb, 0x31, 0xb6, 0x0b, 0xee, 0x31, 0x4a, 0x49, 0xd9, 0xa5,
	0x82, 0xf3, 0xff, 0x2c, 0xc1, 0x3b, 0xc7, 0xa3, 0x44, 0xe2, 0xe0, 0x1b, 0x22, 0xa8, 0x1e, 0xc0,
	0xaf, 0x3a, 0x82, 0x15, 0xf3, 0x5f, 0xce, 0xc1, 0x52, 0x83, 0x04, 0xfd, 0x61, 0xfc, 0xc6, 0x64,
	0xec, 0x36, 0x94, 0x0f, 0x07, 0x31, 0x17, 0xd2, 0x2a, 0x82, 0xdb, 0x50, 0xde, 0xff, 0xc9, 0x5a,
	0x7a, 0x47, 0xb9, 0x12, 0x71, 0x12, 0x16, 0xb9, 0xbe, 0xf0, 0x9a, 0x70, 0x23, 0xbb, 0x17, 0xba,
	0x38, 0xec, 0x0d, 0xa3, 0x70, 0xd6, 0x05, 0x52, 0xf5, 0xda, 0xb6, 0x93, 0x1e, 0x79, 0x49, 0xbf,
	0x43, 0x18, 0xe9, 0x5d, 0x7d, 0xd1, 0x5b, 0x12, 0xbc, 0x51, 0x87, 0x5b, 0x01, 0xab, 0x11, 0x2a,
	0x28, 0x97, 0xb5, 0x24, 0xec, 0xd7, 0x82, 0x88, 0x22, 0x93, 0xb5, 0x70, 0x78, 0x7a, 0xca, 0x6b,
	0x3d, 0x11, 0x07, 0x3a, 0xef, 0x4f, 0x9c, 0xef, 0x2b, 0x35, 0x7d, 0x07, 0x5c, 0x57, 0x3f, 0x3f,
	0x3b, 0xce, 0x0b, 0xc7, 0x39, 0x2d, 0xeb, 0xab, 0xdf, 0xbb, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x13, 0x12, 0xbb, 0x85, 0x22, 0x16, 0x00, 0x00,
}
