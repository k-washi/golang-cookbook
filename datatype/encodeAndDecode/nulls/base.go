package nulls

import (
	"encoding/json"
	"fmt"
)

const (
	jsonBlob     = `{"name": "KWashi"}`
	fulljsonBlob = `{"name": "KWashi", "age":26}`
)

type Example struct {
	Age  int    `json:"age,omitempty"`
	Name string `json:"name"`
}

func BaseEncodeing() error {
	e := Example{}
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Println("unmarshal no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	fmt.Println("Marshal no age:", string(value))

	return nil

}
