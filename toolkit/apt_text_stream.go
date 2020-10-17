package toolkit

/** Text stream is used for message parsing and generation */
type AptTextStream struct {
	/** Text stream */
	text string
	/** Current position in the stream */
	pos *byte
	/** End of stream pointer */
	end *byte
	/** Is end of stream reached */
	isEos bool
}
