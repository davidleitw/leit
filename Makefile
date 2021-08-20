build:
	go build -o leit main.go

set_up_database:
	mkdir ~/.local/share/leit
	touch ~/.local/share/leit/leit.db

clear_database:
	rm ~/.local/share/leit/leit.db

test:
	go test -v . 