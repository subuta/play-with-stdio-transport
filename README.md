# play-with-stdio-transport
Playground for stdio transport between parent-child process

Demo "server" will echo your input to stdout.

## How to develop

```bash
# Run main.go
go run .

# Build executable binary.
make build

# Run binary
./server

# Run binary from Node.js
cd examples/node; npm start

# Run binary from Ruby
cd examples/ruby; ruby ./main.rb
```