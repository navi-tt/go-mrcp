package mpf

/** RTP attributes */
type RtpAttrib = int

const (
	RTP_ATTRIB_RTPMAP RtpAttrib = iota
	RTP_ATTRIB_SENDONLY
	RTP_ATTRIB_RECVONLY
	RTP_ATTRIB_SENDRECV
	RTP_ATTRIB_MID
	RTP_ATTRIB_PTIME

	RTP_ATTRIB_COUNT
	RTP_ATTRIB_UNKNOWN = RTP_ATTRIB_COUNT
)

/** Get audio media attribute name by attribute identifier */
func RtpAttribStrGet(attribId RtpAttrib) string {
	return ""
}

/** Find audio media attribute identifier by attribute name */
func RtpAttribIdFind(attrib string) RtpAttrib {
	return 0
}

/** Get string by RTP direction (send/receive) */
func RtpDirectionStrGet(direction StreamDirection) string {
	return ""
}
