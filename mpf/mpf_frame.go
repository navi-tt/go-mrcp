package mpf

/** Media frame types */
type FrameType = byte

const (
	MEDIA_FRAME_TYPE_NONE  FrameType = 0x0 /**< none */
	MEDIA_FRAME_TYPE_AUDIO FrameType = 0x1 /**< audio frame */
	MEDIA_FRAME_TYPE_VIDEO FrameType = 0x2 /**< video frame */
	MEDIA_FRAME_TYPE_EVENT FrameType = 0x4 /**< named event frame (RFC4733/RFC2833) */
)

/** Media frame marker */
type FrameMarker = int

const (
	MPF_MARKER_NONE           FrameMarker = iota /**< none */
	MPF_MARKER_START_OF_EVENT                    /**< start of event */
	MPF_MARKER_END_OF_EVENT                      /**< end of event */
	MPF_MARKER_NEW_SEGMENT                       /**< start of new segment (long-lasting events) */
)

/** Media frame */
type Frame struct {
	/** frame type (audio/video/named-event) mpf_frame_type_e */
	Type int
	/** frame marker (start-of-eventend-of-event) mpf_frame_marker_e */
	Marker int
	/** codec frame */
	CodecFrame CodecFrame
	/** named-event frame */
	//mpf_named_event_frame_t event_frame;
}
