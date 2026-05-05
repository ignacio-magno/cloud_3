build:
	go build -o bin/server cmd/server/main.go
	go build -o bin/webhook cmd/webhook/main.go

clean:
	rm -rf bin/
