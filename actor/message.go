package actor

type Message struct {
	Sender  *PID
	Payload interface{}
}
