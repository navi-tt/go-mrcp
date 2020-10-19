package apr

import "container/list"

const (
	TABLE_HASH_SIZE = 32
)

/** The opaque string-content table type */
type AprTable struct {
	/* This has to be first to promote backwards compatibility with
	 * older modules which cast a AprTable * to an ArrayHeader *...
	 * they should use the apr_table_elts() function for most of the
	 * cases they do this for.
	 */

	a       ArrayHeader // The underlying array for the table
	Creator interface{} // Who created the array.

	table map[string]*list.Element
}
