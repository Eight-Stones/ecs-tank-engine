package utils

// CheckBitMask check that code contains all bits from conditions.
func CheckBitMask(code int, conditions ...int) bool {
	condition := 0
	for i := range conditions {
		condition |= conditions[i]
	}
	return code&condition == condition
}
