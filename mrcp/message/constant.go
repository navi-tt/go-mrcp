package message

/** MRCP message types */
type MRCPMessageType int

const (
	MRCP_MESSAGE_TYPE_UNKNOWN MRCPMessageType = iota
	MRCP_MESSAGE_TYPE_REQUEST
	MRCP_MESSAGE_TYPE_RESPONSE
	MRCP_MESSAGE_TYPE_EVENT
)

/** Status codes */
type MRCPStatusCode int

const (
	MRCP_STATUS_CODE_UNKNOWN MRCPStatusCode = 0
	/* success codes (2xx) */
	MRCP_STATUS_CODE_SUCCESS             MRCPStatusCode = 200
	MRCP_STATUS_CODE_SUCCESS_WITH_IGNORE MRCPStatusCode = 201
	/* failure codes (4xx) */
	MRCP_STATUS_CODE_METHOD_NOT_ALLOWED        MRCPStatusCode = 401
	MRCP_STATUS_CODE_METHOD_NOT_VALID          MRCPStatusCode = 402
	MRCP_STATUS_CODE_UNSUPPORTED_PARAM         MRCPStatusCode = 403
	MRCP_STATUS_CODE_ILLEGAL_PARAM_VALUE       MRCPStatusCode = 404
	MRCP_STATUS_CODE_NOT_FOUND                 MRCPStatusCode = 405
	MRCP_STATUS_CODE_MISSING_PARAM             MRCPStatusCode = 406
	MRCP_STATUS_CODE_METHOD_FAILED             MRCPStatusCode = 407
	MRCP_STATUS_CODE_UNRECOGNIZED_MESSAGE      MRCPStatusCode = 408
	MRCP_STATUS_CODE_UNSUPPORTED_PARAM_VALUE   MRCPStatusCode = 409
	MRCP_STATUS_CODE_OUT_OF_ORDER              MRCPStatusCode = 410
	MRCP_STATUS_CODE_RESOURCE_SPECIFIC_FAILURE MRCPStatusCode = 421
)

/** Request-states used in MRCP response message */
type MRCPRequestState int

const (
	/** The request was processed to completion and there will be no
	  more events from that resource to the client with that request-id */
	MRCP_REQUEST_STATE_COMPLETE MRCPRequestState = iota
	/** Indicate that further event messages will be delivered with that request-id */
	MRCP_REQUEST_STATE_INPROGRESS
	/** The job has been placed on a queue and will be processed in first-in-first-out order */
	MRCP_REQUEST_STATE_PENDING

	/** Number of request states */
	MRCP_REQUEST_STATE_COUNT
	/** Unknown request state */
	MRCP_REQUEST_STATE_UNKNOWN = MRCP_REQUEST_STATE_COUNT
)
