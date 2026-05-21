#!/usr/bin/env bash
set -euo pipefail

# copilot-configs installer
# https://github.com/woliveiras/copilot-configs
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/woliveiras/copilot-configs/main/install.sh | bash
#   ~/.copilot-configs/install.sh --project
#   ~/.copilot-configs/install.sh --project --configure
#   ~/.copilot-configs/install.sh --configure

REPO_URL="https://github.com/woliveiras/copilot-configs.git"
INSTALL_DIR="$HOME/.copilot-configs"
FORCE=false
PROJECT=false
CONFIGURE=false

# ---------------------------------------------------------------------------
# Helpers
# ---------------------------------------------------------------------------

info()  { printf "\033[1;34m[info]\033[0m  %s\n" "$1"; }
ok()    { printf "\033[1;32m[ok]\033[0m    %s\n" "$1"; }
warn()  { printf "\033[1;33m[warn]\033[0m  %s\n" "$1"; }
error() { printf "\033[1;31m[error]\033[0m %s\n" "$1" >&2; }

detect_vscode_dir() {
  case "$(uname -s)" in
    Darwin)
      echo "$HOME/Library/Application Support/Code/User"
      ;;
    Linux)
      echo "$HOME/.config/Code/User"
      ;;
    *)
      error "Unsupported OS: $(uname -s)"
      exit 1
      ;;
  esac
}

# Files users are expected to customize. Preserved during updates unless --force.
CUSTOMIZABLE_FILES="
AGENTS.md
copilot-instructions.md
hooks/guardrails-rules.txt
"

is_customizable() {
  echo "$CUSTOMIZABLE_FILES" | grep -qxF "$1"
}

# Copy a file, always overwriting. Silently skips if files are identical.
# Usage: safe_copy <src> <dest>
safe_copy() {
  local src="$1" dest="$2"
  if [[ -f "$dest" ]] && diff -q "$src" "$dest" &>/dev/null; then
    return 0
  fi
  local action="Installed"
  [[ -f "$dest" ]] && action="Updated"
  mkdir -p "$(dirname "$dest")"
  cp "$src" "$dest"
  ok "$action: $dest"
}

# Copy a customizable file. Existing customized files are preserved unless
# --force is provided.
# Usage: safe_copy_customizable <src> <dest> <relative_path>
safe_copy_customizable() {
  local src="$1" dest="$2" rel_path="$3"
  if is_customizable "$rel_path" && [[ -f "$dest" ]] && [[ "$FORCE" == "false" ]]; then
    if ! diff -q "$src" "$dest" &>/dev/null; then
      warn "Preserved (customized): $dest"
    fi
  else
    safe_copy "$src" "$dest"
  fi
}

# Copy a directory tree. Customizable files are preserved unless --force.
# Usage: safe_copy_tree <src_dir> <dest_dir> [space-separated exclude subdirs]
safe_copy_tree() {
  local src_dir="${1%/}" dest_dir="$2" excludes="${3:-}"
  find "$src_dir" -type f | while read -r src_file; do
    local rel_path="${src_file#"$src_dir"/}"
    local skip=false
    for exclude in $excludes; do
      if [[ "$rel_path" == "$exclude"/* ]]; then
        skip=true
        break
      fi
    done
    [[ "$skip" == "true" ]] && continue
    if is_customizable "$rel_path" && [[ -f "$dest_dir/$rel_path" ]] && [[ "$FORCE" == "false" ]]; then
      if ! diff -q "$src_file" "$dest_dir/$rel_path" &>/dev/null; then
        warn "Preserved (customized): $dest_dir/$rel_path"
      fi
    else
      safe_copy "$src_file" "$dest_dir/$rel_path"
    fi
  done
}

# ---------------------------------------------------------------------------
# Detection functions
# ---------------------------------------------------------------------------

detect_languages() {
  local langs=()
  [[ -f "package.json" ]] && langs+=("JavaScript/TypeScript")
  [[ -f "go.mod" ]] && langs+=("Go")
  { [[ -f "pyproject.toml" ]] || [[ -f "requirements.txt" ]]; } && langs+=("Python")
  { [[ -f "build.gradle.kts" ]] || [[ -f "build.gradle" ]]; } && langs+=("Kotlin/Java")
  [[ -f "Cargo.toml" ]] && langs+=("Rust")
  [[ -f "Gemfile" ]] && langs+=("Ruby")
  [[ -f "mix.exs" ]] && langs+=("Elixir")
  [[ -f "composer.json" ]] && langs+=("PHP")
  [[ -f "Package.swift" ]] && langs+=("Swift")
  if [[ ${#langs[@]} -eq 0 ]]; then
    echo ""
    return
  fi
  local result
  result="$(printf '%s, ' "${langs[@]}")"
  echo "${result%, }"
}

detect_frameworks() {
  local fws=()
  if [[ -f "package.json" ]]; then
    grep -q '"react"' package.json 2>/dev/null && fws+=("React")
    grep -q '"next"' package.json 2>/dev/null && fws+=("Next.js")
    grep -q '"astro"' package.json 2>/dev/null && fws+=("Astro")
    grep -q '"vue"' package.json 2>/dev/null && fws+=("Vue")
    grep -q '"svelte"' package.json 2>/dev/null && fws+=("Svelte")
    grep -q '"express"' package.json 2>/dev/null && fws+=("Express")
    grep -q '"@nestjs/core"' package.json 2>/dev/null && fws+=("NestJS")
    grep -q '"fastify"' package.json 2>/dev/null && fws+=("Fastify")
    grep -q '"@tanstack/react-query"' package.json 2>/dev/null && fws+=("TanStack Query")
    grep -q '"react-router"' package.json 2>/dev/null && fws+=("React Router")
    grep -q '"tailwindcss"' package.json 2>/dev/null && fws+=("Tailwind CSS")
  fi
  if [[ -f "pyproject.toml" ]]; then
    grep -qi 'fastapi' pyproject.toml 2>/dev/null && fws+=("FastAPI")
    grep -qi 'django' pyproject.toml 2>/dev/null && fws+=("Django")
    grep -qi 'flask' pyproject.toml 2>/dev/null && fws+=("Flask")
  fi
  if [[ -f "go.mod" ]]; then
    grep -q 'gin-gonic' go.mod 2>/dev/null && fws+=("Gin")
    grep -q 'labstack/echo' go.mod 2>/dev/null && fws+=("Echo")
    grep -q 'gofiber' go.mod 2>/dev/null && fws+=("Fiber")
  fi
  if [[ ${#fws[@]} -eq 0 ]]; then
    echo ""
    return
  fi
  local result
  result="$(printf '%s, ' "${fws[@]}")"
  echo "${result%, }"
}

detect_package_manager() {
  if [[ -f "pnpm-lock.yaml" ]]; then echo "pnpm"
  elif [[ -f "yarn.lock" ]]; then echo "yarn"
  elif [[ -f "bun.lockb" ]] || [[ -f "bun.lock" ]]; then echo "bun"
  elif [[ -f "package-lock.json" ]] || [[ -f "package.json" ]]; then echo "npm"
  else echo ""
  fi
}

dir_description() {
  case "$1" in
    src)             echo "Source code" ;;
    test|tests)      echo "Test files" ;;
    docs)            echo "Documentation" ;;
    public|static)   echo "Static assets" ;;
    scripts)         echo "Scripts" ;;
    cmd)             echo "Entry points" ;;
    internal)        echo "Internal packages" ;;
    pkg)             echo "Public packages" ;;
    app)             echo "Application code" ;;
    lib)             echo "Library code" ;;
    config)          echo "Configuration" ;;
    migrations|db)   echo "Database" ;;
    api)             echo "API definitions" ;;
    specs)           echo "Specifications" ;;
    plans)           echo "Implementation plans" ;;
    packages)        echo "Monorepo packages" ;;
    apps)            echo "Monorepo apps" ;;
    components)      echo "UI components" ;;
    features)        echo "Feature modules" ;;
    pages|views)     echo "Page components" ;;
    routes)          echo "Route definitions" ;;
    hooks)           echo "Custom hooks" ;;
    utils|helpers)   echo "Utilities" ;;
    services)        echo "Service layer" ;;
    models)          echo "Data models" ;;
    types)           echo "Type definitions" ;;
    android)         echo "Android app" ;;
    ios)             echo "iOS app" ;;
    *)               echo "" ;;
  esac
}

detect_directory_structure() {
  local exclude_re="^(node_modules|dist|build|target|coverage|vendor|__pycache__)$"
  local result=""
  for dir in */; do
    [[ ! -d "$dir" ]] && continue
    local name="${dir%/}"
    echo "$name" | grep -qE "$exclude_re" && continue
    local desc
    desc=$(dir_description "$name")
    if [[ -n "$desc" ]]; then
      result+="$(printf '%-14s # %s' "${name}/" "$desc")"
    else
      result+="${name}/"
    fi
    result+=$'\n'
  done
  printf '%s' "$result" | sed '/^$/d'
}

detect_build_commands() {
  local lines=()

  # mise
  { [[ -f "mise.toml" ]] || [[ -f ".mise.toml" ]]; } \
    && lines+=("mise install          # Install tool versions")

  # Makefile targets
  if [[ -f "Makefile" ]]; then
    grep -q '^dev[[:space:]]*:' Makefile 2>/dev/null \
      && lines+=("make dev              # Start development")
    grep -q '^build[[:space:]]*:' Makefile 2>/dev/null \
      && lines+=("make build            # Build project")
    grep -q '^test[[:space:]]*:' Makefile 2>/dev/null \
      && lines+=("make test             # Run tests")
    grep -q '^lint[[:space:]]*:' Makefile 2>/dev/null \
      && lines+=("make lint             # Run linters")
  fi

  # package.json scripts (only if no Makefile)
  if [[ -f "package.json" ]] && [[ ! -f "Makefile" ]]; then
    local pm
    pm=$(detect_package_manager)
    local run="${pm} run"
    [[ "$pm" != "npm" ]] && run="$pm"
    grep -q '"dev"' package.json 2>/dev/null \
      && lines+=("${run} dev             # Start development")
    grep -q '"build"' package.json 2>/dev/null \
      && lines+=("${run} build           # Build project")
    grep -q '"test"' package.json 2>/dev/null \
      && lines+=("${run} test            # Run tests")
    grep -q '"lint"' package.json 2>/dev/null \
      && lines+=("${run} lint            # Run linters")
  fi

  # Go (only if no Makefile)
  if [[ -f "go.mod" ]] && [[ ! -f "Makefile" ]]; then
    lines+=("go build ./...        # Build")
    lines+=("go test ./...         # Run tests")
  fi

  # Python (only if no Makefile)
  if [[ -f "pyproject.toml" ]] && [[ ! -f "Makefile" ]]; then
    if command -v uv &>/dev/null; then
      lines+=("uv sync               # Install dependencies")
    else
      lines+=("pip install -e .      # Install dependencies")
    fi
    lines+=("pytest                # Run tests")
  fi

  if [[ ${#lines[@]} -eq 0 ]]; then
    echo ""
    return
  fi
  printf '%s\n' "${lines[@]}"
}

detect_relevant_instructions() {
  local files=()

  # Universal — always included
  files+=("testing.instructions.md")
  files+=("e2e-testing.instructions.md")
  files+=("integration-testing.instructions.md")

  # Docker
  { [[ -f "Dockerfile" ]] || [[ -f "docker-compose.yml" ]] \
    || [[ -f "docker-compose.yaml" ]] || [[ -f "compose.yml" ]] \
    || [[ -f "compose.yaml" ]]; } && files+=("docker.instructions.md")

  # JavaScript/TypeScript ecosystem
  if [[ -f "package.json" ]]; then
    files+=("typescript.instructions.md")
    grep -qE '"(react|next|astro|vue|svelte)"' package.json 2>/dev/null && files+=("web-security.instructions.md")
    { grep -qE '"(express|fastify|next|hono|koa)"' package.json 2>/dev/null \
      || grep -q '"@nestjs/core"' package.json 2>/dev/null; } && files+=("api-security.instructions.md")
    grep -q '"@nestjs/core"' package.json 2>/dev/null && files+=("nestjs.instructions.md")
    grep -q '"fastify"' package.json 2>/dev/null && files+=("fastify.instructions.md")
    grep -q '"react"' package.json 2>/dev/null && files+=("react.instructions.md")
    grep -q '"astro"' package.json 2>/dev/null && files+=("astro-mdx.instructions.md")
    grep -q '"@tanstack/react-query"' package.json 2>/dev/null && files+=("tanstack-query.instructions.md")
    grep -q '"react-router"' package.json 2>/dev/null && files+=("react-router.instructions.md")
    grep -q '"tailwindcss"' package.json 2>/dev/null && files+=("tailwind.instructions.md")
    grep -q '"zod"' package.json 2>/dev/null && files+=("zod.instructions.md")
    grep -q '"zustand"' package.json 2>/dev/null && files+=("zustand.instructions.md")
    grep -q '"xstate"' package.json 2>/dev/null && files+=("xstate.instructions.md")
    grep -qE '"(better-sqlite3|sql\.js|sqlite3)"' package.json 2>/dev/null && {
      files+=("sqlite.instructions.md")
      files+=("node-sqlite.instructions.md")
    }
  fi

  # Go
  if [[ -f "go.mod" ]]; then
    files+=("go.instructions.md")
    files+=("api-security.instructions.md")
    grep -q 'labstack/echo' go.mod 2>/dev/null && files+=("echo.instructions.md")
    grep -qE '(mattn/go-sqlite3|modernc\.org/sqlite)' go.mod 2>/dev/null && {
      files+=("sqlite.instructions.md")
      files+=("go-sqlite.instructions.md")
    }
    find . -name '*.go' -type f -exec grep -l '//go:embed' {} + 2>/dev/null \
      | grep -q . && files+=("go-embed.instructions.md")
    { [[ -f ".air.toml" ]] || [[ -f "air.toml" ]]; } && files+=("air.instructions.md")
  fi

  # Python
  if [[ -f "pyproject.toml" ]] || [[ -f "requirements.txt" ]]; then
    files+=("python.instructions.md")
    if grep -qi 'fastapi' pyproject.toml requirements.txt 2>/dev/null; then
      files+=("fastapi.instructions.md")
      files+=("pydantic.instructions.md")
      files+=("api-security.instructions.md")
    elif grep -qiE '(django|flask)' pyproject.toml requirements.txt 2>/dev/null; then
      files+=("api-security.instructions.md")
    fi
    grep -qiE '(pydantic|pydantic-settings)' pyproject.toml requirements.txt 2>/dev/null \
      && files+=("pydantic.instructions.md")
    grep -qi 'langchain' pyproject.toml requirements.txt 2>/dev/null && {
      files+=("langchain.instructions.md")
      files+=("pydantic.instructions.md")
    }
    grep -qi 'langgraph' pyproject.toml requirements.txt 2>/dev/null && {
      files+=("langgraph.instructions.md")
      files+=("pydantic.instructions.md")
    }
    grep -qiE '(openai|anthropic|google-genai|litellm)' pyproject.toml requirements.txt 2>/dev/null \
      && files+=("llm-service.instructions.md")
    grep -qiE '(sqlite|aiosqlite|sqlalchemy|alembic)' pyproject.toml requirements.txt 2>/dev/null && {
      files+=("sqlite.instructions.md")
      files+=("python-sqlite.instructions.md")
    }
  fi

  # Kotlin/Java (Android)
  if [[ -f "build.gradle.kts" ]] || [[ -f "build.gradle" ]]; then
    files+=("kotlin.instructions.md")
    files+=("android-security.instructions.md")
    grep -rqiE '(androidx\.room|sqlite)' . --include='*.gradle*' 2>/dev/null && {
      files+=("sqlite.instructions.md")
      files+=("android-sqlite.instructions.md")
    }
  fi

  # Deduplicate
  printf '%s\n' "${files[@]}" | sort -u
}

install_instructions() {
  local src_dir="$INSTALL_DIR/project/.github/instructions"
  local dest_dir="./.github/instructions"
  local relevant skipped=0

  relevant="$(detect_relevant_instructions)"

  # If only universal instructions are detected, keep the install minimal.
  local specific_count
  specific_count="$(echo "$relevant" | grep -cvxE '(testing|e2e-testing|integration-testing)\.instructions\.md' || true)"

  if [[ "$specific_count" -eq 0 ]]; then
    info "No language/framework detected — installing universal instructions only."
  fi

  for f in "$src_dir"/*.instructions.md; do
    [[ -f "$f" ]] || continue
    local name
    name="$(basename "$f")"
    if echo "$relevant" | grep -qxF "$name"; then
      safe_copy "$f" "$dest_dir/$name"
    else
      ((skipped++)) || true
    fi
  done

  if [[ "$skipped" -gt 0 ]]; then
    info "Skipped $skipped instruction files (not detected in this project)"
    info "All instructions available at: ~/.copilot-configs/project/.github/instructions/"
  fi
}

skill_dependency() {
  case "$1" in
    migrate-react-router)          echo "react-router" ;;
    manage-state-with-zustand)     echo "zustand" ;;
    model-state-with-xstate)       echo "xstate" ;;
    validate-with-zod)             echo "zod" ;;
    *)                             echo "" ;;
  esac
}

install_skills() {
  local src_dir="$INSTALL_DIR/project/.github/skills"
  local dest_dir="./.github/skills"
  local skipped=0

  for skill_dir in "$src_dir"/*/; do
    [[ ! -d "$skill_dir" ]] && continue
    local skill_name
    skill_name="$(basename "$skill_dir")"

    local dep
    dep="$(skill_dependency "$skill_name")"

    if [[ -n "$dep" ]]; then
      if ! { [[ -f "package.json" ]] && grep -q "\"$dep\"" package.json 2>/dev/null; }; then
        ((skipped++)) || true
        continue
      fi
    fi

    safe_copy_tree "$skill_dir" "$dest_dir/$skill_name"
  done

  if [[ "$skipped" -gt 0 ]]; then
    info "Skipped $skipped skills (not detected in this project)"
    info "All skills available at: ~/.copilot-configs/project/.github/skills/"
  fi
}

# ---------------------------------------------------------------------------
# Section replacement (replaces content between BEGIN/END markers)
# ---------------------------------------------------------------------------

replace_section() {
  local file="$1" marker="$2" content="$3"
  local begin="<!-- BEGIN:${marker} -->"
  local end="<!-- END:${marker} -->"
  local tmp="${file}.tmp"
  local in_section="false"

  while IFS= read -r line || [[ -n "$line" ]]; do
    if [[ "$line" == *"$begin"* ]]; then
      printf '%s\n' "$line"
      [[ -n "$content" ]] && printf '%s\n' "$content"
      in_section="true"
    elif [[ "$line" == *"$end"* ]]; then
      printf '%s\n' "$line"
      in_section="false"
    elif [[ "$in_section" == "false" ]]; then
      printf '%s\n' "$line"
    fi
  done < "$file" > "$tmp"

  mv "$tmp" "$file"
}

code_block() {
  local lang="${1:-}" content="$2"
  if [[ -n "$lang" ]]; then
    # shellcheck disable=SC2016
    printf '```%s\n%s\n```' "$lang" "$content"
  else
    # shellcheck disable=SC2016
    printf '```\n%s\n```' "$content"
  fi
}

# ---------------------------------------------------------------------------
# Configuration flows
# ---------------------------------------------------------------------------

fill_manually() {
  local file="$1"

  echo ""
  info "Answer the following questions about your project:"
  echo ""

  local project_what project_stack project_arch
  read -rp "  What does this project do? " project_what
  read -rp "  Stack (languages, frameworks, key deps): " project_stack
  read -rp "  Architecture (e.g., monolith, microservices, MVC): " project_arch

  local overview
  overview="$(printf '%s\n%s\n%s' \
    "- **What**: ${project_what}" \
    "- **Stack**: ${project_stack}" \
    "- **Architecture**: ${project_arch}")"
  replace_section "$file" "PROJECT_OVERVIEW" "$overview"
  ok "Project overview saved."

  # Directory structure
  echo ""
  info "Detecting directory structure..."
  local detected_dirs
  detected_dirs="$(detect_directory_structure)"

  if [[ -n "$detected_dirs" ]]; then
    echo ""
    echo "${detected_dirs//$'\n'/$'\n  '}" | sed '1s/^/  /'
    echo ""
    local use_dirs
    read -rp "  Use detected structure? [Y/n]: " use_dirs
    case "$use_dirs" in
      [Nn]*)
        info "Enter directory structure (one entry per line, empty line to finish):"
        local custom_dirs=""
        while IFS= read -r dline; do
          [[ -z "$dline" ]] && break
          custom_dirs+="${dline}"$'\n'
        done
        replace_section "$file" "DIRECTORY_STRUCTURE" "$(code_block "" "$custom_dirs")"
        ;;
      *)
        replace_section "$file" "DIRECTORY_STRUCTURE" "$(code_block "" "$detected_dirs")"
        ;;
    esac
  else
    info "No directories detected. Enter structure (one entry per line, empty line to finish):"
    local custom_dirs=""
    while IFS= read -r dline; do
      [[ -z "$dline" ]] && break
      custom_dirs+="${dline}"$'\n'
    done
    replace_section "$file" "DIRECTORY_STRUCTURE" "$(code_block "" "$custom_dirs")"
  fi
  ok "Directory structure saved."

  # Build commands
  echo ""
  info "Detecting build commands..."
  local detected_cmds
  detected_cmds="$(detect_build_commands)"

  if [[ -n "$detected_cmds" ]]; then
    echo ""
    echo "${detected_cmds//$'\n'/$'\n  '}" | sed '1s/^/  /'
    echo ""
    local use_cmds
    read -rp "  Use detected commands? [Y/n]: " use_cmds
    case "$use_cmds" in
      [Nn]*)
        info "Enter build/test commands (one per line, empty line to finish):"
        local custom_cmds=""
        while IFS= read -r cline; do
          [[ -z "$cline" ]] && break
          custom_cmds+="${cline}"$'\n'
        done
        replace_section "$file" "BUILD_COMMANDS" "$(code_block "bash" "$custom_cmds")"
        ;;
      *)
        replace_section "$file" "BUILD_COMMANDS" "$(code_block "bash" "$detected_cmds")"
        ;;
    esac
  else
    info "No build commands detected. Enter commands (one per line, empty line to finish):"
    local custom_cmds=""
    while IFS= read -r cline; do
      [[ -z "$cline" ]] && break
      custom_cmds+="${cline}"$'\n'
    done
    replace_section "$file" "BUILD_COMMANDS" "$(code_block "bash" "$custom_cmds")"
  fi
  ok "Build commands saved."
}

fill_auto_detect() {
  local file="$1"

  info "Analyzing project..."
  echo ""

  local languages frameworks stack
  languages="$(detect_languages)"
  frameworks="$(detect_frameworks)"

  if [[ -n "$languages" ]]; then
    stack="$languages"
    [[ -n "$frameworks" ]] && stack="${languages}, ${frameworks}"
  else
    stack="Not detected"
  fi

  local dirs cmds
  dirs="$(detect_directory_structure)"
  cmds="$(detect_build_commands)"

  # Show summary
  info "Detected configuration:"
  echo ""
  echo "  Stack: $stack"
  if [[ -n "$dirs" ]]; then
    echo ""
    echo "  Directories:"
    echo "    ${dirs//$'\n'/$'\n    '}"
  fi
  if [[ -n "$cmds" ]]; then
    echo ""
    echo "  Build commands:"
    echo "    ${cmds//$'\n'/$'\n    '}"
  fi
  echo ""

  local apply
  read -rp "  Apply detected values? [Y/n]: " apply
  case "$apply" in
    [Nn]*)
      info "Skipped. Edit .github/copilot-instructions.md manually."
      return
      ;;
  esac

  # Apply overview (What and Architecture need manual editing)
  local overview
  overview="$(printf '%s\n%s\n%s' \
    "- **What**: [Edit: add project description]" \
    "- **Stack**: ${stack}" \
    "- **Architecture**: [Edit: add architecture pattern]")"
  replace_section "$file" "PROJECT_OVERVIEW" "$overview"

  if [[ -n "$dirs" ]]; then
    replace_section "$file" "DIRECTORY_STRUCTURE" "$(code_block "" "$dirs")"
  fi

  if [[ -n "$cmds" ]]; then
    replace_section "$file" "BUILD_COMMANDS" "$(code_block "bash" "$cmds")"
  fi

  ok "Configuration applied!"
  warn "Review the file — 'What' and 'Architecture' still need manual editing."
}

configure_instructions() {
  local instructions="./.github/copilot-instructions.md"

  if [[ ! -f "$instructions" ]]; then
    warn "copilot-instructions.md not found. Skipping configuration."
    return
  fi

  if ! grep -q '<!-- BEGIN:PROJECT_OVERVIEW -->' "$instructions" 2>/dev/null; then
    warn "No configuration markers found in copilot-instructions.md."
    warn "The file may have been customized already. Skipping."
    return
  fi

  if [[ ! -t 0 ]]; then
    warn "Not running in a terminal. Skipping interactive configuration."
    return
  fi

  echo ""
  info "Configure copilot-instructions.md"
  echo ""
  echo "  1) Fill interactively"
  echo "  2) Auto-detect from project files"
  echo "  3) Skip (edit later)"
  echo ""

  local choice
  read -rp "  Choose [1/2/3]: " choice

  case "$choice" in
    1) fill_manually "$instructions" ;;
    2) fill_auto_detect "$instructions" ;;
    *) info "Skipping configuration." ;;
  esac

  echo ""
  info "Tip: refine with GitHub Copilot by running 'copilot' and asking:"
  info '  "Review and improve .github/copilot-instructions.md for this project"'
}

# ---------------------------------------------------------------------------
# Parse args
# ---------------------------------------------------------------------------

while [[ $# -gt 0 ]]; do
  case "$1" in
    --project) PROJECT=true; shift ;;
    --force)   FORCE=true; shift ;;
    --configure) CONFIGURE=true; shift ;;
    --help|-h)
      cat <<EOF
copilot-configs installer

Usage:
  install.sh                          Install/update global Copilot configs
  install.sh --project                Apply project template to current directory
  install.sh --project --configure    Apply template and configure placeholders
  install.sh --configure              Configure copilot-instructions.md in current directory
  install.sh --force                  Overwrite existing files

Options:
  --project     Copy .github/ template into the current working directory
  --configure   Interactively configure copilot-instructions.md placeholders
  --force       Overwrite all files, including customized ones (default: preserve)
  --help        Show this help message

Examples:
  # Install global configs
  curl -fsSL https://raw.githubusercontent.com/woliveiras/copilot-configs/main/install.sh | bash

  # Apply project template
  ~/.copilot-configs/install.sh --project

  # Apply project template and configure placeholders
  ~/.copilot-configs/install.sh --project --configure

  # Reconfigure an existing project
  ~/.copilot-configs/install.sh --configure
EOF
      exit 0
      ;;
    *)
      error "Unknown option: $1"
      exit 1
      ;;
  esac
done

# ---------------------------------------------------------------------------
# Project mode: copy template into current directory
# ---------------------------------------------------------------------------

if [[ "$PROJECT" == "true" ]]; then
  if [[ ! -d "$INSTALL_DIR" ]]; then
    error "copilot-configs not installed. Run the installer first:"
    error "  curl -fsSL https://raw.githubusercontent.com/woliveiras/copilot-configs/main/install.sh | bash"
    exit 1
  fi

  info "Applying project template to $(pwd)..."

  # Copy root-level project contract
  safe_copy_customizable "$INSTALL_DIR/project/AGENTS.md" "./AGENTS.md" "AGENTS.md"

  # Copy .github/ template (instructions and skills installed separately based on detection)
  safe_copy_tree "$INSTALL_DIR/project/.github" "./.github" "instructions skills"

  # Install only relevant instruction files and skills
  install_instructions
  install_skills

  echo ""
  ok "Project template applied!"

  if [[ "$CONFIGURE" == "true" ]]; then
    configure_instructions
  fi

  echo ""
  info "Next steps:"
  if [[ "$CONFIGURE" != "true" ]]; then
    info "  - Configure: ~/.copilot-configs/install.sh --configure"
  fi
  info "  - Add more instructions: ls ~/.copilot-configs/project/.github/instructions/"
  info "  - Edit .github/hooks/guardrails-rules.txt to customize guardrails"
  exit 0
fi

# ---------------------------------------------------------------------------
# Configure mode (standalone)
# ---------------------------------------------------------------------------

if [[ "$CONFIGURE" == "true" ]]; then
  configure_instructions
  exit 0
fi

# ---------------------------------------------------------------------------
# Global install mode
# ---------------------------------------------------------------------------

info "Installing copilot-configs..."

# Clone or update
if [[ -d "$INSTALL_DIR/.git" ]]; then
  info "Updating existing installation..."
  git -C "$INSTALL_DIR" pull --ff-only --quiet 2>/dev/null || {
    warn "Could not fast-forward. Re-cloning..."
    rm -rf "$INSTALL_DIR"
    git clone --quiet "$REPO_URL" "$INSTALL_DIR"
  }
  ok "Updated: $INSTALL_DIR"
else
  if [[ -d "$INSTALL_DIR" ]]; then
    warn "Removing non-git directory at $INSTALL_DIR"
    rm -rf "$INSTALL_DIR"
  fi
  git clone --quiet "$REPO_URL" "$INSTALL_DIR"
  ok "Cloned: $INSTALL_DIR"
fi

# Detect VS Code user directory
VSCODE_USER_DIR="$(detect_vscode_dir)"

if [[ ! -d "$VSCODE_USER_DIR" ]]; then
  warn "VS Code user directory not found: $VSCODE_USER_DIR"
  warn "Skipping global prompt installation. Install VS Code first."
else
  # Install global prompts
  info "Installing global prompts..."
  PROMPTS_DIR="$VSCODE_USER_DIR/prompts"
  mkdir -p "$PROMPTS_DIR"

  if [[ -d "$INSTALL_DIR/user/prompts" ]]; then
    for f in "$INSTALL_DIR/user/prompts"/*.prompt.md; do
      [[ -f "$f" ]] || continue
      safe_copy "$f" "$PROMPTS_DIR/$(basename "$f")"
    done
  fi

  if [[ -f "$INSTALL_DIR/user/copilot-instructions.md" ]]; then
    info "Installing global Copilot instruction bootstrap..."
    safe_copy_customizable \
      "$INSTALL_DIR/user/copilot-instructions.md" \
      "$VSCODE_USER_DIR/copilot-instructions.md" \
      "copilot-instructions.md"
  fi
fi

# Check for mise
if ! command -v mise &> /dev/null; then
  echo ""
  info "mise (tool version manager) is not installed."
  info "mise manages Go, Node, pnpm, Python, uv, and other tools."
  info "Install it with: curl https://mise.run | sh"
  info "Learn more: https://mise.jdx.dev"
fi

echo ""
ok "copilot-configs installed!"
info "Global prompts are available in VS Code (type / in Copilot chat)."
info ""
info "To apply project template to a repository:"
info "  cd your-project && ~/.copilot-configs/install.sh --project"
