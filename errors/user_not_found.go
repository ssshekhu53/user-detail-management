package errors

import "fmt"

type UserNotFound struct {
	ID int64
}

func (u UserNotFound) Error() string {
	return fmt.Sprintf("user with ID %v not found", u.ID)
}
