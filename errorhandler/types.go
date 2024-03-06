package errorhandler

type BadRequestError struct {
	Message string
}
type UnauthorizedError struct {
	Message string
}
type NotFoundError struct {
	Message string
}
type InternalServerError struct {
	Message string
}

func (e *BadRequestError) Error() string {
	return e.Message
}
func (e *UnauthorizedError) Error() string {
	return e.Message
}
func (e *NotFoundError) Error()  string {
	return e.Message
}
func (e *InternalServerError) Error() string {
	return e.Message
}
