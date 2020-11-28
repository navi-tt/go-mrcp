package mpf

/** RTP payload types */
type RtpPayloadType = uint8

const (
	RTP_PT_PCMU RtpPayloadType = 0 /**< PCMU           Audio 8kHz 1 */
	RTP_PT_PCMA RtpPayloadType = 8 /**< PCMA           Audio 8kHz 1 */

	RTP_PT_CN RtpPayloadType = 13 /**< Comfort Noise Audio 8kHz 1 */

	RTP_PT_RESERVED RtpPayloadType = 19 /**< Not used for any particular codec */

	RTP_PT_DYNAMIC     RtpPayloadType = 96  /**< Start of dynamic payload types */
	RTP_PT_DYNAMIC_MAX RtpPayloadType = 127 /**< End of dynamic payload types */

	RTP_PT_UNKNOWN RtpPayloadType = 128 /**< Unknown (invalid) payload type */
)
