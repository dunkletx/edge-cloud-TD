// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: operator.proto

package gencmd

import edgeproto "github.com/mobiledgex/edge-cloud/edgeproto"
import "strings"
import "github.com/spf13/cobra"
import "context"
import "os"
import "io"
import "text/tabwriter"
import "github.com/spf13/pflag"
import "github.com/mobiledgex/edge-cloud/protoc-gen-cmd/cmdsup"
import "google.golang.org/grpc/status"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/mobiledgex/edge-cloud/protogen"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT
var OperatorApiCmd edgeproto.OperatorApiClient
var OperatorIn edgeproto.Operator
var OperatorFlagSet = pflag.NewFlagSet("Operator", pflag.ExitOnError)
var OperatorNoConfigFlagSet = pflag.NewFlagSet("OperatorNoConfig", pflag.ExitOnError)

func OperatorKeySlicer(in *edgeproto.OperatorKey) []string {
	s := make([]string, 0, 1)
	s = append(s, in.Name)
	return s
}

func OperatorKeyHeaderSlicer() []string {
	s := make([]string, 0, 1)
	s = append(s, "Name")
	return s
}

func OperatorKeyWriteOutputArray(objs []*edgeproto.OperatorKey) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(OperatorKeyHeaderSlicer(), "\t"))
		for _, obj := range objs {
			fmt.Fprintln(output, strings.Join(OperatorKeySlicer(obj), "\t"))
		}
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(objs)
	}
}

func OperatorKeyWriteOutputOne(obj *edgeproto.OperatorKey) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(OperatorKeyHeaderSlicer(), "\t"))
		fmt.Fprintln(output, strings.Join(OperatorKeySlicer(obj), "\t"))
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(obj)
	}
}
func OperatorSlicer(in *edgeproto.Operator) []string {
	s := make([]string, 0, 2)
	if in.Fields == nil {
		in.Fields = make([]string, 1)
	}
	s = append(s, in.Fields[0])
	s = append(s, in.Key.Name)
	return s
}

func OperatorHeaderSlicer() []string {
	s := make([]string, 0, 2)
	s = append(s, "Fields")
	s = append(s, "Key-Name")
	return s
}

func OperatorWriteOutputArray(objs []*edgeproto.Operator) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(OperatorHeaderSlicer(), "\t"))
		for _, obj := range objs {
			fmt.Fprintln(output, strings.Join(OperatorSlicer(obj), "\t"))
		}
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(objs)
	}
}

func OperatorWriteOutputOne(obj *edgeproto.Operator) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(OperatorHeaderSlicer(), "\t"))
		fmt.Fprintln(output, strings.Join(OperatorSlicer(obj), "\t"))
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(obj)
	}
}

var CreateOperatorCmd = &cobra.Command{
	Use: "CreateOperator",
	RunE: func(cmd *cobra.Command, args []string) error {
		// if we got this far, usage has been met.
		cmd.SilenceUsage = true
		if OperatorApiCmd == nil {
			return fmt.Errorf("OperatorApi client not initialized")
		}
		var err error
		ctx := context.Background()
		obj, err := OperatorApiCmd.CreateOperator(ctx, &OperatorIn)
		if err != nil {
			errstr := err.Error()
			st, ok := status.FromError(err)
			if ok {
				errstr = st.Message()
			}
			return fmt.Errorf("CreateOperator failed: %s", errstr)
		}
		ResultWriteOutputOne(obj)
		return nil
	},
}

var DeleteOperatorCmd = &cobra.Command{
	Use: "DeleteOperator",
	RunE: func(cmd *cobra.Command, args []string) error {
		// if we got this far, usage has been met.
		cmd.SilenceUsage = true
		if OperatorApiCmd == nil {
			return fmt.Errorf("OperatorApi client not initialized")
		}
		var err error
		ctx := context.Background()
		obj, err := OperatorApiCmd.DeleteOperator(ctx, &OperatorIn)
		if err != nil {
			errstr := err.Error()
			st, ok := status.FromError(err)
			if ok {
				errstr = st.Message()
			}
			return fmt.Errorf("DeleteOperator failed: %s", errstr)
		}
		ResultWriteOutputOne(obj)
		return nil
	},
}

var UpdateOperatorCmd = &cobra.Command{
	Use: "UpdateOperator",
	RunE: func(cmd *cobra.Command, args []string) error {
		// if we got this far, usage has been met.
		cmd.SilenceUsage = true
		if OperatorApiCmd == nil {
			return fmt.Errorf("OperatorApi client not initialized")
		}
		var err error
		OperatorSetFields()
		ctx := context.Background()
		obj, err := OperatorApiCmd.UpdateOperator(ctx, &OperatorIn)
		if err != nil {
			errstr := err.Error()
			st, ok := status.FromError(err)
			if ok {
				errstr = st.Message()
			}
			return fmt.Errorf("UpdateOperator failed: %s", errstr)
		}
		ResultWriteOutputOne(obj)
		return nil
	},
}

var ShowOperatorCmd = &cobra.Command{
	Use: "ShowOperator",
	RunE: func(cmd *cobra.Command, args []string) error {
		// if we got this far, usage has been met.
		cmd.SilenceUsage = true
		if OperatorApiCmd == nil {
			return fmt.Errorf("OperatorApi client not initialized")
		}
		var err error
		ctx := context.Background()
		stream, err := OperatorApiCmd.ShowOperator(ctx, &OperatorIn)
		if err != nil {
			errstr := err.Error()
			st, ok := status.FromError(err)
			if ok {
				errstr = st.Message()
			}
			return fmt.Errorf("ShowOperator failed: %s", errstr)
		}
		objs := make([]*edgeproto.Operator, 0)
		for {
			obj, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				return fmt.Errorf("ShowOperator recv failed: %s", err.Error())
			}
			objs = append(objs, obj)
		}
		if len(objs) == 0 {
			return nil
		}
		OperatorWriteOutputArray(objs)
		return nil
	},
}

var OperatorApiCmds = []*cobra.Command{
	CreateOperatorCmd,
	DeleteOperatorCmd,
	UpdateOperatorCmd,
	ShowOperatorCmd,
}

func init() {
	OperatorFlagSet.StringVar(&OperatorIn.Key.Name, "key-name", "", "Key.Name")
	CreateOperatorCmd.Flags().AddFlagSet(OperatorFlagSet)
	DeleteOperatorCmd.Flags().AddFlagSet(OperatorFlagSet)
	UpdateOperatorCmd.Flags().AddFlagSet(OperatorFlagSet)
	ShowOperatorCmd.Flags().AddFlagSet(OperatorFlagSet)
}

func OperatorApiAllowNoConfig() {
	CreateOperatorCmd.Flags().AddFlagSet(OperatorNoConfigFlagSet)
	DeleteOperatorCmd.Flags().AddFlagSet(OperatorNoConfigFlagSet)
	UpdateOperatorCmd.Flags().AddFlagSet(OperatorNoConfigFlagSet)
	ShowOperatorCmd.Flags().AddFlagSet(OperatorNoConfigFlagSet)
}

func OperatorSetFields() {
	OperatorIn.Fields = make([]string, 0)
	if OperatorFlagSet.Lookup("key-name").Changed {
		OperatorIn.Fields = append(OperatorIn.Fields, "2.1")
	}
}
