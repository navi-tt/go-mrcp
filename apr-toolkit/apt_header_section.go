package apr_toolkit

import "container/ring"

/**
 * Header section
 * @remark The header section is a collection of header fields.
 * The header fields are stored in both a ring and an array.
 * The goal is to ensure efficient access and manipulation on the header fields.
 */
type AptHeaderSection struct {
	Ring ring.Ring         // List of header fields (name-value pairs), Ring 的 Value 就是 *AptHeaderField
	Arr  []*AptHeaderField // Array of pointers to header fields
}
