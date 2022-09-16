package validator

func IsValidUnit(unit string) bool {
	switch unit {
	case
		"USD",
		"EUR":
		return true
	}
	return false
}
