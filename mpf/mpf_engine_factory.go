package mpf

import "github.com/navi-tt/go-mrcp/apr"

/** Factory of media engines */
type EngineFactory struct {
	/** Array of pointers to media engines */
	engines apr.ArrayHeader
	/** Index of the current engine */
	index int
}

/** Create factory of media engines. */
func EngineFactoryCreate() *EngineFactory {
	return nil
}

/** Add media engine to factory. */
func (f *EngineFactory) EngineFactoryEngineAdd(mediaEngine *Engine) error {
	return nil
}

/** Determine whether factory is empty. */
func (f *EngineFactory) EngineFactoryIsEmpty() bool {
	return f.engines.Stack.IsEmpty()
}

/** Select next available media engine. */
func (f *EngineFactory) EngineFactoryEngineSelect() *Engine {
	return nil
}

/** Associate media engines with RTP termination factory. */
func (f *EngineFactory) EngineFactoryRtpFactoryAssign(rtpFactory *TerminationFactory) error {
	return nil
}
