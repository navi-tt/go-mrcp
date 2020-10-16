package engine

import "github.com/navi-tt/go-mrcp/mrcp/message"

/** MRCP state machine */
type MRCPStateMachine struct {
	obj    interface{} // External object associated with state machine
	Active bool        // State either active or deactivating

	/** Virtual update */
	Update func(stateMachine *MRCPStateMachine, message *message.MRCPMessage) bool
	/** Deactivate */
	Deactivate func(stateMachine *MRCPStateMachine) bool

	/** Message dispatcher */
	OnDispatch func(stateMachine *MRCPStateMachine, message *message.MRCPMessage) bool
	/** Deactivated */
	OnDeactivate func(stateMachine *MRCPStateMachine) bool
}
