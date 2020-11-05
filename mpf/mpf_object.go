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

/** Initialize object */
func ObjectInit(name string) *Object {
	return &Object{
		Name:    name,
		Destroy: nil,
		Process: nil,
		Trace:   nil,
	}
}

/** Destroy object */
func ObjectDestroy(object *Object) error {
	if object.Destroy != nil {
		return object.Destroy(object)
	}
	return nil
}

/** Process object */
func (object *Object) ObjectProcess() error {
	if object.Process != nil {
		return object.Process(object)
	}
	return nil
}

/** Trace media path */
func (object *Object) ObjectTrace() error {
	if object.Trace != nil {
		return object.Trace(object)
	}
	return nil
}
