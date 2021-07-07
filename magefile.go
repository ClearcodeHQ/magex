//+build mage

package main

import (
	// mage:import go
	_ "github.com/ClearcodeHQ/magex/magefiles/golang"
	// mage:import go:check
	_ "github.com/ClearcodeHQ/magex/magefiles/golang/check"
	// mage:import adr
	_ "github.com/ClearcodeHQ/magex/magefiles/adr"
)
