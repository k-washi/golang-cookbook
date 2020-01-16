package tempfiles

import (
	"fmt"
	"io/ioutil"
	"os"
)

func WorkWithTemp() error {
	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		return nil
	}

	defer os.RemoveAll(t)

	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		return err
	}

	fmt.Println(tf.Name())

	return nil
}
