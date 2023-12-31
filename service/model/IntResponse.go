// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type IntResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsIntResponse(buf []byte, offset flatbuffers.UOffsetT) *IntResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &IntResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsIntResponse(buf []byte, offset flatbuffers.UOffsetT) *IntResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &IntResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *IntResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *IntResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *IntResponse) Output() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *IntResponse) MutateOutput(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *IntResponse) Error() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func IntResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func IntResponseAddOutput(builder *flatbuffers.Builder, output int32) {
	builder.PrependInt32Slot(0, output, 0)
}
func IntResponseAddError(builder *flatbuffers.Builder, error flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(error), 0)
}
func IntResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
