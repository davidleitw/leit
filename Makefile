build:
	go build -o leit main.go

help: leit
	./leit --help