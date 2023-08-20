package main

type Human struct {
	name               string
	satiety, happiness int
}

func (h *Human) Eat() {
	h.satiety += eat
}
