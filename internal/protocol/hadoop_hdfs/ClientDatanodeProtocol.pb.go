// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ClientDatanodeProtocol.proto

package hadoop_hdfs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import hadoop_common "github.com/stubey/hdfs/v2/internal/protocol/hadoop_common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// *
// block - block for which visible length is requested
type GetReplicaVisibleLengthRequestProto struct {
	Block            *ExtendedBlockProto `protobuf:"bytes,1,req,name=block" json:"block,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *GetReplicaVisibleLengthRequestProto) Reset()         { *m = GetReplicaVisibleLengthRequestProto{} }
func (m *GetReplicaVisibleLengthRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetReplicaVisibleLengthRequestProto) ProtoMessage()    {}
func (*GetReplicaVisibleLengthRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{0}
}

func (m *GetReplicaVisibleLengthRequestProto) GetBlock() *ExtendedBlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

// *
// length - visible length of the block
type GetReplicaVisibleLengthResponseProto struct {
	Length           *uint64 `protobuf:"varint,1,req,name=length" json:"length,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetReplicaVisibleLengthResponseProto) Reset()         { *m = GetReplicaVisibleLengthResponseProto{} }
func (m *GetReplicaVisibleLengthResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetReplicaVisibleLengthResponseProto) ProtoMessage()    {}
func (*GetReplicaVisibleLengthResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{1}
}

func (m *GetReplicaVisibleLengthResponseProto) GetLength() uint64 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

// *
// void request
type RefreshNamenodesRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RefreshNamenodesRequestProto) Reset()                    { *m = RefreshNamenodesRequestProto{} }
func (m *RefreshNamenodesRequestProto) String() string            { return proto.CompactTextString(m) }
func (*RefreshNamenodesRequestProto) ProtoMessage()               {}
func (*RefreshNamenodesRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

// *
// void response
type RefreshNamenodesResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RefreshNamenodesResponseProto) Reset()                    { *m = RefreshNamenodesResponseProto{} }
func (m *RefreshNamenodesResponseProto) String() string            { return proto.CompactTextString(m) }
func (*RefreshNamenodesResponseProto) ProtoMessage()               {}
func (*RefreshNamenodesResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

// *
// blockPool - block pool to be deleted
// force - if false, delete the block pool only if it is empty.
//         if true, delete the block pool even if it has blocks.
type DeleteBlockPoolRequestProto struct {
	BlockPool        *string `protobuf:"bytes,1,req,name=blockPool" json:"blockPool,omitempty"`
	Force            *bool   `protobuf:"varint,2,req,name=force" json:"force,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeleteBlockPoolRequestProto) Reset()                    { *m = DeleteBlockPoolRequestProto{} }
func (m *DeleteBlockPoolRequestProto) String() string            { return proto.CompactTextString(m) }
func (*DeleteBlockPoolRequestProto) ProtoMessage()               {}
func (*DeleteBlockPoolRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

func (m *DeleteBlockPoolRequestProto) GetBlockPool() string {
	if m != nil && m.BlockPool != nil {
		return *m.BlockPool
	}
	return ""
}

func (m *DeleteBlockPoolRequestProto) GetForce() bool {
	if m != nil && m.Force != nil {
		return *m.Force
	}
	return false
}

// *
// void response
type DeleteBlockPoolResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DeleteBlockPoolResponseProto) Reset()                    { *m = DeleteBlockPoolResponseProto{} }
func (m *DeleteBlockPoolResponseProto) String() string            { return proto.CompactTextString(m) }
func (*DeleteBlockPoolResponseProto) ProtoMessage()               {}
func (*DeleteBlockPoolResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

// *
// Gets the file information where block and its metadata is stored
// block - block for which path information is being requested
// token - block token
//
// This message is deprecated in favor of file descriptor passing.
type GetBlockLocalPathInfoRequestProto struct {
	Block            *ExtendedBlockProto       `protobuf:"bytes,1,req,name=block" json:"block,omitempty"`
	Token            *hadoop_common.TokenProto `protobuf:"bytes,2,req,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *GetBlockLocalPathInfoRequestProto) Reset()         { *m = GetBlockLocalPathInfoRequestProto{} }
func (m *GetBlockLocalPathInfoRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetBlockLocalPathInfoRequestProto) ProtoMessage()    {}
func (*GetBlockLocalPathInfoRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{6}
}

func (m *GetBlockLocalPathInfoRequestProto) GetBlock() *ExtendedBlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *GetBlockLocalPathInfoRequestProto) GetToken() *hadoop_common.TokenProto {
	if m != nil {
		return m.Token
	}
	return nil
}

// *
// block - block for which file path information is being returned
// localPath - file path where the block data is stored
// localMetaPath - file path where the block meta data is stored
//
// This message is deprecated in favor of file descriptor passing.
type GetBlockLocalPathInfoResponseProto struct {
	Block            *ExtendedBlockProto `protobuf:"bytes,1,req,name=block" json:"block,omitempty"`
	LocalPath        *string             `protobuf:"bytes,2,req,name=localPath" json:"localPath,omitempty"`
	LocalMetaPath    *string             `protobuf:"bytes,3,req,name=localMetaPath" json:"localMetaPath,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *GetBlockLocalPathInfoResponseProto) Reset()         { *m = GetBlockLocalPathInfoResponseProto{} }
func (m *GetBlockLocalPathInfoResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetBlockLocalPathInfoResponseProto) ProtoMessage()    {}
func (*GetBlockLocalPathInfoResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{7}
}

func (m *GetBlockLocalPathInfoResponseProto) GetBlock() *ExtendedBlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *GetBlockLocalPathInfoResponseProto) GetLocalPath() string {
	if m != nil && m.LocalPath != nil {
		return *m.LocalPath
	}
	return ""
}

func (m *GetBlockLocalPathInfoResponseProto) GetLocalMetaPath() string {
	if m != nil && m.LocalMetaPath != nil {
		return *m.LocalMetaPath
	}
	return ""
}

// *
// forUpgrade - if true, clients are advised to wait for restart and quick
//              upgrade restart is instrumented. Otherwise, datanode does
//              the regular shutdown.
type ShutdownDatanodeRequestProto struct {
	ForUpgrade       *bool  `protobuf:"varint,1,req,name=forUpgrade" json:"forUpgrade,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ShutdownDatanodeRequestProto) Reset()                    { *m = ShutdownDatanodeRequestProto{} }
func (m *ShutdownDatanodeRequestProto) String() string            { return proto.CompactTextString(m) }
func (*ShutdownDatanodeRequestProto) ProtoMessage()               {}
func (*ShutdownDatanodeRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{8} }

func (m *ShutdownDatanodeRequestProto) GetForUpgrade() bool {
	if m != nil && m.ForUpgrade != nil {
		return *m.ForUpgrade
	}
	return false
}

type ShutdownDatanodeResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ShutdownDatanodeResponseProto) Reset()                    { *m = ShutdownDatanodeResponseProto{} }
func (m *ShutdownDatanodeResponseProto) String() string            { return proto.CompactTextString(m) }
func (*ShutdownDatanodeResponseProto) ProtoMessage()               {}
func (*ShutdownDatanodeResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{9} }

// * Tell datanode to evict active clients that are writing
type EvictWritersRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *EvictWritersRequestProto) Reset()                    { *m = EvictWritersRequestProto{} }
func (m *EvictWritersRequestProto) String() string            { return proto.CompactTextString(m) }
func (*EvictWritersRequestProto) ProtoMessage()               {}
func (*EvictWritersRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{10} }

type EvictWritersResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *EvictWritersResponseProto) Reset()                    { *m = EvictWritersResponseProto{} }
func (m *EvictWritersResponseProto) String() string            { return proto.CompactTextString(m) }
func (*EvictWritersResponseProto) ProtoMessage()               {}
func (*EvictWritersResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{11} }

// *
// Ping datanode for liveness and quick info
type GetDatanodeInfoRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetDatanodeInfoRequestProto) Reset()                    { *m = GetDatanodeInfoRequestProto{} }
func (m *GetDatanodeInfoRequestProto) String() string            { return proto.CompactTextString(m) }
func (*GetDatanodeInfoRequestProto) ProtoMessage()               {}
func (*GetDatanodeInfoRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{12} }

type GetDatanodeInfoResponseProto struct {
	LocalInfo        *DatanodeLocalInfoProto `protobuf:"bytes,1,req,name=localInfo" json:"localInfo,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *GetDatanodeInfoResponseProto) Reset()                    { *m = GetDatanodeInfoResponseProto{} }
func (m *GetDatanodeInfoResponseProto) String() string            { return proto.CompactTextString(m) }
func (*GetDatanodeInfoResponseProto) ProtoMessage()               {}
func (*GetDatanodeInfoResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{13} }

func (m *GetDatanodeInfoResponseProto) GetLocalInfo() *DatanodeLocalInfoProto {
	if m != nil {
		return m.LocalInfo
	}
	return nil
}

type TriggerBlockReportRequestProto struct {
	Incremental      *bool  `protobuf:"varint,1,req,name=incremental" json:"incremental,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TriggerBlockReportRequestProto) Reset()                    { *m = TriggerBlockReportRequestProto{} }
func (m *TriggerBlockReportRequestProto) String() string            { return proto.CompactTextString(m) }
func (*TriggerBlockReportRequestProto) ProtoMessage()               {}
func (*TriggerBlockReportRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{14} }

func (m *TriggerBlockReportRequestProto) GetIncremental() bool {
	if m != nil && m.Incremental != nil {
		return *m.Incremental
	}
	return false
}

type TriggerBlockReportResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *TriggerBlockReportResponseProto) Reset()         { *m = TriggerBlockReportResponseProto{} }
func (m *TriggerBlockReportResponseProto) String() string { return proto.CompactTextString(m) }
func (*TriggerBlockReportResponseProto) ProtoMessage()    {}
func (*TriggerBlockReportResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{15}
}

type GetBalancerBandwidthRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetBalancerBandwidthRequestProto) Reset()         { *m = GetBalancerBandwidthRequestProto{} }
func (m *GetBalancerBandwidthRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetBalancerBandwidthRequestProto) ProtoMessage()    {}
func (*GetBalancerBandwidthRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{16}
}

// *
// bandwidth - balancer bandwidth value of the datanode.
type GetBalancerBandwidthResponseProto struct {
	Bandwidth        *uint64 `protobuf:"varint,1,req,name=bandwidth" json:"bandwidth,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetBalancerBandwidthResponseProto) Reset()         { *m = GetBalancerBandwidthResponseProto{} }
func (m *GetBalancerBandwidthResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetBalancerBandwidthResponseProto) ProtoMessage()    {}
func (*GetBalancerBandwidthResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{17}
}

func (m *GetBalancerBandwidthResponseProto) GetBandwidth() uint64 {
	if m != nil && m.Bandwidth != nil {
		return *m.Bandwidth
	}
	return 0
}

func init() {
	proto.RegisterType((*GetReplicaVisibleLengthRequestProto)(nil), "hadoop.hdfs.GetReplicaVisibleLengthRequestProto")
	proto.RegisterType((*GetReplicaVisibleLengthResponseProto)(nil), "hadoop.hdfs.GetReplicaVisibleLengthResponseProto")
	proto.RegisterType((*RefreshNamenodesRequestProto)(nil), "hadoop.hdfs.RefreshNamenodesRequestProto")
	proto.RegisterType((*RefreshNamenodesResponseProto)(nil), "hadoop.hdfs.RefreshNamenodesResponseProto")
	proto.RegisterType((*DeleteBlockPoolRequestProto)(nil), "hadoop.hdfs.DeleteBlockPoolRequestProto")
	proto.RegisterType((*DeleteBlockPoolResponseProto)(nil), "hadoop.hdfs.DeleteBlockPoolResponseProto")
	proto.RegisterType((*GetBlockLocalPathInfoRequestProto)(nil), "hadoop.hdfs.GetBlockLocalPathInfoRequestProto")
	proto.RegisterType((*GetBlockLocalPathInfoResponseProto)(nil), "hadoop.hdfs.GetBlockLocalPathInfoResponseProto")
	proto.RegisterType((*ShutdownDatanodeRequestProto)(nil), "hadoop.hdfs.ShutdownDatanodeRequestProto")
	proto.RegisterType((*ShutdownDatanodeResponseProto)(nil), "hadoop.hdfs.ShutdownDatanodeResponseProto")
	proto.RegisterType((*EvictWritersRequestProto)(nil), "hadoop.hdfs.EvictWritersRequestProto")
	proto.RegisterType((*EvictWritersResponseProto)(nil), "hadoop.hdfs.EvictWritersResponseProto")
	proto.RegisterType((*GetDatanodeInfoRequestProto)(nil), "hadoop.hdfs.GetDatanodeInfoRequestProto")
	proto.RegisterType((*GetDatanodeInfoResponseProto)(nil), "hadoop.hdfs.GetDatanodeInfoResponseProto")
	proto.RegisterType((*TriggerBlockReportRequestProto)(nil), "hadoop.hdfs.TriggerBlockReportRequestProto")
	proto.RegisterType((*TriggerBlockReportResponseProto)(nil), "hadoop.hdfs.TriggerBlockReportResponseProto")
	proto.RegisterType((*GetBalancerBandwidthRequestProto)(nil), "hadoop.hdfs.GetBalancerBandwidthRequestProto")
	proto.RegisterType((*GetBalancerBandwidthResponseProto)(nil), "hadoop.hdfs.GetBalancerBandwidthResponseProto")
}

func init() { proto.RegisterFile("ClientDatanodeProtocol.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 804 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4b, 0x6f, 0xdb, 0x46,
	0x10, 0x06, 0xdd, 0xba, 0xb0, 0x46, 0x7d, 0x61, 0xe1, 0xb6, 0x32, 0x2d, 0xd9, 0x32, 0x6d, 0x17,
	0x76, 0x1f, 0x74, 0x2b, 0xc0, 0x3d, 0x1a, 0xb0, 0x6a, 0xc3, 0x28, 0xaa, 0x16, 0x2e, 0xe5, 0x26,
	0x97, 0xe4, 0xb0, 0x22, 0x47, 0xd4, 0xc2, 0x14, 0x97, 0x59, 0xae, 0x6c, 0xe7, 0x12, 0x20, 0xb7,
	0x00, 0x41, 0x0e, 0xf9, 0x05, 0xc9, 0x4f, 0x0d, 0xb8, 0x92, 0xa2, 0x5d, 0x8a, 0xa2, 0x95, 0x20,
	0x27, 0x91, 0x33, 0xdf, 0x3c, 0xf6, 0xdb, 0x99, 0x4f, 0x84, 0xfa, 0x9f, 0x11, 0xc3, 0x58, 0x9e,
	0x51, 0x49, 0x63, 0x1e, 0xe0, 0xa5, 0xe0, 0x92, 0xfb, 0x3c, 0x72, 0x93, 0xec, 0x81, 0x54, 0x07,
	0x34, 0xe0, 0x3c, 0x71, 0x07, 0x41, 0x3f, 0xb5, 0xbf, 0xee, 0xa2, 0x3f, 0x12, 0x4c, 0x3e, 0x1d,
	0x3b, 0x6d, 0xc8, 0xac, 0x93, 0xe7, 0x86, 0x87, 0x3e, 0x8f, 0xfb, 0x2c, 0x1c, 0x09, 0x2a, 0x19,
	0x8f, 0xcd, 0x3c, 0xce, 0x23, 0xd8, 0xbd, 0x40, 0xe9, 0x61, 0x12, 0x31, 0x9f, 0x3e, 0x60, 0x29,
	0xeb, 0x45, 0xd8, 0xc1, 0x38, 0x94, 0x03, 0x0f, 0x9f, 0x8c, 0x30, 0x95, 0x0a, 0x4f, 0x8e, 0x61,
	0xb5, 0x17, 0x71, 0xff, 0xba, 0x66, 0x35, 0x57, 0x0e, 0xaa, 0xad, 0x6d, 0x57, 0x2b, 0xef, 0x9e,
	0xdf, 0x49, 0x8c, 0x03, 0x0c, 0xda, 0x19, 0x42, 0xe1, 0xbd, 0x31, 0xda, 0x39, 0x81, 0xbd, 0x85,
	0xd9, 0xd3, 0x84, 0xc7, 0xe9, 0xf8, 0x58, 0xe4, 0x7b, 0xf8, 0x22, 0x52, 0x66, 0x95, 0xff, 0x73,
	0x6f, 0xf2, 0xe6, 0x6c, 0x41, 0xdd, 0xc3, 0xbe, 0xc0, 0x74, 0xf0, 0x2f, 0x1d, 0x62, 0x46, 0x43,
	0xaa, 0xb7, 0xe5, 0x6c, 0x43, 0x63, 0xde, 0xaf, 0x25, 0x76, 0xfe, 0x83, 0xcd, 0x33, 0x8c, 0x50,
	0xe2, 0xb8, 0x37, 0xce, 0x23, 0xe3, 0x58, 0x75, 0xa8, 0xf4, 0xa6, 0x0e, 0x55, 0xba, 0xe2, 0xcd,
	0x0c, 0x64, 0x1d, 0x56, 0xfb, 0x5c, 0xf8, 0x58, 0x5b, 0x69, 0xae, 0x1c, 0xac, 0x79, 0xe3, 0x97,
	0xac, 0xa7, 0xb9, 0x94, 0x7a, 0xc9, 0x97, 0x16, 0xec, 0x5c, 0xa0, 0x54, 0xde, 0x0e, 0xf7, 0x69,
	0x74, 0x49, 0xe5, 0xe0, 0xaf, 0xb8, 0xcf, 0x3f, 0x01, 0xa1, 0xe4, 0x08, 0x56, 0x25, 0xbf, 0xc6,
	0x58, 0xb5, 0x54, 0x6d, 0x6d, 0x4c, 0xc3, 0x7c, 0x3e, 0x1c, 0xf2, 0xd8, 0xbd, 0xca, 0x7c, 0x93,
	0x00, 0x85, 0x73, 0xde, 0x58, 0xe0, 0x2c, 0xe8, 0x46, 0xbf, 0x80, 0x8f, 0x6c, 0xa7, 0x0e, 0x95,
	0x68, 0x9a, 0x54, 0xb5, 0x54, 0xf1, 0x66, 0x06, 0xb2, 0x07, 0x5f, 0xa9, 0x97, 0x7f, 0x50, 0x52,
	0x85, 0xf8, 0x4c, 0x21, 0x4c, 0xa3, 0x73, 0x02, 0xf5, 0xee, 0x60, 0x24, 0x03, 0x7e, 0x1b, 0x4f,
	0x67, 0xdd, 0x60, 0x6a, 0x0b, 0xa0, 0xcf, 0xc5, 0xff, 0x49, 0x28, 0x68, 0x80, 0xaa, 0xbf, 0x35,
	0x4f, 0xb3, 0x64, 0x33, 0x30, 0x1f, 0xaf, 0x5f, 0x88, 0x0d, 0xb5, 0xf3, 0x1b, 0xe6, 0xcb, 0x87,
	0x82, 0x49, 0x14, 0xe6, 0x00, 0x6d, 0xc2, 0x86, 0xe9, 0xd3, 0x03, 0x1b, 0xb0, 0x79, 0x81, 0xef,
	0x17, 0x30, 0x7f, 0x85, 0x0e, 0x85, 0xfa, 0x9c, 0x5b, 0xe7, 0xf4, 0x74, 0x42, 0x4e, 0xe6, 0x99,
	0xf0, 0xba, 0x6b, 0xf0, 0x3a, 0x0d, 0xed, 0x4c, 0x51, 0x63, 0x6e, 0x67, 0x51, 0x4e, 0x1b, 0xb6,
	0xae, 0x04, 0x0b, 0x43, 0x14, 0x8a, 0x7b, 0x0f, 0x13, 0x2e, 0xa4, 0xc1, 0x4e, 0x13, 0xaa, 0x2c,
	0xf6, 0x05, 0x0e, 0x31, 0x96, 0x34, 0x9a, 0xd0, 0xa3, 0x9b, 0x9c, 0x1d, 0xd8, 0x2e, 0xca, 0xa1,
	0x1f, 0xd4, 0x81, 0x66, 0x36, 0x23, 0x34, 0xa2, 0xb1, 0x8f, 0xa2, 0x4d, 0xe3, 0xe0, 0x96, 0x05,
	0xa6, 0x02, 0x38, 0xa7, 0xe3, 0xa9, 0x9e, 0xc7, 0xe8, 0x47, 0xce, 0xf6, 0x69, 0xea, 0x99, 0xac,
	0xf2, 0xcc, 0xd0, 0x7a, 0x5d, 0x85, 0x46, 0xb1, 0xa8, 0x75, 0x51, 0xdc, 0x30, 0x1f, 0xc9, 0x33,
	0xf8, 0x21, 0x2c, 0xd6, 0x0b, 0xf2, 0x9b, 0x41, 0xdd, 0x12, 0x9a, 0x65, 0xff, 0xbe, 0x5c, 0x84,
	0xde, 0x3f, 0x83, 0x6f, 0x45, 0x4e, 0x4f, 0xc8, 0xa1, 0x91, 0xa6, 0x4c, 0x8e, 0xec, 0x9f, 0xee,
	0x81, 0xea, 0xa5, 0xfa, 0xf0, 0x4d, 0x60, 0xca, 0x08, 0x39, 0x30, 0xa7, 0x63, 0xb1, 0x6e, 0xd9,
	0x87, 0xe5, 0x48, 0xbd, 0xce, 0x1d, 0x7c, 0x17, 0x16, 0xed, 0x3f, 0x71, 0xf3, 0xf4, 0x94, 0x2b,
	0x96, 0x7d, 0xb4, 0x0c, 0x3e, 0x47, 0x66, 0x9a, 0x5b, 0xcc, 0x1c, 0x99, 0x65, 0x7b, 0x9f, 0x23,
	0xb3, 0x74, 0xc5, 0xc9, 0x63, 0xf8, 0x12, 0xb5, 0x35, 0x26, 0xfb, 0xa6, 0x7e, 0x2d, 0xd8, 0x7e,
	0xfb, 0xc7, 0x12, 0x58, 0xee, 0xae, 0x42, 0x73, 0xd3, 0x73, 0x77, 0x55, 0x22, 0x13, 0xb9, 0xbb,
	0x2a, 0x55, 0x8c, 0xe7, 0x16, 0xd4, 0xd4, 0xfc, 0x1b, 0xff, 0xd8, 0x5d, 0x49, 0xe5, 0x28, 0x25,
	0x05, 0xe3, 0x5c, 0x00, 0x33, 0x4a, 0xb7, 0x96, 0x0c, 0xd1, 0x7b, 0x18, 0xc1, 0x7a, 0x2a, 0xa9,
	0xc8, 0x43, 0xc9, 0xaf, 0xe6, 0x75, 0x14, 0x40, 0x8c, 0xd2, 0xee, 0x12, 0x70, 0xbd, 0xec, 0x2b,
	0x0b, 0xea, 0x11, 0x4b, 0x75, 0x50, 0x2f, 0xca, 0x7c, 0x09, 0x0a, 0xc9, 0x30, 0x25, 0xc7, 0x46,
	0xc2, 0x4e, 0x09, 0xd4, 0xe8, 0xe3, 0x8f, 0x0f, 0x08, 0xd3, 0xfb, 0xe1, 0x40, 0xe4, 0x9c, 0x6a,
	0x92, 0x9f, 0x8d, 0x6c, 0xe5, 0xd2, 0x6c, 0xff, 0x72, 0x2f, 0x38, 0xc7, 0x7b, 0x58, 0xa0, 0xaf,
	0x39, 0xde, 0xef, 0x93, 0x69, 0xdb, 0x5d, 0x02, 0xae, 0x95, 0x6d, 0xff, 0x0d, 0xfb, 0x5c, 0x84,
	0x2e, 0x4d, 0xa8, 0x3f, 0x40, 0x23, 0x36, 0x31, 0x3e, 0x14, 0xdb, 0x0b, 0x3e, 0x47, 0xd5, 0x6f,
	0xfa, 0xc2, 0xb2, 0xde, 0x5a, 0xd6, 0xbb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x38, 0x2f, 0xb6,
	0xb3, 0x0a, 0x00, 0x00,
}
