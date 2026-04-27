package main

import (
	"fmt"
	"reflect"
	"reflect_and_interfaces/richerror"
	"time"
)

func main() {
	rErr := richerror.RichError{
		Message: "id is not valid",
		MetaData: map[string]string{
			"id": "0",
		},
		Operation: "main",
		Time:      time.Now(),
	}

	value := reflect.ValueOf(rErr)
	fmt.Println("kind", value.Kind())
	fmt.Println("type", reflect.TypeOf(rErr))

	switch value.Kind() {
	case reflect.Struct:
		fmt.Println("number of field", value.NumField())
		for i := 0; i < value.NumField(); i++ {
			fieldValue := value.Field(i)
			fmt.Printf("field index: %d, type: %s, field-name: %s, value: %s\n", i, fieldValue.Type(), value.Type().Field(i).Name, fieldValue.Interface())
			//fieldType := value.Type().Field(i)

		}
	default:
		fmt.Println("kind", value.Kind())
	}

}
