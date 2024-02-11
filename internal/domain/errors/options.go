package errors

type Option func(*Error)

func Validation(e *Error) {
	e.Kind = KindValidation
}
