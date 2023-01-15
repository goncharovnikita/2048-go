.PHONY: build
build:
	go build -o ./bin/2048 ./main.go

.PHONY: run
run: build
	./bin/2048

