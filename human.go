package main

import "fmt"

type Human struct {
	name               string
	satiety, happiness int
}

func (h *Human) Eat() {
	h.satiety += oneMeal
	fmt.Println(h.name, "покушал")
}
