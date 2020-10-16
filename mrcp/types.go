package mrcp

type Version int

const (
	MRCP_VERSION_UNKNOWN Version = iota
	MRCP_VERSION_1
	MRCP_VERSION_2
)

/** MRCP request identifier */
type MRCPRequestId = uint32

/** Method identifier associated with method name */
type MRCPMethodId = int64

/** Resource identifier associated with resource name */
type MRCPResourceId = int64
