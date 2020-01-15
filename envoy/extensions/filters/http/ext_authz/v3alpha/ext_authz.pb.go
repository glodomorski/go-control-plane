// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/ext_authz/v3alpha/ext_authz.proto

package envoy_extensions_filters_http_ext_authz_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/config/core/v3alpha"
	v3alpha2 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3alpha"
	v3alpha1 "github.com/envoyproxy/go-control-plane/envoy/type/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ExtAuthz struct {
	// Types that are valid to be assigned to Services:
	//	*ExtAuthz_GrpcService
	//	*ExtAuthz_HttpService
	Services                      isExtAuthz_Services               `protobuf_oneof:"services"`
	FailureModeAllow              bool                              `protobuf:"varint,2,opt,name=failure_mode_allow,json=failureModeAllow,proto3" json:"failure_mode_allow,omitempty"`
	HiddenEnvoyDeprecatedUseAlpha bool                              `protobuf:"varint,4,opt,name=hidden_envoy_deprecated_use_alpha,json=hiddenEnvoyDeprecatedUseAlpha,proto3" json:"hidden_envoy_deprecated_use_alpha,omitempty"` // Deprecated: Do not use.
	WithRequestBody               *BufferSettings                   `protobuf:"bytes,5,opt,name=with_request_body,json=withRequestBody,proto3" json:"with_request_body,omitempty"`
	ClearRouteCache               bool                              `protobuf:"varint,6,opt,name=clear_route_cache,json=clearRouteCache,proto3" json:"clear_route_cache,omitempty"`
	StatusOnError                 *v3alpha1.HttpStatus              `protobuf:"bytes,7,opt,name=status_on_error,json=statusOnError,proto3" json:"status_on_error,omitempty"`
	MetadataContextNamespaces     []string                          `protobuf:"bytes,8,rep,name=metadata_context_namespaces,json=metadataContextNamespaces,proto3" json:"metadata_context_namespaces,omitempty"`
	FilterEnabled                 *v3alpha.RuntimeFractionalPercent `protobuf:"bytes,9,opt,name=filter_enabled,json=filterEnabled,proto3" json:"filter_enabled,omitempty"`
	IncludePeerCertificate        bool                              `protobuf:"varint,10,opt,name=include_peer_certificate,json=includePeerCertificate,proto3" json:"include_peer_certificate,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                          `json:"-"`
	XXX_unrecognized              []byte                            `json:"-"`
	XXX_sizecache                 int32                             `json:"-"`
}

func (m *ExtAuthz) Reset()         { *m = ExtAuthz{} }
func (m *ExtAuthz) String() string { return proto.CompactTextString(m) }
func (*ExtAuthz) ProtoMessage()    {}
func (*ExtAuthz) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20c629614df4358, []int{0}
}

func (m *ExtAuthz) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtAuthz.Unmarshal(m, b)
}
func (m *ExtAuthz) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtAuthz.Marshal(b, m, deterministic)
}
func (m *ExtAuthz) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthz.Merge(m, src)
}
func (m *ExtAuthz) XXX_Size() int {
	return xxx_messageInfo_ExtAuthz.Size(m)
}
func (m *ExtAuthz) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthz.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthz proto.InternalMessageInfo

type isExtAuthz_Services interface {
	isExtAuthz_Services()
}

type ExtAuthz_GrpcService struct {
	GrpcService *v3alpha.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3,oneof"`
}

type ExtAuthz_HttpService struct {
	HttpService *HttpService `protobuf:"bytes,3,opt,name=http_service,json=httpService,proto3,oneof"`
}

func (*ExtAuthz_GrpcService) isExtAuthz_Services() {}

func (*ExtAuthz_HttpService) isExtAuthz_Services() {}

func (m *ExtAuthz) GetServices() isExtAuthz_Services {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *ExtAuthz) GetGrpcService() *v3alpha.GrpcService {
	if x, ok := m.GetServices().(*ExtAuthz_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

func (m *ExtAuthz) GetHttpService() *HttpService {
	if x, ok := m.GetServices().(*ExtAuthz_HttpService); ok {
		return x.HttpService
	}
	return nil
}

func (m *ExtAuthz) GetFailureModeAllow() bool {
	if m != nil {
		return m.FailureModeAllow
	}
	return false
}

// Deprecated: Do not use.
func (m *ExtAuthz) GetHiddenEnvoyDeprecatedUseAlpha() bool {
	if m != nil {
		return m.HiddenEnvoyDeprecatedUseAlpha
	}
	return false
}

func (m *ExtAuthz) GetWithRequestBody() *BufferSettings {
	if m != nil {
		return m.WithRequestBody
	}
	return nil
}

func (m *ExtAuthz) GetClearRouteCache() bool {
	if m != nil {
		return m.ClearRouteCache
	}
	return false
}

func (m *ExtAuthz) GetStatusOnError() *v3alpha1.HttpStatus {
	if m != nil {
		return m.StatusOnError
	}
	return nil
}

func (m *ExtAuthz) GetMetadataContextNamespaces() []string {
	if m != nil {
		return m.MetadataContextNamespaces
	}
	return nil
}

func (m *ExtAuthz) GetFilterEnabled() *v3alpha.RuntimeFractionalPercent {
	if m != nil {
		return m.FilterEnabled
	}
	return nil
}

func (m *ExtAuthz) GetIncludePeerCertificate() bool {
	if m != nil {
		return m.IncludePeerCertificate
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExtAuthz) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExtAuthz_GrpcService)(nil),
		(*ExtAuthz_HttpService)(nil),
	}
}

type BufferSettings struct {
	MaxRequestBytes      uint32   `protobuf:"varint,1,opt,name=max_request_bytes,json=maxRequestBytes,proto3" json:"max_request_bytes,omitempty"`
	AllowPartialMessage  bool     `protobuf:"varint,2,opt,name=allow_partial_message,json=allowPartialMessage,proto3" json:"allow_partial_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BufferSettings) Reset()         { *m = BufferSettings{} }
func (m *BufferSettings) String() string { return proto.CompactTextString(m) }
func (*BufferSettings) ProtoMessage()    {}
func (*BufferSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20c629614df4358, []int{1}
}

func (m *BufferSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BufferSettings.Unmarshal(m, b)
}
func (m *BufferSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BufferSettings.Marshal(b, m, deterministic)
}
func (m *BufferSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BufferSettings.Merge(m, src)
}
func (m *BufferSettings) XXX_Size() int {
	return xxx_messageInfo_BufferSettings.Size(m)
}
func (m *BufferSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_BufferSettings.DiscardUnknown(m)
}

var xxx_messageInfo_BufferSettings proto.InternalMessageInfo

func (m *BufferSettings) GetMaxRequestBytes() uint32 {
	if m != nil {
		return m.MaxRequestBytes
	}
	return 0
}

func (m *BufferSettings) GetAllowPartialMessage() bool {
	if m != nil {
		return m.AllowPartialMessage
	}
	return false
}

type HttpService struct {
	ServerUri             *v3alpha.HttpUri       `protobuf:"bytes,1,opt,name=server_uri,json=serverUri,proto3" json:"server_uri,omitempty"`
	PathPrefix            string                 `protobuf:"bytes,2,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	AuthorizationRequest  *AuthorizationRequest  `protobuf:"bytes,7,opt,name=authorization_request,json=authorizationRequest,proto3" json:"authorization_request,omitempty"`
	AuthorizationResponse *AuthorizationResponse `protobuf:"bytes,8,opt,name=authorization_response,json=authorizationResponse,proto3" json:"authorization_response,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *HttpService) Reset()         { *m = HttpService{} }
func (m *HttpService) String() string { return proto.CompactTextString(m) }
func (*HttpService) ProtoMessage()    {}
func (*HttpService) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20c629614df4358, []int{2}
}

func (m *HttpService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpService.Unmarshal(m, b)
}
func (m *HttpService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpService.Marshal(b, m, deterministic)
}
func (m *HttpService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpService.Merge(m, src)
}
func (m *HttpService) XXX_Size() int {
	return xxx_messageInfo_HttpService.Size(m)
}
func (m *HttpService) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpService.DiscardUnknown(m)
}

var xxx_messageInfo_HttpService proto.InternalMessageInfo

func (m *HttpService) GetServerUri() *v3alpha.HttpUri {
	if m != nil {
		return m.ServerUri
	}
	return nil
}

func (m *HttpService) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

func (m *HttpService) GetAuthorizationRequest() *AuthorizationRequest {
	if m != nil {
		return m.AuthorizationRequest
	}
	return nil
}

func (m *HttpService) GetAuthorizationResponse() *AuthorizationResponse {
	if m != nil {
		return m.AuthorizationResponse
	}
	return nil
}

type AuthorizationRequest struct {
	AllowedHeaders       *v3alpha2.ListStringMatcher `protobuf:"bytes,1,opt,name=allowed_headers,json=allowedHeaders,proto3" json:"allowed_headers,omitempty"`
	HeadersToAdd         []*v3alpha.HeaderValue      `protobuf:"bytes,2,rep,name=headers_to_add,json=headersToAdd,proto3" json:"headers_to_add,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *AuthorizationRequest) Reset()         { *m = AuthorizationRequest{} }
func (m *AuthorizationRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorizationRequest) ProtoMessage()    {}
func (*AuthorizationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20c629614df4358, []int{3}
}

func (m *AuthorizationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationRequest.Unmarshal(m, b)
}
func (m *AuthorizationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationRequest.Marshal(b, m, deterministic)
}
func (m *AuthorizationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationRequest.Merge(m, src)
}
func (m *AuthorizationRequest) XXX_Size() int {
	return xxx_messageInfo_AuthorizationRequest.Size(m)
}
func (m *AuthorizationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationRequest proto.InternalMessageInfo

func (m *AuthorizationRequest) GetAllowedHeaders() *v3alpha2.ListStringMatcher {
	if m != nil {
		return m.AllowedHeaders
	}
	return nil
}

func (m *AuthorizationRequest) GetHeadersToAdd() []*v3alpha.HeaderValue {
	if m != nil {
		return m.HeadersToAdd
	}
	return nil
}

type AuthorizationResponse struct {
	AllowedUpstreamHeaders *v3alpha2.ListStringMatcher `protobuf:"bytes,1,opt,name=allowed_upstream_headers,json=allowedUpstreamHeaders,proto3" json:"allowed_upstream_headers,omitempty"`
	AllowedClientHeaders   *v3alpha2.ListStringMatcher `protobuf:"bytes,2,opt,name=allowed_client_headers,json=allowedClientHeaders,proto3" json:"allowed_client_headers,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                    `json:"-"`
	XXX_unrecognized       []byte                      `json:"-"`
	XXX_sizecache          int32                       `json:"-"`
}

func (m *AuthorizationResponse) Reset()         { *m = AuthorizationResponse{} }
func (m *AuthorizationResponse) String() string { return proto.CompactTextString(m) }
func (*AuthorizationResponse) ProtoMessage()    {}
func (*AuthorizationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20c629614df4358, []int{4}
}

func (m *AuthorizationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationResponse.Unmarshal(m, b)
}
func (m *AuthorizationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationResponse.Marshal(b, m, deterministic)
}
func (m *AuthorizationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationResponse.Merge(m, src)
}
func (m *AuthorizationResponse) XXX_Size() int {
	return xxx_messageInfo_AuthorizationResponse.Size(m)
}
func (m *AuthorizationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationResponse proto.InternalMessageInfo

func (m *AuthorizationResponse) GetAllowedUpstreamHeaders() *v3alpha2.ListStringMatcher {
	if m != nil {
		return m.AllowedUpstreamHeaders
	}
	return nil
}

func (m *AuthorizationResponse) GetAllowedClientHeaders() *v3alpha2.ListStringMatcher {
	if m != nil {
		return m.AllowedClientHeaders
	}
	return nil
}

type ExtAuthzPerRoute struct {
	// Types that are valid to be assigned to Override:
	//	*ExtAuthzPerRoute_Disabled
	//	*ExtAuthzPerRoute_CheckSettings
	Override             isExtAuthzPerRoute_Override `protobuf_oneof:"override"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ExtAuthzPerRoute) Reset()         { *m = ExtAuthzPerRoute{} }
func (m *ExtAuthzPerRoute) String() string { return proto.CompactTextString(m) }
func (*ExtAuthzPerRoute) ProtoMessage()    {}
func (*ExtAuthzPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20c629614df4358, []int{5}
}

func (m *ExtAuthzPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtAuthzPerRoute.Unmarshal(m, b)
}
func (m *ExtAuthzPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtAuthzPerRoute.Marshal(b, m, deterministic)
}
func (m *ExtAuthzPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthzPerRoute.Merge(m, src)
}
func (m *ExtAuthzPerRoute) XXX_Size() int {
	return xxx_messageInfo_ExtAuthzPerRoute.Size(m)
}
func (m *ExtAuthzPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthzPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthzPerRoute proto.InternalMessageInfo

type isExtAuthzPerRoute_Override interface {
	isExtAuthzPerRoute_Override()
}

type ExtAuthzPerRoute_Disabled struct {
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}

type ExtAuthzPerRoute_CheckSettings struct {
	CheckSettings *CheckSettings `protobuf:"bytes,2,opt,name=check_settings,json=checkSettings,proto3,oneof"`
}

func (*ExtAuthzPerRoute_Disabled) isExtAuthzPerRoute_Override() {}

func (*ExtAuthzPerRoute_CheckSettings) isExtAuthzPerRoute_Override() {}

func (m *ExtAuthzPerRoute) GetOverride() isExtAuthzPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *ExtAuthzPerRoute) GetDisabled() bool {
	if x, ok := m.GetOverride().(*ExtAuthzPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (m *ExtAuthzPerRoute) GetCheckSettings() *CheckSettings {
	if x, ok := m.GetOverride().(*ExtAuthzPerRoute_CheckSettings); ok {
		return x.CheckSettings
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExtAuthzPerRoute) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExtAuthzPerRoute_Disabled)(nil),
		(*ExtAuthzPerRoute_CheckSettings)(nil),
	}
}

type CheckSettings struct {
	ContextExtensions    map[string]string `protobuf:"bytes,1,rep,name=context_extensions,json=contextExtensions,proto3" json:"context_extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckSettings) Reset()         { *m = CheckSettings{} }
func (m *CheckSettings) String() string { return proto.CompactTextString(m) }
func (*CheckSettings) ProtoMessage()    {}
func (*CheckSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_b20c629614df4358, []int{6}
}

func (m *CheckSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckSettings.Unmarshal(m, b)
}
func (m *CheckSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckSettings.Marshal(b, m, deterministic)
}
func (m *CheckSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckSettings.Merge(m, src)
}
func (m *CheckSettings) XXX_Size() int {
	return xxx_messageInfo_CheckSettings.Size(m)
}
func (m *CheckSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckSettings.DiscardUnknown(m)
}

var xxx_messageInfo_CheckSettings proto.InternalMessageInfo

func (m *CheckSettings) GetContextExtensions() map[string]string {
	if m != nil {
		return m.ContextExtensions
	}
	return nil
}

func init() {
	proto.RegisterType((*ExtAuthz)(nil), "envoy.extensions.filters.http.ext_authz.v3alpha.ExtAuthz")
	proto.RegisterType((*BufferSettings)(nil), "envoy.extensions.filters.http.ext_authz.v3alpha.BufferSettings")
	proto.RegisterType((*HttpService)(nil), "envoy.extensions.filters.http.ext_authz.v3alpha.HttpService")
	proto.RegisterType((*AuthorizationRequest)(nil), "envoy.extensions.filters.http.ext_authz.v3alpha.AuthorizationRequest")
	proto.RegisterType((*AuthorizationResponse)(nil), "envoy.extensions.filters.http.ext_authz.v3alpha.AuthorizationResponse")
	proto.RegisterType((*ExtAuthzPerRoute)(nil), "envoy.extensions.filters.http.ext_authz.v3alpha.ExtAuthzPerRoute")
	proto.RegisterType((*CheckSettings)(nil), "envoy.extensions.filters.http.ext_authz.v3alpha.CheckSettings")
	proto.RegisterMapType((map[string]string)(nil), "envoy.extensions.filters.http.ext_authz.v3alpha.CheckSettings.ContextExtensionsEntry")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/ext_authz/v3alpha/ext_authz.proto", fileDescriptor_b20c629614df4358)
}

var fileDescriptor_b20c629614df4358 = []byte{
	// 1209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xee, 0xda, 0x4e, 0xba, 0x99, 0x34, 0x89, 0x33, 0x24, 0x61, 0x29, 0xa2, 0xa4, 0xa6, 0x40,
	0x54, 0x15, 0x1b, 0x1a, 0x5a, 0x45, 0xee, 0x9f, 0xec, 0xd4, 0x25, 0x0a, 0x2d, 0x58, 0x5b, 0xd2,
	0x0b, 0x6e, 0x46, 0x93, 0xdd, 0x63, 0x7b, 0xe8, 0x7a, 0x77, 0x3b, 0x33, 0xeb, 0xc6, 0x95, 0xb8,
	0xe3, 0x02, 0xf1, 0x04, 0x88, 0x37, 0xe0, 0x05, 0x2a, 0xae, 0xb8, 0xe3, 0x05, 0x78, 0x10, 0x2e,
	0xb8, 0x42, 0x15, 0x12, 0x68, 0x7e, 0xd6, 0x3f, 0x25, 0x2d, 0x71, 0x7b, 0xe7, 0x3d, 0x3f, 0xdf,
	0x39, 0xdf, 0x77, 0xce, 0xcc, 0x18, 0xdd, 0x82, 0x78, 0x90, 0x0c, 0x6b, 0x70, 0x24, 0x21, 0x16,
	0x2c, 0x89, 0x45, 0xad, 0xc3, 0x22, 0x09, 0x5c, 0xd4, 0x7a, 0x52, 0xa6, 0xca, 0x4e, 0x68, 0x26,
	0x7b, 0x4f, 0x6a, 0x83, 0x6d, 0x1a, 0xa5, 0x3d, 0x3a, 0xb6, 0x54, 0x53, 0x9e, 0xc8, 0x04, 0xd7,
	0x34, 0x40, 0x75, 0x0c, 0x50, 0xb5, 0x00, 0x55, 0x05, 0x50, 0x1d, 0x87, 0x5b, 0x80, 0xb3, 0x17,
	0x4c, 0xc5, 0x20, 0x89, 0x3b, 0xac, 0x5b, 0x0b, 0x12, 0x0e, 0x23, 0xec, 0x43, 0x2a, 0xc0, 0xc0,
	0x9e, 0xbd, 0xf4, 0xe2, 0xa8, 0x2e, 0x4f, 0x03, 0x22, 0x80, 0x0f, 0x58, 0x90, 0x47, 0x6f, 0xbd,
	0x38, 0x5a, 0x75, 0x41, 0x32, 0xce, 0x6c, 0xe4, 0x87, 0x26, 0x52, 0x0e, 0x53, 0xa8, 0xf5, 0xa9,
	0x0c, 0x7a, 0xc0, 0x47, 0xa1, 0x42, 0x72, 0x16, 0x77, 0x6d, 0xe0, 0x85, 0x89, 0xc0, 0x29, 0x2c,
	0x21, 0xa9, 0xcc, 0x84, 0x8d, 0x3a, 0x9f, 0x85, 0x29, 0xad, 0xd1, 0x38, 0x4e, 0x24, 0x95, 0x5a,
	0xbe, 0x01, 0x70, 0x25, 0xc3, 0x18, 0xe8, 0x3d, 0x03, 0x34, 0x19, 0x13, 0x42, 0xca, 0x21, 0xd0,
	0x1f, 0x36, 0xe8, 0xcd, 0x01, 0x8d, 0x58, 0x48, 0x25, 0xd4, 0xf2, 0x1f, 0xc6, 0x51, 0x79, 0x3a,
	0x8f, 0xdc, 0xd6, 0x91, 0x6c, 0x28, 0x09, 0xf1, 0xe7, 0xe8, 0xcc, 0x24, 0x79, 0xcf, 0xd9, 0x74,
	0xb6, 0x16, 0x2f, 0x7f, 0x50, 0x35, 0x23, 0x30, 0xec, 0xab, 0x8a, 0x7d, 0x2e, 0x76, 0xf5, 0x33,
	0x9e, 0x06, 0xf7, 0x4d, 0xf4, 0xde, 0x29, 0x7f, 0xb1, 0x3b, 0xfe, 0xc4, 0x14, 0x9d, 0x31, 0x7c,
	0x2c, 0x58, 0x51, 0x83, 0x5d, 0xaf, 0xce, 0x38, 0xcf, 0xea, 0x9e, 0x94, 0xe9, 0x44, 0x89, 0xde,
	0xf8, 0x13, 0x5f, 0x42, 0xb8, 0x43, 0x59, 0x94, 0x71, 0x20, 0xfd, 0x24, 0x04, 0x42, 0xa3, 0x28,
	0x79, 0xec, 0x15, 0x36, 0x9d, 0x2d, 0xd7, 0x2f, 0x5b, 0xcf, 0xbd, 0x24, 0x84, 0x86, 0xb2, 0x63,
	0x1f, 0x9d, 0xef, 0xb1, 0x30, 0x84, 0x98, 0xe8, 0x16, 0x48, 0xae, 0x12, 0x84, 0x24, 0x13, 0x2a,
	0x39, 0xed, 0x51, 0xaf, 0xa4, 0x92, 0x9b, 0xee, 0x2f, 0x7f, 0xfc, 0xf9, 0xfb, 0x9c, 0xe3, 0x39,
	0xfe, 0x3b, 0x26, 0xa5, 0xa5, 0x32, 0x6e, 0x8f, 0x12, 0x0e, 0x04, 0x34, 0x54, 0x38, 0x7e, 0x88,
	0x56, 0x1f, 0x33, 0xd9, 0x23, 0x1c, 0x1e, 0x65, 0x20, 0x24, 0x39, 0x4c, 0xc2, 0xa1, 0x37, 0xa7,
	0x99, 0xde, 0x9a, 0x99, 0x69, 0x33, 0xeb, 0x74, 0x80, 0xdf, 0x07, 0x29, 0x59, 0xdc, 0x15, 0xfe,
	0x8a, 0x42, 0xf6, 0x0d, 0x70, 0x33, 0x09, 0x87, 0xf8, 0x22, 0x5a, 0x0d, 0x22, 0xa0, 0x9c, 0xf0,
	0x24, 0x93, 0x40, 0x02, 0x1a, 0xf4, 0xc0, 0x9b, 0xd7, 0x6c, 0x57, 0xb4, 0xc3, 0x57, 0xf6, 0x5d,
	0x65, 0xc6, 0x77, 0xd0, 0x8a, 0x59, 0x24, 0x92, 0xc4, 0x04, 0x38, 0x4f, 0xb8, 0x77, 0x5a, 0xb7,
	0x75, 0xce, 0xb6, 0xa5, 0x16, 0x6f, 0x5a, 0x63, 0x1d, 0xee, 0x2f, 0x99, 0xb4, 0x2f, 0xe3, 0x96,
	0x4a, 0xc2, 0x37, 0xd1, 0xdb, 0x7d, 0x90, 0x34, 0xa4, 0x92, 0x92, 0x20, 0x89, 0xa5, 0xea, 0x3a,
	0xa6, 0x7d, 0x10, 0x29, 0x0d, 0x40, 0x78, 0xee, 0x66, 0x71, 0x6b, 0xc1, 0x7f, 0x2b, 0x0f, 0xd9,
	0x35, 0x11, 0x5f, 0x8c, 0x02, 0xf0, 0xd7, 0x68, 0xd9, 0xb0, 0x26, 0x10, 0xd3, 0xc3, 0x08, 0x42,
	0x6f, 0x41, 0xb7, 0xb1, 0xfd, 0x92, 0xa5, 0xf2, 0xb3, 0x58, 0xb2, 0x3e, 0xdc, 0xe1, 0x34, 0x50,
	0x2b, 0x4c, 0xa3, 0x36, 0xf0, 0x00, 0x62, 0xe9, 0x2f, 0x19, 0xa8, 0x96, 0x41, 0xc2, 0x3b, 0xc8,
	0x63, 0x71, 0x10, 0x65, 0x21, 0x90, 0x14, 0x80, 0x93, 0x00, 0xb8, 0x64, 0x1d, 0xa6, 0x26, 0xe4,
	0x21, 0x2d, 0xcb, 0x86, 0xf5, 0xb7, 0x01, 0xf8, 0xee, 0xd8, 0x5b, 0xbf, 0xf2, 0xd3, 0x6f, 0xdf,
	0x9f, 0xfb, 0x18, 0x4d, 0xf7, 0x60, 0xc0, 0xff, 0x33, 0x9c, 0xcb, 0xd5, 0xfc, 0x7c, 0x34, 0x11,
	0x72, 0xed, 0x36, 0x8b, 0xca, 0x53, 0x07, 0x2d, 0x4f, 0x0f, 0x0c, 0x6f, 0xa3, 0xd5, 0x3e, 0x3d,
	0x1a, 0xef, 0xc2, 0x50, 0x82, 0xd0, 0x67, 0x68, 0xa9, 0x79, 0xfa, 0x59, 0xb3, 0x74, 0xb1, 0xb0,
	0x79, 0xca, 0x5f, 0xe9, 0xd3, 0xa3, 0x7c, 0xa6, 0xca, 0x8f, 0x2f, 0xa3, 0x75, 0xbd, 0xb6, 0x24,
	0xa5, 0x5c, 0x32, 0x1a, 0x91, 0x3e, 0x08, 0x41, 0xbb, 0x60, 0xd7, 0xf8, 0x0d, 0xed, 0x6c, 0x1b,
	0xdf, 0x3d, 0xe3, 0xaa, 0x5f, 0x53, 0xed, 0x5f, 0x45, 0x9f, 0x9e, 0xac, 0xfd, 0xe9, 0x2e, 0x2b,
	0xbf, 0x16, 0xd1, 0xe2, 0xc4, 0x99, 0xc2, 0x0d, 0x84, 0x14, 0x29, 0xe0, 0xea, 0x16, 0xb3, 0x47,
	0xbe, 0xf2, 0x92, 0xe9, 0xa8, 0xdc, 0x03, 0xce, 0xfc, 0x05, 0x93, 0x75, 0xc0, 0x19, 0x7e, 0x17,
	0x2d, 0xa6, 0x54, 0xf6, 0x48, 0xca, 0xa1, 0xc3, 0x8e, 0x74, 0xe7, 0x0b, 0x3e, 0x52, 0xa6, 0xb6,
	0xb6, 0xe0, 0x27, 0x68, 0x5d, 0xb5, 0x93, 0x70, 0xf6, 0x44, 0xdf, 0x4a, 0xb9, 0x46, 0x76, 0x27,
	0x5b, 0x33, 0x1f, 0x95, 0xc6, 0x24, 0x9a, 0xd5, 0xd3, 0x5f, 0xa3, 0xc7, 0x58, 0xf1, 0xb7, 0x68,
	0xe3, 0xf9, 0xda, 0x22, 0x4d, 0x62, 0x01, 0x9e, 0xab, 0x8b, 0xdf, 0x79, 0xdd, 0xe2, 0x06, 0xcd,
	0x5f, 0xa7, 0xc7, 0x99, 0xeb, 0x3b, 0x6a, 0x56, 0xdb, 0xe8, 0x93, 0x93, 0xcd, 0x6a, 0x62, 0x30,
	0xfb, 0x25, 0xb7, 0x58, 0x2e, 0xed, 0x97, 0xdc, 0x52, 0x79, 0x6e, 0xbf, 0xe4, 0xce, 0x95, 0xe7,
	0xf7, 0x4b, 0xee, 0x7c, 0xf9, 0x74, 0xe5, 0x1f, 0x07, 0xad, 0x1d, 0xc7, 0x1f, 0x3f, 0x40, 0x2b,
	0x7a, 0x5b, 0x20, 0x24, 0x3d, 0xa0, 0x21, 0x70, 0x61, 0xc7, 0xf9, 0xd1, 0xe4, 0x99, 0xb7, 0xaf,
	0xd2, 0x88, 0xcd, 0x5d, 0x26, 0xe4, 0x7d, 0xfd, 0x32, 0xdd, 0x33, 0x1e, 0x7f, 0xd9, 0xa2, 0xec,
	0x19, 0x10, 0x7c, 0x17, 0x2d, 0x5b, 0x3c, 0x22, 0x13, 0x42, 0xc3, 0xd0, 0x2b, 0x6c, 0x16, 0xff,
	0xe7, 0x61, 0x30, 0xb9, 0x0f, 0x68, 0x94, 0x81, 0x7f, 0xc6, 0x66, 0x7f, 0x95, 0x34, 0xc2, 0xb0,
	0xde, 0x50, 0x82, 0x5c, 0x47, 0xf5, 0x93, 0x09, 0x72, 0x1c, 0xd1, 0xca, 0xcf, 0x05, 0xb4, 0x7e,
	0xec, 0x10, 0x70, 0x17, 0x79, 0xb9, 0x04, 0x59, 0x2a, 0x24, 0x07, 0xda, 0x7f, 0x3d, 0x2d, 0x36,
	0x2c, 0xdc, 0x81, 0x45, 0xcb, 0x35, 0x09, 0x50, 0xee, 0x21, 0x41, 0xc4, 0x20, 0x96, 0xa3, 0x32,
	0x85, 0x57, 0x29, 0xb3, 0x66, 0xc1, 0x76, 0x35, 0x96, 0x2d, 0x52, 0x6f, 0x2a, 0xa9, 0x6e, 0xa0,
	0x6b, 0xaf, 0x24, 0x95, 0x51, 0xa4, 0xf2, 0xb7, 0x83, 0xca, 0xf9, 0x05, 0xd6, 0x06, 0xf3, 0x44,
	0xe0, 0xf7, 0x91, 0x1b, 0x32, 0x61, 0xee, 0x63, 0x47, 0xbf, 0x78, 0xea, 0x82, 0xfa, 0xa6, 0xe0,
	0x3a, 0x7b, 0xa7, 0xfc, 0x91, 0x0b, 0x3f, 0x42, 0xcb, 0x41, 0x0f, 0x82, 0x87, 0x44, 0xd8, 0xcb,
	0xc3, 0x92, 0xbb, 0x39, 0xf3, 0x91, 0xd9, 0x55, 0x30, 0xf9, 0x15, 0xd4, 0x74, 0x9f, 0x35, 0xe7,
	0x7e, 0x70, 0x0a, 0x65, 0x55, 0x6d, 0x29, 0x98, 0x74, 0xd5, 0x6f, 0x28, 0xca, 0x3b, 0xe8, 0xea,
	0x6c, 0x37, 0x73, 0x4e, 0xac, 0xb9, 0x82, 0xdc, 0x64, 0x00, 0x9c, 0xb3, 0x10, 0x70, 0xf1, 0xaf,
	0xa6, 0x53, 0xf9, 0xb1, 0x80, 0x96, 0xa6, 0x8a, 0xe3, 0xef, 0x1c, 0x84, 0xf3, 0x97, 0x6c, 0x4c,
	0xc0, 0x73, 0xf4, 0x4a, 0x1f, 0xbc, 0x1e, 0xb3, 0xaa, 0x7d, 0x00, 0x5b, 0xa3, 0xf4, 0x56, 0x2c,
	0xf9, 0xd0, 0x5f, 0x0d, 0x9e, 0xb7, 0x9f, 0xbd, 0x8d, 0x36, 0x8e, 0x0f, 0xc6, 0x65, 0x54, 0x7c,
	0x08, 0x43, 0x3d, 0x97, 0x05, 0x5f, 0xfd, 0xc4, 0x6b, 0x68, 0x6e, 0xa0, 0x4e, 0x92, 0xbd, 0x59,
	0xcd, 0x47, 0xbd, 0xb0, 0xe3, 0xd4, 0xeb, 0x4a, 0xae, 0x2b, 0x68, 0xfb, 0x64, 0x72, 0x4d, 0x4f,
	0xa1, 0x8d, 0x6e, 0xb0, 0xc4, 0xf0, 0x4d, 0x79, 0x72, 0x34, 0x9c, 0x95, 0x7a, 0x73, 0x69, 0x24,
	0xbf, 0xfa, 0x2b, 0xd9, 0x76, 0x0e, 0xe7, 0xf5, 0x7f, 0xca, 0xed, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x75, 0x6f, 0xb5, 0xae, 0xf5, 0x0b, 0x00, 0x00,
}
