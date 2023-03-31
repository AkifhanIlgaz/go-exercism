package purchase

import (
	"fmt"
	"strings"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	switch kind {
	case "car", "truck":
		return true

	}

	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	bestVehicle := func() string {
		if strings.Compare(option1, option2) < 0 {
			return option1
		}
		return option2
	}

	return fmt.Sprintf("%s is clearly the better choice.", bestVehicle())
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var resellPrice float64

	switch {
	case age > 3 && age < 10:
		resellPrice = applyDiscount(originalPrice, 70)
	case age < 3:
		resellPrice = applyDiscount(originalPrice, 80)
	case age >= 10:
		resellPrice = applyDiscount(originalPrice, 50)
	}

	return resellPrice
}

func applyDiscount(originalPrice, percentage float64) float64 {
	return originalPrice * percentage / 100
}
