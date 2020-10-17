package resources

import (
	"github.com/navi-tt/go-mrcp/apr/memory"
	"github.com/navi-tt/go-mrcp/mrcp"
	"github.com/navi-tt/go-mrcp/mrcp/control/resource"
	"github.com/navi-tt/go-mrcp/mrcp/message/header"
)

/**
 * @Author: Liu xiangpeng
 * @Date: 2020/10/17 2:24 下午
 */

/** MRCP recognizer header fields */

type MRCPRecognizerHeaderId = int

const (
	RECOGNIZER_HEADER_CONFIDENCE_THRESHOLD MRCPRecognizerHeaderId = iota
	RECOGNIZER_HEADER_SENSITIVITY_LEVEL
	RECOGNIZER_HEADER_SPEED_VS_ACCURACY
	RECOGNIZER_HEADER_N_BEST_LIST_LENGTH
	RECOGNIZER_HEADER_NO_INPUT_TIMEOUT
	RECOGNIZER_HEADER_RECOGNITION_TIMEOUT
	RECOGNIZER_HEADER_WAVEFORM_URI
	RECOGNIZER_HEADER_COMPLETION_CAUSE
	RECOGNIZER_HEADER_RECOGNIZER_CONTEXT_BLOCK
	RECOGNIZER_HEADER_START_INPUT_TIMERS
	RECOGNIZER_HEADER_SPEECH_COMPLETE_TIMEOUT
	RECOGNIZER_HEADER_SPEECH_INCOMPLETE_TIMEOUT
	RECOGNIZER_HEADER_DTMF_INTERDIGIT_TIMEOUT
	RECOGNIZER_HEADER_DTMF_TERM_TIMEOUT
	RECOGNIZER_HEADER_DTMF_TERM_CHAR
	RECOGNIZER_HEADER_FAILED_URI
	RECOGNIZER_HEADER_FAILED_URI_CAUSE
	RECOGNIZER_HEADER_SAVE_WAVEFORM
	RECOGNIZER_HEADER_NEW_AUDIO_CHANNEL
	RECOGNIZER_HEADER_SPEECH_LANGUAGE

	/** Additional header fields for MRCP v2 */
	RECOGNIZER_HEADER_INPUT_TYPE
	RECOGNIZER_HEADER_INPUT_WAVEFORM_URI
	RECOGNIZER_HEADER_COMPLETION_REASON
	RECOGNIZER_HEADER_MEDIA_TYPE
	RECOGNIZER_HEADER_VER_BUFFER_UTTERANCE
	RECOGNIZER_HEADER_RECOGNITION_MODE
	RECOGNIZER_HEADER_CANCEL_IF_QUEUE
	RECOGNIZER_HEADER_HOTWORD_MAX_DURATION
	RECOGNIZER_HEADER_HOTWORD_MIN_DURATION
	RECOGNIZER_HEADER_INTERPRET_TEXT
	RECOGNIZER_HEADER_DTMF_BUFFER_TIME
	RECOGNIZER_HEADER_CLEAR_DTMF_BUFFER
	RECOGNIZER_HEADER_EARLY_NO_MATCH
	RECOGNIZER_HEADER_NUM_MIN_CONSISTENT_PRONUNCIATIONS
	RECOGNIZER_HEADER_CONSISTENCY_THRESHOLD
	RECOGNIZER_HEADER_CLASH_THRESHOLD
	RECOGNIZER_HEADER_PERSONAL_GRAMMAR_URI
	RECOGNIZER_HEADER_ENROLL_UTTERANCE
	RECOGNIZER_HEADER_PHRASE_ID
	RECOGNIZER_HEADER_PHRASE_NL
	RECOGNIZER_HEADER_WEIGHT
	RECOGNIZER_HEADER_SAVE_BEST_WAVEFORM
	RECOGNIZER_HEADER_NEW_PHRASE_ID
	RECOGNIZER_HEADER_CONFUSABLE_PHRASES_URI
	RECOGNIZER_HEADER_ABORT_PHRASE_ENROLLMENT

	RECOGNIZER_HEADER_COUNT
)

/** MRCP recognizer completion-cause */
type MRCPRecognizerCompletionCause = int

const (
	RECOGNIZER_COMPLETION_CAUSE_SUCCESS                 MRCPRecognizerCompletionCause = 0
	RECOGNIZER_COMPLETION_CAUSE_NO_MATCH                MRCPRecognizerCompletionCause = 1
	RECOGNIZER_COMPLETION_CAUSE_NO_INPUT_TIMEOUT        MRCPRecognizerCompletionCause = 2
	RECOGNIZER_COMPLETION_CAUSE_RECOGNITION_TIMEOUT     MRCPRecognizerCompletionCause = 3
	RECOGNIZER_COMPLETION_CAUSE_GRAM_LOAD_FAILURE       MRCPRecognizerCompletionCause = 4
	RECOGNIZER_COMPLETION_CAUSE_GRAM_COMP_FAILURE       MRCPRecognizerCompletionCause = 5
	RECOGNIZER_COMPLETION_CAUSE_ERROR                   MRCPRecognizerCompletionCause = 6
	RECOGNIZER_COMPLETION_CAUSE_SPEECH_TOO_EARLY        MRCPRecognizerCompletionCause = 7
	RECOGNIZER_COMPLETION_CAUSE_TOO_MUCH_SPEECH_TIMEOUT MRCPRecognizerCompletionCause = 8
	RECOGNIZER_COMPLETION_CAUSE_URI_FAILURE             MRCPRecognizerCompletionCause = 9
	RECOGNIZER_COMPLETION_CAUSE_LANGUAGE_UNSUPPORTED    MRCPRecognizerCompletionCause = 10

	/** Additional completion-cause for MRCP v2 */
	RECOGNIZER_COMPLETION_CAUSE_SEMANTICS_FAILURE       MRCPRecognizerCompletionCause = 12
	RECOGNIZER_COMPLETION_CAUSE_PARTIAL_MATCH           MRCPRecognizerCompletionCause = 13
	RECOGNIZER_COMPLETION_CAUSE_PARTIAL_MATCH_MAXTIME   MRCPRecognizerCompletionCause = 14
	RECOGNIZER_COMPLETION_CAUSE_NO_MATCH_MAXTIME        MRCPRecognizerCompletionCause = 15
	RECOGNIZER_COMPLETION_CAUSE_GRAM_DEFINITION_FAILURE MRCPRecognizerCompletionCause = 16
	RECOGNIZER_COMPLETION_CAUSE_CANCELLED               MRCPRecognizerCompletionCause = 11

	RECOGNIZER_COMPLETION_CAUSE_COUNT   MRCPRecognizerCompletionCause = 17
	RECOGNIZER_COMPLETION_CAUSE_UNKNOWN                               = RECOGNIZER_COMPLETION_CAUSE_COUNT
)

/** MRCP recognizer methods */
type MRCPRecognizerMethodId = int

const (
	RECOGNIZER_SET_PARAMS MRCPRecognizerMethodId = iota
	RECOGNIZER_GET_PARAMS
	RECOGNIZER_DEFINE_GRAMMAR
	RECOGNIZER_RECOGNIZE
	RECOGNIZER_INTERPRET
	RECOGNIZER_GET_RESULT
	RECOGNIZER_START_INPUT_TIMERS
	RECOGNIZER_STOP
	RECOGNIZER_START_PHRASE_ENROLLMENT
	RECOGNIZER_ENROLLMENT_ROLLBACK
	RECOGNIZER_END_PHRASE_ENROLLMENT
	RECOGNIZER_MODIFY_PHRASE
	RECOGNIZER_DELETE_PHRASE

	RECOGNIZER_METHOD_COUNT
)

/** MRCP recognizer events */
type MRCPRecognizerEventId = int

const (
	RECOGNIZER_START_OF_INPUT MRCPRecognizerEventId = iota
	RECOGNIZER_RECOGNITION_COMPLETE
	RECOGNIZER_INTERPRETATION_COMPLETE

	RECOGNIZER_EVENT_COUNT
)

/** MRCP recognizer-header */
type MRCPRecognizerHeader struct {
	/** Tells the recognizer resource what confidence level the client considers a
	  successful match */
	ConfidenceThreshold float64
	/** To filter out background noise and not mistake it for speech */
	SensitivityLevel float64
	/** Tunable towards Performance or Accuracy */
	SpeedVsAccuracy float64
	/** The client, by setting this header, can ask the recognition resource
	to send it more  than 1 alternative */
	NBestListLength int64
	/** The client can use the no-input-timeout header to set this timeout */
	NoInputTimeout int64
	/** The client can use the recognition-timeout header to set this timeout */
	RecognitionTimeout int64
	/** MUST be present in the RECOGNITION-COMPLETE event if the Save-Waveform
	header was set to true */
	WaveformUri string
	/** MUST be part of a RECOGNITION-COMPLETE, event coming from
	  the recognizer resource to the client */
	CompletionCause MRCPRecognizerCompletionCause
	/** MAY be sent as part of the SET-PARAMS or GET-PARAMS request */
	RecognizerContextBlock string
	/** MAY be sent as part of the RECOGNIZE request. A value of false tells
	the recognizer to start recognition, but not to start the no-input timer yet */
	StartInputTimers bool
	/** Specifies the length of silence required following user
	  speech before the speech recognizer finalizes a result */
	SpeechCompleteTimeout int64
	/** Specifies the required length of silence following user
	  speech after which a recognizer finalizes a result */
	SpeechIncompleteTimeout int64
	/** Specifies the inter-digit timeout value to use when
	  recognizing DTMF input */
	DtmfInterdigitTimeout int64
	/** Specifies the terminating timeout to use when
	recognizing DTMF input*/
	DtmfTermTimeout int64
	/** Specifies the terminating DTMF character for DTMF input
	  recognition */
	DtmfTermChar byte
	/** When a recognizer needs to fetch or access a URI and the access fails
	  the server SHOULD provide the failed URI in this header in the method response*/
	FailedUri string
	/** When a recognizer method needs a recognizer to fetch or access a URI
	  and the access fails the server MUST provide the URI specific or
	  protocol specific response code for the URI in the Failed-URI header */
	FailedUriCause string
	/** Allows the client to request the recognizer resource to
	  save the audio input to the recognizer */
	SaveWaveform bool
	/** MAY be specified in a RECOGNIZE request and allows the
	  client to tell the server that, from this point on, further input
	  audio comes from a different audio source */
	NewAudioChannel bool
	/** Specifies the language of recognition grammar data within
	  a session or request, if it is not specified within the data */
	SpeechLanguage string

	/** Additional header fields for MRCP v2 */
	/** Specifies if the input that caused a barge-in was DTMF or speech */
	InputType string
	/** Optional header specifies a URI pointing to audio content to be
	  processed by the RECOGNIZE operation */
	InputWaveformUri string
	/** MAY be specified in a RECOGNITION-COMPLETE event coming from
	  the recognizer resource to the client */
	CompletionReason string
	/** Tells the server resource the Media Type in which to store captured
	audio such as the one captured and returned by the Waveform-URI header */
	MediaType string
	/** Lets the client request the server to buffer the
	  utterance associated with this recognition request into a buffer
	  available to a co-resident verification resource */
	VerBufferUtterance bool
	/** Specifies what mode the RECOGNIZE method will operate in */
	RecognitionMode string
	/** Specifies what will happen if the client attempts to
	  invoke another RECOGNIZE method when this RECOGNIZE request is
	  already in progress for the resource*/
	CancelIfQueue bool
	/** Specifies the maximum length of an utterance (in seconds) that will
	  be considered for Hotword recognition */
	HotWordMaxDuration int64
	/** Specifies the minimum length of an utterance (in seconds) that will
	  be considered for Hotword recognition */
	HotWordMinDuration int64
	/** Provides a pointer to the text for which a natural language interpretation is desired */
	InterpretText string
	/** MAY be specified in a GET-PARAMS or SET-PARAMS method and
	  is used to specify the size in time, in milliseconds, of the
	  typeahead buffer for the recognizer */
	DtmfBufferTime int64
	/** MAY be specified in a RECOGNIZE method and is used to
	  tell the recognizer to clear the DTMF type-ahead buffer before
	  starting the recognize */
	ClearDtmfBuffer bool
	/** MAY be specified in a RECOGNIZE method and is used to
	  tell the recognizer that it MUST not wait for the end of speech
	  before processing the collected speech to match active grammars */
	EarlyNoMatch bool
	/** MAY be specified in a START-PHRASE-ENROLLMENT, "SET-PARAMS", or
	"GET-PARAMS" method and is used to specify the minimum number of
	consistent pronunciations that must be obtained to voice enroll a new phrase */
	NumMinConsistentPronunciations int64
	/** MAY be sent as part of the START-PHRASE-ENROLLMENT,"SET-PARAMS", or
	"GET-PARAMS" method and is used during voice-enrollment to specify how similar
	to a previously enrolled pronunciation of the same phrase an utterance needs
	to be in order to be considered "consistent" */
	ConsistencyThreshold float64
	/** MAY be sent as part of the START-PHRASE-ENROLLMENT, SET-PARAMS, or
	"GET-PARAMS" method and is used during voice-enrollment to specify
	how similar the pronunciations of two different phrases can be
	before they are considered to be clashing */
	ClashThreshold float64
	/** Specifies the speaker-trained grammar to be used or
	referenced during enrollment operations */
	PersonalGrammarUri string
	/** MAY be specified in the RECOGNIZE method. If this header
	is set to "true" and an Enrollment is active, the RECOGNIZE command
	MUST add the collected utterance to the personal grammar that is
	being enrolled */
	EnrollUtterance bool
	/** Identifies a phrase in an existing personal grammar for which
	enrollment is desired.  It is also returned to the client in the
	RECOGNIZE complete event */
	PhraseId string
	/** Specifies the interpreted text to be returned when the
	phrase is recognized */
	PhraseNl string
	/** Represents the occurrence likelihood of a phrase in an enrolled grammar */
	Weight float64
	/** Allows the client to request the recognizer resource to
	save the audio stream for the best repetition of the phrase that was
	used during the enrollment session */
	SaveBestWaveform bool
	/** Replaces the id used to identify the phrase in a personal grammar */
	NewPhraseId string
	/** Specifies a grammar that defines invalid phrases for enrollment */
	ConfusablePhrasesUri string
	/** Can optionally be specified in the END-PHRASE-ENROLLMENT
	method to abort the phrase enrollment, rather than committing the
	phrase to the personal grammar */
	AbortPhraseEnrollment bool
}

/** Get recognizer header vtable */
func MRCPRecognizerHeaderVTableGet(v mrcp.Version) *header.MRCPHeaderVTable {
	return nil
}

/** Get recognizer completion cause string */
func MRCPRecognizerCompletionCauseGet(cause MRCPRecognizerCompletionCause, v mrcp.Version) string {
	return ""
}

/** Create MRCP recognizer resource */
func MRCPRecognizerResourceCreate(pool *memory.AprPool) *resource.MRCPResource {
	return nil
}
