package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float type")
	default:
		fmt.Println("Unknown type: ", t)
	}
}
func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)
}
