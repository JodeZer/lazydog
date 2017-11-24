build:
	go build -o bin/lazydog cmd/lazydog.go

install: build
	cp bin/lazydog /usr/local/bin

jumptest: build
	./bin/lazydog jump -d example
	go run example/*.go

overtest: build
	./bin/lazydog over -d example
	go run example/*.go
