//go:build tools
// +build tools

package tools

// Manage tool dependencies via go.mod.
import (
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/goreleaser/goreleaser/v2"
	_ "golang.org/x/vuln/cmd/govulncheck"
)
