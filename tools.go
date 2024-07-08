//go:build tools
// +build tools

package tools

/* Manage tool dependencies

See: https://rednafi.com/go/omit_dev_dependencies_in_binaries/

*/

import (
	_ "github.com/go-critic/go-critic/cmd/gocritic@"        // dev dependency
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint" // dev dependency
	_ "golang.org/x/lint/golint"                            // dev dependency
	_ "golang.org/x/tools/cmd/goimports"                    // dev dependency
)
