// Code generated by protoc-gen-go. DO NOT EDIT.
// source: squeeze.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// All supported protocols
type Protocol int32

const (
	Protocol_UNKNOWN   Protocol = 0
	Protocol_HTTP      Protocol = 1
	Protocol_TCP       Protocol = 2
	Protocol_UDP       Protocol = 3
	Protocol_REDIS     Protocol = 4
	Protocol_MYSQL     Protocol = 5
	Protocol_MONGO     Protocol = 6
	Protocol_GRPC      Protocol = 7
	Protocol_THRIFT    Protocol = 8
	Protocol_WEBSOCKET Protocol = 9
)

var Protocol_name = map[int32]string{
	0: "UNKNOWN",
	1: "HTTP",
	2: "TCP",
	3: "UDP",
	4: "REDIS",
	5: "MYSQL",
	6: "MONGO",
	7: "GRPC",
	8: "THRIFT",
	9: "WEBSOCKET",
}
var Protocol_value = map[string]int32{
	"UNKNOWN":   0,
	"HTTP":      1,
	"TCP":       2,
	"UDP":       3,
	"REDIS":     4,
	"MYSQL":     5,
	"MONGO":     6,
	"GRPC":      7,
	"THRIFT":    8,
	"WEBSOCKET": 9,
}

func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}
func (Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_squeeze_0884af5490b4c2ad, []int{0}
}

type HeartBeatRequest_Task_Status int32

const (
	HeartBeatRequest_Task_DONE    HeartBeatRequest_Task_Status = 0
	HeartBeatRequest_Task_RUNNING HeartBeatRequest_Task_Status = 1
)

var HeartBeatRequest_Task_Status_name = map[int32]string{
	0: "DONE",
	1: "RUNNING",
}
var HeartBeatRequest_Task_Status_value = map[string]int32{
	"DONE":    0,
	"RUNNING": 1,
}

func (x HeartBeatRequest_Task_Status) String() string {
	return proto.EnumName(HeartBeatRequest_Task_Status_name, int32(x))
}
func (HeartBeatRequest_Task_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_squeeze_0884af5490b4c2ad, []int{0, 1, 0}
}

type ExecuteTaskRequest_Command int32

const (
	ExecuteTaskRequest_START ExecuteTaskRequest_Command = 0
	ExecuteTaskRequest_STOP  ExecuteTaskRequest_Command = 1
)

var ExecuteTaskRequest_Command_name = map[int32]string{
	0: "START",
	1: "STOP",
}
var ExecuteTaskRequest_Command_value = map[string]int32{
	"START": 0,
	"STOP":  1,
}

func (x ExecuteTaskRequest_Command) String() string {
	return proto.EnumName(ExecuteTaskRequest_Command_name, int32(x))
}
func (ExecuteTaskRequest_Command) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_squeeze_0884af5490b4c2ad, []int{3, 0}
}

type ExecuteTaskResponse_Status int32

const (
	ExecuteTaskResponse_SUCC ExecuteTaskResponse_Status = 0
	ExecuteTaskResponse_FAIL ExecuteTaskResponse_Status = 1
)

var ExecuteTaskResponse_Status_name = map[int32]string{
	0: "SUCC",
	1: "FAIL",
}
var ExecuteTaskResponse_Status_value = map[string]int32{
	"SUCC": 0,
	"FAIL": 1,
}

func (x ExecuteTaskResponse_Status) String() string {
	return proto.EnumName(ExecuteTaskResponse_Status_name, int32(x))
}
func (ExecuteTaskResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_squeeze_0884af5490b4c2ad, []int{4, 0}
}

// HeartBeatRequest describes the tasks status on the slave
// and the information about the slave & host.
type HeartBeatRequest struct {
	Task                 *HeartBeatRequest_Task      `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Info                 *HeartBeatRequest_SlaveInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *HeartBeatRequest) Reset()         { *m = HeartBeatRequest{} }
func (m *HeartBeatRequest) String() string { return proto.CompactTextString(m) }
func (*HeartBeatRequest) ProtoMessage()    {}
func (*HeartBeatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_squeeze_0884af5490b4c2ad, []int{0}
}
func (m *HeartBeatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeatRequest.Unmarshal(m, b)
}
func (m *HeartBeatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeatRequest.Marshal(b, m, deterministic)
}
func (dst *HeartBeatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeatRequest.Merge(dst, src)
}
func (m *HeartBeatRequest) XXX_Size() int {
	return xxx_messageInfo_HeartBeatRequest.Size(m)
}
func (m *HeartBeatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeatRequest proto.InternalMessageInfo

func (m *HeartBeatRequest) GetTask() *HeartBeatRequest_Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *HeartBeatRequest) GetInfo() *HeartBeatRequest_SlaveInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type HeartBeatRequest_SlaveInfo struct {
	GrpcPort             uint32   `protobuf:"varint,1,opt,name=grpcPort,proto3" json:"grpcPort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartBeatRequest_SlaveInfo) Reset()         { *m = HeartBeatRequest_SlaveInfo{} }
func (m *HeartBeatRequest_SlaveInfo) String() string { return proto.CompactTextString(m) }
func (*HeartBeatRequest_SlaveInfo) ProtoMessage()    {}
func (*HeartBeatRequest_SlaveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_squeeze_0884af5490b4c2ad, []int{0, 0}
}
func (m *HeartBeatRequest_SlaveInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeatRequest_SlaveInfo.Unmarshal(m, b)
}
func (m *HeartBeatRequest_SlaveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeatRequest_SlaveInfo.Marshal(b, m, deterministic)
}
func (dst *HeartBeatRequest_SlaveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeatRequest_SlaveInfo.Merge(dst, src)
}
func (m *HeartBeatRequest_SlaveInfo) XXX_Size() int {
	return xxx_messageInfo_HeartBeatRequest_SlaveInfo.Size(m)
}
func (m *HeartBeatRequest_SlaveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeatRequest_SlaveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeatRequest_SlaveInfo proto.InternalMessageInfo

func (m *HeartBeatRequest_SlaveInfo) GetGrpcPort() uint32 {
	if m != nil {
		return m.GrpcPort
	}
	return 0
}

type HeartBeatRequest_Task struct {
	Id                   uint32                       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               HeartBeatRequest_Task_Status `protobuf:"varint,2,opt,name=status,proto3,enum=pb.HeartBeatRequest_Task_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *HeartBeatRequest_Task) Reset()         { *m = HeartBeatRequest_Task{} }
func (m *HeartBeatRequest_Task) String() string { return proto.CompactTextString(m) }
func (*HeartBeatRequest_Task) ProtoMessage()    {}
func (*HeartBeatRequest_Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_squeeze_0884af5490b4c2ad, []int{0, 1}
}
func (m *HeartBeatRequest_Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeatRequest_Task.Unmarshal(m, b)
}
func (m *HeartBeatRequest_Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeatRequest_Task.Marshal(b, m, deterministic)
}
func (dst *HeartBeatRequest_Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeatRequest_Task.Merge(dst, src)
}
func (m *HeartBeatRequest_Task) XXX_Size() int {
	return xxx_messageInfo_HeartBeatRequest_Task.Size(m)
}
func (m *HeartBeatRequest_Task) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeatRequest_Task.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeatRequest_Task proto.InternalMessageInfo

func (m *HeartBeatRequest_Task) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HeartBeatRequest_Task) GetStatus() HeartBeatRequest_Task_Status {
	if m != nil {
		return m.Status
	}
	return HeartBeatRequest_Task_DONE
}

// HeartBeatResponse for HeartBeatRequest, is used to send tasks from master to slave.
// !!!Send task is not yet supported!!!
type HeartBeatResponse struct {
	Tasks                []*TaskRequest `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *HeartBeatResponse) Reset()         { *m = HeartBeatResponse{} }
func (m *HeartBeatResponse) String() string { return proto.CompactTextString(m) }
func (*HeartBeatResponse) ProtoMessage()    {}
func (*HeartBeatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_squeeze_0884af5490b4c2ad, []int{1}
}
func (m *HeartBeatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeatResponse.Unmarshal(m, b)
}
func (m *HeartBeatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeatResponse.Marshal(b, m, deterministic)
}
func (dst *HeartBeatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeatResponse.Merge(dst, src)
}
func (m *HeartBeatResponse) XXX_Size() int {
	return xxx_messageInfo_HeartBeatResponse.Size(m)
}
func (m *HeartBeatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeatResponse proto.InternalMessageInfo

func (m *HeartBeatResponse) GetTasks() []*TaskRequest {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type TaskRequest struct {
	Concurrency uint32 `protobuf:"varint,1,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	Requests    uint32 `protobuf:"varint,2,opt,name=requests,proto3" json:"requests,omitempty"`
	RateLimit   uint32 `protobuf:"varint,3,opt,name=rateLimit,proto3" json:"rateLimit,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*TaskRequest_Http
	//	*TaskRequest_Websocket
	Type                 isTaskRequest_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TaskRequest) Reset()         { *m = TaskRequest{} }
func (m *TaskRequest) String() string { return proto.CompactTextString(m) }
func (*TaskRequest) ProtoMessage()    {}
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_squeeze_0884af5490b4c2ad, []int{2}
}
func (m *TaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskRequest.Unmarshal(m, b)
}
func (m *TaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskRequest.Marshal(b, m, deterministic)
}
func (dst *TaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRequest.Merge(dst, src)
}
func (m *TaskRequest) XXX_Size() int {
	return xxx_messageInfo_TaskRequest.Size(m)
}
func (m *TaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRequest proto.InternalMessageInfo

func (m *TaskRequest) GetConcurrency() uint32 {
	if m != nil {
		return m.Concurrency
	}
	return 0
}

func (m *TaskRequest) GetRequests() uint32 {
	if m != nil {
		return m.Requests
	}
	return 0
}

func (m *TaskRequest) GetRateLimit() uint32 {
	if m != nil {
		return m.RateLimit
	}
	return 0
}

type isTaskRequest_Type interface {
	isTaskRequest_Type()
}

type TaskRequest_Http struct {
	Http *HttpTask `protobuf:"bytes,4,opt,name=http,proto3,oneof"`
}

type TaskRequest_Websocket struct {
	Websocket *WebsocketTask `protobuf:"bytes,5,opt,name=websocket,proto3,oneof"`
}

func (*TaskRequest_Http) isTaskRequest_Type() {}

func (*TaskRequest_Websocket) isTaskRequest_Type() {}

func (m *TaskRequest) GetType() isTaskRequest_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *TaskRequest) GetHttp() *HttpTask {
	if x, ok := m.GetType().(*TaskRequest_Http); ok {
		return x.Http
	}
	return nil
}

func (m *TaskRequest) GetWebsocket() *WebsocketTask {
	if x, ok := m.GetType().(*TaskRequest_Websocket); ok {
		return x.Websocket
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TaskRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TaskRequest_OneofMarshaler, _TaskRequest_OneofUnmarshaler, _TaskRequest_OneofSizer, []interface{}{
		(*TaskRequest_Http)(nil),
		(*TaskRequest_Websocket)(nil),
	}
}

func _TaskRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TaskRequest)
	// type
	switch x := m.Type.(type) {
	case *TaskRequest_Http:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Http); err != nil {
			return err
		}
	case *TaskRequest_Websocket:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Websocket); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TaskRequest.Type has unexpected type %T", x)
	}
	return nil
}

func _TaskRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TaskRequest)
	switch tag {
	case 4: // type.http
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HttpTask)
		err := b.DecodeMessage(msg)
		m.Type = &TaskRequest_Http{msg}
		return true, err
	case 5: // type.websocket
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WebsocketTask)
		err := b.DecodeMessage(msg)
		m.Type = &TaskRequest_Websocket{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TaskRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TaskRequest)
	// type
	switch x := m.Type.(type) {
	case *TaskRequest_Http:
		s := proto.Size(x.Http)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TaskRequest_Websocket:
		s := proto.Size(x.Websocket)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// ExecuteTaskRequest describes the task information to be executed.
type ExecuteTaskRequest struct {
	Cmd                  ExecuteTaskRequest_Command `protobuf:"varint,1,opt,name=cmd,proto3,enum=pb.ExecuteTaskRequest_Command" json:"cmd,omitempty"`
	Protocol             Protocol                   `protobuf:"varint,2,opt,name=protocol,proto3,enum=pb.Protocol" json:"protocol,omitempty"`
	Callback             string                     `protobuf:"bytes,3,opt,name=callback,proto3" json:"callback,omitempty"`
	Duration             uint32                     `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Task                 *TaskRequest               `protobuf:"bytes,8,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ExecuteTaskRequest) Reset()         { *m = ExecuteTaskRequest{} }
func (m *ExecuteTaskRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteTaskRequest) ProtoMessage()    {}
func (*ExecuteTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_squeeze_0884af5490b4c2ad, []int{3}
}
func (m *ExecuteTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteTaskRequest.Unmarshal(m, b)
}
func (m *ExecuteTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteTaskRequest.Marshal(b, m, deterministic)
}
func (dst *ExecuteTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteTaskRequest.Merge(dst, src)
}
func (m *ExecuteTaskRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteTaskRequest.Size(m)
}
func (m *ExecuteTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteTaskRequest proto.InternalMessageInfo

func (m *ExecuteTaskRequest) GetCmd() ExecuteTaskRequest_Command {
	if m != nil {
		return m.Cmd
	}
	return ExecuteTaskRequest_START
}

func (m *ExecuteTaskRequest) GetProtocol() Protocol {
	if m != nil {
		return m.Protocol
	}
	return Protocol_UNKNOWN
}

func (m *ExecuteTaskRequest) GetCallback() string {
	if m != nil {
		return m.Callback
	}
	return ""
}

func (m *ExecuteTaskRequest) GetDuration() uint32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *ExecuteTaskRequest) GetTask() *TaskRequest {
	if m != nil {
		return m.Task
	}
	return nil
}

// ExecuteTaskResponse for ExecuteTaskRequest, is used to report the results of this execution task.
type ExecuteTaskResponse struct {
	// Slave address
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	// Task status
	Status ExecuteTaskResponse_Status `protobuf:"varint,2,opt,name=status,proto3,enum=pb.ExecuteTaskResponse_Status" json:"status,omitempty"`
	// When status is equal to FAIL, error will be set
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	// Task results in json format
	Data                 string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteTaskResponse) Reset()         { *m = ExecuteTaskResponse{} }
func (m *ExecuteTaskResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteTaskResponse) ProtoMessage()    {}
func (*ExecuteTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_squeeze_0884af5490b4c2ad, []int{4}
}
func (m *ExecuteTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteTaskResponse.Unmarshal(m, b)
}
func (m *ExecuteTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteTaskResponse.Marshal(b, m, deterministic)
}
func (dst *ExecuteTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteTaskResponse.Merge(dst, src)
}
func (m *ExecuteTaskResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteTaskResponse.Size(m)
}
func (m *ExecuteTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteTaskResponse proto.InternalMessageInfo

func (m *ExecuteTaskResponse) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ExecuteTaskResponse) GetStatus() ExecuteTaskResponse_Status {
	if m != nil {
		return m.Status
	}
	return ExecuteTaskResponse_SUCC
}

func (m *ExecuteTaskResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ExecuteTaskResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type WebsocketTask struct {
	Scheme               string   `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Path                 string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Timeout              uint32   `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Body                 string   `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	MaxResults           int32    `protobuf:"varint,6,opt,name=maxResults,proto3" json:"maxResults,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebsocketTask) Reset()         { *m = WebsocketTask{} }
func (m *WebsocketTask) String() string { return proto.CompactTextString(m) }
func (*WebsocketTask) ProtoMessage()    {}
func (*WebsocketTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_squeeze_0884af5490b4c2ad, []int{5}
}
func (m *WebsocketTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebsocketTask.Unmarshal(m, b)
}
func (m *WebsocketTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebsocketTask.Marshal(b, m, deterministic)
}
func (dst *WebsocketTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebsocketTask.Merge(dst, src)
}
func (m *WebsocketTask) XXX_Size() int {
	return xxx_messageInfo_WebsocketTask.Size(m)
}
func (m *WebsocketTask) XXX_DiscardUnknown() {
	xxx_messageInfo_WebsocketTask.DiscardUnknown(m)
}

var xxx_messageInfo_WebsocketTask proto.InternalMessageInfo

func (m *WebsocketTask) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *WebsocketTask) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *WebsocketTask) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *WebsocketTask) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *WebsocketTask) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *WebsocketTask) GetMaxResults() int32 {
	if m != nil {
		return m.MaxResults
	}
	return 0
}

type HttpTask struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Timeout              uint32   `protobuf:"varint,7,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Http2                bool     `protobuf:"varint,8,opt,name=http2,proto3" json:"http2,omitempty"`
	DisableRedirects     bool     `protobuf:"varint,9,opt,name=disableRedirects,proto3" json:"disableRedirects,omitempty"`
	DisableKeepalives    bool     `protobuf:"varint,10,opt,name=disableKeepalives,proto3" json:"disableKeepalives,omitempty"`
	DisableCompression   bool     `protobuf:"varint,11,opt,name=disableCompression,proto3" json:"disableCompression,omitempty"`
	ProxyAddr            string   `protobuf:"bytes,12,opt,name=proxyAddr,proto3" json:"proxyAddr,omitempty"`
	MaxIdleConn          uint32   `protobuf:"varint,13,opt,name=maxIdleConn,proto3" json:"maxIdleConn,omitempty"`
	Headers              []string `protobuf:"bytes,14,rep,name=headers,proto3" json:"headers,omitempty"`
	ContentType          string   `protobuf:"bytes,15,opt,name=contentType,proto3" json:"contentType,omitempty"`
	MaxResults           int32    `protobuf:"varint,16,opt,name=maxResults,proto3" json:"maxResults,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpTask) Reset()         { *m = HttpTask{} }
func (m *HttpTask) String() string { return proto.CompactTextString(m) }
func (*HttpTask) ProtoMessage()    {}
func (*HttpTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_squeeze_0884af5490b4c2ad, []int{6}
}
func (m *HttpTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpTask.Unmarshal(m, b)
}
func (m *HttpTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpTask.Marshal(b, m, deterministic)
}
func (dst *HttpTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpTask.Merge(dst, src)
}
func (m *HttpTask) XXX_Size() int {
	return xxx_messageInfo_HttpTask.Size(m)
}
func (m *HttpTask) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpTask.DiscardUnknown(m)
}

var xxx_messageInfo_HttpTask proto.InternalMessageInfo

func (m *HttpTask) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *HttpTask) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *HttpTask) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *HttpTask) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *HttpTask) GetHttp2() bool {
	if m != nil {
		return m.Http2
	}
	return false
}

func (m *HttpTask) GetDisableRedirects() bool {
	if m != nil {
		return m.DisableRedirects
	}
	return false
}

func (m *HttpTask) GetDisableKeepalives() bool {
	if m != nil {
		return m.DisableKeepalives
	}
	return false
}

func (m *HttpTask) GetDisableCompression() bool {
	if m != nil {
		return m.DisableCompression
	}
	return false
}

func (m *HttpTask) GetProxyAddr() string {
	if m != nil {
		return m.ProxyAddr
	}
	return ""
}

func (m *HttpTask) GetMaxIdleConn() uint32 {
	if m != nil {
		return m.MaxIdleConn
	}
	return 0
}

func (m *HttpTask) GetHeaders() []string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HttpTask) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *HttpTask) GetMaxResults() int32 {
	if m != nil {
		return m.MaxResults
	}
	return 0
}

func init() {
	proto.RegisterType((*HeartBeatRequest)(nil), "pb.HeartBeatRequest")
	proto.RegisterType((*HeartBeatRequest_SlaveInfo)(nil), "pb.HeartBeatRequest.SlaveInfo")
	proto.RegisterType((*HeartBeatRequest_Task)(nil), "pb.HeartBeatRequest.Task")
	proto.RegisterType((*HeartBeatResponse)(nil), "pb.HeartBeatResponse")
	proto.RegisterType((*TaskRequest)(nil), "pb.TaskRequest")
	proto.RegisterType((*ExecuteTaskRequest)(nil), "pb.ExecuteTaskRequest")
	proto.RegisterType((*ExecuteTaskResponse)(nil), "pb.ExecuteTaskResponse")
	proto.RegisterType((*WebsocketTask)(nil), "pb.WebsocketTask")
	proto.RegisterType((*HttpTask)(nil), "pb.HttpTask")
	proto.RegisterEnum("pb.Protocol", Protocol_name, Protocol_value)
	proto.RegisterEnum("pb.HeartBeatRequest_Task_Status", HeartBeatRequest_Task_Status_name, HeartBeatRequest_Task_Status_value)
	proto.RegisterEnum("pb.ExecuteTaskRequest_Command", ExecuteTaskRequest_Command_name, ExecuteTaskRequest_Command_value)
	proto.RegisterEnum("pb.ExecuteTaskResponse_Status", ExecuteTaskResponse_Status_name, ExecuteTaskResponse_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SqueezeServiceClient is the client API for SqueezeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SqueezeServiceClient interface {
	// ExecuteTask is used to perform tasks on the slave node
	ExecuteTask(ctx context.Context, in *ExecuteTaskRequest, opts ...grpc.CallOption) (*ExecuteTaskResponse, error)
	// HeartBeat is used to report task status to master node
	HeartBeat(ctx context.Context, opts ...grpc.CallOption) (SqueezeService_HeartBeatClient, error)
}

type squeezeServiceClient struct {
	cc *grpc.ClientConn
}

func NewSqueezeServiceClient(cc *grpc.ClientConn) SqueezeServiceClient {
	return &squeezeServiceClient{cc}
}

func (c *squeezeServiceClient) ExecuteTask(ctx context.Context, in *ExecuteTaskRequest, opts ...grpc.CallOption) (*ExecuteTaskResponse, error) {
	out := new(ExecuteTaskResponse)
	err := c.cc.Invoke(ctx, "/pb.SqueezeService/ExecuteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *squeezeServiceClient) HeartBeat(ctx context.Context, opts ...grpc.CallOption) (SqueezeService_HeartBeatClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SqueezeService_serviceDesc.Streams[0], "/pb.SqueezeService/HeartBeat", opts...)
	if err != nil {
		return nil, err
	}
	x := &squeezeServiceHeartBeatClient{stream}
	return x, nil
}

type SqueezeService_HeartBeatClient interface {
	Send(*HeartBeatRequest) error
	Recv() (*HeartBeatResponse, error)
	grpc.ClientStream
}

type squeezeServiceHeartBeatClient struct {
	grpc.ClientStream
}

func (x *squeezeServiceHeartBeatClient) Send(m *HeartBeatRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *squeezeServiceHeartBeatClient) Recv() (*HeartBeatResponse, error) {
	m := new(HeartBeatResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SqueezeServiceServer is the server API for SqueezeService service.
type SqueezeServiceServer interface {
	// ExecuteTask is used to perform tasks on the slave node
	ExecuteTask(context.Context, *ExecuteTaskRequest) (*ExecuteTaskResponse, error)
	// HeartBeat is used to report task status to master node
	HeartBeat(SqueezeService_HeartBeatServer) error
}

func RegisterSqueezeServiceServer(s *grpc.Server, srv SqueezeServiceServer) {
	s.RegisterService(&_SqueezeService_serviceDesc, srv)
}

func _SqueezeService_ExecuteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SqueezeServiceServer).ExecuteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SqueezeService/ExecuteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SqueezeServiceServer).ExecuteTask(ctx, req.(*ExecuteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SqueezeService_HeartBeat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SqueezeServiceServer).HeartBeat(&squeezeServiceHeartBeatServer{stream})
}

type SqueezeService_HeartBeatServer interface {
	Send(*HeartBeatResponse) error
	Recv() (*HeartBeatRequest, error)
	grpc.ServerStream
}

type squeezeServiceHeartBeatServer struct {
	grpc.ServerStream
}

func (x *squeezeServiceHeartBeatServer) Send(m *HeartBeatResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *squeezeServiceHeartBeatServer) Recv() (*HeartBeatRequest, error) {
	m := new(HeartBeatRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SqueezeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SqueezeService",
	HandlerType: (*SqueezeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteTask",
			Handler:    _SqueezeService_ExecuteTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HeartBeat",
			Handler:       _SqueezeService_HeartBeat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "squeeze.proto",
}

func init() { proto.RegisterFile("squeeze.proto", fileDescriptor_squeeze_0884af5490b4c2ad) }

var fileDescriptor_squeeze_0884af5490b4c2ad = []byte{
	// 912 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0x6f, 0x6f, 0xe3, 0xc4,
	0x13, 0xae, 0xf3, 0xdf, 0x93, 0x4b, 0xce, 0xdd, 0x5f, 0x7f, 0x87, 0x89, 0x4e, 0x25, 0x32, 0x42,
	0x44, 0x27, 0x88, 0x8e, 0x20, 0x21, 0xc4, 0x0b, 0x44, 0x9b, 0xe6, 0xda, 0xa8, 0x25, 0x09, 0x6b,
	0x57, 0x15, 0x2f, 0x37, 0xf6, 0x1e, 0xb1, 0x1a, 0x7b, 0xdd, 0xdd, 0x4d, 0x69, 0xf8, 0x16, 0xf0,
	0x1e, 0xde, 0xf3, 0x4d, 0xf8, 0x32, 0x7c, 0x07, 0xb4, 0xeb, 0x4d, 0xea, 0x26, 0xbd, 0x77, 0x33,
	0xcf, 0x3c, 0xeb, 0x99, 0x67, 0x67, 0x76, 0x0c, 0x2d, 0x71, 0xb7, 0xa2, 0xf4, 0x37, 0xda, 0xcf,
	0x38, 0x93, 0x0c, 0x95, 0xb2, 0xb9, 0xf7, 0x7b, 0x09, 0x9c, 0x0b, 0x4a, 0xb8, 0x3c, 0xa5, 0x44,
	0x62, 0x7a, 0xb7, 0xa2, 0x42, 0xa2, 0x2f, 0xa1, 0x22, 0x89, 0xb8, 0x75, 0xad, 0xae, 0xd5, 0x6b,
	0x0e, 0x3e, 0xee, 0x67, 0xf3, 0xfe, 0x2e, 0xa7, 0x1f, 0x10, 0x71, 0x8b, 0x35, 0x0d, 0x0d, 0xa0,
	0x12, 0xa7, 0xef, 0x99, 0x5b, 0xd2, 0xf4, 0xe3, 0x67, 0xe9, 0xfe, 0x92, 0xdc, 0xd3, 0x71, 0xfa,
	0x9e, 0x61, 0xcd, 0xed, 0x7c, 0x0e, 0xf6, 0x16, 0x42, 0x1d, 0x68, 0xfc, 0xc2, 0xb3, 0x70, 0xc6,
	0xb8, 0xd4, 0x39, 0x5b, 0x78, 0xeb, 0x77, 0xee, 0xa0, 0xa2, 0x52, 0xa1, 0x36, 0x94, 0xe2, 0xc8,
	0x44, 0x4b, 0x71, 0x84, 0xbe, 0x85, 0x9a, 0x90, 0x44, 0xae, 0x84, 0x4e, 0xdb, 0x1e, 0x74, 0x3f,
	0x58, 0x65, 0xdf, 0xd7, 0x3c, 0x6c, 0xf8, 0xde, 0x27, 0x50, 0xcb, 0x11, 0xd4, 0x80, 0xca, 0xd9,
	0x74, 0x32, 0x72, 0x0e, 0x50, 0x13, 0xea, 0xf8, 0x7a, 0x32, 0x19, 0x4f, 0xce, 0x1d, 0xcb, 0xfb,
	0x0e, 0x0e, 0x0b, 0x1f, 0x12, 0x19, 0x4b, 0x05, 0x45, 0x9f, 0x41, 0x55, 0x89, 0x15, 0xae, 0xd5,
	0x2d, 0xf7, 0x9a, 0x83, 0x97, 0x2a, 0x9d, 0xbe, 0x83, 0x3c, 0x13, 0xce, 0xa3, 0xde, 0x3f, 0x16,
	0x34, 0x0b, 0x30, 0xea, 0x42, 0x33, 0x64, 0x69, 0xb8, 0xe2, 0x9c, 0xa6, 0xe1, 0xda, 0xd4, 0x5f,
	0x84, 0x94, 0x78, 0x9e, 0x93, 0x73, 0x29, 0x2d, 0xbc, 0xf5, 0xd1, 0x6b, 0xb0, 0x39, 0x91, 0xf4,
	0x2a, 0x4e, 0x62, 0xe9, 0x96, 0x75, 0xf0, 0x11, 0x40, 0x1e, 0x54, 0x16, 0x52, 0x66, 0x6e, 0x45,
	0xdf, 0xfb, 0x0b, 0x7d, 0x01, 0x52, 0x66, 0x2a, 0xfd, 0xc5, 0x01, 0xd6, 0x31, 0xf4, 0x15, 0xd8,
	0xbf, 0xd2, 0xb9, 0x60, 0xe1, 0x2d, 0x95, 0x6e, 0x55, 0x13, 0x0f, 0x15, 0xf1, 0x66, 0x03, 0x1a,
	0xf6, 0x23, 0xeb, 0xb4, 0x06, 0x15, 0xb9, 0xce, 0xa8, 0xf7, 0xaf, 0x05, 0x68, 0xf4, 0x40, 0xc3,
	0x95, 0xa4, 0x45, 0x45, 0x6f, 0xa1, 0x1c, 0x26, 0x79, 0x27, 0xda, 0x79, 0xb3, 0xf7, 0x49, 0xfd,
	0x21, 0x4b, 0x12, 0x92, 0x46, 0x58, 0x51, 0x51, 0x0f, 0x1a, 0x7a, 0xe0, 0x42, 0xb6, 0x34, 0xcd,
	0xd2, 0xb5, 0xce, 0x0c, 0x86, 0xb7, 0x51, 0x75, 0x17, 0x21, 0x59, 0x2e, 0xe7, 0x24, 0xbc, 0xd5,
	0x72, 0x6d, 0xbc, 0xf5, 0x55, 0x2c, 0x5a, 0x71, 0x22, 0x63, 0x96, 0x6a, 0xc5, 0x2d, 0xbc, 0xf5,
	0xd1, 0xa7, 0x66, 0x60, 0x1b, 0x5a, 0xe0, 0x5e, 0x6f, 0x74, 0xd0, 0x3b, 0x86, 0xba, 0x29, 0x0b,
	0xd9, 0x50, 0xf5, 0x83, 0x13, 0x1c, 0x38, 0x07, 0x6a, 0x06, 0xfc, 0x60, 0x3a, 0x73, 0x2c, 0xef,
	0x6f, 0x0b, 0xfe, 0xf7, 0x44, 0x8a, 0xe9, 0x3c, 0x82, 0x0a, 0x89, 0x22, 0xae, 0x15, 0xdb, 0x58,
	0xdb, 0xe8, 0x9b, 0x9d, 0xe9, 0xdb, 0xbf, 0x87, 0xfc, 0xf0, 0xce, 0xec, 0xa1, 0x23, 0xa8, 0x52,
	0xce, 0x19, 0x37, 0xea, 0x72, 0x47, 0x65, 0x88, 0x88, 0x24, 0x5a, 0x96, 0x8d, 0xb5, 0xed, 0xbd,
	0x2e, 0x4e, 0xa9, 0x7f, 0x3d, 0x1c, 0xe6, 0xb5, 0xbe, 0x3b, 0x19, 0x5f, 0x39, 0x96, 0xf7, 0xa7,
	0x05, 0xad, 0x27, 0x2d, 0x44, 0xaf, 0xa0, 0x26, 0xc2, 0x05, 0x4d, 0xa8, 0xa9, 0xd3, 0x78, 0xea,
	0xdb, 0x0b, 0x26, 0xa4, 0xae, 0xd3, 0xc6, 0xda, 0x56, 0x58, 0x46, 0xe4, 0xc2, 0x14, 0xa1, 0x6d,
	0xe4, 0x42, 0x5d, 0xc6, 0x09, 0x65, 0x2b, 0x69, 0x6e, 0x77, 0xe3, 0x2a, 0xf6, 0x9c, 0x45, 0x6b,
	0x3d, 0x3d, 0x36, 0xd6, 0x36, 0x3a, 0x06, 0x48, 0xc8, 0x03, 0xa6, 0x62, 0xb5, 0x94, 0xc2, 0xad,
	0x75, 0xad, 0x5e, 0x15, 0x17, 0x10, 0xef, 0xaf, 0x32, 0x34, 0x36, 0xb3, 0x88, 0x1c, 0x28, 0xaf,
	0xf8, 0xd2, 0xd4, 0xa5, 0x4c, 0x55, 0x6c, 0x42, 0xe5, 0x82, 0x45, 0xa6, 0x2c, 0xe3, 0x6d, 0x53,
	0x95, 0x0b, 0xa9, 0x0a, 0x85, 0xd5, 0x9f, 0x16, 0x76, 0x04, 0x55, 0x35, 0xe3, 0x03, 0xdd, 0xf6,
	0x06, 0xce, 0x1d, 0xf4, 0x06, 0x9c, 0x28, 0x16, 0x64, 0xbe, 0xa4, 0x98, 0x46, 0x31, 0xa7, 0xa1,
	0x14, 0xae, 0xad, 0x09, 0x7b, 0x38, 0xfa, 0x02, 0x0e, 0x0d, 0x76, 0x49, 0x69, 0x46, 0x96, 0xf1,
	0x3d, 0x15, 0x2e, 0x68, 0xf2, 0x7e, 0x00, 0xf5, 0x01, 0x19, 0x70, 0xc8, 0x92, 0x8c, 0x53, 0x21,
	0xd4, 0x2c, 0x36, 0x35, 0xfd, 0x99, 0x88, 0x7a, 0xbd, 0x19, 0x67, 0x0f, 0xeb, 0x13, 0x35, 0x3d,
	0x2f, 0xb4, 0xa4, 0x47, 0x40, 0x6d, 0x86, 0x84, 0x3c, 0x8c, 0x23, 0x75, 0x26, 0x4d, 0xdd, 0x56,
	0xbe, 0x19, 0x0a, 0x90, 0x52, 0xbe, 0xa0, 0x24, 0xa2, 0x5c, 0xb8, 0xed, 0x6e, 0xb9, 0x67, 0xe3,
	0x8d, 0x6b, 0xb6, 0x8a, 0xa4, 0xa9, 0x0c, 0xd6, 0x19, 0x75, 0x5f, 0xea, 0x6f, 0x17, 0xa1, 0x9d,
	0x06, 0x39, 0xbb, 0x0d, 0x7a, 0xb3, 0x86, 0xc6, 0xe6, 0xfd, 0xa9, 0xe5, 0x77, 0x3d, 0xb9, 0x9c,
	0x4c, 0x6f, 0x26, 0xf9, 0x8c, 0x5d, 0x04, 0xc1, 0xcc, 0xb1, 0x50, 0x1d, 0xca, 0xc1, 0x70, 0xe6,
	0x94, 0x94, 0x71, 0x7d, 0x36, 0x73, 0xca, 0xea, 0xd9, 0xe0, 0xd1, 0xd9, 0xd8, 0x77, 0x2a, 0xca,
	0xfc, 0xf1, 0x67, 0xff, 0xa7, 0x2b, 0xa7, 0xaa, 0xcd, 0xe9, 0xe4, 0x7c, 0xea, 0xd4, 0xd4, 0xe1,
	0x73, 0x3c, 0x1b, 0x3a, 0x75, 0x04, 0x50, 0x0b, 0x2e, 0xf0, 0xf8, 0x5d, 0xe0, 0x34, 0x50, 0x0b,
	0xec, 0x9b, 0xd1, 0xa9, 0x3f, 0x1d, 0x5e, 0x8e, 0x02, 0xc7, 0x1e, 0xfc, 0x61, 0x41, 0xdb, 0xcf,
	0x7f, 0x44, 0x3e, 0xe5, 0xf7, 0x71, 0x48, 0xd1, 0x0f, 0xd0, 0x2c, 0x3c, 0x1e, 0xf4, 0xea, 0xf9,
	0xad, 0xd2, 0xf9, 0xe8, 0x03, 0xaf, 0xcc, 0x3b, 0x40, 0xdf, 0x83, 0xbd, 0xdd, 0xd9, 0xe8, 0xe8,
	0xb9, 0x7f, 0x41, 0xe7, 0xff, 0x3b, 0xe8, 0xe6, 0x6c, 0xcf, 0x7a, 0x6b, 0xcd, 0x6b, 0x7a, 0x07,
	0x7d, 0xfd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xa7, 0xca, 0xa7, 0x23, 0x07, 0x00, 0x00,
}
