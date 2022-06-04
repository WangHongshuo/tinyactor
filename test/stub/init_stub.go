package stub

import (
	"log"

	"github.com/WangHongshuo/TinyActor/actor"
)

var ActorCici *TestActor
var ActorCiciPid *actor.PID

var ActorKiki *TestActor
var ActorKikiPid *actor.PID

func init() {
	log.Default().Printf("init test stub.\n")
	ActorSystem = actor.NewActorSystem()
	ActorCiciPid, ActorCici = NewTestActor("cici")
	ActorKikiPid, ActorKiki = NewTestActor("kiki")
}

func GetTestActorCici() (*actor.PID, *TestActor) {
	return ActorCiciPid, ActorCici
}

func GetTestActorKiki() (*actor.PID, *TestActor) {
	return ActorKikiPid, ActorKiki
}
