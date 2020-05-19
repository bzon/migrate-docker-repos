build:
	rm -fr bin/**
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/migrate-docker-repos-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o bin/migrate-docker-repos-darwin-amd64

