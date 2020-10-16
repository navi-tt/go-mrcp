package control

import (
	apr_toolkit "github.com/navi-tt/go-mrcp/apr-toolkit"
	"github.com/navi-tt/go-mrcp/mrcp"
	"github.com/navi-tt/go-mrcp/mrcp/message/header"
)

type MRCPResource struct {
	Id   mrcp.MRCPResourceId // MRCP resource identifier
	Name string              // MRCP resource name

	/** Get string table of methods */
	GetMethodStrTable func(version mrcp.Version) apr_toolkit.AptStrTableItem
	MethodCount       int64 // Number of methods

	/** Get string table of events */
	GetEventStrTable func(version mrcp.Version) apr_toolkit.AptStrTableItem
	EventCount       int64 // Number of events

	/** Get vtable of resource header */
	GetResourceHeaderVTable func(version mrcp.Version) header.MRCPHeaderVTable
}
