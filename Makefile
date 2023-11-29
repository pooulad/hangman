build :
	@go build -o ./bin/hangman ./main.go

run : build
	@./bin/hangman

tidy : 
	go mod tidy