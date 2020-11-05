package mpf

import "github.com/navi-tt/go-mrcp/toolkit"

/** Table of audio stream virtual methods */
type AudioStreamVTable struct {

	/** Virtual destroy method */
	Destroy func(stream *AudioStream) error

	/** Virtual open receiver method */
	OpenRX func(stream *AudioStream, codec *Codec) error
	/** Virtual close receiver method */
	CloseRX func(stream *AudioStream) error
	/** Virtual read frame method */
	ReadFrame func(stream *AudioStream, frame *Frame) error

	/** Virtual open transmitter method */
	OpenTX func(stream *AudioStream, codec *Codec) error
	/** Virtual close transmitter method */
	CloseTX func(stream *AudioStream) error
	/** Virtual write frame method */
	WriteFrame func(stream *AudioStream, frame *Frame) error

	/** Virtual trace method */
	Trace func(stream *AudioStream, direction StreamDirection, output *toolkit.AptTextStream)
}

/** Audio stream */
type AudioStream struct {

	/** External object */
	Obj interface{}
	/** Table of virtual methods */
	VTable *AudioStreamVTable
	/** Back pointer */
	termination *Termination

	/** Stream capabilities */
	Capabilities *StreamCapabilities

	/** Stream direction send/receive (bitmask of mpf_stream_direction_e) */
	direction StreamDirection

	/** Rx codec descriptor */
	RXDescriptor *CodecDescriptor
	/** Rx event descriptor */
	RXEventDescriptor *CodecDescriptor
	/** Tx codec descriptor */
	TXDescriptor *CodecDescriptor
	/** Tx event descriptor */
	TXEventDescriptor *CodecDescriptor
}

/** Video stream */
type VideoStream struct {

	/** Back pointer */
	termination *Termination
	/** Stream direction send/receive (bitmask of mpf_stream_direction_e) */
	direction StreamDirection
}

/** Create audio stream */
func AudioStreamCreate(obj interface{}, vtable *AudioStreamVTable, capabilities *StreamCapabilities) *AudioStream {
	return nil
}

/** Validate audio stream receiver */
func (stream *AudioStream) AudioStreamRXValidate(descriptor, eventDescriptor *CodecDescriptor) error {
	return nil
}

/** Validate audio stream transmitter */
func (stream *AudioStream) AudioStreamTXValidate(descriptor, eventDescriptor *CodecDescriptor) error {
	return nil
}

/** Destroy audio stream */
func AudioStreamDestroy(stream *AudioStream) error {
	if stream.VTable != nil && stream.VTable.Destroy != nil {
		return stream.VTable.Destroy(stream)
	}
	return nil
}

/** Open audio stream receiver */
func (stream *AudioStream) AudioStreamRXOpen(codec *Codec) error {
	if stream.VTable != nil && stream.VTable.OpenRX != nil {
		return stream.VTable.OpenRX(stream, codec)
	}
	return nil
}

/** Close audio stream receiver */
func (stream *AudioStream) AudioStreamRXClose() error {
	if stream.VTable != nil && stream.VTable.CloseRX != nil {
		return stream.VTable.CloseRX(stream)
	}
	return nil
}

/** Read frame */
func (stream *AudioStream) AudioStreamFrameRead(frame *Frame) error {
	if stream.VTable != nil && stream.VTable.ReadFrame != nil {
		return stream.VTable.ReadFrame(stream, frame)
	}
	return nil
}

/** Open audio stream transmitter */
func (stream *AudioStream) AudioStreamTXOpen(codec *Codec) error {
	if stream.VTable != nil && stream.VTable.OpenTX != nil {
		return stream.VTable.OpenTX(stream, codec)
	}
	return nil
}

/** Close audio stream transmitter */
func (stream *AudioStream) AudioStreamTXClose() error {
	if stream.VTable != nil && stream.VTable.CloseTX != nil {
		return stream.VTable.CloseTX(stream)
	}
	return nil
}

/** Write frame */
func (stream *AudioStream) AudioStreamFrameWrite(frame *Frame) error {
	if stream.VTable != nil && stream.VTable.WriteFrame != nil {
		return stream.VTable.WriteFrame(stream, frame)
	}
	return nil
}

/** Trace media path */
func (stream *AudioStream) AudioStreamTrace(direction StreamDirection, output *toolkit.AptTextStream) {
}
