package actor

//
type Actor struct {
	pid     *Pid
	mailbox *MailBox
}

type ActorAbility interface {
	Receive(sender *Pid, msg interface{}) error
	Send(sender, receiver *Pid, msg interface{})
}

//
func (a *Actor) init() {

}

//
func NewActor() *Actor {
	a := &Actor{}
	a.pid = &Pid{}
	a.mailbox = &MailBox{}
	a.init()
	return a
}

func (a *Actor) Receive(sender *Pid, msg interface{}) error {
	return nil
}

func (a *Actor) Send(sender, receiver *Pid, msg interface{}) {

}
