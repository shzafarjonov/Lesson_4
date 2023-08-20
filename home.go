package main

type Home struct {
	money, food, dirty int
	husband, wife      Human
}

type Family struct {
	home Home
}

const (
	minusSatietyPerDay        = 10
	eat                       = 30
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
)

func (h *Home) Work() {
	h.money += addMoney
	h.husband.satiety -= minusSatietyPerDay
}

func (h *Home) BuyFood() {
	h.money -= foodPriceForUnit * addFood
	h.food += addFood
	h.wife.satiety -= minusSatietyPerDay
}

func (h *Home) AddDirtyPerDay() {
	h.dirty += addDirtyPerDay
	if h.dirty > maxAllowableDirty {
		h.husband.happiness -= minesHappiness
		h.wife.happiness -= minesHappiness
	}
}

func (h *Home) Gaming() {
	h.husband.happiness += addHappinessForGaming
	h.husband.satiety -= minusSatietyPerDay
}

func (h *Home) BuyCoat() {
	h.money -= coatPrice
	h.wife.happiness += addHappinessForBuyingCoat
	h.wife.satiety -= minusSatietyPerDay
}

func (h *Home) CleanHome() {
	h.dirty -= cleanDirty
	h.wife.satiety -= minusSatietyPerDay
}
