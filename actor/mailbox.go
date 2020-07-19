package actor

type MailBox struct {
	pid *Pid
}

func NewMailBox() *MailBox {
	return &MailBox{}
}
