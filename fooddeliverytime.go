package piscine

// burger takes 15 min,
// chips takes 10 min
// nuggets takes 12 min
type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	burger := food{15}
	chips := food{10}
	nuggets := food{12}

	switch order {
	case "burger":
		return burger.preptime
	case "chips":
		return chips.preptime
	case "nuggets":
		return nuggets.preptime
	}
	return 404
}
