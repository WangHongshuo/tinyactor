package actor

import "sync"

type ProcessRegistry struct {
	LocalPIDs map[string]*PID
	mutex     sync.Mutex
}

// NewProcessRegistry
func NewProcessRegistry() *ProcessRegistry {
	return &ProcessRegistry{
		LocalPIDs: make(map[string]*PID, 1024),
	}
}

// Register
func (p *ProcessRegistry) Register(pid *PID) {
	if p == nil || pid == nil {
		return
	}
	p.mutex.Lock()
	p.LocalPIDs[pid.Id] = pid
	p.mutex.Unlock()
}
