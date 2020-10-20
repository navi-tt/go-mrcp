package mpf

import (
	"container/list"
	"github.com/navi-tt/go-mrcp/apr"
)

/** Item of the association matrix */
type MatrixItem struct {
	On uint8
}

/** Item of the association matrix header */
type HeaderItem struct {
	termination *Termination
	TXCount     uint8
	RXCount     uint8
}

/** Media processing context */
type Context struct {

	/** Ring entry */
	Link *list.List
	/** Back pointer to the context factory */
	Factory *ContextFactory

	/** Informative name of the context used for debugging */
	Name string
	/** External object */
	Obj interface{}

	/** Max number of terminations in the context */
	Capacity int64
	/** Current number of terminations in the context */
	Count int64
	/** Header of the association matrix */
	header *HeaderItem
	/** Association matrix, which represents the topology */
	matrix **MatrixItem

	/** Array of media processing objects constructed while
	  applying topology based on association matrix */
	mpfObjects *apr.ArrayHeader
}

/**
 * Create factory of media contexts.
 */
func ContextFactoryCreate() *ContextFactory {
	return nil
}

/**
 * Destroy factory of media contexts.
 */
func ContextFactoryDestroy(factory *ContextFactory) error {
	return nil
}

/**
 * Process factory of media contexts.
 */
func ContextFactoryProcess(factory *ContextFactory) error {
	return nil
}

/**
 * Create MPF context.
 * @param factory the factory context belongs to
 * @param name the informative name of the context
 * @param obj the external object associated with context
 * @param max_termination_count the max number of terminations in context
 * @param pool the pool to allocate memory from
 */
func (F *ContextFactory) ContextCreate(name string, obj interface{}, maxTerminationCount int64) *Context {
	return nil
}

/**
 * Destroy MPF context.
 * @param context the context to destroy
 */
func ContextDestroy(context *Context) error {
	return nil
}

/**
 * Get external object associated with MPF context.
 * @param context the context to get object from
 */
func (context *Context) ContextObjectGet() interface{} {
	return nil
}

/**
 * Add termination to context.
 * @param context the context to add termination to
 * @param termination the termination to add
 */
func (context *Context) ContextTerminationAdd(termination *Termination) error {
	return nil
}

/**
 * Subtract termination from context.
 * @param context the context to subtract termination from
 * @param termination the termination to subtract
 */
func (context *Context) ContextTerminationSubtract(termination *Termination) error {
	return nil
}

/**
 * Add association between specified terminations.
 * @param context the context to add association in the scope of
 * @param termination1 the first termination to associate
 * @param termination2 the second termination to associate
 */
func (context *Context) ContextAssociationAdd(termination1, termination2 *Termination) error {
	return nil
}

/**
 * Remove association between specified terminations.
 * @param context the context to remove association in the scope of
 * @param termination1 the first termination
 * @param termination2 the second termination
 */
func (context *Context) ContextAssociationRemove(termination1, termination2 *Termination) error {
	return nil
}

/**
 * Reset assigned associations and destroy applied topology.
 * @param context the context to reset associations for
 */
func (context *Context) ContextAssociationsReset() error {
	return nil
}

/**
 * Apply topology.
 * @param context the context to apply topology for
 */
func (context *Context) ContextTopologyApply() error {
	return nil
}

/**
 * Destroy topology.
 * @param context the context to destroy topology for
 */
func (context *Context) ContextTopologyDestroy() error {
	return nil
}

/**
 * Process context.
 * @param context the context to process
 */
func (context *Context) ContextProcess() error {
	return nil
}
