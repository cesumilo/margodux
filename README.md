# margodux

[![GoDoc](https://godoc.org/github.com/cesumilo/margodux?status.svg)](https://godoc.org/github.com/cesumilo/margodux)

**Version:** v1.0.0

## About

Redux-like module for go.

## Motivation

I started this project because I wanted to make a video game using [pixel](https://github.com/faiface/pixel). However, most operating systems require all graphics and windowing code to be executed from the main thread of the program.

The main purpose of margodux is to provide an effective way of updating the game state without being limited by the frame rate. To do so, I needed a medium to execute code outside of the main thread but still being able to update the game state and the graphical interface.

That's how the idea came to me to build a redux-like module called `margodux`. I found some open-source implementation really close to what I was expecting, but I didn't found the perfect fit. That's why I decided to implement my own version.

## Table of contents

1. [Getting started](#getting-started)
2. [Test](#test)
3. [Changelog](#changelog)
4. [License](#license)

## Getting started

**Creating a new store**

```go
store := margodux.New()
```

**Creating a new reducer**

```go
initialState := margodux.State{"foo": "bar"}
store.Register("test", initialState, func(state State, action Action) State {
  switch action.ID {
  case "test":
    return State{"foo": "oof"}
  default:
    return initialState
  }
})
```

**Getting current state**

```go
currentState := store.GetState()
fmt.Println(currentState["test"]["foo"]) // bar
```

**Dispatching action**

```go
store.Dispatch(Action{ID: "test", Payload: nil, Err: false})

currentState := store.GetState()

fmt.Println(currentState["test"]["foo"]) // oof
```

**Dispatching async action**

```go
type TestAsyncAction struct{}

func (t *TestAsyncAction) Run(s *Store) {
  s.Dispatch(Action{ID: "test", Payload: nil, Err: false})
}

store.Dispatch(&TestAsyncAction{})

time.Sleep(1 * time.Second)

currentState := store.GetState()

fmt.Println(currentState["test"]["foo"]) // oof
```

## Test

**Run test**

```bash
go test
```

**Run test with coverage**

```bash
go test -cover
```

**Run test with coverage report and print**

```bash
go test -coverprofile=coverage.out
go tool cover -func=coverage.out
```

## Changelog

SEE CHANGELOG IN [CHANGELOG.md](CHANGELOG.md).

## License

SEE LICENSE IN [LICENSE](LICENSE).
