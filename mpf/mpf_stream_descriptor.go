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
	direction StreamDirection
	/** Codec capabilities (supported codecs and named events) */
	codecs CodecCapabilities
}

/** Create stream capabilities */
func StreamCapabilitiesCreate(directions StreamDirection) *StreamCapabilities {
	return nil
}

/** Create source stream capabilities */
func SourceStreamCapabilitiesCreate() *StreamCapabilities {
	return StreamCapabilitiesCreate(STREAM_DIRECTION_RECEIVE)
}

/** Create sink stream capabilities */
func SinkStreamCapabilitiesCreate() *StreamCapabilities {
	return StreamCapabilitiesCreate(STREAM_DIRECTION_SEND)
}

/** Clone stream capabilities */
func StreamCapabilitiesClone(srcCapabilities *StreamCapabilities) *StreamCapabilities {
	return nil
}

/** Merge stream capabilities */
func StreamCapabilitiesMerge(capabilities, srcCapabilities *StreamCapabilities) error {
	return nil
}

/** Get reverse direction */
func StreamReverseDirectionGet(direction StreamDirection) StreamDirection {
	revDirection := direction
	if revDirection == STREAM_DIRECTION_SEND {
		revDirection = STREAM_DIRECTION_RECEIVE
	} else if revDirection == STREAM_DIRECTION_RECEIVE {
		revDirection = STREAM_DIRECTION_SEND
	}
	return revDirection
}
