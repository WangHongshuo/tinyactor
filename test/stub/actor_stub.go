package stub

import (
	"fmt"
	"log"
	"time"

	"github.com/WangHongshuo/TinyActor/actor"
)

var ActorSystem *actor.ActorSystem

type TestActor struct {
	ctx     actor.Context
	Spawner actor.SpawnerContext
	Sender  actor.SenderContext
}

func NewTestActor(name string) (*actor.PID, *TestActor) {
	testActor := &TestActor{}
	pid, err := ActorSystem.Root.SpawnFromInstance(name, testActor)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return pid, testActor
}

func (ta *TestActor) Receive(ctx actor.Context) {
	if ta == nil {
		return
	}
	msg := ctx.Message()
	if msg == nil {
		log.Default().Printf("[%v]: nil msg", ctx.Self())
		return
	}
	switch m := msg.(type) {
	case *actor.Started:
		log.Default().Printf("[%v]: Actor Started!\n", ctx.Self())
		ta.ctx = ctx
		ta.Spawner = ctx
		ta.Sender = ctx
	case string:
		log.Default().Printf("[%v]: proc msg from [%v]: %v\n", ctx.Self(), ctx.Sender(), m)
	case *DelayMsg:
		log.Default().Printf("[%v]: wait %v \n", ctx.Self(), m.Duration)
		time.Sleep(m.Duration)
		log.Default().Printf("[%v]: wait %v ok\n", ctx.Self(), m.Duration)
	default:
		log.Default().Printf("[%v]: unsupport msg type[%T]\n", ctx.Self(), m)
	}
}
