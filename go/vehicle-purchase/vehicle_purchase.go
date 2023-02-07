package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}

	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var selectedOption string

	if option2 < option1 {
		selectedOption = option2
	} else {
		selectedOption = option1
	}

	return selectedOption + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var resaleValuePercentage float64

	if age < 3 {
		resaleValuePercentage = 80
	} else if age < 10 {
		resaleValuePercentage = 70
	} else {
		resaleValuePercentage = 50
	}

	return resaleValuePercentage / 100.0 * originalPrice
}
