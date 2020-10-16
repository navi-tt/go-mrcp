package header

import (
	apr_toolkit "github.com/navi-tt/go-mrcp/apr-toolkit"
	"github.com/navi-tt/go-mrcp/apr/memory"
)

/** MRCP channel-identifier */
type MRCPChannelId struct {
	SessionId    string            // Unambiguous string identifying the MRCP session
	ResourceName string            // MRCP resource name
}

/** MRCP message-header */
type MRCPMessageHeader struct {
	GenericHeaderAccessor  MRCPHeaderAccessor           // MRCP generic-header
	ResourceHeaderAccessor MRCPHeaderAccessor           // MRCP resource specific header
	HeaderSection          apr_toolkit.AptHeaderSection // Header section (collection of header fields)
}

/** MRCP header accessor */
type MRCPHeaderAccessor struct {
	Data   []byte            // Actual header data allocated by accessor
	VTable *MRCPHeaderVTable // Header accessor interface
}

/** MRCP header accessor interface */
type MRCPHeaderVTable struct {
	/** Allocate actual header data */
	Allocate func(accessor *MRCPHeaderAccessor, pool *memory.AprPool) interface{}
	/** Destroy header data */
	Destroy func(accessor *MRCPHeaderAccessor)

	/** Parse header field value */
	ParseField func(accessor *MRCPHeaderAccessor, id int64, value string, pool *memory.AprPool) bool
	/** Generate header field value */
	GenerateField func(accessor *MRCPHeaderAccessor, id int64, value string, pool *memory.AprPool) bool
	/** Duplicate header field value */
	DuplicateField func(accessor *MRCPHeaderAccessor, src *MRCPHeaderAccessor, id int64, value string, pool *memory.AprPool) bool

	FieldTable map[string]string // Table of fields
	//FieldCount int64             // Number of fields
}
