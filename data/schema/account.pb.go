// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/data/schema/account.proto
// DO NOT EDIT!

/*
Package schema is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/data/schema/account.proto

It has these top-level messages:
	Account
	OrganizationMember
	Resource
	ResourceSettings
	Team
	TeamMember
	Permission
	Key
*/
package schema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Account Schema Storage Types
type AccountType int32

const (
	AccountType_USER         AccountType = 0
	AccountType_ORGANIZATION AccountType = 1
)

var AccountType_name = map[int32]string{
	0: "USER",
	1: "ORGANIZATION",
}
var AccountType_value = map[string]int32{
	"USER":         0,
	"ORGANIZATION": 1,
}

func (x AccountType) String() string {
	return proto.EnumName(AccountType_name, int32(x))
}
func (AccountType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Reference Data
type ResourceType int32

const (
	ResourceType_STACK   ResourceType = 0
	ResourceType_SERVICE ResourceType = 1
)

var ResourceType_name = map[int32]string{
	0: "STACK",
	1: "SERVICE",
}
var ResourceType_value = map[string]int32{
	"STACK":   0,
	"SERVICE": 1,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}
func (ResourceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GrantType int32

const (
	GrantType_READ   GrantType = 0
	GrantType_WRITE  GrantType = 1
	GrantType_DELETE GrantType = 3
	GrantType_ALL    GrantType = 4
)

var GrantType_name = map[int32]string{
	0: "READ",
	1: "WRITE",
	3: "DELETE",
	4: "ALL",
}
var GrantType_value = map[string]int32{
	"READ":   0,
	"WRITE":  1,
	"DELETE": 3,
	"ALL":    4,
}

func (x GrantType) String() string {
	return proto.EnumName(GrantType_name, int32(x))
}
func (GrantType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Account struct {
	Id           string      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name         string      `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type         AccountType `protobuf:"varint,3,opt,name=type,enum=schema.AccountType" json:"type,omitempty"`
	Email        string      `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	PasswordHash string      `protobuf:"bytes,5,opt,name=password_hash,json=passwordHash" json:"password_hash,omitempty"`
	IsVerified   bool        `protobuf:"varint,6,opt,name=is_verified,json=isVerified" json:"is_verified,omitempty"`
	// TODO Billing Attributes;
	CreateDt *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Account) GetType() AccountType {
	if m != nil {
		return m.Type
	}
	return AccountType_USER
}

func (m *Account) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Account) GetPasswordHash() string {
	if m != nil {
		return m.PasswordHash
	}
	return ""
}

func (m *Account) GetIsVerified() bool {
	if m != nil {
		return m.IsVerified
	}
	return false
}

func (m *Account) GetCreateDt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateDt
	}
	return nil
}

type OrganizationMember struct {
	Id            string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	OrgAccountId  string                     `protobuf:"bytes,2,opt,name=org_account_id,json=orgAccountId" json:"org_account_id,omitempty"`
	UserAccountId string                     `protobuf:"bytes,3,opt,name=user_account_id,json=userAccountId" json:"user_account_id,omitempty"`
	CreateDt      *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
}

func (m *OrganizationMember) Reset()                    { *m = OrganizationMember{} }
func (m *OrganizationMember) String() string            { return proto.CompactTextString(m) }
func (*OrganizationMember) ProtoMessage()               {}
func (*OrganizationMember) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OrganizationMember) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrganizationMember) GetOrgAccountId() string {
	if m != nil {
		return m.OrgAccountId
	}
	return ""
}

func (m *OrganizationMember) GetUserAccountId() string {
	if m != nil {
		return m.UserAccountId
	}
	return ""
}

func (m *OrganizationMember) GetCreateDt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateDt
	}
	return nil
}

type Resource struct {
	Id           string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	OrgAccountId string                     `protobuf:"bytes,2,opt,name=org_account_id,json=orgAccountId" json:"org_account_id,omitempty"`
	TeamId       string                     `protobuf:"bytes,3,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
	Name         string                     `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Type         ResourceType               `protobuf:"varint,5,opt,name=type,enum=schema.ResourceType" json:"type,omitempty"`
	CreateDt     *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
}

func (m *Resource) Reset()                    { *m = Resource{} }
func (m *Resource) String() string            { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()               {}
func (*Resource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Resource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Resource) GetOrgAccountId() string {
	if m != nil {
		return m.OrgAccountId
	}
	return ""
}

func (m *Resource) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resource) GetType() ResourceType {
	if m != nil {
		return m.Type
	}
	return ResourceType_STACK
}

func (m *Resource) GetCreateDt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateDt
	}
	return nil
}

type ResourceSettings struct {
	Id         string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ResourceId string                     `protobuf:"bytes,2,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	Key        string                     `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Value      string                     `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	CreateDt   *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
}

func (m *ResourceSettings) Reset()                    { *m = ResourceSettings{} }
func (m *ResourceSettings) String() string            { return proto.CompactTextString(m) }
func (*ResourceSettings) ProtoMessage()               {}
func (*ResourceSettings) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ResourceSettings) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ResourceSettings) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *ResourceSettings) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ResourceSettings) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ResourceSettings) GetCreateDt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateDt
	}
	return nil
}

type Team struct {
	Id           string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	OrgAccountId string                     `protobuf:"bytes,2,opt,name=org_account_id,json=orgAccountId" json:"org_account_id,omitempty"`
	Name         string                     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Desc         string                     `protobuf:"bytes,4,opt,name=desc" json:"desc,omitempty"`
	CreateDt     *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
}

func (m *Team) Reset()                    { *m = Team{} }
func (m *Team) String() string            { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()               {}
func (*Team) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Team) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Team) GetOrgAccountId() string {
	if m != nil {
		return m.OrgAccountId
	}
	return ""
}

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Team) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Team) GetCreateDt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateDt
	}
	return nil
}

type TeamMember struct {
	Id            string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	UserAccountId string                     `protobuf:"bytes,2,opt,name=user_account_id,json=userAccountId" json:"user_account_id,omitempty"`
	TeamId        string                     `protobuf:"bytes,3,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
	CreateDt      *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
}

func (m *TeamMember) Reset()                    { *m = TeamMember{} }
func (m *TeamMember) String() string            { return proto.CompactTextString(m) }
func (*TeamMember) ProtoMessage()               {}
func (*TeamMember) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TeamMember) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TeamMember) GetUserAccountId() string {
	if m != nil {
		return m.UserAccountId
	}
	return ""
}

func (m *TeamMember) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *TeamMember) GetCreateDt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateDt
	}
	return nil
}

type Permission struct {
	Id         string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	TeamId     string                     `protobuf:"bytes,2,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
	ResourceId string                     `protobuf:"bytes,3,opt,name=resource_id,json=resourceId" json:"resource_id,omitempty"`
	GrantType  GrantType                  `protobuf:"varint,4,opt,name=grant_type,json=grantType,enum=schema.GrantType" json:"grant_type,omitempty"`
	CreateDt   *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
}

func (m *Permission) Reset()                    { *m = Permission{} }
func (m *Permission) String() string            { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()               {}
func (*Permission) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Permission) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Permission) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *Permission) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *Permission) GetGrantType() GrantType {
	if m != nil {
		return m.GrantType
	}
	return GrantType_READ
}

func (m *Permission) GetCreateDt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreateDt
	}
	return nil
}

// Protobuf Type to store the foreign key relationships
type Key struct {
	KeyId string `protobuf:"bytes,1,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
}

func (m *Key) Reset()                    { *m = Key{} }
func (m *Key) String() string            { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()               {}
func (*Key) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Key) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func init() {
	proto.RegisterType((*Account)(nil), "schema.Account")
	proto.RegisterType((*OrganizationMember)(nil), "schema.OrganizationMember")
	proto.RegisterType((*Resource)(nil), "schema.Resource")
	proto.RegisterType((*ResourceSettings)(nil), "schema.ResourceSettings")
	proto.RegisterType((*Team)(nil), "schema.Team")
	proto.RegisterType((*TeamMember)(nil), "schema.TeamMember")
	proto.RegisterType((*Permission)(nil), "schema.Permission")
	proto.RegisterType((*Key)(nil), "schema.Key")
	proto.RegisterEnum("schema.AccountType", AccountType_name, AccountType_value)
	proto.RegisterEnum("schema.ResourceType", ResourceType_name, ResourceType_value)
	proto.RegisterEnum("schema.GrantType", GrantType_name, GrantType_value)
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/data/schema/account.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 654 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xd3, 0x4a,
	0x18, 0xed, 0xc4, 0xce, 0xdf, 0x97, 0xb4, 0xd7, 0x77, 0x6e, 0xaf, 0x88, 0x2a, 0xa4, 0x46, 0x05,
	0x95, 0xd0, 0x85, 0x8d, 0x8a, 0x2a, 0xd6, 0x51, 0x6b, 0x15, 0xab, 0xa5, 0x45, 0x8e, 0x29, 0x12,
	0x1b, 0x6b, 0x62, 0x4f, 0x9d, 0x51, 0xe3, 0x1f, 0xcd, 0x4c, 0x8a, 0xc2, 0x83, 0xb0, 0x63, 0xc7,
	0x8e, 0xa7, 0xe0, 0x0d, 0x78, 0x0f, 0x5e, 0x02, 0xf9, 0xaf, 0x71, 0x48, 0xbb, 0x68, 0x76, 0xdf,
	0x7c, 0x73, 0x66, 0xe6, 0x9c, 0xef, 0x1c, 0x1b, 0x8e, 0x02, 0x26, 0x27, 0xb3, 0xb1, 0xee, 0xc5,
	0xa1, 0x41, 0x92, 0xc4, 0xa3, 0x53, 0xca, 0x89, 0x8c, 0xb9, 0x41, 0xc2, 0xc4, 0xf0, 0x89, 0x24,
	0x86, 0xf0, 0x26, 0x34, 0x24, 0x06, 0xf1, 0xbc, 0x78, 0x16, 0x49, 0x3d, 0xe1, 0xb1, 0x8c, 0x71,
	0x23, 0xef, 0xee, 0xec, 0x06, 0x71, 0x1c, 0x4c, 0xa9, 0x91, 0x75, 0xc7, 0xb3, 0x6b, 0x43, 0xb2,
	0x90, 0x0a, 0x49, 0xc2, 0x24, 0x07, 0xee, 0xfd, 0x46, 0xd0, 0x1c, 0xe6, 0x47, 0xf1, 0x16, 0xd4,
	0x98, 0xdf, 0x43, 0x7d, 0x34, 0x68, 0xdb, 0x35, 0xe6, 0x63, 0x0c, 0x6a, 0x44, 0x42, 0xda, 0xab,
	0x65, 0x9d, 0xac, 0xc6, 0x2f, 0x40, 0x95, 0xf3, 0x84, 0xf6, 0x94, 0x3e, 0x1a, 0x6c, 0x1d, 0xfe,
	0xa7, 0xe7, 0xef, 0xe8, 0xc5, 0x15, 0xce, 0x3c, 0xa1, 0x76, 0x06, 0xc0, 0xdb, 0x50, 0xa7, 0x21,
	0x61, 0xd3, 0x9e, 0x9a, 0x9d, 0xce, 0x17, 0xf8, 0x19, 0x6c, 0x26, 0x44, 0x88, 0xcf, 0x31, 0xf7,
	0xdd, 0x09, 0x11, 0x93, 0x5e, 0x3d, 0xdb, 0xed, 0x96, 0xcd, 0xb7, 0x44, 0x4c, 0xf0, 0x2e, 0x74,
	0x98, 0x70, 0x6f, 0x29, 0x67, 0xd7, 0x8c, 0xfa, 0xbd, 0x46, 0x1f, 0x0d, 0x5a, 0x36, 0x30, 0x71,
	0x55, 0x74, 0xf0, 0x1b, 0x68, 0x7b, 0x9c, 0x12, 0x49, 0x5d, 0x5f, 0xf6, 0x9a, 0x7d, 0x34, 0xe8,
	0x1c, 0xee, 0xe8, 0xb9, 0x52, 0xbd, 0x54, 0xaa, 0x3b, 0xa5, 0x52, 0xbb, 0x95, 0x83, 0x4f, 0xe4,
	0xde, 0x0f, 0x04, 0xf8, 0x92, 0x07, 0x24, 0x62, 0x5f, 0x88, 0x64, 0x71, 0xf4, 0x8e, 0x86, 0x63,
	0xca, 0x57, 0x84, 0x3f, 0x87, 0xad, 0x98, 0x07, 0x6e, 0x31, 0x52, 0x97, 0xf9, 0xc5, 0x08, 0xba,
	0x31, 0x0f, 0x0a, 0xa5, 0x96, 0x8f, 0xf7, 0xe1, 0x9f, 0x99, 0xa0, 0xbc, 0x0a, 0x53, 0x32, 0xd8,
	0x66, 0xda, 0x5e, 0xe0, 0x96, 0xd8, 0xaa, 0x8f, 0x60, 0xfb, 0x0b, 0x41, 0xcb, 0xa6, 0x22, 0x9e,
	0x71, 0x8f, 0xae, 0xc9, 0xf1, 0x09, 0x34, 0x25, 0x25, 0xe1, 0x82, 0x5b, 0x23, 0x5d, 0x5a, 0x0b,
	0x6f, 0xd5, 0x8a, 0xb7, 0x83, 0xc2, 0xdb, 0x7a, 0xe6, 0xed, 0x76, 0xe9, 0x6d, 0x49, 0xa1, 0x62,
	0xee, 0x92, 0xa4, 0xc6, 0x23, 0x24, 0x7d, 0x47, 0xa0, 0x95, 0xf7, 0x8d, 0xa8, 0x94, 0x2c, 0x0a,
	0xc4, 0x8a, 0xb4, 0x5d, 0xe8, 0xf0, 0x02, 0xb3, 0xd0, 0x05, 0x65, 0xcb, 0xf2, 0xb1, 0x06, 0xca,
	0x0d, 0x9d, 0x17, 0x8a, 0xd2, 0x32, 0x4d, 0xdb, 0x2d, 0x99, 0xce, 0x4a, 0x3d, 0xf9, 0x62, 0x99,
	0x66, 0xfd, 0x11, 0x34, 0xbf, 0x21, 0x50, 0x1d, 0x4a, 0xc2, 0x35, 0xa7, 0x5e, 0x0e, 0x57, 0xa9,
	0x0c, 0x17, 0x83, 0xea, 0x53, 0xe1, 0x95, 0x03, 0x4f, 0xeb, 0xf5, 0xf9, 0x7d, 0x45, 0x00, 0x29,
	0xbf, 0x07, 0xf2, 0x7b, 0x4f, 0x32, 0x6b, 0xf7, 0x25, 0xf3, 0xc1, 0x74, 0xac, 0x1d, 0xd9, 0x9f,
	0x08, 0xe0, 0x3d, 0xe5, 0x21, 0x13, 0x82, 0xc5, 0xd1, 0x0a, 0xb1, 0xca, 0x83, 0xb5, 0xa5, 0x07,
	0xff, 0xb2, 0x5c, 0x59, 0xb1, 0xfc, 0x15, 0x40, 0xc0, 0x49, 0x24, 0xdd, 0x2c, 0xa1, 0x6a, 0x96,
	0xd0, 0x7f, 0xcb, 0x84, 0x9e, 0xa6, 0x3b, 0x59, 0x3c, 0xdb, 0x41, 0x59, 0xae, 0x3f, 0xdc, 0xa7,
	0xa0, 0x9c, 0xd1, 0x39, 0xfe, 0x1f, 0x1a, 0x37, 0x74, 0xee, 0xde, 0xf1, 0xaf, 0xdf, 0xd0, 0xb9,
	0xe5, 0x1f, 0xbc, 0x84, 0x4e, 0xe5, 0x67, 0x87, 0x5b, 0xa0, 0x7e, 0x18, 0x99, 0xb6, 0xb6, 0x81,
	0x35, 0xe8, 0x5e, 0xda, 0xa7, 0xc3, 0x0b, 0xeb, 0xd3, 0xd0, 0xb1, 0x2e, 0x2f, 0x34, 0x74, 0xb0,
	0x0f, 0xdd, 0xea, 0xb7, 0x83, 0xdb, 0x50, 0x1f, 0x39, 0xc3, 0xe3, 0x33, 0x6d, 0x03, 0x77, 0xa0,
	0x39, 0x32, 0xed, 0x2b, 0xeb, 0xd8, 0xd4, 0xd0, 0xc1, 0x11, 0xb4, 0xef, 0x14, 0xa4, 0x17, 0xda,
	0xe6, 0xf0, 0x44, 0xdb, 0x48, 0xe1, 0x1f, 0x6d, 0xcb, 0x31, 0x35, 0x84, 0x01, 0x1a, 0x27, 0xe6,
	0xb9, 0xe9, 0x98, 0x9a, 0x82, 0x9b, 0xa0, 0x0c, 0xcf, 0xcf, 0x35, 0x75, 0xdc, 0xc8, 0x54, 0xbc,
	0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xc1, 0xce, 0xbd, 0x23, 0x06, 0x00, 0x00,
}
