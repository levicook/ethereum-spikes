// +build tools

package wally

// These imports ensure build tools are included in Go modules.
// See https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
import (
	_ "github.com/cortesi/modd/cmd/modd"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
