#!/usr/bin/env bash
set -euo pipefail

GO_OUT="../server/gen/proto"
SWAGGER_OUT="./openapi"

BOLD="\033[1m"
GREEN="\033[32m"
NC="\033[0m"

log() {
    echo -e "${BOLD}${GREEN}[OK]${NC} $*"
}

echo -e "${BOLD}=== Generate gRPC + Swagger ===${NC}"

# Clean generated output
rm -rf "$GO_OUT" "$SWAGGER_OUT"

mkdir -p "$GO_OUT"
mkdir -p "$SWAGGER_OUT"

log "Clean output OK"

# Validate proto
echo -e "\n${BOLD}---> Building protobuf${NC}"

buf build

log "Proto build OK"

# Generate all code
echo -e "\n${BOLD}---> Generating Go + gRPC + Swagger${NC}"

buf generate

log "Generation OK"

echo -e "\n${BOLD}=== DONE ===${NC}"

echo "Go      => $GO_OUT"
echo "Swagger => $SWAGGER_OUT"