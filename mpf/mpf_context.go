package mpf

import "container/list"

/** Factory of media contexts */
type ContextFactory struct {

	/** Ring head */
	Head *list.List // List of header fields (name-value pairs), Ring 的 Value 就是 *AptHeaderField head;
}
