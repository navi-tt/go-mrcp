package resources

import (
	"github.com/navi-tt/go-mrcp/mrcp"
	"github.com/navi-tt/go-mrcp/mrcp/control/resource"
	"github.com/navi-tt/go-mrcp/mrcp/message/header"
)

/**
 * @Author: Liu xiangpeng
 * @Date: 2020/10/17 2:23 下午
 */

type MRCPRecorderHeaderId = int

const (
	RECORDER_HEADER_SENSITIVITY_LEVEL MRCPRecorderHeaderId = iota
	RECORDER_HEADER_NO_INPUT_TIMEOUT
	RECORDER_HEADER_COMPLETION_CAUSE
	RECORDER_HEADER_COMPLETION_REASON
	RECORDER_HEADER_FAILED_URI
	RECORDER_HEADER_FAILED_URI_CAUSE
	RECORDER_HEADER_RECORD_URI
	RECORDER_HEADER_MEDIA_TYPE
	RECORDER_HEADER_MAX_TIME
	RECORDER_HEADER_TRIM_LENGTH
	RECORDER_HEADER_FINAL_SILENCE
	RECORDER_HEADER_CAPTURE_ON_SPEECH
	RECORDER_HEADER_VER_BUFFER_UTTERANCE
	RECORDER_HEADER_START_INPUT_TIMERS
	RECORDER_HEADER_NEW_AUDIO_CHANNEL

	RECORDER_HEADER_COUNT
)

/** MRCP recorder completion-cause  */
type MRCPRecorderCompletionCause = int

const (
	RECORDER_COMPLETION_CAUSE_SUCCESS_SILENCE  MRCPRecorderCompletionCause = 0
	RECORDER_COMPLETION_CAUSE_SUCCESS_MAXTIME  MRCPRecorderCompletionCause = 1
	RECORDER_COMPLETION_CAUSE_NO_INPUT_TIMEOUT MRCPRecorderCompletionCause = 2
	RECORDER_COMPLETION_CAUSE_URI_FAILURE      MRCPRecorderCompletionCause = 3
	RECORDER_COMPLETION_CAUSE_ERROR            MRCPRecorderCompletionCause = 4

	RECORDER_COMPLETION_CAUSE_COUNT   MRCPRecorderCompletionCause = 5
	RECORDER_COMPLETION_CAUSE_UNKNOWN MRCPRecorderCompletionCause = RECORDER_COMPLETION_CAUSE_COUNT
)

/** MRCP recorder methods */
type MRCPRecorderMethodId = int

const (
	RECORDER_SET_PARAMS MRCPRecorderMethodId = iota
	RECORDER_GET_PARAMS
	RECORDER_RECORD
	RECORDER_STOP
	RECORDER_START_INPUT_TIMERS

	RECORDER_METHOD_COUNT
)

/** MRCP recorder events */
type MRCPRecorderEventId = int

const (
	RECORDER_START_OF_INPUT MRCPRecorderEventId = iota
	RECORDER_RECORD_COMPLETE

	RECORDER_EVENT_COUNT
)

/** MRCP recorder-header */
type MRCPRecorderHeader struct {
	/** To filter out background noise and not mistake it for speech */
	SensitivityLevel float64
	/** When recording is started and there is no speech detected for a
	certain period of time, the recorder can send a RECORD-COMPLETE event */
	NoInputTimeout int64
	/** MUST be part of a RECORD-COMPLETE, event coming from
	  the recorder resource to the client */
	CompletionCause MRCPRecognizerCompletionCause
	/** MAY be specified in a RECORD-COMPLETE event coming from
	the recorder resource to the client */
	CompletionReason string
	/** When a recorder method needs to post the audio to a URI and access to
	the URI fails, the server MUST provide the failed URI in this header
	in the method response */
	FailedUri string
	/** When a recorder method needs to post the audio to a URI and access to
	the URI fails, the server MUST provide the URI specific or protocol
	specific response code through this header in the method response */
	FailedUriCause string
	/** When a recorder method contains this header the server must capture
	the audio and store it */
	RecordUri string
	/** Tells the server resource the Media Type in which to store captured
	audio such as the one captured and returned by the Waveform-URI header */
	MediaType string
	/** When recording is started this specifies the maximum length of the
	recording in milliseconds, calculated from the time the actual
	capture and store begins and is not necessarily the time the RECORD
	method is received */
	MaxTime int64
	/** This header MAY be sent on a STOP method and specifies the length of
	audio to be trimmed from the end of the recording after the stop */
	TrimLength int64
	/**  When recorder is started and the actual capture begins, this header
	specifies the length of silence in the audio that is to be
	interpreted as the end of the recording*/
	FinalSilence int64
	/** f false, the recorder MUST start capturing immediately when started.
	If true, the recorder MUST wait for the endpointing functionality to
	detect speech before it starts capturing */
	CaptureOnSpeech bool
	/** Tells the server to buffer the utterance associated with this
	recording request into the verification buffer */
	VerBufferUtterance bool
	/** MAY be sent as part of the RECORD request. A value of false tells the
	recorder resource to start the operation, but not to start the no-input
	timer until the client sends a START-INPUT-TIMERS */
	StartInputTimers bool
	/** MAY be specified in a RECORD request and allows the
	client to tell the server that, from this point on, further input
	audio comes from a different audio source */
	NewAudioChannel bool
}

/** Get recorder header vtable */
func MRCPRecorderHeaderVTableGet(v mrcp.Version) *header.MRCPHeaderVTable {
	return nil
}

/** Get recorder completion cause string */
func MRCPRecorderCompletionCauseGet(cause MRCPRecorderCompletionCause, v mrcp.Version) string {
	return ""
}

/** Create MRCP recorder resource */
func MRCPRecorderResourceCreate() *resource.MRCPResource {
	return nil
}
