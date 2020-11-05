package toolkit

import (
	"container/list"
)

type AptHeaderField struct {
	Link  *list.List // Ring entry
	Name  string     // Name of the header field
	Value string     // Value of the header field
	Id    int64      // Numeric identifier associated with name
}

/** Initialize header section (collection of header fields) */
func AptHeaderSectionInit(header *AptHeaderSection) {
	header.Ring = list.New()
	header.Arr = make([]*list.Element, 0)
}

/**
 * Check whether specified header field is set.
 * @param header the header section to use
 * @param id the identifier associated with the header_field to check
 */
func (header *AptHeaderSection) AptHeaderSectionFieldCheck(id int64) bool {
	if id < int64(len(header.Arr)) {
		return header.Arr[id] != nil
	}
	return false
}

/**
 * Get header field by specified identifier.
 * @param header the header section to use
 * @param id the identifier associated with the header_field
 */
func (header *AptHeaderSection) AptHeaderSectionFieldGet(id int64) *AptHeaderField {
	if id < int64(len(header.Arr)) {
		return header.Arr[id].Value.(*AptHeaderField)
	}
	return nil
}

/** Remove header field from header section */
func (header *AptHeaderSection) AptHeaderSectionFieldRemove(headerField *AptHeaderField) error {
	var (
		e *list.Element
	)
	if headerField.Id < int64(len(header.Arr)) {
		e = header.Arr[headerField.Id]
		header.Arr[headerField.Id] = nil
		//todo( 要确认下切片元素只需要设nil, 还是还需要删除元素???)
		//header.Arr = append(header.Arr[:headerField.Id], header.Arr[headerField.Id+1:]...)
	}
	header.Ring.Remove(e)
	return nil
}
