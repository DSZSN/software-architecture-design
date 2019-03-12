package fx_chan

import (
    "testing"
)

func TestPerson(t *testing.T)  {
    var p Person
    
    p.SetName("David").SetAge(12).Print()
}
