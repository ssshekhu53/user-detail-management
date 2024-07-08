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
		msg = fmt.Sprintf("invalid params")

		break

	case 1:
		msg = fmt.Sprintf("invalid param: %s", params)

		break

	default:
		msg = fmt.Sprintf("invalid params: %s", params)
	}

	return msg
}
