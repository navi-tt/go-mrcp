package mpf

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
