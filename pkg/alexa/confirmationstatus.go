package alexa

import "strings"

// ConfirmationStatus represents confirmationStatus in JSON
type ConfirmationStatus int

const (
	// ConfirmationStatusUndefined means incoming value incorrect or not supported
	ConfirmationStatusUndefined ConfirmationStatus = iota
	// ConfirmationStatusNone is constant `NONE`
	ConfirmationStatusNone
	// ConfirmationStatusConfirmed is constant `CONFIRMED`
	ConfirmationStatusConfirmed
	// ConfirmationStatusDenied is constant `DENIED`
	ConfirmationStatusDenied
)

// PropTypes for use outside this module
var confirmationStatusStrings = [...]string{
	"Undefined", // Placeholder - should never be this
	"NONE",
	"CONFIRMED",
	"DENIED",
}

// MarshalJSON Function to handle JSON parsing out
func (s ConfirmationStatus) MarshalJSON() ([]byte, error) {
	j := string(`"` + confirmationStatusStrings[s] + `"`)
	return []byte(j), nil
}

// UnmarshalJSON Function to handle JSON parsing out
func (s *ConfirmationStatus) UnmarshalJSON(data []byte) error {
	cs := ConfirmationStatusUndefined
	// Convert to string whilst removing quotes
	d := string(data)[1 : len(data)-1]
	// Find the type in the range of values
	for i, s := range confirmationStatusStrings {
		if strings.ToLower(s) == strings.ToLower(d) {
			cs = ConfirmationStatus(i)
		}
	}
	*s = cs
	return nil
}
