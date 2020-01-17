package tags

import (
	"fmt"
)

type Person struct {
	Name  string `serialize:"name"`
	City  string `serialize:"city"`
	State string
	Misc  string `serialize:"-"`
	Year  int    `serialize:"year"`
}

func FullStruct() error {
	p := Person{
		Name:  "K washi",
		City:  "Settle",
		State: "WA",
		Misc:  "some fact",
		Year:  2019,
	}

	res, err := SerializeStructStrings(&p)
	if err != nil {
		return err
	}

	fmt.Printf("Full struct: %#v\n", p)
	fmt.Println("Serialize Res:", res)

	return nil
}
