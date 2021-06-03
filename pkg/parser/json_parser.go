package parser

import (
	"encoding/json"
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/subuta/play-with-stdio-transport/pkg/event"
	"log"
)

type JsonParserImpl struct{}

func parseJsonEvent(text string) (*event.JsonEvent, error) {
	var evt event.JsonEvent

	err := json.Unmarshal([]byte(text), &evt)
	if err != nil {
		return nil, err
	}

	return &evt, nil
}

func (parser JsonParserImpl) OnData(text string) error {
	fmt.Print("Got Json: ")

	evt, err := parseJsonEvent(text)
	if err != nil {
		log.Fatal(err)
	}

	pp.Println(evt)

	return nil
}
