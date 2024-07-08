//go:build tools
// +build tools

package tools

/* Manage tool dependencies

See:
- https://pkg.go.dev/go/build#hdr-Build_Constraints
- https://rednafi.com/go/omit_dev_dependencies_in_binaries/
- https://marcofranssen.nl/manage-go-tools-via-go-modules

*/

import (
	_ "github.com/client9/misspell"                         // dev dependency
	_ "github.com/go-critic/go-critic/cmd/gocritic"         // dev dependency
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint" // dev dependency
	_ "golang.org/x/lint/golint"                            // dev dependency
	_ "golang.org/x/tools/cmd/goimports"                    // dev dependency
)
