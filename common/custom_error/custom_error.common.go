package custom_error

type ErrObjectNotFound struct{}

func (e *ErrObjectNotFound) Error() string {
	return "Object not found"
}