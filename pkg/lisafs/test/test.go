// Copyright 2021 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package test holds testing utilites for lisafs.
package test

import (
	"math/rand"

	"gvisor.dev/gvisor/pkg/marshal/primitive"
)

// MsgSimple is a sample packed struct which can be used to test message passing.
//
// +marshal slice:Msg1Slice
type MsgSimple struct {
	A uint16
	B uint16
	C uint32
	D uint64
}

// Randomize randomizes the contents of m.
func (m *MsgSimple) Randomize() {
	m.A = uint16(rand.Uint32())
	m.B = uint16(rand.Uint32())
	m.C = rand.Uint32()
	m.D = rand.Uint64()
}

// MsgDynamic is a sample dynamic struct which can be used to test message passing.
//
// +marshal dynamic
type MsgDynamic struct {
	N   primitive.Uint32
	Arr []MsgSimple
}

// SizeBytes implements marshal.Marshallable.SizeBytes.
func (m *MsgDynamic) SizeBytes() int {
	return m.N.SizeBytes() +
		(int(m.N) * (*MsgSimple)(nil).SizeBytes())
}

// MarshalBytes implements marshal.Marshallable.MarshalBytes.
func (m *MsgDynamic) MarshalBytes(dst []byte) {
	m.N.MarshalUnsafe(dst)
	dst = dst[m.N.SizeBytes():]
	MarshalUnsafeMsg1Slice(m.Arr, dst)
}

// UnmarshalBytes implements marshal.Marshallable.UnmarshalBytes.
func (m *MsgDynamic) UnmarshalBytes(src []byte) {
	m.N.UnmarshalUnsafe(src)
	src = src[m.N.SizeBytes():]
	m.Arr = make([]MsgSimple, m.N)
	UnmarshalUnsafeMsg1Slice(m.Arr, src)
}

// Randomize randomizes the contents of m.
func (m *MsgDynamic) Randomize(arrLen int) {
	m.N = primitive.Uint32(arrLen)
	m.Arr = make([]MsgSimple, arrLen)
	for i := 0; i < arrLen; i++ {
		m.Arr[i].Randomize()
	}
}

// Version mimics p9.TVersion and p9.Rversion.
//
// +marshal dynamic
type Version struct {
	MSize   primitive.Uint32
	Version string
}

// SizeBytes implements marshal.Marshallable.SizeBytes.
func (v *Version) SizeBytes() int {
	return (*primitive.Uint32)(nil).SizeBytes() + (*primitive.Uint16)(nil).SizeBytes() + len(v.Version)
}

// MarshalBytes implements marshal.Marshallable.MarshalBytes.
func (v *Version) MarshalBytes(dst []byte) {
	v.MSize.MarshalUnsafe(dst)
	dst = dst[v.MSize.SizeBytes():]
	versionLen := primitive.Uint16(len(v.Version))
	versionLen.MarshalUnsafe(dst)
	dst = dst[versionLen.SizeBytes():]
	copy(dst, v.Version)
}

// UnmarshalBytes implements marshal.Marshallable.UnmarshalBytes.
func (v *Version) UnmarshalBytes(src []byte) {
	v.MSize.UnmarshalUnsafe(src)
	src = src[v.MSize.SizeBytes():]
	var versionLen primitive.Uint16
	versionLen.UnmarshalUnsafe(src)
	src = src[versionLen.SizeBytes():]
	v.Version = string(src[:versionLen])
}