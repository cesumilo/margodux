/**
 * @author Guillaume Robin <robinguillaume.pro@gmail.com>
 * @file Engine defines a state machine.
 * @desc Created on 2020-06-15 9:30:49 pm
 * @copyright GNU General Public License v3.0
 */
package margodux

import "fmt"

// State defines a current state of a sub part of an app.
type State map[string]interface{}

// GlobalState defines a current state of an app.
type GlobalState map[string]State

// Reducer maps actions to states.
type Reducer func(in chan *Action, out chan *GlobalState)

// Engine manages an app state.
type Engine struct {
	state   GlobalState
	actions chan *Action
	updates chan *GlobalState
}

// Init initializes Engine struct.
func (e *Engine) Init() {
	e.state = make(GlobalState)
	e.actions = make(chan *Action)
	e.updates = make(chan *GlobalState)
}

// Register run a new reducer and pass it input and output channel.
func (e *Engine) Register(key string, initialState State, reducer Reducer) {
	e.state[key] = initialState
	go reducer(e.actions, e.updates)
}

// Quit sends a signal to reducers to stop.
func (e *Engine) Quit() {
	e.Dispatch(Quit())
}

// Run start state machine engine.
func (e *Engine) Run() {
	for {
		select {
		case a := <-e.actions:
			if a.id == GoduxEngineQuit {
				return
			}
		case state := <-e.updates:
			for k, v := range *state {
				e.state[k] = v
			}
		}
	}
}

// Dispatch distributes an action to reducers.
func (e *Engine) Dispatch(action interface{}) {
	switch v := action.(type) {
	case *Action:
		e.actions <- action.(*Action)
	case AsyncAction:
		go action.(AsyncAction).Run(e.actions, &e.state)
	default:
		panic(fmt.Sprintln("unknown: ", v))
	}
}
