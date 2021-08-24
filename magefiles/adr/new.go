package adr

import (
	"github.com/ClearcodeHQ/magex/tools/adr"
)

// New creates a new ADR entry. Specify title of an ADR entry.
// By default adrDir, adrConfigDir, and adrTemplate are set.
// If you wish to override them use github.com/ClearcodeHQ/magex/tools/adr.RunADRCommand directly
func New(title string) error {
	return adr.RunADRCommand("new", title, "", "", "")
}
