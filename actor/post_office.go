package actor

import (
	"fmt"
	"sync"
)

var postOffice *PostOffice

// GetLocalPostOffice
func GetLocalPostOffice() *PostOffice {
	return postOffice
}

// InitLocalPostOffice
func InitLocalPostOffice() *PostOffice {
	if postOffice == nil {
		postOffice = NewPostOffice()
	}
	return postOffice
}

type PostOffice struct {
	mailBox map[string]*MailBox
	mutex   sync.Mutex
}

// NewPostOffice
func NewPostOffice() *PostOffice {
	return &PostOffice{mailBox: make(map[string]*MailBox)}
}

// RegisterMailBox
func (p *PostOffice) RegisterMailBox(m *MailBox) {
	if p == nil || m == nil {
		return
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if _, ok := p.mailBox[m.Pid()]; ok {
		fmt.Printf("Exist Pid[%v] in PostOffice\n", m.Pid())
		return
	}
	p.mailBox[m.Pid()] = m
}

// DeregisterMailBox
func (p *PostOffice) DeregisterMailBox(m *MailBox) {
	if p == nil || m == nil {
		return
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	delete(p.mailBox, m.Pid())
}

// GetMailBox
func (p *PostOffice) GetMailBox(pid string) *MailBox {
	if p == nil || pid == INVALID_PID {
		return nil
	}
	if m, ok := p.mailBox[pid]; ok && m != nil {
		return m
	}
	return nil
}

// SendTo
func (p *PostOffice) SendTo(receiver string, msg *Mail) {
	if p == nil || receiver == INVALID_PID {
		return
	}
	p.GetMailBox(receiver).InBox(msg)
}
