package main

import "fmt"

type Human interface {
	Introduce() string
}

type PoliceOffier struct {
	Name string
	Age  int
}

type Firefighter struct {
	Name string
	Age  int
}

func (p *PoliceOffier) Introduce() string {
	return fmt.Sprintf("警察官: %s, %d歳\n", p.Name, p.Age)
}

func (f *Firefighter) Introduce() string {
	return fmt.Sprintf("消防士: %s, %d歳\n", f.Name, f.Age)
}

var PoliceImp Human = &PoliceOffier{}

var FfighterImp Human = &Firefighter{}

func main() {
	PoliceImp = &PoliceOffier{
		"Aくん",
		24,
	}

	FfighterImp = &Firefighter{
		"Bくん",
		30,
	}

	fmt.Print(PoliceImp.Introduce())
	fmt.Print(FfighterImp.Introduce())
}
