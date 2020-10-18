package mpf

/** Media frame */
type MPFFrame struct {
	/** frame type (audio/video/named-event) mpf_frame_type_e */
	Type int
	/** frame marker (start-of-event,end-of-event) mpf_frame_marker_e */
	Marker int
	/** codec frame */
	CodecFrame MPFCodecFrame
	/** named-event frame */
	//mpf_named_event_frame_t event_frame;
}
