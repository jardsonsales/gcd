default:
	GOOS=linux GOARCH=amd64 go build -o ./bin/gcd ./cmd/gcd/main.go

windows:
	GOOS=windows GOARCH=amd64 go build -o ./bin/gcd.exe ./cmd/gcd/main.go

run:
	go run ./cmd/gcd/main.go

test:
	go test ./...

test-cover:
	go test ./... -cover

symlink:
	sudo cp -f ./bin/gcd /usr/local/bin
