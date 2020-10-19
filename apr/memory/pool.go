package memory

//
//import "sync"
//
//type AprPool struct {
//	parent       *AprPool
//	child        *AprPool
//	sibling      *AprPool
//	ref          []*AprPool
//	Cleanups     []Cleanup
//	FreeCleanups []Cleanup
//}
//
//type AprStatus = int
//
//type Cleanup struct {
//	Next           *Cleanup
//	Data           []byte
//	plainCleanupFn func(data interface{}) AprStatus
//	childCleanupFn func(data interface{}) AprStatus
//}
//
//const __MAX_INDEX = 20
//
//type AprAllocator struct {
//	MaxIndex uint32 // largest used index into free[], always < MAX_INDEX
//	/** Total size (in BOUNDARY_SIZE multiples) of unused memory before
//	 * blocks are given back. @see apr_allocator_max_free_set().
//	 * @note Initialized to APR_ALLOCATOR_MAX_FREE_UNLIMITED,
//	 * which means to never give back blocks.
//	 */
//	MaxFreeIndex uint32
//	/**
//	 * Memory size (in BOUNDARY_SIZE multiples) that currently must be freed
//	 * before blocks are given back. Range: 0..max_free_index
//	 */
//	CurrentFreeIndex uint32
//
//	lock sync.Mutex
//
//	Owner *AprPool
//	/**
//	 * Lists of free nodes. Slot 0 is used for oversized nodes,
//	 * and the slots 1..MAX_INDEX-1 contain nodes of sizes
//	 * (i+1) * BOUNDARY_SIZE. Example for BOUNDARY_INDEX == 12:
//	 * slot  0: nodes larger than 81920
//	 * slot  1: size  8192
//	 * slot  2: size 12288
//	 * ...
//	 * slot 19: size 81920
//	 */
//	Free [__MAX_INDEX][]AprMemNode
//}
//
///** basic memory node structure
// * @note The next, ref and first_avail fields are available for use by the
// *       caller of apr_allocator_alloc(), the remaining fields are read-only.
// *       The next field has to be used with caution and sensibly set when the
// *       memnode is passed back to apr_allocator_free().  See apr_allocator_free()
// *       for details.
// *       The ref and first_avail fields will be properly restored by
// *       apr_allocator_free().
// */
//type AprMemNode struct {
//	Next       *AprMemNode   // next AprMemNode
//	Ref        []*AprMemNode // reference to self
//	Index      uint32        // size
//	FreeIndex  uint32        // how much free
//	FirstAvail byte          // pointer to first free memory
//	EndP       byte          // pointer to end of free memory
//}
