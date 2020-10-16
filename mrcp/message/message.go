package message

import (
	"github.com/navi-tt/go-mrcp/apr/memory"
	"github.com/navi-tt/go-mrcp/mrcp/control/resource"
	"github.com/navi-tt/go-mrcp/mrcp/message/header"
)

/** MRCP message */
type MRCPMessage struct {
	StartLine *MRCPStartLine           // Start-line of MRCP message
	ChannelId header.MRCPChannelId     // Channel-identifier of MRCP message
	Header    header.MRCPMessageHeader // Header of MRCP message
	Body      string                   // Body of MRCP message

	Resource resource.MRCPResource // Associated MRCP resource
	pool     *memory.AprPool       //  Memory pool to allocate memory from
}
