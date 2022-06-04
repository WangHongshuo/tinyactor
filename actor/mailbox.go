package actor

import (
	"log"
)

const defatuMailBoxSize = 10

type MailBox struct {
	mailBox chan interface{}
	ctx     ReceiverContext
}

// NewMailBox
func NewMailBox(ctx Context, size int) *MailBox {
	if ctx == nil {
		return nil
	}
	mb := &MailBox{}
	if size <= 0 {
		size = defatuMailBoxSize
	}
	mb.mailBox = make(chan interface{}, size)
	mb.ctx = ctx
	mb.schedule()
	return mb
}

func (m *MailBox) schedule() {
	log.Default().Printf("[%v]: start new goroutine\n", m.ctx.Self())
	go func() {
		for {
			msg, ok := <-m.mailBox
			if ok {
				if messageEnvelope, ok := msg.(*MessageEnvelope); ok {
					m.ctx.Receive(messageEnvelope)
				}
			} else {
				break
			}
		}
		log.Default().Printf("[%v]: goroutine end\n", m.ctx.Self())
	}()
}

// inBox
func (m *MailBox) inBox(msg interface{}) {
	if m == nil || msg == nil {
		return
	}
	select {
	case m.mailBox <- msg:
		log.Default().Printf("[%v]: MailBox in\n", m.ctx.Self())
	default:
		log.Default().Printf("[%v]: MailBox is full.\n", m.ctx.Self())
		return
	}
}
