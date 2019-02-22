// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/supergloo/api/v1/install.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Install struct {
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by supergloo during validation
	Status core.Status `protobuf:"bytes,100,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,101,opt,name=metadata,proto3" json:"metadata"`
	// disables this install
	// setting this to true will cause supergloo to
	// not to install this mesh, or uninstall an active install
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Types that are valid to be assigned to InstallType:
	//	*Install_Istio_
	//	*Install_Custom_
	InstallType isInstall_InstallType `protobuf_oneof:"install_type"`
	// gzipped inline string containing the applied manifest
	// read-only, set by the server after successful installation.
	// TODO (ilackarms): make sure this is not too large for etcd (value size limit 1.5mb)
	InstalledManifest string `protobuf:"bytes,5,opt,name=installed_manifest,json=installedManifest,proto3" json:"installed_manifest,omitempty"`
	// reference to the Mesh crd that was created from this install
	// read-only, set by the server after successful installation.
	InstalledMesh        *core.ResourceRef `protobuf:"bytes,6,opt,name=installed_mesh,json=installedMesh,proto3" json:"installed_mesh,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Install) Reset()         { *m = Install{} }
func (m *Install) String() string { return proto.CompactTextString(m) }
func (*Install) ProtoMessage()    {}
func (*Install) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b058d98c63047dc, []int{0}
}
func (m *Install) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Install.Unmarshal(m, b)
}
func (m *Install) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Install.Marshal(b, m, deterministic)
}
func (m *Install) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Install.Merge(m, src)
}
func (m *Install) XXX_Size() int {
	return xxx_messageInfo_Install.Size(m)
}
func (m *Install) XXX_DiscardUnknown() {
	xxx_messageInfo_Install.DiscardUnknown(m)
}

var xxx_messageInfo_Install proto.InternalMessageInfo

type isInstall_InstallType interface {
	isInstall_InstallType()
	Equal(interface{}) bool
}

type Install_Istio_ struct {
	Istio *Install_Istio `protobuf:"bytes,2,opt,name=istio,proto3,oneof"`
}
type Install_Custom_ struct {
	Custom *Install_Custom `protobuf:"bytes,3,opt,name=custom,proto3,oneof"`
}

func (*Install_Istio_) isInstall_InstallType()  {}
func (*Install_Custom_) isInstall_InstallType() {}

func (m *Install) GetInstallType() isInstall_InstallType {
	if m != nil {
		return m.InstallType
	}
	return nil
}

func (m *Install) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Install) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Install) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *Install) GetIstio() *Install_Istio {
	if x, ok := m.GetInstallType().(*Install_Istio_); ok {
		return x.Istio
	}
	return nil
}

func (m *Install) GetCustom() *Install_Custom {
	if x, ok := m.GetInstallType().(*Install_Custom_); ok {
		return x.Custom
	}
	return nil
}

func (m *Install) GetInstalledManifest() string {
	if m != nil {
		return m.InstalledManifest
	}
	return ""
}

func (m *Install) GetInstalledMesh() *core.ResourceRef {
	if m != nil {
		return m.InstalledMesh
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Install) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Install_OneofMarshaler, _Install_OneofUnmarshaler, _Install_OneofSizer, []interface{}{
		(*Install_Istio_)(nil),
		(*Install_Custom_)(nil),
	}
}

func _Install_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Install)
	// install_type
	switch x := m.InstallType.(type) {
	case *Install_Istio_:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Istio); err != nil {
			return err
		}
	case *Install_Custom_:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Custom); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Install.InstallType has unexpected type %T", x)
	}
	return nil
}

func _Install_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Install)
	switch tag {
	case 2: // install_type.istio
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Install_Istio)
		err := b.DecodeMessage(msg)
		m.InstallType = &Install_Istio_{msg}
		return true, err
	case 3: // install_type.custom
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Install_Custom)
		err := b.DecodeMessage(msg)
		m.InstallType = &Install_Custom_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Install_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Install)
	// install_type
	switch x := m.InstallType.(type) {
	case *Install_Istio_:
		s := proto.Size(x.Istio)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Install_Custom_:
		s := proto.Size(x.Custom)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Install_Istio struct {
	// which namespace to install to
	InstallationNamespace string `protobuf:"bytes,1,opt,name=installation_namespace,json=installationNamespace,proto3" json:"installation_namespace,omitempty"`
	// which version of the istio helm chart to install
	// ignored if using custom helm chart
	IstioVersion string `protobuf:"bytes,2,opt,name=istio_version,json=istioVersion,proto3" json:"istio_version,omitempty"`
	// enable auto injection of pods
	EnableAutoInject bool `protobuf:"varint,3,opt,name=enable_auto_inject,json=enableAutoInject,proto3" json:"enable_auto_inject,omitempty"`
	// enable mutual tls between pods
	EnableMtls bool `protobuf:"varint,4,opt,name=enable_mtls,json=enableMtls,proto3" json:"enable_mtls,omitempty"`
	// optional. set to use a custom root ca
	// to issue certificates for mtls
	// ignored if mtls is disabled
	CustomRootCert *core.ResourceRef `protobuf:"bytes,9,opt,name=custom_root_cert,json=customRootCert,proto3" json:"custom_root_cert,omitempty"`
	// install grafana with istio
	InstallGrafana bool `protobuf:"varint,6,opt,name=install_grafana,json=installGrafana,proto3" json:"install_grafana,omitempty"`
	// install prometheus with istio
	InstallPrometheus bool `protobuf:"varint,7,opt,name=install_prometheus,json=installPrometheus,proto3" json:"install_prometheus,omitempty"`
	// install jaeger with istio
	InstallJaeger        bool     `protobuf:"varint,8,opt,name=install_jaeger,json=installJaeger,proto3" json:"install_jaeger,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Install_Istio) Reset()         { *m = Install_Istio{} }
func (m *Install_Istio) String() string { return proto.CompactTextString(m) }
func (*Install_Istio) ProtoMessage()    {}
func (*Install_Istio) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b058d98c63047dc, []int{0, 0}
}
func (m *Install_Istio) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Install_Istio.Unmarshal(m, b)
}
func (m *Install_Istio) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Install_Istio.Marshal(b, m, deterministic)
}
func (m *Install_Istio) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Install_Istio.Merge(m, src)
}
func (m *Install_Istio) XXX_Size() int {
	return xxx_messageInfo_Install_Istio.Size(m)
}
func (m *Install_Istio) XXX_DiscardUnknown() {
	xxx_messageInfo_Install_Istio.DiscardUnknown(m)
}

var xxx_messageInfo_Install_Istio proto.InternalMessageInfo

func (m *Install_Istio) GetInstallationNamespace() string {
	if m != nil {
		return m.InstallationNamespace
	}
	return ""
}

func (m *Install_Istio) GetIstioVersion() string {
	if m != nil {
		return m.IstioVersion
	}
	return ""
}

func (m *Install_Istio) GetEnableAutoInject() bool {
	if m != nil {
		return m.EnableAutoInject
	}
	return false
}

func (m *Install_Istio) GetEnableMtls() bool {
	if m != nil {
		return m.EnableMtls
	}
	return false
}

func (m *Install_Istio) GetCustomRootCert() *core.ResourceRef {
	if m != nil {
		return m.CustomRootCert
	}
	return nil
}

func (m *Install_Istio) GetInstallGrafana() bool {
	if m != nil {
		return m.InstallGrafana
	}
	return false
}

func (m *Install_Istio) GetInstallPrometheus() bool {
	if m != nil {
		return m.InstallPrometheus
	}
	return false
}

func (m *Install_Istio) GetInstallJaeger() bool {
	if m != nil {
		return m.InstallJaeger
	}
	return false
}

// note: currently unimplemented
type Install_Custom struct {
	// supergloo will attempt to
	// install this helm chart.
	// format should be a URI accessible to supergloo
	CustomHelmChart string `protobuf:"bytes,9,opt,name=custom_helm_chart,json=customHelmChart,proto3" json:"custom_helm_chart,omitempty"`
	// optional. if specified,
	// these values will be used
	// when rendering the helm chart
	CustomHelmValues     string   `protobuf:"bytes,10,opt,name=custom_helm_values,json=customHelmValues,proto3" json:"custom_helm_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Install_Custom) Reset()         { *m = Install_Custom{} }
func (m *Install_Custom) String() string { return proto.CompactTextString(m) }
func (*Install_Custom) ProtoMessage()    {}
func (*Install_Custom) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b058d98c63047dc, []int{0, 1}
}
func (m *Install_Custom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Install_Custom.Unmarshal(m, b)
}
func (m *Install_Custom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Install_Custom.Marshal(b, m, deterministic)
}
func (m *Install_Custom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Install_Custom.Merge(m, src)
}
func (m *Install_Custom) XXX_Size() int {
	return xxx_messageInfo_Install_Custom.Size(m)
}
func (m *Install_Custom) XXX_DiscardUnknown() {
	xxx_messageInfo_Install_Custom.DiscardUnknown(m)
}

var xxx_messageInfo_Install_Custom proto.InternalMessageInfo

func (m *Install_Custom) GetCustomHelmChart() string {
	if m != nil {
		return m.CustomHelmChart
	}
	return ""
}

func (m *Install_Custom) GetCustomHelmValues() string {
	if m != nil {
		return m.CustomHelmValues
	}
	return ""
}

func init() {
	proto.RegisterType((*Install)(nil), "supergloo.solo.io.Install")
	proto.RegisterType((*Install_Istio)(nil), "supergloo.solo.io.Install.Istio")
	proto.RegisterType((*Install_Custom)(nil), "supergloo.solo.io.Install.Custom")
}

func init() {
	proto.RegisterFile("github.com/solo-io/supergloo/api/v1/install.proto", fileDescriptor_7b058d98c63047dc)
}

var fileDescriptor_7b058d98c63047dc = []byte{
	// 654 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x4e, 0xdb, 0x4c,
	0x14, 0x25, 0x10, 0x42, 0x32, 0xfc, 0x66, 0x04, 0xc8, 0x64, 0x01, 0xf9, 0xf8, 0x54, 0x81, 0x2a,
	0xb0, 0x45, 0xab, 0x4a, 0x88, 0x6e, 0xda, 0x64, 0x01, 0x54, 0xa2, 0xaa, 0xa6, 0x12, 0x8b, 0x6e,
	0xac, 0x89, 0x73, 0xed, 0x0c, 0xd8, 0xbe, 0xd6, 0xcc, 0x98, 0xaa, 0x5b, 0x9e, 0xa1, 0x0f, 0xd1,
	0x47, 0xe9, 0x53, 0xb0, 0xe8, 0x1b, 0xd0, 0x45, 0xd7, 0x95, 0xc7, 0xe3, 0x10, 0xd4, 0x96, 0xd2,
	0x55, 0x32, 0xf7, 0x9c, 0x73, 0xe7, 0x9e, 0x73, 0x6d, 0x93, 0x83, 0x48, 0xe8, 0x51, 0x3e, 0x70,
	0x03, 0x4c, 0x3c, 0x85, 0x31, 0xee, 0x0b, 0xf4, 0x54, 0x9e, 0x81, 0x8c, 0x62, 0x44, 0x8f, 0x67,
	0xc2, 0xbb, 0x3a, 0xf0, 0x44, 0xaa, 0x34, 0x8f, 0x63, 0x37, 0x93, 0xa8, 0x91, 0xb6, 0xc7, 0xb8,
	0x5b, 0x28, 0x5c, 0x81, 0x9d, 0xd5, 0x08, 0x23, 0x34, 0xa8, 0x57, 0xfc, 0x2b, 0x89, 0x9d, 0xcd,
	0x08, 0x31, 0x8a, 0xc1, 0x33, 0xa7, 0x41, 0x1e, 0x7a, 0xc3, 0x5c, 0x72, 0x2d, 0x30, 0xfd, 0x13,
	0xfe, 0x51, 0xf2, 0x2c, 0x03, 0xa9, 0x2c, 0xfe, 0xdb, 0xd9, 0x8a, 0xdf, 0x4b, 0xa1, 0xab, 0xd1,
	0x12, 0xd0, 0x7c, 0xc8, 0x35, 0xb7, 0x12, 0xef, 0x11, 0x12, 0xa5, 0xb9, 0xce, 0xab, 0x3b, 0xf6,
	0x1e, 0x21, 0x90, 0x10, 0xfe, 0xc3, 0x44, 0xd5, 0xb9, 0x94, 0x6c, 0x7f, 0x9e, 0x23, 0x73, 0xa7,
	0x65, 0x7e, 0xf4, 0x98, 0x34, 0xca, 0xcb, 0x9d, 0x61, 0xb7, 0xb6, 0x3b, 0xff, 0x6c, 0xd5, 0x0d,
	0x50, 0x42, 0x95, 0xa2, 0xfb, 0xde, 0x60, 0xbd, 0x8d, 0xaf, 0x37, 0x5b, 0x53, 0xdf, 0x6f, 0xb6,
	0xda, 0x1a, 0x94, 0x1e, 0x8a, 0x30, 0x3c, 0xda, 0x16, 0x51, 0x8a, 0x12, 0xb6, 0x99, 0x95, 0xd3,
	0x43, 0xd2, 0xac, 0x8c, 0x3b, 0x60, 0x5a, 0xad, 0xdf, 0x6f, 0x75, 0x66, 0xd1, 0x5e, 0xbd, 0x68,
	0xc6, 0xc6, 0x6c, 0xda, 0x21, 0xcd, 0xa1, 0x50, 0x7c, 0x10, 0xc3, 0xd0, 0xa9, 0x75, 0x6b, 0xbb,
	0x4d, 0x36, 0x3e, 0xd3, 0x43, 0x32, 0x2b, 0x94, 0x16, 0xe8, 0x4c, 0x9b, 0x96, 0x5d, 0xf7, 0x97,
	0x45, 0xbb, 0xd6, 0x89, 0x7b, 0x5a, 0xf0, 0x4e, 0xa6, 0x58, 0x29, 0xa0, 0x2f, 0x49, 0x23, 0xc8,
	0x95, 0xc6, 0xc4, 0x99, 0x31, 0xd2, 0xff, 0x1e, 0x90, 0xf6, 0x0d, 0xf1, 0x64, 0x8a, 0x59, 0x09,
	0xdd, 0x27, 0xd4, 0x3e, 0x60, 0x30, 0xf4, 0x13, 0x9e, 0x8a, 0x10, 0x94, 0x76, 0x66, 0xbb, 0xb5,
	0xdd, 0x16, 0x6b, 0x8f, 0x91, 0x33, 0x0b, 0xd0, 0x57, 0x64, 0x69, 0x82, 0x0e, 0x6a, 0xe4, 0x34,
	0xcc, 0x9d, 0x1b, 0xf7, 0x13, 0x60, 0xa0, 0x30, 0x97, 0x01, 0x30, 0x08, 0xd9, 0xe2, 0x5d, 0x17,
	0x50, 0xa3, 0xce, 0x8f, 0x69, 0x32, 0x6b, 0x0c, 0xd0, 0x17, 0x64, 0xdd, 0x42, 0xe6, 0xb9, 0xf4,
	0x53, 0x9e, 0x80, 0xca, 0x78, 0x00, 0x26, 0x9b, 0x16, 0x5b, 0x9b, 0x44, 0xdf, 0x56, 0x20, 0xfd,
	0x9f, 0x2c, 0x1a, 0xdf, 0xfe, 0x15, 0x48, 0x25, 0x30, 0x35, 0x81, 0xb5, 0xd8, 0x82, 0x29, 0x9e,
	0x97, 0x35, 0xba, 0x47, 0x28, 0xa4, 0x45, 0xb0, 0x3e, 0xcf, 0x35, 0xfa, 0x22, 0xbd, 0x80, 0x40,
	0x9b, 0x7c, 0x9a, 0x6c, 0xa5, 0x44, 0x5e, 0xe7, 0x1a, 0x4f, 0x4d, 0x9d, 0x6e, 0x91, 0x79, 0xcb,
	0x4e, 0x74, 0xac, 0x9c, 0xba, 0xa1, 0x91, 0xb2, 0x74, 0xa6, 0x63, 0x45, 0xfb, 0x64, 0xa5, 0xcc,
	0xcb, 0x97, 0x88, 0xda, 0x0f, 0x40, 0x6a, 0xa7, 0xf5, 0x37, 0xe3, 0x4b, 0xa5, 0x84, 0x21, 0xea,
	0x3e, 0x48, 0x4d, 0x77, 0xc8, 0xb2, 0x75, 0xe4, 0x47, 0x92, 0x87, 0x3c, 0xe5, 0x26, 0xbc, 0x26,
	0xab, 0x22, 0x3d, 0x2e, 0xab, 0x13, 0x3b, 0xf1, 0x33, 0x89, 0x09, 0xe8, 0x11, 0xe4, 0xca, 0x99,
	0x33, 0xdc, 0x6a, 0x27, 0xef, 0xc6, 0x00, 0x7d, 0x32, 0xde, 0x89, 0x7f, 0xc1, 0x21, 0x02, 0xe9,
	0x34, 0x0d, 0xb5, 0x0a, 0xfe, 0x8d, 0x29, 0x76, 0x06, 0xa4, 0x51, 0x6e, 0x9f, 0x3e, 0x25, 0x6d,
	0xeb, 0x66, 0x04, 0x71, 0xe2, 0x07, 0x23, 0x6e, 0xed, 0xb4, 0xd8, 0x72, 0x09, 0x9c, 0x40, 0x9c,
	0xf4, 0x8b, 0x72, 0x11, 0xe4, 0x24, 0xf7, 0x8a, 0xc7, 0x39, 0x28, 0x87, 0x18, 0xf2, 0xca, 0x1d,
	0xf9, 0xdc, 0xd4, 0x8f, 0xd6, 0xae, 0x6f, 0xeb, 0x33, 0xa4, 0x26, 0xae, 0x6f, 0xeb, 0x84, 0x36,
	0xed, 0xfd, 0xaa, 0xb7, 0x44, 0x16, 0xaa, 0x09, 0xf5, 0xa7, 0x0c, 0x7a, 0xfb, 0x5f, 0xbe, 0x6d,
	0xd6, 0x3e, 0xec, 0x3c, 0xf8, 0xf5, 0xcb, 0x2e, 0x23, 0xfb, 0x52, 0x0f, 0x1a, 0xe6, 0x65, 0x7e,
	0xfe, 0x33, 0x00, 0x00, 0xff, 0xff, 0xae, 0xe1, 0xc6, 0xde, 0x2f, 0x05, 0x00, 0x00,
}

func (this *Install) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Install)
	if !ok {
		that2, ok := that.(Install)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if this.Disabled != that1.Disabled {
		return false
	}
	if that1.InstallType == nil {
		if this.InstallType != nil {
			return false
		}
	} else if this.InstallType == nil {
		return false
	} else if !this.InstallType.Equal(that1.InstallType) {
		return false
	}
	if this.InstalledManifest != that1.InstalledManifest {
		return false
	}
	if !this.InstalledMesh.Equal(that1.InstalledMesh) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Install_Istio_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Install_Istio_)
	if !ok {
		that2, ok := that.(Install_Istio_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Istio.Equal(that1.Istio) {
		return false
	}
	return true
}
func (this *Install_Custom_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Install_Custom_)
	if !ok {
		that2, ok := that.(Install_Custom_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Custom.Equal(that1.Custom) {
		return false
	}
	return true
}
func (this *Install_Istio) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Install_Istio)
	if !ok {
		that2, ok := that.(Install_Istio)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.InstallationNamespace != that1.InstallationNamespace {
		return false
	}
	if this.IstioVersion != that1.IstioVersion {
		return false
	}
	if this.EnableAutoInject != that1.EnableAutoInject {
		return false
	}
	if this.EnableMtls != that1.EnableMtls {
		return false
	}
	if !this.CustomRootCert.Equal(that1.CustomRootCert) {
		return false
	}
	if this.InstallGrafana != that1.InstallGrafana {
		return false
	}
	if this.InstallPrometheus != that1.InstallPrometheus {
		return false
	}
	if this.InstallJaeger != that1.InstallJaeger {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Install_Custom) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Install_Custom)
	if !ok {
		that2, ok := that.(Install_Custom)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CustomHelmChart != that1.CustomHelmChart {
		return false
	}
	if this.CustomHelmValues != that1.CustomHelmValues {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
