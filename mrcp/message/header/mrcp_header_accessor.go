package header

import (
	"github.com/navi-tt/go-mrcp/apr/memory"
	"github.com/navi-tt/go-mrcp/toolkit"
)

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

/** MRCP header accessor */
type MRCPHeaderAccessor struct {
	Data   interface{}       // Actual header data allocated by accessor
	VTable *MRCPHeaderVTable // Header accessor interface
}

/** Initialize header vtable */
func MRCPHeaderVTableInit(vtable *MRCPHeaderVTable) {
	vtable.Allocate = nil
	vtable.Destroy = nil
	vtable.ParseField = nil
	vtable.GenerateField = nil
	vtable.DuplicateField = nil
	vtable.FieldTable = nil
}

/** Validate header vtable */
func MRCPHeaderVTableValidate(vtable *MRCPHeaderVTable) bool {
	if vtable.Allocate != nil &&
		vtable.Destroy != nil &&
		vtable.ParseField != nil &&
		vtable.GenerateField != nil &&
		vtable.DuplicateField != nil &&
		vtable.FieldTable != nil {
		return true
	}
	return false
}

/** Initialize header accessor */
func MRCPHeaderAccessorInit(accessor *MRCPHeaderAccessor) {
	accessor.Data = nil
	accessor.VTable = nil
}

/** Allocate header data */
func MRCPHeaderAllocate(accessor *MRCPHeaderAccessor, pool *memory.AprPool) interface{} {
	if accessor.Data != nil {
		return accessor.Data
	}
	if accessor.VTable == nil || accessor.VTable.Allocate == nil {
		return nil
	}
	return accessor.VTable.Allocate(accessor, pool)
}

/** Destroy header data */
func MRCPHeaderDestroy(accessor *MRCPHeaderAccessor) {
	if accessor.VTable == nil || accessor.VTable.Destroy == nil {
		return
	}
	accessor.VTable.Destroy(accessor)
}

/** Parse header field value */
func (a *MRCPHeaderAccessor) MRCPHeaderFieldValueParse(field *toolkit.AptHeaderField, pool *memory.AprPool) error {
	return nil
}

/** Generate header field value */
func (a *MRCPHeaderAccessor) MRCPHeaderFieldValueGenerate(id int64, emptyValue bool, pool *memory.AprPool) *toolkit.AptHeaderField {
	return nil
}

/** Duplicate header field value */
func (a *MRCPHeaderAccessor) MRCPHeaderFieldValueDuplicate(srcAccessor *MRCPHeaderAccessor, id int64, pool *memory.AprPool) (value string, err error) {
	return "", nil
}
