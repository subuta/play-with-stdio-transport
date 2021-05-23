package main

import (
	"github.com/k0kubun/pp"

	"github.com/subuta/play-with-stdio-transport/internals"
)

type eventParserImpl struct {}

func (parser eventParserImpl) OnData(text string) error {
	pp.Println("Got: ", text)
	return nil
}

func main() {
	internals.LoopRead(eventParserImpl{})
}
