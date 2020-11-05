package mpf

/**
 * Create RTP stream.
 * @param termination the back pointer to hold
 * @param config the configuration of RTP factory
 * @param settings the settings to use
 * @param pool the pool to allocate memory from
 */
func RtpStreamCreate(termination *Termination, config *RtpConfig, settings *RtpSettings) *AudioStream {
	return nil
}

/**
 * Add/enable RTP stream.
 * @param stream RTP stream to add
 */
func RtpStreamAdd(stream *AudioStream) error {
	return nil
}

/**
 * Subtract/disable RTP stream.
 * @param stream RTP stream to subtract
 */
func RtpStreamRemove(stream *AudioStream) error {
	return nil
}

/**
 * Modify RTP stream.
 * @param stream RTP stream to modify
 * @param descriptor the descriptor to modify stream according
 */
func RtpStreamModify(stream *AudioStream, descriptor *RtpStreamDescriptor) error {
	return nil
}
