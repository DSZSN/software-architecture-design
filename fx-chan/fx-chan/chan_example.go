package fx_chan

import (
    "fmt"
)

type Person struct {
    Name string
    Age  int
}

func (p *Person) SetName(name string) *Person {
    p.Name = name
    return p
}

func (p *Person) SetAge(age int) *Person {
    p.Age = age
    return p
}

func (p *Person) Print() {
    fmt.Printf("name:%s, age:%d\n", p.Name, p.Age)
}
