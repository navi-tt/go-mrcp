package message

import (
	"github.com/navi-tt/go-mrcp/apr/memory"
	"github.com/navi-tt/go-mrcp/mrcp"
	"github.com/navi-tt/go-mrcp/mrcp/control/resource"
	"github.com/navi-tt/go-mrcp/mrcp/message/header"
	"github.com/navi-tt/go-mrcp/toolkit"
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

/**
 * Create an MRCP message.
 * @param pool the pool to allocate memory from
 */
func MRCPMessageCreate() *MRCPMessage {
	return &MRCPMessage{}
}

/**
 * Create an MRCP request message.
 * @param resource the MRCP resource to use
 * @param version the MRCP version to use
 * @param method_id the MRCP resource specific method identifier
 * @param pool the pool to allocate memory from
 */
func MRCPRequestCreate(res *resource.MRCPResource, v mrcp.Version, mid mrcp.MRCPMethodId) *MRCPMessage {
	return &MRCPMessage{}
}

/**
 * Create an MRCP response message based on given request message.
 * @param request_message the MRCP request message to create a response for
 * @param pool the pool to allocate memory from
 */
func MRCPResponseCreate(reqMessage *MRCPMessage) *MRCPMessage {
	return &MRCPMessage{}
}

/**
 * Create an MRCP event message based on given requuest message.
 * @param request_message the MRCP request message to create an event for
 * @param event_id the MRCP resource specific event identifier
 * @param pool the pool to allocate memory from
 */
func MRCPEventCreate(reqMessage *MRCPMessage, evevtId mrcp.MRCPMethodId) *MRCPMessage {
	return &MRCPMessage{}
}

/**
 * Associate MRCP resource with message.
 * @param message the message to associate resource with
 * @param resource the resource to associate
 */
func (m *MRCPMessage) MRCPMessageResourceSet(resource *resource.MRCPResource) error {
	return nil
}

/**
 * Validate MRCP message.
 * @param message the message to validate
 */
func (m *MRCPMessage) MRCPMessageValidate() error {
	return nil
}

/**
 * Destroy MRCP message.
 * @param message the message to destroy
 */
func (m *MRCPMessage) MRCPMessageDestroy() error {
	return nil
}

/**
 * Get MRCP generic header.
 * @param message the message to get generic header from
 */
func (m *MRCPMessage) MRCPGenericHeaderGet() *header.MRCPGenericHeader {
	return m.Header.GenericHeaderAccessor.Data.(*header.MRCPGenericHeader)
}

/**
 * Allocate (if not allocated) and get MRCP generic header.
 * @param message the message to prepare generic header for
 */
func (m *MRCPMessage) MRCPGenericHeaderPrepare() *header.MRCPGenericHeader {
	return header.MRCPHeaderAllocate(&m.Header.GenericHeaderAccessor, m.pool).(*header.MRCPGenericHeader)
}

/**
 * Add MRCP generic header field by specified property (numeric identifier).
 * @param message the message to add property for
 * @param id the numeric identifier to add
 */
func (m *MRCPMessage) MRCPGenericHeaderPropertyAdd(id int64) error {
	return nil
}

/**
 * Add only the name of MRCP generic header field specified by property (numeric identifier).
 * @param message the message to add property for
 * @param id the numeric identifier to add
 * @remark Should be used to construct empty header fiedls for GET-PARAMS requests
 */
func (m *MRCPMessage) MRCPGenericHeaderNamePropertyAdd(id int64) error {
	return nil
}

/**
 * Remove MRCP generic header field by specified property (numeric identifier).
 * @param message the message to remove property from
 * @param id the numeric identifier to remove
 */
func (m *MRCPMessage) MRCPGenericHeaderPropertyRemove(id int64) error {
	headerField := m.Header.HeaderSection.AptHeaderSectionFieldGet(id)
	if headerField != nil {
		return m.Header.HeaderSection.AptHeaderSectionFieldRemove(headerField)
	}
	return nil
}

/**
 * Check whether specified by property (numeric identifier) MRCP generic header field is set or not.
 * @param message the message to use
 * @param id the numeric identifier to check
 */
func (m *MRCPMessage) MRCPGenericHeaderPropertyCheck(id int64) bool {
	return m.Header.HeaderSection.AptHeaderSectionFieldCheck(id)
}

/**
 * Get MRCP resource header.
 * @param message the message to get resource header from
 */
func (m *MRCPMessage) MRCPResourceHeaderGet() interface{} {
	return m.Header.ResourceHeaderAccessor.Data
}

/**
 * Allocate (if not allocated) and get MRCP resource header.
 * @param message the message to prepare resource header for
 */
func (m *MRCPMessage) MRCPResourceHeaderPrepare() interface{} {
	return header.MRCPHeaderAllocate(&m.Header.ResourceHeaderAccessor, m.pool)
}

/**
 * Add MRCP resource header field by specified property (numeric identifier).
 * @param message the message to add property for
 * @param id the numeric identifier to add
 */
func (m *MRCPMessage) MRCPResourceHeaderPropertyAdd(id int64) error {
	return nil
}

/**
 * Add only the name of MRCP resource header field specified by property (numeric identifier).
 * @param message the message to add property for
 * @param id the numeric identifier to add
 * @remark Should be used to construct empty header fields for GET-PARAMS requests
 */
func (m *MRCPMessage) MRCPResourceHeaderNamePropertyAdd(id int64) error {
	return nil
}

/**
 * Remove MRCP resource header field by specified property (numeric identifier).
 * @param message the message to remove property from
 * @param id the numeric identifier to remove
 */
func (m *MRCPMessage) MRCPResourceHeaderPropertyRemove(id int64) error {
	headerField := m.Header.HeaderSection.AptHeaderSectionFieldGet(id + int64(header.GENERIC_HEADER_COUNT))
	if headerField != nil {
		return m.Header.HeaderSection.AptHeaderSectionFieldRemove(headerField)
	}
	return nil
}

/**
 * Check whether specified by property (numeric identifier) MRCP resource header field is set or not.
 * @param message the message to use
 * @param id the numeric identifier to check
 */
func (m *MRCPMessage) MRCPResourceHeaderPropertyCheck(id int64) bool {
	return m.Header.HeaderSection.AptHeaderSectionFieldCheck(id + int64(header.GENERIC_HEADER_COUNT))
}

/**
 * Add MRCP header field.
 * @param message the message to add header field for
 * @param header_field the header field to add
 */
func (m *MRCPMessage) MRCPMessageHeaderFieldAdd(headerField *toolkit.AptHeaderField) error {
	return m.Header.MRCPHeaderFieldAdd(headerField, m.pool)
}

/**
 * Get the next MRCP header field.
 * @param message the message to use
 * @param header_field current header field
 * @remark Should be used to iterate on header fields
 *
 *	apt_header_field_t *header_field = NULL;
 *	while( (header_field = mrcp_message_next_header_field_get(message,header_field)) != NULL ) {
 *  }
 */
func (m *MRCPMessage) MRCPMessageNextHeaderFieldGet(headerField *toolkit.AptHeaderField) *toolkit.AptHeaderField {
	return nil
}
