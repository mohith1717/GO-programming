// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type T struct {
// 	A string
// 	B bool
// 	C float64
// 	D int
// }

// func main() {
// 	t := T{"Mohith is", true, 17.5, 2}
// 	typeT:=reflect.TypeOf(t);

// 	for i:=0;i<typeT.NumField();i++{
// 		field:=typeT.Field(i)
// 		fmt.Println(field.Name,field.Type)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"go/types"
// 	"strconv"
// )

// func main(){
// 	var value float64 = 15
// 	value2 := "Go Programming Language"

// 	fmt.Println(Sprint(value))
// 	fmt.Println(Sprint(value2))
// 	fmt.Println(Sprint(10))
// 	fmt.Println(Sprint(true))
// }

// func Sprint(x interface{})string{
// 	switch x:=x.(type){
// 	case string:
// 		return x
// 	case int:
// 		return strconv.Itoa(x)
// 	case bool:
// 		if x {
// 			return "true"
// 		}
// 		return "false"
// 	default:
// 		// array, chan, func, map, pointer, slice, struct
// 		return "???"
// 	}
// 	}
// }

package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
	C float64
}

func main() {
	s := T{A: 20, B: "ABCD", C: 15.24}
	fmt.Println(reflect.ValueOf(&s).Elem().FieldByName("A"))
	reflect.ValueOf(&s).Elem().FieldByName("B").SetString("TEST")
	fmt.Println(s)
}
