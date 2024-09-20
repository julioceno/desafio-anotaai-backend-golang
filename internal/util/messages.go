package util

const (
	BODY_INVALID    = "The request body is invalid."
	REQUEST_INVALID = "The request parameters are invalid."
	NOT_FOUND       = "The requested resource was not found."
	INTERNAL_ERROR  = "An internal error occurred."
)

type _Messages struct {
	BODY_INVALID string
}

var (
	Messages *_Messages
)

func init() {
	Messages = &_Messages{
		BODY_INVALID: BODY_INVALID,
	}
}
