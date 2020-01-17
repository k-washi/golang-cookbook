package checktype

import "fmt"

//CheckType type check example
func CheckType(s interface{}) {
	switch s.(type) {
	case string:
		fmt.Println("String")
	case int:
		fmt.Println("Int")
	default:
		fmt.Println("not sure")
	}
}

/*
//manually check
if val, ok := i.(string); ok {
	fmt.Println("val =", val)
}
*/
