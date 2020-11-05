package mpf

/** Max number of messages grouped in a container */
const MAX_MPF_MESSAGE_COUNT = 5

/** Enumeration of MPF message types */
type MessageType = int

const (
	MPF_MESSAGE_TYPE_REQUEST  MessageType = iota /**< request message */
	MPF_MESSAGE_TYPE_RESPONSE                    /**< response message */
	MPF_MESSAGE_TYPE_EVENT                       /**< event message */
)

/** Enumeration of MPF status codes */
type StatusCode = int

const (
	MPF_STATUS_CODE_SUCCESS StatusCode = iota /**< indicates success */
	MPF_STATUS_CODE_FAILURE                   /**< indicates failure */
)

/** Enumeration of MPF commands */
type CommandType = int

const (
	MPF_ADD_TERMINATION      CommandType = iota /**< add termination to context */
	MPF_MODIFY_TERMINATION                      /**< modify termination properties */
	MPF_SUBTRACT_TERMINATION                    /**< subtract termination from context */
	MPF_ADD_ASSOCIATION                         /**< add association between terminations */
	MPF_REMOVE_ASSOCIATION                      /**< remove association between terminations */
	MPF_RESET_ASSOCIATIONS                      /**< reset associations among terminations (also destroy topology) */
	MPF_APPLY_TOPOLOGY                          /**< apply topology based on assigned associations */
	MPF_DESTROY_TOPOLOGY                        /**< destroy applied topology */
)

/** MPF message definition */
type Message struct {
	/** Message type (request/response/event) */
	messageType MessageType
	/** Command identifier (add, modify, subtract, ...) */
	commandId CommandType
	/** Status code used in responses */
	statusCode StatusCode

	/** Context */
	context *Context
	/** Termination */
	termination *Termination
	/** Associated termination */
	assocTermination *Termination
	/** Termination type dependent descriptor */
	descriptor interface{}
}

/** MPF message container definition */
type MessageContainer struct {

	/** Number of actual messages */
	Count int32
	/** Array of messages */
	messages [MAX_MPF_MESSAGE_COUNT]Message
}
