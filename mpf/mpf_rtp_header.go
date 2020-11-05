package mpf

const RTP_VERSION = 2

type RtpHeader struct {

	//#if (APR_IS_BIGENDIAN == 1)
	///** protocol version */
	//apr_uint32_t version:   2;
	///** padding flag */
	//apr_uint32_t padding:   1;
	///** header extension flag */
	//apr_uint32_t extension: 1;
	///** CSRC count */
	//apr_uint32_t count:     4;
	///** marker bit */
	//apr_uint32_t marker:    1;
	///** payload type */
	//apr_uint32_t type:      7;

	/** CSRC count */
	count uint32 //   4;
	/** header extension flag */
	Extension uint32 // 1;
	/** padding flag */
	Padding uint32 //   1;
	/** protocol version */
	Version uint32 //   2;
	/** payload type */
	Type uint32 //     7;
	/** marker bit */
	Marker uint32 //    1;

	/** sequence number */
	sequence uint32 //   16;
	/** timestamp */
	timestamp uint32
	/** synchronization source */
	ssrc uint32
}

/** RTP extension header */
type RtpExtensionHeader struct {
	/** profile */
	profile uint16
	/** length */
	length uint16
}
