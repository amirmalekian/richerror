package log

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect_and_interfaces/richerror"
)

type Log struct {
	Errors []richerror.RichError
}

func (l Log) Print() {
	for i, e := range l.Errors {
		fmt.Println("i", i, "error", e)
	}
}

func (l Log) Append(r *richerror.RichError) {
	l.Errors = append(l.Errors, *r)
}

func (l Log) Save() {
	f, _ := os.OpenFile("errors.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	defer f.Close()

	data, _ := json.Marshal(l.Errors)
	f.Write(data)

}
