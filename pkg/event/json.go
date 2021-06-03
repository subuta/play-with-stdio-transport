package event

type JsonEvent struct {
	Type string `json:"type"`
	Body map[string]interface{} `json:"body"`
}