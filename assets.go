package geremmyas

import "embed"

// EmbeddedFiles contains the pack catalog and all installable templates.
// The CLI copies files from this filesystem into the target repository.
//
//go:embed catalog/** content/** targets/**
var EmbeddedFiles embed.FS
