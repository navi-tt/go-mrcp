package mpf

/** Table of audio stream virtual methods */
type MPFAudioStreamVTable struct {

	/** Virtual destroy method */
	Destroy func(stream *MPFAudioStream) error

	/** Virtual open receiver method */
	OpenRX func(stream *MPFAudioStream, codec *MPFCodec) error
	/** Virtual close receiver method */
	CloseRX func(stream *MPFAudioStream) error
	/** Virtual read frame method */
	ReadFrame func(stream *MPFAudioStream, frame *MPFFrame) error

	/** Virtual open transmitter method */
	OpenTX func(stream *MPFAudioStream, codec *MPFCodec) error
	/** Virtual close transmitter method */
	CloseTX func(stream *MPFAudioStream) error
	/** Virtual write frame method */
	WriteFrame func(stream *MPFAudioStream, frame *MPFFrame) error

	/** Virtual trace method */
	//Trace func(mpf_audio_stream_t *stream, mpf_stream_direction_e direction, apt_text_stream_t *output);
}

/** Audio stream */
type MPFAudioStream struct {

	/** External object */
	Obj interface{}
	/** Table of virtual methods */
	VTable *MPFAudioStreamVTable
	/** Back pointer */
	termination *MPFTermination

	/** Stream capabilities */
	//const mpf_stream_capabilities_t *capabilities;
	/** Stream direction send/receive (bitmask of mpf_stream_direction_e) */
	//mpf_stream_direction_e           direction;

	/** Rx codec descriptor */
	RXDescriptor *MPFCodecDescriptor
	/** Rx event descriptor */
	RXEventDescriptor *MPFCodecDescriptor
	/** Tx codec descriptor */
	TXDescriptor *MPFCodecDescriptor
	/** Tx event descriptor */
	TXEventDescriptor *MPFCodecDescriptor
}
