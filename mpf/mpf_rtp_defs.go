package mpf

/** Used to calculate actual number of received packets (32bit) in
 * case seq number (16bit) wrapped around */
const RTP_SEQ_MOD = (1 << 16)

/** Number of max dropout packets (seq numbers) is used to trigger
 * either a drift in the seq numbers or a misorder packet */
const MAX_DROPOUT = 3000

/** Number of max misorder packets (seq numbers) is used to
 * differentiate a drift in the seq numbers from a misorder packet */
const MAX_MISORDER = 100

/** Restart receiver if threshold is reached */
const DISCARDED_TO_RECEIVED_RATIO_THRESHOLD = 30 /* 30% */
/** Deviation threshold is used to trigger a drift in timestamps */
const DEVIATION_THRESHOLD = 4000

/** This threshold is used to detect a new talkspurt */
const INTER_TALKSPURT_GAP = 1000 /* msec */

/** History of RTP receiver */
type RtpRXHistory struct {

	/** Updated on every seq num wrap around */
	seqCycles uint32

	/** First seq num received */
	seqNumBase uint16
	/** Max seq num received */
	seqNumMax uint16

	/** Last timestamp received */
	tsLast uint32
	/** Local time measured on last packet received */
	timeLast int64

	/** New ssrc, which is in probation */
	ssrcNew uint32
	/** Period of ssrc probation */
	ssrcProbation byte
}

/** Periodic history of RTP receiver (initialized after every N packets) */
type RtpRXPeriodicHistory struct {

	/** Number of packets received */
	receivedPrior uint32
	/** Number of packets expected */
	expectedPrior uint32
	/** Number of packets discarded */
	discardedPrior uint32

	/** Min jitter */
	jitterMin uint32
	/** Max jitter */
	jitterMax uint32
}

/** Reset RTP receiver history */
func RtpRXHistoryReset(rxHistory *RtpRXHistory) {
	rxHistory = new(RtpRXHistory)
}

/** Reset RTP receiver periodic history */
func RtpRXPeriodicHistoryReset(rxPeriodicHistory *RtpRXPeriodicHistory) {
	rxPeriodicHistory = new(RtpRXPeriodicHistory)
}

/** RTP receiver */
type RtpReceiver struct {

	/** Jitter buffer */
	jb *JitterBuffer

	/** RTCP statistics used in RR */
	rrStat RtcpRRStat
	/** RTP receiver statistics */
	stat RtpRXStat
	/** RTP history */
	history RtpRXHistory
	/** RTP periodic history */
	periodicHistory RtpRXPeriodicHistory
}

/** RTP transmitter */
type RtpTransmitter struct {
	/** PacketTization time in msec */
	ptime uint16

	/** Number of frames in a packet */
	packetFrames uint16
	/** Current number of frames */
	currentFrames uint16
	/** Samples in frames in timestamp units */
	samplesPerFrame uint32

	/** Indicate silence period among the talkspurts */
	inactivity byte
	/** Last seq number sent */
	lastSeqNum uint16
	/** Current timestamp (samples processed) */
	timestamp uint32
	/** Event timestamp base */
	timestampBase uint32

	/** RTP packet payload */
	packetData []byte
	/** RTP packet payload size */
	packetSize int64

	/** RTCP statistics used in SR */
	srStat RtcpSRStat
}

/** Initialize RTP receiver */
func RtpReceiverInit(receiver *RtpReceiver) {
	receiver.jb = nil

	RtcpRRStatReset(&(receiver.rrStat))
	RtpRXStatReset(&(receiver.stat))
	RtpRXHistoryReset(&(receiver.history))
	RtpRXPeriodicHistoryReset(&(receiver.periodicHistory))
}

/** Initialize RTP transmitter */
func RtpTransmitterInit(transmitter *RtpTransmitter) {
	transmitter.ptime = 0

	transmitter.packetFrames = 0
	transmitter.currentFrames = 0
	transmitter.samplesPerFrame = 0

	transmitter.inactivity = 0
	transmitter.lastSeqNum = 0
	transmitter.timestamp = 0
	transmitter.timestampBase = 0

	transmitter.packetData = nil
	transmitter.packetSize = 0

	RtcpSRStatReset(&(transmitter.srStat))
}
