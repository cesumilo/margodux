/**
 * @author Guillaume Robin <robinguillaume.pro@gmail.com>
 * @file Store defines a state machine.
 * @desc Created on 2020-06-15 9:30:49 pm
 * @copyright GNU General Public License v3.0
 */
package margodux

import (
	"fmt"
	"sync"
)

// State defines a current state of a sub part of an app.
type State map[string]interface{}

// GlobalState defines a current state of an app.
type GlobalState map[string]State

// Reducer maps actions to states.
type Reducer func(State, Action) State

// Store manages an app state.
type Store struct {
	mu       sync.RWMutex
	state    GlobalState
	reducers map[string]Reducer
}

// New create a new store.
func New() *Store {
	store := Store{
		state:    make(GlobalState),
		reducers: make(map[string]Reducer),
	}
	return &store
}

// Register run a new reducer and pass it input and output channel.
func (s *Store) Register(key string, initialState State, reducer Reducer) {
	defer s.mu.Unlock()
	s.mu.Lock()
	s.state[key] = initialState
	s.reducers[key] = reducer
}

// GetState returns a copy of the current state.
func (s *Store) GetState() GlobalState {
	defer s.mu.RUnlock()
	s.mu.RLock()
	return s.state
}

// Dispatch distributes an action to reducers.
func (s *Store) Dispatch(action interface{}) {
	switch v := action.(type) {
	case Action:
		defer s.mu.Unlock()
		s.mu.Lock()
		for key := range s.reducers {
			s.state[key] = s.reducers[key](s.state[key], action.(Action))
		}
	case AsyncAction:
		go action.(AsyncAction).Run(s)
	default:
		panic(fmt.Sprintln("unknown: ", v))
	}
}
