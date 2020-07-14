gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb --go_opt=paths=source_relative

clean:
	rm pb/*.go

run:
	go run main.go

test:
	go test -cover -race ./...