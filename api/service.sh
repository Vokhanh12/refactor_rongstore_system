#!/usr/bin/env bash
set -euo pipefail

PROTO_DIR="proto"
GO_OUT="../server/api"
DART_OUT="../client/lib/features/api"
SWAGGER_OUT="./openapi"
SERVICE_PROTO_DIR="$PROTO_DIR/iam/v1/services"

BOLD="\033[1m"; GREEN="\033[32m"; YELLOW="\033[33m"; RED="\033[31m"; NC="\033[0m"

log()  { echo -e "${BOLD}${GREEN}[OK]${NC} $*"; }
warn() { echo -e "${BOLD}${YELLOW}[WARN]${NC} $*"; }
err()  { echo -e "${BOLD}${RED}[ERR]${NC} $*" 1>&2; }

try_run() {
  set +e
  "$@"
  local status=$?
  set -e
  return $status
}

echo -e "${BOLD}=== Generate gRPC + Swagger ===${NC}"

rm -rf "$GO_OUT" "$DART_OUT" "$SWAGGER_OUT" 2>/dev/null || true
mkdir -p "$GO_OUT" "$DART_OUT" "$SWAGGER_OUT"
log "Clean output OK"

### Collect proto files
PROTO_FILES=()
while IFS= read -r f; do
  PROTO_FILES+=("$f")
done < <(
  find "$PROTO_DIR" -name '*.proto' ! -path "$PROTO_DIR/google/*" | sort
)

### Go (critical)
echo -e "\n${BOLD}---> Go${NC}"
protoc \
  -I "$PROTO_DIR" \
  -I "$PROTO_DIR/google" \
  --go_out="$GO_OUT" \
  --go_opt=paths=source_relative \
  --go-grpc_out="$GO_OUT" \
  --go-grpc_opt=paths=source_relative \
  "${PROTO_FILES[@]}"
log "Go generated"

### Dart (optional)
echo -e "\n${BOLD}---> Dart (optional)${NC}"
if command -v protoc-gen-dart >/dev/null 2>&1; then
  try_run protoc \
    -I "$PROTO_DIR" \
    -I "$PROTO_DIR/google" \
    --dart_out=grpc:"$DART_OUT" \
    "${PROTO_FILES[@]}" \
    && log "Dart generated" \
    || warn "Dart failed"
else
  warn "Skip Dart (plugin not installed)"
fi

### Swagger (🔥 MUST)
echo -e "\n${BOLD}---> Swagger (OpenAPI)${NC}"

SERVICE_PROTOS=()
while IFS= read -r f; do
  SERVICE_PROTOS+=("$f")
done < <(
  find "$SERVICE_PROTO_DIR" -name "*.proto" | sort
)

[[ ${#SERVICE_PROTOS[@]} -gt 0 ]] || { err "No service proto found"; exit 1; }

protoc \
  -I "$PROTO_DIR" \
  -I "$PROTO_DIR/google" \
  --openapiv2_out "$SWAGGER_OUT" \
  --openapiv2_opt logtostderr=true \
  "${SERVICE_PROTOS[@]}"

log "Swagger generated"

echo -e "\n${BOLD}=== DONE ===${NC}"
echo "Go      => $GO_OUT"
echo "Dart    => $DART_OUT"
echo "Swagger => $SWAGGER_OUT"
