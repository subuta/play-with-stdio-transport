package parser

type EventParser interface {
	OnData(text string) error
}
