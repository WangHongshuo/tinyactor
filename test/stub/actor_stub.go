package stub

import (
	"fmt"
	"strings"
	"time"

	"github.com/WangHongshuo/TinyActor/actor"
)

type TestActor struct {
	a *actor.ActorBase
}

func NewTestActor(name string, po *actor.PostOffice, mailBoxSize int) *TestActor {
	if po == nil {
		return nil
	}
	ta := &TestActor{a: actor.NewActor(name)}
	ta.a.Init(po, mailBoxSize, ta)
	return ta
}

func (ta *TestActor) Receive(msg *actor.Mail) {
	if ta == nil {
		return
	}
	if msg == nil {
		fmt.Printf("Pid[%v]: nil msg", ta.Pid())
		return
	}
	m, ok := msg.Msg.(string)
	if !ok {
		fmt.Printf("Pid[%v]: unsupport msg type[%t]\n", ta.Pid(), msg.Msg)
	}
	fmt.Printf("Pid[%v]: proc msg from [%v]: %v\n", ta.Pid(), msg.Sender, m)
	if m == "wait 1s" {
		fmt.Printf("Pid[%v]: wait 1s \n", ta.Pid())
		time.Sleep(time.Second * 1)
		fmt.Printf("Pid[%v]: wait 1s ok\n", ta.Pid())
	} else if strings.Contains(m, "Hello") {
		ta.SendTo(msg.Sender, "Hi, "+msg.Sender)
	}
}

func (ta *TestActor) SendTo(receiver string, msg interface{}) {
	ta.a.SendTo(receiver, msg)
}

func (ta *TestActor) Pid() string {
	return ta.a.Pid()
}
