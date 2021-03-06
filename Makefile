BINARY=engine
get: 
	go get

run: 
	go run main.go

test: 
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o ${BINARY} app/*.go
