package mpf

import (
	"bytes"
)

/** Codec */
type Codec struct {

	/** Codec manipulators (encode, decode, dissect) */
	VTable *CodecVTable
	/** Codec attributes (capabilities) */
	Attribs *CodecAttribs
	/** Optional static codec descriptor (pt < 96) */
	StaticDescriptor *CodecDescriptor
}

/** Table of codec virtual methods */
type CodecVTable struct {

	/** Virtual open method */
	Open func(codec *Codec) error
	/** Virtual close method */
	Close func(codec *Codec) error

	/** Virtual encode method */
	Encode func(codec *Codec, frameIn, frameOut *CodecFrame) error
	/** Virtual decode method */
	Decode func(codec *Codec, frameIn, frameOut *CodecFrame) error

	/** Virtual dissect method */
	Dissect func(codec *Codec, buffer *bytes.Buffer, frame *CodecFrame) error

	/** Virtual initialize method */
	Initialize func(codec *Codec, frameOut *CodecFrame) error
}

/**
 * Create codec.
 * @param vtable the table of virtual mthods
 * @param attribs the codec attributes
 * @param descriptor the codec descriptor
 * @param pool the pool to allocate memory from
 */
func CodecCreate(vtable *CodecVTable, attribs *CodecAttribs, descriptor *CodecDescriptor) *Codec {
	return &Codec{
		VTable:           vtable,
		Attribs:          attribs,
		StaticDescriptor: descriptor,
	}
}

/**
 * Clone codec.
 * @param src_codec the source (original) codec to clone
 * @param pool the pool to allocate memory from
 */
func (c *Codec) CodecClone() *Codec {
	codec := &Codec{
		VTable:           c.VTable,
		Attribs:          c.Attribs,
		StaticDescriptor: c.StaticDescriptor,
	}
	return codec
}

/** Open codec */
func (c *Codec) CodecOpen() error {
	if c.VTable != nil && c.VTable.Open != nil {
		return c.VTable.Open(c)
	}
	return nil
}

/** Close codec */
func (c *Codec) CodecClose() error {
	if c.VTable != nil && c.VTable.Close != nil {
		return c.VTable.Close(c)
	}
	return nil
}

/** Encode codec frame */
func (c *Codec) CodecEncode(frameIn, frameOut *CodecFrame) error {
	if c.VTable != nil && c.VTable.Encode != nil {
		return c.VTable.Encode(c, frameIn, frameOut)
	}
	return nil
}

/** Decode codec frame */
func (c *Codec) CodecDecode(frameIn, frameOut *CodecFrame) error {
	if c.VTable != nil && c.VTable.Decode != nil {
		return c.VTable.Decode(c, frameIn, frameOut)
	}
	return nil
}

/** Dissect codec frame (navigate through codec frames in a buffer, which may contain multiple frames) */
func (c *Codec) CodecDissect(buffer *bytes.Buffer, frame *CodecFrame) error {
	if c.VTable != nil && c.VTable.Dissect != nil {
		return c.VTable.Dissect(c, buffer, frame)
	} else {
		/* default dissector */
		if frame.Buffer != nil && frame.Buffer.Len() > 0 && buffer.Len() >= frame.Buffer.Len() {
			frame.Buffer.Reset()
			_, err := frame.Buffer.Write(buffer.Bytes())
			return err
		}
	}
	return nil
}

/** Initialize (fill) codec frame with silence */
func (c *Codec) CodecInitialize(frameOut *CodecFrame) error {
	if c.VTable != nil && c.VTable.Initialize != nil {
		return c.VTable.Initialize(c, frameOut)
	} else {
		frameOut.Buffer.Reset()
	}
	return nil
}
