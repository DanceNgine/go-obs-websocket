package obsws

// Request is a request to obs-websocket.
type Request interface {
	ID() string
	Type() string
}

// Response is a response from obs-websocket.
type Response interface {
	ID() string
	Stat() string
	Err() string
}

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#requests
type _request struct {
	MessageID   string `json:"message-id"`
	RequestType string `json:"request-type"`
}

func (r _request) ID() string { return r.MessageID }

func (r _request) Type() string { return r.RequestType }

// https://github.com/Palakis/obs-websocket/blob/master/docs/generated/protocol.md#requests
type _response struct {
	MessageID string `json:"message-id"`
	Status    string `json:"status"`
	Error     string `json:"error"`
}

func (r _response) ID() string { return r.MessageID }

func (r _response) Stat() string { return r.Status }

func (r _response) Err() string { return r.Error }
