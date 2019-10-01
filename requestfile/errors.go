package requestfile

import (
	"fmt"
)

type ErrUnableToParseRequest string

func (e ErrUnableToParseRequest) Error() string {
	return fmt.Sprintf("unable to parse request: %s", e)
}
