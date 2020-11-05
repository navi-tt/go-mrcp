package resource

import (
	"github.com/navi-tt/go-mrcp/mrcp"
	"github.com/navi-tt/go-mrcp/mrcp/message/header"
	"github.com/navi-tt/go-mrcp/toolkit"
)

type MRCPResource struct {
	Id   mrcp.MRCPResourceId // MRCP resource identifier
	Name string              // MRCP resource name

	/** Get string table of methods */
	GetMethodStrTable func(version mrcp.Version) toolkit.AptStrTableItem
	MethodCount       int64 // Number of methods

	/** Get string table of events */
	GetEventStrTable func(version mrcp.Version) toolkit.AptStrTableItem
	EventCount       int64 // Number of events

	/** Get vtable of resource header */
	GetResourceHeaderVTable func(version mrcp.Version) header.MRCPHeaderVTable
}

/** Initialize MRCP resource */
func MRCPResourceCreate() *MRCPResource {
	resource := MRCPResource{
		Id:                      0,
		Name:                    "",
		GetMethodStrTable:       nil,
		MethodCount:             0,
		GetEventStrTable:        nil,
		EventCount:              0,
		GetResourceHeaderVTable: nil,
	}
	return &resource
}

/** Validate MRCP resource */
func MRCPResourceValidate(resource *MRCPResource) bool {
	if resource.MethodCount > 0 && resource.EventCount > 0 &&
		resource.GetMethodStrTable != nil && resource.GetEventStrTable != nil &&
		resource.GetResourceHeaderVTable != nil &&
		len(resource.Name) > 0 {
		return true
	}
	return false
}
