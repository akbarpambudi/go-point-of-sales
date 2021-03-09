package ptrval

func StringVal(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func Float64Val(ptr *float64) float64 {
	if ptr == nil {
		return 0
	}
	return *ptr
}

func Float32Val(ptr *float32) float32 {
	if ptr == nil {
		return 0
	}
	return *ptr
}
