package engine

import (
	"container/list"
	"fmt"

	"github.com/navi-tt/go-mrcp/mrcp"
)

/** Engine factory declaration */
type MRCPEngineFactory struct {
	/* C:apr_hash_t -> Go:map+list, key 用engine的Id */
	enginesList *list.List
	enginesMap  map[string]*list.Element
	//pool        *memory.AprPool
}

/** Create engine factory */
func MRCPEngineFactoryCreate() *MRCPEngineFactory {
	return &MRCPEngineFactory{
		enginesList: list.New(),
		enginesMap:  make(map[string]*list.Element),
		//pool:        &memory.AprPool{},
	}
}

/** Destroy registered engines and the factory */
func MRCPEngineFactoryDestroy(factory *MRCPEngineFactory) error {
	if factory == nil {
		return nil
	}
	var (
		engine *MRCPEngine
	)

	e := factory.MRCPEngineFactoryEngineFirst()
	for ; e != nil; e = e.Next() {
		engine = e.Value.(*MRCPEngine)
		if engine != nil {
			// MRCPEngineVirtualDestroy(engine) todo()
			delete(factory.enginesMap, engine.Id)
		}
	}
	factory.enginesList = nil
	factory.enginesMap = nil
	return nil
}

/** Open registered engines */
func (f *MRCPEngineFactory) MRCPEngineFactoryOpen() error {
	return nil
}

/** Close registered engines */
func (f *MRCPEngineFactory) MRCPEngineFactoryClose() error {
	return nil
}

/** Register engine */
func (f *MRCPEngineFactory) MRCPEngineFactoryRegister(engine *MRCPEngine) error {
	if engine == nil {
		return fmt.Errorf("invalid engine")
	}
	if len(engine.Id) == 0 {
		return fmt.Errorf("invalid engine, not id")
	}
	switch engine.ResourceId {
	case mrcp.MRCP_SYNTHESIZER_RESOURCE:
	//engine.CreateStateMachine = mrcp_synth_state_machine_create
	case mrcp.MRCP_RECOGNIZER_RESOURCE:
		//engine.CreateStateMachine = mrcp_recog_state_machine_create
		break
	case mrcp.MRCP_RECORDER_RESOURCE:
		//engine.CreateStateMachine = mrcp_recorder_state_machine_create
		break
	case mrcp.MRCP_VERIFIER_RESOURCE:
		//engine.CreateStateMachine = mrcp_verifier_state_machine_create
		break
	default:
		break
	}

	if engine.CreateStateMachine == nil {
		return fmt.Errorf("invalid engine, CreateStateMachine is nil")
	}

	e := f.enginesList.PushBack(engine)
	f.enginesMap[engine.Id] = e
	return nil
}

/** Get engine by name */
func (f *MRCPEngineFactory) MRCPEngineFactoryEngineGet(name string) *MRCPEngine {
	return f.enginesMap[name].Value.(*MRCPEngine)
}

/** Find engine by resource identifier */
func (f *MRCPEngineFactory) MRCPEngineFactoryEngineFind(rid mrcp.MRCPResourceId) *MRCPEngine {
	e := f.MRCPEngineFactoryEngineFirst()
	for ; e != nil; e = e.Next() {
		engine := e.Value.(*MRCPEngine)
		if engine.ResourceId == rid {
			return engine
		}
	}
	return nil
}

/** Start iterating over the engines in a factory */
func (f *MRCPEngineFactory) MRCPEngineFactoryEngineFirst() *list.Element {
	return f.enginesList.Front()
}
