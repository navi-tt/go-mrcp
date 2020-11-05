package mpf

/** MPF termination factory */
type TerminationFactory struct {
	/** Virtual create */
	CreateTermination func(factory *TerminationFactory, obj interface{}) *Termination
	/** Virtual assign engine */
	AssignEngine func(factory *TerminationFactory, mediaEngine *Engine) error
}
