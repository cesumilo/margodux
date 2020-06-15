/**
 * @author Guillaume Robin <robinguillaume.pro@gmail.com>
 * @file Implements margodux.actions tests suite.
 * @desc Created on 2020-06-15 9:45:12 pm
 * @copyright GNU General Public License v3.0
 */
package margodux

import "testing"

func TestQuit(t *testing.T) {
	res := Quit()
	if res.err {
		t.Error("res.err is true")
	}
	if res.id != GoduxEngineQuit {
		t.Errorf("res.id is different than: %s", GoduxEngineQuit)
	}
}
