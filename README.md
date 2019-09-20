# gsg_interfaces



```go
package main

import "fmt"
import "reflect"


// Here's a new custom datatype, called 'Dog', which is based on the base datatype of 'string'
type Dog string

// Here's a method that can requires our new custom datatype as an input parameter. 
func (dog Dog) Walk() string{
    return "Dog can walk"
}


func main(){
    // Here we initialise a new variable that belongs to our new datatype
    var dog Dog
    
    fmt.Println(reflect.TypeOf(dog))
    
    // then apply a method to it. 
    fmt.Println((dog.Walk()))
    
}
```
[Run this code here](https://play.golang.org/p/gcAcLTnnySp)

This outputs: 

```
main.Dog
Dog can walk
```



```go

```



```go

```






## Useful links

https://dev.to/rapulu/basic-understanding-of-golang-interface-59cn
