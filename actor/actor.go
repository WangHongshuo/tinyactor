package actor

type Actor interface {
	Pid() string
	Sender
	Receiver
}

type Sender interface {
	SendTo(receiver string, msg interface{})
}

type Receiver interface {
	Receive(msg *Mail)
}
