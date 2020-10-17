package engine

import "github.com/navi-tt/go-mrcp/mrcp/message"

/** MRCP state machine */
type MRCPStateMachine struct {
	obj    interface{} // External object associated with state machine
	Active bool        // State either active or deactivating

	/** Virtual update */
	Update func(machine *MRCPStateMachine, message *message.MRCPMessage) bool
	/** Deactivate */
	Deactivate func(machine *MRCPStateMachine) bool

	/** Message dispatcher */
	OnDispatch func(machine *MRCPStateMachine, message *message.MRCPMessage) bool
	/** Deactivated */
	OnDeactivate func(machine *MRCPStateMachine) bool
}
