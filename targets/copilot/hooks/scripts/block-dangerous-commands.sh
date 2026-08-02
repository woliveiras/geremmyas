#!/usr/bin/env bash
# block-dangerous-commands.sh
#
# Copilot hook script for PreToolUse events.
# Reads the tool invocation from stdin (JSON), checks the command
# against rules in guardrails-rules.txt, and returns a permissionDecision.
#
# Exit codes:
#   0 = allowed (or tool is not a terminal command)
#
# Output (JSON to stdout):
#   {"permissionDecision": "allow"}
#   {"permissionDecision": "deny", "reason": "..."}
#   {"permissionDecision": "ask", "reason": "..."}

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
RULES_FILE="${SCRIPT_DIR}/../guardrails-rules.txt"

# Read the hook payload from stdin
PAYLOAD=$(cat)

json_string_field() {
  local field="$1"
  if command -v jq >/dev/null 2>&1; then
    printf '%s' "$PAYLOAD" | jq -er --arg field "$field" '.[$field] | select(type == "string")' 2>/dev/null
  elif command -v python3 >/dev/null 2>&1; then
    printf '%s' "$PAYLOAD" | python3 -c 'import json, sys; value = json.load(sys.stdin).get(sys.argv[1]); assert isinstance(value, str); print(value)' "$field" 2>/dev/null
  else
    return 1
  fi
}

if ! TOOL_NAME=$(json_string_field toolName); then
  echo '{"permissionDecision": "deny", "reason": "Unable to decode terminal policy input."}'
  exit 0
fi

# Only check terminal commands
if [[ "$TOOL_NAME" != *"terminal"* && "$TOOL_NAME" != *"shell"* && "$TOOL_NAME" != "runInTerminal" ]]; then
  echo '{"permissionDecision": "allow"}'
  exit 0
fi

if ! COMMAND=$(json_string_field command) || [[ -z "$COMMAND" ]]; then
  echo '{"permissionDecision": "deny", "reason": "Unable to decode terminal command."}'
  exit 0
fi

# Missing policy is a broken safety harness; fail closed for terminal commands.
if [[ ! -f "$RULES_FILE" ]]; then
  echo '{"permissionDecision": "deny", "reason": "Geremmyas guardrail rules are missing."}'
  exit 0
fi

# Check command against each rule
while IFS= read -r line; do
  # Skip comments and blank lines
  [[ "$line" =~ ^[[:space:]]*# ]] && continue
  [[ "$line" =~ ^[[:space:]]*$ ]] && continue

  # Parse action and pattern
  ACTION="${line%%[[:space:]]*}"
  PATTERN="${line#*[[:space:]]}"
  # Trim leading whitespace from pattern
  PATTERN="${PATTERN#"${PATTERN%%[![:space:]]*}"}"

  # Skip malformed lines
  [[ -z "$ACTION" || -z "$PATTERN" ]] && continue

  # Check if command matches pattern
  if echo "$COMMAND" | grep -qiE "$PATTERN"; then
    case "$ACTION" in
      BLOCK)
        echo "{\"permissionDecision\": \"deny\", \"reason\": \"Blocked by guardrail: pattern '$PATTERN' matched\"}"
        exit 0
        ;;
      ASK)
        echo "{\"permissionDecision\": \"ask\", \"reason\": \"Protected boundary requires explicit authority: pattern '$PATTERN' matched\"}"
        exit 0
        ;;
      ALLOW)
        echo '{"permissionDecision": "allow"}'
        exit 0
        ;;
    esac
  fi
done < "$RULES_FILE"

# No rules matched — allow
echo '{"permissionDecision": "allow"}'
exit 0
