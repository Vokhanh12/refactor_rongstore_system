#!/usr/bin/env bash
set -euo pipefail

# Directories
PROTO_DIR="proto"
GO_OUT="../server/gen/proto"
SWAGGER_OUT="./openapi"
SERVICE_PROTO_DIR="$PROTO_DIR/iam/v1/services"

# Colors
BOLD="\033[1m"; GREEN="\033[32m"; YELLOW="\033[33m"; RED="\033[31m"; NC="\033[0m"

# Logging helpers
log()  { echo -e "${BOLD}${GREEN}[OK]${NC} $*"; }
warn() { echo -e "${BOLD}${YELLOW}[WARN]${NC} $*"; }
err()  { echo -e "${BOLD}${RED}[ERR]${NC} $*" 1>&2; }

# Try run command without failing immediately
try_run() {
  set +e
  "$@"
  local status=$?
  set -e
  return $status
}

echo -e "${BOLD}=== Generate gRPC + Swagger ===${NC}"

# Clean output directories
rm -rf "$GO_OUT" "$SWAGGER_OUT"
mkdir -p "$GO_OUT" "$SWAGGER_OUT"
log "Clean output OK"

# Collect all proto files except google/*.proto
PROTO_FILES=()
while IFS= read -r f; do
  PROTO_FILES+=("$f")
done < <(find "$PROTO_DIR" -name '*.proto' ! -path "$PROTO_DIR/google/*" | sort)

# --- Generate Go code ---
echo -e "\n${BOLD}---> Generating Go gRPC code${NC}"
protoc \
  -I "$PROTO_DIR" \
  -I "$PROTO_DIR/google" \
  --go_out="$GO_OUT" \
  --go_opt=paths=source_relative \
  --go-grpc_out="$GO_OUT" \
  --go-grpc_opt=paths=source_relative \
  "${PROTO_FILES[@]}"
log "Go code generated"

# --- Generate Swagger/OpenAPI ---
echo -e "\n${BOLD}---> Generating Swagger (OpenAPI)${NC}"
SERVICE_PROTOS=()
while IFS= read -r f; do
  SERVICE_PROTOS+=("$f")
done < <(find "$SERVICE_PROTO_DIR" -name "*.proto" | sort)

[[ ${#SERVICE_PROTOS[@]} -gt 0 ]] || { err "No service proto found in $SERVICE_PROTO_DIR"; exit 1; }

protoc \
  -I "$PROTO_DIR" \
  -I "$PROTO_DIR/google" \
  --openapiv2_out "$SWAGGER_OUT" \
  --openapiv2_opt logtostderr=true \
  "${SERVICE_PROTOS[@]}"

log "Swagger generated"

echo -e "\n${BOLD}=== DONE ===${NC}"
echo "Go      => $GO_OUT"
echo "Swagger => $SWAGGER_OUT"