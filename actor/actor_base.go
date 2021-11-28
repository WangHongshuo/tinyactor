package actor

import (
	"fmt"
	"sync/atomic"
)

const INVALID_PID = ""

const (
	INVALID = iota
	IDLE
	RUNNING
)

//
type ActorBase struct {
	pid        string
	mailbox    *MailBox
	postOffice *PostOffice
	status     uint32
}

// Init
func (a *ActorBase) Init(p *PostOffice, mailBoxSize int, actor Actor) {
	if a == nil || actor == nil {
		return
	}
	if p == nil {
		fmt.Println("PostOffice is nil, get local postoffice")
		localPostOffice := GetLocalPostOffice()
		if localPostOffice == nil {
			InitLocalPostOffice()
		}
		p = localPostOffice
	}
	a.mailbox = NewMailBox(actor, a, mailBoxSize)
	a.postOffice = p
	p.RegisterMailBox(a.mailbox)
}

// NewActor
func NewActor(pid string) *ActorBase {

	a := &ActorBase{}
	a.pid = pid
	a.status = uint32(IDLE)
	return a
}

// Send
func (a *ActorBase) SendTo(receiver string, msg interface{}) {
	if a == nil || receiver == INVALID_PID {
		return
	}
	mail := &Mail{
		Sender: a.Self(),
		Msg:    msg,
	}
	a.postOffice.SendTo(receiver, mail)
}

// Self
func (a *ActorBase) Self() string {
	if a == nil {
		return INVALID_PID
	}
	return a.pid
}

// IsRunning
func (a *ActorBase) IsRunning() bool {
	if a == nil {
		return false
	}
	return atomic.LoadUint32(&a.status) == RUNNING
}

// IsIdle
func (a *ActorBase) IsIdle() bool {
	if a == nil {
		return false
	}
	return atomic.LoadUint32(&a.status) == IDLE
}

// ChangeStatusToRunning
func (a *ActorBase) ChangeStatusToRunning() {
	if a == nil {
		return
	}
	atomic.StoreUint32(&a.status, RUNNING)
}

// ChangeStatusToIdle
func (a *ActorBase) ChangeStatusToIdle() {
	if a == nil {
		return
	}
	atomic.StoreUint32(&a.status, IDLE)
}

// Pid
func (a *ActorBase) Pid() string {
	if a == nil {
		return INVALID_PID
	}
	return a.pid
}
