// Package portfolio is the root package. It exists solely to embed
// runtime assets (data, templates, static files) into the binary.
package portfolio

import "embed"

//go:embed data web
var FS embed.FS
