package main

import "fmt"

type Home struct {
	money, food, dirty int
	husband, wife      Human
}

const (
	minusSatietyPerDay        = 10
	oneMeal                   = 30
	addMoney                  = 150
	addDirtyPerDay            = 5
	cleanDirty                = 100
	maxAllowableDirty         = 90
	minesHappiness            = 10
	addHappinessForGaming     = 20
	addHappinessForBuyingCoat = 60
	coatPrice                 = 60
	addFood                   = 50
	foodPriceForUnit          = 1
	minHappinessBeforeDeath   = 5
	minMoneyToBuyProducts     = 50
)

func (h *Home) Work() {
	h.money += addMoney
	h.husband.satiety -= minusSatietyPerDay
	fmt.Println(h.husband.name, "сходил на работу")
}

func (h *Home) BuyFood() {
	h.money -= foodPriceForUnit * addFood
	h.food += addFood
	h.wife.satiety -= minusSatietyPerDay
	fmt.Println(h.wife.name, "купила домой продукты")
}

func (h *Home) AddDirtyPerDay() {
	h.dirty += addDirtyPerDay
	fmt.Println("количество грязи в доме увеличилось")
	if h.dirty > maxAllowableDirty {
		h.husband.happiness -= minesHappiness
		h.wife.happiness -= minesHappiness
		fmt.Println("уровень счастья уменьшился из-за грязи")
	}
}

func (h *Home) Gaming() {
	h.husband.happiness += addHappinessForGaming
	h.husband.satiety -= minusSatietyPerDay
	fmt.Println(h.husband.name, "сыграл в компьютерные игры")
}

func (h *Home) BuyCoat() {
	h.money -= coatPrice
	h.wife.happiness += addHappinessForBuyingCoat
	h.wife.satiety -= minusSatietyPerDay
	fmt.Println(h.wife.name, "купила себе шубу")
}

func (h *Home) CleanHome() {
	h.dirty -= cleanDirty
	h.wife.satiety -= minusSatietyPerDay
	fmt.Println(h.wife.name, "убралась дома")
}

func (h *Home) HusbandDayAction(totalFood *int, totalMoney *int) {
	if h.husband.satiety <= oneMeal && h.food >= oneMeal {
		h.husband.Eat()
		*totalFood += oneMeal
		return
	}
	if h.husband.happiness <= minHappinessBeforeDeath {
		h.Gaming()
		return
	}
	if h.money <= addFood*foodPriceForUnit {
		h.Work()
		*totalMoney += addMoney
		return
	}
}

func (h *Home) WifeDayAction(totalFood *int, totalCoat *int) {
	if h.wife.satiety <= oneMeal && h.food >= oneMeal {
		h.wife.Eat()
		*totalFood += oneMeal
		return
	}
	if h.dirty >= maxAllowableDirty {
		h.CleanHome()
		return
	}
	if h.food <= 2*oneMeal && h.money >= foodPriceForUnit*addFood {
		h.BuyFood()
		return
	}
	if h.wife.happiness <= minHappinessBeforeDeath {
		h.BuyCoat()
		*totalCoat += 1
		return
	}
}

func (h *Home) FamilyOneDay(totalFood *int, totalMoney *int, totalCoat *int) {
	h.HusbandDayAction(totalFood, totalMoney)
	h.WifeDayAction(totalFood, totalCoat)
	h.AddDirtyPerDay()
}
