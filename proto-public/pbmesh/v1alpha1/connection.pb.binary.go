// Code generated by protoc-gen-go-binary. DO NOT EDIT.
// source: pbmesh/v1alpha1/connection.proto

package meshv1alpha1

import (
	"google.golang.org/protobuf/proto"
)

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *ConnectionConfig) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *ConnectionConfig) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *InboundConnectionsConfig) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *InboundConnectionsConfig) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}
