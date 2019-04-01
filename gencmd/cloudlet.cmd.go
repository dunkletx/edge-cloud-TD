// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cloudlet.proto

package gencmd

import distributed_match_engine "github.com/mobiledgex/edge-cloud/d-match-engine/dme-proto"
import edgeproto "github.com/mobiledgex/edge-cloud/edgeproto"
import "strings"
import "strconv"
import "github.com/spf13/cobra"
import "context"
import "os"
import "io"
import "text/tabwriter"
import "github.com/spf13/pflag"
import "errors"
import "github.com/mobiledgex/edge-cloud/protoc-gen-cmd/cmdsup"
import "google.golang.org/grpc/status"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/mobiledgex/edge-cloud/protogen"
import _ "github.com/mobiledgex/edge-cloud/protoc-gen-cmd/protocmd"
import _ "github.com/mobiledgex/edge-cloud/d-match-engine/dme-proto"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT
var CloudletApiCmd edgeproto.CloudletApiClient
var CloudletInfoApiCmd edgeproto.CloudletInfoApiClient
var CloudletMetricsApiCmd edgeproto.CloudletMetricsApiClient
var CloudletIn edgeproto.Cloudlet
var CloudletFlagSet = pflag.NewFlagSet("Cloudlet", pflag.ExitOnError)
var CloudletNoConfigFlagSet = pflag.NewFlagSet("CloudletNoConfig", pflag.ExitOnError)
var CloudletInIpSupport string
var CloudletInfoIn edgeproto.CloudletInfo
var CloudletInfoFlagSet = pflag.NewFlagSet("CloudletInfo", pflag.ExitOnError)
var CloudletInfoNoConfigFlagSet = pflag.NewFlagSet("CloudletInfoNoConfig", pflag.ExitOnError)
var CloudletInfoInState string
var CloudletMetricsIn edgeproto.CloudletMetrics
var CloudletMetricsFlagSet = pflag.NewFlagSet("CloudletMetrics", pflag.ExitOnError)
var CloudletMetricsNoConfigFlagSet = pflag.NewFlagSet("CloudletMetricsNoConfig", pflag.ExitOnError)
var CloudletStateStrings = []string{
	"CloudletStateUnknown",
	"CloudletStateErrors",
	"CloudletStateReady",
	"CloudletStateOffline",
	"CloudletStateNotPresent",
}

func CloudletKeySlicer(in *edgeproto.CloudletKey) []string {
	s := make([]string, 0, 2)
	s = append(s, in.OperatorKey.Name)
	s = append(s, in.Name)
	return s
}

func CloudletKeyHeaderSlicer() []string {
	s := make([]string, 0, 2)
	s = append(s, "OperatorKey-Name")
	s = append(s, "Name")
	return s
}

func CloudletKeyWriteOutputArray(objs []*edgeproto.CloudletKey) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(CloudletKeyHeaderSlicer(), "\t"))
		for _, obj := range objs {
			fmt.Fprintln(output, strings.Join(CloudletKeySlicer(obj), "\t"))
		}
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(objs)
	}
}

func CloudletKeyWriteOutputOne(obj *edgeproto.CloudletKey) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(CloudletKeyHeaderSlicer(), "\t"))
		fmt.Fprintln(output, strings.Join(CloudletKeySlicer(obj), "\t"))
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(obj)
	}
}
func CloudletInfraCommonSlicer(in *edgeproto.CloudletInfraCommon) []string {
	s := make([]string, 0, 7)
	s = append(s, in.DockerRegistry)
	s = append(s, in.DNSZone)
	s = append(s, in.RegistryFileServer)
	s = append(s, in.CFKey)
	s = append(s, in.CFUser)
	s = append(s, in.DockerRegPass)
	s = append(s, in.NetworkScheme)
	return s
}

func CloudletInfraCommonHeaderSlicer() []string {
	s := make([]string, 0, 7)
	s = append(s, "DockerRegistry")
	s = append(s, "DNSZone")
	s = append(s, "RegistryFileServer")
	s = append(s, "CFKey")
	s = append(s, "CFUser")
	s = append(s, "DockerRegPass")
	s = append(s, "NetworkScheme")
	return s
}

func CloudletInfraCommonWriteOutputArray(objs []*edgeproto.CloudletInfraCommon) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(CloudletInfraCommonHeaderSlicer(), "\t"))
		for _, obj := range objs {
			fmt.Fprintln(output, strings.Join(CloudletInfraCommonSlicer(obj), "\t"))
		}
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(objs)
	}
}

func CloudletInfraCommonWriteOutputOne(obj *edgeproto.CloudletInfraCommon) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(CloudletInfraCommonHeaderSlicer(), "\t"))
		fmt.Fprintln(output, strings.Join(CloudletInfraCommonSlicer(obj), "\t"))
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(obj)
	}
}
func AzurePropertiesSlicer(in *edgeproto.AzureProperties) []string {
	s := make([]string, 0, 4)
	s = append(s, in.Location)
	s = append(s, in.ResourceGroup)
	s = append(s, in.UserName)
	s = append(s, in.Password)
	return s
}

func AzurePropertiesHeaderSlicer() []string {
	s := make([]string, 0, 4)
	s = append(s, "Location")
	s = append(s, "ResourceGroup")
	s = append(s, "UserName")
	s = append(s, "Password")
	return s
}

func AzurePropertiesWriteOutputArray(objs []*edgeproto.AzureProperties) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(AzurePropertiesHeaderSlicer(), "\t"))
		for _, obj := range objs {
			fmt.Fprintln(output, strings.Join(AzurePropertiesSlicer(obj), "\t"))
		}
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(objs)
	}
}

func AzurePropertiesWriteOutputOne(obj *edgeproto.AzureProperties) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(AzurePropertiesHeaderSlicer(), "\t"))
		fmt.Fprintln(output, strings.Join(AzurePropertiesSlicer(obj), "\t"))
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(obj)
	}
}
func GcpPropertiesSlicer(in *edgeproto.GcpProperties) []string {
	s := make([]string, 0, 2)
	s = append(s, in.Project)
	s = append(s, in.Zone)
	return s
}

func GcpPropertiesHeaderSlicer() []string {
	s := make([]string, 0, 2)
	s = append(s, "Project")
	s = append(s, "Zone")
	return s
}

func GcpPropertiesWriteOutputArray(objs []*edgeproto.GcpProperties) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(GcpPropertiesHeaderSlicer(), "\t"))
		for _, obj := range objs {
			fmt.Fprintln(output, strings.Join(GcpPropertiesSlicer(obj), "\t"))
		}
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(objs)
	}
}

func GcpPropertiesWriteOutputOne(obj *edgeproto.GcpProperties) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(GcpPropertiesHeaderSlicer(), "\t"))
		fmt.Fprintln(output, strings.Join(GcpPropertiesSlicer(obj), "\t"))
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(obj)
	}
}
func OpenStackPropertiesSlicer(in *edgeproto.OpenStackProperties) []string {
	s := make([]string, 0, 5)
	s = append(s, in.OSExternalNetworkName)
	s = append(s, in.OSImageName)
	s = append(s, in.OSExternalRouterName)
	s = append(s, in.OSMexNetwork)
	return s
}

func OpenStackPropertiesHeaderSlicer() []string {
	s := make([]string, 0, 5)
	s = append(s, "OSExternalNetworkName")
	s = append(s, "OSImageName")
	s = append(s, "OSExternalRouterName")
	s = append(s, "OSMexNetwork")
	return s
}

func OpenStackPropertiesWriteOutputArray(objs []*edgeproto.OpenStackProperties) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(OpenStackPropertiesHeaderSlicer(), "\t"))
		for _, obj := range objs {
			fmt.Fprintln(output, strings.Join(OpenStackPropertiesSlicer(obj), "\t"))
		}
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(objs)
	}
}

func OpenStackPropertiesWriteOutputOne(obj *edgeproto.OpenStackProperties) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(OpenStackPropertiesHeaderSlicer(), "\t"))
		fmt.Fprintln(output, strings.Join(OpenStackPropertiesSlicer(obj), "\t"))
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(obj)
	}
}
func CloudletInfraPropertiesSlicer(in *edgeproto.CloudletInfraProperties) []string {
	s := make([]string, 0, 5)
	s = append(s, in.CloudletKind)
	s = append(s, in.MexosContainerImageName)
	if in.OpenstackProperties == nil {
		in.OpenstackProperties = &edgeproto.OpenStackProperties{}
	}
	s = append(s, in.OpenstackProperties.OSExternalNetworkName)
	s = append(s, in.OpenstackProperties.OSImageName)
	s = append(s, in.OpenstackProperties.OSExternalRouterName)
	s = append(s, in.OpenstackProperties.OSMexNetwork)
	if in.AzureProperties == nil {
		in.AzureProperties = &edgeproto.AzureProperties{}
	}
	s = append(s, in.AzureProperties.Location)
	s = append(s, in.AzureProperties.ResourceGroup)
	s = append(s, in.AzureProperties.UserName)
	s = append(s, in.AzureProperties.Password)
	if in.GcpProperties == nil {
		in.GcpProperties = &edgeproto.GcpProperties{}
	}
	s = append(s, in.GcpProperties.Project)
	s = append(s, in.GcpProperties.Zone)
	return s
}

func CloudletInfraPropertiesHeaderSlicer() []string {
	s := make([]string, 0, 5)
	s = append(s, "CloudletKind")
	s = append(s, "MexosContainerImageName")
	s = append(s, "OpenstackProperties-OSExternalNetworkName")
	s = append(s, "OpenstackProperties-OSImageName")
	s = append(s, "OpenstackProperties-OSExternalRouterName")
	s = append(s, "OpenstackProperties-OSMexNetwork")
	s = append(s, "AzureProperties-Location")
	s = append(s, "AzureProperties-ResourceGroup")
	s = append(s, "AzureProperties-UserName")
	s = append(s, "AzureProperties-Password")
	s = append(s, "GcpProperties-Project")
	s = append(s, "GcpProperties-Zone")
	return s
}

func CloudletInfraPropertiesWriteOutputArray(objs []*edgeproto.CloudletInfraProperties) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(CloudletInfraPropertiesHeaderSlicer(), "\t"))
		for _, obj := range objs {
			fmt.Fprintln(output, strings.Join(CloudletInfraPropertiesSlicer(obj), "\t"))
		}
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(objs)
	}
}

func CloudletInfraPropertiesWriteOutputOne(obj *edgeproto.CloudletInfraProperties) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(CloudletInfraPropertiesHeaderSlicer(), "\t"))
		fmt.Fprintln(output, strings.Join(CloudletInfraPropertiesSlicer(obj), "\t"))
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(obj)
	}
}
func CloudletSlicer(in *edgeproto.Cloudlet) []string {
	s := make([]string, 0, 7)
	if in.Fields == nil {
		in.Fields = make([]string, 1)
	}
	s = append(s, in.Fields[0])
	s = append(s, in.Key.OperatorKey.Name)
	s = append(s, in.Key.Name)
	s = append(s, in.AccessUri)
	s = append(s, strconv.FormatFloat(float64(in.Location.Latitude), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.Location.Longitude), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.Location.HorizontalAccuracy), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.Location.VerticalAccuracy), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.Location.Altitude), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.Location.Course), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.Location.Speed), 'e', -1, 32))
	if in.Location.Timestamp == nil {
		in.Location.Timestamp = &distributed_match_engine.Timestamp{}
	}
	s = append(s, strconv.FormatUint(uint64(in.Location.Timestamp.Seconds), 10))
	s = append(s, strconv.FormatUint(uint64(in.Location.Timestamp.Nanos), 10))
	s = append(s, edgeproto.IpSupport_name[int32(in.IpSupport)])
	s = append(s, in.StaticIps)
	s = append(s, strconv.FormatUint(uint64(in.NumDynamicIps), 10))
	return s
}

func CloudletHeaderSlicer() []string {
	s := make([]string, 0, 7)
	s = append(s, "Fields")
	s = append(s, "Key-OperatorKey-Name")
	s = append(s, "Key-Name")
	s = append(s, "AccessUri")
	s = append(s, "Location-Latitude")
	s = append(s, "Location-Longitude")
	s = append(s, "Location-HorizontalAccuracy")
	s = append(s, "Location-VerticalAccuracy")
	s = append(s, "Location-Altitude")
	s = append(s, "Location-Course")
	s = append(s, "Location-Speed")
	s = append(s, "Location-Timestamp-Seconds")
	s = append(s, "Location-Timestamp-Nanos")
	s = append(s, "IpSupport")
	s = append(s, "StaticIps")
	s = append(s, "NumDynamicIps")
	return s
}

func CloudletWriteOutputArray(objs []*edgeproto.Cloudlet) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(CloudletHeaderSlicer(), "\t"))
		for _, obj := range objs {
			fmt.Fprintln(output, strings.Join(CloudletSlicer(obj), "\t"))
		}
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(objs)
	}
}

func CloudletWriteOutputOne(obj *edgeproto.Cloudlet) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(CloudletHeaderSlicer(), "\t"))
		fmt.Fprintln(output, strings.Join(CloudletSlicer(obj), "\t"))
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(obj)
	}
}
func FlavorInfoSlicer(in *edgeproto.FlavorInfo) []string {
	s := make([]string, 0, 4)
	s = append(s, in.Name)
	s = append(s, strconv.FormatUint(uint64(in.Vcpus), 10))
	s = append(s, strconv.FormatUint(uint64(in.Ram), 10))
	s = append(s, strconv.FormatUint(uint64(in.Disk), 10))
	return s
}

func FlavorInfoHeaderSlicer() []string {
	s := make([]string, 0, 4)
	s = append(s, "Name")
	s = append(s, "Vcpus")
	s = append(s, "Ram")
	s = append(s, "Disk")
	return s
}

func FlavorInfoWriteOutputArray(objs []*edgeproto.FlavorInfo) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(FlavorInfoHeaderSlicer(), "\t"))
		for _, obj := range objs {
			fmt.Fprintln(output, strings.Join(FlavorInfoSlicer(obj), "\t"))
		}
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(objs)
	}
}

func FlavorInfoWriteOutputOne(obj *edgeproto.FlavorInfo) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(FlavorInfoHeaderSlicer(), "\t"))
		fmt.Fprintln(output, strings.Join(FlavorInfoSlicer(obj), "\t"))
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(obj)
	}
}
func CloudletInfoSlicer(in *edgeproto.CloudletInfo) []string {
	s := make([]string, 0, 10)
	if in.Fields == nil {
		in.Fields = make([]string, 1)
	}
	s = append(s, in.Fields[0])
	s = append(s, in.Key.OperatorKey.Name)
	s = append(s, in.Key.Name)
	s = append(s, edgeproto.CloudletState_name[int32(in.State)])
	s = append(s, strconv.FormatUint(uint64(in.NotifyId), 10))
	s = append(s, in.Controller)
	s = append(s, strconv.FormatUint(uint64(in.OsMaxRam), 10))
	s = append(s, strconv.FormatUint(uint64(in.OsMaxVcores), 10))
	s = append(s, strconv.FormatUint(uint64(in.OsMaxVolGb), 10))
	if in.Errors == nil {
		in.Errors = make([]string, 1)
	}
	s = append(s, in.Errors[0])
	if in.Flavors == nil {
		in.Flavors = make([]*edgeproto.FlavorInfo, 1)
	}
	if in.Flavors[0] == nil {
		in.Flavors[0] = &edgeproto.FlavorInfo{}
	}
	s = append(s, in.Flavors[0].Name)
	s = append(s, strconv.FormatUint(uint64(in.Flavors[0].Vcpus), 10))
	s = append(s, strconv.FormatUint(uint64(in.Flavors[0].Ram), 10))
	s = append(s, strconv.FormatUint(uint64(in.Flavors[0].Disk), 10))
	return s
}

func CloudletInfoHeaderSlicer() []string {
	s := make([]string, 0, 10)
	s = append(s, "Fields")
	s = append(s, "Key-OperatorKey-Name")
	s = append(s, "Key-Name")
	s = append(s, "State")
	s = append(s, "NotifyId")
	s = append(s, "Controller")
	s = append(s, "OsMaxRam")
	s = append(s, "OsMaxVcores")
	s = append(s, "OsMaxVolGb")
	s = append(s, "Errors")
	s = append(s, "Flavors-Name")
	s = append(s, "Flavors-Vcpus")
	s = append(s, "Flavors-Ram")
	s = append(s, "Flavors-Disk")
	return s
}

func CloudletInfoWriteOutputArray(objs []*edgeproto.CloudletInfo) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(CloudletInfoHeaderSlicer(), "\t"))
		for _, obj := range objs {
			fmt.Fprintln(output, strings.Join(CloudletInfoSlicer(obj), "\t"))
		}
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(objs)
	}
}

func CloudletInfoWriteOutputOne(obj *edgeproto.CloudletInfo) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(CloudletInfoHeaderSlicer(), "\t"))
		fmt.Fprintln(output, strings.Join(CloudletInfoSlicer(obj), "\t"))
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(obj)
	}
}
func CloudletMetricsSlicer(in *edgeproto.CloudletMetrics) []string {
	s := make([]string, 0, 1)
	s = append(s, strconv.FormatUint(uint64(in.Foo), 10))
	return s
}

func CloudletMetricsHeaderSlicer() []string {
	s := make([]string, 0, 1)
	s = append(s, "Foo")
	return s
}

func CloudletMetricsWriteOutputArray(objs []*edgeproto.CloudletMetrics) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(CloudletMetricsHeaderSlicer(), "\t"))
		for _, obj := range objs {
			fmt.Fprintln(output, strings.Join(CloudletMetricsSlicer(obj), "\t"))
		}
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(objs)
	}
}

func CloudletMetricsWriteOutputOne(obj *edgeproto.CloudletMetrics) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(CloudletMetricsHeaderSlicer(), "\t"))
		fmt.Fprintln(output, strings.Join(CloudletMetricsSlicer(obj), "\t"))
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(obj)
	}
}

var CreateCloudletCmd = &cobra.Command{
	Use: "CreateCloudlet",
	RunE: func(cmd *cobra.Command, args []string) error {
		// if we got this far, usage has been met.
		cmd.SilenceUsage = true
		err := parseCloudletEnums()
		if err != nil {
			return fmt.Errorf("CreateCloudlet failed: %s", err.Error())
		}
		return CreateCloudlet(&CloudletIn)
	},
}

func CreateCloudlet(in *edgeproto.Cloudlet) error {
	if CloudletApiCmd == nil {
		return fmt.Errorf("CloudletApi client not initialized")
	}
	ctx := context.Background()
	stream, err := CloudletApiCmd.CreateCloudlet(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("CreateCloudlet failed: %s", errstr)
	}
	objs := make([]*edgeproto.Result, 0)
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("CreateCloudlet recv failed: %s", err.Error())
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		return nil
	}
	ResultWriteOutputArray(objs)
	return nil
}

func CreateCloudlets(data []edgeproto.Cloudlet, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("CreateCloudlet %v\n", data[ii])
		myerr := CreateCloudlet(&data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var DeleteCloudletCmd = &cobra.Command{
	Use: "DeleteCloudlet",
	RunE: func(cmd *cobra.Command, args []string) error {
		// if we got this far, usage has been met.
		cmd.SilenceUsage = true
		err := parseCloudletEnums()
		if err != nil {
			return fmt.Errorf("DeleteCloudlet failed: %s", err.Error())
		}
		return DeleteCloudlet(&CloudletIn)
	},
}

func DeleteCloudlet(in *edgeproto.Cloudlet) error {
	if CloudletApiCmd == nil {
		return fmt.Errorf("CloudletApi client not initialized")
	}
	ctx := context.Background()
	stream, err := CloudletApiCmd.DeleteCloudlet(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("DeleteCloudlet failed: %s", errstr)
	}
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("DeleteCloudlet recv failed: %s", err.Error())
		}
		ResultWriteOutputOne(obj)
	}
	return nil
}

func DeleteCloudlets(data []edgeproto.Cloudlet, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("DeleteCloudlet %v\n", data[ii])
		myerr := DeleteCloudlet(&data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var UpdateCloudletCmd = &cobra.Command{
	Use: "UpdateCloudlet",
	RunE: func(cmd *cobra.Command, args []string) error {
		// if we got this far, usage has been met.
		cmd.SilenceUsage = true
		err := parseCloudletEnums()
		if err != nil {
			return fmt.Errorf("UpdateCloudlet failed: %s", err.Error())
		}
		CloudletSetFields()
		return UpdateCloudlet(&CloudletIn)
	},
}

func UpdateCloudlet(in *edgeproto.Cloudlet) error {
	if CloudletApiCmd == nil {
		return fmt.Errorf("CloudletApi client not initialized")
	}
	ctx := context.Background()
	stream, err := CloudletApiCmd.UpdateCloudlet(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("UpdateCloudlet failed: %s", errstr)
	}
	objs := make([]*edgeproto.Result, 0)
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("UpdateCloudlet recv failed: %s", err.Error())
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		return nil
	}
	ResultWriteOutputArray(objs)
	return nil
}

func UpdateCloudlets(data []edgeproto.Cloudlet, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("UpdateCloudlet %v\n", data[ii])
		myerr := UpdateCloudlet(&data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var ShowCloudletCmd = &cobra.Command{
	Use: "ShowCloudlet",
	RunE: func(cmd *cobra.Command, args []string) error {
		// if we got this far, usage has been met.
		cmd.SilenceUsage = true
		err := parseCloudletEnums()
		if err != nil {
			return fmt.Errorf("ShowCloudlet failed: %s", err.Error())
		}
		return ShowCloudlet(&CloudletIn)
	},
}

func ShowCloudlet(in *edgeproto.Cloudlet) error {
	if CloudletApiCmd == nil {
		return fmt.Errorf("CloudletApi client not initialized")
	}
	ctx := context.Background()
	stream, err := CloudletApiCmd.ShowCloudlet(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("ShowCloudlet failed: %s", errstr)
	}
	objs := make([]*edgeproto.Cloudlet, 0)
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("ShowCloudlet recv failed: %s", err.Error())
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		return nil
	}
	CloudletWriteOutputArray(objs)
	return nil
}

func ShowCloudlets(data []edgeproto.Cloudlet, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("ShowCloudlet %v\n", data[ii])
		myerr := ShowCloudlet(&data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var CloudletApiCmds = []*cobra.Command{
	CreateCloudletCmd,
	DeleteCloudletCmd,
	UpdateCloudletCmd,
	ShowCloudletCmd,
}

var ShowCloudletInfoCmd = &cobra.Command{
	Use: "ShowCloudletInfo",
	RunE: func(cmd *cobra.Command, args []string) error {
		// if we got this far, usage has been met.
		cmd.SilenceUsage = true
		err := parseCloudletInfoEnums()
		if err != nil {
			return fmt.Errorf("ShowCloudletInfo failed: %s", err.Error())
		}
		return ShowCloudletInfo(&CloudletInfoIn)
	},
}

func ShowCloudletInfo(in *edgeproto.CloudletInfo) error {
	if CloudletInfoApiCmd == nil {
		return fmt.Errorf("CloudletInfoApi client not initialized")
	}
	ctx := context.Background()
	stream, err := CloudletInfoApiCmd.ShowCloudletInfo(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("ShowCloudletInfo failed: %s", errstr)
	}
	objs := make([]*edgeproto.CloudletInfo, 0)
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("ShowCloudletInfo recv failed: %s", err.Error())
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		return nil
	}
	CloudletInfoWriteOutputArray(objs)
	return nil
}

func ShowCloudletInfos(data []edgeproto.CloudletInfo, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("ShowCloudletInfo %v\n", data[ii])
		myerr := ShowCloudletInfo(&data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var InjectCloudletInfoCmd = &cobra.Command{
	Use: "InjectCloudletInfo",
	RunE: func(cmd *cobra.Command, args []string) error {
		// if we got this far, usage has been met.
		cmd.SilenceUsage = true
		err := parseCloudletInfoEnums()
		if err != nil {
			return fmt.Errorf("InjectCloudletInfo failed: %s", err.Error())
		}
		return InjectCloudletInfo(&CloudletInfoIn)
	},
}

func InjectCloudletInfo(in *edgeproto.CloudletInfo) error {
	if CloudletInfoApiCmd == nil {
		return fmt.Errorf("CloudletInfoApi client not initialized")
	}
	ctx := context.Background()
	obj, err := CloudletInfoApiCmd.InjectCloudletInfo(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("InjectCloudletInfo failed: %s", errstr)
	}
	ResultWriteOutputOne(obj)
	return nil
}

func InjectCloudletInfos(data []edgeproto.CloudletInfo, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("InjectCloudletInfo %v\n", data[ii])
		myerr := InjectCloudletInfo(&data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var EvictCloudletInfoCmd = &cobra.Command{
	Use: "EvictCloudletInfo",
	RunE: func(cmd *cobra.Command, args []string) error {
		// if we got this far, usage has been met.
		cmd.SilenceUsage = true
		err := parseCloudletInfoEnums()
		if err != nil {
			return fmt.Errorf("EvictCloudletInfo failed: %s", err.Error())
		}
		return EvictCloudletInfo(&CloudletInfoIn)
	},
}

func EvictCloudletInfo(in *edgeproto.CloudletInfo) error {
	if CloudletInfoApiCmd == nil {
		return fmt.Errorf("CloudletInfoApi client not initialized")
	}
	ctx := context.Background()
	obj, err := CloudletInfoApiCmd.EvictCloudletInfo(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("EvictCloudletInfo failed: %s", errstr)
	}
	ResultWriteOutputOne(obj)
	return nil
}

func EvictCloudletInfos(data []edgeproto.CloudletInfo, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("EvictCloudletInfo %v\n", data[ii])
		myerr := EvictCloudletInfo(&data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var CloudletInfoApiCmds = []*cobra.Command{
	ShowCloudletInfoCmd,
	InjectCloudletInfoCmd,
	EvictCloudletInfoCmd,
}

var ShowCloudletMetricsCmd = &cobra.Command{
	Use: "ShowCloudletMetrics",
	RunE: func(cmd *cobra.Command, args []string) error {
		// if we got this far, usage has been met.
		cmd.SilenceUsage = true
		return ShowCloudletMetrics(&CloudletMetricsIn)
	},
}

func ShowCloudletMetrics(in *edgeproto.CloudletMetrics) error {
	if CloudletMetricsApiCmd == nil {
		return fmt.Errorf("CloudletMetricsApi client not initialized")
	}
	ctx := context.Background()
	stream, err := CloudletMetricsApiCmd.ShowCloudletMetrics(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("ShowCloudletMetrics failed: %s", errstr)
	}
	objs := make([]*edgeproto.CloudletMetrics, 0)
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("ShowCloudletMetrics recv failed: %s", err.Error())
		}
		objs = append(objs, obj)
	}
	if len(objs) == 0 {
		return nil
	}
	CloudletMetricsWriteOutputArray(objs)
	return nil
}

func ShowCloudletMetricss(data []edgeproto.CloudletMetrics, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("ShowCloudletMetrics %v\n", data[ii])
		myerr := ShowCloudletMetrics(&data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var CloudletMetricsApiCmds = []*cobra.Command{
	ShowCloudletMetricsCmd,
}

func init() {
	CloudletFlagSet.StringVar(&CloudletIn.Key.OperatorKey.Name, "key-operatorkey-name", "", "Key.OperatorKey.Name")
	CloudletFlagSet.StringVar(&CloudletIn.Key.Name, "key-name", "", "Key.Name")
	CloudletFlagSet.StringVar(&CloudletIn.AccessUri, "accessuri", "", "AccessUri")
	CloudletFlagSet.Float64Var(&CloudletIn.Location.Latitude, "location-latitude", 0, "Location.Latitude")
	CloudletFlagSet.Float64Var(&CloudletIn.Location.Longitude, "location-longitude", 0, "Location.Longitude")
	CloudletNoConfigFlagSet.Float64Var(&CloudletIn.Location.HorizontalAccuracy, "location-horizontalaccuracy", 0, "Location.HorizontalAccuracy")
	CloudletNoConfigFlagSet.Float64Var(&CloudletIn.Location.VerticalAccuracy, "location-verticalaccuracy", 0, "Location.VerticalAccuracy")
	CloudletFlagSet.Float64Var(&CloudletIn.Location.Altitude, "location-altitude", 0, "Location.Altitude")
	CloudletNoConfigFlagSet.Float64Var(&CloudletIn.Location.Course, "location-course", 0, "Location.Course")
	CloudletNoConfigFlagSet.Float64Var(&CloudletIn.Location.Speed, "location-speed", 0, "Location.Speed")
	CloudletIn.Location.Timestamp = &distributed_match_engine.Timestamp{}
	CloudletNoConfigFlagSet.Int64Var(&CloudletIn.Location.Timestamp.Seconds, "location-timestamp-seconds", 0, "Location.Timestamp.Seconds")
	CloudletNoConfigFlagSet.Int32Var(&CloudletIn.Location.Timestamp.Nanos, "location-timestamp-nanos", 0, "Location.Timestamp.Nanos")
	CloudletFlagSet.StringVar(&CloudletInIpSupport, "ipsupport", "", "one of [IpSupportUnknown IpSupportStatic IpSupportDynamic]")
	CloudletFlagSet.StringVar(&CloudletIn.StaticIps, "staticips", "", "StaticIps")
	CloudletFlagSet.Int32Var(&CloudletIn.NumDynamicIps, "numdynamicips", 0, "NumDynamicIps")
	CloudletInfoFlagSet.StringVar(&CloudletInfoIn.Key.OperatorKey.Name, "key-operatorkey-name", "", "Key.OperatorKey.Name")
	CloudletInfoFlagSet.StringVar(&CloudletInfoIn.Key.Name, "key-name", "", "Key.Name")
	CloudletInfoFlagSet.StringVar(&CloudletInfoInState, "state", "", "one of [CloudletStateUnknown CloudletStateErrors CloudletStateReady CloudletStateOffline CloudletStateNotPresent]")
	CloudletInfoFlagSet.Int64Var(&CloudletInfoIn.NotifyId, "notifyid", 0, "NotifyId")
	CloudletInfoFlagSet.StringVar(&CloudletInfoIn.Controller, "controller", "", "Controller")
	CloudletInfoFlagSet.Uint64Var(&CloudletInfoIn.OsMaxRam, "osmaxram", 0, "OsMaxRam")
	CloudletInfoFlagSet.Uint64Var(&CloudletInfoIn.OsMaxVcores, "osmaxvcores", 0, "OsMaxVcores")
	CloudletInfoFlagSet.Uint64Var(&CloudletInfoIn.OsMaxVolGb, "osmaxvolgb", 0, "OsMaxVolGb")
	CloudletMetricsFlagSet.Uint64Var(&CloudletMetricsIn.Foo, "foo", 0, "Foo")
	CreateCloudletCmd.Flags().AddFlagSet(CloudletFlagSet)
	DeleteCloudletCmd.Flags().AddFlagSet(CloudletFlagSet)
	UpdateCloudletCmd.Flags().AddFlagSet(CloudletFlagSet)
	ShowCloudletCmd.Flags().AddFlagSet(CloudletFlagSet)
	ShowCloudletInfoCmd.Flags().AddFlagSet(CloudletInfoFlagSet)
	InjectCloudletInfoCmd.Flags().AddFlagSet(CloudletInfoFlagSet)
	EvictCloudletInfoCmd.Flags().AddFlagSet(CloudletInfoFlagSet)
	ShowCloudletMetricsCmd.Flags().AddFlagSet(CloudletMetricsFlagSet)
}

func CloudletApiAllowNoConfig() {
	CreateCloudletCmd.Flags().AddFlagSet(CloudletNoConfigFlagSet)
	DeleteCloudletCmd.Flags().AddFlagSet(CloudletNoConfigFlagSet)
	UpdateCloudletCmd.Flags().AddFlagSet(CloudletNoConfigFlagSet)
	ShowCloudletCmd.Flags().AddFlagSet(CloudletNoConfigFlagSet)
}

func CloudletInfoApiAllowNoConfig() {
	ShowCloudletInfoCmd.Flags().AddFlagSet(CloudletInfoNoConfigFlagSet)
	InjectCloudletInfoCmd.Flags().AddFlagSet(CloudletInfoNoConfigFlagSet)
	EvictCloudletInfoCmd.Flags().AddFlagSet(CloudletInfoNoConfigFlagSet)
}

func CloudletMetricsApiAllowNoConfig() {
	ShowCloudletMetricsCmd.Flags().AddFlagSet(CloudletMetricsNoConfigFlagSet)
}

func CloudletSetFields() {
	CloudletIn.Fields = make([]string, 0)
	if CloudletFlagSet.Lookup("key-operatorkey-name").Changed {
		CloudletIn.Fields = append(CloudletIn.Fields, "2.1.1")
	}
	if CloudletFlagSet.Lookup("key-name").Changed {
		CloudletIn.Fields = append(CloudletIn.Fields, "2.2")
	}
	if CloudletFlagSet.Lookup("accessuri").Changed {
		CloudletIn.Fields = append(CloudletIn.Fields, "4")
	}
	if CloudletFlagSet.Lookup("location-latitude").Changed {
		CloudletIn.Fields = append(CloudletIn.Fields, "5.1")
	}
	if CloudletFlagSet.Lookup("location-longitude").Changed {
		CloudletIn.Fields = append(CloudletIn.Fields, "5.2")
	}
	if CloudletNoConfigFlagSet.Lookup("location-horizontalaccuracy").Changed {
		CloudletIn.Fields = append(CloudletIn.Fields, "5.3")
	}
	if CloudletNoConfigFlagSet.Lookup("location-verticalaccuracy").Changed {
		CloudletIn.Fields = append(CloudletIn.Fields, "5.4")
	}
	if CloudletFlagSet.Lookup("location-altitude").Changed {
		CloudletIn.Fields = append(CloudletIn.Fields, "5.5")
	}
	if CloudletNoConfigFlagSet.Lookup("location-course").Changed {
		CloudletIn.Fields = append(CloudletIn.Fields, "5.6")
	}
	if CloudletNoConfigFlagSet.Lookup("location-speed").Changed {
		CloudletIn.Fields = append(CloudletIn.Fields, "5.7")
	}
	if CloudletNoConfigFlagSet.Lookup("location-timestamp-seconds").Changed {
		CloudletIn.Fields = append(CloudletIn.Fields, "5.8.1")
	}
	if CloudletNoConfigFlagSet.Lookup("location-timestamp-nanos").Changed {
		CloudletIn.Fields = append(CloudletIn.Fields, "5.8.2")
	}
	if CloudletFlagSet.Lookup("ipsupport").Changed {
		CloudletIn.Fields = append(CloudletIn.Fields, "6")
	}
	if CloudletFlagSet.Lookup("staticips").Changed {
		CloudletIn.Fields = append(CloudletIn.Fields, "7")
	}
	if CloudletFlagSet.Lookup("numdynamicips").Changed {
		CloudletIn.Fields = append(CloudletIn.Fields, "8")
	}
}

func CloudletInfoSetFields() {
	CloudletInfoIn.Fields = make([]string, 0)
	if CloudletInfoFlagSet.Lookup("key-operatorkey-name").Changed {
		CloudletInfoIn.Fields = append(CloudletInfoIn.Fields, "2.1.1")
	}
	if CloudletInfoFlagSet.Lookup("key-name").Changed {
		CloudletInfoIn.Fields = append(CloudletInfoIn.Fields, "2.2")
	}
	if CloudletInfoFlagSet.Lookup("state").Changed {
		CloudletInfoIn.Fields = append(CloudletInfoIn.Fields, "3")
	}
	if CloudletInfoFlagSet.Lookup("notifyid").Changed {
		CloudletInfoIn.Fields = append(CloudletInfoIn.Fields, "4")
	}
	if CloudletInfoFlagSet.Lookup("controller").Changed {
		CloudletInfoIn.Fields = append(CloudletInfoIn.Fields, "5")
	}
	if CloudletInfoFlagSet.Lookup("osmaxram").Changed {
		CloudletInfoIn.Fields = append(CloudletInfoIn.Fields, "6")
	}
	if CloudletInfoFlagSet.Lookup("osmaxvcores").Changed {
		CloudletInfoIn.Fields = append(CloudletInfoIn.Fields, "7")
	}
	if CloudletInfoFlagSet.Lookup("osmaxvolgb").Changed {
		CloudletInfoIn.Fields = append(CloudletInfoIn.Fields, "8")
	}
}

func parseCloudletEnums() error {
	if CloudletInIpSupport != "" {
		switch CloudletInIpSupport {
		case "IpSupportUnknown":
			CloudletIn.IpSupport = edgeproto.IpSupport(0)
		case "IpSupportStatic":
			CloudletIn.IpSupport = edgeproto.IpSupport(1)
		case "IpSupportDynamic":
			CloudletIn.IpSupport = edgeproto.IpSupport(2)
		default:
			return errors.New("Invalid value for CloudletInIpSupport")
		}
	}
	return nil
}

func parseCloudletInfoEnums() error {
	if CloudletInfoInState != "" {
		switch CloudletInfoInState {
		case "CloudletStateUnknown":
			CloudletInfoIn.State = edgeproto.CloudletState(0)
		case "CloudletStateErrors":
			CloudletInfoIn.State = edgeproto.CloudletState(1)
		case "CloudletStateReady":
			CloudletInfoIn.State = edgeproto.CloudletState(2)
		case "CloudletStateOffline":
			CloudletInfoIn.State = edgeproto.CloudletState(3)
		case "CloudletStateNotPresent":
			CloudletInfoIn.State = edgeproto.CloudletState(4)
		default:
			return errors.New("Invalid value for CloudletInfoInState")
		}
	}
	return nil
}
