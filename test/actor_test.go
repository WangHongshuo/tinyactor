package test

import (
	"testing"
	"time"

	"github.com/WangHongshuo/TinyActor/actor"
	"github.com/WangHongshuo/TinyActor/test/stub"
)

func Test_Actor(t *testing.T) {
	pidCici, actorCici := stub.GetTestActorCici()
	pidKiki, actorKiki := stub.GetTestActorKiki()
	actorCici.Sender.Send(pidKiki, &stub.DelayMsg{Duration: 1 * time.Second})
	actorCici.Sender.Send(pidKiki, &stub.DelayMsg{Duration: 2 * time.Second})
	actorCici.Sender.Send(pidKiki, "Hello, "+pidKiki.Id)
	actorKiki.Sender.Send(pidCici, "Hello, "+pidCici.Id)

	actorCiciChild := actor.Actor(&stub.TestActor{})
	pidCiciChild, _ := actorCici.Spawner.SpawnFromInstance("child", actorCiciChild)
	actorKiki.Sender.Send(pidCiciChild, "Hello, "+pidCiciChild.Id)

	time.Sleep(time.Second * 5)
}
