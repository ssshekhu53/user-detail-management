package errors

import (
	"fmt"
	"strings"
)

type InvalidParams struct {
	Params []string
}

func (i InvalidParams) Error() string {
	var msg string

	params := strings.TrimSuffix(strings.Join(i.Params, ", "), ", ")

	switch len(i.Params) {
	case 0:
		msg = fmt.Sprintf("Invalid Params")

		break

	case 1:
		msg = fmt.Sprintf("Invalid Param: %s", params)

		break

	default:
		msg = fmt.Sprintf("Invalid Params: %s", params)
	}

	return msg
}
