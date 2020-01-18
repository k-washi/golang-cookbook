package globl

func UserLog() error {
	/*
		if err := Init(); err != nil {
			fmt.Println(err)
			//return err
		}
	*/

	//{"key":"value","level":"debug","msg":"hello","time":"2020-01-18T22:31:36+09:00"}
	WithField("key", "value").Debug("hello")

	//{"level":"debug","msg":"[test test2]","time":"2020-01-18T22:37:30+09:00"}
	Debug("test", "test2")
	return nil
}
