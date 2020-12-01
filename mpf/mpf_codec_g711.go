package mpf

import (
	"bytes"
	"github.com/navi-tt/go-mrcp/mpf/codecs/g711"
)

const (
	G711U_CODEC_NAME = "PCMU"
	G711A_CODEC_NAME = "PCMA"
)

func G711Open(codec *Codec) error {
	return nil
}

func G711Close(codec *Codec) error {
	return nil
}

func G711UEncode(codec *Codec, frameIn, frameOut *CodecFrame) error {
	data := g711.EncodeUlaw(frameIn.Buffer.Bytes())
	n, err := frameOut.Buffer.Write(data)
	frameOut.Size = int64(n)
	return err
}

func G711UDecode(codec *Codec, frameIn, frameOut *CodecFrame) error {
	data := g711.DecodeUlaw(frameIn.Buffer.Bytes())
	n, err := frameOut.Buffer.Write(data)
	frameOut.Size = int64(n)
	return err
}

func G711UInit(codec *Codec, frameOut *CodecFrame) error {
	buf := bytes.NewBuffer(make([]byte, frameOut.Size))
	data := g711.EncodeUlaw(buf.Bytes())
	n, err := frameOut.Buffer.Write(data)
	frameOut.Size = int64(n)
	return err
}

func G711AEncode(codec *Codec, frameIn, frameOut *CodecFrame) error {
	data := g711.EncodeAlaw(frameIn.Buffer.Bytes())
	n, err := frameOut.Buffer.Write(data)
	frameOut.Size = int64(n)
	return err
}

func G711ADecode(codec *Codec, frameIn, frameOut *CodecFrame) error {
	data := g711.DecodeAlaw(frameIn.Buffer.Bytes())
	n, err := frameOut.Buffer.Write(data)
	frameOut.Size = int64(n)
	return err
}

func G711AInit(codec *Codec, frameOut *CodecFrame) error {
	buf := bytes.NewBuffer(make([]byte, frameOut.Size))
	data := g711.EncodeAlaw(buf.Bytes())
	n, err := frameOut.Buffer.Write(data)
	frameOut.Size = int64(n)
	return err
}

var g711UVTable = CodecVTable{
	Open:       G711Open,
	Close:      G711Close,
	Encode:     G711UEncode,
	Decode:     G711UDecode,
	Dissect:    nil,
	Initialize: G711UInit,
}

var g711AVTable = CodecVTable{
	Open:       G711Open,
	Close:      G711Close,
	Encode:     G711AEncode,
	Decode:     G711ADecode,
	Dissect:    nil,
	Initialize: G711AInit,
}

var g711UDescriptor = CodecDescriptor{
	PayloadType:  RTP_PT_PCMU,
	Name:         G711U_CODEC_NAME,
	SamplingRate: 8000,
	ChannelCount: 1,
	Format:       "",
	Enabled:      true,
}

var g711ADescriptor = CodecDescriptor{
	PayloadType:  RTP_PT_PCMA,
	Name:         G711A_CODEC_NAME,
	SamplingRate: 8000,
	ChannelCount: 1,
	Format:       "",
	Enabled:      true,
}

var g711UAttribs = CodecAttribs{
	Name:          G711U_CODEC_NAME,
	BitsPerSample: 8,
	SampleRates:   MPF_SAMPLE_RATE_8000 | MPF_SAMPLE_RATE_16000 | MPF_SAMPLE_RATE_32000 | MPF_SAMPLE_RATE_48000,
}

var g711AAttribs = CodecAttribs{
	Name:          G711A_CODEC_NAME,
	BitsPerSample: 8,
	SampleRates:   MPF_SAMPLE_RATE_8000 | MPF_SAMPLE_RATE_16000 | MPF_SAMPLE_RATE_32000 | MPF_SAMPLE_RATE_48000,
}

func CodecG711UCreate() *Codec {
	return CodecCreate(&g711UVTable, &g711UAttribs, &g711UDescriptor)
}

func CodecG711ACreate() *Codec {
	return CodecCreate(&g711AVTable, &g711AAttribs, &g711ADescriptor)
}
