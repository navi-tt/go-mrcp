package mpf

import "sync"

/** Max DTMF digits waiting to be sent */
const MPF_DTMFGEN_QUEUE_LEN = 32

/** See RFC4733 */
const DTMF_EVENT_VOLUME = 10

/** Amplitude of single sine wave from tone generator */
const DTMF_SINE_AMPLITUDE = 12288

/** DTMF generator band */
type DtmfGeneratorBand = int

const (
	/** Generate tones in-band */
	MPF_DTMF_GENERATOR_INBAND DtmfGeneratorBand = 0x1
	/** Generate named events out-of-band */
	MPF_DTMF_GENERATOR_OUTBAND DtmfGeneratorBand = 0x2
	/** Generate both tones and named events */
	MPF_DTMF_GENERATOR_BOTH = MPF_DTMF_GENERATOR_INBAND | MPF_DTMF_GENERATOR_OUTBAND
)

/** State of the DTMF generator */
type DtmfGeneratorState = int

const (
	/** Ready to generate next digit in queue */
	DTMF_GEN_STATE_IDLE DtmfGeneratorState = iota
	/** Generating tone */
	DTMF_GEN_STATE_TONE
	/** Retransmitting final RTP packet */
	DTMF_GEN_STATE_ENDING
	/** Generating silence between tones */
	DTMF_GEN_STATE_SILENCE
)

/**
 * Sine wave generator (second-order IIR filter) state:
 *
 * s(t) = Amp*sin(2*pi*f_tone/f_sampling*t)
 *
 * s(t) = coef * s(t-1) - s(t-2); s(0)=0; s(1)=Amp*sin(2*pi*f_tone/f_sampling)
 */
type SineState struct {
	/** coef = cos(2*pi*f_tone/f_sampling) */
	Coef float64
	/** s(t-2) @see sine_state_t */
	S1 float64
	/** s(t-1) @see sine_state_t */
	S2 float64
}

/** Media Processing Framework's Dual Tone Multiple Frequncy generator */
type DtmfGenerator struct {

	/** Generator state */
	state DtmfGeneratorState
	/** In-band or out-of-band */
	band DtmfGeneratorBand
	/** Mutex to guard the queue */
	mutex sync.Mutex
	/** Queue of digits to generate */
	queue [MPF_DTMFGEN_QUEUE_LEN + 1]byte
	/** DTMF event_id according to RFC4733 */
	EventId uint8
	/** Duration in RTP units: (sample_rate / 1000) * milliseconds */
	ToneDuration uint32
	/** Duration of inter-digit silence @see tone_duration */
	SilenceDuration uint32
	/** Multipurpose counter; mostly in RTP time units */
	Counter uint32
	/** Frame duration in RTP units */
	FrameDuration uint32
	/** RTP named event duration (0..0xFFFF) */
	EventDuration uint32
	/** Set MPF_MARKER_NEW_SEGMENT in the next event frame */
	NewSegment bool
	/** Lower frequency generator */
	sine1 SineState
	/** Higher frequency generator */
	sine2 SineState
	/** Sampling rate of audio in Hz; used in tone generator */
	SampleRateAudio uint32
	/** Sampling rate of telephone-events in Hz; used for timing */
	SampleRateEvents uint32
	/** How often to issue event packet */
	EventsPtime uint32
	/** Milliseconds elapsed since last event packet */
	SinceLastEvent uint32
}

/**
 * Create MPF DTMF generator (advanced).
 * @param stream      A stream to transport digits via.
 * @param band        MPF_DTMF_GENERATOR_INBAND or MPF_DTMF_GENERATOR_OUTBAND
 * @param tone_ms     Tone duration in milliseconds.
 * @param silence_ms  Inter-digit silence in milliseconds.
 * @param pool        Memory pool to allocate DTMF generator from.
 * @return The object or NULL on error.
 * @see mpf_dtmf_generator_create
 */
func DtmfGeneratorCreateEx(stream *AudioStream, band DtmfGeneratorBand, toneMs uint32, silenceMs uint32) *DtmfGenerator {
	return nil
}

/**
 * Create MPF DTMF generator (simple). Calls mpf_dtmf_generator_create_ex
 * with band = MPF_DTMF_GENERATOR_OUTBAND if supported by the stream or
 * MPF_DTMF_GENERATOR_INBAND otherwise, tone_ms = 70, silence_ms = 50.
 * @param stream      A stream to transport digits via.
 * @param pool        Memory pool to allocate DTMF generator from.
 * @return The object or NULL on error.
 * @see mpf_dtmf_generator_create_ex
 */
func DtmfGeneratorCreate(stream *AudioStream) *DtmfGenerator {
	var band DtmfGeneratorBand = MPF_DTMF_GENERATOR_INBAND
	if stream.RXEventDescriptor != nil {
		band = MPF_DTMF_GENERATOR_OUTBAND
	}
	return DtmfGeneratorCreateEx(stream, band, 70, 50)
}

/**
 * Add DTMF digits to the queue.
 * @param generator The generator.
 * @param digits    DTMF character sequence [0-9*#A-D].
 * @return TRUE if ok, FALSE if there are too many digits.
 */
func (g *DtmfGenerator) DtmfGeneratorEnqueue(digits byte) error {
	return nil
}

/**
 * Empty the queue and immediately stop generating.
 * @param generator The generator.
 */
func (g *DtmfGenerator) DtmfGeneratorReset() {}

/**
 * Check state of the generator.
 * @param generator The generator.
 * @return TRUE if generating a digit or there are digits waiting in queue.
 * FALSE if the queue is empty or generating silence after the last digit.
 */
func (g *DtmfGenerator) DtmfGeneratorSending() bool {
	return false
}

/**
 * Put frame into the stream.
 * @param generator The generator.
 * @param frame     Frame object passed in stream_read().
 * @return TRUE if frame with tone (both in-band and out-of-band) was generated,
 * FALSE otherwise. In contrast to mpf_dtmf_generator_sending, returns FALSE even
 * if generating inter-digit silence. In other words returns TRUE iff the frame
 * object was filled with data. This method MUST be called for each frame for
 * proper timing.
 */
func (g *DtmfGenerator) DtmfGeneratorPutFrame(frame *Frame) error {
	return nil
}

/**
 * Free all resources associated with the generator.
 * @param generator The generator.
 */
func DtmfGeneratorDestroy(g *DtmfGenerator) {

}
