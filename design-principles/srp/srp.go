package srp

import(
    "fmt"
)

type Person struct{}

func (p *Person) Run() {
    fmt.Println("running")
}

func (p *Person) Jump() {
    fmt.Println("jumping")
}


type BritshPerson struct {}

func (b *BritshPerson) SpeakEnglish() {
    fmt.Println("speak english")
}


type ChinesePerson struct {}

func (c *ChinesePerson) SpeakChinese() {
   fmt.Println("speak chinese.") 
}

type Diplomat struct {
    Person
    ChinesePerson
    BritshPerson
}

func NewDiplomat() *Diplomat {
    return &Diplomat{}
}
