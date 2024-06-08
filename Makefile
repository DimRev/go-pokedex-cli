build: 
	@echo building...
	@go build -o ./bin/go-pokemon-cli
	@echo build succeeded

start:
	@echo starting...
	@./bin/go-pokemon-cli
	@echo stopped!

run: build start