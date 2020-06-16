/**
 * @author Guillaume Robin <guillaume@inarix.com>
 * @file Implements Store tests suite.
 * @desc Created on 2020-06-15 10:09:17 pm
 * @copyright Inarix
 */
package margodux

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	store := New()
	if store == nil {
		t.Error("returned nil")
	}
	if store.reducers == nil {
		t.Error("store.reducers is nil")
	}
	if store.state == nil {
		t.Error("store.state is nil")
	}
}

func TestRegister(t *testing.T) {
	store := New()
	initialState := State{"foo": "bar"}
	store.Register("test", initialState, func(state State, action Action) State {
		switch action.ID {
		case "test":
			return State{"foo": "oof"}
		default:
			return initialState
		}
	})
	if store.state["test"]["foo"] != "bar" {
		t.Error("store.Register didn't initialize state")
	}
}

func TestGetState(t *testing.T) {
	store := New()
	initialState := State{"foo": "bar"}
	store.Register("test", initialState, func(state State, action Action) State {
		switch action.ID {
		case "test":
			return State{"foo": "oof"}
		default:
			return initialState
		}
	})
	currentState := store.GetState()
	if currentState["test"]["foo"] != "bar" {
		t.Error("store.GetState didn't return current state")
	}
	if &store.state == &currentState {
		t.Error("store.GetState didn't return a copy of store's state")
	}
}

func TestDispatch(t *testing.T) {
	store := New()
	initialState := State{"foo": "bar"}
	store.Register("test", initialState, func(state State, action Action) State {
		switch action.ID {
		case "test":
			return State{"foo": "oof"}
		default:
			return initialState
		}
	})
	store.Dispatch(Action{ID: "test", Payload: nil, Err: false})
	if store.GetState()["test"]["foo"] != "oof" {
		t.Error("store.Dispatch didn't execute action")
	}
}

type TestAsyncAction struct{}

func (t *TestAsyncAction) Run(s *Store) {
	s.Dispatch(Action{ID: "test", Payload: nil, Err: false})
}

func TestDispatchAsyncAction(t *testing.T) {
	store := New()
	initialState := State{"foo": "bar"}
	store.Register("test", initialState, func(state State, action Action) State {
		switch action.ID {
		case "test":
			return State{"foo": "oof"}
		default:
			return initialState
		}
	})
	store.Dispatch(&TestAsyncAction{})
	time.Sleep(1 * time.Second)
	if store.GetState()["test"]["foo"] != "oof" {
		t.Error("store.Dispatch didn't execute action")
	}
}

func TestDispatchPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	store := New()
	store.Dispatch("test")
}
