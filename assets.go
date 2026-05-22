package geremmyas

import "embed"

// EmbeddedFiles contains the pack catalog and all installable templates.
// The CLI copies files from this filesystem into the target repository.
//
//go:embed catalog/** project/** user/**
var EmbeddedFiles embed.FS
