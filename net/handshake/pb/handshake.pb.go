// Code generated by protoc-gen-gogo.
// source: handshake.proto
// DO NOT EDIT!

/*
Package handshake_pb is a generated protocol buffer package.

It is generated from these files:
	handshake.proto

It has these top-level messages:
	Handshake1
	Handshake3
*/
package handshake_pb

import proto "github.com/jbenet/go-ipfs/Godeps/_workspace/src/code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import mux "github.com/jbenet/go-ipfs/net/mux/mux.pb"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// Handshake1 is delivered _before_ the secure channel is initialized
type Handshake1 struct {
	// protocolVersion determines compatibility between peers
	ProtocolVersion *string `protobuf:"bytes,1,opt,name=protocolVersion" json:"protocolVersion,omitempty"`
	// agentVersion is like a UserAgent string in browsers, or client version in bittorrent
	// includes the client name and client. e.g.   "go-ipfs/0.1.0"
	AgentVersion     *string `protobuf:"bytes,2,opt,name=agentVersion" json:"agentVersion,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Handshake1) Reset()         { *m = Handshake1{} }
func (m *Handshake1) String() string { return proto.CompactTextString(m) }
func (*Handshake1) ProtoMessage()    {}

func (m *Handshake1) GetProtocolVersion() string {
	if m != nil && m.ProtocolVersion != nil {
		return *m.ProtocolVersion
	}
	return ""
}

func (m *Handshake1) GetAgentVersion() string {
	if m != nil && m.AgentVersion != nil {
		return *m.AgentVersion
	}
	return ""
}

// Handshake3 is delivered _after_ the secure channel is initialized
type Handshake3 struct {
	// listenAddrs are the multiaddrs this node listens for open connections on
	ListenAddrs      [][]byte `protobuf:"bytes,2,rep,name=listenAddrs" json:"listenAddrs,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Handshake3) Reset()         { *m = Handshake3{} }
func (m *Handshake3) String() string { return proto.CompactTextString(m) }
func (*Handshake3) ProtoMessage()    {}

func (m *Handshake3) GetListenAddrs() [][]byte {
	if m != nil {
		return m.ListenAddrs
	}
	return nil
}

func init() {
}
