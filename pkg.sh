GOARCH=386      GOOS=linux      go build -o bin/archiver32      cmd/main.go
GOARCH=amd64    GOOS=linux      go build -o bin/archiver        cmd/main.go
GOARCH=amd64    GOOS=windows    go build -o bin/archiver.exe    cmd/main.go