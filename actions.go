/**
 * @author Guillaume Robin <robinguillaume.pro@gmail.com>
 * @file Defines Store action types and functions.
 * @desc Created on 2020-06-15 9:33:43 pm
 * @copyright GNU General Public License v3.0
 */
package margodux

// Payload defines the payload contained within an action.
type Payload interface{}

// AsyncAction defines an action that needs to run to produce an action.
type AsyncAction interface {
	Run(*Store)
}

// Action defines a synchronous action to perform in order to update the global state.
type Action struct {
	ID      string
	Payload Payload
	Err     bool
}
