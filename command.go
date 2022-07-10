package instrumentation

type CommandID string

type CommandKind string

type CommandPayload []byte

type Command struct {
	ID      CommandID      `json:"id"`
	Kind    CommandKind    `json:"kind"`
	Payload CommandPayload `json:"payload"`
}
