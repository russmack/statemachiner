// Package statemachine provides a state machine framework.
package statemachiner

// StateMachine is the structure which represents the state machine.
type StateMachine struct {
	StartState StateFn
}

// StateMachiner is a useful interface.
type StateMachiner interface {
	Start()
}

// StateFn is a user-defined function type.  All state functions in the client
// must adhere to this signature.
type StateFn func(interface{}) StateFn

// NewStateMachine returns a new StateMachine.
func NewStateMachine() *StateMachine {
	return &StateMachine{}
}

// Start starts the StateMachine execution loop, which iterates from the StartState
// state function, through subsequent state functions, until one returns nil.
func (s *StateMachine) Start(cargo interface{}) {
	for state := s.StartState; state != nil; {
		state = state(cargo)
	}
}
