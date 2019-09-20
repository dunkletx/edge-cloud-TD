// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: operator.proto

package gencmd

import edgeproto "github.com/mobiledgex/edge-cloud/edgeproto"
import "strings"
import "github.com/spf13/cobra"
import "context"
import "io"
import "github.com/mobiledgex/edge-cloud/cli"
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

var CreateOperatorCmd = &cli.Command{
	Use:          "CreateOperator",
	RequiredArgs: strings.Join(OperatorRequiredArgs, " "),
	OptionalArgs: strings.Join(OperatorOptionalArgs, " "),
	AliasArgs:    strings.Join(OperatorAliasArgs, " "),
	SpecialArgs:  &OperatorSpecialArgs,
	Comments:     OperatorComments,
	ReqData:      &edgeproto.Operator{},
	ReplyData:    &edgeproto.Result{},
	Run:          runCreateOperator,
}

func runCreateOperator(c *cli.Command, args []string) error {
	obj := c.ReqData.(*edgeproto.Operator)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return CreateOperator(c, obj)
}

func CreateOperator(c *cli.Command, in *edgeproto.Operator) error {
	if OperatorApiCmd == nil {
		return fmt.Errorf("OperatorApi client not initialized")
	}
	ctx := context.Background()
	obj, err := OperatorApiCmd.CreateOperator(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("CreateOperator failed: %s", errstr)
	}
	c.WriteOutput(obj, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func CreateOperators(c *cli.Command, data []edgeproto.Operator, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("CreateOperator %v\n", data[ii])
		myerr := CreateOperator(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var DeleteOperatorCmd = &cli.Command{
	Use:          "DeleteOperator",
	RequiredArgs: strings.Join(OperatorRequiredArgs, " "),
	OptionalArgs: strings.Join(OperatorOptionalArgs, " "),
	AliasArgs:    strings.Join(OperatorAliasArgs, " "),
	SpecialArgs:  &OperatorSpecialArgs,
	Comments:     OperatorComments,
	ReqData:      &edgeproto.Operator{},
	ReplyData:    &edgeproto.Result{},
	Run:          runDeleteOperator,
}

func runDeleteOperator(c *cli.Command, args []string) error {
	obj := c.ReqData.(*edgeproto.Operator)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return DeleteOperator(c, obj)
}

func DeleteOperator(c *cli.Command, in *edgeproto.Operator) error {
	if OperatorApiCmd == nil {
		return fmt.Errorf("OperatorApi client not initialized")
	}
	ctx := context.Background()
	obj, err := OperatorApiCmd.DeleteOperator(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("DeleteOperator failed: %s", errstr)
	}
	c.WriteOutput(obj, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func DeleteOperators(c *cli.Command, data []edgeproto.Operator, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("DeleteOperator %v\n", data[ii])
		myerr := DeleteOperator(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var UpdateOperatorCmd = &cli.Command{
	Use:          "UpdateOperator",
	RequiredArgs: strings.Join(OperatorRequiredArgs, " "),
	OptionalArgs: strings.Join(OperatorOptionalArgs, " "),
	AliasArgs:    strings.Join(OperatorAliasArgs, " "),
	SpecialArgs:  &OperatorSpecialArgs,
	Comments:     OperatorComments,
	ReqData:      &edgeproto.Operator{},
	ReplyData:    &edgeproto.Result{},
	Run:          runUpdateOperator,
}

func runUpdateOperator(c *cli.Command, args []string) error {
	obj := c.ReqData.(*edgeproto.Operator)
	jsonMap, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	obj.Fields = cli.GetSpecifiedFields(jsonMap, c.ReqData, cli.JsonNamespace)
	return UpdateOperator(c, obj)
}

func UpdateOperator(c *cli.Command, in *edgeproto.Operator) error {
	if OperatorApiCmd == nil {
		return fmt.Errorf("OperatorApi client not initialized")
	}
	ctx := context.Background()
	obj, err := OperatorApiCmd.UpdateOperator(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("UpdateOperator failed: %s", errstr)
	}
	c.WriteOutput(obj, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func UpdateOperators(c *cli.Command, data []edgeproto.Operator, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("UpdateOperator %v\n", data[ii])
		myerr := UpdateOperator(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var ShowOperatorCmd = &cli.Command{
	Use:          "ShowOperator",
	OptionalArgs: strings.Join(append(OperatorRequiredArgs, OperatorOptionalArgs...), " "),
	AliasArgs:    strings.Join(OperatorAliasArgs, " "),
	SpecialArgs:  &OperatorSpecialArgs,
	Comments:     OperatorComments,
	ReqData:      &edgeproto.Operator{},
	ReplyData:    &edgeproto.Operator{},
	Run:          runShowOperator,
}

func runShowOperator(c *cli.Command, args []string) error {
	obj := c.ReqData.(*edgeproto.Operator)
	_, err := c.ParseInput(args)
	if err != nil {
		return err
	}
	return ShowOperator(c, obj)
}

func ShowOperator(c *cli.Command, in *edgeproto.Operator) error {
	if OperatorApiCmd == nil {
		return fmt.Errorf("OperatorApi client not initialized")
	}
	ctx := context.Background()
	stream, err := OperatorApiCmd.ShowOperator(ctx, in)
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
	c.WriteOutput(objs, cli.OutputFormat)
	return nil
}

// this supports "Create" and "Delete" commands on ApplicationData
func ShowOperators(c *cli.Command, data []edgeproto.Operator, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("ShowOperator %v\n", data[ii])
		myerr := ShowOperator(c, &data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var OperatorApiCmds = []*cobra.Command{
	CreateOperatorCmd.GenCmd(),
	DeleteOperatorCmd.GenCmd(),
	UpdateOperatorCmd.GenCmd(),
	ShowOperatorCmd.GenCmd(),
}

var OperatorKeyRequiredArgs = []string{}
var OperatorKeyOptionalArgs = []string{
	"name",
}
var OperatorKeyAliasArgs = []string{}
var OperatorKeyComments = map[string]string{
	"name": "Company or Organization name of the operator",
}
var OperatorKeySpecialArgs = map[string]string{}
var OperatorRequiredArgs = []string{
	"name",
}
var OperatorOptionalArgs = []string{}
var OperatorAliasArgs = []string{
	"name=key.name",
}
var OperatorComments = map[string]string{
	"name": "Company or Organization name of the operator",
}
var OperatorSpecialArgs = map[string]string{}
