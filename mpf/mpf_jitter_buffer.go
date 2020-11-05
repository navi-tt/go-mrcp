package mpf

/** Jitter buffer write result */
type JbResult = int

const (
	JB_OK                   JbResult = iota /**< successful write */
	JB_DISCARD_NOT_ALLIGNED                 /**< discarded write (frame isn't alligned to CODEC_FRAME_TIME_BASE) */
	JB_DISCARD_TOO_LATE                     /**< discarded write (frame is arrived too late) */
	JB_DISCARD_TOO_EARLY                    /**< discarded write (frame is arrived too early, buffer is full) */
)

type JitterBuffer struct {

	/* jitter buffer config */
	config *JbConfig
	/* codec to be used to dissect payload */
	codec *Codec

	/* cyclic raw data */
	RawData []byte
	/* frames (out of raw data) */
	frames []*Frame
	/* number of frames */
	frameCount int64
	/* frame timestamp units (samples) */
	frameTs int64
	/* frame size in bytes */
	frameSize int64

	/* playout delay in timetsamp units */
	playoutDelayTs uint32
	/* max playout delay in timetsamp units */
	maxPlayOutDelayTs uint32

	/* write should be synchronized (offset calculated) */
	writeSync byte
	/* write timestamp offset */
	writeTsOffset int32

	/* write pointer in timestamp units */
	writeTs uint32
	/* read pointer in timestamp units */
	readTs uint32

	/* min length of the buffer in timestamp units */
	minLengthTs int32
	/* max length of the buffer in timestamp units */
	maxLengthTs int32
	/* number of statistical measurements made */
	measurementCount uint32

	/* timestamp event starts at */
	eventWriteBaseTs uint32
	/* the first (base) frame of the event */
	eventWriteBase NamedEventFrame
	/* the last received update for the event */
	eventWriteUpdate *NamedEventFrame
}

/** Create jitter buffer */
func JitterBufferCreate(jbConfig *JbConfig, descriptor *CodecDescriptor, codec *Codec) *JitterBuffer {
	return nil
}

/** Destroy jitter buffer */
func JitterBufferDestroy(jb *JitterBuffer) error {
	return nil
}

/** Restart jitter buffer */
func JitterBufferRestart(jb *JitterBuffer) error {
	return nil
}

/** Write audio data to jitter buffer */
func (jb *JitterBuffer) JitterBufferWrite(buffer []byte, ts uint32, marker byte) JbResult {
	return 0
}

/** Write named event to jitter buffer */
func (jb *JitterBuffer) JitterBufferEventWrite(namedEvent *NamedEventFrame, ts uint32, marker byte) JbResult {
	return 0
}

/** Read media frame from jitter buffer */
func (jb *JitterBuffer) JitterBufferRead(mediaFrame *Frame) error {
	return nil
}

/** Get current playout delay */
func (jb *JitterBuffer) JitterBufferPlayOutDelayGet() uint32 {
	return 0
}
