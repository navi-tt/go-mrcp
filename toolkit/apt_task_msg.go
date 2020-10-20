package toolkit

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
