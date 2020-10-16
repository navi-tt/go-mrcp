package apr

import "github.com/navi-tt/go-mrcp/apr/memory"

/** An opaque array type */
type AprArrayHeader struct {
	pool        *memory.AprPool // The pool the array is allocated out of
	EltSize     int             // The amount of memory allocated for each element of the array
	NumElt      int             // The number of active elements in the array
	NumEltAlloc int             // The number of elements allocated in the array
	Elements    string          //  The elements in the array
}
