package errors

type UserAlreadyExists struct{}

func (u UserAlreadyExists) Error() string {
	return "user already exists with given combination"
}
