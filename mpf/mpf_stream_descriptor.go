package mpf

/** Stream directions (none, send, receive, duplex) */
type StreamDirection = int

const (
	STREAM_DIRECTION_NONE    StreamDirection = 0x0 /**< none */
	STREAM_DIRECTION_SEND    StreamDirection = 0x1 /**< send (sink) */
	STREAM_DIRECTION_RECEIVE StreamDirection = 0x2 /**< receive (source) */

	STREAM_DIRECTION_DUPLEX StreamDirection = STREAM_DIRECTION_SEND | STREAM_DIRECTION_RECEIVE /**< duplex */
)

/** Stream capabilities */
type StreamCapabilities struct {
	/** Supported directions either send, receive or bidirectional stream (bitmask of mpf_stream_direction_e) */
	//mpf_stream_direction_e   direction;
	/** Codec capabilities (supported codecs and named events) */
	//mpf_codec_capabilities_t codecs;
}
