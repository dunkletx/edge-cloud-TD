// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sample.proto

/*
Package gencmd is a generated protocol buffer package.

It is generated from these files:
	sample.proto

It has these top-level messages:
	NestedMessage
	IncludeMessage
	IncludeFields
	TestGen
*/
package gencmd

import distributed_match_engine "github.com/mobiledgex/edge-cloud/d-match-engine/dme-proto"
import testgen "github.com/mobiledgex/edge-cloud/testgen"
import "strings"
import "strconv"
import "github.com/spf13/cobra"
import "context"
import "os"
import "text/tabwriter"
import "github.com/spf13/pflag"
import "errors"
import "github.com/mobiledgex/edge-cloud/protoc-gen-cmd/cmdsup"
import "google.golang.org/grpc/status"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/mobiledgex/edge-cloud/d-match-engine/dme-proto"
import _ "github.com/mobiledgex/edge-cloud/protogen"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT
var TestApiCmd testgen.TestApiClient
var TestGenIn testgen.TestGen
var TestGenFlagSet = pflag.NewFlagSet("TestGen", pflag.ExitOnError)
var TestGenNoConfigFlagSet = pflag.NewFlagSet("TestGenNoConfig", pflag.ExitOnError)
var TestGenInOuterEn string
var TestGenInInnerEn string
var OuterEnumStrings = []string{
	"OUTER0",
	"OUTER1",
	"OUTER2",
	"OUTER3",
}

var InnerEnumStrings = []string{
	"INNER0",
	"INNER1",
	"INNER2",
	"INNER3",
}

func NestedMessageSlicer(in *testgen.NestedMessage) []string {
	s := make([]string, 0, 1)
	s = append(s, in.Name)
	return s
}

func NestedMessageHeaderSlicer() []string {
	s := make([]string, 0, 1)
	s = append(s, "Name")
	return s
}

func NestedMessageWriteOutputArray(objs []*testgen.NestedMessage) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(NestedMessageHeaderSlicer(), "\t"))
		for _, obj := range objs {
			fmt.Fprintln(output, strings.Join(NestedMessageSlicer(obj), "\t"))
		}
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(objs)
	}
}

func NestedMessageWriteOutputOne(obj *testgen.NestedMessage) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(NestedMessageHeaderSlicer(), "\t"))
		fmt.Fprintln(output, strings.Join(NestedMessageSlicer(obj), "\t"))
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(obj)
	}
}
func IncludeMessageSlicer(in *testgen.IncludeMessage) []string {
	s := make([]string, 0, 3)
	s = append(s, in.Name)
	s = append(s, strconv.FormatUint(uint64(in.Id), 10))
	if in.NestedMsg == nil {
		in.NestedMsg = &testgen.NestedMessage{}
	}
	s = append(s, in.NestedMsg.Name)
	return s
}

func IncludeMessageHeaderSlicer() []string {
	s := make([]string, 0, 3)
	s = append(s, "Name")
	s = append(s, "Id")
	s = append(s, "NestedMsg-Name")
	return s
}

func IncludeMessageWriteOutputArray(objs []*testgen.IncludeMessage) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(IncludeMessageHeaderSlicer(), "\t"))
		for _, obj := range objs {
			fmt.Fprintln(output, strings.Join(IncludeMessageSlicer(obj), "\t"))
		}
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(objs)
	}
}

func IncludeMessageWriteOutputOne(obj *testgen.IncludeMessage) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(IncludeMessageHeaderSlicer(), "\t"))
		fmt.Fprintln(output, strings.Join(IncludeMessageSlicer(obj), "\t"))
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(obj)
	}
}
func IncludeFieldsSlicer(in *testgen.IncludeFields) []string {
	s := make([]string, 0, 2)
	s = append(s, "")
	for _, b := range in.Fields {
		s[len(s)-1] += fmt.Sprintf("%v", b)
	}
	s = append(s, in.Name)
	return s
}

func IncludeFieldsHeaderSlicer() []string {
	s := make([]string, 0, 2)
	s = append(s, "Fields")
	s = append(s, "Name")
	return s
}

func IncludeFieldsWriteOutputArray(objs []*testgen.IncludeFields) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(IncludeFieldsHeaderSlicer(), "\t"))
		for _, obj := range objs {
			fmt.Fprintln(output, strings.Join(IncludeFieldsSlicer(obj), "\t"))
		}
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(objs)
	}
}

func IncludeFieldsWriteOutputOne(obj *testgen.IncludeFields) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(IncludeFieldsHeaderSlicer(), "\t"))
		fmt.Fprintln(output, strings.Join(IncludeFieldsSlicer(obj), "\t"))
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(obj)
	}
}
func TestGenSlicer(in *testgen.TestGen) []string {
	s := make([]string, 0, 38)
	if in.Fields == nil {
		in.Fields = make([]string, 1)
	}
	s = append(s, in.Fields[0])
	s = append(s, in.Name)
	s = append(s, strconv.FormatFloat(float64(in.Db), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.Fl), 'e', -1, 32))
	s = append(s, strconv.FormatUint(uint64(in.I32), 10))
	s = append(s, strconv.FormatUint(uint64(in.I64), 10))
	s = append(s, strconv.FormatUint(uint64(in.U32), 10))
	s = append(s, strconv.FormatUint(uint64(in.U64), 10))
	s = append(s, strconv.FormatUint(uint64(in.S32), 10))
	s = append(s, strconv.FormatUint(uint64(in.S64), 10))
	s = append(s, strconv.FormatUint(uint64(in.F32), 10))
	s = append(s, strconv.FormatUint(uint64(in.F64), 10))
	s = append(s, strconv.FormatUint(uint64(in.Sf32), 10))
	s = append(s, strconv.FormatUint(uint64(in.Sf64), 10))
	s = append(s, strconv.FormatBool(in.Bb))
	s = append(s, testgen.OuterEnum_name[int32(in.OuterEn)])
	s = append(s, testgen.TestGen_InnerEnum_name[int32(in.InnerEn)])
	if in.InnerMsg == nil {
		in.InnerMsg = &testgen.TestGen_InnerMessage{}
	}
	s = append(s, in.InnerMsg.Url)
	s = append(s, strconv.FormatUint(uint64(in.InnerMsg.Id), 10))
	s = append(s, in.InnerMsgNonnull.Url)
	s = append(s, strconv.FormatUint(uint64(in.InnerMsgNonnull.Id), 10))
	if in.IncludeMsg == nil {
		in.IncludeMsg = &testgen.IncludeMessage{}
	}
	s = append(s, in.IncludeMsg.Name)
	s = append(s, strconv.FormatUint(uint64(in.IncludeMsg.Id), 10))
	if in.IncludeMsg.NestedMsg == nil {
		in.IncludeMsg.NestedMsg = &testgen.NestedMessage{}
	}
	s = append(s, in.IncludeMsg.NestedMsg.Name)
	s = append(s, in.IncludeMsgNonnull.Name)
	s = append(s, strconv.FormatUint(uint64(in.IncludeMsgNonnull.Id), 10))
	if in.IncludeMsgNonnull.NestedMsg == nil {
		in.IncludeMsgNonnull.NestedMsg = &testgen.NestedMessage{}
	}
	s = append(s, in.IncludeMsgNonnull.NestedMsg.Name)
	if in.IncludeFields == nil {
		in.IncludeFields = &testgen.IncludeFields{}
	}
	s = append(s, "")
	for _, b := range in.IncludeFields.Fields {
		s[len(s)-1] += fmt.Sprintf("%v", b)
	}
	s = append(s, in.IncludeFields.Name)
	s = append(s, "")
	for _, b := range in.IncludeFieldsNonnull.Fields {
		s[len(s)-1] += fmt.Sprintf("%v", b)
	}
	s = append(s, in.IncludeFieldsNonnull.Name)
	if in.Loc == nil {
		in.Loc = &distributed_match_engine.Loc{}
	}
	s = append(s, strconv.FormatFloat(float64(in.Loc.Lat), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.Loc.Long), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.Loc.HorizontalAccuracy), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.Loc.VerticalAccuracy), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.Loc.Altitude), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.Loc.Course), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.Loc.Speed), 'e', -1, 32))
	if in.Loc.Timestamp == nil {
		in.Loc.Timestamp = &distributed_match_engine.Timestamp{}
	}
	s = append(s, strconv.FormatUint(uint64(in.Loc.Timestamp.Seconds), 10))
	s = append(s, strconv.FormatUint(uint64(in.Loc.Timestamp.Nanos), 10))
	s = append(s, strconv.FormatFloat(float64(in.LocNonnull.Lat), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.LocNonnull.Long), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.LocNonnull.HorizontalAccuracy), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.LocNonnull.VerticalAccuracy), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.LocNonnull.Altitude), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.LocNonnull.Course), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.LocNonnull.Speed), 'e', -1, 32))
	if in.LocNonnull.Timestamp == nil {
		in.LocNonnull.Timestamp = &distributed_match_engine.Timestamp{}
	}
	s = append(s, strconv.FormatUint(uint64(in.LocNonnull.Timestamp.Seconds), 10))
	s = append(s, strconv.FormatUint(uint64(in.LocNonnull.Timestamp.Nanos), 10))
	if in.RepeatedInt == nil {
		in.RepeatedInt = make([]int64, 1)
	}
	s = append(s, strconv.FormatUint(uint64(in.RepeatedInt[0]), 10))
	s = append(s, "")
	for i, b := range in.Ip {
		s[len(s)-1] += fmt.Sprintf("%v", b)
		if i < 3 {
			s[len(s)-1] += "."
		}
	}
	if in.Names == nil {
		in.Names = make([]string, 1)
	}
	s = append(s, in.Names[0])
	if in.RepeatedMsg == nil {
		in.RepeatedMsg = make([]*testgen.IncludeMessage, 1)
	}
	if in.RepeatedMsg[0] == nil {
		in.RepeatedMsg[0] = &testgen.IncludeMessage{}
	}
	s = append(s, in.RepeatedMsg[0].Name)
	s = append(s, strconv.FormatUint(uint64(in.RepeatedMsg[0].Id), 10))
	if in.RepeatedMsg[0].NestedMsg == nil {
		in.RepeatedMsg[0].NestedMsg = &testgen.NestedMessage{}
	}
	s = append(s, in.RepeatedMsg[0].NestedMsg.Name)
	if in.RepeatedMsgNonnull == nil {
		in.RepeatedMsgNonnull = make([]testgen.IncludeMessage, 1)
	}
	s = append(s, in.RepeatedMsgNonnull[0].Name)
	s = append(s, strconv.FormatUint(uint64(in.RepeatedMsgNonnull[0].Id), 10))
	if in.RepeatedMsgNonnull[0].NestedMsg == nil {
		in.RepeatedMsgNonnull[0].NestedMsg = &testgen.NestedMessage{}
	}
	s = append(s, in.RepeatedMsgNonnull[0].NestedMsg.Name)
	if in.RepeatedFields == nil {
		in.RepeatedFields = make([]*testgen.IncludeFields, 1)
	}
	if in.RepeatedFields[0] == nil {
		in.RepeatedFields[0] = &testgen.IncludeFields{}
	}
	s = append(s, "")
	for _, b := range in.RepeatedFields[0].Fields {
		s[len(s)-1] += fmt.Sprintf("%v", b)
	}
	s = append(s, in.RepeatedFields[0].Name)
	if in.RepeatedFieldsNonnull == nil {
		in.RepeatedFieldsNonnull = make([]testgen.IncludeFields, 1)
	}
	s = append(s, "")
	for _, b := range in.RepeatedFieldsNonnull[0].Fields {
		s[len(s)-1] += fmt.Sprintf("%v", b)
	}
	s = append(s, in.RepeatedFieldsNonnull[0].Name)
	if in.RepeatedInnerMsg == nil {
		in.RepeatedInnerMsg = make([]*testgen.TestGen_InnerMessage, 1)
	}
	if in.RepeatedInnerMsg[0] == nil {
		in.RepeatedInnerMsg[0] = &testgen.TestGen_InnerMessage{}
	}
	s = append(s, in.RepeatedInnerMsg[0].Url)
	s = append(s, strconv.FormatUint(uint64(in.RepeatedInnerMsg[0].Id), 10))
	if in.RepeatedInnerMsgNonnull == nil {
		in.RepeatedInnerMsgNonnull = make([]testgen.TestGen_InnerMessage, 1)
	}
	s = append(s, in.RepeatedInnerMsgNonnull[0].Url)
	s = append(s, strconv.FormatUint(uint64(in.RepeatedInnerMsgNonnull[0].Id), 10))
	if in.RepeatedLoc == nil {
		in.RepeatedLoc = make([]*distributed_match_engine.Loc, 1)
	}
	if in.RepeatedLoc[0] == nil {
		in.RepeatedLoc[0] = &distributed_match_engine.Loc{}
	}
	s = append(s, strconv.FormatFloat(float64(in.RepeatedLoc[0].Lat), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.RepeatedLoc[0].Long), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.RepeatedLoc[0].HorizontalAccuracy), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.RepeatedLoc[0].VerticalAccuracy), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.RepeatedLoc[0].Altitude), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.RepeatedLoc[0].Course), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.RepeatedLoc[0].Speed), 'e', -1, 32))
	if in.RepeatedLoc[0].Timestamp == nil {
		in.RepeatedLoc[0].Timestamp = &distributed_match_engine.Timestamp{}
	}
	s = append(s, strconv.FormatUint(uint64(in.RepeatedLoc[0].Timestamp.Seconds), 10))
	s = append(s, strconv.FormatUint(uint64(in.RepeatedLoc[0].Timestamp.Nanos), 10))
	if in.RepeatedLocNonnull == nil {
		in.RepeatedLocNonnull = make([]distributed_match_engine.Loc, 1)
	}
	s = append(s, strconv.FormatFloat(float64(in.RepeatedLocNonnull[0].Lat), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.RepeatedLocNonnull[0].Long), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.RepeatedLocNonnull[0].HorizontalAccuracy), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.RepeatedLocNonnull[0].VerticalAccuracy), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.RepeatedLocNonnull[0].Altitude), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.RepeatedLocNonnull[0].Course), 'e', -1, 32))
	s = append(s, strconv.FormatFloat(float64(in.RepeatedLocNonnull[0].Speed), 'e', -1, 32))
	if in.RepeatedLocNonnull[0].Timestamp == nil {
		in.RepeatedLocNonnull[0].Timestamp = &distributed_match_engine.Timestamp{}
	}
	s = append(s, strconv.FormatUint(uint64(in.RepeatedLocNonnull[0].Timestamp.Seconds), 10))
	s = append(s, strconv.FormatUint(uint64(in.RepeatedLocNonnull[0].Timestamp.Nanos), 10))
	return s
}

func TestGenHeaderSlicer() []string {
	s := make([]string, 0, 38)
	s = append(s, "Fields")
	s = append(s, "Name")
	s = append(s, "Db")
	s = append(s, "Fl")
	s = append(s, "I32")
	s = append(s, "I64")
	s = append(s, "U32")
	s = append(s, "U64")
	s = append(s, "S32")
	s = append(s, "S64")
	s = append(s, "F32")
	s = append(s, "F64")
	s = append(s, "Sf32")
	s = append(s, "Sf64")
	s = append(s, "Bb")
	s = append(s, "OuterEn")
	s = append(s, "InnerEn")
	s = append(s, "InnerMsg-Url")
	s = append(s, "InnerMsg-Id")
	s = append(s, "InnerMsgNonnull-Url")
	s = append(s, "InnerMsgNonnull-Id")
	s = append(s, "IncludeMsg-Name")
	s = append(s, "IncludeMsg-Id")
	s = append(s, "IncludeMsg-NestedMsg-Name")
	s = append(s, "IncludeMsgNonnull-Name")
	s = append(s, "IncludeMsgNonnull-Id")
	s = append(s, "IncludeMsgNonnull-NestedMsg-Name")
	s = append(s, "IncludeFields-Fields")
	s = append(s, "IncludeFields-Name")
	s = append(s, "IncludeFieldsNonnull-Fields")
	s = append(s, "IncludeFieldsNonnull-Name")
	s = append(s, "Loc-Lat")
	s = append(s, "Loc-Long")
	s = append(s, "Loc-HorizontalAccuracy")
	s = append(s, "Loc-VerticalAccuracy")
	s = append(s, "Loc-Altitude")
	s = append(s, "Loc-Course")
	s = append(s, "Loc-Speed")
	s = append(s, "Loc-Timestamp-Seconds")
	s = append(s, "Loc-Timestamp-Nanos")
	s = append(s, "LocNonnull-Lat")
	s = append(s, "LocNonnull-Long")
	s = append(s, "LocNonnull-HorizontalAccuracy")
	s = append(s, "LocNonnull-VerticalAccuracy")
	s = append(s, "LocNonnull-Altitude")
	s = append(s, "LocNonnull-Course")
	s = append(s, "LocNonnull-Speed")
	s = append(s, "LocNonnull-Timestamp-Seconds")
	s = append(s, "LocNonnull-Timestamp-Nanos")
	s = append(s, "RepeatedInt")
	s = append(s, "Ip")
	s = append(s, "Names")
	s = append(s, "RepeatedMsg-Name")
	s = append(s, "RepeatedMsg-Id")
	s = append(s, "RepeatedMsg-NestedMsg-Name")
	s = append(s, "RepeatedMsgNonnull-Name")
	s = append(s, "RepeatedMsgNonnull-Id")
	s = append(s, "RepeatedMsgNonnull-NestedMsg-Name")
	s = append(s, "RepeatedFields-Fields")
	s = append(s, "RepeatedFields-Name")
	s = append(s, "RepeatedFieldsNonnull-Fields")
	s = append(s, "RepeatedFieldsNonnull-Name")
	s = append(s, "RepeatedInnerMsg-Url")
	s = append(s, "RepeatedInnerMsg-Id")
	s = append(s, "RepeatedInnerMsgNonnull-Url")
	s = append(s, "RepeatedInnerMsgNonnull-Id")
	s = append(s, "RepeatedLoc-Lat")
	s = append(s, "RepeatedLoc-Long")
	s = append(s, "RepeatedLoc-HorizontalAccuracy")
	s = append(s, "RepeatedLoc-VerticalAccuracy")
	s = append(s, "RepeatedLoc-Altitude")
	s = append(s, "RepeatedLoc-Course")
	s = append(s, "RepeatedLoc-Speed")
	s = append(s, "RepeatedLoc-Timestamp-Seconds")
	s = append(s, "RepeatedLoc-Timestamp-Nanos")
	s = append(s, "RepeatedLocNonnull-Lat")
	s = append(s, "RepeatedLocNonnull-Long")
	s = append(s, "RepeatedLocNonnull-HorizontalAccuracy")
	s = append(s, "RepeatedLocNonnull-VerticalAccuracy")
	s = append(s, "RepeatedLocNonnull-Altitude")
	s = append(s, "RepeatedLocNonnull-Course")
	s = append(s, "RepeatedLocNonnull-Speed")
	s = append(s, "RepeatedLocNonnull-Timestamp-Seconds")
	s = append(s, "RepeatedLocNonnull-Timestamp-Nanos")
	return s
}

func TestGenWriteOutputArray(objs []*testgen.TestGen) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(TestGenHeaderSlicer(), "\t"))
		for _, obj := range objs {
			fmt.Fprintln(output, strings.Join(TestGenSlicer(obj), "\t"))
		}
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(objs)
	}
}

func TestGenWriteOutputOne(obj *testgen.TestGen) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(TestGenHeaderSlicer(), "\t"))
		fmt.Fprintln(output, strings.Join(TestGenSlicer(obj), "\t"))
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(obj)
	}
}
func TestGen_InnerMessageSlicer(in *testgen.TestGen_InnerMessage) []string {
	s := make([]string, 0, 2)
	s = append(s, in.Url)
	s = append(s, strconv.FormatUint(uint64(in.Id), 10))
	return s
}

func TestGen_InnerMessageHeaderSlicer() []string {
	s := make([]string, 0, 2)
	s = append(s, "Url")
	s = append(s, "Id")
	return s
}

func InnerMessageWriteOutputArray(objs []*testgen.TestGen_InnerMessage) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(TestGen_InnerMessageHeaderSlicer(), "\t"))
		for _, obj := range objs {
			fmt.Fprintln(output, strings.Join(TestGen_InnerMessageSlicer(obj), "\t"))
		}
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(objs)
	}
}

func InnerMessageWriteOutputOne(obj *testgen.TestGen_InnerMessage) {
	if cmdsup.OutputFormat == cmdsup.OutputFormatTable {
		output := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
		fmt.Fprintln(output, strings.Join(TestGen_InnerMessageHeaderSlicer(), "\t"))
		fmt.Fprintln(output, strings.Join(TestGen_InnerMessageSlicer(obj), "\t"))
		output.Flush()
	} else {
		cmdsup.WriteOutputGeneric(obj)
	}
}

var RequestCmd = &cobra.Command{
	Use: "Request",
	RunE: func(cmd *cobra.Command, args []string) error {
		// if we got this far, usage has been met.
		cmd.SilenceUsage = true
		err := parseTestGenEnums()
		if err != nil {
			return fmt.Errorf("Request failed: %s", err.Error())
		}
		return Request(&TestGenIn)
	},
}

func Request(in *testgen.TestGen) error {
	if TestApiCmd == nil {
		return fmt.Errorf("TestApi client not initialized")
	}
	ctx := context.Background()
	obj, err := TestApiCmd.Request(ctx, in)
	if err != nil {
		errstr := err.Error()
		st, ok := status.FromError(err)
		if ok {
			errstr = st.Message()
		}
		return fmt.Errorf("Request failed: %s", errstr)
	}
	TestGenWriteOutputOne(obj)
	return nil
}

func Requests(data []testgen.TestGen, err *error) {
	if *err != nil {
		return
	}
	for ii, _ := range data {
		fmt.Printf("Request %v\n", data[ii])
		myerr := Request(&data[ii])
		if myerr != nil {
			*err = myerr
			break
		}
	}
}

var TestApiCmds = []*cobra.Command{
	RequestCmd,
}

func init() {
	TestGenFlagSet.StringVar(&TestGenIn.Name, "name", "", "Name")
	TestGenFlagSet.Float64Var(&TestGenIn.Db, "db", 0, "Db")
	TestGenFlagSet.Float32Var(&TestGenIn.Fl, "fl", 0, "Fl")
	TestGenFlagSet.Int32Var(&TestGenIn.I32, "i32", 0, "I32")
	TestGenFlagSet.Int64Var(&TestGenIn.I64, "i64", 0, "I64")
	TestGenFlagSet.Uint32Var(&TestGenIn.U32, "u32", 0, "U32")
	TestGenFlagSet.Uint64Var(&TestGenIn.U64, "u64", 0, "U64")
	TestGenFlagSet.Int32Var(&TestGenIn.S32, "s32", 0, "S32")
	TestGenFlagSet.Int64Var(&TestGenIn.S64, "s64", 0, "S64")
	TestGenFlagSet.Uint32Var(&TestGenIn.F32, "f32", 0, "F32")
	TestGenFlagSet.Uint64Var(&TestGenIn.F64, "f64", 0, "F64")
	TestGenFlagSet.Int32Var(&TestGenIn.Sf32, "sf32", 0, "Sf32")
	TestGenFlagSet.Int64Var(&TestGenIn.Sf64, "sf64", 0, "Sf64")
	TestGenFlagSet.BoolVar(&TestGenIn.Bb, "bb", false, "Bb")
	TestGenFlagSet.StringVar(&TestGenInOuterEn, "outeren", "", "one of [OUTER0 OUTER1 OUTER2 OUTER3]")
	TestGenFlagSet.StringVar(&TestGenInInnerEn, "inneren", "", "one of [INNER0 INNER1 INNER2 INNER3]")
	TestGenIn.InnerMsg = &testgen.TestGen_InnerMessage{}
	TestGenFlagSet.StringVar(&TestGenIn.InnerMsg.Url, "innermsg-url", "", "InnerMsg.Url")
	TestGenFlagSet.Int64Var(&TestGenIn.InnerMsg.Id, "innermsg-id", 0, "InnerMsg.Id")
	TestGenFlagSet.StringVar(&TestGenIn.InnerMsgNonnull.Url, "innermsgnonnull-url", "", "InnerMsgNonnull.Url")
	TestGenFlagSet.Int64Var(&TestGenIn.InnerMsgNonnull.Id, "innermsgnonnull-id", 0, "InnerMsgNonnull.Id")
	TestGenIn.IncludeMsg = &testgen.IncludeMessage{}
	TestGenFlagSet.StringVar(&TestGenIn.IncludeMsg.Name, "includemsg-name", "", "IncludeMsg.Name")
	TestGenFlagSet.Uint64Var(&TestGenIn.IncludeMsg.Id, "includemsg-id", 0, "IncludeMsg.Id")
	TestGenIn.IncludeMsg.NestedMsg = &testgen.NestedMessage{}
	TestGenFlagSet.StringVar(&TestGenIn.IncludeMsg.NestedMsg.Name, "includemsg-nestedmsg-name", "", "IncludeMsg.NestedMsg.Name")
	TestGenFlagSet.StringVar(&TestGenIn.IncludeMsgNonnull.Name, "includemsgnonnull-name", "", "IncludeMsgNonnull.Name")
	TestGenFlagSet.Uint64Var(&TestGenIn.IncludeMsgNonnull.Id, "includemsgnonnull-id", 0, "IncludeMsgNonnull.Id")
	TestGenIn.IncludeMsgNonnull.NestedMsg = &testgen.NestedMessage{}
	TestGenFlagSet.StringVar(&TestGenIn.IncludeMsgNonnull.NestedMsg.Name, "includemsgnonnull-nestedmsg-name", "", "IncludeMsgNonnull.NestedMsg.Name")
	TestGenIn.IncludeFields = &testgen.IncludeFields{}
	TestGenFlagSet.StringVar(&TestGenIn.IncludeFields.Name, "includefields-name", "", "IncludeFields.Name")
	TestGenFlagSet.StringVar(&TestGenIn.IncludeFieldsNonnull.Name, "includefieldsnonnull-name", "", "IncludeFieldsNonnull.Name")
	TestGenIn.Loc = &distributed_match_engine.Loc{}
	TestGenFlagSet.Float64Var(&TestGenIn.Loc.Lat, "loc-lat", 0, "Loc.Lat")
	TestGenFlagSet.Float64Var(&TestGenIn.Loc.Long, "loc-long", 0, "Loc.Long")
	TestGenFlagSet.Float64Var(&TestGenIn.Loc.HorizontalAccuracy, "loc-horizontalaccuracy", 0, "Loc.HorizontalAccuracy")
	TestGenFlagSet.Float64Var(&TestGenIn.Loc.VerticalAccuracy, "loc-verticalaccuracy", 0, "Loc.VerticalAccuracy")
	TestGenFlagSet.Float64Var(&TestGenIn.Loc.Altitude, "loc-altitude", 0, "Loc.Altitude")
	TestGenFlagSet.Float64Var(&TestGenIn.Loc.Course, "loc-course", 0, "Loc.Course")
	TestGenFlagSet.Float64Var(&TestGenIn.Loc.Speed, "loc-speed", 0, "Loc.Speed")
	TestGenIn.Loc.Timestamp = &distributed_match_engine.Timestamp{}
	TestGenFlagSet.Int64Var(&TestGenIn.Loc.Timestamp.Seconds, "loc-timestamp-seconds", 0, "Loc.Timestamp.Seconds")
	TestGenFlagSet.Int32Var(&TestGenIn.Loc.Timestamp.Nanos, "loc-timestamp-nanos", 0, "Loc.Timestamp.Nanos")
	TestGenFlagSet.Float64Var(&TestGenIn.LocNonnull.Lat, "locnonnull-lat", 0, "LocNonnull.Lat")
	TestGenFlagSet.Float64Var(&TestGenIn.LocNonnull.Long, "locnonnull-long", 0, "LocNonnull.Long")
	TestGenFlagSet.Float64Var(&TestGenIn.LocNonnull.HorizontalAccuracy, "locnonnull-horizontalaccuracy", 0, "LocNonnull.HorizontalAccuracy")
	TestGenFlagSet.Float64Var(&TestGenIn.LocNonnull.VerticalAccuracy, "locnonnull-verticalaccuracy", 0, "LocNonnull.VerticalAccuracy")
	TestGenFlagSet.Float64Var(&TestGenIn.LocNonnull.Altitude, "locnonnull-altitude", 0, "LocNonnull.Altitude")
	TestGenFlagSet.Float64Var(&TestGenIn.LocNonnull.Course, "locnonnull-course", 0, "LocNonnull.Course")
	TestGenFlagSet.Float64Var(&TestGenIn.LocNonnull.Speed, "locnonnull-speed", 0, "LocNonnull.Speed")
	TestGenIn.LocNonnull.Timestamp = &distributed_match_engine.Timestamp{}
	TestGenFlagSet.Int64Var(&TestGenIn.LocNonnull.Timestamp.Seconds, "locnonnull-timestamp-seconds", 0, "LocNonnull.Timestamp.Seconds")
	TestGenFlagSet.Int32Var(&TestGenIn.LocNonnull.Timestamp.Nanos, "locnonnull-timestamp-nanos", 0, "LocNonnull.Timestamp.Nanos")
	TestGenFlagSet.BytesHexVar(&TestGenIn.Ip, "ip", nil, "Ip")
	RequestCmd.Flags().AddFlagSet(TestGenFlagSet)
}

func TestApiAllowNoConfig() {
	RequestCmd.Flags().AddFlagSet(TestGenNoConfigFlagSet)
}

func TestGenSetFields() {
	TestGenIn.Fields = make([]string, 0)
	if TestGenFlagSet.Lookup("name").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "2")
	}
	if TestGenFlagSet.Lookup("db").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "3")
	}
	if TestGenFlagSet.Lookup("fl").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "4")
	}
	if TestGenFlagSet.Lookup("i32").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "5")
	}
	if TestGenFlagSet.Lookup("i64").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "6")
	}
	if TestGenFlagSet.Lookup("u32").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "7")
	}
	if TestGenFlagSet.Lookup("u64").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "8")
	}
	if TestGenFlagSet.Lookup("s32").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "9")
	}
	if TestGenFlagSet.Lookup("s64").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "10")
	}
	if TestGenFlagSet.Lookup("f32").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "11")
	}
	if TestGenFlagSet.Lookup("f64").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "12")
	}
	if TestGenFlagSet.Lookup("sf32").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "13")
	}
	if TestGenFlagSet.Lookup("sf64").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "14")
	}
	if TestGenFlagSet.Lookup("bb").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "15")
	}
	if TestGenFlagSet.Lookup("outeren").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "16")
	}
	if TestGenFlagSet.Lookup("inneren").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "17")
	}
	if TestGenFlagSet.Lookup("innermsg-url").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "18.1")
	}
	if TestGenFlagSet.Lookup("innermsg-id").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "18.2")
	}
	if TestGenFlagSet.Lookup("innermsgnonnull-url").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "19.1")
	}
	if TestGenFlagSet.Lookup("innermsgnonnull-id").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "19.2")
	}
	if TestGenFlagSet.Lookup("includemsg-name").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "20.1")
	}
	if TestGenFlagSet.Lookup("includemsg-id").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "20.2")
	}
	if TestGenFlagSet.Lookup("includemsg-nestedmsg-name").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "20.3.1")
	}
	if TestGenFlagSet.Lookup("includemsgnonnull-name").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "21.1")
	}
	if TestGenFlagSet.Lookup("includemsgnonnull-id").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "21.2")
	}
	if TestGenFlagSet.Lookup("includemsgnonnull-nestedmsg-name").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "21.3.1")
	}
	if TestGenFlagSet.Lookup("includefields-name").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "22.2")
	}
	if TestGenFlagSet.Lookup("includefieldsnonnull-name").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "23.2")
	}
	if TestGenFlagSet.Lookup("loc-lat").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "24.1")
	}
	if TestGenFlagSet.Lookup("loc-long").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "24.2")
	}
	if TestGenFlagSet.Lookup("loc-horizontalaccuracy").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "24.3")
	}
	if TestGenFlagSet.Lookup("loc-verticalaccuracy").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "24.4")
	}
	if TestGenFlagSet.Lookup("loc-altitude").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "24.5")
	}
	if TestGenFlagSet.Lookup("loc-course").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "24.6")
	}
	if TestGenFlagSet.Lookup("loc-speed").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "24.7")
	}
	if TestGenFlagSet.Lookup("loc-timestamp-seconds").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "24.8.1")
	}
	if TestGenFlagSet.Lookup("loc-timestamp-nanos").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "24.8.2")
	}
	if TestGenFlagSet.Lookup("locnonnull-lat").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "25.1")
	}
	if TestGenFlagSet.Lookup("locnonnull-long").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "25.2")
	}
	if TestGenFlagSet.Lookup("locnonnull-horizontalaccuracy").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "25.3")
	}
	if TestGenFlagSet.Lookup("locnonnull-verticalaccuracy").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "25.4")
	}
	if TestGenFlagSet.Lookup("locnonnull-altitude").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "25.5")
	}
	if TestGenFlagSet.Lookup("locnonnull-course").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "25.6")
	}
	if TestGenFlagSet.Lookup("locnonnull-speed").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "25.7")
	}
	if TestGenFlagSet.Lookup("locnonnull-timestamp-seconds").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "25.8.1")
	}
	if TestGenFlagSet.Lookup("locnonnull-timestamp-nanos").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "25.8.2")
	}
	if TestGenFlagSet.Lookup("ip").Changed {
		TestGenIn.Fields = append(TestGenIn.Fields, "27")
	}
}

func parseTestGenEnums() error {
	if TestGenInOuterEn != "" {
		switch TestGenInOuterEn {
		case "OUTER0":
			TestGenIn.OuterEn = testgen.OuterEnum(0)
		case "OUTER1":
			TestGenIn.OuterEn = testgen.OuterEnum(1)
		case "OUTER2":
			TestGenIn.OuterEn = testgen.OuterEnum(2)
		case "OUTER3":
			TestGenIn.OuterEn = testgen.OuterEnum(3)
		default:
			return errors.New("Invalid value for TestGenInOuterEn")
		}
	}
	if TestGenInInnerEn != "" {
		switch TestGenInInnerEn {
		case "INNER0":
			TestGenIn.InnerEn = testgen.TestGen_InnerEnum(0)
		case "INNER1":
			TestGenIn.InnerEn = testgen.TestGen_InnerEnum(1)
		case "INNER2":
			TestGenIn.InnerEn = testgen.TestGen_InnerEnum(2)
		case "INNER3":
			TestGenIn.InnerEn = testgen.TestGen_InnerEnum(3)
		default:
			return errors.New("Invalid value for TestGenInInnerEn")
		}
	}
	return nil
}
