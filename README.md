# go-compare

Lightweight package provide comparison function on unknown types, or interface{} in other words.

Provide ability to choose which comparison oparator based on its mathematical string representation. E.g: "lt", "le", ">=", ...

### Quick tour

```go
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
```

### Supported operators

| Operator | Provided function from package                    | Comparison meaning        |
|:--------:| ------------------------------------------------- | ------------------------- |
| ==       | func Equal(left, right interface{}) bool          | Equals to                 |
| !=       |                                                   | Not equals to             |
| <        | func LessThan(left, right interface{}) bool       | Less than                 |
| <=       | func LessOrEqual(left, right interface{}) bool    | Less than or equals to    |
| >        | func GreaterThan(left, right interface{}) bool    | Greater than              |
| >=       | func GreaterOrEqual(left, right interface{}) bool | Greater than or equals to |

### Supported types

| Type        | Equal/NotEqual     | Less/Greater       |
| ----------- |:------------------:|:------------------:|
| float32     | :heavy_check_mark: | :heavy_check_mark: |
| float64     | :heavy_check_mark: | :heavy_check_mark: |
| int         | :heavy_check_mark: | :heavy_check_mark: |
| uint        | :heavy_check_mark: | :heavy_check_mark: |
| int8        | :heavy_check_mark: | :heavy_check_mark: |
| uint8       | :heavy_check_mark: | :heavy_check_mark: |
| int16       | :heavy_check_mark: | :heavy_check_mark: |
| uint16      | :heavy_check_mark: | :heavy_check_mark: |
| int32       | :heavy_check_mark: | :heavy_check_mark: |
| uint32      | :heavy_check_mark: | :heavy_check_mark: |
| int64       | :heavy_check_mark: | :heavy_check_mark: |
| uint64      | :heavy_check_mark: | :heavy_check_mark: |
| string      | :heavy_check_mark: | :heavy_check_mark: |
| bool        | :heavy_check_mark: |                    |
| nil         | :heavy_check_mark: |                    |

### Implementation idea is reference from

https://github.com/openprovider/assert

https://github.com/stretchr/testify/blob/master/assert

https://cs.opensource.google/go/go/+/master:src/text/template
