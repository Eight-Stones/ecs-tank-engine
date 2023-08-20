package pkg

// CheckBitMask проверяет, что code содержит все биты из списка conditions.
func CheckBitMask(code int, conditions ...int) bool {
	condition := 0
	for i := range conditions {
		condition |= conditions[i]
	}
	return code&condition == condition
}
