package domain_enums

import (
	"fmt"
)

type CommandStatus int

const (
	Pending CommandStatus = iota // 0
	Process                      // 1
	Done                         // 2
	Failed                       // 3
)

func NewCommandStatus(v int) (CommandStatus, error) {
	switch CommandStatus(v) {
	case Pending, Process, Done, Failed:
		return CommandStatus(v), nil
	default:
		return -1, fmt.Errorf("invalid CommandStatus value: %d", v)
	}
}

func (e CommandStatus) Label() string {
	switch e {
	case Pending:
		return "pending"
	case Process:
		return "process"
	case Done:
		return "done"
	case Failed:
		return "failed"
	default:
		return "unknown"
	}
}
