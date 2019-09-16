// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: appcommon.proto

package gencmd

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT
var AppPortRequiredArgs = []string{}
var AppPortOptionalArgs = []string{
	"proto",
	"internalport",
	"publicport",
	"pathprefix",
	"fqdnprefix",
	"endport",
}
var AppPortAliasArgs = []string{}
var AppPortComments = map[string]string{
	"proto":        "TCP (L4), UDP (L4), or HTTP (L7) protocol, one of LProtoUnknown, LProtoTcp, LProtoUdp, LProtoHttp",
	"internalport": "Container port",
	"publicport":   "Public facing port for TCP/UDP (may be mapped on shared LB reverse proxy)",
	"pathprefix":   "Public facing path for HTTP L7 access.",
	"fqdnprefix":   "FQDN prefix to append to base FQDN in FindCloudlet response. May be empty.",
	"endport":      "A non-zero end port indicates this is a port range from internal port to end port, inclusive.",
}
var AppPortSpecialArgs = map[string]string{}
