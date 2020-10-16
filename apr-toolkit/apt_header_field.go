package apr_toolkit

import "container/ring"

type AptHeaderField struct {
	Link  *ring.Ring // Ring entry
	Name  string     // Name of the header field
	Value string     // Value of the header field
	Id    int64      // Numeric identifier associated with name
}
