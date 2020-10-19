package header

import (
	"github.com/navi-tt/go-mrcp/apr/memory"
	"github.com/navi-tt/go-mrcp/toolkit"
)

/** MRCP channel-identifier */
type MRCPChannelId struct {
	SessionId    string // Unambiguous string identifying the MRCP session
	ResourceName string // MRCP resource name
}

/** MRCP message-header */
type MRCPMessageHeader struct {
	GenericHeaderAccessor  MRCPHeaderAccessor       // MRCP generic-header
	ResourceHeaderAccessor MRCPHeaderAccessor       // MRCP resource specific header
	HeaderSection          toolkit.AptHeaderSection // Header section (collection of header fields)
}

/** Initialize MRCP message-header */
func (header *MRCPMessageHeader) MRCPMessageHeaderInit() {
	MRCPHeaderAccessorInit(&header.GenericHeaderAccessor)
	MRCPHeaderAccessorInit(&header.ResourceHeaderAccessor)
	// todo(AptHeaderSectionInit ring 初始化得再确认下)
	toolkit.AptHeaderSectionInit(&header.HeaderSection)
}

/** Allocate MRCP message-header data */
func (header *MRCPMessageHeader) MRCPMessageHeaderDataAlloc(generic, resource *MRCPHeaderVTable) error {
	return nil
}

/** Create MRCP message-header */
func MRCPMessageHeaderCreate(generic, resource *MRCPHeaderVTable) *MRCPMessageHeader {
	return nil
}

/** Destroy MRCP message-header */
func (header *MRCPMessageHeader) MRCPMessageHeaderDestroy() {
	MRCPHeaderDestroy(&header.GenericHeaderAccessor)
	MRCPHeaderDestroy(&header.ResourceHeaderAccessor)
}

/** Add MRCP header field */
func (header *MRCPMessageHeader) MRCPHeaderFieldAdd(field *toolkit.AptHeaderField) error {
	return nil
}

/** Set (copy) MRCP header fields */
func (header *MRCPMessageHeader) MRCPHeaderFieldsSet(srcHeader *MRCPMessageHeader) error {
	return nil
}

/** Get (copy) MRCP header fields */
func (header *MRCPMessageHeader) MRCPHeaderFieldsGet(srcHeader, maskHeader *MRCPMessageHeader) error {
	return nil
}

/** Inherit (copy) MRCP header fields */
func (header *MRCPMessageHeader) MRCPHeaderFieldsInherit(srcHeader *MRCPMessageHeader) error {
	return nil
}

/** Parse MRCP header fields */
func (header *MRCPMessageHeader) MRCPHeaderFieldsParse() error {
	return nil
}

/** Initialize MRCP channel-identifier */
func MRCPChannelIdInit(channelId *MRCPChannelId) {

}

/** Parse MRCP channel-identifier */
func (cid *MRCPChannelId) MRCPChannelIdParse(header *MRCPMessageHeader) error {
	return nil
}

/** Generate MRCP channel-identifier */
func (cid *MRCPChannelId) MRCPChannelIdGenerate(textStream *toolkit.AptTextStream) error {
	return nil
}
