#!/usr/bin/env bash
set -euo pipefail

# geremmyas binary installer
# https://github.com/woliveiras/geremmyas
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/woliveiras/geremmyas/main/install.sh | bash
#   ./install.sh install
#   ./install.sh update
#   ./install.sh uninstall

RELEASE_BASE_URL="https://github.com/woliveiras/geremmyas"
BINARY_NAME="geremmyas"
INSTALL_SOURCE="${GEREMMYAS_INSTALL_SOURCE:-auto}"

info()  { printf "\033[1;34m[info]\033[0m  %s\n" "$1"; }
ok()    { printf "\033[1;32m[ok]\033[0m    %s\n" "$1"; }
warn()  { printf "\033[1;33m[warn]\033[0m  %s\n" "$1"; }
error() { printf "\033[1;31m[error]\033[0m %s\n" "$1" >&2; }

detect_bin_dir() {
  if [[ -n "${XDG_BIN_HOME:-}" ]]; then
    echo "$XDG_BIN_HOME"
  else
    echo "$HOME/.local/bin"
  fi
}

detect_goos() {
  case "$(uname -s)" in
    Darwin) echo "darwin" ;;
    Linux)  echo "linux" ;;
    *)      echo "" ;;
  esac
}

detect_goarch() {
  case "$(uname -m)" in
    x86_64|amd64) echo "amd64" ;;
    arm64|aarch64) echo "arm64" ;;
    *) echo "" ;;
  esac
}

script_dir() {
  cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd
}

ensure_path_hint() {
  local bin_dir="$1"
  case ":$PATH:" in
    *":$bin_dir:"*) ;;
    *)
      warn "$bin_dir is not in PATH."
      warn "Add it to your shell profile to run $BINARY_NAME from anywhere."
      ;;
  esac
}

install_from_release() {
  local bin_path="$1" goos="$2" goarch="$3"
  local url tmp

  if [[ -z "$goos" ]] || [[ -z "$goarch" ]] || ! command -v curl &>/dev/null; then
    return 1
  fi

  url="$RELEASE_BASE_URL/releases/latest/download/${BINARY_NAME}-${goos}-${goarch}"
  tmp="${bin_path}.tmp"

  info "Downloading $BINARY_NAME from latest release..."
  if curl -fsSL "$url" -o "$tmp"; then
    mv "$tmp" "$bin_path"
    chmod +x "$bin_path"
    return 0
  fi

  rm -f "$tmp"
  return 1
}

install_from_checkout() {
  local bin_path="$1"
  local root cache_dir

  root="$(script_dir)"
  cache_dir="${XDG_CACHE_HOME:-$HOME/.cache}/geremmyas"

  if [[ ! -d "$root/cmd/geremmyas" ]]; then
    return 1
  fi
  if ! command -v go &>/dev/null; then
    return 1
  fi

  info "Building $BINARY_NAME from local checkout..."
  mkdir -p "$cache_dir/go-build" "$cache_dir/gomod"
  GOCACHE="$cache_dir/go-build" GOMODCACHE="$cache_dir/gomod" go build -o "$bin_path" "$root/cmd/geremmyas"
}

install_binary() {
  local bin_dir bin_path goos goarch action="${1:-install}"

  bin_dir="$(detect_bin_dir)"
  bin_path="$bin_dir/$BINARY_NAME"
  goos="$(detect_goos)"
  goarch="$(detect_goarch)"

  mkdir -p "$bin_dir"

  if [[ "$INSTALL_SOURCE" != "checkout" ]] && install_from_release "$bin_path" "$goos" "$goarch"; then
    ok "Installed: $bin_path"
    ensure_path_hint "$bin_dir"
    return
  fi

  if [[ "$INSTALL_SOURCE" == "release" ]]; then
    error "Could not install from release asset."
    exit 1
  fi

  if [[ "$INSTALL_SOURCE" == "checkout" ]]; then
    info "Using local checkout source."
  else
    warn "Could not install from release asset. Falling back to local build."
  fi
  if install_from_checkout "$bin_path"; then
    ok "Installed: $bin_path"
    ensure_path_hint "$bin_dir"
    return
  fi

  error "Could not $action $BINARY_NAME."
  error "No release binary was available, or this is not a local checkout with Go installed."
  exit 1
}

uninstall_binary() {
  local bin_path
  bin_path="$(detect_bin_dir)/$BINARY_NAME"

  if [[ -f "$bin_path" ]]; then
    rm "$bin_path"
    ok "Removed: $bin_path"
  else
    warn "$BINARY_NAME is not installed at $bin_path"
  fi
}

print_help() {
  cat <<HELP
geremmyas installer

Usage:
  install.sh              Install or update the geremmyas binary
  install.sh install      Install the geremmyas binary
  install.sh update       Update the geremmyas binary
  install.sh uninstall    Remove the geremmyas binary
  install.sh --help       Show this help message

Environment:
  XDG_BIN_HOME              Override binary install directory
  GEREMMYAS_INSTALL_SOURCE  auto, release, or checkout

After install:
  geremmyas init
  geremmyas sync
HELP
}

command="${1:-install}"

case "$command" in
  install|update)
    install_binary "$command"
    ;;
  uninstall|remove)
    uninstall_binary
    ;;
  --help|-h|help)
    print_help
    ;;
  *)
    error "Unknown command: $command"
    print_help
    exit 1
    ;;
esac
