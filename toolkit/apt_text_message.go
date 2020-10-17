package toolkit

/** Stage of text message processing (parsing/generation) */
type AptMessageStage = int

const (
	APT_MESSAGE_STAGE_START_LINE AptMessageStage = iota
	APT_MESSAGE_STAGE_HEADER
	APT_MESSAGE_STAGE_BODY
)

/** Text message parser */
type AptMessageParser struct {
}

/** Text message generator */
type AptMessageGenerator struct {
}
