redis:
	docker compose up -d

run:
	go run src/main.go

gen:
	buf generate
