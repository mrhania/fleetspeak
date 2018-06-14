// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fleetspeak/src/server/plugins/proto/plugins/config.proto

/*
Package fleetspeak_plugins is a generated protocol buffer package.

It is generated from these files:
	fleetspeak/src/server/plugins/proto/plugins/config.proto

It has these top-level messages:
	Plugin
	Config
*/
package fleetspeak_plugins

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A Plugin message describes how to instantiate a FS component from a go plugin
// file, as defined by go's plugin library: https://golang.org/pkg/plugin/
type Plugin struct {
	// The full path to the plugin file containing this component.
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// The name of a factory function exported by the file. It will be passed the
	// config string and must be one of the *Factory types defined in the
	// Fleetspeak server.plugins package.
	FactoryName string `protobuf:"bytes,2,opt,name=factory_name,json=factoryName" json:"factory_name,omitempty"`
	Config      string `protobuf:"bytes,3,opt,name=config" json:"config,omitempty"`
}

func (m *Plugin) Reset()                    { *m = Plugin{} }
func (m *Plugin) String() string            { return proto.CompactTextString(m) }
func (*Plugin) ProtoMessage()               {}
func (*Plugin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Plugin) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Plugin) GetFactoryName() string {
	if m != nil {
		return m.FactoryName
	}
	return ""
}

func (m *Plugin) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

type Config struct {
	Datastore      *Plugin   `protobuf:"bytes,1,opt,name=datastore" json:"datastore,omitempty"`
	ServiceFactory []*Plugin `protobuf:"bytes,2,rep,name=service_factory,json=serviceFactory" json:"service_factory,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Config) GetDatastore() *Plugin {
	if m != nil {
		return m.Datastore
	}
	return nil
}

func (m *Config) GetServiceFactory() []*Plugin {
	if m != nil {
		return m.ServiceFactory
	}
	return nil
}

func init() {
	proto.RegisterType((*Plugin)(nil), "fleetspeak.plugins.Plugin")
	proto.RegisterType((*Config)(nil), "fleetspeak.plugins.Config")
}

func init() {
	proto.RegisterFile("fleetspeak/src/server/plugins/proto/plugins/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcb, 0x49, 0x4d,
	0x2d, 0x29, 0x2e, 0x48, 0x4d, 0xcc, 0xd6, 0x2f, 0x2e, 0x4a, 0xd6, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b,
	0x2d, 0xd2, 0x2f, 0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0x2b, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x87,
	0xf3, 0x92, 0xf3, 0xf3, 0xd2, 0x32, 0xd3, 0xf5, 0xc0, 0x82, 0x42, 0x42, 0x08, 0x9d, 0x7a, 0x50,
	0x05, 0x4a, 0xe1, 0x5c, 0x6c, 0x01, 0x60, 0xa6, 0x90, 0x10, 0x17, 0x4b, 0x41, 0x62, 0x49, 0x86,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x2d, 0xa4, 0xc8, 0xc5, 0x93, 0x96, 0x98, 0x5c,
	0x92, 0x5f, 0x54, 0x19, 0x9f, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0x04, 0x96, 0xe3, 0x86, 0x8a, 0xf9,
	0x25, 0xe6, 0xa6, 0x0a, 0x89, 0x71, 0xb1, 0x41, 0x2c, 0x91, 0x60, 0x06, 0x4b, 0x42, 0x79, 0x4a,
	0xed, 0x8c, 0x5c, 0x6c, 0xce, 0x60, 0xa6, 0x90, 0x05, 0x17, 0x67, 0x4a, 0x62, 0x49, 0x62, 0x71,
	0x49, 0x7e, 0x51, 0x2a, 0xd8, 0x78, 0x6e, 0x23, 0x29, 0x3d, 0x4c, 0xb7, 0xe8, 0x41, 0x1c, 0x12,
	0x84, 0x50, 0x2c, 0xe4, 0xcc, 0xc5, 0x0f, 0xf2, 0x5e, 0x66, 0x72, 0x6a, 0x3c, 0xd4, 0x4e, 0x09,
	0x26, 0x05, 0x66, 0x02, 0xfa, 0xf9, 0xa0, 0x5a, 0xdc, 0x20, 0x3a, 0x92, 0xd8, 0xc0, 0xbe, 0x37,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x18, 0xa5, 0xf7, 0x39, 0x01, 0x00, 0x00,
}
