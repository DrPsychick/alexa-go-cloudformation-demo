package alexa

import (
	"strings"
)

// DirectiveType represents confirmvarious Directive Types
type DirectiveType int

const (
	// DirectiveTypeUndefined means incoming value incorrect or not supported
	DirectiveTypeUndefined DirectiveType = iota
	// DirectiveTypeDialogDelegate is constant `Dialog.Delegate`
	DirectiveTypeDialogDelegate
	// DirectiveTypeDialogElicitSlot is constant `Dialog.ElicitSlot`
	DirectiveTypeDialogElicitSlot
	// DirectiveTypeDialogConfirmSlot is constant `Dialog.ConfirmSlot`
	DirectiveTypeDialogConfirmSlot
	// DirectiveTypeDialogConfirmIntent is constant `Dialog.ConfirmIntent`
	DirectiveTypeDialogConfirmIntent
)

// directiveTypeStrings for use outside this module
var directiveTypeStrings = [...]string{
	"Undefined", // Placeholder - should never be this
	"Dialog.Delegate",
	"Dialog.ElicitSlot",
	"Dialog.ConfirmSlot",
	"Dialog.ConfirmIntent",
}

// Directives is imformation
type Directives struct {
	Type          DirectiveType  `json:"type,omitempty"`
	SlotToElicit  string         `json:"slotToElicit,omitempty"`
	UpdatedIntent *UpdatedIntent `json:"UpdatedIntent,omitempty"`
	PlayBehavior  string         `json:"playBehavior,omitempty"`
	AudioItem     struct {
		Stream struct {
			Token                string `json:"token,omitempty"`
			URL                  string `json:"url,omitempty"`
			OffsetInMilliseconds int    `json:"offsetInMilliseconds,omitempty"`
		} `json:"stream,omitempty"`
	} `json:"audioItem,omitempty"`
}

// MarshalJSON Function to handle JSON parsing out
func (d DirectiveType) MarshalJSON() ([]byte, error) {
	j := string(`"` + directiveTypeStrings[d] + `"`)
	return []byte(j), nil
}

// UnmarshalJSON Function to handle JSON parsing out
func (d *DirectiveType) UnmarshalJSON(data []byte) error {
	dt := DirectiveTypeUndefined
	// Convert to string whilst removing quotes
	x := string(data)[1 : len(data)-1]
	// Find the type in the range of values
	for i, s := range directiveTypeStrings {
		if strings.ToLower(s) == strings.ToLower(x) {
			dt = DirectiveType(i)
		}
	}
	*d = dt
	return nil
}
