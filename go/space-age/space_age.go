package space

type Planet string

const (
	oneEarthYearInSeconds = 31557600
	oneEarthYearInDays    = 365.25
)

func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Earth":
		return calculateAge(seconds, 1)
	case "Mercury":
		return calculateAge(seconds, 0.2408467)
	case "Venus":
		return calculateAge(seconds, 0.61519726)
	case "Mars":
		return calculateAge(seconds, 1.8808158)
	case "Jupiter":
		return calculateAge(seconds, 11.862615)
	case "Saturn":
		return calculateAge(seconds, 29.447498)
	case "Uranus":
		return calculateAge(seconds, 84.016846)
	case "Neptune":
		return calculateAge(seconds, 164.79132)
	default:
		return -1
	}
}

func calculateAge(seconds float64, orbitalPeriod float64) float64 {
	earthYear := (seconds) / oneEarthYearInSeconds
	return earthYear / orbitalPeriod
}
