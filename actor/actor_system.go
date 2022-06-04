package actor

type ActorSystem struct {
	ProcessRegistry *ProcessRegistry
	Root            *actorContext
}

// NewActorSystem
func NewActorSystem() *ActorSystem {
	system := &ActorSystem{}
	system.ProcessRegistry = NewProcessRegistry()
	system.Root = NewActorContext(NewPID("ROOT"), system)
	system.ProcessRegistry.Register(system.Root.Pid())
	return system
}
