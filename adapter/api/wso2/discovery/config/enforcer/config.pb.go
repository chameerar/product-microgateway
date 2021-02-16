// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.9.1
// source: wso2/discovery/config/enforcer/config.proto

package enforcer

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Enforcer config model
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JwtTokenConfig   []*Issuer      `protobuf:"bytes,1,rep,name=jwtTokenConfig,proto3" json:"jwtTokenConfig,omitempty"`
	Keystore         *CertStore     `protobuf:"bytes,2,opt,name=keystore,proto3" json:"keystore,omitempty"`
	Truststore       *CertStore     `protobuf:"bytes,3,opt,name=truststore,proto3" json:"truststore,omitempty"`
	Eventhub         *EventHub      `protobuf:"bytes,4,opt,name=eventhub,proto3" json:"eventhub,omitempty"`
	AuthService      *AuthService   `protobuf:"bytes,5,opt,name=authService,proto3" json:"authService,omitempty"`
	ApimCredentials  *AmCredentials `protobuf:"bytes,6,opt,name=apimCredentials,proto3" json:"apimCredentials,omitempty"`
	JwtGenerator     *JWTGenerator  `protobuf:"bytes,7,opt,name=jwtGenerator,proto3" json:"jwtGenerator,omitempty"`
	ThrottlingConfig *Throttling    `protobuf:"bytes,8,opt,name=throttlingConfig,proto3" json:"throttlingConfig,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_config_enforcer_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_config_enforcer_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_config_enforcer_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetJwtTokenConfig() []*Issuer {
	if x != nil {
		return x.JwtTokenConfig
	}
	return nil
}

func (x *Config) GetKeystore() *CertStore {
	if x != nil {
		return x.Keystore
	}
	return nil
}

func (x *Config) GetTruststore() *CertStore {
	if x != nil {
		return x.Truststore
	}
	return nil
}

func (x *Config) GetEventhub() *EventHub {
	if x != nil {
		return x.Eventhub
	}
	return nil
}

func (x *Config) GetAuthService() *AuthService {
	if x != nil {
		return x.AuthService
	}
	return nil
}

func (x *Config) GetApimCredentials() *AmCredentials {
	if x != nil {
		return x.ApimCredentials
	}
	return nil
}

func (x *Config) GetJwtGenerator() *JWTGenerator {
	if x != nil {
		return x.JwtGenerator
	}
	return nil
}

func (x *Config) GetThrottlingConfig() *Throttling {
	if x != nil {
		return x.ThrottlingConfig
	}
	return nil
}

var File_wso2_discovery_config_enforcer_config_proto protoreflect.FileDescriptor

var file_wso2_discovery_config_enforcer_config_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x77,
	0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x1a, 0x29, 0x77,
	0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f, 0x63, 0x65,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x75, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f, 0x61, 0x6d, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x77, 0x73, 0x6f, 0x32,
	0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x77,
	0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2f, 0x6a, 0x77,
	0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x72, 0x2f, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x82, 0x05, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4e, 0x0a,
	0x0e, 0x6a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x52, 0x0e, 0x6a,
	0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x45, 0x0a,
	0x08, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72,
	0x2e, 0x43, 0x65, 0x72, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x52, 0x0a, 0x74, 0x72, 0x75, 0x73, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12,
	0x44, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x75, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x72, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x75, 0x62, 0x52, 0x08, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x68, 0x75, 0x62, 0x12, 0x4d, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x77, 0x73, 0x6f,
	0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0f, 0x61, 0x70, 0x69, 0x6d, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x41,
	0x6d, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x0f, 0x61, 0x70,
	0x69, 0x6d, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x50, 0x0a,
	0x0c, 0x6a, 0x77, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x72, 0x2e, 0x4a, 0x57, 0x54, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x0c, 0x6a, 0x77, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x56, 0x0a, 0x10, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x77, 0x73, 0x6f, 0x32,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x54, 0x68, 0x72, 0x6f, 0x74,
	0x74, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x8b, 0x01, 0x0a, 0x2a, 0x6f, 0x72, 0x67, 0x2e,
	0x77, 0x73, 0x6f, 0x32, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x42, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x77, 0x73,
	0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x3b, 0x65, 0x6e, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wso2_discovery_config_enforcer_config_proto_rawDescOnce sync.Once
	file_wso2_discovery_config_enforcer_config_proto_rawDescData = file_wso2_discovery_config_enforcer_config_proto_rawDesc
)

func file_wso2_discovery_config_enforcer_config_proto_rawDescGZIP() []byte {
	file_wso2_discovery_config_enforcer_config_proto_rawDescOnce.Do(func() {
		file_wso2_discovery_config_enforcer_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_wso2_discovery_config_enforcer_config_proto_rawDescData)
	})
	return file_wso2_discovery_config_enforcer_config_proto_rawDescData
}

var file_wso2_discovery_config_enforcer_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wso2_discovery_config_enforcer_config_proto_goTypes = []interface{}{
	(*Config)(nil),        // 0: wso2.discovery.config.enforcer.Config
	(*Issuer)(nil),        // 1: wso2.discovery.config.enforcer.Issuer
	(*CertStore)(nil),     // 2: wso2.discovery.config.enforcer.CertStore
	(*EventHub)(nil),      // 3: wso2.discovery.config.enforcer.EventHub
	(*AuthService)(nil),   // 4: wso2.discovery.config.enforcer.AuthService
	(*AmCredentials)(nil), // 5: wso2.discovery.config.enforcer.AmCredentials
	(*JWTGenerator)(nil),  // 6: wso2.discovery.config.enforcer.JWTGenerator
	(*Throttling)(nil),    // 7: wso2.discovery.config.enforcer.Throttling
}
var file_wso2_discovery_config_enforcer_config_proto_depIdxs = []int32{
	1, // 0: wso2.discovery.config.enforcer.Config.jwtTokenConfig:type_name -> wso2.discovery.config.enforcer.Issuer
	2, // 1: wso2.discovery.config.enforcer.Config.keystore:type_name -> wso2.discovery.config.enforcer.CertStore
	2, // 2: wso2.discovery.config.enforcer.Config.truststore:type_name -> wso2.discovery.config.enforcer.CertStore
	3, // 3: wso2.discovery.config.enforcer.Config.eventhub:type_name -> wso2.discovery.config.enforcer.EventHub
	4, // 4: wso2.discovery.config.enforcer.Config.authService:type_name -> wso2.discovery.config.enforcer.AuthService
	5, // 5: wso2.discovery.config.enforcer.Config.apimCredentials:type_name -> wso2.discovery.config.enforcer.AmCredentials
	6, // 6: wso2.discovery.config.enforcer.Config.jwtGenerator:type_name -> wso2.discovery.config.enforcer.JWTGenerator
	7, // 7: wso2.discovery.config.enforcer.Config.throttlingConfig:type_name -> wso2.discovery.config.enforcer.Throttling
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_wso2_discovery_config_enforcer_config_proto_init() }
func file_wso2_discovery_config_enforcer_config_proto_init() {
	if File_wso2_discovery_config_enforcer_config_proto != nil {
		return
	}
	file_wso2_discovery_config_enforcer_cert_proto_init()
	file_wso2_discovery_config_enforcer_issuer_proto_init()
	file_wso2_discovery_config_enforcer_event_hub_proto_init()
	file_wso2_discovery_config_enforcer_am_credentials_proto_init()
	file_wso2_discovery_config_enforcer_auth_service_proto_init()
	file_wso2_discovery_config_enforcer_jwt_generator_proto_init()
	file_wso2_discovery_config_enforcer_throttling_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wso2_discovery_config_enforcer_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_wso2_discovery_config_enforcer_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wso2_discovery_config_enforcer_config_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_config_enforcer_config_proto_depIdxs,
		MessageInfos:      file_wso2_discovery_config_enforcer_config_proto_msgTypes,
	}.Build()
	File_wso2_discovery_config_enforcer_config_proto = out.File
	file_wso2_discovery_config_enforcer_config_proto_rawDesc = nil
	file_wso2_discovery_config_enforcer_config_proto_goTypes = nil
	file_wso2_discovery_config_enforcer_config_proto_depIdxs = nil
}
