package richerror

const (
	NotFound KindError = iota + 1
	Forbidden
	Unexpected
	Invalid
)

type KindError int

type RichError struct {
	Operation    string
	WrappedError error
	Message      string
	Kind         KindError
	Meta         map[string]interface{}
}

func (r RichError) Error() string {
	return r.Message
}
