package toolkit

import "container/list"

/** Internal states of the task */
type TaskState = int

const (
	TASK_STATE_IDLE                TaskState = iota /**< no task activity */
	TASK_STATE_START_REQUESTED                      /**< start of the task has been requested, but it's not running yet */
	TASK_STATE_RUNNING                              /**< task is running */
	TASK_STATE_TERMINATE_REQUESTED                  /**< termination of the task has been requested, but it's still running */
)

type Task struct {
	link *list.List    /* entry to parent task ring */
	head *list.Element /* head of child tasks ring */

	Name string      /* name of the task */
	Obj  interface{} /* external object associated with the task */
	//apt_task_msg_pool_t *msg_pool;      /* message pool to allocate task messages from */
	//apr_thread_mutex_t  *data_guard;    /* mutex to protect task data */
	//apr_thread_t        *thread_handle; /* thread handle */
	//apt_task_state_e     state;         /* current task state */
	//apt_task_vtable_t    vtable;        /* table of virtual methods */
	//apt_task_t          *parent_task;   /* parent (master) task */
	//apr_size_t           pending_start; /* number of pending start requests */
	//apr_size_t           pending_term;  /* number of pending terminate requests */
	//apr_size_t           pending_off;   /* number of pending taking-offline requests */
	//apr_size_t           pending_on;    /* number of pending bringing-online requests */
	//apt_bool_t           running;       /* task is running (TRUE if even terminate has already been requested) */
	//apt_bool_t           auto_ready;    /* if TRUE, task is implicitly ready to process messages */
}
