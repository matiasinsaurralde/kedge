// Code generated by protoc-gen-go. DO NOT EDIT.
// source: winch/config/auth.proto

/*
Package winch_config is a generated protocol buffer package.

It is generated from these files:
	winch/config/auth.proto
	winch/config/mapper.proto

It has these top-level messages:
	AuthConfig
	AuthSource
	KubernetesAccess
	OIDCAccess
	DummyAccess
	TokenAccess
	MapperConfig
	Route
	DirectRoute
	RegexpRoute
*/
package winch_config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// / AuthConfig is the top level configuration message for a winch auth.
type AuthConfig struct {
	AuthSources []*AuthSource `protobuf:"bytes,1,rep,name=auth_sources,json=authSources" json:"auth_sources,omitempty"`
}

func (m *AuthConfig) Reset()                    { *m = AuthConfig{} }
func (m *AuthConfig) String() string            { return proto.CompactTextString(m) }
func (*AuthConfig) ProtoMessage()               {}
func (*AuthConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthConfig) GetAuthSources() []*AuthSource {
	if m != nil {
		return m.AuthSources
	}
	return nil
}

// / AuthSource specifies the kind of the backend auth we need to inject on winch reqeuest.
type AuthSource struct {
	// name is an ID of auth source. It can be referenced inside winch routing.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*AuthSource_Dummy
	//	*AuthSource_Kube
	//	*AuthSource_Oidc
	//	*AuthSource_Token
	Type isAuthSource_Type `protobuf_oneof:"type"`
}

func (m *AuthSource) Reset()                    { *m = AuthSource{} }
func (m *AuthSource) String() string            { return proto.CompactTextString(m) }
func (*AuthSource) ProtoMessage()               {}
func (*AuthSource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isAuthSource_Type interface {
	isAuthSource_Type()
}

type AuthSource_Dummy struct {
	Dummy *DummyAccess `protobuf:"bytes,2,opt,name=dummy,oneof"`
}
type AuthSource_Kube struct {
	Kube *KubernetesAccess `protobuf:"bytes,3,opt,name=kube,oneof"`
}
type AuthSource_Oidc struct {
	Oidc *OIDCAccess `protobuf:"bytes,4,opt,name=oidc,oneof"`
}
type AuthSource_Token struct {
	Token *TokenAccess `protobuf:"bytes,5,opt,name=token,oneof"`
}

func (*AuthSource_Dummy) isAuthSource_Type() {}
func (*AuthSource_Kube) isAuthSource_Type()  {}
func (*AuthSource_Oidc) isAuthSource_Type()  {}
func (*AuthSource_Token) isAuthSource_Type() {}

func (m *AuthSource) GetType() isAuthSource_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *AuthSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuthSource) GetDummy() *DummyAccess {
	if x, ok := m.GetType().(*AuthSource_Dummy); ok {
		return x.Dummy
	}
	return nil
}

func (m *AuthSource) GetKube() *KubernetesAccess {
	if x, ok := m.GetType().(*AuthSource_Kube); ok {
		return x.Kube
	}
	return nil
}

func (m *AuthSource) GetOidc() *OIDCAccess {
	if x, ok := m.GetType().(*AuthSource_Oidc); ok {
		return x.Oidc
	}
	return nil
}

func (m *AuthSource) GetToken() *TokenAccess {
	if x, ok := m.GetType().(*AuthSource_Token); ok {
		return x.Token
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AuthSource) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AuthSource_OneofMarshaler, _AuthSource_OneofUnmarshaler, _AuthSource_OneofSizer, []interface{}{
		(*AuthSource_Dummy)(nil),
		(*AuthSource_Kube)(nil),
		(*AuthSource_Oidc)(nil),
		(*AuthSource_Token)(nil),
	}
}

func _AuthSource_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AuthSource)
	// type
	switch x := m.Type.(type) {
	case *AuthSource_Dummy:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Dummy); err != nil {
			return err
		}
	case *AuthSource_Kube:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Kube); err != nil {
			return err
		}
	case *AuthSource_Oidc:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Oidc); err != nil {
			return err
		}
	case *AuthSource_Token:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Token); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AuthSource.Type has unexpected type %T", x)
	}
	return nil
}

func _AuthSource_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AuthSource)
	switch tag {
	case 2: // type.dummy
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DummyAccess)
		err := b.DecodeMessage(msg)
		m.Type = &AuthSource_Dummy{msg}
		return true, err
	case 3: // type.kube
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KubernetesAccess)
		err := b.DecodeMessage(msg)
		m.Type = &AuthSource_Kube{msg}
		return true, err
	case 4: // type.oidc
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(OIDCAccess)
		err := b.DecodeMessage(msg)
		m.Type = &AuthSource_Oidc{msg}
		return true, err
	case 5: // type.token
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TokenAccess)
		err := b.DecodeMessage(msg)
		m.Type = &AuthSource_Token{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AuthSource_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AuthSource)
	// type
	switch x := m.Type.(type) {
	case *AuthSource_Dummy:
		s := proto.Size(x.Dummy)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AuthSource_Kube:
		s := proto.Size(x.Kube)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AuthSource_Oidc:
		s := proto.Size(x.Oidc)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AuthSource_Token:
		s := proto.Size(x.Token)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// / KubernetesAccess is an convenient way of specifying auth for backend. It grabs the data inside already used
// / ~/.kube/config (or any specified config path) and deducts the auth type based on that. NOTE that only these types are
// / supported:
// / - OIDC
type KubernetesAccess struct {
	// User to reference access credentials from.
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	// By default ~/.kube/config as usual.
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *KubernetesAccess) Reset()                    { *m = KubernetesAccess{} }
func (m *KubernetesAccess) String() string            { return proto.CompactTextString(m) }
func (*KubernetesAccess) ProtoMessage()               {}
func (*KubernetesAccess) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *KubernetesAccess) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *KubernetesAccess) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type OIDCAccess struct {
	Provider string   `protobuf:"bytes,1,opt,name=provider" json:"provider,omitempty"`
	ClientId string   `protobuf:"bytes,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	Secret   string   `protobuf:"bytes,3,opt,name=secret" json:"secret,omitempty"`
	Scopes   []string `protobuf:"bytes,4,rep,name=scopes" json:"scopes,omitempty"`
	Path     string   `protobuf:"bytes,5,opt,name=path" json:"path,omitempty"`
	// login_callback_path specified URL path for redirect URL to specify when doing OIDC login.
	// If empty login will be disabled which means in case of no refresh token or not valid one, error will be returned
	// thus not needing user interaction.
	LoginCallbackPath string `protobuf:"bytes,6,opt,name=login_callback_path,json=loginCallbackPath" json:"login_callback_path,omitempty"`
}

func (m *OIDCAccess) Reset()                    { *m = OIDCAccess{} }
func (m *OIDCAccess) String() string            { return proto.CompactTextString(m) }
func (*OIDCAccess) ProtoMessage()               {}
func (*OIDCAccess) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OIDCAccess) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *OIDCAccess) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *OIDCAccess) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *OIDCAccess) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *OIDCAccess) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *OIDCAccess) GetLoginCallbackPath() string {
	if m != nil {
		return m.LoginCallbackPath
	}
	return ""
}

// DummyAccess just directly passes specified value into auth header. If value is not specified it will return error.
type DummyAccess struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *DummyAccess) Reset()                    { *m = DummyAccess{} }
func (m *DummyAccess) String() string            { return proto.CompactTextString(m) }
func (*DummyAccess) ProtoMessage()               {}
func (*DummyAccess) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DummyAccess) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// TokenAccess passes specified token into auth header as a bearer.
type TokenAccess struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *TokenAccess) Reset()                    { *m = TokenAccess{} }
func (m *TokenAccess) String() string            { return proto.CompactTextString(m) }
func (*TokenAccess) ProtoMessage()               {}
func (*TokenAccess) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TokenAccess) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthConfig)(nil), "winch.config.AuthConfig")
	proto.RegisterType((*AuthSource)(nil), "winch.config.AuthSource")
	proto.RegisterType((*KubernetesAccess)(nil), "winch.config.KubernetesAccess")
	proto.RegisterType((*OIDCAccess)(nil), "winch.config.OIDCAccess")
	proto.RegisterType((*DummyAccess)(nil), "winch.config.DummyAccess")
	proto.RegisterType((*TokenAccess)(nil), "winch.config.TokenAccess")
}

func init() { proto.RegisterFile("winch/config/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0xc9, 0x36, 0x8d, 0xe8, 0x64, 0x0f, 0x60, 0x10, 0x84, 0x15, 0x5a, 0xa2, 0xec, 0xa5,
	0x12, 0xda, 0x44, 0x2c, 0x88, 0x0b, 0xa7, 0x6d, 0xf7, 0x40, 0xc5, 0x01, 0x14, 0xb8, 0x57, 0x8e,
	0x63, 0x1a, 0x2b, 0x7f, 0x1c, 0xc5, 0x76, 0xab, 0x7d, 0x4c, 0x9e, 0x00, 0x89, 0x67, 0xe0, 0x01,
	0x50, 0xc6, 0xdd, 0xa4, 0x5d, 0xf6, 0x36, 0x9e, 0xef, 0x37, 0xe3, 0x6f, 0xc6, 0x86, 0x97, 0x3b,
	0xd1, 0xb0, 0x22, 0x61, 0xb2, 0xf9, 0x29, 0x36, 0x09, 0x35, 0xba, 0x88, 0xdb, 0x4e, 0x6a, 0x49,
	0x4e, 0x51, 0x88, 0xad, 0x70, 0xf6, 0x71, 0x23, 0x74, 0x61, 0xb2, 0x98, 0xc9, 0x3a, 0xa9, 0x77,
	0x42, 0x97, 0x72, 0x97, 0x6c, 0xe4, 0x25, 0xa2, 0x97, 0x5b, 0x5a, 0x89, 0x9c, 0x6a, 0xd9, 0xa9,
	0x64, 0x08, 0x6d, 0x97, 0x68, 0x05, 0x70, 0x6d, 0x74, 0xb1, 0xc4, 0x2e, 0xe4, 0x13, 0x9c, 0xf6,
	0x37, 0xac, 0x95, 0x34, 0x1d, 0xe3, 0x2a, 0x70, 0xc2, 0xc9, 0xdc, 0xbf, 0x0a, 0xe2, 0xc3, 0xab,
	0xe2, 0x9e, 0xff, 0x8e, 0x40, 0xea, 0xd3, 0x21, 0x56, 0xd1, 0x5f, 0xc7, 0xf6, 0xb2, 0x67, 0x42,
	0xc0, 0x6d, 0x68, 0xcd, 0x03, 0x27, 0x74, 0xe6, 0xb3, 0x14, 0x63, 0xf2, 0x0e, 0xa6, 0xb9, 0xa9,
	0xeb, 0xdb, 0xe0, 0x24, 0x74, 0xe6, 0xfe, 0xd5, 0xab, 0xe3, 0xc6, 0x37, 0xbd, 0x74, 0xcd, 0x18,
	0x57, 0xea, 0xf3, 0xa3, 0xd4, 0x92, 0xe4, 0x03, 0xb8, 0xa5, 0xc9, 0x78, 0x30, 0xc1, 0x8a, 0xf3,
	0xe3, 0x8a, 0x2f, 0x26, 0xe3, 0x5d, 0xc3, 0x35, 0x57, 0x43, 0x19, 0xd2, 0x24, 0x06, 0x57, 0x8a,
	0x9c, 0x05, 0x2e, 0x56, 0xdd, 0x1b, 0xe0, 0xeb, 0xea, 0x66, 0x39, 0xf2, 0x3d, 0xd7, 0x1b, 0xd3,
	0xb2, 0xe4, 0x4d, 0x30, 0x7d, 0xc8, 0xd8, 0x8f, 0x5e, 0x1a, 0x8d, 0x21, 0xb9, 0xf0, 0xc0, 0xd5,
	0xb7, 0x2d, 0x8f, 0x16, 0xf0, 0xe4, 0xbe, 0x0d, 0x72, 0x06, 0xae, 0x51, 0xbc, 0xb3, 0xb3, 0x2f,
	0xbc, 0x3f, 0xbf, 0xdf, 0x9c, 0x84, 0x4e, 0x8a, 0xb9, 0x7e, 0x2f, 0x2d, 0xd5, 0x05, 0xae, 0x60,
	0x96, 0x62, 0x1c, 0xfd, 0x72, 0x00, 0x46, 0x57, 0x24, 0x82, 0xc7, 0x6d, 0x27, 0xb7, 0x22, 0xff,
	0xaf, 0xc5, 0x90, 0x27, 0x17, 0x30, 0x63, 0x95, 0xe0, 0x8d, 0x5e, 0x8b, 0xdc, 0xf6, 0x1a, 0x21,
	0x2b, 0xac, 0x72, 0x72, 0x0e, 0x9e, 0xe2, 0xac, 0xe3, 0x1a, 0xd7, 0x37, 0x12, 0xfb, 0x2c, 0x79,
	0x01, 0x9e, 0x62, 0xb2, 0xe5, 0x2a, 0x70, 0xc3, 0xc9, 0x7c, 0x96, 0xee, 0x4f, 0x83, 0xc7, 0xe9,
	0xe8, 0x91, 0xc4, 0xf0, 0xac, 0x92, 0x1b, 0xd1, 0xac, 0x19, 0xad, 0xaa, 0x8c, 0xb2, 0x72, 0x8d,
	0x88, 0x87, 0xc8, 0x53, 0x94, 0x96, 0x7b, 0xe5, 0x5b, 0x3f, 0xd3, 0x05, 0xf8, 0x07, 0x0f, 0x4a,
	0x9e, 0xc3, 0x74, 0x4b, 0x2b, 0x73, 0xf7, 0x1f, 0xec, 0x21, 0x7a, 0x0b, 0xfe, 0xc1, 0x72, 0xc9,
	0xeb, 0xbb, 0x67, 0x38, 0x9e, 0xda, 0x26, 0x33, 0x0f, 0xbf, 0xec, 0xfb, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x12, 0xb9, 0x93, 0x6e, 0x13, 0x03, 0x00, 0x00,
}
