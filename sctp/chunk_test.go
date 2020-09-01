package sctp

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestInitChunk(t *testing.T) {
	pkt := &packet{}
	rawPkt := []byte{0x13, 0x88, 0x13, 0x88, 0x00, 0x00, 0x00, 0x00, 0x81, 0x46, 0x9d, 0xfc, 0x01, 0x00, 0x00, 0x56, 0x55,
		0xb9, 0x64, 0xa5, 0x00, 0x02, 0x00, 0x00, 0x04, 0x00, 0x08, 0x00, 0xe8, 0x6d, 0x10, 0x30, 0xc0, 0x00, 0x00, 0x04, 0x80,
		0x08, 0x00, 0x09, 0xc0, 0x0f, 0xc1, 0x80, 0x82, 0x00, 0x00, 0x00, 0x80, 0x02, 0x00, 0x24, 0x9f, 0xeb, 0xbb, 0x5c, 0x50,
		0xc9, 0xbf, 0x75, 0x9c, 0xb1, 0x2c, 0x57, 0x4f, 0xa4, 0x5a, 0x51, 0xba, 0x60, 0x17, 0x78, 0x27, 0x94, 0x5c, 0x31, 0xe6,
		0x5d, 0x5b, 0x09, 0x47, 0xe2, 0x22, 0x06, 0x80, 0x04, 0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x80, 0x03, 0x00, 0x06, 0x80, 0xc1, 0x00, 0x00}
	err := pkt.unmarshal(rawPkt)
	if err != nil {
		t.Error(errors.Wrap(err, "Unmarshal failed, has chunk"))
	}

	i, ok := pkt.chunks[0].(*chunkInit)
	if !ok {
		t.Error("Failed to cast Chunk -> Init")
	}

	if err != nil {
		t.Error(errors.Wrap(err, "Unmarshal init Chunk failed"))
	} else if i.initiateTag != 1438213285 {
		t.Error(errors.Errorf("Unmarshal passed for SCTP packet, but got incorrect initiate tag exp: %d act: %d", 1438213285, i.initiateTag))
	} else if i.advertisedReceiverWindowCredit != 131072 {
		t.Error(errors.Errorf("Unmarshal passed for SCTP packet, but got incorrect advertisedReceiverWindowCredit exp: %d act: %d", 131072, i.advertisedReceiverWindowCredit))
	} else if i.numOutboundStreams != 1024 {
		t.Error(errors.Errorf("Unmarshal passed for SCTP packet, but got incorrect numOutboundStreams tag exp: %d act: %d", 1024, i.numOutboundStreams))
	} else if i.numInboundStreams != 2048 {
		t.Error(errors.Errorf("Unmarshal passed for SCTP packet, but got incorrect numInboundStreams exp: %d act: %d", 2048, i.numInboundStreams))
	} else if i.initialTSN != uint32(3899461680) {
		t.Error(errors.Errorf("Unmarshal passed for SCTP packet, but got incorrect initialTSN exp: %d act: %d", uint32(3899461680), i.initialTSN))
	}
}

func TestInitAck(t *testing.T) {
	pkt := &packet{}
	rawPkt := []byte{0x13, 0x88, 0x13, 0x88, 0xce, 0x15, 0x79, 0xa2, 0x96, 0x19, 0xe8, 0xb2, 0x02, 0x00, 0x00, 0x1c, 0xeb, 0x81, 0x4e, 0x01, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x08, 0x00, 0x50, 0xdf, 0x90, 0xd9, 0x00, 0x07, 0x00, 0x08, 0x94, 0x06, 0x2f, 0x93}
	err := pkt.unmarshal(rawPkt)
	if err != nil {
		t.Error(errors.Wrap(err, "Unmarshal failed, has chunk"))
	}

	_, ok := pkt.chunks[0].(*chunkInitAck)
	if !ok {
		t.Error("Failed to cast Chunk -> Init")
	} else if err != nil {
		t.Error(errors.Wrap(err, "Unmarshal init Chunk failed"))
	}
}

func TestChromeChunk1Init(t *testing.T) {
	pkt := &packet{}
	rawPkt := []byte{0x13, 0x88, 0x13, 0x88, 0x00, 0x00, 0x00, 0x00, 0xbc, 0xb3, 0x45, 0xa2, 0x01, 0x00, 0x00, 0x56, 0xce, 0x15, 0x79, 0xa2, 0x00, 0x02, 0x00, 0x00, 0x04, 0x00, 0x08, 0x00, 0x94, 0x57, 0x95, 0xc0, 0xc0, 0x00, 0x00, 0x04, 0x80, 0x08, 0x00, 0x09, 0xc0, 0x0f, 0xc1, 0x80, 0x82, 0x00, 0x00, 0x00, 0x80, 0x02, 0x00, 0x24, 0xff, 0x5c, 0x49, 0x19, 0x4a, 0x94, 0xe8, 0x2a, 0xec, 0x58, 0x55, 0x62, 0x29, 0x1f, 0x8e, 0x23, 0xcd, 0x7c, 0xe8, 0x46, 0xba, 0x58, 0x1b, 0x3d, 0xab, 0xd7, 0x7e, 0x50, 0xf2, 0x41, 0xb1, 0x2e, 0x80, 0x04, 0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x80, 0x03, 0x00, 0x06, 0x80, 0xc1, 0x00, 0x00}
	err := pkt.unmarshal(rawPkt)
	if err != nil {
		t.Error(errors.Wrap(err, "Unmarshal failed, has chunk"))
	}

	rawPkt2, err := pkt.marshal()
	if err != nil {
		t.Error(errors.Wrap(err, "Remarshal failed"))
	}

	assert.Equal(t, rawPkt, rawPkt2)
}

func TestChromeChunk2InitAck(t *testing.T) {
	pkt := &packet{}
	rawPkt := []byte{0x13, 0x88, 0x13, 0x88, 0xce, 0x15, 0x79, 0xa2, 0xb5, 0xdb, 0x2d, 0x93, 0x02, 0x00, 0x01, 0x90, 0x9b, 0xd5, 0xb3, 0x6f, 0x00, 0x02, 0x00, 0x00, 0x04, 0x00, 0x08, 0x00, 0xef, 0xb4, 0x72, 0x87, 0xc0, 0x00, 0x00, 0x04, 0x80, 0x08, 0x00, 0x09, 0xc0, 0x0f, 0xc1, 0x80, 0x82, 0x00, 0x00, 0x00, 0x80, 0x02, 0x00, 0x24, 0x2e, 0xf9, 0x9c, 0x10, 0x63, 0x72, 0xed, 0x0d, 0x33, 0xc2, 0xdc, 0x7f, 0x9f, 0xd7, 0xef, 0x1b, 0xc9, 0xc4, 0xa7, 0x41, 0x9a, 0x07, 0x68, 0x6b, 0x66, 0xfb, 0x6a, 0x4e, 0x32, 0x5d, 0xe4, 0x25, 0x80, 0x04, 0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x80, 0x03, 0x00, 0x06, 0x80, 0xc1, 0x00, 0x00, 0x00, 0x07, 0x01, 0x38, 0x4b, 0x41, 0x4d, 0x45, 0x2d, 0x42, 0x53, 0x44, 0x20, 0x31, 0x2e, 0x31, 0x00, 0x00, 0x00, 0x00, 0x9c, 0x1e, 0x49, 0x5b, 0x00, 0x00, 0x00, 0x00, 0xd2, 0x42, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x60, 0xea, 0x00, 0x00, 0xc4, 0x13, 0x3d, 0xe9, 0x86, 0xb1, 0x85, 0x75, 0xa2, 0x79, 0x15, 0xce, 0x9b, 0xd5, 0xb3, 0x6f, 0x20, 0xe0, 0x9f, 0x89, 0xe0, 0x27, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x20, 0xe0, 0x9f, 0x89, 0xe0, 0x27, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x13, 0x88, 0x13, 0x88, 0x00, 0x00, 0x01, 0x00, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x56, 0xce, 0x15, 0x79, 0xa2, 0x00, 0x02, 0x00, 0x00, 0x04, 0x00, 0x08, 0x00, 0x94, 0x57, 0x95, 0xc0, 0xc0, 0x00, 0x00, 0x04, 0x80, 0x08, 0x00, 0x09, 0xc0, 0x0f, 0xc1, 0x80, 0x82, 0x00, 0x00, 0x00, 0x80, 0x02, 0x00, 0x24, 0xff, 0x5c, 0x49, 0x19, 0x4a, 0x94, 0xe8, 0x2a, 0xec, 0x58, 0x55, 0x62, 0x29, 0x1f, 0x8e, 0x23, 0xcd, 0x7c, 0xe8, 0x46, 0xba, 0x58, 0x1b, 0x3d, 0xab, 0xd7, 0x7e, 0x50, 0xf2, 0x41, 0xb1, 0x2e, 0x80, 0x04, 0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x80, 0x03, 0x00, 0x06, 0x80, 0xc1, 0x00, 0x00, 0x02, 0x00, 0x01, 0x90, 0x9b, 0xd5, 0xb3, 0x6f, 0x00, 0x02, 0x00, 0x00, 0x04, 0x00, 0x08, 0x00, 0xef, 0xb4, 0x72, 0x87, 0xc0, 0x00, 0x00, 0x04, 0x80, 0x08, 0x00, 0x09, 0xc0, 0x0f, 0xc1, 0x80, 0x82, 0x00, 0x00, 0x00, 0x80, 0x02, 0x00, 0x24, 0x2e, 0xf9, 0x9c, 0x10, 0x63, 0x72, 0xed, 0x0d, 0x33, 0xc2, 0xdc, 0x7f, 0x9f, 0xd7, 0xef, 0x1b, 0xc9, 0xc4, 0xa7, 0x41, 0x9a, 0x07, 0x68, 0x6b, 0x66, 0xfb, 0x6a, 0x4e, 0x32, 0x5d, 0xe4, 0x25, 0x80, 0x04, 0x00, 0x06, 0x00, 0x01, 0x00, 0x00, 0x80, 0x03, 0x00, 0x06, 0x80, 0xc1, 0x00, 0x00, 0xca, 0x0c, 0x21, 0x11, 0xce, 0xf4, 0xfc, 0xb3, 0x66, 0x99, 0x4f, 0xdb, 0x4f, 0x95, 0x6b, 0x6f, 0x3b, 0xb1, 0xdb, 0x5a}
	err := pkt.unmarshal(rawPkt)
	if err != nil {
		t.Error(errors.Wrap(err, "Unmarshal failed, has chunk"))
	}

	rawPkt2, err := pkt.marshal()
	if err != nil {
		t.Error(errors.Wrap(err, "Remarshal failed"))
	}

	assert.Equal(t, rawPkt, rawPkt2)
}

func TestInitMarshalUnmarshal(t *testing.T) {
	p := &packet{}
	p.destinationPort = 1
	p.sourcePort = 1
	p.verificationTag = 123

	initAck := &chunkInitAck{}

	initAck.initialTSN = 123
	initAck.numOutboundStreams = 1
	initAck.numInboundStreams = 1
	initAck.initiateTag = 123
	initAck.advertisedReceiverWindowCredit = 1024
	initAck.params = []param{newRandomStateCookie()}

	p.chunks = []chunk{initAck}
	rawPkt, err := p.marshal()
	if err != nil {
		t.Error(errors.Wrap(err, "Failed to marshal packet"))
	}

	pkt := &packet{}
	err = pkt.unmarshal(rawPkt)
	if err != nil {
		t.Error(errors.Wrap(err, "Unmarshal failed, has chunk"))
	}

	i, ok := pkt.chunks[0].(*chunkInitAck)
	if !ok {
		t.Error("Failed to cast Chunk -> InitAck")
	}

	if err != nil {
		t.Error(errors.Wrap(err, "Unmarshal init ack Chunk failed"))
	} else if i.initiateTag != 123 {
		t.Error(errors.Errorf("Unmarshal passed for SCTP packet, but got incorrect initiate tag exp: %d act: %d", 123, i.initiateTag))
	} else if i.advertisedReceiverWindowCredit != 1024 {
		t.Error(errors.Errorf("Unmarshal passed for SCTP packet, but got incorrect advertisedReceiverWindowCredit exp: %d act: %d", 1024, i.advertisedReceiverWindowCredit))
	} else if i.numOutboundStreams != 1 {
		t.Error(errors.Errorf("Unmarshal passed for SCTP packet, but got incorrect numOutboundStreams tag exp: %d act: %d", 1, i.numOutboundStreams))
	} else if i.numInboundStreams != 1 {
		t.Error(errors.Errorf("Unmarshal passed for SCTP packet, but got incorrect numInboundStreams exp: %d act: %d", 1, i.numInboundStreams))
	} else if i.initialTSN != 123 {
		t.Error(errors.Errorf("Unmarshal passed for SCTP packet, but got incorrect initialTSN exp: %d act: %d", 123, i.initialTSN))
	}
}

func TestPayloadDataMarshalUnmarshal(t *testing.T) {
	pkt := &packet{}
	rawPkt := []byte{0x13, 0x88, 0x13, 0x88, 0xfc, 0xd6, 0x3f, 0xc6, 0xbe, 0xfa, 0xdc, 0x52, 0x0a, 0x00, 0x00, 0x24, 0x9b, 0x28, 0x7e, 0x48, 0xa3, 0x7b, 0xc1, 0x83, 0xc4, 0x4b, 0x41, 0x04, 0xa4, 0xf7, 0xed, 0x4c, 0x93, 0x62, 0xc3, 0x49, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x1f, 0xa8, 0x79, 0xa1, 0xc7, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x32, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x66, 0x6f, 0x6f, 0x00}
	err := pkt.unmarshal(rawPkt)
	if err != nil {
		t.Error(errors.Wrap(err, "Unmarshal failed, has chunk"))
	}

	_, ok := pkt.chunks[1].(*chunkPayloadData)
	if !ok {
		t.Error("Failed to cast Chunk -> PayloadData")
	}
}

func TestSelectAckChunk(t *testing.T) {
	pkt := &packet{}
	rawPkt := []byte{0x13, 0x88, 0x13, 0x88, 0xc2, 0x98, 0x98, 0x0f, 0x42, 0x31, 0xea, 0x78, 0x03, 0x00, 0x00, 0x14, 0x87, 0x73, 0xbd, 0xa4, 0x00, 0x01, 0xfe, 0x74, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x02}
	err := pkt.unmarshal(rawPkt)
	if err != nil {
		t.Error(errors.Wrap(err, "Unmarshal failed, has chunk"))
	}

	_, ok := pkt.chunks[0].(*chunkSelectiveAck)
	if !ok {
		t.Error("Failed to cast Chunk -> SelectiveAck")
	}
}

func TestReconfigChunk(t *testing.T) {
	pkt := &packet{}
	rawPkt := []byte{0x13, 0x88, 0x13, 0x88, 0xb6, 0xa5, 0x12, 0xe5, 0x75, 0x3b, 0x12, 0xd3, 0x82, 0x0, 0x0, 0x16, 0x0, 0xd, 0x0, 0x12, 0x4e, 0x1c, 0xb9, 0xe6, 0x3a, 0x74, 0x8d, 0xff, 0x4e, 0x1c, 0xb9, 0xe6, 0x0, 0x1, 0x0, 0x0}
	err := pkt.unmarshal(rawPkt)
	if err != nil {
		t.Error(errors.Wrap(err, "Unmarshal failed, has chunk"))
	}

	c, ok := pkt.chunks[0].(*chunkReconfig)
	if !ok {
		t.Error("Failed to cast Chunk -> Reconfig")
	}

	if c.paramA.(*paramOutgoingResetRequest).streamIdentifiers[0] != uint16(1) {
		t.Error(errors.Errorf("unexpected stream identifier: %d", c.paramA.(*paramOutgoingResetRequest).streamIdentifiers[0]))
	}
}

func TestForwardTSNChunk(t *testing.T) {
	pkt := &packet{}
	rawPkt := append([]byte{0x13, 0x88, 0x13, 0x88, 0xb6, 0xa5, 0x12, 0xe5, 0x1f, 0x9d, 0xa0, 0xfb}, testChunkForwardTSN...)
	err := pkt.unmarshal(rawPkt)
	if err != nil {
		t.Error(errors.Wrap(err, "Unmarshal failed, has chunk"))
	}

	c, ok := pkt.chunks[0].(*chunkForwardTSN)
	if !ok {
		t.Error("Failed to cast Chunk -> Forward TSN")
	}

	if c.newCumulativeTSN != uint32(3) {
		t.Errorf("unexpected New Cumulative TSN: %d", c.newCumulativeTSN)
	}
}
