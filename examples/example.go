// This is a contrived example of using the statemachine package.
package main

import (
	"fmt"
	"github.com/russmack/statemachiner"
)

// Setup and start the state machine.
func main() {
	s := statemachiner.NewStateMachine()
	s.StartState = dispenseDogBiscuits
	c := NewHome()
	s.Start(c)
}

// Home is the machine's state structure.
type Home struct {
	DispenseDogBiscuits int    `json:"dispense_dog_biscuits"`
	TotalDogBiscuits    int    `json:"dog_biscuits"`
	Lights              bool   `json:"lights"`
	Kettle              bool   `json:"kettle"`
	VacuumRoom          string `json:"vacuum_room"`
}

// NewHome returns a new Home object.
func NewHome() *Home {
	return &Home{}
}

// dispdispenseDogBiscuits is a state function.  This function updates state, and
// returns the next state.
func dispenseDogBiscuits(cargo interface{}) statemachiner.StateFn {
	cargo.(*Home).DispenseDogBiscuits = 10
	cargo.(*Home).TotalDogBiscuits += cargo.(*Home).DispenseDogBiscuits
	fmt.Printf("%+v\n", cargo)
	return lights
}

// lights is a state function.
func lights(cargo interface{}) statemachiner.StateFn {
	cargo.(*Home).Lights = true
	fmt.Printf("%+v\n", cargo)
	return kettle
}

// kettle is a state function.
func kettle(cargo interface{}) statemachiner.StateFn {
	cargo.(*Home).Kettle = true
	fmt.Printf("%+v\n", cargo)
	return vacuumRoom
}

// vacuumRoom is a state function which returns one of two possible state functions.
func vacuumRoom(cargo interface{}) statemachiner.StateFn {
	if cargo.(*Home).TotalDogBiscuits < 20 {
		return dispenseDogBiscuits
	} else {
		cargo.(*Home).VacuumRoom = "kitchen"
	}
	fmt.Printf("%+v\n", cargo)
	return nil
}
