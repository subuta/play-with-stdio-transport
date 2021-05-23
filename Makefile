build:
	go build -o server
	cp ./server ./examples/node
	cp ./server ./examples/ruby