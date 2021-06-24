package utils

import "strconv"

func ParseFloatToStringWithAccuracy(floatValue float64, bits int) string {
	var stringValue = strconv.FormatFloat(floatValue, 'f', bits, 64)
	return stringValue
}
