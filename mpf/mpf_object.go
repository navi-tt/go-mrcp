package mpf

/** Media processing objects base */
type Object struct {
	/** Informative name used for debugging */
	Name string
	/** Virtual destroy */
	Destroy func(object *Object) error
	/** Virtual process */
	Process func(object *Object) error
	/** Virtual trace of media path */
	Trace func(object *Object) error
}
