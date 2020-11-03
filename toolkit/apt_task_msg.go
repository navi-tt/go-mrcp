package toolkit

/** Enumeration of task message types */
type TaskMsgType = int

const (
	TASK_MSG_CORE TaskMsgType = iota /**< core task message type */
	TASK_MSG_USER                    /**< user defined task messages start from here */
)

type CoreTaskMsgType = int

const (
	CORE_TASK_MSG_NONE                 CoreTaskMsgType = iota /**< indefinite message */
	CORE_TASK_MSG_START_COMPLETE                              /**< start-complete message */
	CORE_TASK_MSG_TERMINATE_REQUEST                           /**< terminate-request message */
	CORE_TASK_MSG_TERMINATE_COMPLETE                          /**< terminate-complete message */
	CORE_TASK_MSG_TAKEOFFLINE_REQUEST                         /**< take-offline-request message */
	CORE_TASK_MSG_TAKEOFFLINE_COMPLETE                        /**< take-offline-complete message */
	CORE_TASK_MSG_BRINGONLINE_REQUEST                         /**< bring-online-request message */
	CORE_TASK_MSG_BRINGONLINE_COMPLETE                        /**< bring-online-complete message */
)

/** Task message is used for inter task communication */
type TaskMsg struct {
	/** Message pool the task message is allocated from */
	//apt_task_msg_pool_t *msg_pool;
	/** Task msg type */
	Type int
	/** Task msg sub type */
	SubType int
	/** Context specific data */
	data [1]byte
}
