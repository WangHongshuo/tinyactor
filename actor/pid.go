package actor

const INVALID_PID = ""

type PID struct {
	Id      string
	mailBox *MailBox
}

// NewPid
func NewPID(id string) *PID {
	if id == INVALID_PID {
		return nil
	}
	return &PID{Id: id}
}

// String
func (p PID) String() string {
	return p.Id
}
