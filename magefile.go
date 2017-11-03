// +build mage

package main

import (
	"github.com/magefile/mage/sh"
	"time"
)

var (
	packageName = "github.com/kevingimbel/license"
	LDFLAGS     = "-X $PACKAGE/lib.version=$VERSION -X $PACKAGE/lib.buildDate=$BUILD_DATE"
)

// Proves a flag environment for LDFLAGS
func flagEnv() map[string]string {
	return map[string]string{
		"PACKAGE":    packageName,
		"BUILD_DATE": time.Now().Format("2006-01-02T15:04:05Z0700"),
		"VERSION": "1.0.0",
	}
}

// Run tests
func Test() error {
	return sh.Run("go", "test")
}

// Build the binary
func Build() error {
	return sh.RunWith(flagEnv(), "go", "build", "-ldflags", LDFLAGS, packageName)
}

// Run "go fmt" on all files
func Format() error {
	return sh.Run("go", "fmt")
}