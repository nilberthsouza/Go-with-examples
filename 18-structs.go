package main

import "fmt"

type person struct {
    name string 
    age int
}

func NewPerson(name string) *person{
    p := person{name: name}
    p.age  = 42
    return &p
}

func main(){
    fmt.Println(person{"Bob",20})
    
    fmt.Println(person{name:"Nilberth",age:21})
    
    fmt.Println(person{name:"Fred"})
    
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)

}

