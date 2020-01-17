package main

import (
	"github.com/k-washi/golang-cookbook/datatype/reflection_tags/tags"
)

/*
Full struct: tags.Person{Name:"K washi", City:"Settle", State:"WA", Misc:"some fact", Year:2019}
Serialize Res: name:K washi;city:Settle;State:WA;
*/

func main() {
	if err := tags.FullStruct(); err != nil {
		panic(err)
	}
}
