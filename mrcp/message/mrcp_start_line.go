package message

import (
	"github.com/navi-tt/go-mrcp/apr/memory"
	"github.com/navi-tt/go-mrcp/mrcp"
	"github.com/navi-tt/go-mrcp/toolkit"
)

type MRCPStartLine struct {
	MessageType  MRCPMessageType    // MRCP message type
	Version      mrcp.Version       // Version of protocol in use
	Length       int64              // Specify the length of the message, including the start-line (v2)
	RequestId    mrcp.MRCPRequestId // Unique identifier among client and server
	MethodName   string             // MRCP method name
	MethodId     mrcp.MRCPMethodId  // MRCP method id (associated with method name)
	StatusCode   MRCPStatusCode     // Success or failure or other status of the request
	RequestState MRCPRequestState   // The state of the job initiated by the request
}

/** Initialize MRCP start-line */
func MRCPStartLineInit(statLine *MRCPStartLine) error {
	return nil
}

/** Parse MRCP start-line */
func (statLine *MRCPStartLine) MRCPStartLineParse(str string) (string, error) {
	return "", nil
}

/** Generate MRCP start-line */
func (statLine *MRCPStartLine) MRCPStartLineGenerate(textStream *toolkit.AptTextStream) error {
	return nil
}

/** Finalize MRCP start-line generation */
func (statLine *MRCPStartLine) MRCPStartLineFinalize(contentLength int64, textStream *toolkit.AptTextStream) error {
	return nil
}

/** Parse MRCP request-id */
func MRCPRequestIdParse(field string) mrcp.MRCPRequestId {
	return 0
}

/** Generate MRCP request-id */
func MRCPRequestIdGenerate(rid mrcp.MRCPRequestId, stream *toolkit.AptTextStream) error {
	return nil
}
