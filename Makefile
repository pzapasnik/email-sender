templ:
	templ generate

run: templ
	go run cmd/email-sender/main.go

build: templ
	go build -o bin/email-sender cmd/email-sender/main.go
