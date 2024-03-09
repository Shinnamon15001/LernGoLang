package player

import "fmt"

type IPlayer interface {
	Attack() bool
	Walking() bool
	PickUp() bool
	Throw(x, y float32) bool
}

type Warrior struct {
	Name   string
	Level  int8
	Health int
}

func (w Warrior) Attack() bool {
	fmt.Println("Attacking")
	return true
}

func (w Warrior) Walking() bool {
	fmt.Println("Walking")
	return true
}

func (w Warrior) PickUp() bool {
	fmt.Println("Picking")
	return true
}

func (w Warrior) Thow(x, y float32) bool {
	fmt.Println("")
	return true
}

type Mage struct {
	Name   string
	Level  int8
	Health int
}

func (w Mage) Attack() bool {
	fmt.Println("Attacking")
	return true
}

func (w Mage) Walking() bool {
	fmt.Println("Walking")
	return true
}

func (w Mage) PickUp() bool {
	fmt.Println("Picking")
	return true
}

func (w Mage) Thow(x, y float32) bool {
	fmt.Println("")
	return true
}
