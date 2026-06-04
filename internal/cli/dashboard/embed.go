package dashboard

import "embed"

// Assets holds dashboard HTML templates and static files.
//
//go:embed all:dashboard_assets
var Assets embed.FS
