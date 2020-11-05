package mrcp

type Version = int

const (
	MRCP_VERSION_UNKNOWN Version = iota
	MRCP_VERSION_1
	MRCP_VERSION_2
)

/** Enumeration of MRCP resource types */
type MRCPResourceType = int64

const (
	MRCP_SYNTHESIZER_RESOURCE MRCPResourceType = iota /**< Synthesizer resource */
	MRCP_RECOGNIZER_RESOURCE                          /**< Recognizer resource */
	MRCP_RECORDER_RESOURCE                            /**< Recorder resource */
	MRCP_VERIFIER_RESOURCE                            /**< Verifier resource */

	MRCP_RESOURCE_TYPE_COUNT /**< Number of resources */
)

/** MRCP request identifier */
type MRCPRequestId = uint32

/** Method identifier associated with method name */
type MRCPMethodId = int64

/** Resource identifier associated with resource name */
type MRCPResourceId = int64
