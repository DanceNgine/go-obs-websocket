package obsws

import (
	"errors"
)

var ErrUnknownEventType = errors.New("unknown event type")

// AddEventHandler adds a handler function for a given event type.
func (c *Client) AddEventHandler(eventType string, handler func(Event)) error {
	if eventMap[eventType] == nil {
		return ErrUnknownEventType
	}
	c.handlers[eventType] = handler
	return nil
}

// RemoveEventHandler removes the handler for a given event type.
func (c *Client) RemoveEventHandler(eventType string) {
	delete(c.handlers, eventType)
}

// handleEvent runs an event's handler if it exists.
func (c *Client) handleEvent(m map[string]interface{}) {
	event := eventMap[m["update-type"].(string)]
	if event == nil {
		logger.Warning("unknown event type:", m["update-type"])
		return
	}

	if err := mapToStruct(m, event); err != nil {
		logger.Warning("event handler failed:", err)
		return
	}

	handler := c.handlers[event.Type()]
	if handler != nil {
		go handler(event)
	}
}
