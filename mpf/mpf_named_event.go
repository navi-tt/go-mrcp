package mpf

/** Named event (RFC4733/RFC2833, out-of-band DTMF) */
type NamedEventFrame struct {

	/** event (DTMF, tone) identifier */
	EventId uint32 // 8
	//#if (APR_IS_BIGENDIAN == 1)
	///** end of event */
	//apr_uint32_t edge:     1;
	///** reserved */
	//apr_uint32_t reserved: 1;
	///** tone volume */
	//apr_uint32_t volume:   6;
	//#else

	// 小端
	/** tone volume */
	Volume uint32 //    6;
	/** reserved */
	Reserved uint32 // 1;
	/** end of event */
	Edge uint32 //    1;

	/** event duration */
	Duration uint32 // 16;
}

/** Create named event descriptor */
func EventDescriptorCreate(samplingRate uint16) *CodecDescriptor {
	return nil
}

/** Check whether the specified descriptor is named event one */
func EventDescriptorCheck(descriptor *CodecDescriptor) bool {
	return false
}

/** Convert DTMF character to event identifier */
func DtmfCharToEventId(dtmfChar byte) uint32 {
	return 0
}

/** Convert event identifier to DTMF character */
func EventIdToDtmfChar(eventId uint32) byte {
	return 0
}
