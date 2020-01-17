package tags

import "reflect"

func SerializeStructStrings(s interface{}) (string, error) {
	result := ""
	r := reflect.TypeOf(s)
	value := reflect.ValueOf(s)

	//Typeがポインタなら、r, valueにポインタを格納
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
		value = value.Elem()
	}

	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		key := field.Name

		//Serializeが-ならスキップ
		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			if serialize == "-" {
				continue
			}
			key = serialize
		} else {
			//serializeがないものはスキップ
			continue
		}

		//値がstringならresultを拡張
		switch value.Field(i).Kind() {
		case reflect.String:
			result += key + ":" + value.Field(i).String() + ";"
		default:
			continue
		}
	}
	return result, nil

}
