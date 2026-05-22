#!/usr/bin/env bash
set -euo pipefail

# geremmyas binary uninstaller

BINARY_NAME="geremmyas"

info() { printf "\033[1;34m[info]\033[0m  %s\n" "$1"; }
ok()   { printf "\033[1;32m[ok]\033[0m    %s\n" "$1"; }
warn() { printf "\033[1;33m[warn]\033[0m  %s\n" "$1"; }

detect_bin_dir() {
  if [[ -n "${XDG_BIN_HOME:-}" ]]; then
    echo "$XDG_BIN_HOME"
  else
    echo "$HOME/.local/bin"
  fi
}

info "Uninstalling $BINARY_NAME..."

bin_path="$(detect_bin_dir)/$BINARY_NAME"
if [[ -f "$bin_path" ]]; then
  rm "$bin_path"
  ok "Removed: $bin_path"
else
  warn "$BINARY_NAME is not installed at $bin_path"
fi

ok "$BINARY_NAME uninstalled."
