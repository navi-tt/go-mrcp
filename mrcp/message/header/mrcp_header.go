package header

import (
	apr_toolkit "github.com/navi-tt/go-mrcp/apr-toolkit"
	"github.com/navi-tt/go-mrcp/apr/memory"
)

/** MRCP channel-identifier */
type MRCPChannelId struct {
	SessionId    string // Unambiguous string identifying the MRCP session
	ResourceName string // MRCP resource name
}

/** MRCP message-header */
type MRCPMessageHeader struct {
	GenericHeaderAccessor  MRCPHeaderAccessor           // MRCP generic-header
	ResourceHeaderAccessor MRCPHeaderAccessor           // MRCP resource specific header
	HeaderSection          apr_toolkit.AptHeaderSection // Header section (collection of header fields)
}

/** Initialize MRCP message-header */
func (header *MRCPMessageHeader) MRCPMessageHeaderInit() {
	MRCPHeaderAccessorInit(&header.GenericHeaderAccessor)
	MRCPHeaderAccessorInit(&header.ResourceHeaderAccessor)
	// todo(AptHeaderSectionInit ring 初始化得再确认下)
	apr_toolkit.AptHeaderSectionInit(&header.HeaderSection)
}

/** Allocate MRCP message-header data */
func (header *MRCPMessageHeader) MRCPMessageHeaderDataAlloc(generic, resource *MRCPHeaderVTable, pool *memory.AprPool) error {
	return nil
}

/** Create MRCP message-header */
func MRCPMessageHeaderCreate(generic, resource *MRCPHeaderVTable, pool *memory.AprPool) *MRCPMessageHeader {
	return nil
}

/** Destroy MRCP message-header */
func (header *MRCPMessageHeader) MRCPMessageHeaderDestroy() {
	MRCPHeaderDestroy(&header.GenericHeaderAccessor)
	MRCPHeaderDestroy(&header.ResourceHeaderAccessor)
}

/** Add MRCP header field */
func (header *MRCPMessageHeader) MRCPHeaderFieldAdd(field *apr_toolkit.AptHeaderField, pool *memory.AprPool) error {
	return nil
}

/** Set (copy) MRCP header fields */
func (header *MRCPMessageHeader) MRCPHeaderFieldsSet(srcHeader *MRCPMessageHeader, pool *memory.AprPool) error {
	return nil
}

/** Get (copy) MRCP header fields */
func (header *MRCPMessageHeader) MRCPHeaderFieldsGet(srcHeader, maskHeader *MRCPMessageHeader, pool *memory.AprPool) error {
	return nil
}

/** Inherit (copy) MRCP header fields */
func (header *MRCPMessageHeader) MRCPHeaderFieldsInherit(srcHeader *MRCPMessageHeader, pool *memory.AprPool) error {
	return nil
}

/** Parse MRCP header fields */
func (header *MRCPMessageHeader) MRCPHeaderFieldsParse(pool *memory.AprPool) error {
	return nil
}

/** Initialize MRCP channel-identifier */
func MRCPChannelIdInit(channelId *MRCPChannelId) {

}

/** Parse MRCP channel-identifier */
func (cid *MRCPChannelId) MRCPChannelIdParse(header *MRCPMessageHeader, pool *memory.AprPool) error {
	return nil
}

/** Generate MRCP channel-identifier */
func (cid *MRCPChannelId) MRCPChannelIdGenerate(textStream *apr_toolkit.AptTextStream) error {
	return nil
}
