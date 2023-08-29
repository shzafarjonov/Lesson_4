package main

import "fmt"

const (
	startMoney     = 100
	startFood      = 50
	startDirty     = 0
	startSatiety   = 30
	startHappiness = 100
	daysInYear     = 365
)

func main() {
	var totalFood, totalMoney, totalCoat int = 0, 0, 0
	myHome := Home{
		money: startMoney,
		food:  startFood,
		dirty: startDirty,
		husband: Human{
			name:      "Husband",
			satiety:   startSatiety,
			happiness: startHappiness,
		},
		wife: Human{
			name:      "Wife",
			satiety:   startSatiety,
			happiness: startHappiness,
		},
	}
	for i := 1; i <= daysInYear; i++ {
		fmt.Println(i, "- день")
		myHome.FamilyOneDay(&totalFood, &totalMoney, &totalCoat)
	}
	fmt.Println(totalFood, totalMoney, totalCoat)
}
