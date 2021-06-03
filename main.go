package main

import (
	"github.com/subuta/play-with-stdio-transport/internals"
	"github.com/subuta/play-with-stdio-transport/pkg/parser"
)

func main() {
	internals.LoopRead(parser.JsonParserImpl{})
}
