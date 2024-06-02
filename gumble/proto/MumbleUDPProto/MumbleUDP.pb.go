// Copyright 2022-2023 The Mumble Developers. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file at the root of the
// Mumble source tree or at <https://www.mumble.info/LICENSE>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: MumbleUDP.proto

package MumbleUDPProto

import (
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

type Audio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Header:
	//
	//	*Audio_Target
	//	*Audio_Context
	Header isAudio_Header `protobuf_oneof:"Header"`
	// The session of the client (sender) this audio was originally sent from. This field is not required when sending
	// audio to the server, but will always be set when receiving audio from the server.
	SenderSession uint32 `protobuf:"varint,3,opt,name=sender_session,json=senderSession,proto3" json:"sender_session,omitempty"`
	// The number of the first contained audio frame (indicating the position of that frame in the overall audio stream)
	FrameNumber uint64 `protobuf:"varint,4,opt,name=frame_number,json=frameNumber,proto3" json:"frame_number,omitempty"`
	// The actual voice data payload in the Opus format.
	OpusData []byte `protobuf:"bytes,5,opt,name=opus_data,json=opusData,proto3" json:"opus_data,omitempty"`
	// Optional positional data indicating the speaker's position in a virtual world (in meters). This "list" is really
	// expected to be an array of size 3 containing the X, Y and Z coordinates of the position (in that order).
	PositionalData []float32 `protobuf:"fixed32,6,rep,packed,name=positional_data,json=positionalData,proto3" json:"positional_data,omitempty"`
	// A volume adjustment determined by the server for this audio packet. It is up to the client to apply this adjustment to
	// the resulting audio (or not). Note: A value of 0 means that this field is unset.
	VolumeAdjustment float32 `protobuf:"fixed32,7,opt,name=volume_adjustment,json=volumeAdjustment,proto3" json:"volume_adjustment,omitempty"`
	// A flag indicating whether this audio packet represents the end of transmission for the current audio stream
	IsTerminator bool `protobuf:"varint,16,opt,name=is_terminator,json=isTerminator,proto3" json:"is_terminator,omitempty"`
}

func (x *Audio) Reset() {
	*x = Audio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MumbleUDP_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Audio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Audio) ProtoMessage() {}

func (x *Audio) ProtoReflect() protoreflect.Message {
	mi := &file_MumbleUDP_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Audio.ProtoReflect.Descriptor instead.
func (*Audio) Descriptor() ([]byte, []int) {
	return file_MumbleUDP_proto_rawDescGZIP(), []int{0}
}

func (m *Audio) GetHeader() isAudio_Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (x *Audio) GetTarget() uint32 {
	if x, ok := x.GetHeader().(*Audio_Target); ok {
		return x.Target
	}
	return 0
}

func (x *Audio) GetContext() uint32 {
	if x, ok := x.GetHeader().(*Audio_Context); ok {
		return x.Context
	}
	return 0
}

func (x *Audio) GetSenderSession() uint32 {
	if x != nil {
		return x.SenderSession
	}
	return 0
}

func (x *Audio) GetFrameNumber() uint64 {
	if x != nil {
		return x.FrameNumber
	}
	return 0
}

func (x *Audio) GetOpusData() []byte {
	if x != nil {
		return x.OpusData
	}
	return nil
}

func (x *Audio) GetPositionalData() []float32 {
	if x != nil {
		return x.PositionalData
	}
	return nil
}

func (x *Audio) GetVolumeAdjustment() float32 {
	if x != nil {
		return x.VolumeAdjustment
	}
	return 0
}

func (x *Audio) GetIsTerminator() bool {
	if x != nil {
		return x.IsTerminator
	}
	return false
}

type isAudio_Header interface {
	isAudio_Header()
}

type Audio_Target struct {
	// When this audio is sent by the client to the server, this is set to the target of the audio data. This target
	// is a number in the range [0, 2^{32} - 1], where 0 means "normal talking", 2^{5} - 1 means "server loopback"
	// and all other targets are understood as shout/whisper targets that have previously been registered via a
	// VoiceTarget message (via TCP).
	Target uint32 `protobuf:"varint,1,opt,name=target,proto3,oneof"`
}

type Audio_Context struct {
	// When this audio is sent by the server to the client, this indicates the context in which the audio has been sent.
	// 0: Normal speech
	// 1: Shout to channel
	// 2: Whisper to user
	// 3: Received via channel listener
	Context uint32 `protobuf:"varint,2,opt,name=context,proto3,oneof"`
}

func (*Audio_Target) isAudio_Header() {}

func (*Audio_Context) isAudio_Header() {}

// *
// Ping message for checking UDP connectivity (and roundtrip ping) and potentially obtaining further server
// details (e.g. version).
type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timestamp as encoded by the client. A server is not supposed to attempt to decode or modify this field. Therefore,
	// clients may choose an arbitrary format for this timestamp (as long as it fits into a uint64 field).
	Timestamp uint64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// A flag set by the sending client, if it wants to obtain additional information about the server.
	RequestExtendedInformation bool `protobuf:"varint,2,opt,name=request_extended_information,json=requestExtendedInformation,proto3" json:"request_extended_information,omitempty"`
	// The version of the server in the new version format.
	// The new protobuf Ping packet introduced with 1.5 drops support for the legacy version format
	// since both server and client have to support this new format.
	// (See https://github.com/mumble-voip/mumble/issues/5827)
	ServerVersionV2 uint64 `protobuf:"varint,3,opt,name=server_version_v2,json=serverVersionV2,proto3" json:"server_version_v2,omitempty"`
	// The amount of users currently connected to the server
	UserCount uint32 `protobuf:"varint,4,opt,name=user_count,json=userCount,proto3" json:"user_count,omitempty"`
	// The maximum amount of users permitted on this server
	MaxUserCount uint32 `protobuf:"varint,5,opt,name=max_user_count,json=maxUserCount,proto3" json:"max_user_count,omitempty"`
	// The maximum bandwidth each user is allowed to use for sending audio to the server
	MaxBandwidthPerUser uint32 `protobuf:"varint,6,opt,name=max_bandwidth_per_user,json=maxBandwidthPerUser,proto3" json:"max_bandwidth_per_user,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MumbleUDP_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_MumbleUDP_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_MumbleUDP_proto_rawDescGZIP(), []int{1}
}

func (x *Ping) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Ping) GetRequestExtendedInformation() bool {
	if x != nil {
		return x.RequestExtendedInformation
	}
	return false
}

func (x *Ping) GetServerVersionV2() uint64 {
	if x != nil {
		return x.ServerVersionV2
	}
	return 0
}

func (x *Ping) GetUserCount() uint32 {
	if x != nil {
		return x.UserCount
	}
	return 0
}

func (x *Ping) GetMaxUserCount() uint32 {
	if x != nil {
		return x.MaxUserCount
	}
	return 0
}

func (x *Ping) GetMaxBandwidthPerUser() uint32 {
	if x != nil {
		return x.MaxBandwidthPerUser
	}
	return 0
}

var File_MumbleUDP_proto protoreflect.FileDescriptor

var file_MumbleUDP_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x4d, 0x75, 0x6d, 0x62, 0x6c, 0x65, 0x55, 0x44, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x4d, 0x75, 0x6d, 0x62, 0x6c, 0x65, 0x55, 0x44, 0x50, 0x22, 0xa9, 0x02, 0x0a,
	0x05, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x18, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x12, 0x1a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x75, 0x73, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6f, 0x70, 0x75, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0e, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x11,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x61, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x41,
	0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x73, 0x5f,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x69, 0x73, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x08,
	0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x8c, 0x02, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x40, 0x0a, 0x1c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x32, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e,
	0x6d, 0x61, 0x78, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x33, 0x0a, 0x16, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x13, 0x6d, 0x61, 0x78, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x50, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x42, 0x02, 0x48, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_MumbleUDP_proto_rawDescOnce sync.Once
	file_MumbleUDP_proto_rawDescData = file_MumbleUDP_proto_rawDesc
)

func file_MumbleUDP_proto_rawDescGZIP() []byte {
	file_MumbleUDP_proto_rawDescOnce.Do(func() {
		file_MumbleUDP_proto_rawDescData = protoimpl.X.CompressGZIP(file_MumbleUDP_proto_rawDescData)
	})
	return file_MumbleUDP_proto_rawDescData
}

var file_MumbleUDP_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_MumbleUDP_proto_goTypes = []interface{}{
	(*Audio)(nil), // 0: MumbleUDP.Audio
	(*Ping)(nil),  // 1: MumbleUDP.Ping
}
var file_MumbleUDP_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MumbleUDP_proto_init() }
func file_MumbleUDP_proto_init() {
	if File_MumbleUDP_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MumbleUDP_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Audio); i {
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
		file_MumbleUDP_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
	file_MumbleUDP_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Audio_Target)(nil),
		(*Audio_Context)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MumbleUDP_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MumbleUDP_proto_goTypes,
		DependencyIndexes: file_MumbleUDP_proto_depIdxs,
		MessageInfos:      file_MumbleUDP_proto_msgTypes,
	}.Build()
	File_MumbleUDP_proto = out.File
	file_MumbleUDP_proto_rawDesc = nil
	file_MumbleUDP_proto_goTypes = nil
	file_MumbleUDP_proto_depIdxs = nil
}