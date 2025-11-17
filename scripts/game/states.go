package game

type State interface {
	Enter(g Game) error
	Exit(g Game) error
}

type statemachine struct {
	current State
}

func NewStateMachine() *statemachine {
	return &statemachine{}
}

func (sm *statemachine) ChangeState(g Game, newState State) error {
	if sm.current != nil {
		if err := sm.current.Exit(g); err != nil {
			return err
		}
	}
	sm.current = newState
	if sm.current != nil {
		if err := sm.current.Enter(g); err != nil {
			return err
		}
	}
	return nil
}

func (sm *statemachine) CurrentState() State {
	return sm.current
}
