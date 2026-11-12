.PHONY: build test check release-snapshot

build:
	go build -o bin/fireweed .

test:
	go test ./...

check:
	go test -race ./...
	go vet ./...

# Builds all four archives and the tap formula without tokens or publication.
release-snapshot:
	goreleaser check
	goreleaser release --snapshot --clean
	python3 scripts/homebrew.py dist > dist/fireweed.rb
