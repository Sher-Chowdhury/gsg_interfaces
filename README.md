# gsg_interfaces


In this demo we have a pets child package, pets/pets.go:

```go
package pets

import "fmt"

type Dog struct {
    Name  string
    Breed string
}

func (d Dog) Walk() {
    fmt.Println(d.Name, "walks across the room")
}
func (d Dog) Sit() {
    fmt.Println(d.Name, "sits down")
}
func (d Dog) Fetch() {
    fmt.Println(d.Name, "fetches a toy")
}

type Cat struct {
    Name  string
    Breed string
}

func (c Cat) Walk() {
    fmt.Println(c.Name, "walks across the room")
}
func (c Cat) Sit() {
    fmt.Println(c.Name, "sits down")
}
func (c Cat) Purr() {
    fmt.Println(c.Name, "purrs")
}

```

Notice here we have 2 methods called 'Walk'. The except for the method's body, the only other difference between them are the input parameter datatype. One accepts the 'Dog' datatype and the other is 'Cat'.



We also have the 'Sit' method with a similar situation.


To call these functions, we could have:

```go
package main

import "pets"

func DemoDog(dog pets.Dog) {
    dog.Walk()
    dog.Sit()
}

func DemoCat(cat pets.Cat) {
    cat.Walk()
    cat.Sit()
}

func main() {
    dog := pets.Dog{"Fido", "Terrier"}
    cat := pets.Cat{"Fluffy", "Siamese"}
    DemoDog(dog)    // Fido walks across the room
                    // Fido sits down

    DemoCat(cat)    // Fluffy walks across the room
                    // Fluffy sits down
}
```

Here we have two functions, DemoDog and DemoCat. They both do exactly the same thing, but for their respective input parameter data type. It would be nicer if we could somehow create a more generic function to replace these function, e.g. something like:

```go
func DemoAnimal(animal pets.$Animal) {
    $Animal.Walk()
    $Animal.Sit()
}
```

And then call it by running somthing like:

```
DemoAnimal(dog)
DemoAnimal(cat)
```

This approach would cut down with duplicate code, especially if you scale up to include more animals, rabbits, hamsters,...etc. Luckily that is possible using the 'interface' type. With interfaces we can end up with:

```go
package main

import "github.com/Sher-Chowdhury/gsg_interfaces/pets"

// Here we define an interface
type FourLegged interface {
    Walk()
    Sit()
    // Fetch()    // uncommenting this would cause cat to fail, because a cat.Fetch() method doesn't exist. 
}

// We can replace DemoDog and DemoCat
// with this single function.
func Demo(animal FourLegged) {
    animal.Walk()
    animal.Sit()
}

func main() {
    dog := pets.Dog{"Fido", "Terrier"}
    cat := pets.Cat{"Fluffy", "Siamese"}
    Demo(dog)
    Demo(cat)
}
```

This outputs:

```
Fido walks across the room
Fido sits down
Fluffy walks across the room
Fluffy sits down
```

So when we run Demo(cat), the Demo(animal FourLegged) function checks with 'type FourLegged interface' to see whether it can treat "cat" as an animal, which will only be true if a 'Walk' AND 'Sit' methods exists for the cat object, since those are the minimum requirements listed in the interface definitions. So it would fail if Fetch() is listed in the interface, even though the Demo() function doesn't actually calls the fetch() method. 


## Useful links

https://blog.teamtreehouse.com/go-interfaces-awesome

https://www.callicoder.com/golang-interfaces/