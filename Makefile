build:
	mkdir "bin" && GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main main.go