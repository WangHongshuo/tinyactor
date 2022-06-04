package actor

type MessageEnvelope struct {
	Sender  *PID
	Message interface{}
}

// WarpEnvelope
func WarpEnvelope(sender *PID, message interface{}) *MessageEnvelope {
	if m, ok := message.(*MessageEnvelope); ok {
		return m
	}
	return &MessageEnvelope{Sender: sender, Message: message}
}

// UnwarpEnvelope
func UnwarpEnvelope(message interface{}) (interface{}, *PID) {
	if m, ok := message.(*MessageEnvelope); ok {
		return m.Message, m.Sender
	}
	return message, nil
}

// UnwarpEnvelopeMessage
func UnwarpEnvelopeMessage(message interface{}) interface{} {
	if m, ok := message.(*MessageEnvelope); ok {
		return m.Message
	}
	return message
}

// UnwrapEnvelopeSender
func UnwrapEnvelopeSender(message interface{}) *PID {
	if m, ok := message.(*MessageEnvelope); ok {
		return m.Sender
	}
	return nil
}
