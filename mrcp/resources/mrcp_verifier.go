package resources

import (
	"github.com/navi-tt/go-mrcp/apr/memory"
	"github.com/navi-tt/go-mrcp/mrcp"
	"github.com/navi-tt/go-mrcp/mrcp/control/resource"
	"github.com/navi-tt/go-mrcp/mrcp/message/header"
)

/** MRCP verifier header fields */
type MRCPVerifierHeaderId = int

const (
	VERIFIER_HEADER_REPOSITORY_URIstring MRCPVerifierHeaderId = iota
	VERIFIER_HEADER_VOICEPRINT_IDENTIFIERstring
	VERIFIER_HEADER_VERIFICATION_MODEstring
	VERIFIER_HEADER_ADAPT_MODELstring
	VERIFIER_HEADER_ABORT_MODELstring
	VERIFIER_HEADER_MIN_VERIFICATION_SCOREstring
	VERIFIER_HEADER_NUM_MIN_VERIFICATION_PHRASESstring
	VERIFIER_HEADER_NUM_MAX_VERIFICATION_PHRASESstring
	VERIFIER_HEADER_NO_INPUT_TIMEOUTstring
	VERIFIER_HEADER_SAVE_WAVEFORMstring
	VERIFIER_HEADER_MEDIA_TYPEstring
	VERIFIER_HEADER_WAVEFORM_URIstring
	VERIFIER_HEADER_VOICEPRINT_EXISTSstring
	VERIFIER_HEADER_VER_BUFFER_UTTERANCEstring
	VERIFIER_HEADER_INPUT_WAVEFORM_URIstring
	VERIFIER_HEADER_COMPLETION_CAUSEstring
	VERIFIER_HEADER_COMPLETION_REASONstring
	VERIFIER_HEADER_SPEECH_COMPLETE_TIMEOUTstring
	VERIFIER_HEADER_NEW_AUDIO_CHANNELstring
	VERIFIER_HEADER_ABORT_VERIFICATIONstring
	VERIFIER_HEADER_START_INPUT_TIMERSstring

	VERIFIER_HEADER_COUNT
)

/** MRCP verifier completion-cause  */
type MRCPVerifierCompletionCause = int

const (
	VERIFIER_COMPLETION_CAUSE_SUCCESS MRCPVerifierCompletionCause = iota
	VERIFIER_COMPLETION_CAUSE_ERROR
	VERIFIER_COMPLETION_CAUSE_NO_INPUT_TIMEOUT
	VERIFIER_COMPLETION_CAUSE_TOO_MUCH_SPEECH_TIMEOUT
	VERIFIER_COMPLETION_CAUSE_SPEECH_TOO_EARLY
	VERIFIER_COMPLETION_CAUSE_BUFFER_EMPTY
	VERIFIER_COMPLETION_CAUSE_OUT_OF_SEQUENCE
	VERIFIER_COMPLETION_CAUSE_REPOSITORY_URI_FAILURE
	VERIFIER_COMPLETION_CAUSE_REPOSITORY_URI_MISSING
	VERIFIER_COMPLETION_CAUSE_VOICEPRINT_ID_MISSING
	VERIFIER_COMPLETION_CAUSE_VOICEPRINT_ID_NOT_EXIST
	VERIFIER_COMPLETION_CAUSE_SPEECH_NOT_USABLE

	VERIFIER_COMPLETION_CAUSE_COUNT
	VERIFIER_COMPLETION_CAUSE_UNKNOWN = VERIFIER_COMPLETION_CAUSE_COUNT
)

/** MRCP verifier methods */
type MRCPVerifierMethodId = int

const (
	VERIFIER_SET_PARAMS MRCPVerifierMethodId = iota
	VERIFIER_GET_PARAMS
	VERIFIER_START_SESSION
	VERIFIER_END_SESSION
	VERIFIER_QUERY_VOICEPRINT
	VERIFIER_DELETE_VOICEPRINT
	VERIFIER_VERIFY
	VERIFIER_VERIFY_FROM_BUFFER
	VERIFIER_VERIFY_ROLLBACK
	VERIFIER_STOP
	VERIFIER_CLEAR_BUFFER
	VERIFIER_START_INPUT_TIMERS
	VERIFIER_GET_INTERMIDIATE_RESULT

	VERIFIER_METHOD_COUNT
)

/** MRCP verifier events */
type MRCPVerifierEventId = int

const (
	VERIFIER_START_OF_INPUT MRCPVerifierEventId = iota
	VERIFIER_VERIFICATION_COMPLETE

	VERIFIER_EVENT_COUNT
)

/** MRCP verifier-header */
type MRCPVerifierHeader struct {

	/** Specifies the voiceprint repository to be used or referenced during
	  speaker verification or identification operations */
	RepositoryUri string
	/** Specifies the claimed identity for verification applications */
	VoiceprintIdentifier string
	/** Specifies the mode of the verification resource */
	VerificationMode string
	/** Indicates the desired behavior of the verification resource
	  after a successful verification operation */
	AdaptModel bool
	/** Indicates the desired behavior of the verification resource
	  upon session termination */
	AbortModel bool
	/** Determines the minimum verification score for which a verification
	  decision of "accepted" may be declared by the server */
	MinVerificationScore float64
	/** Specifies the minimum number of valid utterances
	  before a positive decision is given for verification */
	NumMinVerificationPhrases int64
	/** Specifies the number of valid utterances required
	  before a decision is forced for verification */
	NumMaxVerificationPhrases int64
	/** Sets the length of time from the start of the verification timers
	  (see START-INPUT-TIMERS) until the declaration of a no-input event
	  in the VERIFICATION-COMPLETE server event message */
	NoInputTimeout int64
	/** Allows the client to request the verification resource to save
	  the audio stream that was used for verification/identification */
	SaveWaveform bool
	/** Tells the server resource the Media Type of the captured audio or video
	  such as the one captured and returned by the Waveform-URI header field */
	MediaType string
	/** If the Save-Waveform header field is set to true, the verification resource
	  MUST attempt to record the incoming audio stream of the verification into
	  a file and provide a URI for the client to access it */
	WaveformUri string
	/** Shows the status of the voiceprint specified
	  in the QUERY-VOICEPRINT method */
	VoiceprintExists bool
	/** Indicates that this utterance could be
	  later considered for Speaker Verification */
	VerBufferUtterance bool
	/** Specifies stored audio content that the client requests the server
	  to fetch and process according to the current verification mode,
	  either to train the voiceprint or verify a claimed identity */
	InputWaveformUri string
	/** Indicates the cause of VERIFY or VERIFY-FROM-BUFFER method completion */
	CompletionCause MRCPVerifierCompletionCause
	/** MAY be specified in a VERIFICATION-COMPLETE event
	  coming from the verifier resource to the client */
	CompletionReason string
	/** Specifies the length of silence required following user
	  speech before the speech verifier finalizes a result */
	SpeechCompleteTimeout int64
	/** MAY be specified in a VERIFIER request and allows the
	  client to tell the server that, from this point on, further input
	  audio comes from a different audio source */
	NewAudioChannel bool
	/** MUST be sent in a STOP request to indicate
	  whether or not to abort a VERIFY method in progress */
	AbortVerification bool
	/** MAY be sent as part of a VERIFY request. A value of false
	  tells the verification resource to start the VERIFY operation,
	  but not to start the no-input timer yet */
	StartInputTimers bool
}

/** Get verifier header vtable */
func MRCPVerifierHeaderVTableGet(v mrcp.Version) *header.MRCPHeaderVTable {
	return nil
}

/** Get verifier completion cause string */
func MRCPVerifierCompletionCauseGet(cause MRCPVerifierCompletionCause, v mrcp.Version) string {
	return ""
}

/** Create MRCP verifier resource */
func MRCPVerifierResourceCreate() *resource.MRCPResource {
	return nil
}
