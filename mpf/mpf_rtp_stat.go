package mpf

/** RTP receiver statistics */
type RtpRXStat struct {

	/** number of valid RTP packets received */
	receivedPackets uint32
	/** number of invalid RTP packets received */
	invalidPackets uint32

	/** number of discarded in jitter buffer packets */
	discardedPackets uint32
	/** number of ignored packets */
	ignoredPackets uint32

	/** number of lost in network packets */
	lostPackets uint32

	/** number of restarts */
	restarts byte
}

/** RTCP statistics used in Sender Report (SR)  */
type RtcpSRStat struct {
	/** sender source identifier */
	ssrc uint32
	/** NTP timestamp (seconds) */
	ntpSec uint32
	/** NTP timestamp (fractions) */
	ntpFrac uint32
	/** RTP timestamp */
	rtpTs uint32
	/** packets sent */
	sentPackets uint32
	/** octets (bytes) sent */
	sentOctets uint32
}

/** RTCP statistics used in Receiver Report (RR) */
type RtcpRRStat struct {

	/** source identifier of RTP stream being received */
	ssrc uint32
	/** fraction lost since last SR/RR */
	fraction uint32 // 8
	/** cumulative number of packets lost (signed!) */
	lost int32 //24;
	/** extended last sequence number received */
	lastSeq uint32
	/** interArrival jitter (RFC3550) */
	jitter uint32
	/** last SR packet from this source */
	lsr uint32
	/** delay since last SR packet */
	dlsr uint32
}

/** Reset RTCP SR statistics */
func RtcpSRStatReset(srStat *RtcpSRStat) {
	srStat = new(RtcpSRStat)
}

/** Reset RTCP RR statistics */
func RtcpRRStatReset(rrStat *RtcpRRStat) {
	rrStat = new(RtcpRRStat)
}

/** Reset RTP receiver statistics */
func RtpRXStatReset(rxStat *RtpRXStat) {
	rxStat = new(RtpRXStat)
}
