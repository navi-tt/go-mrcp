package apr

const (
	TABLE_HASH_SIZE = 32
)

/** The opaque string-content table type */
type AprTable struct {
	/* This has to be first to promote backwards compatibility with
	 * older modules which cast a AprTable * to an AprArrayHeader *...
	 * they should use the apr_table_elts() function for most of the
	 * cases they do this for.
	 */

	a       AprArrayHeader // The underlying array for the table
	Creator interface{}    // Who created the array.

	/* An index to speed up table lookups.  The way this works is:
	 *   - Hash the key into the index:
	 *     - index_first[TABLE_HASH(key)] is the offset within
	 *       the table of the first entry with that key
	 *     - index_last[TABLE_HASH(key)] is the offset within
	 *       the table of the last entry with that key
	 *   - If (and only if) there is no entry in the table whose
	 *     key hashes to index element i, then the i'th bit
	 *     of index_initialized will be zero.  (Check this before
	 *     trying to use index_first[i] or index_last[i]!)
	 */
	IndexInitialized uint32
	IndexFirst       [TABLE_HASH_SIZE]int
	IndexLast        [TABLE_HASH_SIZE]int
}
