package message

import (
	"github.com/navi-tt/go-mrcp/apr"
	"github.com/navi-tt/go-mrcp/mrcp/control"
	"github.com/navi-tt/go-mrcp/mrcp/message/header"
)

/** MRCP message */
type MRCPMessage struct {
	StartLine *MRCPStartLine           // Start-line of MRCP message
	ChannelId header.MRCPChannelId     // Channel-identifier of MRCP message
	Header    header.MRCPMessageHeader // Header of MRCP message
	Body      string                   // Body of MRCP message

	Resource control.MRCPResource // Associated MRCP resource
	pool     *apr.AprPool         //  Memory pool to allocate memory from
}
