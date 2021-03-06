package transport

// ----------------------------------------------------------------------------

// Error is the error type of the GAS package.
//
// It implements the error interface.
type Error int

const (
	// ErrMaxRetries is returned when the called function failed after the
	// maximum number of allowed tries.
	ErrMaxRetries Error = 0x01
	// ErrEmptyLinkRepository is returned when user trying to init connections no init links
	ErrEmptyLinkRepository Error = 0x02
	// ErrWrongNodeMode is an error which appears when i user set wrong node mode
	ErrWrongNodeMode Error = 0x03
)

// ----------------------------------------------------------------------------

// Error returns the error as a string.
func (e Error) Error() string {
	switch e {
	case ErrMaxRetries:
		return "ErrMaxRetries"
	case ErrEmptyLinkRepository:
		return "could not init connections: no datalinks."
	case ErrWrongNodeMode:
		return "Wrong mode in node; will be ignored."
	default:
		return "unknown error"
	}
}

var ErrStubFunction = "Function is empty (Stub realization of interface)."
