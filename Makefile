redis:
	docker compose up -d

run:
	@go run cmd/main.go

gen:
	buf generate

.PHONY: shell

shell:
	@if ! command -v devbox >/dev/null 2>&1; then curl -fsSL https://get.jetpack.io/devbox | bash; fi
	@devbox install
	@devbox shell
