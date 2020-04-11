package main

import "github.com/Sher-Chowdhury/gsg_interfaces/pets"

// This interface represents any type
// that has both Walk and Sit methods.
type FourLegged interface {
	Walk()
	Sit()
}

// We can replace DemoDog and DemoCat
// with this single function.
func FunctionDemo(animal FourLegged) {
	animal.Walk()
	animal.Sit()
}

// Note you can't rewrite the above funciton
// in the form of a method, since you're not
// allowed to use an interface as a method receiver
// only structs are allowed as method receivers.
// func (animal FourLegged) MethodDemo() {
//  	animal.Walk()
//  	animal.Sit()
// }

func main() {
	var dog FourLegged = pets.Dog{"Fido", "Terrier"}

	cat := pets.Cat{"Fluffy", "Siamese"}

	FunctionDemo(dog)
	// Fido walks across the room
	// Fido sits down

	// Note this still works if you pass in a struct instead of a interface.
	// That's because golang is smart enough to know that that's what you intended.
	FunctionDemo(cat)
	// Fluffy walks across the room
	// Fluffy sits down

}
