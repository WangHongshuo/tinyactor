package actor

import "testing"

func Test_Actor(t *testing.T) {
	a := NewActor()
	a.Send(nil, a.pid, nil)
	a.Receive(nil, nil)
}
