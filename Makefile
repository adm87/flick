.PHONY: aseprite-keys

aseprite-keys:
	go run ./cmd/tools/gen-aseprite-keys/main.go --output ./generated/aseprite --working-dir .