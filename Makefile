redis:
	docker compose up -d

app:
	go run src/main.go

gen:
	buf generate
