package mpf

import (
	"github.com/navi-tt/go-mrcp/utils/binaryx"
)

/* linear 16-bit PCM (RFC3551) */
const L16_CODEC_NAME = "L16"

func L16Open(codec *Codec) error {
	return nil
}

func L16Close(codec *Codec) error {
	return nil
}

func L16Encode(codec *Codec, frameIn, frameOut *CodecFrame) error {
	ins, err := binaryx.ByteSliceToInt16Slice(frameIn.Buffer.Bytes())
	if err != nil {
		return err
	}

	var outs = make([]int16, 0, len(ins))
	for i := 0; i < len(ins); i++ {
		out := binaryx.HToNS(uint16(ins[i]))
		outs = append(outs, int16(out))
	}

	data := binaryx.Int16SliceToByteSlice(outs)
	frameOut.Buffer.Write(data)
	return err
}

func L16Decode(codec *Codec, frameIn, frameOut *CodecFrame) error {
	ins, err := binaryx.ByteSliceToInt16Slice(frameIn.Buffer.Bytes())
	if err != nil {
		return err
	}

	var outs = make([]int16, 0, len(ins))
	for i := 0; i < len(ins); i++ {
		out := binaryx.NToHS(uint16(ins[i]))
		outs = append(outs, int16(out))
	}

	data := binaryx.Int16SliceToByteSlice(outs)
	frameOut.Buffer.Write(data)
	return err
}

var l16VTable = CodecVTable{
	Open:       L16Open,
	Close:      L16Close,
	Encode:     L16Encode,
	Decode:     L16Decode,
	Dissect:    nil,
	Initialize: nil,
}

var l16Attribs = CodecAttribs{
	Name:          L16_CODEC_NAME,
	BitsPerSample: 16,
	SampleRates:   MPF_SAMPLE_RATE_8000 | MPF_SAMPLE_RATE_16000 | MPF_SAMPLE_RATE_32000 | MPF_SAMPLE_RATE_48000, /* supported sampling rates */
}

func CodecL16Create() *Codec {
	return CodecCreate(&l16VTable, &l16Attribs, nil)
}
