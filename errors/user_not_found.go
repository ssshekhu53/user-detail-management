package errors

import "fmt"

type UserNotFound struct {
	ID int
}

func (u UserNotFound) Error() string {
	if u.ID == 0 {
		return "user not found"
	}

	return fmt.Sprintf("user with ID %v not found", u.ID)
}
