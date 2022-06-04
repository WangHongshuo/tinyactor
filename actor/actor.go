package actor

type Actor interface {
	Receive(c Context)
}
