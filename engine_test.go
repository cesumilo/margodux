/**
 * @author Guillaume Robin <guillaume@inarix.com>
 * @file Implements Engine tests suite.
 * @desc Created on 2020-06-15 10:09:17 pm
 * @copyright Inarix
 */
package margodux

import (
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	engine := Engine{}
	engine.Init()

	go (func() {
		for action := range engine.actions {
			switch action.id {
			case GoduxEngineQuit:
				return
			default:
				t.Error("Not sending quit signal")
				return
			}
		}
	})()

	engine.Quit()
}

func TestRun(t *testing.T) {
	engine := Engine{}
	engine.Init()

	engine.Register("app", State{"foo": "bar"}, func(in chan *Action, out chan *GlobalState) {
		for action := range in {
			switch action.id {
			case GoduxEngineQuit:
				return
			default:
				t.Error("Not sending quit signal")
				return
			}
		}
	})

	go engine.Run()

	engine.Quit()
}

type AsyncTest struct {
}

func (x AsyncTest) Run(out chan *Action, state *GlobalState) {
	out <- &Action{id: "test", err: false, payload: nil}
}

func TestDispatchAsyncAction(t *testing.T) {
	engine := Engine{}
	engine.Init()

	engine.Register("app", State{"foo": "bar"}, func(in chan *Action, out chan *GlobalState) {
		for action := range in {
			switch action.id {
			case GoduxEngineQuit:
				return
			case "test":
				out <- &GlobalState{"app": State{"foo": "oof"}}
			default:
				t.Error("Not sending quit signal")
				return
			}
		}
	})

	go engine.Run()

	test := AsyncTest{}

	engine.Dispatch(test)

	time.Sleep(1 * time.Second)

	if engine.state["app"]["foo"] != "oof" {
		t.Error("Didn't update state")
	}

	engine.Quit()
}

func TestDispatchPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	engine := Engine{}
	engine.Init()

	engine.Register("app", State{"foo": "bar"}, func(in chan *Action, out chan *GlobalState) {
		for action := range in {
			switch action.id {
			case GoduxEngineQuit:
				return
			default:
				t.Error("Not sending quit signal")
				return
			}
		}
	})

	go engine.Run()

	engine.Dispatch("test")

	engine.Quit()
}
