package mpf

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
	//Trace func(mpf_audio_stream_t *stream, mpf_stream_direction_e direction, apt_text_stream_t *output);
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
	//const mpf_stream_capabilities_t *capabilities;
	/** Stream direction send/receive (bitmask of mpf_stream_direction_e) */
	//mpf_stream_direction_e           direction;

	/** Rx codec descriptor */
	RXDescriptor *CodecDescriptor
	/** Rx event descriptor */
	RXEventDescriptor *CodecDescriptor
	/** Tx codec descriptor */
	TXDescriptor *CodecDescriptor
	/** Tx event descriptor */
	TXEventDescriptor *CodecDescriptor
}
