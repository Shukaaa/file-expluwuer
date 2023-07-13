go build -o bin/main main.go # Build the binary
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/main.exe main.go # Build the executable