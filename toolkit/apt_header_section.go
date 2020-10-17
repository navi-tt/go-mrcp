package toolkit

import (
	"container/list"
)

/**
 * Header section
 * @remark The header section is a collection of header fields.
 * The header fields are stored in both a ring and an array.
 * The goal is to ensure efficient access and manipulation on the header fields.
 */
type AptHeaderSection struct {
	//todo(是否换成container.List)
	Ring *list.List      // List of header fields (name-value pairs), Ring 的 Value 就是 *AptHeaderField
	Arr  []*list.Element // Array of pointers to header fields
}
