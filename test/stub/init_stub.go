package stub

import (
	"fmt"

	"github.com/WangHongshuo/TinyActor/actor"
)

var mockPostOffice *actor.PostOffice
var mockTestActor1 actor.Actor
var mockTestActor2 actor.Actor

func Init() {
	fmt.Println("init test stub.")
	mockPostOffice = actor.NewPostOffice()
	mockTestActor1 = NewTestActor("kiki", mockPostOffice, 10)
	mockTestActor2 = NewTestActor("cici", mockPostOffice, 10)
}

func GetTestActor1() actor.Actor {
	return mockTestActor1
}

func GetTestActor2() actor.Actor {
	return mockTestActor2
}
