package mpf

import "github.com/navi-tt/go-mrcp/utils/binaryx"

/** RTCP payload (packet) types */
type RtcpType = int

const (
	RTCP_SR   RtcpType = 200
	RTCP_RR   RtcpType = 201
	RTCP_SDES RtcpType = 202
	RTCP_BYE  RtcpType = 203
	RTCP_APP  RtcpType = 204
)

/** RTCP SDES types */
type RtcpSdesType = int

const (
	RTCP_SDES_END RtcpSdesType = iota
	RTCP_SDES_CNAME
	RTCP_SDES_NAME
	RTCP_SDES_EMAIL
	RTCP_SDES_PHONE
	RTCP_SDES_LOC
	RTCP_SDES_TOOL
	RTCP_SDES_NOTE
	RTCP_SDES_PRIV
)

/** RTCP header */
type RtcpHeader struct {

	/* 大端 */
	///** protocol version */
	//unsigned int version: 2;
	///** padding flag */
	//unsigned int padding: 1;
	///** varies by packet type */
	//unsigned int count:   5;
	///** packet type */
	//unsigned int pt:      8;

	/** varies by packet type */
	count uint //   5;
	/** padding flag */
	padding uint // 1;
	/** protocol version */
	version uint // 2;
	/** packet type */
	pt uint // 8;

	/** packet length in words, w/o this word */
	length uint16 //  16;
}

/** SDES item */
type RtcpSdesItem struct {
	/** type of item (rtcp_sdes_type_t) */
	Type int
	/** length of item (in octets) */
	length byte
	/** text, not null-terminated */
	data [1]byte
}

/** RTCP packet */
type RtcpPacket struct {

	/** common header */
	header RtcpHeader
	/** union of RTCP reports */
	r struct {
		/** sender report (SR) */
		sr struct {
			/** sr stat */
			srStat RtcpSRStat
			/** variable-length list rr stats */
			rrStat [1]RtcpRRStat
		}

		/** reception report (RR) */
		rr struct {
			/** receiver generating this report */
			ssrc uint32
			/** variable-length list rr stats */
			rrStat [1]RtcpRRStat
		}

		/** source description (SDES) */
		sdes struct {
			/** first SSRC/CSRC */
			ssrc uint32
			/** list of SDES items */
			item [1]RtcpSdesItem
		}

		/** BYE */
		bye struct {
			/** list of sources */
			ssrc [1]uint32
			/* optional length of reason string (in octets) */
			length byte
			/* optional reason string, not null-terminated */
			data [1]byte
		}
	}
}

/** Initialize RTCP header */
func RtcpHeaderInit(header *RtcpHeader, pt RtcpType) {
	header.version = RTP_VERSION
	header.padding = 0
	header.count = 0
	header.pt = uint(pt)
	header.length = 0
}

func (header *RtcpHeader) RtcpHeaderLengthSet(length int64) {
	/* htons 主机序转网络序  */
	header.length = binaryx.HToNS(uint16(length)/4 - 1)
}

/**
  htonl()--"Host to Network Long"  	小端输入, 大端输出   	uint32

  ntohl()--"Network to Host Long"  	大端输入, 小端输出   	uint32

  htons()--"Host to Network Short" 	小端输入, 大端输出	uint16

  ntohs()--"Network to Host Short" 	大端输入, 小端输出	uint16
*/

func RtcpSRHToN(srStat *RtcpSRStat) {
	srStat.ssrc = binaryx.HToNL(srStat.ssrc)
	srStat.ntpSec = binaryx.HToNL(srStat.ntpSec)
	srStat.ntpFrac = binaryx.HToNL(srStat.ntpFrac)
	srStat.rtpTs = binaryx.HToNL(srStat.rtpTs)
	srStat.sentPackets = binaryx.HToNL(srStat.sentPackets)
	srStat.sentOctets = binaryx.HToNL(srStat.sentOctets)
}

func RtcpSRNToH(srStat *RtcpSRStat) {
	srStat.ssrc = binaryx.NToHL(srStat.ssrc)
	srStat.ntpSec = binaryx.NToHL(srStat.ntpSec)
	srStat.ntpFrac = binaryx.NToHL(srStat.ntpFrac)
	srStat.rtpTs = binaryx.NToHL(srStat.rtpTs)
	srStat.sentPackets = binaryx.NToHL(srStat.sentPackets)
	srStat.sentOctets = binaryx.NToHL(srStat.sentOctets)
}

func RtcpRRHToN(rrStat *RtcpRRStat) {
	rrStat.ssrc = binaryx.HToNL(rrStat.ssrc)
	rrStat.lastSeq = binaryx.HToNL(rrStat.lastSeq)
	rrStat.jitter = binaryx.HToNL(rrStat.jitter)

	//#if (APR_IS_BIGENDIAN == 0)
	//rr_stat->lost = ((rr_stat->lost >> 16) & 0x000000ff) |
	//(rr_stat->lost & 0x0000ff00) |
	//((rr_stat->lost << 16) & 0x00ff0000);
	//#endif
}

func RtcpRRNToH(rrStat *RtcpRRStat) {
	rrStat.ssrc = binaryx.NToHL(rrStat.ssrc)
	rrStat.lastSeq = binaryx.NToHL(rrStat.lastSeq)
	rrStat.jitter = binaryx.NToHL(rrStat.jitter)

	//#if (APR_IS_BIGENDIAN == 0)
	//rr_stat->lost = ((rr_stat->lost >> 16) & 0x000000ff) |
	//(rr_stat->lost & 0x0000ff00) |
	//((rr_stat->lost << 16) & 0x00ff0000);
	//#endif
}
