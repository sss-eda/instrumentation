package instrumentation

type EventID string

type EventKind string

type EventPayload []byte

type Event struct {
	ID      EventID      `json:"id"`
	Kind    EventKind    `json:"kind"`
	Payload EventPayload `json:"payload"`
}
