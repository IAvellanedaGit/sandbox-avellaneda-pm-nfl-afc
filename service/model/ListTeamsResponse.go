// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ListTeamsResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsListTeamsResponse(buf []byte, offset flatbuffers.UOffsetT) *ListTeamsResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ListTeamsResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsListTeamsResponse(buf []byte, offset flatbuffers.UOffsetT) *ListTeamsResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ListTeamsResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ListTeamsResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ListTeamsResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ListTeamsResponse) Teams(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *ListTeamsResponse) TeamsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ListTeamsResponse) Error() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ListTeamsResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func ListTeamsResponseAddTeams(builder *flatbuffers.Builder, teams flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(teams), 0)
}
func ListTeamsResponseStartTeamsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ListTeamsResponseAddError(builder *flatbuffers.Builder, error flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(error), 0)
}
func ListTeamsResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
