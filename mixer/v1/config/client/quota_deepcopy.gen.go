// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/v1/config/client/quota.proto

package client

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using QuotaSpec within kubernetes types, where deepcopy-gen is used.
func (in *QuotaSpec) DeepCopyInto(out *QuotaSpec) {
	proto.Merge(out, in)
}

// DeepCopyInto supports using QuotaSpecBinding within kubernetes types, where deepcopy-gen is used.
func (in *QuotaSpecBinding) DeepCopyInto(out *QuotaSpecBinding) {
	proto.Merge(out, in)
}
