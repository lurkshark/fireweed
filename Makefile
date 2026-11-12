.PHONY: build test check release-snapshot

# Fixture subprocesses need a stable identity and bytewise filename ordering.
test check: export GIT_AUTHOR_NAME := Fireweed Test
test check: export GIT_AUTHOR_EMAIL := fixture@example.invalid
test check: export GIT_COMMITTER_NAME := Fireweed Test
test check: export GIT_COMMITTER_EMAIL := fixture@example.invalid
test check: export LC_ALL := C

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
