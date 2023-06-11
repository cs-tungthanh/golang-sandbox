redis:
	docker compose up -d

app:
	go run src/main.go

gen:
    protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
