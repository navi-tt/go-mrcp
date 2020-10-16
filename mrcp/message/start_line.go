package message

import "github.com/navi-tt/go-mrcp/mrcp"

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
