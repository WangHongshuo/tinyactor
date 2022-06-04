package actor

type Context interface {
	infoPart
	basePart
	senderPart
	receiverPart
	messagePart
	spawnerPart
}

type SpawnerContext interface {
	infoPart
	spawnerPart
}

type SenderContext interface {
	infoPart
	senderPart
	messagePart
}

type ReceiverContext interface {
	infoPart
	receiverPart
	messagePart
}

type infoPart interface {
	Parent() *PID
	Self() *PID
	Actor() Actor
}

type basePart interface {
	Children() []*PID
}

type senderPart interface {
	Sender() *PID
	Send(pid *PID, msg interface{})
}

type receiverPart interface {
	Receive(msg *MessageEnvelope)
}

type spawnerPart interface {
	SpawnFromInstance(id string, actor Actor) (*PID, error)
}

type messagePart interface {
	Message() interface{}
}
