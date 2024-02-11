package errors

type Kind uint

const (
	KindInternal Kind = iota
	KindNotFound
	KindValidation
	KindConflict
	KindForbidden
)

type Error struct {
	Kind    Kind
	Message string
	Data    map[string]any
}

func New(msg string, opts ...Option) Error {
	e := Error{
		Kind:    KindInternal,
		Message: msg,
		Data:    map[string]any{},
	}

	for _, opt := range opts {
		opt(&e)
	}

	return e
}

func (e Error) Error() string {
	return e.Message
}
