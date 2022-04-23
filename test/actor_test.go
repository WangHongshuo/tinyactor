package test

import (
	"testing"
	"time"

	"github.com/WangHongshuo/TinyActor/test/stub"
)

func Test_Actor(t *testing.T) {
	stub.Init()
	actor1 := stub.GetTestActor1()
	actor2 := stub.GetTestActor2()
	actor1.SendTo(actor2.Pid(), &stub.DelayMsg{Duration: 1 * time.Second})
	actor1.SendTo(actor2.Pid(), &stub.DelayMsg{Duration: 2 * time.Second})
	actor1.SendTo(actor2.Pid(), "Hello, "+actor2.Pid())
	actor2.SendTo(actor1.Pid(), "Hello, "+actor1.Pid())

	time.Sleep(time.Second * 5)
}
