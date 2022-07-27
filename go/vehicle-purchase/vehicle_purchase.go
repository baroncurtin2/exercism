package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var car string

	if option1 <= option2 {
		car = option1
	} else {
		car = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", car)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var value float64

	switch {
		case age < 3:
			value = .8
		case age > 3 && age < 10:
			value = .7
		case age >= 10:
			value = .5

	}

	return originalPrice * value
}
