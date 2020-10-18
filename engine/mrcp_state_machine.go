package engine

import "github.com/navi-tt/go-mrcp/mrcp/message"

/** MRCP state machine */
type MRCPStateMachine struct {
	obj    interface{} // External object associated with state machine
	Active bool        // State either active or deactivating

	/** Virtual update */
	Update func(machine *MRCPStateMachine, message *message.MRCPMessage) error
	/** Deactivate */
	Deactivate func(machine *MRCPStateMachine) error

	/** Message dispatcher */
	OnDispatch func(machine *MRCPStateMachine, message *message.MRCPMessage) error
	/** Deactivated */
	OnDeactivate func(machine *MRCPStateMachine) error
}

/** Initialize MRCP state machine */
func MRCPStateMachineInit(sMachine *MRCPStateMachine, obj interface{}) error {
	sMachine.obj = obj
	sMachine.Active = true
	sMachine.Update = nil
	sMachine.Deactivate = nil
	sMachine.OnDispatch = nil
	sMachine.OnDeactivate = nil
	return nil
}

/** Update MRCP state machine */
func (m *MRCPStateMachine) MRCPStateMachineUpdate(message *message.MRCPMessage) error {
	if m.Update != nil {
		return m.Update(m, message)
	}
	return nil
}

/** Deactivate MRCP state machine */
func (m *MRCPStateMachine) MRCPStateMachineDeactivate() error {
	if m.Deactivate != nil {
		return m.Deactivate(m)
	}
	return nil
}
