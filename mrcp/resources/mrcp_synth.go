package resources

import (
	"github.com/navi-tt/go-mrcp/mrcp"
	"github.com/navi-tt/go-mrcp/mrcp/control/resource"
	"github.com/navi-tt/go-mrcp/mrcp/message/header"
)

/**
 * @Author: Liu xiangpeng
 * @Date: 2020/10/17 2:24 下午
 */

/** MRCP synthesizer header fields */
type MRCPSynthesizerHeaderId = int

const (
	SYNTHESIZER_HEADER_JUMP_SIZE MRCPSynthesizerHeaderId = iota
	SYNTHESIZER_HEADER_KILL_ON_BARGE_IN
	SYNTHESIZER_HEADER_SPEAKER_PROFILE
	SYNTHESIZER_HEADER_COMPLETION_CAUSE
	SYNTHESIZER_HEADER_COMPLETION_REASON
	SYNTHESIZER_HEADER_VOICE_GENDER
	SYNTHESIZER_HEADER_VOICE_AGE
	SYNTHESIZER_HEADER_VOICE_VARIANT
	SYNTHESIZER_HEADER_VOICE_NAME
	SYNTHESIZER_HEADER_PROSODY_VOLUME
	SYNTHESIZER_HEADER_PROSODY_RATE
	SYNTHESIZER_HEADER_SPEECH_MARKER
	SYNTHESIZER_HEADER_SPEECH_LANGUAGE
	SYNTHESIZER_HEADER_FETCH_HINT
	SYNTHESIZER_HEADER_AUDIO_FETCH_HINT
	SYNTHESIZER_HEADER_FAILED_URI
	SYNTHESIZER_HEADER_FAILED_URI_CAUSE
	SYNTHESIZER_HEADER_SPEAK_RESTART
	SYNTHESIZER_HEADER_SPEAK_LENGTH
	SYNTHESIZER_HEADER_LOAD_LEXICON
	SYNTHESIZER_HEADER_LEXICON_SEARCH_ORDER

	SYNTHESIZER_HEADER_COUNT
)

/** Speech-units */
type MRCPSpeechUnit = int

const (
	SPEECH_UNIT_SECOND MRCPSpeechUnit = iota
	SPEECH_UNIT_WORD
	SPEECH_UNIT_SENTENCE
	SPEECH_UNIT_PARAGRAPH

	SPEECH_UNIT_COUNT
)

/** Speech-length types */
type MRCPSpeechLengthType = int

const (
	SPEECH_LENGTH_TYPE_TEXT MRCPSpeechLengthType = iota
	SPEECH_LENGTH_TYPE_NUMERIC_POSITIVE
	SPEECH_LENGTH_TYPE_NUMERIC_NEGATIVE

	SPEECH_LENGTH_TYPE_UNKNOWN
)

/** MRCP voice-gender */
type MRCPVoiceGender = int

const (
	VOICE_GENDER_MALE MRCPVoiceGender = iota
	VOICE_GENDER_FEMALE
	VOICE_GENDER_NEUTRAL

	VOICE_GENDER_COUNT
	VOICE_GENDER_UNKNOWN = VOICE_GENDER_COUNT
)

/** Prosody-volume type */
type MRCPProsodyVolumeType = int

const (
	PROSODY_VOLUME_TYPE_LABEL MRCPProsodyVolumeType = iota
	PROSODY_VOLUME_TYPE_NUMERIC
	PROSODY_VOLUME_TYPE_RELATIVE_CHANGE

	PROSODY_VOLUME_TYPE_UNKNOWN
)

/** Prosody-rate type */
type MRCPProsodyRateType = int

const (
	PROSODY_RATE_TYPE_LABEL MRCPProsodyRateType = iota
	PROSODY_RATE_TYPE_RELATIVE_CHANGE

	PROSODY_RATE_TYPE_UNKNOWN
)

/** Prosody-volume */
type MRCPProsodyVolumeLabel = int

const (
	PROSODY_VOLUME_SILENT MRCPProsodyVolumeLabel = iota
	PROSODY_VOLUME_XSOFT
	PROSODY_VOLUME_SOFT
	PROSODY_VOLUME_MEDIUM
	PROSODY_VOLUME_LOUD
	PROSODY_VOLUME_XLOUD
	PROSODY_VOLUME_DEFAULT

	PROSODY_VOLUME_COUNT
	PROSODY_VOLUME_UNKNOWN = PROSODY_VOLUME_COUNT
)

/** Prosody-rate */
type MRCPProsodyRateLabel = int

const (
	PROSODY_RATE_XSLOW MRCPProsodyRateLabel = iota
	PROSODY_RATE_SLOW
	PROSODY_RATE_MEDIUM
	PROSODY_RATE_FAST
	PROSODY_RATE_XFAST
	PROSODY_RATE_DEFAULT

	PROSODY_RATE_COUNT
	PROSODY_RATE_UNKNOWN = PROSODY_RATE_COUNT
)

/** Synthesizer completion-cause specified in SPEAK-COMPLETE event */
type MRCPSynthCompletionCause = int

const (
	SYNTHESIZER_COMPLETION_CAUSE_NORMAL               MRCPSynthCompletionCause = 0
	SYNTHESIZER_COMPLETION_CAUSE_BARGE_IN             MRCPSynthCompletionCause = 1
	SYNTHESIZER_COMPLETION_CAUSE_PARSE_FAILURE        MRCPSynthCompletionCause = 2
	SYNTHESIZER_COMPLETION_CAUSE_URI_FAILURE          MRCPSynthCompletionCause = 3
	SYNTHESIZER_COMPLETION_CAUSE_ERROR                MRCPSynthCompletionCause = 4
	SYNTHESIZER_COMPLETION_CAUSE_LANGUAGE_UNSUPPORTED MRCPSynthCompletionCause = 5
	SYNTHESIZER_COMPLETION_CAUSE_LEXICON_LOAD_FAILURE MRCPSynthCompletionCause = 6
	SYNTHESIZER_COMPLETION_CAUSE_CANCELLED            MRCPSynthCompletionCause = 7

	SYNTHESIZER_COMPLETION_CAUSE_COUNT   MRCPSynthCompletionCause = 8
	SYNTHESIZER_COMPLETION_CAUSE_UNKNOWN MRCPSynthCompletionCause = SYNTHESIZER_COMPLETION_CAUSE_COUNT
)

/** MRCP synthesizer methods */
type MRCPSynthesizerMethodId = int

const (
	SYNTHESIZER_SET_PARAMS MRCPSynthesizerMethodId = iota
	SYNTHESIZER_GET_PARAMS
	SYNTHESIZER_SPEAK
	SYNTHESIZER_STOP
	SYNTHESIZER_PAUSE
	SYNTHESIZER_RESUME
	SYNTHESIZER_BARGE_IN_OCCURRED
	SYNTHESIZER_CONTROL
	SYNTHESIZER_DEFINE_LEXICON

	SYNTHESIZER_METHOD_COUNT
)

/** MRCP synthesizer events */
type MRCPSynthesizerEventId = int

const (
	SYNTHESIZER_SPEECH_MARKER MRCPSynthesizerEventId = iota
	SYNTHESIZER_SPEAK_COMPLETE

	SYNTHESIZER_EVENT_COUNT
)

/** Numeric speech-length */
type MRCPNumericSpeechLength struct {
	Length int64          // The length
	Unit   MRCPSpeechUnit //  The unit (second/word/sentence/paragraph)
}

/** Definition of speech-length value */
type MRCPSpeechLengthValue struct {
	Type MRCPSpeechLengthType // Speech-length type (numeric/text)

	Value struct { // Speech-length value (either numeric or text)
		Tag     string                  // Text speech-length
		Numeric MRCPNumericSpeechLength // Numeric speech-length
	}
}

/** MRCP voice-param */
type MRCPVoiceParam struct {
	Gender  MRCPVoiceGender // Voice gender (male/female/neutral)
	Age     int64           // Voice age
	Variant int64           // Voice variant
	Name    string          // Voice name
}

/** MRCP prosody-volume */
type MRCPProsodyVolume struct {
	Type MRCPProsodyVolumeType // prosody-volume type (one of label,numeric,relative change)

	Value struct { // prosody-volume value
		Label    MRCPProsodyVolumeLabel // one of "silent", "x-soft", ...
		Numeric  float64                // numeric value
		Relative float64                // relative change
	}
}

/** MRCP prosody-rate */
type MRCPProsodyRate struct {
	Type MRCPProsodyRateType // prosody-rate type (one of label, relative change)

	Value struct { // prosody-rate value
		Label    MRCPProsodyRateLabel // one of "x-slow", "slow", ...
		Relative float64              // relative change
	}
}

/** MRCP prosody-param */
type MRCPProsodyParam struct {
	Volume MRCPProsodyVolume // Prosofy volume
	rate   MRCPProsodyRate   // Prosofy rate
}

/** MRCP synthesizer-header */
type MRCPSynthHeader struct {

	/** MAY be specified in a CONTROL method and controls the
	  amount to jump forward or backward in an active "SPEAK" request */
	JumpSize MRCPSpeechLengthValue
	/** MAY be sent as part of the "SPEAK" method to enable kill-
	  on-barge-in support */
	KillOnBargeIn bool
	/** MAY be part of the "SET-PARAMS"/"GET-PARAMS" or "SPEAK"
	  request from the client to the server and specifies a URI which
	  references the profile of the speaker */
	SpeakerProfile string
	/** MUST be specified in a "SPEAK-COMPLETE" event coming from
	  the synthesizer resource to the client */
	CompletionCause MRCPSynthCompletionCause
	/** MAY be specified in a "SPEAK-COMPLETE" event coming from
	  the synthesizer resource to the client */
	CompletionReason string
	/** This set of header fields defines the voice of the speaker */
	voice_param MRCPVoiceParam
	/** This set of header fields defines the prosody of the speech */
	prosody_param MRCPProsodyParam
	/** Contains timestamp information in a "timestamp" field */
	SpeechMarker string
	/** specifies the default language of the speech data if the
	  language is not specified in the markup */
	SpeechLanguage string
	/** When the synthesizer needs to fetch documents or other resources like
	  speech markup or audio files, this header controls the corresponding
	  URI access properties */
	FetchHint string
	/** When the synthesizer needs to fetch documents or other resources like
	  speech audio files, this header controls the corresponding URI access
	  properties */
	AudioFetchHint string
	/** When a synthesizer method needs a synthesizer to fetch or access a
	  URI and the access fails, the server SHOULD provide the failed URI in
	  this header in the method response */
	FailedUri string
	/** When a synthesizer method needs a synthesizer to fetch or access a
	  URI and the access fails the server MUST provide the URI specific or
	  protocol specific response code for the URI in the Failed-URI header
	  in the method response through this header */
	FailedUriCause string
	/** When a CONTROL request to jump backward is issued to a currently
	  speaking synthesizer resource, and the target jump point is before
	  the start of the current "SPEAK" request, the current "SPEAK" request
	  MUST restart */
	SpeakRestart bool
	/** MAY be specified in a CONTROL method to control the
	  length of speech to speak, relative to the current speaking point in
	  the currently active "SPEAK" request */
	SpeakLength MRCPSpeechLengthValue
	/** Used to indicate whether a lexicon has to be loaded or unloaded */
	LoadLexicon bool
	/** Used to specify a list of active Lexicon URIs and the
	  search order among the active lexicons */
	LexiconSearchOrder string
}

/** Get synthesizer header vtable */
func MRCPSynthHeaderVTableGet(v mrcp.Version) *header.MRCPHeaderVTable {
	return nil
}

/** Get synthesizer completion cause string */
func MRCPSynthCompletionCauseGet(cause MRCPSynthCompletionCause, v mrcp.Version) string {
	return ""
}

/** Create MRCP synthesizer resource */
func MRCPSynthResourceCreate() *resource.MRCPResource {
	return nil
}
