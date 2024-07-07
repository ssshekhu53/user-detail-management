package errors

import (
	"fmt"
	"strings"
)

type MissingParams struct {
	Params []string
}

func (m MissingParams) Error() string {
	var msg string

	params := strings.TrimSuffix(strings.Join(m.Params, ", "), ", ")

	switch len(m.Params) {
	case 0:
		msg = fmt.Sprintf("Missing Params")

		break

	case 1:
		msg = fmt.Sprintf("Missing Param: %s", params)

		break

	default:
		msg = fmt.Sprintf("Missing Params: %s", params)
	}

	return msg
}
