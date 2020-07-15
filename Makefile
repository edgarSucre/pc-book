gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb --go_opt=paths=source_relative

clean:
	rm pb/*.go

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

server:
	go run cmd/server/main.go -port 8080

test:
	go test -cover -race ./...