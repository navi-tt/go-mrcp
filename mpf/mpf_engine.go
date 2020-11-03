package mpf

import (
	"github.com/eapache/queue"
	"github.com/navi-tt/go-mrcp/toolkit"
	"sync"
)

type Engine struct {
	Task              *toolkit.TaskMsg
	TaskMsgType       toolkit.TaskMsgType
	requestQueueGuard sync.Mutex
	RequestQueue      *queue.Queue
	contextFactory    *ContextFactory
	scheduler         *Scheduler
	timerQueue        *toolkit.TimerQueue
	CodecManager      *CodecManager
}

/**
* Create MPF engine.
* @param id the identifier of the engine
* @param pool the pool to allocate memory from
 */
func EngineCreate(id string) *Engine {
	return nil
}

/**
* Create MPF codec manager.
* @param pool the pool to allocate memory from
 */
func EngineCodecManagerCreate() *CodecManager {
	return nil
}

/**
* Register MPF codec manager.
* @param engine the engine to register codec manager for
* @param codec_manager the codec manager to register
 */
func (engine *Engine) EngineCodecManagerRegister(codecManager *CodecManager) error {
	return nil
}

/**
* Create MPF context.
* @param engine the engine to create context for
* @param name the informative name of the context
* @param obj the external object associated with context
* @param max_termination_count the max number of terminations in context
* @param pool the pool to allocate memory from
 */
func (engine *Engine) EngineContextCreate(name string, obj interface{}, maxTerminationCount uint32) *Context {
	return nil
}

/**
* Destroy MPF context.
* @param context the context to destroy
 */
func EngineContextDestroy(context *Context) error {
	return nil
}

/**
* Get external object associated with MPF context.
* @param context the context to get object from
 */
func EngineContextObjectGet(context *Context) error {
	return nil
}

/**
* Get task.
* @param engine the engine to get task from
 */
func (engine *Engine) TaskGet() *toolkit.Task {
	return nil
}

/**
* Set task msg type to send responses and events with.
* @param engine the engine to set task msg type for
* @param type the type to set
 */
func (engine *Engine) EngineTaskMsgTypeSet(taskMsgType toolkit.TaskMsgType) {}

/**
* Create task message(if not created) and add MPF termination message to it.
* @param engine the engine task message belongs to
* @param command_id the MPF command identifier
* @param context the context to add termination to
* @param termination the termination to add
* @param descriptor the termination dependent descriptor
* @param task_msg the task message to create and add constructed MPF message to
 */
func (engine *Engine) EngineTerminationMessageAdd(commandId CommandType,
	context *Context, termination *Termination, descriptor *interface{}, taskMsg *toolkit.TaskMsg) error {
	return nil
}

/**
* Create task message(if not created) and add MPF association message to it.
* @param engine the engine task message belongs to
* @param command_id the MPF command identifier
* @param context the context to add association of terminations for
* @param termination the termination to associate
* @param assoc_termination the termination to associate
* @param task_msg the task message to create and add constructed MPF message to
 */
func (engine *Engine) EngineAssocMessageAdd(commandId CommandType, context *Context, termination *Termination, assocTermination *Termination, taskMsg *toolkit.TaskMsg) error {
	return nil
}

/**
* Create task message(if not created) and add MPF topology message to it.
* @param engine the engine task message belongs to
* @param command_id the MPF command identifier
* @param context the context to modify topology for
* @param task_msg the task message to create and add constructed MPF message to
 */
func (engine *Engine) EngineTopologyMessageAdd(commandId CommandType, context *Context, taskMsg *toolkit.TaskMsg) error {
	return nil
}

/**
* Send MPF task message.
* @param engine the engine to send task message to
* @param task_msg the task message to send
 */
func (engine *Engine) EngineMessageSend(taskMsg *toolkit.TaskMsg) error {
	return nil
}

/**
* Set scheduler rate.
* @param engine the engine to set rate for
* @param rate the rate (n times faster than real-time)
 */
func (engine *Engine) EngineSchedulerRateSet(rate uint64) {

}

/**
* Get the identifier of the engine .
* @param engine the engine to get name of
 */
func (engine *Engine) EngineIdGet() string {
	return ""
}
