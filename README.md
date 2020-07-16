# gsg_interfaces

Useful link: https://medium.com/rungo/interfaces-in-go-ab1601159b3a

So far we've seen that a box can only store one type of data.


```go
package main

import (
    "fmt"
    "reflect"
)

type Rect struct {
    width  float64
    height float64
}

type Circle struct {
    radius float64
}

func main() {

    myShape := Rect{
        width: 5,
        height: 10,
    }

    // This indicates that this box can only store variables
    // of the data type "main.Rect"
    fmt.Println(reflect.TypeOf(myShape)) // main.Rect

    // Here we tried to replace the content of this box with
    // a variable of the data type 'Circle'. As expected,
    // this ends up failing with the following error message
    myShape = Circle{
        radius: 6,
    }       // cannot use Circle literal (type Circle) as type Rect in assignment

}

```
See: https://play.golang.org/p/vA0pJmdZVxU




Now let's create a couple of methods for each of our struct types:

```go
package main

import (
    "fmt"
    "reflect"
    "math"
)

type Rect struct {
    width  float64
    height float64
}


type Circle struct {
    radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}



func main() {

    myShape := Rect{
        width: 5,
        height: 10,
    }

    fmt.Println(reflect.TypeOf(myShape)) // main.Rect


}

```
See: https://play.golang.org/p/pXx1Za0sLw0


Notice here that:

- Circle and Rect structs both have methods with the same name, i.e. "Perimeter" and "Area"
- These method's have identical parameter signatures. For example both Area methods doesn't require any input parameters and only returns a single float64 output parameter. The same is true for the perimeter methods too.


The only difference are the method receivers, one accepts a Circle struct while the other accept's a square struct, and the methods body.


Now an "interface" datatype is a datatype that let's you store other datatypes inside it.

```
type Shape interface {
	Area() float64
	Perimeter() float64
}
```

Here we're we defining a new (interface) data type. This says that we can create a box for storing Shape objects. And this box will actually store any objects of other data types, as long as those data types have a corresponding methods called  'Area' and 'Perimeter' and they have method-parameter signatures as shown above.

Now lets give this a try.

```go
package main

import (
    "fmt"
    "reflect"
    "math"
)

type Rect struct {
    width  float64
    height float64
}


type Circle struct {
    radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func main() {

    // here we're creating a new variable called 'myshape' based
    // on the interface-based datatype type
    var myShape Shape = Rect{
        width: 5,
        height: 10,
    }

    fmt.Println(reflect.TypeOf(myShape)) // main.Rect
}
```

Here we declared an interface datatype called "Shape". We also created an object from this datatype called 'myShape'. In this example, this interface object's, box is storing a Rect struct object. Now we

Note, we have to use the 'var myShape Shape =' rather than the ":=" syntax, because golang isn't smart enough to work out whether we want to create a "Shape" based object or a "Rect" object.

Golang knows that myShape's box is actually housing a variable of the Rect struct type. So when we call a method against this object, e.g.:

```go
myShape.Area()
```

It will check the the interface definition to see if "Area()" is listed, and if so it will call the corresponding method for the underlying Rect struct, with the matching paramter signature of "Area() float64".


```go
package main

import (
    "fmt"
    "reflect"
    "math"
)

type Rect struct {
    width  float64
    height float64
}


type Circle struct {
    radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func main() {

    var myShape Shape = Rect{
        width: 5,
        height: 10,
    }

    fmt.Println(reflect.TypeOf(myShape)) // main.Rect

    // this is the same as the last example but has this line.
    fmt.Println(myShape.Area()) // 50

}
```
see: https://play.golang.org/p/ooYwys0jVON




We can also replace the content of the myShape box with a variable from a different compatible datatype, for example the Circle struct datatype:


```go
package main

import (
    "fmt"
    "reflect"
    "math"
)

type Rect struct {
    width  float64
    height float64
}


type Circle struct {
    radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func main() {

    var myShape Shape = Rect{
        width: 5,
        height: 10,
    }

    fmt.Println(reflect.TypeOf(myShape)) // main.Rect
    fmt.Println(myShape.Area()) // 50

    // This bit is the new section
    myShape = Circle{
        radius: 6,
    }
    fmt.Println(reflect.TypeOf(myShape)) // main.Circle
    fmt.Println(myShape.Area()) // 113.09733552923255
}
```

Notice how the interface can store object from different types of structs! That's the power of interfaces.

Note, Rect structs can only implement the Shape interface if it has all the methods defined, as listed in interface defintion. i.e. if you delete the Rect's Perimeter() function definition then the code will fail.



























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

Notice here we have 2 methods called 'Walk'. The are similar, except for a couple of things:

- the method's body,
- input parameter datatype. One accepts the 'Dog' struct datatype and the other is 'Cat'.


the same is true for the 'Sit' functions.

To call these functions, we could have:

```go
package main

import "github.com/Sher-Chowdhury/gsg_interfaces/pets"

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

Here we have two functions, DemoDog and DemoCat. They both do exactly the same thing, but for their respective input parameter data type. It would be nicer if we could somehow create a more generic function to replace these demo* functions, e.g. something like:

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
    // Fetch()    // uncommenting because this would cause cat to fail, because a cat.Fetch() method doesn't exist.
}
// golang automatically works out that both cat and dog structs are allowed to masqaurade as an animal object, since they both meet the requirements of having corresponding Walk() and Sit() methods defined for them.



// We have replaced DemoDog and DemoCat functions
// with this single, more-generic function.
// This function accepts an interface as an input parameter.
// in otherwords, this is like saying that this function
func (animal FourLegged) Demo() {
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
