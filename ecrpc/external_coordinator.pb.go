// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: ecrpc/external_coordinator.proto

package ecrpc

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// RegisterMissionControlRequest is the request message for registering mission
// control data.
type RegisterMissionControlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pairs []*PairHistory `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"`
}

func (x *RegisterMissionControlRequest) Reset() {
	*x = RegisterMissionControlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecrpc_external_coordinator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterMissionControlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterMissionControlRequest) ProtoMessage() {}

func (x *RegisterMissionControlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ecrpc_external_coordinator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterMissionControlRequest.ProtoReflect.Descriptor instead.
func (*RegisterMissionControlRequest) Descriptor() ([]byte, []int) {
	return file_ecrpc_external_coordinator_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterMissionControlRequest) GetPairs() []*PairHistory {
	if x != nil {
		return x.Pairs
	}
	return nil
}

// RegisterMissionControlResponse is the response message for registering
// mission control data.
type RegisterMissionControlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Success message indicating the number of pairs successfully registered
	// and stale pairs removed (if any).
	SuccessMessage string `protobuf:"bytes,1,opt,name=success_message,json=successMessage,proto3" json:"success_message,omitempty"`
}

func (x *RegisterMissionControlResponse) Reset() {
	*x = RegisterMissionControlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecrpc_external_coordinator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterMissionControlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterMissionControlResponse) ProtoMessage() {}

func (x *RegisterMissionControlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ecrpc_external_coordinator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterMissionControlResponse.ProtoReflect.Descriptor instead.
func (*RegisterMissionControlResponse) Descriptor() ([]byte, []int) {
	return file_ecrpc_external_coordinator_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterMissionControlResponse) GetSuccessMessage() string {
	if x != nil {
		return x.SuccessMessage
	}
	return ""
}

// QueryAggregatedMissionControlRequest is the request message for querying
// aggregated mission control data.
type QueryAggregatedMissionControlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryAggregatedMissionControlRequest) Reset() {
	*x = QueryAggregatedMissionControlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecrpc_external_coordinator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAggregatedMissionControlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAggregatedMissionControlRequest) ProtoMessage() {}

func (x *QueryAggregatedMissionControlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ecrpc_external_coordinator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAggregatedMissionControlRequest.ProtoReflect.Descriptor instead.
func (*QueryAggregatedMissionControlRequest) Descriptor() ([]byte, []int) {
	return file_ecrpc_external_coordinator_proto_rawDescGZIP(), []int{2}
}

// QueryAggregatedMissionControlResponse is the response message for querying
// aggregated mission control data.
//
// NOTE: This is the same message that is found in LND.
type QueryAggregatedMissionControlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pairs []*PairHistory `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"`
}

func (x *QueryAggregatedMissionControlResponse) Reset() {
	*x = QueryAggregatedMissionControlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecrpc_external_coordinator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAggregatedMissionControlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAggregatedMissionControlResponse) ProtoMessage() {}

func (x *QueryAggregatedMissionControlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ecrpc_external_coordinator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAggregatedMissionControlResponse.ProtoReflect.Descriptor instead.
func (*QueryAggregatedMissionControlResponse) Descriptor() ([]byte, []int) {
	return file_ecrpc_external_coordinator_proto_rawDescGZIP(), []int{3}
}

func (x *QueryAggregatedMissionControlResponse) GetPairs() []*PairHistory {
	if x != nil {
		return x.Pairs
	}
	return nil
}

// PairHistory contains the mission control state for a particular node pair.
type PairHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The source node pubkey of the pair.
	NodeFrom []byte `protobuf:"bytes,1,opt,name=node_from,json=nodeFrom,proto3" json:"node_from,omitempty"`
	// The destination node pubkey of the pair.
	NodeTo []byte `protobuf:"bytes,2,opt,name=node_to,json=nodeTo,proto3" json:"node_to,omitempty"`
	// History data for the pair.
	History *PairData `protobuf:"bytes,3,opt,name=history,proto3" json:"history,omitempty"`
}

func (x *PairHistory) Reset() {
	*x = PairHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecrpc_external_coordinator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PairHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PairHistory) ProtoMessage() {}

func (x *PairHistory) ProtoReflect() protoreflect.Message {
	mi := &file_ecrpc_external_coordinator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PairHistory.ProtoReflect.Descriptor instead.
func (*PairHistory) Descriptor() ([]byte, []int) {
	return file_ecrpc_external_coordinator_proto_rawDescGZIP(), []int{4}
}

func (x *PairHistory) GetNodeFrom() []byte {
	if x != nil {
		return x.NodeFrom
	}
	return nil
}

func (x *PairHistory) GetNodeTo() []byte {
	if x != nil {
		return x.NodeTo
	}
	return nil
}

func (x *PairHistory) GetHistory() *PairData {
	if x != nil {
		return x.History
	}
	return nil
}

// PairData contains the detailed history data for a node pair.
type PairData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time of last failure.
	FailTime int64 `protobuf:"varint,1,opt,name=fail_time,json=failTime,proto3" json:"fail_time,omitempty"`
	// Lowest amount that failed to forward rounded to whole sats. This may be
	// set to zero if the failure is independent of amount.
	FailAmtSat int64 `protobuf:"varint,2,opt,name=fail_amt_sat,json=failAmtSat,proto3" json:"fail_amt_sat,omitempty"`
	// Lowest amount that failed to forward in millisats. This may be set to
	// zero if the failure is independent of amount.
	FailAmtMsat int64 `protobuf:"varint,3,opt,name=fail_amt_msat,json=failAmtMsat,proto3" json:"fail_amt_msat,omitempty"`
	// Time of last success.
	SuccessTime int64 `protobuf:"varint,4,opt,name=success_time,json=successTime,proto3" json:"success_time,omitempty"`
	// Highest amount that we could successfully forward rounded to whole sats.
	SuccessAmtSat int64 `protobuf:"varint,5,opt,name=success_amt_sat,json=successAmtSat,proto3" json:"success_amt_sat,omitempty"`
	// Highest amount that we could successfully forward in millisats.
	SuccessAmtMsat int64 `protobuf:"varint,6,opt,name=success_amt_msat,json=successAmtMsat,proto3" json:"success_amt_msat,omitempty"`
}

func (x *PairData) Reset() {
	*x = PairData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecrpc_external_coordinator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PairData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PairData) ProtoMessage() {}

func (x *PairData) ProtoReflect() protoreflect.Message {
	mi := &file_ecrpc_external_coordinator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PairData.ProtoReflect.Descriptor instead.
func (*PairData) Descriptor() ([]byte, []int) {
	return file_ecrpc_external_coordinator_proto_rawDescGZIP(), []int{5}
}

func (x *PairData) GetFailTime() int64 {
	if x != nil {
		return x.FailTime
	}
	return 0
}

func (x *PairData) GetFailAmtSat() int64 {
	if x != nil {
		return x.FailAmtSat
	}
	return 0
}

func (x *PairData) GetFailAmtMsat() int64 {
	if x != nil {
		return x.FailAmtMsat
	}
	return 0
}

func (x *PairData) GetSuccessTime() int64 {
	if x != nil {
		return x.SuccessTime
	}
	return 0
}

func (x *PairData) GetSuccessAmtSat() int64 {
	if x != nil {
		return x.SuccessAmtSat
	}
	return 0
}

func (x *PairData) GetSuccessAmtMsat() int64 {
	if x != nil {
		return x.SuccessAmtMsat
	}
	return 0
}

var File_ecrpc_external_coordinator_proto protoreflect.FileDescriptor

var file_ecrpc_external_coordinator_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x63, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x65, 0x63, 0x72, 0x70, 0x63, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x1d, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x70, 0x61, 0x69, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x63, 0x72, 0x70, 0x63, 0x2e,
	0x50, 0x61, 0x69, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x05, 0x70, 0x61, 0x69,
	0x72, 0x73, 0x22, 0x49, 0x0a, 0x1e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a,
	0x24, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64,
	0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a, 0x25, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x65, 0x63, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73, 0x22, 0x6e, 0x0a, 0x0b, 0x50, 0x61, 0x69, 0x72,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x54, 0x6f, 0x12, 0x29, 0x0a,
	0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x65, 0x63, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0xe2, 0x01, 0x0a, 0x08, 0x50, 0x61, 0x69,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x6d, 0x74, 0x5f, 0x73,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x61, 0x69, 0x6c, 0x41, 0x6d,
	0x74, 0x53, 0x61, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x6d, 0x74,
	0x5f, 0x6d, 0x73, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x61, 0x69,
	0x6c, 0x41, 0x6d, 0x74, 0x4d, 0x73, 0x61, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x61, 0x6d, 0x74, 0x5f, 0x73, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x41, 0x6d, 0x74,
	0x53, 0x61, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x61,
	0x6d, 0x74, 0x5f, 0x6d, 0x73, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x41, 0x6d, 0x74, 0x4d, 0x73, 0x61, 0x74, 0x32, 0xcc, 0x02,
	0x0a, 0x13, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x8c, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x12, 0x24, 0x2e, 0x65, 0x63, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x65, 0x63, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x12, 0xa5, 0x01, 0x0a, 0x1d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x2b, 0x2e, 0x65, 0x63, 0x72, 0x70, 0x63, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x65, 0x63, 0x72, 0x70, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x42, 0x41, 0x5a, 0x3f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x67, 0x67, 0x69,
	0x65, 0x31, 0x39, 0x38, 0x34, 0x2f, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x64, 0x2d, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x66, 0x6f, 0x72, 0x2d, 0x4c, 0x4e, 0x44, 0x2f, 0x65, 0x63, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ecrpc_external_coordinator_proto_rawDescOnce sync.Once
	file_ecrpc_external_coordinator_proto_rawDescData = file_ecrpc_external_coordinator_proto_rawDesc
)

func file_ecrpc_external_coordinator_proto_rawDescGZIP() []byte {
	file_ecrpc_external_coordinator_proto_rawDescOnce.Do(func() {
		file_ecrpc_external_coordinator_proto_rawDescData = protoimpl.X.CompressGZIP(file_ecrpc_external_coordinator_proto_rawDescData)
	})
	return file_ecrpc_external_coordinator_proto_rawDescData
}

var file_ecrpc_external_coordinator_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ecrpc_external_coordinator_proto_goTypes = []interface{}{
	(*RegisterMissionControlRequest)(nil),         // 0: ecrpc.RegisterMissionControlRequest
	(*RegisterMissionControlResponse)(nil),        // 1: ecrpc.RegisterMissionControlResponse
	(*QueryAggregatedMissionControlRequest)(nil),  // 2: ecrpc.QueryAggregatedMissionControlRequest
	(*QueryAggregatedMissionControlResponse)(nil), // 3: ecrpc.QueryAggregatedMissionControlResponse
	(*PairHistory)(nil),                           // 4: ecrpc.PairHistory
	(*PairData)(nil),                              // 5: ecrpc.PairData
}
var file_ecrpc_external_coordinator_proto_depIdxs = []int32{
	4, // 0: ecrpc.RegisterMissionControlRequest.pairs:type_name -> ecrpc.PairHistory
	4, // 1: ecrpc.QueryAggregatedMissionControlResponse.pairs:type_name -> ecrpc.PairHistory
	5, // 2: ecrpc.PairHistory.history:type_name -> ecrpc.PairData
	0, // 3: ecrpc.ExternalCoordinator.RegisterMissionControl:input_type -> ecrpc.RegisterMissionControlRequest
	2, // 4: ecrpc.ExternalCoordinator.QueryAggregatedMissionControl:input_type -> ecrpc.QueryAggregatedMissionControlRequest
	1, // 5: ecrpc.ExternalCoordinator.RegisterMissionControl:output_type -> ecrpc.RegisterMissionControlResponse
	3, // 6: ecrpc.ExternalCoordinator.QueryAggregatedMissionControl:output_type -> ecrpc.QueryAggregatedMissionControlResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ecrpc_external_coordinator_proto_init() }
func file_ecrpc_external_coordinator_proto_init() {
	if File_ecrpc_external_coordinator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ecrpc_external_coordinator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterMissionControlRequest); i {
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
		file_ecrpc_external_coordinator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterMissionControlResponse); i {
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
		file_ecrpc_external_coordinator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAggregatedMissionControlRequest); i {
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
		file_ecrpc_external_coordinator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAggregatedMissionControlResponse); i {
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
		file_ecrpc_external_coordinator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PairHistory); i {
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
		file_ecrpc_external_coordinator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PairData); i {
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
			RawDescriptor: file_ecrpc_external_coordinator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ecrpc_external_coordinator_proto_goTypes,
		DependencyIndexes: file_ecrpc_external_coordinator_proto_depIdxs,
		MessageInfos:      file_ecrpc_external_coordinator_proto_msgTypes,
	}.Build()
	File_ecrpc_external_coordinator_proto = out.File
	file_ecrpc_external_coordinator_proto_rawDesc = nil
	file_ecrpc_external_coordinator_proto_goTypes = nil
	file_ecrpc_external_coordinator_proto_depIdxs = nil
}
