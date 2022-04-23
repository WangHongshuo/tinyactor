package stub

import (
	"fmt"
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
	switch msg.Msg.(type) {
	case string:
		fmt.Printf("Pid[%v]: proc msg from [%v]: %v\n", ta.Pid(), msg.Sender, msg.Msg)
	case *DelayMsg:
		duration := msg.Msg.(*DelayMsg).Duration
		fmt.Printf("Pid[%v]: wait %v \n", ta.Pid(), duration)
		time.Sleep(duration)
		fmt.Printf("Pid[%v]: wait %v ok\n", ta.Pid(), duration)
	default:
		fmt.Printf("Pid[%v]: unsupport msg type[%t]\n", ta.Pid(), msg.Msg)
	}
}

func (ta *TestActor) SendTo(receiver string, msg interface{}) {
	ta.a.SendTo(receiver, msg)
}

func (ta *TestActor) Pid() string {
	return ta.a.Pid()
}
