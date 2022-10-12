/*
Lightweight package provide comparison function on unknown type, or interface{} in other words.

Provide ability to choose which comparison oparator basic out its string representation. E.g: "lt", "le", ">=", ...

Example 1:

package main

import (
	"encoding/json"
	"fmt"

	compare "github.com/trungqudinh/go-compare"
)

func main() {
	fmt.Println(compare.Compare(">=", 1, 2)) //false
	fmt.Println(compare.Compare("lt", 3, 4)) //true
	//fmt.Println(compare.Compare("<<", 3, 4)) //panic: Invalid operator! The parsed operator should be in ['>','>=','<','<=','=='], received ['<<']
}
*/
package compare

import (
	"encoding/json"
	"fmt"
	"strings"
)

var (
	operatorMapping = map[string]func(a, b interface{}) bool{
		">":  GreaterThan,
		">=": GreaterOrEqual,
		"<":  LessThan,
		"<=": LessOrEqual,
		"==": EqualTo,
		"!=": func(left, right interface{}) bool { return !EqualTo(left, right) },
	}
)

// Compare "left" and "right" using convert comparision "operator" string
func Compare(operator string, a, b interface{}) bool {
	f, ok := operatorMapping[operator]
	if !ok {
		keys := make([]string, 0, len(operatorMapping))
		for k := range operatorMapping {
			keys = append(keys, k)
		}
		panic(fmt.Sprintf("Invalid operator! The parsed operator should be in ['%s'], received ['%s']", strings.Join(keys, "','"), operator))
	}
	return f(a, b)
}

// Checks if values equal to each other
func EqualTo(left, right interface{}) bool {
	return isCompareTrue(equal, left, right)
}

// For '<' operator, checks if value "left" less than value "right".
func LessThan(left, right interface{}) bool {
	return isCompareTrue(lessThan, left, right)
}

// For '>' operator, checks if value "left" greater than value "right"
func GreaterThan(left, right interface{}) bool {
	return isCompareTrue(greaterThan, left, right)
}

// For '<=' operator, checks if value "left" less than or equal to value "right"
func LessOrEqual(left, right interface{}) bool {
	return isCompareTrue(lessOrEqual, left, right)
}

// For '>=' operator, checks if value "left" greater than or equal to value "right"
func GreaterOrEqual(left, right interface{}) bool {
	return isCompareTrue(greaterOrEqual, left, right)
}

// For '⊆' operator, checks if array value "right" contains value "left".
const (
	equal = iota
	lessThan
	greaterThan
	lessOrEqual
	greaterOrEqual
)

func isCompareTrue(comparison uint8, valueLeft, valueRight interface{}) bool {
	if valueLeft == nil && valueRight == nil && comparison == equal {
		return true
	}
	switch vLeft := valueLeft.(type) {
	case json.Number:
		if left, err := vLeft.Float64(); err == nil {
			switch vRight := valueRight.(type) {
			case json.Number:
				if right, err := vRight.Float64(); err == nil {
					switch comparison {
					case equal:
						if left == right {
							return true
						}
					case lessThan:
						if left < right {
							return true
						}
					case greaterThan:
						if left > right {
							return true
						}
					case lessOrEqual:
						if left <= right {
							return true
						}
					case greaterOrEqual:
						if left >= right {
							return true
						}
					}
				}
			case float32, float64, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
				right := float(valueRight)
				switch comparison {
				case equal:
					if left == right {
						return true
					}
				case lessThan:
					if left < right {
						return true
					}
				case greaterThan:
					if left > right {
						return true
					}
				case lessOrEqual:
					if left <= right {
						return true
					}
				case greaterOrEqual:
					if left >= right {
						return true
					}
				}
			}
		}
	case float32, float64, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		left := float(valueLeft)
		switch vRight := valueRight.(type) {
		case json.Number:
			if right, err := vRight.Float64(); err == nil {
				switch comparison {
				case equal:
					if left == right {
						return true
					}
				case lessThan:
					if left < right {
						return true
					}
				case greaterThan:
					if left > right {
						return true
					}
				case lessOrEqual:
					if left <= right {
						return true
					}
				case greaterOrEqual:
					if left >= right {
						return true
					}
				}
			}
		case float32, float64, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
			right := float(valueRight)
			switch comparison {
			case equal:
				if left == right {
					return true
				}
			case lessThan:
				if left < right {
					return true
				}
			case greaterThan:
				if left > right {
					return true
				}
			case lessOrEqual:
				if left <= right {
					return true
				}
			case greaterOrEqual:
				if left >= right {
					return true
				}
			}
		}
	case string:
		left := vLeft
		switch right := valueRight.(type) {
		case string:
			switch comparison {
			case equal:
				if left == right {
					return true
				}
			case lessThan:
				if left < right {
					return true
				}
			case greaterThan:
				if left > right {
					return true
				}
			case lessOrEqual:
				if left <= right {
					return true
				}
			case greaterOrEqual:
				if left >= right {
					return true
				}
			}
		}
	case bool:
		left := vLeft
		switch right := valueRight.(type) {
		case bool:
			switch comparison {
			case equal:
				if left == right {
					return true
				}
			}
		}
	}

	return false
}

func float(value interface{}) float64 {

	switch v := value.(type) {
	case float32:
		return float64(v)
	case float64:
		return float64(v)
	case int:
		return float64(v)
	case uint:
		return float64(v)
	case int8:
		return float64(v)
	case uint8:
		return float64(v)
	case int16:
		return float64(v)
	case uint16:
		return float64(v)
	case int32:
		return float64(v)
	case uint32:
		return float64(v)
	case int64:
		return float64(v)
	case uint64:
		return float64(v)
	}

	panic("never happen in that implementation")
}
