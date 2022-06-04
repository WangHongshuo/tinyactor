package actor

import (
	"errors"
	"fmt"
	"strings"
)

const (
	INVALID = iota
	IDLE
	RUNNING
)

//
type actorContext struct {
	actor       Actor
	actorSystem *ActorSystem
	parent      *PID
	self        *PID
	children    []*PID
	message     interface{}
}

// NewActorContext
func NewActorContext(pid *PID, as *ActorSystem) *actorContext {
	if pid == nil || as == nil {
		return nil
	}

	a := &actorContext{
		self:        pid,
		actorSystem: as,
	}
	if pid.mailBox == nil {
		pid.mailBox = NewMailBox(a, 10)
	}
	return a
}

// Context: InfoPart

// Parent
func (a *actorContext) Parent() *PID {
	if a == nil {
		return nil
	}
	return a.parent
}

// Self
func (a *actorContext) Self() *PID {
	if a == nil {
		return nil
	}
	return a.self
}

// Actor
func (a *actorContext) Actor() Actor {
	if a == nil {
		return nil
	}
	return a.actor
}

// Context: BasePart

func (a *actorContext) Children() []*PID {
	if a == nil {
		return nil
	}
	return a.children
}

// Context: SenderPart

// Send
func (a *actorContext) Send(pid *PID, msg interface{}) {
	if a == nil || pid == nil {
		return
	}
	if m, ok := msg.(*MessageEnvelope); ok {
		pid.mailBox.inBox(m)
		return
	}
	pid.mailBox.inBox(WarpEnvelope(a.self, msg))

}

// Sender
func (a *actorContext) Sender() *PID {
	if a == nil {
		return nil
	}
	return UnwrapEnvelopeSender(a.message)
}

// Context: ReceiverPart
// Receive
func (a *actorContext) Receive(message *MessageEnvelope) {
	if a == nil || message == nil {
		return
	}
	a.message = message
	a.defaultReceive()
	a.message = nil
}

func (a *actorContext) defaultReceive() {
	a.actor.Receive(a)
}

// Context: SpawnerPart
// SpawnFromInstance
func (a *actorContext) SpawnFromInstance(id string, actor Actor) (*PID, error) {
	if a == nil {
		return nil, errors.New("actorContext is nil")
	}
	if id == INVALID_PID || strings.Contains(id, "/") {
		return nil, fmt.Errorf("invalid id: %v", id)
	}
	pid := NewPID(a.self.Id + "/" + id)
	actorContext := NewActorContext(pid, a.actorSystem)
	actorContext.actor = actor
	actorContext.self = pid
	actorContext.parent = a.self
	a.children = append(a.children, pid)
	a.actorSystem.ProcessRegistry.Register(pid)
	pid.mailBox.inBox(WarpEnvelope(pid, &Started{}))
	return pid, nil
}

// Context: MessagePart
// Message
func (a *actorContext) Message() interface{} {
	return UnwarpEnvelopeMessage(a.message)
}

// Pid
func (a *actorContext) Pid() *PID {
	return a.self
}
