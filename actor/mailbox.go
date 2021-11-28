package actor

import "fmt"

const defatuMailBoxSize = 10

type MailBox struct {
	actor    *ActorBase
	mailBox  chan *Mail
	receiver Receiver
}

// NewMailBox
func NewMailBox(a Actor, ab *ActorBase, size int) *MailBox {
	if a == nil || ab == nil {
		return nil
	}
	mb := &MailBox{}
	if size <= 0 {
		size = defatuMailBoxSize
	}
	mb.actor = ab
	mb.mailBox = make(chan *Mail, size)
	mb.receiver = a
	return mb
}

// InBox
func (m *MailBox) InBox(msg *Mail) {
	if m == nil || msg == nil {
		return
	}
	select {
	case m.mailBox <- msg:
		fmt.Printf("Pid[%v]: recv msg from [%v]\n", m.Pid(), msg.Sender)
	default:
		fmt.Printf("MailBox[%v]: box is full.\n", m.Pid())
		return
	}
	if m.actor.IsIdle() {
		fmt.Printf("Pid[%v]: start new goroutine\n", m.Pid())
		m.actor.ChangeStatusToRunning()
		go func() {
			for {
				msg, ok := <-m.mailBox
				if ok {
					m.receiver.Receive(msg)
				} else {
					break
				}
			}
			fmt.Printf("Pid[%v]: goroutine end\n", m.Pid())
			m.actor.ChangeStatusToIdle()
		}()
	}
}

// Pid
func (m *MailBox) Pid() string {
	if m == nil {
		return INVALID_PID
	}
	return m.actor.Pid()
}
