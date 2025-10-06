package errors

import "errors"

// these variables are used to give us access to existing
// functions in the std lib errors package. We can also
// wrap them in custom functionnality as needed if we want,
// or mock them during testing
var (
	As = errors.As
	Is = errors.Is
)
